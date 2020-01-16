package gogmi

import (
	"encoding/json"
	"strconv"
)

type CustomDate string

// Just a current workaround since the property sometimes returns a int instead of a string
func (prim *CustomDate) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		*prim = CustomDate(s)
		return nil
	}

	var n int
	if err := json.Unmarshal(b, &n); err != nil {
		return err
	}

	ns := strconv.Itoa(n)
	*prim = CustomDate(ns)
	return nil
}

// Company is the model for GMI-company
type Company struct {
	PrimUID     uint   `json:"prim_uid"`
	Name        string `json:"name"`
	CompanyType string `json:"company_type"`
	Note        string `json:"note"`
	Tags        string `json:"tags"`
}

// Companies is a slice of companies
type Companies []Company

type ApiStatus struct {
	Success bool `json:"success,omitempty"`
}

// Invoice represent a record
type Document struct {
	PrimUID         uint   `json:"prim_uid,omitempty"`
	CreatedAt       string `json:"created_at,omitempty"`
	CompanyUID      uint   `json:"company_uid,omitempty"`
	DocumentType    string `json:"document_type,omitempty"`
	DocumentNumber  string `json:"document_number,omitempty"`
	DocumentDate    string `json:"document_date,omitempty"`
	DocumentDueDate string `json:"document_due_date,omitempty"`
	NetAmount       string `json:"net_amount,omitempty"`
	Vat             string `json:"vat,omitempty"`
	GrossAmount     string `json:"gross_amount,omitempty"`
	Currency        string `json:"currency,omitempty"`
	IsArchived      string `json:"is_archived,omitempty"`
	IsOcrCompleted  int    `json:"is_ocr_completed,omitempty"`
	Tags            string `json:"tags,omitempty"`
	Note            string `json:"note,omitempty"`
	Source          string `json:"source,omitempty"`
	Filename        string `json:"filename,omitempty"`
	FileSize        string `json:"file_size,omitempty"`
	PaymentStatus   string `json:"payment_status,omitempty"`
	PaymentMethod   string `json:"payment_method,omitempty"`
	PaymentDetails  struct {
		SenderEmail          string     `json:"sender_email,omitempty"`
		RecipientEmail       string     `json:"recipient_email,omitempty"`
		TransactionId        string     `json:"transaction_id,omitempty"`
		CardNumber           string     `json:"card_number,omitempty"`
		SepaCreditorId       string     `json:"sepa_creditor_id,omitempty"`
		SepaMandateReference string     `json:"sepa_mandate_reference,omitempty"`
		PurposeOfUsage       string     `json:"purpose_of_usage,omitempty"`
		Iban                 string     `json:"iban,omitempty"`
		Bic                  string     `json:"bic,omitempty"`
		AccountHolderName    string     `json:"account_holder_name,omitempty"`
		AccountNumber        string     `json:"account_number,omitempty"`
		BankName             string     `json:"bank_name,omitempty"`
		BankAddress          string     `json:"bank_address,omitempty"`
		SortCode             string     `json:"sort_code,omitempty"`
		RoutingNumber        string     `json:"routing_number,omitempty"`
		IfscCode             string     `json:"ifsc_code,omitempty"`
		RoutingCode          string     `json:"routing_code,omitempty"`
		CashDiscountDate     CustomDate `json:"cash_discount_date,omitempty"`
		CashDiscountValue    string     `json:"cash_discount_value,omitempty"`
	} `json:"payment_details,omitempty"`
}

// RecordsRack holds all records
type RecordsRack struct {
	Documents  []Document `json:"records"`
	TotalCount string     `json:"total_count"`
	Start      int        `json:"start"`
	Offset     int        `json:"offset"`
}

// Countries is a slice of Country
type Countries []Country

// Country represent a country
type Country struct {
	PrimUID     uint   `json:"prim_uid"`
	Name        string `json:"name"`
	CountryCode string `json:"country_code"`
	Vat         int    `json:"vat"`
	IsEu        int    `json:"is_eu"`
}
