// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelOrderHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListHotelOrderHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListHotelOrderHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListHotelOrderHeaders
	GetAuthorization() *string
}

type ListHotelOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListHotelOrderHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListHotelOrderHeaders) GoString() string {
	return s.String()
}

func (s *ListHotelOrderHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListHotelOrderHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListHotelOrderHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListHotelOrderHeaders) SetCommonHeaders(v map[string]*string) *ListHotelOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotelOrderHeaders) SetXAcsAligenieAccessToken(v string) *ListHotelOrderHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListHotelOrderHeaders) SetAuthorization(v string) *ListHotelOrderHeaders {
	s.Authorization = &v
	return s
}

func (s *ListHotelOrderHeaders) Validate() error {
	return dara.Validate(s)
}
