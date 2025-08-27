// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightRefundDetailHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IntlFlightRefundDetailHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *IntlFlightRefundDetailHeaders
	GetXAcsBtripCorpToken() *string
}

type IntlFlightRefundDetailHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s IntlFlightRefundDetailHeaders) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundDetailHeaders) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundDetailHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IntlFlightRefundDetailHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *IntlFlightRefundDetailHeaders) SetCommonHeaders(v map[string]*string) *IntlFlightRefundDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IntlFlightRefundDetailHeaders) SetXAcsBtripCorpToken(v string) *IntlFlightRefundDetailHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *IntlFlightRefundDetailHeaders) Validate() error {
	return dara.Validate(s)
}
