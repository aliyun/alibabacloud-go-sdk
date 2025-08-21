// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceStatusDetailHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDeviceStatusDetailHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetDeviceStatusDetailHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetDeviceStatusDetailHeaders
	GetAuthorization() *string
}

type GetDeviceStatusDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetDeviceStatusDetailHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceStatusDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusDetailHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDeviceStatusDetailHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetDeviceStatusDetailHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetDeviceStatusDetailHeaders) SetCommonHeaders(v map[string]*string) *GetDeviceStatusDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeviceStatusDetailHeaders) SetXAcsAligenieAccessToken(v string) *GetDeviceStatusDetailHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetDeviceStatusDetailHeaders) SetAuthorization(v string) *GetDeviceStatusDetailHeaders {
	s.Authorization = &v
	return s
}

func (s *GetDeviceStatusDetailHeaders) Validate() error {
	return dara.Validate(s)
}
