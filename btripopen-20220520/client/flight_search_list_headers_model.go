// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightSearchListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightSearchListHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *FlightSearchListHeaders
	GetXAcsBtripCorpToken() *string
}

type FlightSearchListHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s FlightSearchListHeaders) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListHeaders) GoString() string {
	return s.String()
}

func (s *FlightSearchListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightSearchListHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *FlightSearchListHeaders) SetCommonHeaders(v map[string]*string) *FlightSearchListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FlightSearchListHeaders) SetXAcsBtripCorpToken(v string) *FlightSearchListHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *FlightSearchListHeaders) Validate() error {
	return dara.Validate(s)
}
