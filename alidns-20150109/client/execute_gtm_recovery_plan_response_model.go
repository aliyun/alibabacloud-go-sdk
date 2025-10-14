// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteGtmRecoveryPlanResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteGtmRecoveryPlanResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteGtmRecoveryPlanResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteGtmRecoveryPlanResponseBody) *ExecuteGtmRecoveryPlanResponse
  GetBody() *ExecuteGtmRecoveryPlanResponseBody 
}

type ExecuteGtmRecoveryPlanResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteGtmRecoveryPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteGtmRecoveryPlanResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteGtmRecoveryPlanResponse) GoString() string {
  return s.String()
}

func (s *ExecuteGtmRecoveryPlanResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteGtmRecoveryPlanResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteGtmRecoveryPlanResponse) GetBody() *ExecuteGtmRecoveryPlanResponseBody  {
  return s.Body
}

func (s *ExecuteGtmRecoveryPlanResponse) SetHeaders(v map[string]*string) *ExecuteGtmRecoveryPlanResponse {
  s.Headers = v
  return s
}

func (s *ExecuteGtmRecoveryPlanResponse) SetStatusCode(v int32) *ExecuteGtmRecoveryPlanResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteGtmRecoveryPlanResponse) SetBody(v *ExecuteGtmRecoveryPlanResponseBody) *ExecuteGtmRecoveryPlanResponse {
  s.Body = v
  return s
}

func (s *ExecuteGtmRecoveryPlanResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

