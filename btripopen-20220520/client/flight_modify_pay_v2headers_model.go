// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightModifyPayV2Headers interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightModifyPayV2Headers
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *FlightModifyPayV2Headers
	GetXAcsBtripCorpToken() *string
}

type FlightModifyPayV2Headers struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s FlightModifyPayV2Headers) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyPayV2Headers) GoString() string {
	return s.String()
}

func (s *FlightModifyPayV2Headers) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightModifyPayV2Headers) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *FlightModifyPayV2Headers) SetCommonHeaders(v map[string]*string) *FlightModifyPayV2Headers {
	s.CommonHeaders = v
	return s
}

func (s *FlightModifyPayV2Headers) SetXAcsBtripCorpToken(v string) *FlightModifyPayV2Headers {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *FlightModifyPayV2Headers) Validate() error {
	return dara.Validate(s)
}
