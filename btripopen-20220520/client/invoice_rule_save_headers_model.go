// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceRuleSaveHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InvoiceRuleSaveHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *InvoiceRuleSaveHeaders
	GetXAcsBtripSoCorpToken() *string
}

type InvoiceRuleSaveHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s InvoiceRuleSaveHeaders) String() string {
	return dara.Prettify(s)
}

func (s InvoiceRuleSaveHeaders) GoString() string {
	return s.String()
}

func (s *InvoiceRuleSaveHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InvoiceRuleSaveHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *InvoiceRuleSaveHeaders) SetCommonHeaders(v map[string]*string) *InvoiceRuleSaveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvoiceRuleSaveHeaders) SetXAcsBtripSoCorpToken(v string) *InvoiceRuleSaveHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *InvoiceRuleSaveHeaders) Validate() error {
	return dara.Validate(s)
}
