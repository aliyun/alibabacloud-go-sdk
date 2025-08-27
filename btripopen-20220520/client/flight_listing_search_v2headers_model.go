// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightListingSearchV2Headers interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightListingSearchV2Headers
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *FlightListingSearchV2Headers
	GetXAcsBtripCorpToken() *string
}

type FlightListingSearchV2Headers struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// Dj2laAwE00
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s FlightListingSearchV2Headers) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchV2Headers) GoString() string {
	return s.String()
}

func (s *FlightListingSearchV2Headers) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightListingSearchV2Headers) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *FlightListingSearchV2Headers) SetCommonHeaders(v map[string]*string) *FlightListingSearchV2Headers {
	s.CommonHeaders = v
	return s
}

func (s *FlightListingSearchV2Headers) SetXAcsBtripCorpToken(v string) *FlightListingSearchV2Headers {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *FlightListingSearchV2Headers) Validate() error {
	return dara.Validate(s)
}
