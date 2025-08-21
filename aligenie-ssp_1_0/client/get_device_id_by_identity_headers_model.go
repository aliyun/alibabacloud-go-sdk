// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceIdByIdentityHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDeviceIdByIdentityHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetDeviceIdByIdentityHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetDeviceIdByIdentityHeaders
	GetAuthorization() *string
}

type GetDeviceIdByIdentityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetDeviceIdByIdentityHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceIdByIdentityHeaders) GoString() string {
	return s.String()
}

func (s *GetDeviceIdByIdentityHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDeviceIdByIdentityHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetDeviceIdByIdentityHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetDeviceIdByIdentityHeaders) SetCommonHeaders(v map[string]*string) *GetDeviceIdByIdentityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeviceIdByIdentityHeaders) SetXAcsAligenieAccessToken(v string) *GetDeviceIdByIdentityHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetDeviceIdByIdentityHeaders) SetAuthorization(v string) *GetDeviceIdByIdentityHeaders {
	s.Authorization = &v
	return s
}

func (s *GetDeviceIdByIdentityHeaders) Validate() error {
	return dara.Validate(s)
}
