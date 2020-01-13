package gogmi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)


// Version 2 of your GMI Client
type GMI struct {
	APIKey     string
}

func (gmi *GMI) do(path string, methode string, in map[string]interface{}, out interface{}) error {
	// Create client
	client := &http.Client{}

	// add key
	if in == nil {
		in = map[string]interface{}{}
	}
	in["api_key"] = gmi.APIKey

	// Turn the struct into JSON bytes
	b, _ := json.Marshal(&in)
	// Post JSON request to FreshDesk
	req, _ := http.NewRequest(methode, fmt.Sprintf("https://api.getmyinvoices.com/accounts/v2/%s", path), bytes.NewReader(b))
	req.Header.Add("Content-type", "application/json")
	res, e := client.Do(req)
	if e != nil {
		return e
	}
	defer res.Body.Close()

	// Check the status
	if res.StatusCode != 200 {
		return errors.New("server didn't like the request")
	}
	// Grab the JSON response
	if e = json.NewDecoder(res.Body).Decode(out); e != nil {
		return e
	}
	return nil
}

// ListCompanies give a list of all companies
func (gmi *GMI) ListCompanies() (companies Companies, err error) {
	err = gmi.do("listCompanies", http.MethodPost, nil, &companies)
	return
}

// GetCompany returns a specific company
func (gmi *GMI) GetCompany(primUID int) (company Company, err error) {
	err = gmi.do("getCompany", http.MethodPost, map[string]interface{}{"company_id": primUID}, &company)
	return
}

// ListDocuments returns all documents. ListDocuments is equivalent to the ListInvoices in v1, but gives additional data.
func (gmi *GMI) ListDocuments() (documents []Document, err error) {
	var rack RecordsRack
	err = gmi.do("listDocuments", http.MethodPost, nil, &rack)
	documents = rack.Documents
	return
}

// ListDocuments returns all documents with the given parameters. For a list of valid parameters have a look at https://api.getmyinvoices.com/accounts/v2/doc/index.html#listdocuments_post
func (gmi *GMI) ListDocumentsWithParams() (documents []Document, params map[string]interface{}, err error) {
	var rack RecordsRack
	err = gmi.do("listDocuments", http.MethodPost, params, &rack)
	documents = rack.Documents
	return
}

// ListDocumentsFilterByDate returns all Documents with a filter by date
func (gmi *GMI) ListDocumentsFilterByDate(startDate time.Time) (documents []Document, err error) {
	var rack RecordsRack
	in := map[string]interface{}{}
	in["start_date_filter"] = startDate.Format("2006-01-02")
	err = gmi.do("listDocuments", http.MethodPost, nil, &rack)
	documents = rack.Documents
	return
}

// GetDocument returns specific document
func (gmi *GMI) GetDocument(primUID PrimUID) (rack interface{}, err error) {
	err = gmi.do("getDocument", http.MethodPost, map[string]interface{}{"document_prim_uid": primUID}, &rack)
	return
}

// GetCountries returns a slice of countries
func (gmi *GMI) GetCountries() (countries Countries, err error) {
	err = gmi.do("getCountries", http.MethodPost, nil, &countries)
	return
}
