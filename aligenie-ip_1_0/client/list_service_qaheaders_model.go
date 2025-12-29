// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceQAHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListServiceQAHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListServiceQAHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListServiceQAHeaders
	GetAuthorization() *string
}

type ListServiceQAHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListServiceQAHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListServiceQAHeaders) GoString() string {
	return s.String()
}

func (s *ListServiceQAHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListServiceQAHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListServiceQAHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListServiceQAHeaders) SetCommonHeaders(v map[string]*string) *ListServiceQAHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListServiceQAHeaders) SetXAcsAligenieAccessToken(v string) *ListServiceQAHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListServiceQAHeaders) SetAuthorization(v string) *ListServiceQAHeaders {
	s.Authorization = &v
	return s
}

func (s *ListServiceQAHeaders) Validate() error {
	return dara.Validate(s)
}
