// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHotelHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateHotelHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *CreateHotelHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *CreateHotelHeaders
	GetAuthorization() *string
}

type CreateHotelHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CreateHotelHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateHotelHeaders) GoString() string {
	return s.String()
}

func (s *CreateHotelHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateHotelHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *CreateHotelHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *CreateHotelHeaders) SetCommonHeaders(v map[string]*string) *CreateHotelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateHotelHeaders) SetXAcsAligenieAccessToken(v string) *CreateHotelHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *CreateHotelHeaders) SetAuthorization(v string) *CreateHotelHeaders {
	s.Authorization = &v
	return s
}

func (s *CreateHotelHeaders) Validate() error {
	return dara.Validate(s)
}
