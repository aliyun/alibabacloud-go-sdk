// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomQAV2Headers interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddCustomQAV2Headers
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *AddCustomQAV2Headers
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *AddCustomQAV2Headers
	GetAuthorization() *string
}

type AddCustomQAV2Headers struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AddCustomQAV2Headers) String() string {
	return dara.Prettify(s)
}

func (s AddCustomQAV2Headers) GoString() string {
	return s.String()
}

func (s *AddCustomQAV2Headers) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddCustomQAV2Headers) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *AddCustomQAV2Headers) GetAuthorization() *string {
	return s.Authorization
}

func (s *AddCustomQAV2Headers) SetCommonHeaders(v map[string]*string) *AddCustomQAV2Headers {
	s.CommonHeaders = v
	return s
}

func (s *AddCustomQAV2Headers) SetXAcsAligenieAccessToken(v string) *AddCustomQAV2Headers {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AddCustomQAV2Headers) SetAuthorization(v string) *AddCustomQAV2Headers {
	s.Authorization = &v
	return s
}

func (s *AddCustomQAV2Headers) Validate() error {
	return dara.Validate(s)
}
