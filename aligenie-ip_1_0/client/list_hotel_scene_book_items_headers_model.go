// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelSceneBookItemsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListHotelSceneBookItemsHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListHotelSceneBookItemsHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListHotelSceneBookItemsHeaders
	GetAuthorization() *string
}

type ListHotelSceneBookItemsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListHotelSceneBookItemsHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneBookItemsHeaders) GoString() string {
	return s.String()
}

func (s *ListHotelSceneBookItemsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListHotelSceneBookItemsHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListHotelSceneBookItemsHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListHotelSceneBookItemsHeaders) SetCommonHeaders(v map[string]*string) *ListHotelSceneBookItemsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotelSceneBookItemsHeaders) SetXAcsAligenieAccessToken(v string) *ListHotelSceneBookItemsHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListHotelSceneBookItemsHeaders) SetAuthorization(v string) *ListHotelSceneBookItemsHeaders {
	s.Authorization = &v
	return s
}

func (s *ListHotelSceneBookItemsHeaders) Validate() error {
	return dara.Validate(s)
}
