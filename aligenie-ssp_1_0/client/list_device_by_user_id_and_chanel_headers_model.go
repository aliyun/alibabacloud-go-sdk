// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceByUserIdAndChanelHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListDeviceByUserIdAndChanelHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListDeviceByUserIdAndChanelHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListDeviceByUserIdAndChanelHeaders
	GetAuthorization() *string
}

type ListDeviceByUserIdAndChanelHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListDeviceByUserIdAndChanelHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceByUserIdAndChanelHeaders) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdAndChanelHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListDeviceByUserIdAndChanelHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListDeviceByUserIdAndChanelHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListDeviceByUserIdAndChanelHeaders) SetCommonHeaders(v map[string]*string) *ListDeviceByUserIdAndChanelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDeviceByUserIdAndChanelHeaders) SetXAcsAligenieAccessToken(v string) *ListDeviceByUserIdAndChanelHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelHeaders) SetAuthorization(v string) *ListDeviceByUserIdAndChanelHeaders {
	s.Authorization = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelHeaders) Validate() error {
	return dara.Validate(s)
}
