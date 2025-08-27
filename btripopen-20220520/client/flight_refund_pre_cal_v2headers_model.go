// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundPreCalV2Headers interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightRefundPreCalV2Headers
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *FlightRefundPreCalV2Headers
	GetXAcsBtripCorpToken() *string
}

type FlightRefundPreCalV2Headers struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s FlightRefundPreCalV2Headers) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundPreCalV2Headers) GoString() string {
	return s.String()
}

func (s *FlightRefundPreCalV2Headers) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightRefundPreCalV2Headers) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *FlightRefundPreCalV2Headers) SetCommonHeaders(v map[string]*string) *FlightRefundPreCalV2Headers {
	s.CommonHeaders = v
	return s
}

func (s *FlightRefundPreCalV2Headers) SetXAcsBtripCorpToken(v string) *FlightRefundPreCalV2Headers {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *FlightRefundPreCalV2Headers) Validate() error {
	return dara.Validate(s)
}
