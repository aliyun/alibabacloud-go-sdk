// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteSmartHomeSceneResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteSmartHomeSceneResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteSmartHomeSceneResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteSmartHomeSceneResponseBody) *ExecuteSmartHomeSceneResponse
  GetBody() *ExecuteSmartHomeSceneResponseBody 
}

type ExecuteSmartHomeSceneResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteSmartHomeSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteSmartHomeSceneResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSmartHomeSceneResponse) GoString() string {
  return s.String()
}

func (s *ExecuteSmartHomeSceneResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteSmartHomeSceneResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteSmartHomeSceneResponse) GetBody() *ExecuteSmartHomeSceneResponseBody  {
  return s.Body
}

func (s *ExecuteSmartHomeSceneResponse) SetHeaders(v map[string]*string) *ExecuteSmartHomeSceneResponse {
  s.Headers = v
  return s
}

func (s *ExecuteSmartHomeSceneResponse) SetStatusCode(v int32) *ExecuteSmartHomeSceneResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteSmartHomeSceneResponse) SetBody(v *ExecuteSmartHomeSceneResponseBody) *ExecuteSmartHomeSceneResponse {
  s.Body = v
  return s
}

func (s *ExecuteSmartHomeSceneResponse) Validate() error {
  return dara.Validate(s)
}

