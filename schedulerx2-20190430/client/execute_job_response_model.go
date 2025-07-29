// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteJobResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteJobResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteJobResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteJobResponseBody) *ExecuteJobResponse
  GetBody() *ExecuteJobResponseBody 
}

type ExecuteJobResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteJobResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteJobResponse) GoString() string {
  return s.String()
}

func (s *ExecuteJobResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteJobResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteJobResponse) GetBody() *ExecuteJobResponseBody  {
  return s.Body
}

func (s *ExecuteJobResponse) SetHeaders(v map[string]*string) *ExecuteJobResponse {
  s.Headers = v
  return s
}

func (s *ExecuteJobResponse) SetStatusCode(v int32) *ExecuteJobResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteJobResponse) SetBody(v *ExecuteJobResponseBody) *ExecuteJobResponse {
  s.Body = v
  return s
}

func (s *ExecuteJobResponse) Validate() error {
  return dara.Validate(s)
}

