// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOrderListQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightOrderListQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *FlightOrderListQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type FlightOrderListQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s FlightOrderListQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryHeaders) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightOrderListQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *FlightOrderListQueryHeaders) SetCommonHeaders(v map[string]*string) *FlightOrderListQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FlightOrderListQueryHeaders) SetXAcsBtripSoCorpToken(v string) *FlightOrderListQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *FlightOrderListQueryHeaders) Validate() error {
	return dara.Validate(s)
}
