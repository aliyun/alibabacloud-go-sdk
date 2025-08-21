// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceBasicInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDeviceBasicInfoHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetDeviceBasicInfoHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetDeviceBasicInfoHeaders
	GetAuthorization() *string
}

type GetDeviceBasicInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetDeviceBasicInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceBasicInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetDeviceBasicInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDeviceBasicInfoHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetDeviceBasicInfoHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetDeviceBasicInfoHeaders) SetCommonHeaders(v map[string]*string) *GetDeviceBasicInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeviceBasicInfoHeaders) SetXAcsAligenieAccessToken(v string) *GetDeviceBasicInfoHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetDeviceBasicInfoHeaders) SetAuthorization(v string) *GetDeviceBasicInfoHeaders {
	s.Authorization = &v
	return s
}

func (s *GetDeviceBasicInfoHeaders) Validate() error {
	return dara.Validate(s)
}
