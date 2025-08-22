// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTerraformDestroyResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteTerraformDestroyResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteTerraformDestroyResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteTerraformDestroyResponseBody) *ExecuteTerraformDestroyResponse
  GetBody() *ExecuteTerraformDestroyResponseBody 
}

type ExecuteTerraformDestroyResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteTerraformDestroyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteTerraformDestroyResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTerraformDestroyResponse) GoString() string {
  return s.String()
}

func (s *ExecuteTerraformDestroyResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteTerraformDestroyResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteTerraformDestroyResponse) GetBody() *ExecuteTerraformDestroyResponseBody  {
  return s.Body
}

func (s *ExecuteTerraformDestroyResponse) SetHeaders(v map[string]*string) *ExecuteTerraformDestroyResponse {
  s.Headers = v
  return s
}

func (s *ExecuteTerraformDestroyResponse) SetStatusCode(v int32) *ExecuteTerraformDestroyResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteTerraformDestroyResponse) SetBody(v *ExecuteTerraformDestroyResponseBody) *ExecuteTerraformDestroyResponse {
  s.Body = v
  return s
}

func (s *ExecuteTerraformDestroyResponse) Validate() error {
  return dara.Validate(s)
}

