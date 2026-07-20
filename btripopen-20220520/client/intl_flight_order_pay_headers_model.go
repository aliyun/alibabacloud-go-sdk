// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightOrderPayHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IntlFlightOrderPayHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *IntlFlightOrderPayHeaders
	GetXAcsBtripCorpToken() *string
}

type IntlFlightOrderPayHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s IntlFlightOrderPayHeaders) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderPayHeaders) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderPayHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IntlFlightOrderPayHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *IntlFlightOrderPayHeaders) SetCommonHeaders(v map[string]*string) *IntlFlightOrderPayHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IntlFlightOrderPayHeaders) SetXAcsBtripCorpToken(v string) *IntlFlightOrderPayHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *IntlFlightOrderPayHeaders) Validate() error {
	return dara.Validate(s)
}
