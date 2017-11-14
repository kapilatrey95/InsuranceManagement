/*  completed till  coinsurer agree to the lead insurer quote,
 next step is to provide client to set the capacity of different
  coinsurer and then policy and all  and one more thing ONLY BROKER HAS THE FUCNTIONALITY TO PROVIDE RFQ    */


package main

import (
	"crypto/sha256"
	//"reflect"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"strconv"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	mspprotos "github.com/hyperledger/fabric/protos/msp"
	pb "github.com/hyperledger/fabric/protos/peer"
)


type InsuranceManagement struct {
}

//=============================main=====================================================

func main() {
	fmt.Println("hello world")
	err := shim.Start(new(InsuranceManagement))
	if err != nil {
		fmt.Println("error starting chaincode :%s", err)
	}
}



//======================================init========================================

func (t *InsuranceManagement) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Init called")
	_, args := stub.GetFunctionAndParameters()
	if len(args) != 0 {
		return shim.Error(fmt.Sprintf("chaincode:Init::Wrong number of arguments"))
	}
	return shim.Success(nil)
}


//==============================invoke====================================================

func (t *InsuranceManagement) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()

	if function == "initClient" {

		return t.InitClient(stub, args)   //done
		                
	} else if function == "initInsurer" {

		return t.InitInsurer(stub, args)  //done
	
	} else if function == "generateRFQ" {

		return t.GenerateRFQ(stub, args)   //done

	} else if function == "provideQuote" {

		return t.ProvideQuote(stub, args)   //done

	} else if function == "initBroker" {

		return t.InitBroker(stub, args)    //done

	}else if function == "generateRFQByBroker" {

		return t.GenerateRFQByBroker(stub, args)    //done

	}else if function == "initClientByBroker" {

		return t.InitClientByBroker(stub, args)    //done

	}else if function == "selectLeadInsurer" {

		return t.SelectLeadInsurer(stub, args)    //done
	}

	return shim.Error(fmt.Sprintf("chaincode:Invoke::NO such function exists"))

}

//=========================================================InitClient=================================================

func (t *InsuranceManagement) InitClient(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("init client called")
	fmt.Println("=========================================")
	creator, err := stub.GetCreator()
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:InitClient::couldn't get creator"))
	}
	id := &mspprotos.SerializedIdentity{}
	err = proto.Unmarshal(creator, id)

	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:InitClient::error unmarshalling"))
	}

	block, _ := pem.Decode(id.GetIdBytes())
	// if err !=nil {
	// 	return shim.Error(fmt.Sprintf("couldn decode"));
	// }
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return shim.Error("chaincode:InitClient::couldn pasre ParseCertificate")
	}

	invokerHash := sha256.Sum256([]byte(cert.Subject.CommonName + cert.Issuer.CommonName))
	clientAddress := hex.EncodeToString(invokerHash[:])

	checkClientAsBytes, err := stub.GetState(clientAddress)
	if err != nil || len(checkClientAsBytes) != 0 {
		return shim.Error(fmt.Sprintf("chaincode:InitClient::client already exist"))
	}

	client := Client{}
	client.ClientId = clientAddress
	client.ClientName = cert.Subject.CommonName

	clientAsBytes, err := json.Marshal(client)
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:InitClient::couldn't Unmarsh creator"))
	}

	err = stub.PutState(clientAddress, clientAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:InitClient::couldn't write state "))
	}
	return shim.Success(nil)

}


//=========================================================InitBroker=================================================

func (t *InsuranceManagement) InitBroker(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("init broker called")
	fmt.Println("=======================================")
	creator, err := stub.GetCreator()
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:InitBroker::couldn't get creator"))
	}
	id := &mspprotos.SerializedIdentity{}
	err = proto.Unmarshal(creator, id)

	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:InitBroker::error unmarshalling"))
	}

	block, _ := pem.Decode(id.GetIdBytes())
	// if err !=nil {
	// 	return shim.Error(fmt.Sprintf("couldn decode"));
	// }
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return shim.Error("chaincode:InitBroker::couldn pasre ParseCertificate")
	}

	invokerHash := sha256.Sum256([]byte(cert.Subject.CommonName + cert.Issuer.CommonName))
	brokerAddress := hex.EncodeToString(invokerHash[:])

	checkBrokerAsBytes, err := stub.GetState(brokerAddress)
	if err != nil || len(checkBrokerAsBytes) != 0 {
		return shim.Error(fmt.Sprintf("chaincode:InitBroker::broker already exist"))
	}

	broker := Broker{}
	broker.BrokerId = brokerAddress
	broker.BrokerName = cert.Subject.CommonName

	brokerAsBytes, err := json.Marshal(broker)
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:InitBroker::couldn't Unmarsh creator"))
	}

	err = stub.PutState(brokerAddress, brokerAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:InitBroker::couldn't write state "))
	}
	return shim.Success(nil)

}


//========================================================InitInsurer============


func (t *InsuranceManagement) InitInsurer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("init Insurer called")
	fmt.Println("=========================================")
	creator, err := stub.GetCreator()
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:InitInsurer::couldn't get creator"))
	}
	id := &mspprotos.SerializedIdentity{}
	err = proto.Unmarshal(creator, id)

	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:InitInsurer::error unmarshalling"))
	}

	block, _ := pem.Decode(id.GetIdBytes())
	// if err !=nil {
	// 	return shim.Error(fmt.Sprintf("couldn decode"));
	// }
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return shim.Error("chaincode:InitInsurer::couldn pasre ParseCertificate")
	}

	insurerHash := sha256.Sum256([]byte(cert.Subject.CommonName + cert.Issuer.CommonName))
	insurerAddress := hex.EncodeToString(insurerHash[:])

	checkInsurerAsBytes, err := stub.GetState(insurerAddress)
	if err != nil || len(checkInsurerAsBytes) != 0 {
		return shim.Error(fmt.Sprintf("chaincode:InitInsurer::Insurer already exist"))
	}

	insurer := Insurer{}
	insurer.InsurerId = insurerAddress
	insurer.InsurerName = cert.Subject.CommonName

	insurerAsBytes, err := json.Marshal(insurer)
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:InitInsurer::couldn't Unmarsh creator"))
	}

	err = stub.PutState(insurerAddress, insurerAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:InitInsurer::couldn't write state "))
	}
	return shim.Success(nil)

}




//=========================================================InitClientByBroker=================================================

func (t *InsuranceManagement) InitClientByBroker(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("init InitClientByBroker called")
	fmt.Println("=======================================")
	creator, err := stub.GetCreator()
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:InitClientByBroker::couldn't get creator"))
	}
	id := &mspprotos.SerializedIdentity{}
	err = proto.Unmarshal(creator, id)

	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:InitClientByBroker::error unmarshalling"))
	}

	block, _ := pem.Decode(id.GetIdBytes())
	// if err !=nil {
	// 	return shim.Error(fmt.Sprintf("couldn decode"));
	// }
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return shim.Error("chaincode:InitClientByBroker::couldn pasre ParseCertificate")
	}

	brokerHash := sha256.Sum256([]byte(cert.Subject.CommonName + cert.Issuer.CommonName))
	brokerAddress := hex.EncodeToString(brokerHash[:])

	checkBrokerAsBytes, err := stub.GetState(brokerAddress)
	if err != nil  {
		return shim.Error(fmt.Sprintf("chaincode:InitClientByBroker::broker didnt found "))
	}


	broker := Broker{}
	

	 err = json.Unmarshal(checkBrokerAsBytes,&broker)
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:InitClientByBroker::couldn't Unmarshal creator"))
	}

	tym := time.Now()
	clientHash := sha256.Sum256([]byte(fmt.Sprintf("%s",tym) + cert.Issuer.CommonName))
	clientAddress := hex.EncodeToString(clientHash[:])

	client :=Client{}
	client.ClientId=clientAddress;
	client.ClientName=args[0]

	clientAsBytes,err:=json.Marshal(client);
	if err !=nil {
		return shim.Error(fmt.Sprintf("chaincode:InitClientByBroker::couldn't Marshal client ")) 
	}

	err=stub.PutState(clientAddress,clientAsBytes);
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:InitClientByBroker::couldn't write client state "))
	}

	broker.Clients=append(broker.Clients,clientAddress);

	finalBrokerAsBytes,err:=json.Marshal(broker)
	if err !=nil {
		return shim.Error(fmt.Sprintf("chaincode:InitClientByBroker::couldn't Marshal broker ")) 
	}


	err = stub.PutState(brokerAddress, finalBrokerAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:InitClientByBroker::couldn't write state "))
	}
	return shim.Success(nil)

}



//=======================================================GenerateRFQ================================


func (t *InsuranceManagement) GenerateRFQ(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	//args[0]=RFQID generated on the client side
	//args[1]=ClientId
	//args[2]=InsurerName
	//args[3]=TypeOFinsurance
	//args[4]=RiskAmount
	//args[5]=number of insurer
	//args[6].args[7]......insurer addresses

	NumberOfInsurer, err := strconv.Atoi(args[5])
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:GenerateRFQ::number of insurer is not int"))
	}
	if NumberOfInsurer < 1 {
		return shim.Error("chaincode:GenerateRFQ::provide atleast one insurer")
	}

	creator, err := stub.GetCreator() // it'll give the certificate of the invoker
	id := &mspprotos.SerializedIdentity{}
	err = proto.Unmarshal(creator, id)

	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:GenerateRFQ::couldnt unmarshal creator"))
	}
	block, _ := pem.Decode(id.GetIdBytes())
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:GenerateRFQ::couldnt parse certificate"))
	}
	invokerhash := sha256.Sum256([]byte(cert.Subject.CommonName + cert.Issuer.CommonName))
	clientAddress := hex.EncodeToString(invokerhash[:])

	clientAsBytes, err := stub.GetState(clientAddress)
	if err != nil || clientAsBytes == nil {
		shim.Error(fmt.Sprintf("chaincode:GenerateRFQ::account doesnt exists"))

	}
	client := Client{}

	err = json.Unmarshal(clientAsBytes, &client)
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:GenerateRFQ::couldnt unmarshal client "))
	}
	tym := time.Now()
	tym.Format("Mon Jan _2 15:04:05 2006")
	rfq := RFQ{}
	rfq.ClientId = clientAddress
	rfq.RFQId = args[0]
	rfq.RiskAmount = args[4]
	rfq.TypeOfInsurance = args[3]
	rfq.InsuredName = args[2]
	rfq.Status = "RFQ fired on " + tym.String()

	//var insurerArray []string

	for i := 0; i < NumberOfInsurer; i++ {
		rfq.SelectedInsurer= append(rfq.SelectedInsurer, args[6+i])
		insurerAsBytes,err:=stub.GetState(args[6+i])
		if err!=nil {
			return shim.Error(fmt.Sprintf("Chaincode:generateRFQ:can't get %dth insurer provided",i+1))
		}
		insurer:=Insurer{}
		err = json.Unmarshal(insurerAsBytes, &insurer)
		if err != nil {
			return shim.Error(fmt.Sprintf("chaincode:GenerateRFQ::couldnt unmarshal client "))
		}
		insurer.RFQArray=append(insurer.RFQArray,args[0]);
		finalInsurerAsBytes,err:=json.Marshal(insurer)
		if err!=nil {
			return shim.Error(fmt.Sprintf("Chaincode:generateRFQ:can't marshal the finalInsurerAsBytes "))
		}
		err=stub.PutState(args[6+i],finalInsurerAsBytes)
		if err!=nil {
			return shim.Error(fmt.Sprintf("Chaincode:generateRFQ:couldnt putstate the finalInsurerAsBytes "))
		}

		
	}

	rfqAsBytes,err:=json.Marshal(rfq);
	if err !=nil {
		return shim.Error(fmt.Sprintf("chaincode:generateRfQ couldnt marshal rfq"))
	}

	client.RFQArray=append(client.RFQArray,args[0]);

	finalClientAsBytes,err:=json.Marshal(client);
	if err !=nil {
		return shim.Error(fmt.Sprintf("chaincode:generateRfQ couldnt marshal rfq"))
	}

	err=stub.PutState(args[0],rfqAsBytes);
	if err !=nil {
		return shim.Error(fmt.Sprintf("chaincode:generateRfQ couldnt putstate rfq"))
	}

	err = stub.PutState(clientAddress,finalClientAsBytes)
	if err !=nil {
		return shim.Error(fmt.Sprintf("chaincode:generateRfQ couldnt putstate client"))
	}

	return shim.Success(nil)


	}




//=======================================================GenerateRFQByBroker================================


func (t *InsuranceManagement) GenerateRFQByBroker(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	//args[0]=RFQID generated on the client side
	//args[1]=ClientId
	//args[2]=InsurerName
	//args[3]=TypeOFinsurance
	//args[4]=RiskAmount
	//args[5]=number of insurer
	//args[6].args[7]......insurer addresses

	NumberOfInsurer, err := strconv.Atoi(args[5])
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:GenerateRFQByBroker:number of insurer is not int"))
	}
	if NumberOfInsurer < 1 {
		return shim.Error("chaincode:GenerateRFQByBroker:provide atleast one insurer")
	}

	creator, err := stub.GetCreator() // it'll give the certificate of the invoker
	id := &mspprotos.SerializedIdentity{}
	err = proto.Unmarshal(creator, id)

	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:GenerateRFQByBroker:couldnt unmarshal creator"))
	}
	block, _ := pem.Decode(id.GetIdBytes())
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:GenerateRFQByBroker::couldnt parse certificate"))
	}
	invokerhash := sha256.Sum256([]byte(cert.Subject.CommonName + cert.Issuer.CommonName))
	brokerAddress := hex.EncodeToString(invokerhash[:])

	brokerAsBytes, err := stub.GetState(brokerAddress)
	if err != nil || brokerAsBytes == nil {
		shim.Error(fmt.Sprintf("chaincode:GenerateRFQByBroker::account doesnt exists"))

	}
	broker := Broker{}

	err = json.Unmarshal(brokerAsBytes, &broker)
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:GenerateRFQByBroker:couldnt unmarshal client "))
	}

	clientId:=args[1]
	flag:=1
	brokerClientArray:=broker.Clients;
	lengthOfBrokerClientArray:=len(brokerClientArray)
	for i:=0;i<lengthOfBrokerClientArray;i++ {
		if brokerClientArray[i] == clientId {
			flag=0
			break;
		}
	}

	if flag==1 {
		return shim.Error(fmt.Sprintf("chaincode:GenerateRFQByBroker:couldnt find the client in your stack "))
	}


	clientAsBytes,err:=stub.GetState(args[1]);
	if err !=nil {
		return shim.Error(fmt.Sprintf("chaincode:GenerateRFQByBroker::couldnt get client provided by broker"))
	}

	client := Client{}

	err = json.Unmarshal(clientAsBytes, &client)
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:GenerateRFQByBroker::couldnt unmarshal client "))
	}






	tym := time.Now()
	tym.Format("Mon Jan _2 15:04:05 2006")
	rfq := RFQ{}
	rfq.ClientId = clientId
	rfq.RFQId = args[0]
	rfq.RiskAmount = args[4]
	rfq.TypeOfInsurance = args[3]
	rfq.InsuredName = args[2]
	rfq.Status = "RFQ fired on " + tym.String()

	//var insurerArray []string

	for i := 0; i < NumberOfInsurer; i++ {
		rfq.SelectedInsurer= append(rfq.SelectedInsurer, args[6+i])
		
	}

	rfqAsBytes,err:=json.Marshal(rfq);
	if err !=nil {
		return shim.Error(fmt.Sprintf("chaincode:GenerateRFQByBroker: couldnt marshal rfq"))
	}

	client.RFQArray=append(client.RFQArray,args[0]);

	finalClientAsBytes,err:=json.Marshal(client);
	if err !=nil {
		return shim.Error(fmt.Sprintf("chaincode:GenerateRFQByBroker: couldnt marshal rfq"))
	}

	err=stub.PutState(args[0],rfqAsBytes);
	if err !=nil {
		return shim.Error(fmt.Sprintf("chaincode:GenerateRFQByBroker:couldnt putstate rfq"))
	}

	err = stub.PutState(args[1],finalClientAsBytes)
	if err !=nil {
		return shim.Error(fmt.Sprintf("chaincode:GenerateRFQByBroker: couldnt putstate client"))
	}

	return shim.Success(nil)


	}




	//============================provideQuote==============================================



func (t *InsuranceManagement) ProvideQuote(stub shim.ChaincodeStubInterface, args []string) pb.Response {

		//args[0]=RFQID
		//args[1]=Premium
		//args[2]=Capacity

	creator, err := stub.GetCreator() // it'll give the certificate of the invoker
	id := &mspprotos.SerializedIdentity{}
	err = proto.Unmarshal(creator, id)

	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:ProvideQuote:couldnt unmarshal creator"))
	}
	block, _ := pem.Decode(id.GetIdBytes())
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:ProvideQuote:couldnt parse certificate"))
	}
	invokerhash := sha256.Sum256([]byte(cert.Subject.CommonName + cert.Issuer.CommonName))
	insurerAddress := hex.EncodeToString(invokerhash[:])

	insurerAsBytes, err := stub.GetState(insurerAddress)
	if err != nil || insurerAsBytes == nil {
		shim.Error(fmt.Sprintf("chaincode:ProvideQuote::account doesnt exists"))

	}
	insurer:=Insurer{}

	err = json.Unmarshal(insurerAsBytes, &insurer)
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:ProvideQuote:couldnt unmarshal insurer "))
	}


	rfqAsBytes,err:=stub.GetState(args[0])
	if err !=nil {
		return shim.Error(fmt.Sprintf("chaincode:ProvideQuote::generateRfQ:RFQ doesnt exists"))
	}

	rfq:=RFQ{};

	err = json.Unmarshal(rfqAsBytes, &rfq)
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:ProvideQuote::couldnt unmarshal rfq "))
	}



	quote :=Quote{}
	quoteHash := sha256.Sum256([]byte(cert.Subject.CommonName +args[0] ))
	quoteAddress := hex.EncodeToString(quoteHash[:])

	quote.QuoteId=quoteAddress;
	quote.InsurerName=insurer.InsurerName
	quote.InsurerId=insurerAddress
	quote.Premium=args[1]
	quote.Capacity =args[2]
	quote.RFQId = rfq.RFQId

	quoteAsBytes,err:=json.Marshal(quote)
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:ProvideQuote:couldnt marshal quote "))
	}

	err=stub.PutState(quoteAddress,quoteAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:ProvideQuote:couldnt put state quote "))
	}

	rfq.Quotes=append(rfq.Quotes,quoteAddress)

	finalRFQAsBytes,err:=json.Marshal(rfq)
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:ProvideQuote::couldnt marshal RFQ "))
	}

	err=stub.PutState(args[0],finalRFQAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:ProvideQuote:couldnt put state rfq "))
	}
	insurer.Quotes=append(insurer.Quotes,quote.QuoteId)
	finalInsurerAsBytes,err:=json.Marshal(insurer);
	if err !=nil {
		return shim.Error(fmt.Sprintf("chaincode:ProvideQuote:couldnt marshal RFQ "))
	}
	err=stub.PutState(insurerAddress,finalInsurerAsBytes)
	if err !=nil {
		return shim.Error(fmt.Sprintf("chaincode:ProvideQuote:couldnt put state client ")) 
	}


return shim.Success(nil)

}


//============================================SelectLeadInsurer=====================================================================



func (t *InsuranceManagement) SelectLeadInsurer(stub shim.ChaincodeStubInterface, args []string) pb.Response {

		//args[0]=RFQId
		//args[1]=QuoteId


	creator, err := stub.GetCreator() // it'll give the certificate of the invoker
	id := &mspprotos.SerializedIdentity{}
	err = proto.Unmarshal(creator, id)

	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:SelectLeadInsurer::couldnt unmarshal creator"))
	}
	block, _ := pem.Decode(id.GetIdBytes())
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:SelectLeadInsurer::couldnt parse certificate"))
	}
	invokerhash := sha256.Sum256([]byte(cert.Subject.CommonName + cert.Issuer.CommonName))
	clientAddress := hex.EncodeToString(invokerhash[:])

	clientAsBytes, err := stub.GetState(clientAddress)
	if err != nil || clientAsBytes == nil {
		shim.Error(fmt.Sprintf("chaincode:SelectLeadInsurer::account doesnt exists"))

	}
	client := Client{}

	err = json.Unmarshal(clientAsBytes, &client)
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:SelectLeadInsurer::couldnt unmarshal client "))
	}

	RFQArrayLength:=len(client.RFQArray)
	flag:=1
	for i:=0;i<RFQArrayLength;i++ {
		if client.RFQArray[i]==args[0] {
			flag=0
			break;
		}
	}
	if flag==1 {
		return shim.Error(fmt.Sprintf("chaincode:SelectLeadInsurer::invalid RFQ ID"))
	}

	rfqAsBytes,err:=stub.GetState(args[0]);
	if err !=nil {
		return shim.Error(fmt.Sprintf("chaincode:SelectLeadInsurer::couldnt get RFQ from the state "))
	}

	rfq:=RFQ{}

	err=json.Unmarshal(rfqAsBytes,&rfq)
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:SelectLeadInsurer::couldnt unmarshal rfq "))
	}

	quoteArrayLength := len(rfq.Quotes);
	for i:=0;i<quoteArrayLength;i++ {
		if rfq.Quotes[i] == args[1] {
			flag=0
			break;
		}
	}

	if flag==1 {
		return shim.Error(fmt.Sprintf("chaincode:SelectLeadInsurer::invalid Quote Id"))
	}

	if len(rfq.LeadInsurer)!=0 {
		return shim.Error(fmt.Sprintf("chaincode:SelectLeadInsurer::lead insurer already selected for this RFQ"))
	}

	rfq.LeadInsurer=args[1]

	finalRFQAsBytes,err:=json.Marshal(rfq);
	if err!=nil {
		return shim.Error(fmt.Sprintf("chaincode:SelectLeadInsurer::couldnt marshal finalrfqasbytes "))
	}

	err=stub.PutState(args[0],finalRFQAsBytes)
	if err !=nil {
		return shim.Error(fmt.Sprintf("chaincode:SelectLeadInsurer::couldnt put RFQ to the state "))
	}



return shim.Success(nil);

}


//============================================AcceptLeadQuote=====================================================================



func (t *InsuranceManagement) AcceptLeadQuote(stub shim.ChaincodeStubInterface, args []string) pb.Response {
		//args[0]=rfqId
		//args[1]=quoteId


	creator, err := stub.GetCreator() // it'll give the certificate of the invoker
	id := &mspprotos.SerializedIdentity{}
	err = proto.Unmarshal(creator, id)

	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode::AcceptLeadQuote:couldnt unmarshal creator"))
	}
	block, _ := pem.Decode(id.GetIdBytes())
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode:AcceptLeadQuote:couldnt parse certificate"))
	}
	invokerhash := sha256.Sum256([]byte(cert.Subject.CommonName + cert.Issuer.CommonName))
	insurerAddress := hex.EncodeToString(invokerhash[:])

	insurerAsBytes, err := stub.GetState(insurerAddress)
	if err != nil || insurerAsBytes == nil {
		shim.Error(fmt.Sprintf("chaincode:AcceptLeadQuote:account doesnt exists"))

	}
	insurer:=Insurer{}

	err = json.Unmarshal(insurerAsBytes, &insurer)
	if err != nil {
		return shim.Error(fmt.Sprintf("Chaincode:AcceptLeadQuote:couldnt unmarshal insurer "))
	}

	rfqAsBytes,err:=stub.GetState(args[0])
	if err !=nil {
		return shim.Error(fmt.Sprintf("chaincode:AcceptLeadQuote::RFQ doesnt exists"))
	}

	rfq:=RFQ{};

	err = json.Unmarshal(rfqAsBytes, &rfq)
	if err != nil {
		return shim.Error(fmt.Sprintf("Chaincode:AcceptLeadQuote:couldnt unmarshal rfq "))
	}

	rfq.FinalInsurer=append(rfq.FinalInsurer,insurerAddress)

	finalRfqAsBytes,err:=json.Marshal(rfq);
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode::AcceptLeadQuote:couldnt marshal rfq"))
	}

	err=stub.PutState(args[0],finalRfqAsBytes);
	if err != nil {
		return shim.Error(fmt.Sprintf("chaincode::AcceptLeadQuote:couldnt putstate rfq"))
	}

	return shim.Success(nil)


}



