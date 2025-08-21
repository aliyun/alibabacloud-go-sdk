// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEcologyOpennessSendVerificationCodeHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *EcologyOpennessSendVerificationCodeHeaders
  GetCommonHeaders() map[string]*string 
  SetXAcsAligenieAccessToken(v string) *EcologyOpennessSendVerificationCodeHeaders
  GetXAcsAligenieAccessToken() *string 
  SetAuthorization(v string) *EcologyOpennessSendVerificationCodeHeaders
  GetAuthorization() *string 
}

type EcologyOpennessSendVerificationCodeHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  XAcsAligenieAccessToken *string `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
  Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s EcologyOpennessSendVerificationCodeHeaders) String() string {
  return dara.Prettify(s)
}

func (s EcologyOpennessSendVerificationCodeHeaders) GoString() string {
  return s.String()
}

func (s *EcologyOpennessSendVerificationCodeHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *EcologyOpennessSendVerificationCodeHeaders) GetXAcsAligenieAccessToken() *string  {
  return s.XAcsAligenieAccessToken
}

func (s *EcologyOpennessSendVerificationCodeHeaders) GetAuthorization() *string  {
  return s.Authorization
}

func (s *EcologyOpennessSendVerificationCodeHeaders) SetCommonHeaders(v map[string]*string) *EcologyOpennessSendVerificationCodeHeaders {
  s.CommonHeaders = v
  return s
}

func (s *EcologyOpennessSendVerificationCodeHeaders) SetXAcsAligenieAccessToken(v string) *EcologyOpennessSendVerificationCodeHeaders {
  s.XAcsAligenieAccessToken = &v
  return s
}

func (s *EcologyOpennessSendVerificationCodeHeaders) SetAuthorization(v string) *EcologyOpennessSendVerificationCodeHeaders {
  s.Authorization = &v
  return s
}

func (s *EcologyOpennessSendVerificationCodeHeaders) Validate() error {
  return dara.Validate(s)
}

