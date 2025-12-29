// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelMessageTemplateHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListHotelMessageTemplateHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListHotelMessageTemplateHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListHotelMessageTemplateHeaders
	GetAuthorization() *string
}

type ListHotelMessageTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListHotelMessageTemplateHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListHotelMessageTemplateHeaders) GoString() string {
	return s.String()
}

func (s *ListHotelMessageTemplateHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListHotelMessageTemplateHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListHotelMessageTemplateHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListHotelMessageTemplateHeaders) SetCommonHeaders(v map[string]*string) *ListHotelMessageTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotelMessageTemplateHeaders) SetXAcsAligenieAccessToken(v string) *ListHotelMessageTemplateHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListHotelMessageTemplateHeaders) SetAuthorization(v string) *ListHotelMessageTemplateHeaders {
	s.Authorization = &v
	return s
}

func (s *ListHotelMessageTemplateHeaders) Validate() error {
	return dara.Validate(s)
}
