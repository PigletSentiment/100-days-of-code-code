// SPDX-License-Identifier: CC-By-1.0
// Creative Commons Attribution 1.0 Generic

// pragma solidity ^0.7.0;
pragma solidity ^0.8.0;

contract Utility {
    address owner;
    event errorEvent(string errorEvent_Description);
    modifier onlyOwner {
        if (msg.sender != owner) {
            emit errorEvent("Only the owner of this contract instance can perform this function!");
        } else {
            _;
        }
    }
}

contract Album is Utility {
    
    struct musicAlbum {
        string artist;
        string albumTitle;
        uint tracks;    
    }
    musicAlbum public currentAlbum;
    mapping(address => musicAlbum) public userAlbums;
    
    string public constant contractAuthor = 'Andrew Park';

    // Event which will be raised anytime the current album information is updated.
    event albumEvent(string albumEvent_Description, string albumEvent_Artist, string albumEvent_Title, uint albumEvent_Tracks);

    constructor() {
        currentAlbum.artist = 'Nirvana';
        currentAlbum.albumTitle = 'Nevermind';
        currentAlbum.tracks = 13;
        
        owner = msg.sender;
    }

    function getCurrentAlbum() public view returns (string memory, string memory, uint) {
        return (currentAlbum.artist, currentAlbum.albumTitle, currentAlbum.tracks);
    }
    
    function setCurrentAlbum(string memory _artist, string memory _albumTitle, uint _tracks) onlyOwner public {
        currentAlbum.artist = _artist;
        currentAlbum.albumTitle = _albumTitle;
        currentAlbum.tracks = _tracks;
        
        // Raise the albumEvent to let any subscribers know the current album information has changed.
        emit albumEvent("The current album information has been updated", _artist, _albumTitle, _tracks);
    }
    
    function getUsersFavoriteAlbum() public view returns (string memory, string memory, uint) {
        return (userAlbums[msg.sender].artist, userAlbums[msg.sender].albumTitle, userAlbums[msg.sender].tracks);
    }
    
    function setUsersFavoriteAlbum(string memory _artist, string memory _albumTitle, uint _tracks) public {
        userAlbums[msg.sender].artist = _artist;
        userAlbums[msg.sender].albumTitle = _albumTitle;
        userAlbums[msg.sender].tracks = _tracks;
        
        emit albumEvent("You have updated your personal favorite album information.", _artist, _albumTitle, _tracks);
    }
    
    
    
} // Album

