// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAutoDisposeRecordsResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteAutoDisposeRecordsResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteAutoDisposeRecordsResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteAutoDisposeRecordsResponseBody) *ExecuteAutoDisposeRecordsResponse
  GetBody() *ExecuteAutoDisposeRecordsResponseBody 
}

type ExecuteAutoDisposeRecordsResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteAutoDisposeRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAutoDisposeRecordsResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAutoDisposeRecordsResponse) GoString() string {
  return s.String()
}

func (s *ExecuteAutoDisposeRecordsResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteAutoDisposeRecordsResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteAutoDisposeRecordsResponse) GetBody() *ExecuteAutoDisposeRecordsResponseBody  {
  return s.Body
}

func (s *ExecuteAutoDisposeRecordsResponse) SetHeaders(v map[string]*string) *ExecuteAutoDisposeRecordsResponse {
  s.Headers = v
  return s
}

func (s *ExecuteAutoDisposeRecordsResponse) SetStatusCode(v int32) *ExecuteAutoDisposeRecordsResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteAutoDisposeRecordsResponse) SetBody(v *ExecuteAutoDisposeRecordsResponseBody) *ExecuteAutoDisposeRecordsResponse {
  s.Body = v
  return s
}

func (s *ExecuteAutoDisposeRecordsResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

