// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOtaItemDetailHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightOtaItemDetailHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *FlightOtaItemDetailHeaders
	GetXAcsBtripCorpToken() *string
}

type FlightOtaItemDetailHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// Dj2laAwE00
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s FlightOtaItemDetailHeaders) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaItemDetailHeaders) GoString() string {
	return s.String()
}

func (s *FlightOtaItemDetailHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightOtaItemDetailHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *FlightOtaItemDetailHeaders) SetCommonHeaders(v map[string]*string) *FlightOtaItemDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FlightOtaItemDetailHeaders) SetXAcsBtripCorpToken(v string) *FlightOtaItemDetailHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *FlightOtaItemDetailHeaders) Validate() error {
	return dara.Validate(s)
}
