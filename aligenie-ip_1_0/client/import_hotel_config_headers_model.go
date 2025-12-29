// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportHotelConfigHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ImportHotelConfigHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ImportHotelConfigHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ImportHotelConfigHeaders
	GetAuthorization() *string
}

type ImportHotelConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ImportHotelConfigHeaders) String() string {
	return dara.Prettify(s)
}

func (s ImportHotelConfigHeaders) GoString() string {
	return s.String()
}

func (s *ImportHotelConfigHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ImportHotelConfigHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ImportHotelConfigHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ImportHotelConfigHeaders) SetCommonHeaders(v map[string]*string) *ImportHotelConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ImportHotelConfigHeaders) SetXAcsAligenieAccessToken(v string) *ImportHotelConfigHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ImportHotelConfigHeaders) SetAuthorization(v string) *ImportHotelConfigHeaders {
	s.Authorization = &v
	return s
}

func (s *ImportHotelConfigHeaders) Validate() error {
	return dara.Validate(s)
}
