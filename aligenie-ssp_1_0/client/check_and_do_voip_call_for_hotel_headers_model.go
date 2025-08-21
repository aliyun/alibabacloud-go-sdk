// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAndDoVoipCallForHotelHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CheckAndDoVoipCallForHotelHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *CheckAndDoVoipCallForHotelHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *CheckAndDoVoipCallForHotelHeaders
	GetAuthorization() *string
}

type CheckAndDoVoipCallForHotelHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CheckAndDoVoipCallForHotelHeaders) String() string {
	return dara.Prettify(s)
}

func (s CheckAndDoVoipCallForHotelHeaders) GoString() string {
	return s.String()
}

func (s *CheckAndDoVoipCallForHotelHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CheckAndDoVoipCallForHotelHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *CheckAndDoVoipCallForHotelHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *CheckAndDoVoipCallForHotelHeaders) SetCommonHeaders(v map[string]*string) *CheckAndDoVoipCallForHotelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CheckAndDoVoipCallForHotelHeaders) SetXAcsAligenieAccessToken(v string) *CheckAndDoVoipCallForHotelHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelHeaders) SetAuthorization(v string) *CheckAndDoVoipCallForHotelHeaders {
	s.Authorization = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelHeaders) Validate() error {
	return dara.Validate(s)
}
