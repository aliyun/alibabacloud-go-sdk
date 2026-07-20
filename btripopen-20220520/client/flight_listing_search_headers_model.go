// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightListingSearchHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightListingSearchHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *FlightListingSearchHeaders
	GetXAcsBtripCorpToken() *string
}

type FlightListingSearchHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s FlightListingSearchHeaders) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchHeaders) GoString() string {
	return s.String()
}

func (s *FlightListingSearchHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightListingSearchHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *FlightListingSearchHeaders) SetCommonHeaders(v map[string]*string) *FlightListingSearchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FlightListingSearchHeaders) SetXAcsBtripCorpToken(v string) *FlightListingSearchHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *FlightListingSearchHeaders) Validate() error {
	return dara.Validate(s)
}
