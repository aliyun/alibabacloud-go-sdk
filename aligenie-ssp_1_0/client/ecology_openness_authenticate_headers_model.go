// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEcologyOpennessAuthenticateHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *EcologyOpennessAuthenticateHeaders
  GetCommonHeaders() map[string]*string 
  SetXAcsAligenieAccessToken(v string) *EcologyOpennessAuthenticateHeaders
  GetXAcsAligenieAccessToken() *string 
  SetAuthorization(v string) *EcologyOpennessAuthenticateHeaders
  GetAuthorization() *string 
}

type EcologyOpennessAuthenticateHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  XAcsAligenieAccessToken *string `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
  Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s EcologyOpennessAuthenticateHeaders) String() string {
  return dara.Prettify(s)
}

func (s EcologyOpennessAuthenticateHeaders) GoString() string {
  return s.String()
}

func (s *EcologyOpennessAuthenticateHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *EcologyOpennessAuthenticateHeaders) GetXAcsAligenieAccessToken() *string  {
  return s.XAcsAligenieAccessToken
}

func (s *EcologyOpennessAuthenticateHeaders) GetAuthorization() *string  {
  return s.Authorization
}

func (s *EcologyOpennessAuthenticateHeaders) SetCommonHeaders(v map[string]*string) *EcologyOpennessAuthenticateHeaders {
  s.CommonHeaders = v
  return s
}

func (s *EcologyOpennessAuthenticateHeaders) SetXAcsAligenieAccessToken(v string) *EcologyOpennessAuthenticateHeaders {
  s.XAcsAligenieAccessToken = &v
  return s
}

func (s *EcologyOpennessAuthenticateHeaders) SetAuthorization(v string) *EcologyOpennessAuthenticateHeaders {
  s.Authorization = &v
  return s
}

func (s *EcologyOpennessAuthenticateHeaders) Validate() error {
  return dara.Validate(s)
}

