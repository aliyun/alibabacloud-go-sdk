// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundDetailV2Headers interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightRefundDetailV2Headers
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *FlightRefundDetailV2Headers
	GetXAcsBtripCorpToken() *string
}

type FlightRefundDetailV2Headers struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s FlightRefundDetailV2Headers) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundDetailV2Headers) GoString() string {
	return s.String()
}

func (s *FlightRefundDetailV2Headers) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightRefundDetailV2Headers) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *FlightRefundDetailV2Headers) SetCommonHeaders(v map[string]*string) *FlightRefundDetailV2Headers {
	s.CommonHeaders = v
	return s
}

func (s *FlightRefundDetailV2Headers) SetXAcsBtripCorpToken(v string) *FlightRefundDetailV2Headers {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *FlightRefundDetailV2Headers) Validate() error {
	return dara.Validate(s)
}
