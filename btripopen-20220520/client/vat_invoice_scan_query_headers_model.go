// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVatInvoiceScanQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *VatInvoiceScanQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *VatInvoiceScanQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type VatInvoiceScanQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s VatInvoiceScanQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s VatInvoiceScanQueryHeaders) GoString() string {
	return s.String()
}

func (s *VatInvoiceScanQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *VatInvoiceScanQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *VatInvoiceScanQueryHeaders) SetCommonHeaders(v map[string]*string) *VatInvoiceScanQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *VatInvoiceScanQueryHeaders) SetXAcsBtripSoCorpToken(v string) *VatInvoiceScanQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *VatInvoiceScanQueryHeaders) Validate() error {
	return dara.Validate(s)
}
