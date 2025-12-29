// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelOrderDetailHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetHotelOrderDetailHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetHotelOrderDetailHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetHotelOrderDetailHeaders
	GetAuthorization() *string
}

type GetHotelOrderDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelOrderDetailHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetHotelOrderDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelOrderDetailHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetHotelOrderDetailHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetHotelOrderDetailHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetHotelOrderDetailHeaders) SetCommonHeaders(v map[string]*string) *GetHotelOrderDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelOrderDetailHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelOrderDetailHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelOrderDetailHeaders) SetAuthorization(v string) *GetHotelOrderDetailHeaders {
	s.Authorization = &v
	return s
}

func (s *GetHotelOrderDetailHeaders) Validate() error {
	return dara.Validate(s)
}
