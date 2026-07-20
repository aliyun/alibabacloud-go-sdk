// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCitySearchHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CitySearchHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripSoCorpToken(v string) *CitySearchHeaders
	GetXAcsBtripSoCorpToken() *string
}

type CitySearchHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripSoCorpToken *string `json:"x-acs-btrip-so-corp-token,omitempty" xml:"x-acs-btrip-so-corp-token,omitempty"`
}

func (s CitySearchHeaders) String() string {
	return dara.Prettify(s)
}

func (s CitySearchHeaders) GoString() string {
	return s.String()
}

func (s *CitySearchHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CitySearchHeaders) GetXAcsBtripSoCorpToken() *string {
	return s.XAcsBtripSoCorpToken
}

func (s *CitySearchHeaders) SetCommonHeaders(v map[string]*string) *CitySearchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CitySearchHeaders) SetXAcsBtripSoCorpToken(v string) *CitySearchHeaders {
	s.XAcsBtripSoCorpToken = &v
	return s
}

func (s *CitySearchHeaders) Validate() error {
	return dara.Validate(s)
}
