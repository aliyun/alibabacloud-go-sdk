// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelSceneItemHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListHotelSceneItemHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListHotelSceneItemHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListHotelSceneItemHeaders
	GetAuthorization() *string
}

type ListHotelSceneItemHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListHotelSceneItemHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneItemHeaders) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListHotelSceneItemHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListHotelSceneItemHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListHotelSceneItemHeaders) SetCommonHeaders(v map[string]*string) *ListHotelSceneItemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotelSceneItemHeaders) SetXAcsAligenieAccessToken(v string) *ListHotelSceneItemHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListHotelSceneItemHeaders) SetAuthorization(v string) *ListHotelSceneItemHeaders {
	s.Authorization = &v
	return s
}

func (s *ListHotelSceneItemHeaders) Validate() error {
	return dara.Validate(s)
}
