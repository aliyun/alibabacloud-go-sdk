// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceAddHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InvoiceAddHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *InvoiceAddHeaders
	GetXAcsBtripSoCorpToken() *string
}

type InvoiceAddHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s InvoiceAddHeaders) String() string {
	return dara.Prettify(s)
}

func (s InvoiceAddHeaders) GoString() string {
	return s.String()
}

func (s *InvoiceAddHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InvoiceAddHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *InvoiceAddHeaders) SetCommonHeaders(v map[string]*string) *InvoiceAddHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvoiceAddHeaders) SetXAcsBtripSoCorpToken(v string) *InvoiceAddHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *InvoiceAddHeaders) Validate() error {
	return dara.Validate(s)
}
