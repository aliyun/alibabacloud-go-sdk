// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecutePlatformTaskResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecutePlatformTaskResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecutePlatformTaskResponse
  GetStatusCode() *int32 
  SetBody(v *ExecutePlatformTaskResponseBody) *ExecutePlatformTaskResponse
  GetBody() *ExecutePlatformTaskResponseBody 
}

type ExecutePlatformTaskResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecutePlatformTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecutePlatformTaskResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecutePlatformTaskResponse) GoString() string {
  return s.String()
}

func (s *ExecutePlatformTaskResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecutePlatformTaskResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecutePlatformTaskResponse) GetBody() *ExecutePlatformTaskResponseBody  {
  return s.Body
}

func (s *ExecutePlatformTaskResponse) SetHeaders(v map[string]*string) *ExecutePlatformTaskResponse {
  s.Headers = v
  return s
}

func (s *ExecutePlatformTaskResponse) SetStatusCode(v int32) *ExecutePlatformTaskResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecutePlatformTaskResponse) SetBody(v *ExecutePlatformTaskResponseBody) *ExecutePlatformTaskResponse {
  s.Body = v
  return s
}

func (s *ExecutePlatformTaskResponse) Validate() error {
  return dara.Validate(s)
}

