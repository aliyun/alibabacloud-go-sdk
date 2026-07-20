// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightPayOrderV2Headers interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightPayOrderV2Headers
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *FlightPayOrderV2Headers
	GetXAcsBtripCorpToken() *string
}

type FlightPayOrderV2Headers struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s FlightPayOrderV2Headers) String() string {
	return dara.Prettify(s)
}

func (s FlightPayOrderV2Headers) GoString() string {
	return s.String()
}

func (s *FlightPayOrderV2Headers) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightPayOrderV2Headers) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *FlightPayOrderV2Headers) SetCommonHeaders(v map[string]*string) *FlightPayOrderV2Headers {
	s.CommonHeaders = v
	return s
}

func (s *FlightPayOrderV2Headers) SetXAcsBtripCorpToken(v string) *FlightPayOrderV2Headers {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *FlightPayOrderV2Headers) Validate() error {
	return dara.Validate(s)
}
