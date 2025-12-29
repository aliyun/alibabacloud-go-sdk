// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRoomCheckOutHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *RoomCheckOutHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *RoomCheckOutHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *RoomCheckOutHeaders
	GetAuthorization() *string
}

type RoomCheckOutHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s RoomCheckOutHeaders) String() string {
	return dara.Prettify(s)
}

func (s RoomCheckOutHeaders) GoString() string {
	return s.String()
}

func (s *RoomCheckOutHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *RoomCheckOutHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *RoomCheckOutHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *RoomCheckOutHeaders) SetCommonHeaders(v map[string]*string) *RoomCheckOutHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RoomCheckOutHeaders) SetXAcsAligenieAccessToken(v string) *RoomCheckOutHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *RoomCheckOutHeaders) SetAuthorization(v string) *RoomCheckOutHeaders {
	s.Authorization = &v
	return s
}

func (s *RoomCheckOutHeaders) Validate() error {
	return dara.Validate(s)
}
