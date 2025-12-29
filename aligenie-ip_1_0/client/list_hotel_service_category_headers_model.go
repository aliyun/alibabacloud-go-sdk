// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelServiceCategoryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListHotelServiceCategoryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListHotelServiceCategoryHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListHotelServiceCategoryHeaders
	GetAuthorization() *string
}

type ListHotelServiceCategoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListHotelServiceCategoryHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListHotelServiceCategoryHeaders) GoString() string {
	return s.String()
}

func (s *ListHotelServiceCategoryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListHotelServiceCategoryHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListHotelServiceCategoryHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListHotelServiceCategoryHeaders) SetCommonHeaders(v map[string]*string) *ListHotelServiceCategoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotelServiceCategoryHeaders) SetXAcsAligenieAccessToken(v string) *ListHotelServiceCategoryHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListHotelServiceCategoryHeaders) SetAuthorization(v string) *ListHotelServiceCategoryHeaders {
	s.Authorization = &v
	return s
}

func (s *ListHotelServiceCategoryHeaders) Validate() error {
	return dara.Validate(s)
}
