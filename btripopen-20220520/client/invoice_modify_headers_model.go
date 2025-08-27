// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceModifyHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InvoiceModifyHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *InvoiceModifyHeaders
	GetXAcsBtripSoCorpToken() *string
}

type InvoiceModifyHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s InvoiceModifyHeaders) String() string {
	return dara.Prettify(s)
}

func (s InvoiceModifyHeaders) GoString() string {
	return s.String()
}

func (s *InvoiceModifyHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InvoiceModifyHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *InvoiceModifyHeaders) SetCommonHeaders(v map[string]*string) *InvoiceModifyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvoiceModifyHeaders) SetXAcsBtripSoCorpToken(v string) *InvoiceModifyHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *InvoiceModifyHeaders) Validate() error {
	return dara.Validate(s)
}
