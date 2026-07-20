// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBaseCityInfoSearchHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *BaseCityInfoSearchHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripAccessToken(v string) *BaseCityInfoSearchHeaders
	GetXAcsBtripAccessToken() *string
}

type BaseCityInfoSearchHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripAccessToken *string `json:"x-acs-btrip-access-token,omitempty" xml:"x-acs-btrip-access-token,omitempty"`
}

func (s BaseCityInfoSearchHeaders) String() string {
	return dara.Prettify(s)
}

func (s BaseCityInfoSearchHeaders) GoString() string {
	return s.String()
}

func (s *BaseCityInfoSearchHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *BaseCityInfoSearchHeaders) GetXAcsBtripAccessToken() *string {
	return s.XAcsBtripAccessToken
}

func (s *BaseCityInfoSearchHeaders) SetCommonHeaders(v map[string]*string) *BaseCityInfoSearchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BaseCityInfoSearchHeaders) SetXAcsBtripAccessToken(v string) *BaseCityInfoSearchHeaders {
	s.XAcsBtripAccessToken = &v
	return s
}

func (s *BaseCityInfoSearchHeaders) Validate() error {
	return dara.Validate(s)
}
