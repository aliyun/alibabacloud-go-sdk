// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundPreCalHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightRefundPreCalHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *FlightRefundPreCalHeaders
	GetXAcsBtripCorpToken() *string
}

type FlightRefundPreCalHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s FlightRefundPreCalHeaders) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundPreCalHeaders) GoString() string {
	return s.String()
}

func (s *FlightRefundPreCalHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightRefundPreCalHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *FlightRefundPreCalHeaders) SetCommonHeaders(v map[string]*string) *FlightRefundPreCalHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FlightRefundPreCalHeaders) SetXAcsBtripCorpToken(v string) *FlightRefundPreCalHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *FlightRefundPreCalHeaders) Validate() error {
	return dara.Validate(s)
}
