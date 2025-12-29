// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRoomStatusHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryRoomStatusHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *QueryRoomStatusHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *QueryRoomStatusHeaders
	GetAuthorization() *string
}

type QueryRoomStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s QueryRoomStatusHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryRoomStatusHeaders) GoString() string {
	return s.String()
}

func (s *QueryRoomStatusHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryRoomStatusHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *QueryRoomStatusHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *QueryRoomStatusHeaders) SetCommonHeaders(v map[string]*string) *QueryRoomStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryRoomStatusHeaders) SetXAcsAligenieAccessToken(v string) *QueryRoomStatusHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *QueryRoomStatusHeaders) SetAuthorization(v string) *QueryRoomStatusHeaders {
	s.Authorization = &v
	return s
}

func (s *QueryRoomStatusHeaders) Validate() error {
	return dara.Validate(s)
}
