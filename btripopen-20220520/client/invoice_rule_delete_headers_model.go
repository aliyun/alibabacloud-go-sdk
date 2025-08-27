// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceRuleDeleteHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InvoiceRuleDeleteHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *InvoiceRuleDeleteHeaders
	GetXAcsBtripCorpToken() *string
}

type InvoiceRuleDeleteHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s InvoiceRuleDeleteHeaders) String() string {
	return dara.Prettify(s)
}

func (s InvoiceRuleDeleteHeaders) GoString() string {
	return s.String()
}

func (s *InvoiceRuleDeleteHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InvoiceRuleDeleteHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *InvoiceRuleDeleteHeaders) SetCommonHeaders(v map[string]*string) *InvoiceRuleDeleteHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvoiceRuleDeleteHeaders) SetXAcsBtripCorpToken(v string) *InvoiceRuleDeleteHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *InvoiceRuleDeleteHeaders) Validate() error {
	return dara.Validate(s)
}
