
package main

//import nesessage packages here

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
//the struct SimpleChaincode is used for pointer receivers
type SimpleChaincode struct {
}

// ============================================================================================================================
// Main Method for the go file
// ============================================================================================================================

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}


// ============================================================================================================================
// Init resets all the things
// ============================================================================================================================

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Inside Init")
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	err := stub.PutState("hello_world", []byte(args[0]))
    	if err != nil {
        	return nil, err
    	}


	return nil, nil
}


// ============================================================================================================================
// Invoke is our entry point to invoke a chaincode function
// ============================================================================================================================



func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Inside Invoke")
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {	//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	}else if function == "write" {
        	return t.write(stub, args)
    	}
    	fmt.Println("invoke did not find func: " + function)					//error

	return nil, errors.New("Received unknown function invocation: " + function)
}


// ============================================================================================================================
// Query is our entry point for queries
// ============================================================================================================================

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Inside Query")
	fmt.Println("query is running " + function)

	// Handle different functions
	
    	if function == "read" {                            //read a variable
        	return t.read(stub, args)
    	}
    	fmt.Println("query did not find func: " + function)

    	return nil, errors.New("Received unknown function query: " + function)
}

// ============================================================================================================================
// write method changes the ledger state as it writes into the chain
// ============================================================================================================================

func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Inside Custom method Write")
	var key, value string
	var err error
	fmt.Println("running write()")

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}

	key = args[0]                            //rename for fun
	value = args[1]
	err = stub.PutState(key, []byte(value))  //write the variable into the chaincode state
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// ============================================================================================================================
// read method does not changes the ledger state as it just reads from the chain
// ============================================================================================================================

func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Inside Custom method Read")
	var key, jsonResp string
	var err error

    	if len(args) != 1 {
        	return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
   	 }

   	 key = args[0]
    	valAsbytes, err := stub.GetState(key)
    	if err != nil {
        	jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        	return nil, errors.New(jsonResp)
    	}

    	return valAsbytes, nil
}
