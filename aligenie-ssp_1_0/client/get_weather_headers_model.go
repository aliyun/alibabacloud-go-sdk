// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWeatherHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetWeatherHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetWeatherHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetWeatherHeaders
	GetAuthorization() *string
}

type GetWeatherHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetWeatherHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetWeatherHeaders) GoString() string {
	return s.String()
}

func (s *GetWeatherHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetWeatherHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetWeatherHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetWeatherHeaders) SetCommonHeaders(v map[string]*string) *GetWeatherHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetWeatherHeaders) SetXAcsAligenieAccessToken(v string) *GetWeatherHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetWeatherHeaders) SetAuthorization(v string) *GetWeatherHeaders {
	s.Authorization = &v
	return s
}

func (s *GetWeatherHeaders) Validate() error {
	return dara.Validate(s)
}
