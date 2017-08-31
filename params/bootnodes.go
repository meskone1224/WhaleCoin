// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
  "enode://173d88a03853011bd7505b205f40cc144646bbf843f93682457447c4e19407d9f2389a8657ce5f4eceede4076d4bf329a3630b554a6f353ae8d397b150b8c5e6@34.231.64.228:30373",

}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
/*
{
  enode: "enode://cd15a411c96b2dceb43a7365247bfb6d0017b3b04f45498804e1b8ea3c27c8b1329cf73a5a87b0fad71bf6cc84ddf59b19213980b71c81181d3400259e372ef5@[::]:30373?discport=0",
  id: "cd15a411c96b2dceb43a7365247bfb6d0017b3b04f45498804e1b8ea3c27c8b1329cf73a5a87b0fad71bf6cc84ddf59b19213980b71c81181d3400259e372ef5",
  ip: "::",
  listenAddr: "[::]:30373",
  name: "gwhale/v1.6.7-stable-26c696fc/linux-amd64/go1.8.3",
  ports: {
    discovery: 0,
    listener: 30373
  },
  protocols: {
    eth: {
      difficulty: 711726653,
      genesis: "0xb3756bec461cefd484dea6fbb115fc2db0cd1c06fae53c39cfc0bff9f3f40bf7",
      head: "0x1ecb7960ebb203981ba38331078e14e2a642626b9e0b81622be918f363328086",
      network: 1211
    }
  }
}
*/
var TestnetBootnodes = []string{

}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{

}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{

}
