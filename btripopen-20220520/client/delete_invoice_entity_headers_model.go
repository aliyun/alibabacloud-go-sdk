// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInvoiceEntityHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteInvoiceEntityHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *DeleteInvoiceEntityHeaders
	GetXAcsBtripCorpToken() *string
}

type DeleteInvoiceEntityHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s DeleteInvoiceEntityHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteInvoiceEntityHeaders) GoString() string {
	return s.String()
}

func (s *DeleteInvoiceEntityHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteInvoiceEntityHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *DeleteInvoiceEntityHeaders) SetCommonHeaders(v map[string]*string) *DeleteInvoiceEntityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteInvoiceEntityHeaders) SetXAcsBtripCorpToken(v string) *DeleteInvoiceEntityHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *DeleteInvoiceEntityHeaders) Validate() error {
	return dara.Validate(s)
}
