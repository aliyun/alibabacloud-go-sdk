// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecMetaProfilingTaskRunResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v bool) *ExecMetaProfilingTaskRunResponseBody
  GetData() *bool 
  SetErrCode(v string) *ExecMetaProfilingTaskRunResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *ExecMetaProfilingTaskRunResponseBody
  GetErrMessage() *string 
  SetSuccess(v bool) *ExecMetaProfilingTaskRunResponseBody
  GetSuccess() *bool 
}

type ExecMetaProfilingTaskRunResponseBody struct {
  // example:
  // 
  // 100
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

func (s ExecMetaProfilingTaskRunResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecMetaProfilingTaskRunResponseBody) GoString() string {
  return s.String()
}

func (s *ExecMetaProfilingTaskRunResponseBody) GetData() *bool  {
  return s.Data
}

func (s *ExecMetaProfilingTaskRunResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *ExecMetaProfilingTaskRunResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *ExecMetaProfilingTaskRunResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecMetaProfilingTaskRunResponseBody) SetData(v bool) *ExecMetaProfilingTaskRunResponseBody {
  s.Data = &v
  return s
}

func (s *ExecMetaProfilingTaskRunResponseBody) SetErrCode(v string) *ExecMetaProfilingTaskRunResponseBody {
  s.ErrCode = &v
  return s
}

func (s *ExecMetaProfilingTaskRunResponseBody) SetErrMessage(v string) *ExecMetaProfilingTaskRunResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *ExecMetaProfilingTaskRunResponseBody) SetSuccess(v bool) *ExecMetaProfilingTaskRunResponseBody {
  s.Success = &v
  return s
}

func (s *ExecMetaProfilingTaskRunResponseBody) Validate() error {
  return dara.Validate(s)
}

