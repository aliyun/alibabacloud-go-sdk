// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAutoDisposeRecordsRequest interface {
  dara.Model
  String() string
  GoString() string
  SetLang(v string) *ExecuteAutoDisposeRecordsRequest
  GetLang() *string 
  SetSelectedEntityList(v []*ExecuteAutoDisposeRecordsRequestSelectedEntityList) *ExecuteAutoDisposeRecordsRequest
  GetSelectedEntityList() []*ExecuteAutoDisposeRecordsRequestSelectedEntityList 
  SetUnSelectedEntityList(v []*ExecuteAutoDisposeRecordsRequestUnSelectedEntityList) *ExecuteAutoDisposeRecordsRequest
  GetUnSelectedEntityList() []*ExecuteAutoDisposeRecordsRequestUnSelectedEntityList 
}

type ExecuteAutoDisposeRecordsRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // zh
  Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
  SelectedEntityList []*ExecuteAutoDisposeRecordsRequestSelectedEntityList `json:"SelectedEntityList,omitempty" xml:"SelectedEntityList,omitempty" type:"Repeated"`
  UnSelectedEntityList []*ExecuteAutoDisposeRecordsRequestUnSelectedEntityList `json:"UnSelectedEntityList,omitempty" xml:"UnSelectedEntityList,omitempty" type:"Repeated"`
}

func (s ExecuteAutoDisposeRecordsRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAutoDisposeRecordsRequest) GoString() string {
  return s.String()
}

func (s *ExecuteAutoDisposeRecordsRequest) GetLang() *string  {
  return s.Lang
}

func (s *ExecuteAutoDisposeRecordsRequest) GetSelectedEntityList() []*ExecuteAutoDisposeRecordsRequestSelectedEntityList  {
  return s.SelectedEntityList
}

func (s *ExecuteAutoDisposeRecordsRequest) GetUnSelectedEntityList() []*ExecuteAutoDisposeRecordsRequestUnSelectedEntityList  {
  return s.UnSelectedEntityList
}

func (s *ExecuteAutoDisposeRecordsRequest) SetLang(v string) *ExecuteAutoDisposeRecordsRequest {
  s.Lang = &v
  return s
}

func (s *ExecuteAutoDisposeRecordsRequest) SetSelectedEntityList(v []*ExecuteAutoDisposeRecordsRequestSelectedEntityList) *ExecuteAutoDisposeRecordsRequest {
  s.SelectedEntityList = v
  return s
}

func (s *ExecuteAutoDisposeRecordsRequest) SetUnSelectedEntityList(v []*ExecuteAutoDisposeRecordsRequestUnSelectedEntityList) *ExecuteAutoDisposeRecordsRequest {
  s.UnSelectedEntityList = v
  return s
}

func (s *ExecuteAutoDisposeRecordsRequest) Validate() error {
  if s.SelectedEntityList != nil {
    for _, item := range s.SelectedEntityList {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  if s.UnSelectedEntityList != nil {
    for _, item := range s.UnSelectedEntityList {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExecuteAutoDisposeRecordsRequestSelectedEntityList struct {
  // example:
  // 
  // 0000089b040b8935fed2e24ca2ec8335
  AutoDisposeRecordId *string `json:"AutoDisposeRecordId,omitempty" xml:"AutoDisposeRecordId,omitempty"`
  // example:
  // 
  // 9938fc2708ddc7b7651f3a19e4f09962
  EntityUuid *string `json:"EntityUuid,omitempty" xml:"EntityUuid,omitempty"`
}

func (s ExecuteAutoDisposeRecordsRequestSelectedEntityList) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAutoDisposeRecordsRequestSelectedEntityList) GoString() string {
  return s.String()
}

func (s *ExecuteAutoDisposeRecordsRequestSelectedEntityList) GetAutoDisposeRecordId() *string  {
  return s.AutoDisposeRecordId
}

func (s *ExecuteAutoDisposeRecordsRequestSelectedEntityList) GetEntityUuid() *string  {
  return s.EntityUuid
}

func (s *ExecuteAutoDisposeRecordsRequestSelectedEntityList) SetAutoDisposeRecordId(v string) *ExecuteAutoDisposeRecordsRequestSelectedEntityList {
  s.AutoDisposeRecordId = &v
  return s
}

func (s *ExecuteAutoDisposeRecordsRequestSelectedEntityList) SetEntityUuid(v string) *ExecuteAutoDisposeRecordsRequestSelectedEntityList {
  s.EntityUuid = &v
  return s
}

func (s *ExecuteAutoDisposeRecordsRequestSelectedEntityList) Validate() error {
  return dara.Validate(s)
}

type ExecuteAutoDisposeRecordsRequestUnSelectedEntityList struct {
  // example:
  // 
  // 0000089b040b8935fed2e24ca2ec8335
  AutoDisposeRecordId *string `json:"AutoDisposeRecordId,omitempty" xml:"AutoDisposeRecordId,omitempty"`
  // example:
  // 
  // ae6ac3e1c9ada174eb8dadd029a2e9d1
  EntityUuid *string `json:"EntityUuid,omitempty" xml:"EntityUuid,omitempty"`
}

func (s ExecuteAutoDisposeRecordsRequestUnSelectedEntityList) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAutoDisposeRecordsRequestUnSelectedEntityList) GoString() string {
  return s.String()
}

func (s *ExecuteAutoDisposeRecordsRequestUnSelectedEntityList) GetAutoDisposeRecordId() *string  {
  return s.AutoDisposeRecordId
}

func (s *ExecuteAutoDisposeRecordsRequestUnSelectedEntityList) GetEntityUuid() *string  {
  return s.EntityUuid
}

func (s *ExecuteAutoDisposeRecordsRequestUnSelectedEntityList) SetAutoDisposeRecordId(v string) *ExecuteAutoDisposeRecordsRequestUnSelectedEntityList {
  s.AutoDisposeRecordId = &v
  return s
}

func (s *ExecuteAutoDisposeRecordsRequestUnSelectedEntityList) SetEntityUuid(v string) *ExecuteAutoDisposeRecordsRequestUnSelectedEntityList {
  s.EntityUuid = &v
  return s
}

func (s *ExecuteAutoDisposeRecordsRequestUnSelectedEntityList) Validate() error {
  return dara.Validate(s)
}

