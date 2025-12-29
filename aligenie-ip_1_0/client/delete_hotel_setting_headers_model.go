// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHotelSettingHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteHotelSettingHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *DeleteHotelSettingHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *DeleteHotelSettingHeaders
	GetAuthorization() *string
}

type DeleteHotelSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteHotelSettingHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteHotelSettingHeaders) GoString() string {
	return s.String()
}

func (s *DeleteHotelSettingHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteHotelSettingHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *DeleteHotelSettingHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *DeleteHotelSettingHeaders) SetCommonHeaders(v map[string]*string) *DeleteHotelSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteHotelSettingHeaders) SetXAcsAligenieAccessToken(v string) *DeleteHotelSettingHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *DeleteHotelSettingHeaders) SetAuthorization(v string) *DeleteHotelSettingHeaders {
	s.Authorization = &v
	return s
}

func (s *DeleteHotelSettingHeaders) Validate() error {
	return dara.Validate(s)
}
