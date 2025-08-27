// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOrderDetailInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightOrderDetailInfoHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *FlightOrderDetailInfoHeaders
	GetXAcsBtripCorpToken() *string
}

type FlightOrderDetailInfoHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s FlightOrderDetailInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderDetailInfoHeaders) GoString() string {
	return s.String()
}

func (s *FlightOrderDetailInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightOrderDetailInfoHeaders) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *FlightOrderDetailInfoHeaders) SetCommonHeaders(v map[string]*string) *FlightOrderDetailInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FlightOrderDetailInfoHeaders) SetXAcsBtripCorpToken(v string) *FlightOrderDetailInfoHeaders {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *FlightOrderDetailInfoHeaders) Validate() error {
	return dara.Validate(s)
}
