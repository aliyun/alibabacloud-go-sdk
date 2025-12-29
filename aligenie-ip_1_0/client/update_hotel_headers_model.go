// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotelHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateHotelHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *UpdateHotelHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *UpdateHotelHeaders
	GetAuthorization() *string
}

type UpdateHotelHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UpdateHotelHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelHeaders) GoString() string {
	return s.String()
}

func (s *UpdateHotelHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateHotelHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *UpdateHotelHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *UpdateHotelHeaders) SetCommonHeaders(v map[string]*string) *UpdateHotelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateHotelHeaders) SetXAcsAligenieAccessToken(v string) *UpdateHotelHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UpdateHotelHeaders) SetAuthorization(v string) *UpdateHotelHeaders {
	s.Authorization = &v
	return s
}

func (s *UpdateHotelHeaders) Validate() error {
	return dara.Validate(s)
}
