// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteComponentRequest interface {
  dara.Model
  String() string
  GoString() string
  SetComponentActionName(v string) *ExecuteComponentRequest
  GetComponentActionName() *string 
  SetComponentAssetUuid(v string) *ExecuteComponentRequest
  GetComponentAssetUuid() *string 
  SetComponentInput(v string) *ExecuteComponentRequest
  GetComponentInput() *string 
  SetComponentName(v string) *ExecuteComponentRequest
  GetComponentName() *string 
  SetLang(v string) *ExecuteComponentRequest
  GetLang() *string 
  SetPlayBookNodeName(v string) *ExecuteComponentRequest
  GetPlayBookNodeName() *string 
  SetPlaybookUuid(v string) *ExecuteComponentRequest
  GetPlaybookUuid() *string 
}

type ExecuteComponentRequest struct {
  // example:
  // 
  // doRequest
  ComponentActionName *string `json:"ComponentActionName,omitempty" xml:"ComponentActionName,omitempty"`
  // example:
  // 
  // 1C5F11E9-****-51F0-****-43BB312A0557
  ComponentAssetUuid *string `json:"ComponentAssetUuid,omitempty" xml:"ComponentAssetUuid,omitempty"`
  // example:
  // 
  // {}
  ComponentInput *string `json:"ComponentInput,omitempty" xml:"ComponentInput,omitempty"`
  // example:
  // 
  // SLS
  ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
  // example:
  // 
  // zh
  Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
  // example:
  // 
  // node1
  PlayBookNodeName *string `json:"PlayBookNodeName,omitempty" xml:"PlayBookNodeName,omitempty"`
  // example:
  // 
  // ac343acc-1a61-4084-9a1cxxxxx
  PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s ExecuteComponentRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteComponentRequest) GoString() string {
  return s.String()
}

func (s *ExecuteComponentRequest) GetComponentActionName() *string  {
  return s.ComponentActionName
}

func (s *ExecuteComponentRequest) GetComponentAssetUuid() *string  {
  return s.ComponentAssetUuid
}

func (s *ExecuteComponentRequest) GetComponentInput() *string  {
  return s.ComponentInput
}

func (s *ExecuteComponentRequest) GetComponentName() *string  {
  return s.ComponentName
}

func (s *ExecuteComponentRequest) GetLang() *string  {
  return s.Lang
}

func (s *ExecuteComponentRequest) GetPlayBookNodeName() *string  {
  return s.PlayBookNodeName
}

func (s *ExecuteComponentRequest) GetPlaybookUuid() *string  {
  return s.PlaybookUuid
}

func (s *ExecuteComponentRequest) SetComponentActionName(v string) *ExecuteComponentRequest {
  s.ComponentActionName = &v
  return s
}

func (s *ExecuteComponentRequest) SetComponentAssetUuid(v string) *ExecuteComponentRequest {
  s.ComponentAssetUuid = &v
  return s
}

func (s *ExecuteComponentRequest) SetComponentInput(v string) *ExecuteComponentRequest {
  s.ComponentInput = &v
  return s
}

func (s *ExecuteComponentRequest) SetComponentName(v string) *ExecuteComponentRequest {
  s.ComponentName = &v
  return s
}

func (s *ExecuteComponentRequest) SetLang(v string) *ExecuteComponentRequest {
  s.Lang = &v
  return s
}

func (s *ExecuteComponentRequest) SetPlayBookNodeName(v string) *ExecuteComponentRequest {
  s.PlayBookNodeName = &v
  return s
}

func (s *ExecuteComponentRequest) SetPlaybookUuid(v string) *ExecuteComponentRequest {
  s.PlaybookUuid = &v
  return s
}

func (s *ExecuteComponentRequest) Validate() error {
  return dara.Validate(s)
}

