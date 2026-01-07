// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountSeparateEaResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseAccountSeparateEaResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseAccountSeparateEaResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseAccountSeparateEaResponseBody) *EnterpriseAccountSeparateEaResponse
  GetBody() *EnterpriseAccountSeparateEaResponseBody 
}

type EnterpriseAccountSeparateEaResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseAccountSeparateEaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountSeparateEaResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountSeparateEaResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountSeparateEaResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseAccountSeparateEaResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseAccountSeparateEaResponse) GetBody() *EnterpriseAccountSeparateEaResponseBody  {
  return s.Body
}

func (s *EnterpriseAccountSeparateEaResponse) SetHeaders(v map[string]*string) *EnterpriseAccountSeparateEaResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseAccountSeparateEaResponse) SetStatusCode(v int32) *EnterpriseAccountSeparateEaResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseAccountSeparateEaResponse) SetBody(v *EnterpriseAccountSeparateEaResponseBody) *EnterpriseAccountSeparateEaResponse {
  s.Body = v
  return s
}

func (s *EnterpriseAccountSeparateEaResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

