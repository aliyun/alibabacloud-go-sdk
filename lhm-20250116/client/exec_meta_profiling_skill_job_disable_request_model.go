// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecMetaProfilingSkillJobDisableRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDsId(v int64) *ExecMetaProfilingSkillJobDisableRequest
  GetDsId() *int64 
}

type ExecMetaProfilingSkillJobDisableRequest struct {
  // example:
  // 
  // 290
  DsId *int64 `json:"dsId,omitempty" xml:"dsId,omitempty"`
}

func (s ExecMetaProfilingSkillJobDisableRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecMetaProfilingSkillJobDisableRequest) GoString() string {
  return s.String()
}

func (s *ExecMetaProfilingSkillJobDisableRequest) GetDsId() *int64  {
  return s.DsId
}

func (s *ExecMetaProfilingSkillJobDisableRequest) SetDsId(v int64) *ExecMetaProfilingSkillJobDisableRequest {
  s.DsId = &v
  return s
}

func (s *ExecMetaProfilingSkillJobDisableRequest) Validate() error {
  return dara.Validate(s)
}

