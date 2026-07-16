// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopCreateHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IntlFlightReShopCreateHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *IntlFlightReShopCreateHeaders
	GetXAcsBtripCorpToken() *string
}

type IntlFlightReShopCreateHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s IntlFlightReShopCreateHeaders) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopCreateHeaders) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopCreateHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IntlFlightReShopCreateHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *IntlFlightReShopCreateHeaders) SetCommonHeaders(v map[string]*string) *IntlFlightReShopCreateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IntlFlightReShopCreateHeaders) SetXAcsBtripCorpToken(v string) *IntlFlightReShopCreateHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *IntlFlightReShopCreateHeaders) Validate() error {
	return dara.Validate(s)
}
