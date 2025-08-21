// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeviceControlHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeviceControlHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *DeviceControlHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *DeviceControlHeaders
	GetAuthorization() *string
}

type DeviceControlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeviceControlHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeviceControlHeaders) GoString() string {
	return s.String()
}

func (s *DeviceControlHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeviceControlHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *DeviceControlHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *DeviceControlHeaders) SetCommonHeaders(v map[string]*string) *DeviceControlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeviceControlHeaders) SetXAcsAligenieAccessToken(v string) *DeviceControlHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *DeviceControlHeaders) SetAuthorization(v string) *DeviceControlHeaders {
	s.Authorization = &v
	return s
}

func (s *DeviceControlHeaders) Validate() error {
	return dara.Validate(s)
}
