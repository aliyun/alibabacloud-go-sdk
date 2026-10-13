// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecMetaProfilingSkillJobDisableResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecMetaProfilingSkillJobDisableResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecMetaProfilingSkillJobDisableResponse
  GetStatusCode() *int32 
  SetBody(v *ExecMetaProfilingSkillJobDisableResponseBody) *ExecMetaProfilingSkillJobDisableResponse
  GetBody() *ExecMetaProfilingSkillJobDisableResponseBody 
}

type ExecMetaProfilingSkillJobDisableResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecMetaProfilingSkillJobDisableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecMetaProfilingSkillJobDisableResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecMetaProfilingSkillJobDisableResponse) GoString() string {
  return s.String()
}

func (s *ExecMetaProfilingSkillJobDisableResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecMetaProfilingSkillJobDisableResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecMetaProfilingSkillJobDisableResponse) GetBody() *ExecMetaProfilingSkillJobDisableResponseBody  {
  return s.Body
}

func (s *ExecMetaProfilingSkillJobDisableResponse) SetHeaders(v map[string]*string) *ExecMetaProfilingSkillJobDisableResponse {
  s.Headers = v
  return s
}

func (s *ExecMetaProfilingSkillJobDisableResponse) SetStatusCode(v int32) *ExecMetaProfilingSkillJobDisableResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecMetaProfilingSkillJobDisableResponse) SetBody(v *ExecMetaProfilingSkillJobDisableResponseBody) *ExecMetaProfilingSkillJobDisableResponse {
  s.Body = v
  return s
}

func (s *ExecMetaProfilingSkillJobDisableResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

