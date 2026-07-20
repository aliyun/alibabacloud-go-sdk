// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightExceedApplyQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FlightExceedApplyQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *FlightExceedApplyQueryHeaders
	GetXAcsBtripSoCorpToken() *string
}

type FlightExceedApplyQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s FlightExceedApplyQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s FlightExceedApplyQueryHeaders) GoString() string {
	return s.String()
}

func (s *FlightExceedApplyQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FlightExceedApplyQueryHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *FlightExceedApplyQueryHeaders) SetCommonHeaders(v map[string]*string) *FlightExceedApplyQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FlightExceedApplyQueryHeaders) SetXAcsBtripSoCorpToken(v string) *FlightExceedApplyQueryHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *FlightExceedApplyQueryHeaders) Validate() error {
	return dara.Validate(s)
}
