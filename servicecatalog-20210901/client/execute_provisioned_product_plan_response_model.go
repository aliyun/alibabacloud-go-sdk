// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteProvisionedProductPlanResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteProvisionedProductPlanResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteProvisionedProductPlanResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteProvisionedProductPlanResponseBody) *ExecuteProvisionedProductPlanResponse
  GetBody() *ExecuteProvisionedProductPlanResponseBody 
}

type ExecuteProvisionedProductPlanResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteProvisionedProductPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteProvisionedProductPlanResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteProvisionedProductPlanResponse) GoString() string {
  return s.String()
}

func (s *ExecuteProvisionedProductPlanResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteProvisionedProductPlanResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteProvisionedProductPlanResponse) GetBody() *ExecuteProvisionedProductPlanResponseBody  {
  return s.Body
}

func (s *ExecuteProvisionedProductPlanResponse) SetHeaders(v map[string]*string) *ExecuteProvisionedProductPlanResponse {
  s.Headers = v
  return s
}

func (s *ExecuteProvisionedProductPlanResponse) SetStatusCode(v int32) *ExecuteProvisionedProductPlanResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteProvisionedProductPlanResponse) SetBody(v *ExecuteProvisionedProductPlanResponseBody) *ExecuteProvisionedProductPlanResponse {
  s.Body = v
  return s
}

func (s *ExecuteProvisionedProductPlanResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

