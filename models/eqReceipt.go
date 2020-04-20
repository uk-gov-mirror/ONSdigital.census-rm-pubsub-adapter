package models

type EqReceiptMetadata struct {
	TransactionID   string `json:"tx_id"`
	QuestionnaireID string `json:"questionnaire_id"`
	CaseID          string `json:"caseId,omitempty"`
}

type EqReceipt struct {
	TimeCreated string            `json:"timeCreated"`
	Metadata    EqReceiptMetadata `json:"metadata"`
}
