// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserByDeviceIdHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetUserByDeviceIdHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetUserByDeviceIdHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetUserByDeviceIdHeaders
	GetAuthorization() *string
}

type GetUserByDeviceIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetUserByDeviceIdHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetUserByDeviceIdHeaders) GoString() string {
	return s.String()
}

func (s *GetUserByDeviceIdHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetUserByDeviceIdHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetUserByDeviceIdHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetUserByDeviceIdHeaders) SetCommonHeaders(v map[string]*string) *GetUserByDeviceIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserByDeviceIdHeaders) SetXAcsAligenieAccessToken(v string) *GetUserByDeviceIdHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetUserByDeviceIdHeaders) SetAuthorization(v string) *GetUserByDeviceIdHeaders {
	s.Authorization = &v
	return s
}

func (s *GetUserByDeviceIdHeaders) Validate() error {
	return dara.Validate(s)
}
