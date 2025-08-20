// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBusAppConfigHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetBusAppConfigHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetBusAppConfigHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetBusAppConfigHeaders
	GetAuthorization() *string
}

type GetBusAppConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetBusAppConfigHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetBusAppConfigHeaders) GoString() string {
	return s.String()
}

func (s *GetBusAppConfigHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetBusAppConfigHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetBusAppConfigHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetBusAppConfigHeaders) SetCommonHeaders(v map[string]*string) *GetBusAppConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetBusAppConfigHeaders) SetXAcsAligenieAccessToken(v string) *GetBusAppConfigHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetBusAppConfigHeaders) SetAuthorization(v string) *GetBusAppConfigHeaders {
	s.Authorization = &v
	return s
}

func (s *GetBusAppConfigHeaders) Validate() error {
	return dara.Validate(s)
}
