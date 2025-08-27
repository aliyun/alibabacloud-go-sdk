// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopConsultHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IntlFlightReShopConsultHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *IntlFlightReShopConsultHeaders
	GetXAcsBtripCorpToken() *string
}

type IntlFlightReShopConsultHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s IntlFlightReShopConsultHeaders) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopConsultHeaders) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopConsultHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IntlFlightReShopConsultHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *IntlFlightReShopConsultHeaders) SetCommonHeaders(v map[string]*string) *IntlFlightReShopConsultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IntlFlightReShopConsultHeaders) SetXAcsBtripCorpToken(v string) *IntlFlightReShopConsultHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *IntlFlightReShopConsultHeaders) Validate() error {
	return dara.Validate(s)
}
