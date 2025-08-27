// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightPayOrderHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightPayOrderHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *FlightPayOrderHeaders
	GetXAcsBtripCorpToken() *string
}

type FlightPayOrderHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s FlightPayOrderHeaders) String() string {
	return dara.Prettify(s)
}

func (s FlightPayOrderHeaders) GoString() string {
	return s.String()
}

func (s *FlightPayOrderHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightPayOrderHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *FlightPayOrderHeaders) SetCommonHeaders(v map[string]*string) *FlightPayOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FlightPayOrderHeaders) SetXAcsBtripCorpToken(v string) *FlightPayOrderHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *FlightPayOrderHeaders) Validate() error {
	return dara.Validate(s)
}
