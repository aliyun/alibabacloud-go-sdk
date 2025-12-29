// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelSampleUtterancesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetHotelSampleUtterancesHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetHotelSampleUtterancesHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetHotelSampleUtterancesHeaders
	GetAuthorization() *string
}

type GetHotelSampleUtterancesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelSampleUtterancesHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetHotelSampleUtterancesHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelSampleUtterancesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetHotelSampleUtterancesHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetHotelSampleUtterancesHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetHotelSampleUtterancesHeaders) SetCommonHeaders(v map[string]*string) *GetHotelSampleUtterancesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelSampleUtterancesHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelSampleUtterancesHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelSampleUtterancesHeaders) SetAuthorization(v string) *GetHotelSampleUtterancesHeaders {
	s.Authorization = &v
	return s
}

func (s *GetHotelSampleUtterancesHeaders) Validate() error {
	return dara.Validate(s)
}
