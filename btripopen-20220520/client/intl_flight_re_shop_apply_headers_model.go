// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopApplyHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IntlFlightReShopApplyHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *IntlFlightReShopApplyHeaders
	GetXAcsBtripCorpToken() *string
}

type IntlFlightReShopApplyHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s IntlFlightReShopApplyHeaders) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopApplyHeaders) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopApplyHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IntlFlightReShopApplyHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *IntlFlightReShopApplyHeaders) SetCommonHeaders(v map[string]*string) *IntlFlightReShopApplyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IntlFlightReShopApplyHeaders) SetXAcsBtripCorpToken(v string) *IntlFlightReShopApplyHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *IntlFlightReShopApplyHeaders) Validate() error {
	return dara.Validate(s)
}
