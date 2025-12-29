// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iControlRoomDeviceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ControlRoomDeviceHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ControlRoomDeviceHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ControlRoomDeviceHeaders
	GetAuthorization() *string
}

type ControlRoomDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ControlRoomDeviceHeaders) String() string {
	return dara.Prettify(s)
}

func (s ControlRoomDeviceHeaders) GoString() string {
	return s.String()
}

func (s *ControlRoomDeviceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ControlRoomDeviceHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ControlRoomDeviceHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ControlRoomDeviceHeaders) SetCommonHeaders(v map[string]*string) *ControlRoomDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ControlRoomDeviceHeaders) SetXAcsAligenieAccessToken(v string) *ControlRoomDeviceHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ControlRoomDeviceHeaders) SetAuthorization(v string) *ControlRoomDeviceHeaders {
	s.Authorization = &v
	return s
}

func (s *ControlRoomDeviceHeaders) Validate() error {
	return dara.Validate(s)
}
