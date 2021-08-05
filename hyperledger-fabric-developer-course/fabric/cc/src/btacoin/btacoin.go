package main

import (
    "bytes"
    "encoding/json"
    "encoding/gob"
    "fmt"
    "strconv"
    "github.com/hyperledger/fabric-chaincode-go/shim"
    peer "github.com/hyperledger/fabric-protos-go/peer"
)

type CoinChainCode struct {

}

type btaCoin struct {
    CoinID int `json:"coinID"`
    FullName string `json:"fullName"`
    CourseName string `json:"courseName"`
    Score int `json:"score"`
}

func (Contract *CoinChainCode) Init(stub shim.ChaincodeStubInterface) peer.Response {
}

func (Contract *CoinChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
    fnType, args := stub.GetFunctionAndParameters()

    if len(args) == 0 {
        return shim.Error("Please rerun w/ an approved method ... Methods are: returnCoinBank \n, listCoinsByCourse \n,findByName \n,createCoin \n, transferCoin \n, or startLedger \n")
    }

    if fnType == "initLedger" {
        Contract.initLedger(stub)
    } else if fnType == "transferCoin" {
        Contract.transferCoin(stub, args)
    } else if fnType == "createCoin" {
        Contract.createCoin(stub, args)
    } else if fnType == "revokeCoin" {
        Contract.revokeCoin(stub, args)
    } else if fnType == "findByID" {
        Contract.findByID(stub, args)
    } else if fnType == "returnCoinBank" {
        Contract.returnCoinBank(stub)
    } else if fnType == "findByAttributes" {
        Contract.findByAttributes(stub, args)
    } else if fnType == "showCoinsAbovePercentage" {
        Contract.showCoinsAbovePercentage(stub, args)
    } else {
        return shim.Error("Please rerun w/ an approved method ... Methods are: returnCoinBank \n,listCoinsByCourse \n,findByName \n,createCoin \n, transferCoin \n, or startLedger \n")
    }

    return shim.Success(nil)
}

func (Contract *CoinChainCode) createCoin(stub shim.ChaincodeStubInterface, args, []string) peer.Response
{
    params := stub.GetStringArgs()
    fmt.Println("The params passed in are:", params)

    if len(params) != 3 {
        fmt.Println("Please resubmit in this particular order: createCoin, COINID, Course Name")

        return shim.Error("Please resubmit in this particular order: createCoin, COINID, Course Name")
    }

    coinUID := params[1]
    coinuid, err := strconv.Atoi(coinUID)

    if err != nil {
        return shim.Error("Something wentn wrong!" + err.Error()
    }

    var coinDetails = btaCoin(CoinID: coinuid, FullName: "Unassigned", CourseName: params[2], Score: 0)

    oldKeyVal, err := stub.GetState(coinUID)

    if oldKeyVal != nil {
        fmt.Println("There's a coin with the same ID here. Use a different ID. Error Code: " + err.Error())
    }

    newCoin, err := json.Marshal(coinDetails)

    if err != nil {
        shim.Error("Something went wrong in serializing. Error: " + err.Error())
    } else {
        err = stub.PutState(args[0], newCoin)

        if err != nil {
            shim.Error("Something went wrong while adding the coin. Error: " + err.Error())
        }
    }
    return shim.Success(nil)

}

// func (Contract *CoinChainCode) transferCoin (stub shim.ChaincodeStubInterface, args []string) peer.Response {
//     args = stub.GetStringArgs()
//     if len(args) != 5 {
//         fmt.Println("There is an error... Please rerun with 4 arguments")
//         return shim.Error("Incorrect number of arguments. Expecting 4")
//     }
// }

func main() {
    err := shim.Start (new(CoinChainCode))
    if err != nil {
        fmt.Printf("Chaincode could not be started: %s", err)
    }
}




