// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOrderDetailV2Headers interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightOrderDetailV2Headers
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *FlightOrderDetailV2Headers
	GetXAcsBtripCorpToken() *string
}

type FlightOrderDetailV2Headers struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// Dj2laAwE00
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s FlightOrderDetailV2Headers) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderDetailV2Headers) GoString() string {
	return s.String()
}

func (s *FlightOrderDetailV2Headers) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightOrderDetailV2Headers) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *FlightOrderDetailV2Headers) SetCommonHeaders(v map[string]*string) *FlightOrderDetailV2Headers {
	s.CommonHeaders = v
	return s
}

func (s *FlightOrderDetailV2Headers) SetXAcsBtripCorpToken(v string) *FlightOrderDetailV2Headers {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *FlightOrderDetailV2Headers) Validate() error {
	return dara.Validate(s)
}
