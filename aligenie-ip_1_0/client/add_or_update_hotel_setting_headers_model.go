// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrUpdateHotelSettingHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddOrUpdateHotelSettingHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *AddOrUpdateHotelSettingHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *AddOrUpdateHotelSettingHeaders
	GetAuthorization() *string
}

type AddOrUpdateHotelSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AddOrUpdateHotelSettingHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddOrUpdateHotelSettingHeaders) GoString() string {
	return s.String()
}

func (s *AddOrUpdateHotelSettingHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddOrUpdateHotelSettingHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *AddOrUpdateHotelSettingHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *AddOrUpdateHotelSettingHeaders) SetCommonHeaders(v map[string]*string) *AddOrUpdateHotelSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddOrUpdateHotelSettingHeaders) SetXAcsAligenieAccessToken(v string) *AddOrUpdateHotelSettingHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AddOrUpdateHotelSettingHeaders) SetAuthorization(v string) *AddOrUpdateHotelSettingHeaders {
	s.Authorization = &v
	return s
}

func (s *AddOrUpdateHotelSettingHeaders) Validate() error {
	return dara.Validate(s)
}
