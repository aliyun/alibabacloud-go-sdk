// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserDeviceListByTmeUserIdHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryUserDeviceListByTmeUserIdHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *QueryUserDeviceListByTmeUserIdHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *QueryUserDeviceListByTmeUserIdHeaders
	GetAuthorization() *string
}

type QueryUserDeviceListByTmeUserIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s QueryUserDeviceListByTmeUserIdHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryUserDeviceListByTmeUserIdHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserDeviceListByTmeUserIdHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryUserDeviceListByTmeUserIdHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *QueryUserDeviceListByTmeUserIdHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *QueryUserDeviceListByTmeUserIdHeaders) SetCommonHeaders(v map[string]*string) *QueryUserDeviceListByTmeUserIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserDeviceListByTmeUserIdHeaders) SetXAcsAligenieAccessToken(v string) *QueryUserDeviceListByTmeUserIdHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *QueryUserDeviceListByTmeUserIdHeaders) SetAuthorization(v string) *QueryUserDeviceListByTmeUserIdHeaders {
	s.Authorization = &v
	return s
}

func (s *QueryUserDeviceListByTmeUserIdHeaders) Validate() error {
	return dara.Validate(s)
}
