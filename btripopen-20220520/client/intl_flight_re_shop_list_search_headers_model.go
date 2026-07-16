// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopListSearchHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IntlFlightReShopListSearchHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *IntlFlightReShopListSearchHeaders
	GetXAcsBtripCorpToken() *string
}

type IntlFlightReShopListSearchHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s IntlFlightReShopListSearchHeaders) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopListSearchHeaders) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopListSearchHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IntlFlightReShopListSearchHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *IntlFlightReShopListSearchHeaders) SetCommonHeaders(v map[string]*string) *IntlFlightReShopListSearchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IntlFlightReShopListSearchHeaders) SetXAcsBtripCorpToken(v string) *IntlFlightReShopListSearchHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *IntlFlightReShopListSearchHeaders) Validate() error {
	return dara.Validate(s)
}
