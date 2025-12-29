// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelHomeBackImageAndModesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetHotelHomeBackImageAndModesHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetHotelHomeBackImageAndModesHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetHotelHomeBackImageAndModesHeaders
	GetAuthorization() *string
}

type GetHotelHomeBackImageAndModesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelHomeBackImageAndModesHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetHotelHomeBackImageAndModesHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelHomeBackImageAndModesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetHotelHomeBackImageAndModesHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetHotelHomeBackImageAndModesHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetHotelHomeBackImageAndModesHeaders) SetCommonHeaders(v map[string]*string) *GetHotelHomeBackImageAndModesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelHomeBackImageAndModesHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelHomeBackImageAndModesHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesHeaders) SetAuthorization(v string) *GetHotelHomeBackImageAndModesHeaders {
	s.Authorization = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesHeaders) Validate() error {
	return dara.Validate(s)
}
