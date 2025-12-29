// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelSceneItemsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListHotelSceneItemsHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListHotelSceneItemsHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListHotelSceneItemsHeaders
	GetAuthorization() *string
}

type ListHotelSceneItemsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListHotelSceneItemsHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneItemsHeaders) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListHotelSceneItemsHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListHotelSceneItemsHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListHotelSceneItemsHeaders) SetCommonHeaders(v map[string]*string) *ListHotelSceneItemsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotelSceneItemsHeaders) SetXAcsAligenieAccessToken(v string) *ListHotelSceneItemsHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListHotelSceneItemsHeaders) SetAuthorization(v string) *ListHotelSceneItemsHeaders {
	s.Authorization = &v
	return s
}

func (s *ListHotelSceneItemsHeaders) Validate() error {
	return dara.Validate(s)
}
