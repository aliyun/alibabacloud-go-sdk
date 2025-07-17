// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAdhocWorkflowInstanceResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteAdhocWorkflowInstanceResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteAdhocWorkflowInstanceResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteAdhocWorkflowInstanceResponseBody) *ExecuteAdhocWorkflowInstanceResponse
  GetBody() *ExecuteAdhocWorkflowInstanceResponseBody 
}

type ExecuteAdhocWorkflowInstanceResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteAdhocWorkflowInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAdhocWorkflowInstanceResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAdhocWorkflowInstanceResponse) GoString() string {
  return s.String()
}

func (s *ExecuteAdhocWorkflowInstanceResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteAdhocWorkflowInstanceResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteAdhocWorkflowInstanceResponse) GetBody() *ExecuteAdhocWorkflowInstanceResponseBody  {
  return s.Body
}

func (s *ExecuteAdhocWorkflowInstanceResponse) SetHeaders(v map[string]*string) *ExecuteAdhocWorkflowInstanceResponse {
  s.Headers = v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceResponse) SetStatusCode(v int32) *ExecuteAdhocWorkflowInstanceResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceResponse) SetBody(v *ExecuteAdhocWorkflowInstanceResponseBody) *ExecuteAdhocWorkflowInstanceResponse {
  s.Body = v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceResponse) Validate() error {
  return dara.Validate(s)
}

