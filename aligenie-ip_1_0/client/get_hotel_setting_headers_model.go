// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelSettingHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetHotelSettingHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetHotelSettingHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetHotelSettingHeaders
	GetAuthorization() *string
}

type GetHotelSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelSettingHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetHotelSettingHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelSettingHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetHotelSettingHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetHotelSettingHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetHotelSettingHeaders) SetCommonHeaders(v map[string]*string) *GetHotelSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelSettingHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelSettingHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelSettingHeaders) SetAuthorization(v string) *GetHotelSettingHeaders {
	s.Authorization = &v
	return s
}

func (s *GetHotelSettingHeaders) Validate() error {
	return dara.Validate(s)
}
