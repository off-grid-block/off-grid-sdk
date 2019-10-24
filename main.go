package main
import (
	"fmt"
	"codesdk/blockchainSDK"
	"os"
	"bufio"
	"strings"
)

//function to read string
func userInput(request string) (inputval string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(request)
	inputval, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}
	returnval := strings.TrimRight (inputval, "\n")
	return returnval
}


func main() {

	// Definition of the Fabric SDK properties
	ms := blockchainSDK.SetupSDK {
		// Network parameters
		OrdererID: "orderer.example.com",

		// Channel parameters
		//ChannelID: "mychannel",    
		ChannelID: "",
		ChannelConfig: "",

		// Chaincode parameters
		ChainCodeID:     "",
		//ChainCodeID:     chaincodename,
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/hyperledger/codesdk/chaincode/chaincode_v1",
		OrgAdmin:        "Admin",
		OrgName:         "org1",
		ConfigFile:      "config.yaml",//"first-network/connection-org1.yaml",

		// User parameters
		UserName: "User1",
	}


	// Initialization of the SDK 
	err := ms.Initialization()
	if err != nil {
		fmt.Printf("SDK innitialization failed: %v\n", err)
		return
	}
	// Close SDK
	defer ms.CloseSDK()


	looper := 0
	for (looper<1) {
		//select action
		action := userInput("Enter which action to take (options are: addcc addchannel invoke query delete searchbyowner quit) :")
		fmt.Printf("selected %s\n", action)
		//select query
		if (action  == "query"){
			// setup channel client and client registration
			err =ms.ClientSetup()
			if err!= nil {
				fmt.Printf("Channel client or Event client installation failed")
				return
			}
			ID := userInput("Enter ID:")
			response, err := ms.ReadEntrySDK(ID)
			if err !=nil {
				fmt.Printf("unable to query: %v\n", err)
				//return
			}
			fmt.Printf("response is %s\n", response)
		}
		//select invoke
		if (action  == "invoke"){

			ID		:= userInput("Enter ID:")
			Hash		:= userInput("Enter Hash:")
			Application	:= userInput("Enter application name:")
			NodeIP		:= userInput("Enter node IP:")
			Owner		:= userInput("Enter Owner name:")
			Updated		:= userInput("Enter update status (0/1)")

			// setup channel client and client registration
			err =ms.ClientSetup()
			if err!= nil {
				fmt.Printf("Channel client or Event client installation failed")
				return
			}
			response, err := ms.InitEntrySDK(ID, Hash, Application, NodeIP, Owner, Updated)
			if err !=nil {
				fmt.Printf("unable to invoke: %v\n", err)
				//return
			}
			fmt.Printf("response is %s\n", response)
		}
		//select delete
		if (action == "delete"){
			// setup channel client and client registration
			err =ms.ClientSetup()
			if err!= nil {
				fmt.Printf("Channel client or Event client installation failed")
				return
			}
			ID := userInput("Enter ID")
			response, err := ms.DeleteEntrySDK(ID)
			if err != nil {
				fmt.Printf("unable to delete: %v\n", err)
			}
			fmt.Print("response is %s\n", response)
		}

		//select searchbyowner
		if (action  == "searchbyowner"){
			// setup channel client and client registration
			err =ms.ClientSetup()
			if err!= nil {
				fmt.Printf("Channel client or Event client installation failed")
				return
			}
			Owner := userInput("Enter Owner:")
			response, err := ms.SearchByOwnerSDK(Owner)
			if err !=nil {
				fmt.Printf("unable to query: %v\n", err)
				//return
			}
			fmt.Printf("response is %s\n", response)
		}

		//add new chaincode
		if (action == "addcc") {

			ms.ChainCodeID = userInput("Enter the chaincode name: ")
			fmt.Printf("ChainCodeID is named %s\n", ms.ChainCodeID)
			// Installation and instantiation of the chaincode
			err = ms.ChainCodeInstallationInstantiation()
			if err != nil {
				fmt.Printf("Installation/Innitialization of chaincode: %v\n", err)
				return
			}
		}

		//add new channel
		if (action == "addchannel") {

			ms.ChannelID = userInput("Enter the channel name: ")
			fmt.Printf("ChannelID is named %s\n", ms.ChannelID)
			// Installation and instantiation of the chaincode
			err = ms.ChannelSetup()
			if err != nil {
				fmt.Printf("Installation of channel failed %v\n", err)
				return
			}
		}

		//select to end
		if (action =="quit"){
			looper = looper+1
		}
	}

}
