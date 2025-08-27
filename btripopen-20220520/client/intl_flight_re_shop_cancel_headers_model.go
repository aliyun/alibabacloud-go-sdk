// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopCancelHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IntlFlightReShopCancelHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *IntlFlightReShopCancelHeaders
	GetXAcsBtripCorpToken() *string
}

type IntlFlightReShopCancelHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s IntlFlightReShopCancelHeaders) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopCancelHeaders) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopCancelHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IntlFlightReShopCancelHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *IntlFlightReShopCancelHeaders) SetCommonHeaders(v map[string]*string) *IntlFlightReShopCancelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IntlFlightReShopCancelHeaders) SetXAcsBtripCorpToken(v string) *IntlFlightReShopCancelHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *IntlFlightReShopCancelHeaders) Validate() error {
	return dara.Validate(s)
}
