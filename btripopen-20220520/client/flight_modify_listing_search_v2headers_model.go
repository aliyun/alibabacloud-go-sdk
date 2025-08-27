// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightModifyListingSearchV2Headers interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightModifyListingSearchV2Headers
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *FlightModifyListingSearchV2Headers
	GetXAcsBtripCorpToken() *string
}

type FlightModifyListingSearchV2Headers struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s FlightModifyListingSearchV2Headers) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2Headers) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2Headers) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightModifyListingSearchV2Headers) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *FlightModifyListingSearchV2Headers) SetCommonHeaders(v map[string]*string) *FlightModifyListingSearchV2Headers {
	s.CommonHeaders = v
	return s
}

func (s *FlightModifyListingSearchV2Headers) SetXAcsBtripCorpToken(v string) *FlightModifyListingSearchV2Headers {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *FlightModifyListingSearchV2Headers) Validate() error {
	return dara.Validate(s)
}
