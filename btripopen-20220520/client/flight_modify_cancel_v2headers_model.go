// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightModifyCancelV2Headers interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightModifyCancelV2Headers
	GetCommonHeaders() map[string]*string
	SetXAcsBtripCorpToken(v string) *FlightModifyCancelV2Headers
	GetXAcsBtripCorpToken() *string
}

type FlightModifyCancelV2Headers struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripCorpToken *string `json:"x-acs-btrip-corp-token,omitempty" xml:"x-acs-btrip-corp-token,omitempty"`
}

func (s FlightModifyCancelV2Headers) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyCancelV2Headers) GoString() string {
	return s.String()
}

func (s *FlightModifyCancelV2Headers) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightModifyCancelV2Headers) GetXAcsBtripCorpToken() *string {
	return s.XAcsBtripCorpToken
}

func (s *FlightModifyCancelV2Headers) SetCommonHeaders(v map[string]*string) *FlightModifyCancelV2Headers {
	s.CommonHeaders = v
	return s
}

func (s *FlightModifyCancelV2Headers) SetXAcsBtripCorpToken(v string) *FlightModifyCancelV2Headers {
	s.XAcsBtripCorpToken = &v
	return s
}

func (s *FlightModifyCancelV2Headers) Validate() error {
	return dara.Validate(s)
}
