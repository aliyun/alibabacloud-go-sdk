// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommonCateFirstFloorHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListCommonCateFirstFloorHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListCommonCateFirstFloorHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListCommonCateFirstFloorHeaders
	GetAuthorization() *string
}

type ListCommonCateFirstFloorHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListCommonCateFirstFloorHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListCommonCateFirstFloorHeaders) GoString() string {
	return s.String()
}

func (s *ListCommonCateFirstFloorHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListCommonCateFirstFloorHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListCommonCateFirstFloorHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListCommonCateFirstFloorHeaders) SetCommonHeaders(v map[string]*string) *ListCommonCateFirstFloorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListCommonCateFirstFloorHeaders) SetXAcsAligenieAccessToken(v string) *ListCommonCateFirstFloorHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListCommonCateFirstFloorHeaders) SetAuthorization(v string) *ListCommonCateFirstFloorHeaders {
	s.Authorization = &v
	return s
}

func (s *ListCommonCateFirstFloorHeaders) Validate() error {
	return dara.Validate(s)
}
