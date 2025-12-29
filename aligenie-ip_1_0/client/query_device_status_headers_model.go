// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceStatusHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryDeviceStatusHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *QueryDeviceStatusHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *QueryDeviceStatusHeaders
	GetAuthorization() *string
}

type QueryDeviceStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s QueryDeviceStatusHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceStatusHeaders) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryDeviceStatusHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *QueryDeviceStatusHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *QueryDeviceStatusHeaders) SetCommonHeaders(v map[string]*string) *QueryDeviceStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDeviceStatusHeaders) SetXAcsAligenieAccessToken(v string) *QueryDeviceStatusHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *QueryDeviceStatusHeaders) SetAuthorization(v string) *QueryDeviceStatusHeaders {
	s.Authorization = &v
	return s
}

func (s *QueryDeviceStatusHeaders) Validate() error {
	return dara.Validate(s)
}
