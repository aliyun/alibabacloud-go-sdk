// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopOtaSearchHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IntlFlightReShopOtaSearchHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *IntlFlightReShopOtaSearchHeaders
	GetXAcsBtripCorpToken() *string
}

type IntlFlightReShopOtaSearchHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s IntlFlightReShopOtaSearchHeaders) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchHeaders) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IntlFlightReShopOtaSearchHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *IntlFlightReShopOtaSearchHeaders) SetCommonHeaders(v map[string]*string) *IntlFlightReShopOtaSearchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IntlFlightReShopOtaSearchHeaders) SetXAcsBtripCorpToken(v string) *IntlFlightReShopOtaSearchHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *IntlFlightReShopOtaSearchHeaders) Validate() error {
	return dara.Validate(s)
}
