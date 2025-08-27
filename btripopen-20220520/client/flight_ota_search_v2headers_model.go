// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOtaSearchV2Headers interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightOtaSearchV2Headers
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *FlightOtaSearchV2Headers
	GetXAcsBtripCorpToken() *string
}

type FlightOtaSearchV2Headers struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// Dj2laAwE00
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s FlightOtaSearchV2Headers) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchV2Headers) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchV2Headers) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightOtaSearchV2Headers) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *FlightOtaSearchV2Headers) SetCommonHeaders(v map[string]*string) *FlightOtaSearchV2Headers {
	s.CommonHeaders = v
	return s
}

func (s *FlightOtaSearchV2Headers) SetXAcsBtripCorpToken(v string) *FlightOtaSearchV2Headers {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *FlightOtaSearchV2Headers) Validate() error {
	return dara.Validate(s)
}
