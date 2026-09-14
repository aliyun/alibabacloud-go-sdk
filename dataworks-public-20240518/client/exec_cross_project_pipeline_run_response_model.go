// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecCrossProjectPipelineRunResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecCrossProjectPipelineRunResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecCrossProjectPipelineRunResponse
  GetStatusCode() *int32 
  SetBody(v *ExecCrossProjectPipelineRunResponseBody) *ExecCrossProjectPipelineRunResponse
  GetBody() *ExecCrossProjectPipelineRunResponseBody 
}

type ExecCrossProjectPipelineRunResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecCrossProjectPipelineRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecCrossProjectPipelineRunResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecCrossProjectPipelineRunResponse) GoString() string {
  return s.String()
}

func (s *ExecCrossProjectPipelineRunResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecCrossProjectPipelineRunResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecCrossProjectPipelineRunResponse) GetBody() *ExecCrossProjectPipelineRunResponseBody  {
  return s.Body
}

func (s *ExecCrossProjectPipelineRunResponse) SetHeaders(v map[string]*string) *ExecCrossProjectPipelineRunResponse {
  s.Headers = v
  return s
}

func (s *ExecCrossProjectPipelineRunResponse) SetStatusCode(v int32) *ExecCrossProjectPipelineRunResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecCrossProjectPipelineRunResponse) SetBody(v *ExecCrossProjectPipelineRunResponseBody) *ExecCrossProjectPipelineRunResponse {
  s.Body = v
  return s
}

func (s *ExecCrossProjectPipelineRunResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

