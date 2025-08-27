// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOrderListQueryV2Headers interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightOrderListQueryV2Headers
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *FlightOrderListQueryV2Headers
	GetXAcsBtripCorpToken() *string
}

type FlightOrderListQueryV2Headers struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s FlightOrderListQueryV2Headers) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2Headers) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2Headers) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightOrderListQueryV2Headers) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *FlightOrderListQueryV2Headers) SetCommonHeaders(v map[string]*string) *FlightOrderListQueryV2Headers {
	s.CommonHeaders = v
	return s
}

func (s *FlightOrderListQueryV2Headers) SetXAcsBtripCorpToken(v string) *FlightOrderListQueryV2Headers {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *FlightOrderListQueryV2Headers) Validate() error {
	return dara.Validate(s)
}
