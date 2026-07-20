// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightCancelOrderHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightCancelOrderHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *FlightCancelOrderHeaders
	GetXAcsBtripCorpToken() *string
}

type FlightCancelOrderHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s FlightCancelOrderHeaders) String() string {
	return dara.Prettify(s)
}

func (s FlightCancelOrderHeaders) GoString() string {
	return s.String()
}

func (s *FlightCancelOrderHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightCancelOrderHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *FlightCancelOrderHeaders) SetCommonHeaders(v map[string]*string) *FlightCancelOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FlightCancelOrderHeaders) SetXAcsBtripCorpToken(v string) *FlightCancelOrderHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *FlightCancelOrderHeaders) Validate() error {
	return dara.Validate(s)
}
