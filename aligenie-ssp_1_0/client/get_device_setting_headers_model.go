// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceSettingHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDeviceSettingHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetDeviceSettingHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetDeviceSettingHeaders
	GetAuthorization() *string
}

type GetDeviceSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetDeviceSettingHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceSettingHeaders) GoString() string {
	return s.String()
}

func (s *GetDeviceSettingHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDeviceSettingHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetDeviceSettingHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetDeviceSettingHeaders) SetCommonHeaders(v map[string]*string) *GetDeviceSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeviceSettingHeaders) SetXAcsAligenieAccessToken(v string) *GetDeviceSettingHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetDeviceSettingHeaders) SetAuthorization(v string) *GetDeviceSettingHeaders {
	s.Authorization = &v
	return s
}

func (s *GetDeviceSettingHeaders) Validate() error {
	return dara.Validate(s)
}
