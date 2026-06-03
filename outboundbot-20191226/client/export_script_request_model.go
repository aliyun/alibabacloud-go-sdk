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
  // This parameter is required.
  // 
  // example:
  // 
  // 12f3dd08-0c55-44ce-9b64-e69d35ed3a76
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // d004cfd2-6a81-491c-83c6-cbe186620c95
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

