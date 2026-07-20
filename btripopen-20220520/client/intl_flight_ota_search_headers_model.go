// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightOtaSearchHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IntlFlightOtaSearchHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *IntlFlightOtaSearchHeaders
	GetXAcsBtripCorpToken() *string
}

type IntlFlightOtaSearchHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s IntlFlightOtaSearchHeaders) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaSearchHeaders) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaSearchHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IntlFlightOtaSearchHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *IntlFlightOtaSearchHeaders) SetCommonHeaders(v map[string]*string) *IntlFlightOtaSearchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IntlFlightOtaSearchHeaders) SetXAcsBtripCorpToken(v string) *IntlFlightOtaSearchHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *IntlFlightOtaSearchHeaders) Validate() error {
	return dara.Validate(s)
}
