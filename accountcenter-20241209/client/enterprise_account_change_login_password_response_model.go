// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountChangeLoginPasswordResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseAccountChangeLoginPasswordResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseAccountChangeLoginPasswordResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseAccountChangeLoginPasswordResponseBody) *EnterpriseAccountChangeLoginPasswordResponse
  GetBody() *EnterpriseAccountChangeLoginPasswordResponseBody 
}

type EnterpriseAccountChangeLoginPasswordResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseAccountChangeLoginPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountChangeLoginPasswordResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountChangeLoginPasswordResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountChangeLoginPasswordResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseAccountChangeLoginPasswordResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseAccountChangeLoginPasswordResponse) GetBody() *EnterpriseAccountChangeLoginPasswordResponseBody  {
  return s.Body
}

func (s *EnterpriseAccountChangeLoginPasswordResponse) SetHeaders(v map[string]*string) *EnterpriseAccountChangeLoginPasswordResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseAccountChangeLoginPasswordResponse) SetStatusCode(v int32) *EnterpriseAccountChangeLoginPasswordResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseAccountChangeLoginPasswordResponse) SetBody(v *EnterpriseAccountChangeLoginPasswordResponseBody) *EnterpriseAccountChangeLoginPasswordResponse {
  s.Body = v
  return s
}

func (s *EnterpriseAccountChangeLoginPasswordResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

