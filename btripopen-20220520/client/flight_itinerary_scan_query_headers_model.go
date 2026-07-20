// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightItineraryScanQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightItineraryScanQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *FlightItineraryScanQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type FlightItineraryScanQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s FlightItineraryScanQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s FlightItineraryScanQueryHeaders) GoString() string {
	return s.String()
}

func (s *FlightItineraryScanQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightItineraryScanQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *FlightItineraryScanQueryHeaders) SetCommonHeaders(v map[string]*string) *FlightItineraryScanQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FlightItineraryScanQueryHeaders) SetXAcsBtripSoCorpToken(v string) *FlightItineraryScanQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *FlightItineraryScanQueryHeaders) Validate() error {
	return dara.Validate(s)
}
