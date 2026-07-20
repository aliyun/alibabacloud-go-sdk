// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceDeleteHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InvoiceDeleteHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *InvoiceDeleteHeaders
	GetXAcsBtripSoCorpToken() *string
}

type InvoiceDeleteHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s InvoiceDeleteHeaders) String() string {
	return dara.Prettify(s)
}

func (s InvoiceDeleteHeaders) GoString() string {
	return s.String()
}

func (s *InvoiceDeleteHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InvoiceDeleteHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *InvoiceDeleteHeaders) SetCommonHeaders(v map[string]*string) *InvoiceDeleteHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvoiceDeleteHeaders) SetXAcsBtripSoCorpToken(v string) *InvoiceDeleteHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *InvoiceDeleteHeaders) Validate() error {
	return dara.Validate(s)
}
