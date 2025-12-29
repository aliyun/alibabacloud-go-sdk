// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInfraredDeviceBrandsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListInfraredDeviceBrandsHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListInfraredDeviceBrandsHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListInfraredDeviceBrandsHeaders
	GetAuthorization() *string
}

type ListInfraredDeviceBrandsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListInfraredDeviceBrandsHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListInfraredDeviceBrandsHeaders) GoString() string {
	return s.String()
}

func (s *ListInfraredDeviceBrandsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListInfraredDeviceBrandsHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListInfraredDeviceBrandsHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListInfraredDeviceBrandsHeaders) SetCommonHeaders(v map[string]*string) *ListInfraredDeviceBrandsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListInfraredDeviceBrandsHeaders) SetXAcsAligenieAccessToken(v string) *ListInfraredDeviceBrandsHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListInfraredDeviceBrandsHeaders) SetAuthorization(v string) *ListInfraredDeviceBrandsHeaders {
	s.Authorization = &v
	return s
}

func (s *ListInfraredDeviceBrandsHeaders) Validate() error {
	return dara.Validate(s)
}
