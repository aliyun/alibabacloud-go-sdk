// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindDeviceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UnbindDeviceHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *UnbindDeviceHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *UnbindDeviceHeaders
	GetAuthorization() *string
}

type UnbindDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UnbindDeviceHeaders) String() string {
	return dara.Prettify(s)
}

func (s UnbindDeviceHeaders) GoString() string {
	return s.String()
}

func (s *UnbindDeviceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UnbindDeviceHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *UnbindDeviceHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *UnbindDeviceHeaders) SetCommonHeaders(v map[string]*string) *UnbindDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UnbindDeviceHeaders) SetXAcsAligenieAccessToken(v string) *UnbindDeviceHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UnbindDeviceHeaders) SetAuthorization(v string) *UnbindDeviceHeaders {
	s.Authorization = &v
	return s
}

func (s *UnbindDeviceHeaders) Validate() error {
	return dara.Validate(s)
}
