// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecMetaProfilingSkillJobDisableResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v bool) *ExecMetaProfilingSkillJobDisableResponseBody
  GetData() *bool 
  SetErrCode(v string) *ExecMetaProfilingSkillJobDisableResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecMetaProfilingSkillJobDisableResponseBody
  GetErrMessage() *string 
  SetSuccess(v bool) *ExecMetaProfilingSkillJobDisableResponseBody
  GetSuccess() *bool 
}

type ExecMetaProfilingSkillJobDisableResponseBody struct {
  // example:
  // 
  // demo
  Data *bool `json:"data,omitempty" xml:"data,omitempty"`
  // example:
  // 
  // None
  ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
  // example:
  // 
  // success
  ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
  // example:
  // 
  // True
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecMetaProfilingSkillJobDisableResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecMetaProfilingSkillJobDisableResponseBody) GoString() string {
  return s.String()
}

func (s *ExecMetaProfilingSkillJobDisableResponseBody) GetData() *bool  {
  return s.Data
}

func (s *ExecMetaProfilingSkillJobDisableResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecMetaProfilingSkillJobDisableResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecMetaProfilingSkillJobDisableResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecMetaProfilingSkillJobDisableResponseBody) SetData(v bool) *ExecMetaProfilingSkillJobDisableResponseBody {
  s.Data = &v
  return s
}

func (s *ExecMetaProfilingSkillJobDisableResponseBody) SetErrCode(v string) *ExecMetaProfilingSkillJobDisableResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecMetaProfilingSkillJobDisableResponseBody) SetErrMessage(v string) *ExecMetaProfilingSkillJobDisableResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecMetaProfilingSkillJobDisableResponseBody) SetSuccess(v bool) *ExecMetaProfilingSkillJobDisableResponseBody {
  s.Success = &v
  return s
}

func (s *ExecMetaProfilingSkillJobDisableResponseBody) Validate() error {
  return dara.Validate(s)
}

