// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightCreateOrderV2Headers interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightCreateOrderV2Headers
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *FlightCreateOrderV2Headers
	GetXAcsBtripCorpToken() *string
}

type FlightCreateOrderV2Headers struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwis
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s FlightCreateOrderV2Headers) String() string {
	return dara.Prettify(s)
}

func (s FlightCreateOrderV2Headers) GoString() string {
	return s.String()
}

func (s *FlightCreateOrderV2Headers) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightCreateOrderV2Headers) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *FlightCreateOrderV2Headers) SetCommonHeaders(v map[string]*string) *FlightCreateOrderV2Headers {
	s.CommonHeaders = v
	return s
}

func (s *FlightCreateOrderV2Headers) SetXAcsBtripCorpToken(v string) *FlightCreateOrderV2Headers {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *FlightCreateOrderV2Headers) Validate() error {
	return dara.Validate(s)
}
