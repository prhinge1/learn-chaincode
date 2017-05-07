/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	err	:=stub.PutState("Idea_Products",[]byte(args[0]))
	if err !=nil {
	return nil,err
	}

	err1 :=stub.PutState("Distributer",[]byte(args[1]))
	if err1 !=nil {
	return nil,err1
	}

	err2	:=stub.PutState("Retailer",[]byte(args[2]))
	if err2 !=nil {
	return nil,err2
	}
	return nil, nil
}

// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running =====>" + function)

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	}else if function == "write" {
		return t.write(stub,args)
	}
	fmt.Println("invoke did not find func: " + function)					//error

	return nil, errors.New("Received unknown function invocation: " + function)
}

//Write is custome function inserted for sample
func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error){
	
	//var srcParty,,targBalstr,srcBalstr,strsrcBal  string
	var srcParty,trgParty,strsrcBal,strtargBal,targBalstr  string
   var transferAmt,srcBal,targBal int
   
   
   var bytesinfo []byte
   var bytestarbal []byte
      
    //var err,err3,err4,err5,err6,err7,err8,err1 error 
	var err3,err4,err1,err2,err5 error
    
    if len(args) != 3 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the variable and value to set")
    }
	//fmt.Println(" Argument--------> "+ len(args))
    srcParty = args[0]
	trgParty = args[1]                           //rename for fun
    transferAmt, err3 = strconv.Atoi(args[2])
	if err3 !=nil{
		fmt.Println(err3)
	}
	bytesinfo, err1 =stub.GetState(srcParty)
	if err1 !=nil{
		fmt.Println(err1)
	}
	strsrcBal =string(bytesinfo)
	srcBal, err4 = strconv.Atoi(strsrcBal)
	if err4 !=nil{
		fmt.Println(err4)
	}
	
	bytestarbal, err2 =stub.GetState(trgParty)
	if err2 !=nil{
		fmt.Println(err2)
	}
	strtargBal=string(bytestarbal)
	targBal, err5 = strconv.Atoi(strtargBal)
	if err5 !=nil{
		fmt.Println(err5)
	}
	
	if(transferAmt<srcBal){
		targBal=targBal + transferAmt
		srcBal=srcBal-transferAmt
		targBalstr=string(targBal)
		srcBalstr:=string(srcBal)

		err := stub.PutState(trgParty, []byte(targBalstr))  //write the variable into the chaincode state
		if err != nil {
			return nil, err
		}

		err1 := stub.PutState(srcParty, []byte(srcBalstr))  //write the variable into the chaincode state
		if err1 != nil {
			return nil, err1
		}
	
	}

  return nil, nil


}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" {											//read a variable
		return t.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)						//error

	return nil, errors.New("Received unknown function query: " + function)
}


//read function newly added
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var name, jsonResp string
    var err error

    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the var to query")
    }

    name = args[0]
    valAsbytes, err := stub.GetState(name)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + name + "\"}"
        return nil, errors.New(jsonResp)
    }

    return valAsbytes, nil
}