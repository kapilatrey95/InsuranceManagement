package main

type Client struct {
	ClientId string `json:"insureId"`
	ClientName string `json:"clientName"`
	Policies []string `json:"policies"`
	RFQArray []string `json:"rfqArray"`

}

type Insurer struct {
	InsurerId string `json:"insurerId"`
	InsurerName string `json:"insurerName"`
	RFQArray []string `json:"rfqArray"`
	Quotes []string `json:"quotes"`
	Policies []string `json:"policies"`

}

type RFQ struct {
	RFQId string `json:"rfqId"`
	ClientId string `json:"insurerId"`
	InsuredName string `json:"insuredName"`
	TypeOfInsurance string `json:"typeOfInsurance"`
	RiskAmount string `json:"riskAmount"`
	Status string `json:"status"`
	Quotes []string `json:"quotes"`
	SelectedInsurer []string `json:"selectedInsurer"` 
	LeadInsurer string `json:"leadInsurerQuote"`
	FinalInsurer []string `json:"finalInsurer"`
	// ClientProposal string `json:clientProposal`

}

// type Policy struct {
// 	PolicyNumber string `json:"policyNumber"`
// 	InsuredName string `json:"insuredName"`
// 	SumInsured string `json:"sumInsured"`
// 	Premium string `json:"premium"`
// 	StartDate string `json:"startDate"`
// 	EndDate string `json:"endDate"`
	
// }


type Broker struct {
	BrokerId string `json:"brokerId"`
	BrokerName string `json:"brokerName"`
	Clients []string `json:"clients"`

}

type Quote struct {
	QuoteId string `json:"quoteId"`
	InsurerName string `json:"insurerName"`
	InsurerId string `json:"insurerId"`
	Premium string `json:"premium"`
	Capacity string `json:"capacity"`
	RFQId string `json:"rfqId"`


}
// type RevisedQuote struct {
// 	RevisedQuoteId string `json:"revisedQuoteId"`  //rfq+quoteId
// 	InsurerName string `json:"insurerName"`
// 	InsurerId string `json:"insurerId"`
// 	Premium string `json:"premium"`
// 	Capacity string `json:"capacity"`
// 	RFQId string `json:"rfqId"`


// }


// type ClientProposal struct {
// LeadInsurerRevisedQuote string
// CoInsurerRevisedQuote []string

// }