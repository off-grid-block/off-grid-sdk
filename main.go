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

	//take input values from the user
	channelname := userInput("Enter the channel name: ")
	fmt.Printf("Channel is named %s\n", channelname)

	channelstatus := userInput("Press 1 to join an existing channel (new channel is created by default): ")
	fmt.Printf("Choice recorded: %s\n", channelstatus)

	chaincodename := userInput("Enter the chaincode name: ")
	fmt.Printf("ChainCodeID is named %s\n", chaincodename)	

	chaincodestatus :=  userInput("Press 1 to avoid installing a chaincode (new chaincode added by default): ")
	fmt.Printf("Choice recorded: %s\n", chaincodestatus)
	
	// Definition of the Fabric SDK properties
	ms := blockchainSDK.SetupSDK {
		// Network parameters
		OrdererID: "orderer.example.com",

		// Channel parameters
		//ChannelID: "mychannel",    
		ChannelID: channelname,
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/hyperledger/codesdk/first-network/channel-artifacts/channel.tx",

		// Chaincode parameters
		//ChainCodeID:     "ourcode",
		ChainCodeID:     chaincodename,
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/hyperledger/codesdk/chaincode/chaincode_v1",
		OrgAdmin:        "Admin",
		OrgName:         "org1",
		ConfigFile:      "config.yaml",//"first-network/connection-org1.yaml",

		// User parameters
		UserName: "User1",
	}

	// Initialization of the SDK 
	err := ms.Initialization(channelstatus)
	if err != nil {
		fmt.Printf("SDK innitialization failed: %v\n", err)
		return
	}
	// Close SDK
	defer ms.CloseSDK()

	// Installation and instantiation of the chaincode
	if (chaincodestatus != "1") {
		err = ms.ChainCodeInstallationInstantiation()
		if err != nil {
			fmt.Printf("Installation/Innitialization of SDK failed: %v\n", err)
			return
		}
	}

	// setup channel client and client registration
	err =ms.ClientSetup()
	if err!= nil {
		fmt.Printf("Channel client or Event client installation failed")
		return
	}

	fmt.Printf("Done ****Innitiating \n")

	looper := 0
	for (looper<1) {
		//select action
		action := userInput("Enter which action to take (options are: invoke query delete searchbyowner quit) :")
		fmt.Printf("selected %s\n", action)
		//select query
		if (action  == "query"){
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

			ID 		:= userInput("Enter ID:")
			Hash 		:= userInput("Enter Hash:")
			Application 	:= userInput("Enter application name:")
			NodeIP 		:= userInput("Enter node IP:")
			Owner 		:= userInput("Enter Owner name:")
			Updated 	:= userInput ("Enter update status (0/1)")
			
			response, err := ms.InitEntrySDK(ID, Hash, Application, NodeIP, Owner, Updated)
			if err !=nil {
				fmt.Printf("unable to invoke: %v\n", err)
				//return
			}
			fmt.Printf("response is %s\n", response)
		}
		//select delete
		if (action == "delete"){
			ID := userInput("Enter ID")
			response, err := ms.DeleteEntrySDK(ID)
			if err != nil {
				fmt.Printf("unable to delete: %v\n", err)
			}
			fmt.Print("response is %s\n", response)
		}	
		//select searchbyowner
		if (action  == "searchbyowner"){
			Owner := userInput("Enter Owner:") 
			response, err := ms.SearchByOwnerSDK(Owner)
			if err !=nil {
				fmt.Printf("unable to query: %v\n", err)
				//return
			}
			fmt.Printf("response is %s\n", response)
		} 
		//select to end
		if (action =="quit"){
			looper = looper+1
		}
	}

}	
