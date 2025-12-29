// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllProvincesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListAllProvincesHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListAllProvincesHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListAllProvincesHeaders
	GetAuthorization() *string
}

type ListAllProvincesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListAllProvincesHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListAllProvincesHeaders) GoString() string {
	return s.String()
}

func (s *ListAllProvincesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListAllProvincesHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListAllProvincesHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListAllProvincesHeaders) SetCommonHeaders(v map[string]*string) *ListAllProvincesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAllProvincesHeaders) SetXAcsAligenieAccessToken(v string) *ListAllProvincesHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListAllProvincesHeaders) SetAuthorization(v string) *ListAllProvincesHeaders {
	s.Authorization = &v
	return s
}

func (s *ListAllProvincesHeaders) Validate() error {
	return dara.Validate(s)
}
