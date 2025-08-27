// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightCreateOrderHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightCreateOrderHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *FlightCreateOrderHeaders
	GetXAcsBtripCorpToken() *string
}

type FlightCreateOrderHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s FlightCreateOrderHeaders) String() string {
	return dara.Prettify(s)
}

func (s FlightCreateOrderHeaders) GoString() string {
	return s.String()
}

func (s *FlightCreateOrderHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightCreateOrderHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *FlightCreateOrderHeaders) SetCommonHeaders(v map[string]*string) *FlightCreateOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FlightCreateOrderHeaders) SetXAcsBtripCorpToken(v string) *FlightCreateOrderHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *FlightCreateOrderHeaders) Validate() error {
	return dara.Validate(s)
}
