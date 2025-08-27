// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddInvoiceEntityHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddInvoiceEntityHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *AddInvoiceEntityHeaders
	GetXAcsBtripCorpToken() *string
}

type AddInvoiceEntityHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s AddInvoiceEntityHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddInvoiceEntityHeaders) GoString() string {
	return s.String()
}

func (s *AddInvoiceEntityHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddInvoiceEntityHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *AddInvoiceEntityHeaders) SetCommonHeaders(v map[string]*string) *AddInvoiceEntityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddInvoiceEntityHeaders) SetXAcsBtripCorpToken(v string) *AddInvoiceEntityHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *AddInvoiceEntityHeaders) Validate() error {
	return dara.Validate(s)
}
