// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightModifyOtaSearchV2Headers interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightModifyOtaSearchV2Headers
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *FlightModifyOtaSearchV2Headers
	GetXAcsBtripCorpToken() *string
}

type FlightModifyOtaSearchV2Headers struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s FlightModifyOtaSearchV2Headers) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOtaSearchV2Headers) GoString() string {
	return s.String()
}

func (s *FlightModifyOtaSearchV2Headers) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightModifyOtaSearchV2Headers) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *FlightModifyOtaSearchV2Headers) SetCommonHeaders(v map[string]*string) *FlightModifyOtaSearchV2Headers {
	s.CommonHeaders = v
	return s
}

func (s *FlightModifyOtaSearchV2Headers) SetXAcsBtripCorpToken(v string) *FlightModifyOtaSearchV2Headers {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *FlightModifyOtaSearchV2Headers) Validate() error {
	return dara.Validate(s)
}
