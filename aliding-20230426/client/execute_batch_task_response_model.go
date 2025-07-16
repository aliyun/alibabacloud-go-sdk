// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteBatchTaskResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteBatchTaskResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteBatchTaskResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteBatchTaskResponseBody) *ExecuteBatchTaskResponse
  GetBody() *ExecuteBatchTaskResponseBody 
}

type ExecuteBatchTaskResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteBatchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteBatchTaskResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteBatchTaskResponse) GoString() string {
  return s.String()
}

func (s *ExecuteBatchTaskResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteBatchTaskResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteBatchTaskResponse) GetBody() *ExecuteBatchTaskResponseBody  {
  return s.Body
}

func (s *ExecuteBatchTaskResponse) SetHeaders(v map[string]*string) *ExecuteBatchTaskResponse {
  s.Headers = v
  return s
}

func (s *ExecuteBatchTaskResponse) SetStatusCode(v int32) *ExecuteBatchTaskResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteBatchTaskResponse) SetBody(v *ExecuteBatchTaskResponseBody) *ExecuteBatchTaskResponse {
  s.Body = v
  return s
}

func (s *ExecuteBatchTaskResponse) Validate() error {
  return dara.Validate(s)
}

