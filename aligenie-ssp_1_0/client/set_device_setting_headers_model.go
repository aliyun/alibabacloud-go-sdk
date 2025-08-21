// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDeviceSettingHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SetDeviceSettingHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *SetDeviceSettingHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *SetDeviceSettingHeaders
	GetAuthorization() *string
}

type SetDeviceSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s SetDeviceSettingHeaders) String() string {
	return dara.Prettify(s)
}

func (s SetDeviceSettingHeaders) GoString() string {
	return s.String()
}

func (s *SetDeviceSettingHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SetDeviceSettingHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *SetDeviceSettingHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *SetDeviceSettingHeaders) SetCommonHeaders(v map[string]*string) *SetDeviceSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetDeviceSettingHeaders) SetXAcsAligenieAccessToken(v string) *SetDeviceSettingHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *SetDeviceSettingHeaders) SetAuthorization(v string) *SetDeviceSettingHeaders {
	s.Authorization = &v
	return s
}

func (s *SetDeviceSettingHeaders) Validate() error {
	return dara.Validate(s)
}
