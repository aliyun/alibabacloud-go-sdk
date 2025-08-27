// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllBaseCityInfoQueryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AllBaseCityInfoQueryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsBtripAccessToken(v string) *AllBaseCityInfoQueryHeaders
	GetXAcsBtripAccessToken() *string
}

type AllBaseCityInfoQueryHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// feth00jqwls
	XAcsBtripAccessToken *string `json:"x-acs-btrip-access-token,omitempty" xml:"x-acs-btrip-access-token,omitempty"`
}

func (s AllBaseCityInfoQueryHeaders) String() string {
	return dara.Prettify(s)
}

func (s AllBaseCityInfoQueryHeaders) GoString() string {
	return s.String()
}

func (s *AllBaseCityInfoQueryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AllBaseCityInfoQueryHeaders) GetXAcsBtripAccessToken() *string {
	return s.XAcsBtripAccessToken
}

func (s *AllBaseCityInfoQueryHeaders) SetCommonHeaders(v map[string]*string) *AllBaseCityInfoQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AllBaseCityInfoQueryHeaders) SetXAcsBtripAccessToken(v string) *AllBaseCityInfoQueryHeaders {
	s.XAcsBtripAccessToken = &v
	return s
}

func (s *AllBaseCityInfoQueryHeaders) Validate() error {
	return dara.Validate(s)
}
