// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteSceneHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *ExecuteSceneHeaders
  GetCommonHeaders() map[string]*string 
  SetXAcsAligenieAccessToken(v string) *ExecuteSceneHeaders
  GetXAcsAligenieAccessToken() *string 
  SetAuthorization(v string) *ExecuteSceneHeaders
  GetAuthorization() *string 
}

type ExecuteSceneHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  XAcsAligenieAccessToken *string `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
  Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ExecuteSceneHeaders) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSceneHeaders) GoString() string {
  return s.String()
}

func (s *ExecuteSceneHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *ExecuteSceneHeaders) GetXAcsAligenieAccessToken() *string  {
  return s.XAcsAligenieAccessToken
}

func (s *ExecuteSceneHeaders) GetAuthorization() *string  {
  return s.Authorization
}

func (s *ExecuteSceneHeaders) SetCommonHeaders(v map[string]*string) *ExecuteSceneHeaders {
  s.CommonHeaders = v
  return s
}

func (s *ExecuteSceneHeaders) SetXAcsAligenieAccessToken(v string) *ExecuteSceneHeaders {
  s.XAcsAligenieAccessToken = &v
  return s
}

func (s *ExecuteSceneHeaders) SetAuthorization(v string) *ExecuteSceneHeaders {
  s.Authorization = &v
  return s
}

func (s *ExecuteSceneHeaders) Validate() error {
  return dara.Validate(s)
}

