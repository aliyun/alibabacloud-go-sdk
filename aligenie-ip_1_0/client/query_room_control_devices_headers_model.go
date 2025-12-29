// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRoomControlDevicesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryRoomControlDevicesHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *QueryRoomControlDevicesHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *QueryRoomControlDevicesHeaders
	GetAuthorization() *string
}

type QueryRoomControlDevicesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s QueryRoomControlDevicesHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryRoomControlDevicesHeaders) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryRoomControlDevicesHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *QueryRoomControlDevicesHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *QueryRoomControlDevicesHeaders) SetCommonHeaders(v map[string]*string) *QueryRoomControlDevicesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryRoomControlDevicesHeaders) SetXAcsAligenieAccessToken(v string) *QueryRoomControlDevicesHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *QueryRoomControlDevicesHeaders) SetAuthorization(v string) *QueryRoomControlDevicesHeaders {
	s.Authorization = &v
	return s
}

func (s *QueryRoomControlDevicesHeaders) Validate() error {
	return dara.Validate(s)
}
