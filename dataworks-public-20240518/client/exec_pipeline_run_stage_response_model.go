// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecPipelineRunStageResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecPipelineRunStageResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecPipelineRunStageResponse
  GetStatusCode() *int32 
  SetBody(v *ExecPipelineRunStageResponseBody) *ExecPipelineRunStageResponse
  GetBody() *ExecPipelineRunStageResponseBody 
}

type ExecPipelineRunStageResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecPipelineRunStageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecPipelineRunStageResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecPipelineRunStageResponse) GoString() string {
  return s.String()
}

func (s *ExecPipelineRunStageResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecPipelineRunStageResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecPipelineRunStageResponse) GetBody() *ExecPipelineRunStageResponseBody  {
  return s.Body
}

func (s *ExecPipelineRunStageResponse) SetHeaders(v map[string]*string) *ExecPipelineRunStageResponse {
  s.Headers = v
  return s
}

func (s *ExecPipelineRunStageResponse) SetStatusCode(v int32) *ExecPipelineRunStageResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecPipelineRunStageResponse) SetBody(v *ExecPipelineRunStageResponseBody) *ExecPipelineRunStageResponse {
  s.Body = v
  return s
}

func (s *ExecPipelineRunStageResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

