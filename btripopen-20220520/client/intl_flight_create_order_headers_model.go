// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightCreateOrderHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IntlFlightCreateOrderHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *IntlFlightCreateOrderHeaders
	GetXAcsBtripCorpToken() *string
}

type IntlFlightCreateOrderHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s IntlFlightCreateOrderHeaders) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightCreateOrderHeaders) GoString() string {
	return s.String()
}

func (s *IntlFlightCreateOrderHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IntlFlightCreateOrderHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *IntlFlightCreateOrderHeaders) SetCommonHeaders(v map[string]*string) *IntlFlightCreateOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IntlFlightCreateOrderHeaders) SetXAcsBtripCorpToken(v string) *IntlFlightCreateOrderHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *IntlFlightCreateOrderHeaders) Validate() error {
	return dara.Validate(s)
}
