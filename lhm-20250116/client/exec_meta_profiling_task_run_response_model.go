// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecMetaProfilingTaskRunResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecMetaProfilingTaskRunResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecMetaProfilingTaskRunResponse
  GetStatusCode() *int32 
  SetBody(v *ExecMetaProfilingTaskRunResponseBody) *ExecMetaProfilingTaskRunResponse
  GetBody() *ExecMetaProfilingTaskRunResponseBody 
}

type ExecMetaProfilingTaskRunResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecMetaProfilingTaskRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecMetaProfilingTaskRunResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecMetaProfilingTaskRunResponse) GoString() string {
  return s.String()
}

func (s *ExecMetaProfilingTaskRunResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecMetaProfilingTaskRunResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecMetaProfilingTaskRunResponse) GetBody() *ExecMetaProfilingTaskRunResponseBody  {
  return s.Body
}

func (s *ExecMetaProfilingTaskRunResponse) SetHeaders(v map[string]*string) *ExecMetaProfilingTaskRunResponse {
  s.Headers = v
  return s
}

func (s *ExecMetaProfilingTaskRunResponse) SetStatusCode(v int32) *ExecMetaProfilingTaskRunResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecMetaProfilingTaskRunResponse) SetBody(v *ExecMetaProfilingTaskRunResponseBody) *ExecMetaProfilingTaskRunResponse {
  s.Body = v
  return s
}

func (s *ExecMetaProfilingTaskRunResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

