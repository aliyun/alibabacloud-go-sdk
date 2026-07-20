// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceSearchHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InvoiceSearchHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *InvoiceSearchHeaders
	GetXAcsBtripSoCorpToken() *string
}

type InvoiceSearchHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s InvoiceSearchHeaders) String() string {
	return dara.Prettify(s)
}

func (s InvoiceSearchHeaders) GoString() string {
	return s.String()
}

func (s *InvoiceSearchHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InvoiceSearchHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *InvoiceSearchHeaders) SetCommonHeaders(v map[string]*string) *InvoiceSearchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvoiceSearchHeaders) SetXAcsBtripSoCorpToken(v string) *InvoiceSearchHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *InvoiceSearchHeaders) Validate() error {
	return dara.Validate(s)
}
