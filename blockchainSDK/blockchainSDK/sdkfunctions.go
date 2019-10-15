package blockchainSDK

import(
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"time"
)

// add entry using SDK
func (s *SetupSDK) InitEntrySDK(ID string, Hash string, Application string, NodeIP string, Owner string, Updated string) (string, error) {
        
	eventID := "initEvent"

        // register chaincode event
        registered, notifier, err := s.event.RegisterChaincodeEvent(s.ChainCodeID, eventID)
        if err != nil {
                return "", err
        }
        defer s.event.Unregister(registered)

        // Create a request for entry init and send it
        response, err := s.client.Execute(channel.Request{ChaincodeID: s.ChainCodeID, Fcn: "initEntry", Args: [][]byte{[]byte(ID), []byte(Hash), []byte(Application), []byte(NodeIP), []byte(Owner), []byte(Updated) }})
        if err != nil {
                return "", fmt.Errorf("failed to innitiate: %v", err)
        }

        // Wait for the result of the submission
        var ccEvent *fab.CCEvent
        select {
        case ccEvent = <-notifier:
                fmt.Printf("Received CC event: %v\n", ccEvent)
        case <-time.After(time.Second * 20):
                return "", fmt.Errorf("did NOT receive CC event for eventId(%s)", eventID)
        }

        return string(response.Payload), nil
}

//read entry on chaincode using SDK
func (s *SetupSDK) ReadEntrySDK(ID string) (string, error) {

	//creat and send request for reading an entry
        response, err := s.client.Query(channel.Request{ChaincodeID: s.ChainCodeID, Fcn: "readEntry",  Args: [][]byte{[]byte(ID)}})
        if err != nil {
                return "", fmt.Errorf("failed to query: %v", err)
        }

        return string(response.Payload), nil
}

//delete entry on chaincode using SDK
func (s *SetupSDK) DeleteEntrySDK(ID string) (string, error) {

	//register event
	eventID := "deleteEvent"
	reg, notifier, err := s.event.RegisterChaincodeEvent(s.ChainCodeID, eventID)
	if err != nil {
		return "", err
	}
	defer s.event.Unregister(reg)

	//create a request for deletion and sent it
	resp, err := s.client.Execute(channel.Request{ChaincodeID: s.ChainCodeID, Fcn: "deleteEntry", Args: [][]byte{[]byte(ID)} })
	if err != nil {
		return "", fmt.Errorf("failed to delete: %v",err)
	}

	// Wait for the result of the submission
        var ccEvent *fab.CCEvent
        select {
        case ccEvent = <-notifier:
                fmt.Printf("Received CC event: %v\n", ccEvent)
        case <-time.After(time.Second * 20):
                return "", fmt.Errorf("did NOT receive CC event for eventId(%s)", eventID)
        }

	return string(resp.Payload), nil
}

//search by username on chaincode using SDK
func (s *SetupSDK) SearchByOwnerSDK(Owner string) (string, error) {

	//creat and send request for reading an entry
        response, err := s.client.Query(channel.Request{ChaincodeID: s.ChainCodeID, Fcn: "searchByOwner",  Args: [][]byte{[]byte(Owner)}})
        if err != nil {
                return "", fmt.Errorf("failed to query: %v", err)
        }

        return string(response.Payload), nil
}

