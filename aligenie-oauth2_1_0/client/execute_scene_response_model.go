// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteSceneResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteSceneResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteSceneResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteSceneResponseBody) *ExecuteSceneResponse
  GetBody() *ExecuteSceneResponseBody 
}

type ExecuteSceneResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteSceneResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSceneResponse) GoString() string {
  return s.String()
}

func (s *ExecuteSceneResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteSceneResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteSceneResponse) GetBody() *ExecuteSceneResponseBody  {
  return s.Body
}

func (s *ExecuteSceneResponse) SetHeaders(v map[string]*string) *ExecuteSceneResponse {
  s.Headers = v
  return s
}

func (s *ExecuteSceneResponse) SetStatusCode(v int32) *ExecuteSceneResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteSceneResponse) SetBody(v *ExecuteSceneResponseBody) *ExecuteSceneResponse {
  s.Body = v
  return s
}

func (s *ExecuteSceneResponse) Validate() error {
  return dara.Validate(s)
}

