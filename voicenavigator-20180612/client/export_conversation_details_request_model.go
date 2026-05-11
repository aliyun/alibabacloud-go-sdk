// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportConversationDetailsRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBeginTimeLeftRange(v int64) *ExportConversationDetailsRequest
  GetBeginTimeLeftRange() *int64 
  SetBeginTimeRightRange(v int64) *ExportConversationDetailsRequest
  GetBeginTimeRightRange() *int64 
  SetCallingNumber(v string) *ExportConversationDetailsRequest
  GetCallingNumber() *string 
  SetDebugConversation(v int32) *ExportConversationDetailsRequest
  GetDebugConversation() *int32 
  SetInstanceId(v string) *ExportConversationDetailsRequest
  GetInstanceId() *string 
  SetOptions(v []*string) *ExportConversationDetailsRequest
  GetOptions() []*string 
  SetResult(v int32) *ExportConversationDetailsRequest
  GetResult() *int32 
  SetRoundsLeftRange(v int32) *ExportConversationDetailsRequest
  GetRoundsLeftRange() *int32 
  SetRoundsRightRange(v int32) *ExportConversationDetailsRequest
  GetRoundsRightRange() *int32 
}

type ExportConversationDetailsRequest struct {
  // example:
  // 
  // 1582266750353
  BeginTimeLeftRange *int64 `json:"BeginTimeLeftRange,omitempty" xml:"BeginTimeLeftRange,omitempty"`
  // example:
  // 
  // 1640793599000
  BeginTimeRightRange *int64 `json:"BeginTimeRightRange,omitempty" xml:"BeginTimeRightRange,omitempty"`
  // example:
  // 
  // 13581588**
  CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
  DebugConversation *int32 `json:"DebugConversation,omitempty" xml:"DebugConversation,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 6c01a99f-1b72-4f75-a8bd-3875766bd19d
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  Options []*string `json:"Options,omitempty" xml:"Options,omitempty" type:"Repeated"`
  Result *int32 `json:"Result,omitempty" xml:"Result,omitempty"`
  RoundsLeftRange *int32 `json:"RoundsLeftRange,omitempty" xml:"RoundsLeftRange,omitempty"`
  RoundsRightRange *int32 `json:"RoundsRightRange,omitempty" xml:"RoundsRightRange,omitempty"`
}

func (s ExportConversationDetailsRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportConversationDetailsRequest) GoString() string {
  return s.String()
}

func (s *ExportConversationDetailsRequest) GetBeginTimeLeftRange() *int64  {
  return s.BeginTimeLeftRange
}

func (s *ExportConversationDetailsRequest) GetBeginTimeRightRange() *int64  {
  return s.BeginTimeRightRange
}

func (s *ExportConversationDetailsRequest) GetCallingNumber() *string  {
  return s.CallingNumber
}

func (s *ExportConversationDetailsRequest) GetDebugConversation() *int32  {
  return s.DebugConversation
}

func (s *ExportConversationDetailsRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *ExportConversationDetailsRequest) GetOptions() []*string  {
  return s.Options
}

func (s *ExportConversationDetailsRequest) GetResult() *int32  {
  return s.Result
}

func (s *ExportConversationDetailsRequest) GetRoundsLeftRange() *int32  {
  return s.RoundsLeftRange
}

func (s *ExportConversationDetailsRequest) GetRoundsRightRange() *int32  {
  return s.RoundsRightRange
}

func (s *ExportConversationDetailsRequest) SetBeginTimeLeftRange(v int64) *ExportConversationDetailsRequest {
  s.BeginTimeLeftRange = &v
  return s
}

func (s *ExportConversationDetailsRequest) SetBeginTimeRightRange(v int64) *ExportConversationDetailsRequest {
  s.BeginTimeRightRange = &v
  return s
}

func (s *ExportConversationDetailsRequest) SetCallingNumber(v string) *ExportConversationDetailsRequest {
  s.CallingNumber = &v
  return s
}

func (s *ExportConversationDetailsRequest) SetDebugConversation(v int32) *ExportConversationDetailsRequest {
  s.DebugConversation = &v
  return s
}

func (s *ExportConversationDetailsRequest) SetInstanceId(v string) *ExportConversationDetailsRequest {
  s.InstanceId = &v
  return s
}

func (s *ExportConversationDetailsRequest) SetOptions(v []*string) *ExportConversationDetailsRequest {
  s.Options = v
  return s
}

func (s *ExportConversationDetailsRequest) SetResult(v int32) *ExportConversationDetailsRequest {
  s.Result = &v
  return s
}

func (s *ExportConversationDetailsRequest) SetRoundsLeftRange(v int32) *ExportConversationDetailsRequest {
  s.RoundsLeftRange = &v
  return s
}

func (s *ExportConversationDetailsRequest) SetRoundsRightRange(v int32) *ExportConversationDetailsRequest {
  s.RoundsRightRange = &v
  return s
}

func (s *ExportConversationDetailsRequest) Validate() error {
  return dara.Validate(s)
}

