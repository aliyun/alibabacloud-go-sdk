// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushHotelMessageHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PushHotelMessageHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *PushHotelMessageHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *PushHotelMessageHeaders
	GetAuthorization() *string
}

type PushHotelMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PushHotelMessageHeaders) String() string {
	return dara.Prettify(s)
}

func (s PushHotelMessageHeaders) GoString() string {
	return s.String()
}

func (s *PushHotelMessageHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PushHotelMessageHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *PushHotelMessageHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *PushHotelMessageHeaders) SetCommonHeaders(v map[string]*string) *PushHotelMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PushHotelMessageHeaders) SetXAcsAligenieAccessToken(v string) *PushHotelMessageHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *PushHotelMessageHeaders) SetAuthorization(v string) *PushHotelMessageHeaders {
	s.Authorization = &v
	return s
}

func (s *PushHotelMessageHeaders) Validate() error {
	return dara.Validate(s)
}
