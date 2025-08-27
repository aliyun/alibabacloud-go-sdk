// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceRuleAddHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InvoiceRuleAddHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *InvoiceRuleAddHeaders
	GetXAcsBtripCorpToken() *string
}

type InvoiceRuleAddHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s InvoiceRuleAddHeaders) String() string {
	return dara.Prettify(s)
}

func (s InvoiceRuleAddHeaders) GoString() string {
	return s.String()
}

func (s *InvoiceRuleAddHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InvoiceRuleAddHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *InvoiceRuleAddHeaders) SetCommonHeaders(v map[string]*string) *InvoiceRuleAddHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvoiceRuleAddHeaders) SetXAcsBtripCorpToken(v string) *InvoiceRuleAddHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *InvoiceRuleAddHeaders) Validate() error {
	return dara.Validate(s)
}
