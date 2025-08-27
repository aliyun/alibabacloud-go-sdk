// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundDetailHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightRefundDetailHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *FlightRefundDetailHeaders
	GetXAcsBtripCorpToken() *string
}

type FlightRefundDetailHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s FlightRefundDetailHeaders) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundDetailHeaders) GoString() string {
	return s.String()
}

func (s *FlightRefundDetailHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightRefundDetailHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *FlightRefundDetailHeaders) SetCommonHeaders(v map[string]*string) *FlightRefundDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FlightRefundDetailHeaders) SetXAcsBtripCorpToken(v string) *FlightRefundDetailHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *FlightRefundDetailHeaders) Validate() error {
	return dara.Validate(s)
}
