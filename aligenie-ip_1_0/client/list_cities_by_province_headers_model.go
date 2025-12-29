// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCitiesByProvinceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListCitiesByProvinceHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListCitiesByProvinceHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListCitiesByProvinceHeaders
	GetAuthorization() *string
}

type ListCitiesByProvinceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListCitiesByProvinceHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListCitiesByProvinceHeaders) GoString() string {
	return s.String()
}

func (s *ListCitiesByProvinceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListCitiesByProvinceHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListCitiesByProvinceHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListCitiesByProvinceHeaders) SetCommonHeaders(v map[string]*string) *ListCitiesByProvinceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListCitiesByProvinceHeaders) SetXAcsAligenieAccessToken(v string) *ListCitiesByProvinceHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListCitiesByProvinceHeaders) SetAuthorization(v string) *ListCitiesByProvinceHeaders {
	s.Authorization = &v
	return s
}

func (s *ListCitiesByProvinceHeaders) Validate() error {
	return dara.Validate(s)
}
