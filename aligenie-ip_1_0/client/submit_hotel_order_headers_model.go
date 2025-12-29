// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitHotelOrderHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SubmitHotelOrderHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *SubmitHotelOrderHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *SubmitHotelOrderHeaders
	GetAuthorization() *string
}

type SubmitHotelOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s SubmitHotelOrderHeaders) String() string {
	return dara.Prettify(s)
}

func (s SubmitHotelOrderHeaders) GoString() string {
	return s.String()
}

func (s *SubmitHotelOrderHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SubmitHotelOrderHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *SubmitHotelOrderHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *SubmitHotelOrderHeaders) SetCommonHeaders(v map[string]*string) *SubmitHotelOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SubmitHotelOrderHeaders) SetXAcsAligenieAccessToken(v string) *SubmitHotelOrderHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *SubmitHotelOrderHeaders) SetAuthorization(v string) *SubmitHotelOrderHeaders {
	s.Authorization = &v
	return s
}

func (s *SubmitHotelOrderHeaders) Validate() error {
	return dara.Validate(s)
}
