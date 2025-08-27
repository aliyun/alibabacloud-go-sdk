// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopDetailHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IntlFlightReShopDetailHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *IntlFlightReShopDetailHeaders
	GetXAcsBtripCorpToken() *string
}

type IntlFlightReShopDetailHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s IntlFlightReShopDetailHeaders) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailHeaders) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IntlFlightReShopDetailHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *IntlFlightReShopDetailHeaders) SetCommonHeaders(v map[string]*string) *IntlFlightReShopDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IntlFlightReShopDetailHeaders) SetXAcsBtripCorpToken(v string) *IntlFlightReShopDetailHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *IntlFlightReShopDetailHeaders) Validate() error {
	return dara.Validate(s)
}
