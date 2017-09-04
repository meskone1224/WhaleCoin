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
// the mian WhaleCoin network.
var MainnetBootnodes = []string{

  "enode://173d88a03853011bd7505b205f40cc144646bbf843f93682457447c4e19407d9f2389a8657ce5f4eceede4076d4bf329a3630b554a6f353ae8d397b150b8c5e6@34.231.64.228:30371",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// WhaleCoin test network.

var TestnetBootnodes = []string{

	"enode://70c31da0ebef98f40b50ea5e27c067ad22a671e04c7b2df508c0d3cdfe35aab4eed8b8dd67ba5a733fbc79644f026894a0bfe35859c79e3e1b8433ceede86e48@34.231.48.74:30373",
	// "enode://e6a4a7b8e6f0adf290af9c7a979c3bd332bf638e57a4b3d2a641cbafd4717f8e442e65ffac711e6ae5e84eae3161dc824fd407685ab21aa2a7957eb1f98800c8@34.192.101.147:30373",

}



// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{

	"enode://b6efb6c953df1c90fd727e9970e109b90a471953dd5e67843b6492c0315d7665346ccaeea5099293e5efe77bd9f49d921be8c4516c4de4f5910a797045f35a77@34.231.48.74:30373",
	// "enode://271d3c7b3422e5447df4d6b8478b6c913c41bbd9490e34b92cdf7910b958ec938001ba2d60955f53812f66206fbedd0b9daea1c44621ddfb7df25e24dfaba51c@34.192.101.147:30373",
}

