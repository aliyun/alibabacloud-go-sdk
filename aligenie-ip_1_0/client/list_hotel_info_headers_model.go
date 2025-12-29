// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListHotelInfoHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListHotelInfoHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListHotelInfoHeaders
	GetAuthorization() *string
}

type ListHotelInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListHotelInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListHotelInfoHeaders) GoString() string {
	return s.String()
}

func (s *ListHotelInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListHotelInfoHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListHotelInfoHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListHotelInfoHeaders) SetCommonHeaders(v map[string]*string) *ListHotelInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotelInfoHeaders) SetXAcsAligenieAccessToken(v string) *ListHotelInfoHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListHotelInfoHeaders) SetAuthorization(v string) *ListHotelInfoHeaders {
	s.Authorization = &v
	return s
}

func (s *ListHotelInfoHeaders) Validate() error {
	return dara.Validate(s)
}
