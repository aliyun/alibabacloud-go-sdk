// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsInvoiceScanQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InsInvoiceScanQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *InsInvoiceScanQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type InsInvoiceScanQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s InsInvoiceScanQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s InsInvoiceScanQueryHeaders) GoString() string {
	return s.String()
}

func (s *InsInvoiceScanQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InsInvoiceScanQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *InsInvoiceScanQueryHeaders) SetCommonHeaders(v map[string]*string) *InsInvoiceScanQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsInvoiceScanQueryHeaders) SetXAcsBtripSoCorpToken(v string) *InsInvoiceScanQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *InsInvoiceScanQueryHeaders) Validate() error {
	return dara.Validate(s)
}
