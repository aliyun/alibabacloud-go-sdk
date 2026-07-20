// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyInvoiceTaskHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ApplyInvoiceTaskHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *ApplyInvoiceTaskHeaders
	GetXAcsBtripSoCorpToken() *string
}

type ApplyInvoiceTaskHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s ApplyInvoiceTaskHeaders) String() string {
	return dara.Prettify(s)
}

func (s ApplyInvoiceTaskHeaders) GoString() string {
	return s.String()
}

func (s *ApplyInvoiceTaskHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ApplyInvoiceTaskHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *ApplyInvoiceTaskHeaders) SetCommonHeaders(v map[string]*string) *ApplyInvoiceTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ApplyInvoiceTaskHeaders) SetXAcsBtripSoCorpToken(v string) *ApplyInvoiceTaskHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *ApplyInvoiceTaskHeaders) Validate() error {
	return dara.Validate(s)
}
