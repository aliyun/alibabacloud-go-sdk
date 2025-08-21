// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommonCateSecondFloorHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListCommonCateSecondFloorHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListCommonCateSecondFloorHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListCommonCateSecondFloorHeaders
	GetAuthorization() *string
}

type ListCommonCateSecondFloorHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListCommonCateSecondFloorHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListCommonCateSecondFloorHeaders) GoString() string {
	return s.String()
}

func (s *ListCommonCateSecondFloorHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListCommonCateSecondFloorHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListCommonCateSecondFloorHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListCommonCateSecondFloorHeaders) SetCommonHeaders(v map[string]*string) *ListCommonCateSecondFloorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListCommonCateSecondFloorHeaders) SetXAcsAligenieAccessToken(v string) *ListCommonCateSecondFloorHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListCommonCateSecondFloorHeaders) SetAuthorization(v string) *ListCommonCateSecondFloorHeaders {
	s.Authorization = &v
	return s
}

func (s *ListCommonCateSecondFloorHeaders) Validate() error {
	return dara.Validate(s)
}
