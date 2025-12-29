// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelNoticeHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetHotelNoticeHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetHotelNoticeHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetHotelNoticeHeaders
	GetAuthorization() *string
}

type GetHotelNoticeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelNoticeHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetHotelNoticeHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelNoticeHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetHotelNoticeHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetHotelNoticeHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetHotelNoticeHeaders) SetCommonHeaders(v map[string]*string) *GetHotelNoticeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelNoticeHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelNoticeHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelNoticeHeaders) SetAuthorization(v string) *GetHotelNoticeHeaders {
	s.Authorization = &v
	return s
}

func (s *GetHotelNoticeHeaders) Validate() error {
	return dara.Validate(s)
}
