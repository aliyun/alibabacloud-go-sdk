// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountRemoveMfaResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseAccountRemoveMfaResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseAccountRemoveMfaResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseAccountRemoveMfaResponseBody) *EnterpriseAccountRemoveMfaResponse
  GetBody() *EnterpriseAccountRemoveMfaResponseBody 
}

type EnterpriseAccountRemoveMfaResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseAccountRemoveMfaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountRemoveMfaResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountRemoveMfaResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountRemoveMfaResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseAccountRemoveMfaResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseAccountRemoveMfaResponse) GetBody() *EnterpriseAccountRemoveMfaResponseBody  {
  return s.Body
}

func (s *EnterpriseAccountRemoveMfaResponse) SetHeaders(v map[string]*string) *EnterpriseAccountRemoveMfaResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseAccountRemoveMfaResponse) SetStatusCode(v int32) *EnterpriseAccountRemoveMfaResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseAccountRemoveMfaResponse) SetBody(v *EnterpriseAccountRemoveMfaResponseBody) *EnterpriseAccountRemoveMfaResponse {
  s.Body = v
  return s
}

func (s *EnterpriseAccountRemoveMfaResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

