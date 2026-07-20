// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOtaSearchHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightOtaSearchHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *FlightOtaSearchHeaders
	GetXAcsBtripCorpToken() *string
}

type FlightOtaSearchHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s FlightOtaSearchHeaders) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchHeaders) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightOtaSearchHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *FlightOtaSearchHeaders) SetCommonHeaders(v map[string]*string) *FlightOtaSearchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FlightOtaSearchHeaders) SetXAcsBtripCorpToken(v string) *FlightOtaSearchHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *FlightOtaSearchHeaders) Validate() error {
	return dara.Validate(s)
}
