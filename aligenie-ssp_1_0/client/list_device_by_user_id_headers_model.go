// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceByUserIdHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListDeviceByUserIdHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListDeviceByUserIdHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListDeviceByUserIdHeaders
	GetAuthorization() *string
}

type ListDeviceByUserIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListDeviceByUserIdHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceByUserIdHeaders) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListDeviceByUserIdHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListDeviceByUserIdHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListDeviceByUserIdHeaders) SetCommonHeaders(v map[string]*string) *ListDeviceByUserIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDeviceByUserIdHeaders) SetXAcsAligenieAccessToken(v string) *ListDeviceByUserIdHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListDeviceByUserIdHeaders) SetAuthorization(v string) *ListDeviceByUserIdHeaders {
	s.Authorization = &v
	return s
}

func (s *ListDeviceByUserIdHeaders) Validate() error {
	return dara.Validate(s)
}
