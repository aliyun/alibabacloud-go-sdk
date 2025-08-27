// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundApplyHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightRefundApplyHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *FlightRefundApplyHeaders
	GetXAcsBtripCorpToken() *string
}

type FlightRefundApplyHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s FlightRefundApplyHeaders) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundApplyHeaders) GoString() string {
	return s.String()
}

func (s *FlightRefundApplyHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightRefundApplyHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *FlightRefundApplyHeaders) SetCommonHeaders(v map[string]*string) *FlightRefundApplyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FlightRefundApplyHeaders) SetXAcsBtripCorpToken(v string) *FlightRefundApplyHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *FlightRefundApplyHeaders) Validate() error {
	return dara.Validate(s)
}
