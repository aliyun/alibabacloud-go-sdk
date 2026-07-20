// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightOtaItemDetailHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IntlFlightOtaItemDetailHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *IntlFlightOtaItemDetailHeaders
	GetXAcsBtripCorpToken() *string
}

type IntlFlightOtaItemDetailHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s IntlFlightOtaItemDetailHeaders) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaItemDetailHeaders) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaItemDetailHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IntlFlightOtaItemDetailHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *IntlFlightOtaItemDetailHeaders) SetCommonHeaders(v map[string]*string) *IntlFlightOtaItemDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IntlFlightOtaItemDetailHeaders) SetXAcsBtripCorpToken(v string) *IntlFlightOtaItemDetailHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *IntlFlightOtaItemDetailHeaders) Validate() error {
	return dara.Validate(s)
}
