// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAirportSearchHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AirportSearchHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *AirportSearchHeaders
	GetXAcsBtripSoCorpToken() *string
}

type AirportSearchHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s AirportSearchHeaders) String() string {
	return dara.Prettify(s)
}

func (s AirportSearchHeaders) GoString() string {
	return s.String()
}

func (s *AirportSearchHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AirportSearchHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *AirportSearchHeaders) SetCommonHeaders(v map[string]*string) *AirportSearchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AirportSearchHeaders) SetXAcsBtripSoCorpToken(v string) *AirportSearchHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *AirportSearchHeaders) Validate() error {
	return dara.Validate(s)
}
