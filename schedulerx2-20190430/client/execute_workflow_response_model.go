// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteWorkflowResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteWorkflowResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteWorkflowResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteWorkflowResponseBody) *ExecuteWorkflowResponse
  GetBody() *ExecuteWorkflowResponseBody 
}

type ExecuteWorkflowResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteWorkflowResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteWorkflowResponse) GoString() string {
  return s.String()
}

func (s *ExecuteWorkflowResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteWorkflowResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteWorkflowResponse) GetBody() *ExecuteWorkflowResponseBody  {
  return s.Body
}

func (s *ExecuteWorkflowResponse) SetHeaders(v map[string]*string) *ExecuteWorkflowResponse {
  s.Headers = v
  return s
}

func (s *ExecuteWorkflowResponse) SetStatusCode(v int32) *ExecuteWorkflowResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteWorkflowResponse) SetBody(v *ExecuteWorkflowResponseBody) *ExecuteWorkflowResponse {
  s.Body = v
  return s
}

func (s *ExecuteWorkflowResponse) Validate() error {
  return dara.Validate(s)
}

