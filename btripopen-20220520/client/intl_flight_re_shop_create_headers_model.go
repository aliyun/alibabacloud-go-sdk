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
	// The enterprise access token.
	//
	// - When calling this operation over HTTP, this parameter is required and must be appended to the request URL. For more information about how to obtain the token, see [Enterprise access token](https://openapi.alibtrip.com/doc/toDocDetail?spm=openapi-amp.newDocPublishment.0.0.5e2a281frQyDQ8&docId=3769985).
	//
	// - When appending the token, use crop_token=value instead.
	//
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
