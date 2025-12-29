// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRoomControlDevicesAndStatusHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryRoomControlDevicesAndStatusHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *QueryRoomControlDevicesAndStatusHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *QueryRoomControlDevicesAndStatusHeaders
	GetAuthorization() *string
}

type QueryRoomControlDevicesAndStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s QueryRoomControlDevicesAndStatusHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryRoomControlDevicesAndStatusHeaders) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesAndStatusHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryRoomControlDevicesAndStatusHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *QueryRoomControlDevicesAndStatusHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *QueryRoomControlDevicesAndStatusHeaders) SetCommonHeaders(v map[string]*string) *QueryRoomControlDevicesAndStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryRoomControlDevicesAndStatusHeaders) SetXAcsAligenieAccessToken(v string) *QueryRoomControlDevicesAndStatusHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusHeaders) SetAuthorization(v string) *QueryRoomControlDevicesAndStatusHeaders {
	s.Authorization = &v
	return s
}

func (s *QueryRoomControlDevicesAndStatusHeaders) Validate() error {
	return dara.Validate(s)
}
