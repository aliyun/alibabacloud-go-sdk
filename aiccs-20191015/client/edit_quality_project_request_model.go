// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditQualityProjectRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAnalysisIds(v []*int64) *EditQualityProjectRequest
  GetAnalysisIds() []*int64 
  SetChannelTouchType(v []*int32) *EditQualityProjectRequest
  GetChannelTouchType() []*int32 
  SetCheckFreqType(v int32) *EditQualityProjectRequest
  GetCheckFreqType() *int32 
  SetDepList(v []*int64) *EditQualityProjectRequest
  GetDepList() []*int64 
  SetGroupList(v []*int64) *EditQualityProjectRequest
  GetGroupList() []*int64 
  SetInstanceId(v string) *EditQualityProjectRequest
  GetInstanceId() *string 
  SetProjectId(v int64) *EditQualityProjectRequest
  GetProjectId() *int64 
  SetProjectName(v string) *EditQualityProjectRequest
  GetProjectName() *string 
  SetProjectVersion(v int32) *EditQualityProjectRequest
  GetProjectVersion() *int32 
  SetScopeType(v int32) *EditQualityProjectRequest
  GetScopeType() *int32 
  SetServicerList(v []*string) *EditQualityProjectRequest
  GetServicerList() []*string 
  SetTimeRangeEnd(v string) *EditQualityProjectRequest
  GetTimeRangeEnd() *string 
  SetTimeRangeStart(v string) *EditQualityProjectRequest
  GetTimeRangeStart() *string 
}

type EditQualityProjectRequest struct {
  // This parameter is required.
  AnalysisIds []*int64 `json:"AnalysisIds,omitempty" xml:"AnalysisIds,omitempty" type:"Repeated"`
  // This parameter is required.
  ChannelTouchType []*int32 `json:"ChannelTouchType,omitempty" xml:"ChannelTouchType,omitempty" type:"Repeated"`
  // This parameter is required.
  CheckFreqType *int32 `json:"CheckFreqType,omitempty" xml:"CheckFreqType,omitempty"`
  DepList []*int64 `json:"DepList,omitempty" xml:"DepList,omitempty" type:"Repeated"`
  GroupList []*int64 `json:"GroupList,omitempty" xml:"GroupList,omitempty" type:"Repeated"`
  // This parameter is required.
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // This parameter is required.
  ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
  // This parameter is required.
  ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
  // This parameter is required.
  ProjectVersion *int32 `json:"ProjectVersion,omitempty" xml:"ProjectVersion,omitempty"`
  // This parameter is required.
  ScopeType *int32 `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
  ServicerList []*string `json:"ServicerList,omitempty" xml:"ServicerList,omitempty" type:"Repeated"`
  TimeRangeEnd *string `json:"TimeRangeEnd,omitempty" xml:"TimeRangeEnd,omitempty"`
  TimeRangeStart *string `json:"TimeRangeStart,omitempty" xml:"TimeRangeStart,omitempty"`
}

func (s EditQualityProjectRequest) String() string {
  return dara.Prettify(s)
}

func (s EditQualityProjectRequest) GoString() string {
  return s.String()
}

func (s *EditQualityProjectRequest) GetAnalysisIds() []*int64  {
  return s.AnalysisIds
}

func (s *EditQualityProjectRequest) GetChannelTouchType() []*int32  {
  return s.ChannelTouchType
}

func (s *EditQualityProjectRequest) GetCheckFreqType() *int32  {
  return s.CheckFreqType
}

func (s *EditQualityProjectRequest) GetDepList() []*int64  {
  return s.DepList
}

func (s *EditQualityProjectRequest) GetGroupList() []*int64  {
  return s.GroupList
}

func (s *EditQualityProjectRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EditQualityProjectRequest) GetProjectId() *int64  {
  return s.ProjectId
}

func (s *EditQualityProjectRequest) GetProjectName() *string  {
  return s.ProjectName
}

func (s *EditQualityProjectRequest) GetProjectVersion() *int32  {
  return s.ProjectVersion
}

func (s *EditQualityProjectRequest) GetScopeType() *int32  {
  return s.ScopeType
}

func (s *EditQualityProjectRequest) GetServicerList() []*string  {
  return s.ServicerList
}

func (s *EditQualityProjectRequest) GetTimeRangeEnd() *string  {
  return s.TimeRangeEnd
}

func (s *EditQualityProjectRequest) GetTimeRangeStart() *string  {
  return s.TimeRangeStart
}

func (s *EditQualityProjectRequest) SetAnalysisIds(v []*int64) *EditQualityProjectRequest {
  s.AnalysisIds = v
  return s
}

func (s *EditQualityProjectRequest) SetChannelTouchType(v []*int32) *EditQualityProjectRequest {
  s.ChannelTouchType = v
  return s
}

func (s *EditQualityProjectRequest) SetCheckFreqType(v int32) *EditQualityProjectRequest {
  s.CheckFreqType = &v
  return s
}

func (s *EditQualityProjectRequest) SetDepList(v []*int64) *EditQualityProjectRequest {
  s.DepList = v
  return s
}

func (s *EditQualityProjectRequest) SetGroupList(v []*int64) *EditQualityProjectRequest {
  s.GroupList = v
  return s
}

func (s *EditQualityProjectRequest) SetInstanceId(v string) *EditQualityProjectRequest {
  s.InstanceId = &v
  return s
}

func (s *EditQualityProjectRequest) SetProjectId(v int64) *EditQualityProjectRequest {
  s.ProjectId = &v
  return s
}

func (s *EditQualityProjectRequest) SetProjectName(v string) *EditQualityProjectRequest {
  s.ProjectName = &v
  return s
}

func (s *EditQualityProjectRequest) SetProjectVersion(v int32) *EditQualityProjectRequest {
  s.ProjectVersion = &v
  return s
}

func (s *EditQualityProjectRequest) SetScopeType(v int32) *EditQualityProjectRequest {
  s.ScopeType = &v
  return s
}

func (s *EditQualityProjectRequest) SetServicerList(v []*string) *EditQualityProjectRequest {
  s.ServicerList = v
  return s
}

func (s *EditQualityProjectRequest) SetTimeRangeEnd(v string) *EditQualityProjectRequest {
  s.TimeRangeEnd = &v
  return s
}

func (s *EditQualityProjectRequest) SetTimeRangeStart(v string) *EditQualityProjectRequest {
  s.TimeRangeStart = &v
  return s
}

func (s *EditQualityProjectRequest) Validate() error {
  return dara.Validate(s)
}

