/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
	"errors"
	"fmt"
	"strconv"
	"encoding/json"
	"time"
	"strings"

	"github.com/openblockchain/obc-peer/openchain/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

var tradeIndexStr = "_tradeindex"				//name for the key/value that will store a list of all known trades

type Trade struct {
	TradeDate string `json:"tradedate"`
	ValueDate string `json:"valuedate"`
	Operation string `json:"operation"`
	Quantity int `json:"quantity"`
	Security string `json:"security"`
	Price string `json:"price"`
	Counterparty string `json:"counterparty"`
	User string `json:"user"`
	Timestamp int64 `json:"timestamp"`			// utc timestamp of creation
	Settled int `json:"settled"`				// enriched & settled
	NeedsRevision int `json:"needsrevision"`	// returned to client for revision
}

// ============================================================================================================================
// Init - reset all the things
// ============================================================================================================================
func (t *SimpleChaincode) init(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var Aval int
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	// Initialize the chaincode
	Aval, err = strconv.Atoi(args[0])
	if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
	}

	// Write the state to the ledger
	err = stub.PutState("abc", []byte(strconv.Itoa(Aval)))				//making a test var "abc", I find it handy to read/write to it right away to test the network
	if err != nil {
		return nil, err
	}
	
	var empty []string
	jsonAsBytes, _ := json.Marshal(empty)								//marshal an emtpy array of strings to clear the index
	err = stub.PutState(tradeIndexStr, jsonAsBytes)
	if err != nil {
		return nil, err
	}
	
	return nil, nil
}

// ============================================================================================================================
// Run - Our entry point
// ============================================================================================================================
func (t *SimpleChaincode) Run(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {

	fmt.Println("ruslan: run is running " + function)

	// Handle different functions
	if function == "init" {												// initialize the chaincode state, used as reset
		return t.init(stub, args)
	} else if function == "init_trade" {									// create a new trade
		return t.init_trade(stub, args)
	} else if function == "submit_for_enrichment" {							// submit for enrichment
		return t.submit_for_enrichment(stub, args)
	} else if function == "mark_revision_needed" {						
		return t.mark_revision_needed(stub, args)
	}

	// else if function == "write" {										// writes a value to the chaincode state
	// 	return t.Write(stub, args)
	// } 

	fmt.Println("run did not find func: " + function)						// error

	return nil, errors.New("Received unknown function invocation")
}

// ============================================================================================================================
// Query - read a variable from chaincode state - (aka read)
// ============================================================================================================================
func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	if function != "query" {
		return nil, errors.New("Invalid query function name. Expecting \"query\"")
	}
	var security, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting security of the trade to query")
	}

	security = args[0]
	valAsbytes, err := stub.GetState(security)									//get the var from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for security " + security + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil													//send it onward
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// ============================================================================================================================
// Write - write variable into chaincode state
// ============================================================================================================================
// func (t *SimpleChaincode) Write(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
// 	var security, value string // Entities
// 	var err error
// 	fmt.Println("running write()")

// 	if len(args) != 2 {
// 		return nil, errors.New("Incorrect number of arguments. Expecting 2. security of the variable and value to set")
// 	}

// 	security = args[0]															//rename for funsies
// 	value = args[1]
// 	err = stub.PutState(security, []byte(value))								//write the variable into the chaincode state
// 	if err != nil {
// 		return nil, err
// 	}
// 	return nil, nil
// }

// ============================================================================================================================
// Init Trade - create a new trade, store into chaincode state
// ============================================================================================================================
func (t *SimpleChaincode) init_trade(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var err error

	if len(args) != 11 {
		return nil, errors.New("Incorrect number of arguments. Expecting 11")
	}

	fmt.Println("- start init trade")
	if len(args[0]) <= 0 {
		return nil, errors.New("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return nil, errors.New("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return nil, errors.New("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return nil, errors.New("4th argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return nil, errors.New("5th argument must be a non-empty string")
	}
	if len(args[5]) <= 0 {
		return nil, errors.New("6th argument must be a non-empty string")
	}
	if len(args[6]) <= 0 {
		return nil, errors.New("7th argument must be a non-empty string")
	}
	if len(args[7]) <= 0 {
		return nil, errors.New("8th argument must be a non-empty string")
	}
	if len(args[8]) <= 0 {
		return nil, errors.New("9th argument must be a non-empty string")
	}
	if len(args[9]) <= 0 {
		return nil, errors.New("10th argument must be a non-empty string")
	}
	if len(args[10]) <= 0 {
		return nil, errors.New("11th argument must be a non-empty string")
	}
	
	// TradeDate string `json:"tradedate"`
	// ValueDate string `json:"valuedate"`
	// Operation string `json:"operation"`
	// Quantity int `json:"quantity"`
	// Security string `json:"security"`
	// Price string `json:"price"`
	// Counterparty string `json:"counterparty"`
	// User string `json:"user"`
	// Timestamp int64 `json:"timestamp"`			// utc timestamp of creation
	// Settled int `json:"settled"`				// enriched & settled
	// NeedsRevision int `json:"needsrevision"`	// returned to client for revision

	tradedate := strings.ToLower(args[0])
	valuedate := strings.ToLower(args[1])
	operation := strings.ToLower(args[2])

	quantity, err := strconv.Atoi(args[3])

	if err != nil {
		return nil, errors.New("4th argument must be a numeric string")
	}

	security := strings.ToLower(args[4])
	price := strings.ToLower(args[5])
	counterparty := strings.ToLower(args[6])
	user := strings.ToLower(args[7])

	timestamp := makeTimestamp()

	settled, err := strconv.Atoi(args[9])

	if err != nil {
		return nil, errors.New("10th argument must be a numeric string")
	}

	needsrevision, err := strconv.Atoi(args[10])

	if err != nil {
		return nil, errors.New("11th argument must be a numeric string")
	}

	str := `{"tradedate": "` + tradedate + `", "valuedate": "` + valuedate + `", "operation": "` + operation + `", "quantity": ` + quantity + `, "security": "` + security + `", "price": "` + price + `", "counterparty": "` + counterparty + `", "user": "` + user + `", "timestamp": "` + timestamp + `", "settled": "` + settled + `", "needsrevision": "` + needsrevision + `"}`

	err = stub.PutState(args[0], []byte(str))								//store marble with id as key

	if err != nil {
		return nil, err
	}
	
	//get the trade index
	tradesAsBytes, err := stub.GetState(tradeIndexStr)

	if err != nil {
		return nil, errors.New("Failed to get trade index")
	}

	var tradeIndex []string

	json.Unmarshal(tradesAsBytes, &tradeIndex)							//un stringify it aka JSON.parse()
	
	//append
	tradeIndex = append(tradeIndex, args[8])								//add trade timestamp to index list
	fmt.Println("! trade index: ", tradeIndex)
	jsonAsBytes, _ := json.Marshal(tradeIndex)
	err = stub.PutState(tradeIndexStr, jsonAsBytes)						//store timestamp of trade

	fmt.Println("- end init trade")
	return nil, nil
	
}

// ============================================================================================================================
// Set User Permission on Trade
// ============================================================================================================================
func (t *SimpleChaincode) set_user(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var err error
	
	//   0       1
	// "name", "bob"
	if len(args) < 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}
	
	fmt.Println("- start set user")
	fmt.Println(args[0] + " - " + args[1])
	marbleAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return nil, errors.New("Failed to get thing")
	}
	res := Marble{}
	json.Unmarshal(marbleAsBytes, &res)										//un stringify it aka JSON.parse()
	res.User = args[1]														//change the user
	
	jsonAsBytes, _ := json.Marshal(res)
	err = stub.PutState(args[0], jsonAsBytes)								//rewrite the marble with id as key
	if err != nil {
		return nil, err
	}
	
	fmt.Println("- end set user")
	return nil, nil
}

// ============================================================================================================================
// Open Trade - create an open trade for a marble you want with marbles you have 
// ============================================================================================================================
// func (t *SimpleChaincode) open_trade(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
// 	var err error
// 	var will_size int
// 	var trade_away Description
	
// 	//	0        1      2     3      4      5       6
// 	//["bob", "blue", "16", "red", "16"] *"blue", "35*
// 	if len(args) < 5 {
// 		return nil, errors.New("Incorrect number of arguments. Expecting like 5?")
// 	}
// 	if len(args)%2 == 0{
// 		return nil, errors.New("Incorrect number of arguments. Expecting an odd number")
// 	}

// 	size1, err := strconv.Atoi(args[2])
// 	if err != nil {
// 		return nil, errors.New("3rd argument must be a numeric string")
// 	}

// 	open := AnOpenTrade{}
// 	open.User = args[0]
// 	open.Timestamp = makeTimestamp()											//use timestamp as an ID
// 	open.Want.Color = args[1]
// 	open.Want.Size =  size1
// 	fmt.Println("- start open trade")
// 	jsonAsBytes, _ := json.Marshal(open)
// 	err = stub.PutState("_debug1", jsonAsBytes)

// 	for i:=3; i < len(args); i++ {												//create and append each willing trade
// 		will_size, err = strconv.Atoi(args[i + 1])
// 		if err != nil {
// 			msg := "is not a numeric string " + args[i + 1]
// 			fmt.Println(msg)
// 			return nil, errors.New(msg)
// 		}
		
// 		trade_away = Description{}
// 		trade_away.Color = args[i]
// 		trade_away.Size =  will_size
// 		fmt.Println("! created trade_away: " + args[i])
// 		jsonAsBytes, _ = json.Marshal(trade_away)
// 		err = stub.PutState("_debug2", jsonAsBytes)
		
// 		open.Willing = append(open.Willing, trade_away)
// 		fmt.Println("! appended willing to open")
// 		i++;
// 	}
	
// 	//get the open trade struct
// 	tradesAsBytes, err := stub.GetState(openTradesStr)
// 	if err != nil {
// 		return nil, errors.New("Failed to get opentrades")
// 	}
// 	var trades AllTrades
// 	json.Unmarshal(tradesAsBytes, &trades)										//un stringify it aka JSON.parse()
	
// 	trades.OpenTrades = append(trades.OpenTrades, open);						//append to open trades
// 	fmt.Println("! appended open to trades")
// 	jsonAsBytes, _ = json.Marshal(trades)
// 	err = stub.PutState(openTradesStr, jsonAsBytes)								//rewrite open orders
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Println("- end open trade")
// 	return nil, nil
// }

// ============================================================================================================================
// Perform Trade - close an open trade and move ownership
// ============================================================================================================================
// func (t *SimpleChaincode) perform_trade(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
// 	var err error
	
// 	//	0		1					2					3				4					5
// 	//[data.id, data.closer.user, data.closer.name, data.opener.user, data.opener.color, data.opener.size]
// 	if len(args) < 6 {
// 		return nil, errors.New("Incorrect number of arguments. Expecting 6")
// 	}
	
// 	fmt.Println("- start close trade")
// 	timestamp, err := strconv.ParseInt(args[0], 10, 64)
// 	if err != nil {
// 		return nil, errors.New("1st argument must be a numeric string")
// 	}
	
// 	size, err := strconv.Atoi(args[5])
// 	if err != nil {
// 		return nil, errors.New("6th argument must be a numeric string")
// 	}
	
// 	//get the open trade struct
// 	tradesAsBytes, err := stub.GetState(openTradesStr)
// 	if err != nil {
// 		return nil, errors.New("Failed to get opentrades")
// 	}
// 	var trades AllTrades
// 	json.Unmarshal(tradesAsBytes, &trades)																//un stringify it aka JSON.parse()
	
// 	for i := range trades.OpenTrades{																		//look for the trade
// 		fmt.Println("looking at " + strconv.FormatInt(trades.OpenTrades[i].Timestamp, 10) + " for " + strconv.FormatInt(timestamp, 10))
// 		if trades.OpenTrades[i].Timestamp == timestamp{
// 			fmt.Println("found the trade");
			
// 			marble, e := findMarble4Trade(stub, trades.OpenTrades[i].User, args[4], size)					//find a marble that is suitable from opener
// 			if(e == nil){
// 				fmt.Println("! no errors, proceeding")

// 				t.set_user(stub, []string{args[2], trades.OpenTrades[i].User})								//change owner of selected marble, closer -> opener
// 				t.set_user(stub, []string{marble.Name, args[1]})											//change owner of selected marble, opener -> closer
			
// 				trades.OpenTrades = append(trades.OpenTrades[:i], trades.OpenTrades[i+1:]...)				//remove trade
// 				jsonAsBytes, _ := json.Marshal(trades)
// 				err = stub.PutState(openTradesStr, jsonAsBytes)												//rewrite open orders
// 				if err != nil {
// 					return nil, err
// 				}
// 			}
// 		}
// 	}
// 	fmt.Println("- end close trade")
// 	return nil, nil
// }

// ============================================================================================================================
// findMarble4Trade - look for a matching marble that this user owns and return it
// ============================================================================================================================
// func findMarble4Trade(stub *shim.ChaincodeStub, user string, color string, size int )(m Marble, err error){
// 	var fail Marble;
// 	fmt.Println("- start find marble 4 trade")
// 	fmt.Println("looking for " + user + ", " + color + ", " + strconv.Itoa(size));

// 	//get the marble index
// 	marblesAsBytes, err := stub.GetState(marbleIndexStr)
// 	if err != nil {
// 		return fail, errors.New("Failed to get marble index")
// 	}
// 	var marbleIndex []string
// 	json.Unmarshal(marblesAsBytes, &marbleIndex)								//un stringify it aka JSON.parse()
	
// 	for i:= range marbleIndex{													//iter through all the marbles
// 		//fmt.Println("looking @ marble name: " + marbleIndex[i]);

// 		marbleAsBytes, err := stub.GetState(marbleIndex[i])						//grab this marble
// 		if err != nil {
// 			return fail, errors.New("Failed to get marble")
// 		}
// 		res := Marble{}
// 		json.Unmarshal(marbleAsBytes, &res)										//un stringify it aka JSON.parse()
// 		//fmt.Println("looking @ " + res.User + ", " + res.Color + ", " + strconv.Itoa(res.Size));
		
// 		//check for user && color && size
// 		if strings.ToLower(res.User) == strings.ToLower(user) && strings.ToLower(res.Color) == strings.ToLower(color) && res.Size == size{
// 			fmt.Println("found a marble: " + res.Name)
// 			fmt.Println("! end find marble 4 trade")
// 			return res, nil
// 		}
// 	}
	
// 	fmt.Println("- end find marble 4 trade - error")
// 	return fail, errors.New("Did not find marble to use in this trade")
// }

// ============================================================================================================================
// Make Timestamp - create a timestamp in ms
// ============================================================================================================================
func makeTimestamp() int64 {
    return time.Now().UnixNano() / (int64(time.Millisecond)/int64(time.Nanosecond))
}

// ============================================================================================================================
// Remove Open Trade - close an open trade
// ============================================================================================================================
// func (t *SimpleChaincode) remove_trade(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
// 	var err error
	
// 	//	0
// 	//[data.id]
// 	if len(args) < 1 {
// 		return nil, errors.New("Incorrect number of arguments. Expecting 1")
// 	}
	
// 	fmt.Println("- start remove trade")
// 	timestamp, err := strconv.ParseInt(args[0], 10, 64)
// 	if err != nil {
// 		return nil, errors.New("1st argument must be a numeric string")
// 	}
	
// 	//get the open trade struct
// 	tradesAsBytes, err := stub.GetState(openTradesStr)
// 	if err != nil {
// 		return nil, errors.New("Failed to get opentrades")
// 	}
// 	var trades AllTrades
// 	json.Unmarshal(tradesAsBytes, &trades)																//un stringify it aka JSON.parse()
	
// 	for i := range trades.OpenTrades{																	//look for the trade
// 		//fmt.Println("looking at " + strconv.FormatInt(trades.OpenTrades[i].Timestamp, 10) + " for " + strconv.FormatInt(timestamp, 10))
// 		if trades.OpenTrades[i].Timestamp == timestamp{
// 			fmt.Println("found the trade");
// 			trades.OpenTrades = append(trades.OpenTrades[:i], trades.OpenTrades[i+1:]...)				//remove this trade
// 			jsonAsBytes, _ := json.Marshal(trades)
// 			err = stub.PutState(openTradesStr, jsonAsBytes)												//rewrite open orders
// 			if err != nil {
// 				return nil, err
// 			}
// 			break
// 		}
// 	}
	
// 	fmt.Println("- end remove trade")
// 	return nil, nil
// }

// ============================================================================================================================
// Clean Up Open Trades - make sure open trades are still possible, remove choices that are no longer possible, remove trades that have no valid choices
// ============================================================================================================================
// func cleanTrades(stub *shim.ChaincodeStub)(err error){
// 	var didWork = false
// 	fmt.Println("- start clean trades")
	
// 	//get the open trade struct
// 	tradesAsBytes, err := stub.GetState(openTradesStr)
// 	if err != nil {
// 		return errors.New("Failed to get opentrades")
// 	}
// 	var trades AllTrades
// 	json.Unmarshal(tradesAsBytes, &trades)																		//un stringify it aka JSON.parse()
	
// 	fmt.Println("# trades " + strconv.Itoa(len(trades.OpenTrades)))
// 	for i:=0; i<len(trades.OpenTrades); {																		//iter over all the known open trades
// 		fmt.Println(strconv.Itoa(i) + ": looking at trade " + strconv.FormatInt(trades.OpenTrades[i].Timestamp, 10))
		
// 		fmt.Println("# options " + strconv.Itoa(len(trades.OpenTrades[i].Willing)))
// 		for x:=0; x<len(trades.OpenTrades[i].Willing); {														//find a marble that is suitable
// 			fmt.Println("! on next option " + strconv.Itoa(i) + ":" + strconv.Itoa(x))
// 			_, e := findMarble4Trade(stub, trades.OpenTrades[i].User, trades.OpenTrades[i].Willing[x].Color, trades.OpenTrades[i].Willing[x].Size)
// 			if(e != nil){
// 				fmt.Println("! errors with this option, removing option")
// 				didWork = true
// 				trades.OpenTrades[i].Willing = append(trades.OpenTrades[i].Willing[:x], trades.OpenTrades[i].Willing[x+1:]...)	//remove this option
// 				x--;
// 			}else{
// 				fmt.Println("! this option is fine")
// 			}
			
// 			x++
// 			fmt.Println("! x:" + strconv.Itoa(x))
// 			if x >= len(trades.OpenTrades[i].Willing) {														//things might have shifted, recalcuate
// 				break
// 			}
// 		}
		
// 		if len(trades.OpenTrades[i].Willing) == 0 {
// 			fmt.Println("! no more options for this trade, removing trade")
// 			didWork = true
// 			trades.OpenTrades = append(trades.OpenTrades[:i], trades.OpenTrades[i+1:]...)					//remove this trade
// 			i--;
// 		}
		
// 		i++
// 		fmt.Println("! i:" + strconv.Itoa(i))
// 		if i >= len(trades.OpenTrades) {																	//things might have shifted, recalcuate
// 			break
// 		}
// 	}

// 	if(didWork){
// 		fmt.Println("! saving open trade changes")
// 		jsonAsBytes, _ := json.Marshal(trades)
// 		err = stub.PutState(openTradesStr, jsonAsBytes)														//rewrite open orders
// 		if err != nil {
// 			return err
// 		}
// 	}else{
// 		fmt.Println("! all open trades are fine")
// 	}

// 	fmt.Println("- end clean trades")
// 	return nil
// }