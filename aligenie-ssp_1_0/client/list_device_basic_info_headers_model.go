// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceBasicInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListDeviceBasicInfoHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListDeviceBasicInfoHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListDeviceBasicInfoHeaders
	GetAuthorization() *string
}

type ListDeviceBasicInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListDeviceBasicInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceBasicInfoHeaders) GoString() string {
	return s.String()
}

func (s *ListDeviceBasicInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListDeviceBasicInfoHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListDeviceBasicInfoHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListDeviceBasicInfoHeaders) SetCommonHeaders(v map[string]*string) *ListDeviceBasicInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDeviceBasicInfoHeaders) SetXAcsAligenieAccessToken(v string) *ListDeviceBasicInfoHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListDeviceBasicInfoHeaders) SetAuthorization(v string) *ListDeviceBasicInfoHeaders {
	s.Authorization = &v
	return s
}

func (s *ListDeviceBasicInfoHeaders) Validate() error {
	return dara.Validate(s)
}
