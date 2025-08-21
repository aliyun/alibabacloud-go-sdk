// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceStatusInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDeviceStatusInfoHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetDeviceStatusInfoHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetDeviceStatusInfoHeaders
	GetAuthorization() *string
}

type GetDeviceStatusInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetDeviceStatusInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceStatusInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetDeviceStatusInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDeviceStatusInfoHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetDeviceStatusInfoHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetDeviceStatusInfoHeaders) SetCommonHeaders(v map[string]*string) *GetDeviceStatusInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeviceStatusInfoHeaders) SetXAcsAligenieAccessToken(v string) *GetDeviceStatusInfoHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetDeviceStatusInfoHeaders) SetAuthorization(v string) *GetDeviceStatusInfoHeaders {
	s.Authorization = &v
	return s
}

func (s *GetDeviceStatusInfoHeaders) Validate() error {
	return dara.Validate(s)
}
