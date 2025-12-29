// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSceneCategoryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListSceneCategoryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListSceneCategoryHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListSceneCategoryHeaders
	GetAuthorization() *string
}

type ListSceneCategoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListSceneCategoryHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListSceneCategoryHeaders) GoString() string {
	return s.String()
}

func (s *ListSceneCategoryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListSceneCategoryHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListSceneCategoryHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListSceneCategoryHeaders) SetCommonHeaders(v map[string]*string) *ListSceneCategoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSceneCategoryHeaders) SetXAcsAligenieAccessToken(v string) *ListSceneCategoryHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListSceneCategoryHeaders) SetAuthorization(v string) *ListSceneCategoryHeaders {
	s.Authorization = &v
	return s
}

func (s *ListSceneCategoryHeaders) Validate() error {
	return dara.Validate(s)
}
