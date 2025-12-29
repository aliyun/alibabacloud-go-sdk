// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveHotelHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *RemoveHotelHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *RemoveHotelHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *RemoveHotelHeaders
	GetAuthorization() *string
}

type RemoveHotelHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s RemoveHotelHeaders) String() string {
	return dara.Prettify(s)
}

func (s RemoveHotelHeaders) GoString() string {
	return s.String()
}

func (s *RemoveHotelHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *RemoveHotelHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *RemoveHotelHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *RemoveHotelHeaders) SetCommonHeaders(v map[string]*string) *RemoveHotelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveHotelHeaders) SetXAcsAligenieAccessToken(v string) *RemoveHotelHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *RemoveHotelHeaders) SetAuthorization(v string) *RemoveHotelHeaders {
	s.Authorization = &v
	return s
}

func (s *RemoveHotelHeaders) Validate() error {
	return dara.Validate(s)
}
