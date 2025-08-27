// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightRefundApplyHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *IntlFlightRefundApplyHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *IntlFlightRefundApplyHeaders
	GetXAcsBtripCorpToken() *string
}

type IntlFlightRefundApplyHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s IntlFlightRefundApplyHeaders) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundApplyHeaders) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundApplyHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *IntlFlightRefundApplyHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *IntlFlightRefundApplyHeaders) SetCommonHeaders(v map[string]*string) *IntlFlightRefundApplyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IntlFlightRefundApplyHeaders) SetXAcsBtripCorpToken(v string) *IntlFlightRefundApplyHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *IntlFlightRefundApplyHeaders) Validate() error {
	return dara.Validate(s)
}
