// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelNoticeV2Headers interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetHotelNoticeV2Headers
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetHotelNoticeV2Headers
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetHotelNoticeV2Headers
	GetAuthorization() *string
}

type GetHotelNoticeV2Headers struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelNoticeV2Headers) String() string {
	return dara.Prettify(s)
}

func (s GetHotelNoticeV2Headers) GoString() string {
	return s.String()
}

func (s *GetHotelNoticeV2Headers) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetHotelNoticeV2Headers) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetHotelNoticeV2Headers) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetHotelNoticeV2Headers) SetCommonHeaders(v map[string]*string) *GetHotelNoticeV2Headers {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelNoticeV2Headers) SetXAcsAligenieAccessToken(v string) *GetHotelNoticeV2Headers {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelNoticeV2Headers) SetAuthorization(v string) *GetHotelNoticeV2Headers {
	s.Authorization = &v
	return s
}

func (s *GetHotelNoticeV2Headers) Validate() error {
	return dara.Validate(s)
}
