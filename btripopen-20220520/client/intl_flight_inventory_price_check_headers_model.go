// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightInventoryPriceCheckHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IntlFlightInventoryPriceCheckHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *IntlFlightInventoryPriceCheckHeaders
	GetXAcsBtripCorpToken() *string
}

type IntlFlightInventoryPriceCheckHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s IntlFlightInventoryPriceCheckHeaders) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightInventoryPriceCheckHeaders) GoString() string {
	return s.String()
}

func (s *IntlFlightInventoryPriceCheckHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IntlFlightInventoryPriceCheckHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *IntlFlightInventoryPriceCheckHeaders) SetCommonHeaders(v map[string]*string) *IntlFlightInventoryPriceCheckHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IntlFlightInventoryPriceCheckHeaders) SetXAcsBtripCorpToken(v string) *IntlFlightInventoryPriceCheckHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckHeaders) Validate() error {
	return dara.Validate(s)
}
