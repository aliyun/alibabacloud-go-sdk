// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportScriptRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceId(v string) *ExportScriptRequest
  GetInstanceId() *string 
  SetScriptId(v string) *ExportScriptRequest
  GetScriptId() *string 
}

type ExportScriptRequest struct {
  // example:
  // 
  // c28fc549-d88f-4f6e-85ad-a0806e5e39c0
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // example:
  // 
  // 64241e64-190c-45d1-af66-06f51c07b090
  ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s ExportScriptRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportScriptRequest) GoString() string {
  return s.String()
}

func (s *ExportScriptRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *ExportScriptRequest) GetScriptId() *string  {
  return s.ScriptId
}

func (s *ExportScriptRequest) SetInstanceId(v string) *ExportScriptRequest {
  s.InstanceId = &v
  return s
}

func (s *ExportScriptRequest) SetScriptId(v string) *ExportScriptRequest {
  s.ScriptId = &v
  return s
}

func (s *ExportScriptRequest) Validate() error {
  return dara.Validate(s)
}

