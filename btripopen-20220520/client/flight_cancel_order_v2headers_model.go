// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightCancelOrderV2Headers interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightCancelOrderV2Headers
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *FlightCancelOrderV2Headers
	GetXAcsBtripCorpToken() *string
}

type FlightCancelOrderV2Headers struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwis
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s FlightCancelOrderV2Headers) String() string {
	return dara.Prettify(s)
}

func (s FlightCancelOrderV2Headers) GoString() string {
	return s.String()
}

func (s *FlightCancelOrderV2Headers) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightCancelOrderV2Headers) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *FlightCancelOrderV2Headers) SetCommonHeaders(v map[string]*string) *FlightCancelOrderV2Headers {
	s.CommonHeaders = v
	return s
}

func (s *FlightCancelOrderV2Headers) SetXAcsBtripCorpToken(v string) *FlightCancelOrderV2Headers {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *FlightCancelOrderV2Headers) Validate() error {
	return dara.Validate(s)
}
