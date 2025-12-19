// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteOperationASyncResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteOperationASyncResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteOperationASyncResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteOperationASyncResponseBody) *ExecuteOperationASyncResponse
  GetBody() *ExecuteOperationASyncResponseBody 
}

type ExecuteOperationASyncResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteOperationASyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteOperationASyncResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteOperationASyncResponse) GoString() string {
  return s.String()
}

func (s *ExecuteOperationASyncResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteOperationASyncResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteOperationASyncResponse) GetBody() *ExecuteOperationASyncResponseBody  {
  return s.Body
}

func (s *ExecuteOperationASyncResponse) SetHeaders(v map[string]*string) *ExecuteOperationASyncResponse {
  s.Headers = v
  return s
}

func (s *ExecuteOperationASyncResponse) SetStatusCode(v int32) *ExecuteOperationASyncResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteOperationASyncResponse) SetBody(v *ExecuteOperationASyncResponseBody) *ExecuteOperationASyncResponse {
  s.Body = v
  return s
}

func (s *ExecuteOperationASyncResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

