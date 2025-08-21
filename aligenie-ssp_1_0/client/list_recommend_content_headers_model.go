// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecommendContentHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListRecommendContentHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListRecommendContentHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListRecommendContentHeaders
	GetAuthorization() *string
}

type ListRecommendContentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListRecommendContentHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListRecommendContentHeaders) GoString() string {
	return s.String()
}

func (s *ListRecommendContentHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListRecommendContentHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListRecommendContentHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListRecommendContentHeaders) SetCommonHeaders(v map[string]*string) *ListRecommendContentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListRecommendContentHeaders) SetXAcsAligenieAccessToken(v string) *ListRecommendContentHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListRecommendContentHeaders) SetAuthorization(v string) *ListRecommendContentHeaders {
	s.Authorization = &v
	return s
}

func (s *ListRecommendContentHeaders) Validate() error {
	return dara.Validate(s)
}
