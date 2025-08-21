// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceTagHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDeviceTagHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetDeviceTagHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetDeviceTagHeaders
	GetAuthorization() *string
}

type GetDeviceTagHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetDeviceTagHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceTagHeaders) GoString() string {
	return s.String()
}

func (s *GetDeviceTagHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDeviceTagHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetDeviceTagHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetDeviceTagHeaders) SetCommonHeaders(v map[string]*string) *GetDeviceTagHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeviceTagHeaders) SetXAcsAligenieAccessToken(v string) *GetDeviceTagHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetDeviceTagHeaders) SetAuthorization(v string) *GetDeviceTagHeaders {
	s.Authorization = &v
	return s
}

func (s *GetDeviceTagHeaders) Validate() error {
	return dara.Validate(s)
}
