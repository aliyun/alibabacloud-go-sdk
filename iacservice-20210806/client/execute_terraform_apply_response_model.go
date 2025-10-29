// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTerraformApplyResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteTerraformApplyResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteTerraformApplyResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteTerraformApplyResponseBody) *ExecuteTerraformApplyResponse
  GetBody() *ExecuteTerraformApplyResponseBody 
}

type ExecuteTerraformApplyResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteTerraformApplyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteTerraformApplyResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTerraformApplyResponse) GoString() string {
  return s.String()
}

func (s *ExecuteTerraformApplyResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteTerraformApplyResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteTerraformApplyResponse) GetBody() *ExecuteTerraformApplyResponseBody  {
  return s.Body
}

func (s *ExecuteTerraformApplyResponse) SetHeaders(v map[string]*string) *ExecuteTerraformApplyResponse {
  s.Headers = v
  return s
}

func (s *ExecuteTerraformApplyResponse) SetStatusCode(v int32) *ExecuteTerraformApplyResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteTerraformApplyResponse) SetBody(v *ExecuteTerraformApplyResponseBody) *ExecuteTerraformApplyResponse {
  s.Body = v
  return s
}

func (s *ExecuteTerraformApplyResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

