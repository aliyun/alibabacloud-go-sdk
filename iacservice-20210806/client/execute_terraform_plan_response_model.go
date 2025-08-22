// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTerraformPlanResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteTerraformPlanResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteTerraformPlanResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteTerraformPlanResponseBody) *ExecuteTerraformPlanResponse
  GetBody() *ExecuteTerraformPlanResponseBody 
}

type ExecuteTerraformPlanResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteTerraformPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteTerraformPlanResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTerraformPlanResponse) GoString() string {
  return s.String()
}

func (s *ExecuteTerraformPlanResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteTerraformPlanResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteTerraformPlanResponse) GetBody() *ExecuteTerraformPlanResponseBody  {
  return s.Body
}

func (s *ExecuteTerraformPlanResponse) SetHeaders(v map[string]*string) *ExecuteTerraformPlanResponse {
  s.Headers = v
  return s
}

func (s *ExecuteTerraformPlanResponse) SetStatusCode(v int32) *ExecuteTerraformPlanResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteTerraformPlanResponse) SetBody(v *ExecuteTerraformPlanResponseBody) *ExecuteTerraformPlanResponse {
  s.Body = v
  return s
}

func (s *ExecuteTerraformPlanResponse) Validate() error {
  return dara.Validate(s)
}

