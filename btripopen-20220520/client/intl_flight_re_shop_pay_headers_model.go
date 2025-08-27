// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopPayHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IntlFlightReShopPayHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *IntlFlightReShopPayHeaders
	GetXAcsBtripCorpToken() *string
}

type IntlFlightReShopPayHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s IntlFlightReShopPayHeaders) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopPayHeaders) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopPayHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IntlFlightReShopPayHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *IntlFlightReShopPayHeaders) SetCommonHeaders(v map[string]*string) *IntlFlightReShopPayHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IntlFlightReShopPayHeaders) SetXAcsBtripCorpToken(v string) *IntlFlightReShopPayHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *IntlFlightReShopPayHeaders) Validate() error {
	return dara.Validate(s)
}
