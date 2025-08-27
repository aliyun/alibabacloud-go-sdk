// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundApplyV2Headers interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightRefundApplyV2Headers
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *FlightRefundApplyV2Headers
	GetXAcsBtripCorpToken() *string
}

type FlightRefundApplyV2Headers struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s FlightRefundApplyV2Headers) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundApplyV2Headers) GoString() string {
	return s.String()
}

func (s *FlightRefundApplyV2Headers) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightRefundApplyV2Headers) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *FlightRefundApplyV2Headers) SetCommonHeaders(v map[string]*string) *FlightRefundApplyV2Headers {
	s.CommonHeaders = v
	return s
}

func (s *FlightRefundApplyV2Headers) SetXAcsBtripCorpToken(v string) *FlightRefundApplyV2Headers {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *FlightRefundApplyV2Headers) Validate() error {
	return dara.Validate(s)
}
