// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOrderQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightOrderQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *FlightOrderQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type FlightOrderQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s FlightOrderQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderQueryHeaders) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightOrderQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *FlightOrderQueryHeaders) SetCommonHeaders(v map[string]*string) *FlightOrderQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FlightOrderQueryHeaders) SetXAcsBtripSoCorpToken(v string) *FlightOrderQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *FlightOrderQueryHeaders) Validate() error {
	return dara.Validate(s)
}
