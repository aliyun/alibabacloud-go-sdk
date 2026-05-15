// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQualityProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnalysisIds(v []*int64) *CreateQualityProjectRequest
	GetAnalysisIds() []*int64
	SetChannelTouchType(v []*int32) *CreateQualityProjectRequest
	GetChannelTouchType() []*int32
	SetCheckFreqType(v int32) *CreateQualityProjectRequest
	GetCheckFreqType() *int32
	SetDepList(v []*int64) *CreateQualityProjectRequest
	GetDepList() []*int64
	SetGroupList(v []*int64) *CreateQualityProjectRequest
	GetGroupList() []*int64
	SetInstanceId(v string) *CreateQualityProjectRequest
	GetInstanceId() *string
	SetProjectName(v string) *CreateQualityProjectRequest
	GetProjectName() *string
	SetScopeType(v int32) *CreateQualityProjectRequest
	GetScopeType() *int32
	SetServicerList(v []*string) *CreateQualityProjectRequest
	GetServicerList() []*string
	SetTimeRangeEnd(v string) *CreateQualityProjectRequest
	GetTimeRangeEnd() *string
	SetTimeRangeStart(v string) *CreateQualityProjectRequest
	GetTimeRangeStart() *string
}

type CreateQualityProjectRequest struct {
	// This parameter is required.
	AnalysisIds      []*int64 `json:"AnalysisIds,omitempty" xml:"AnalysisIds,omitempty" type:"Repeated"`
	ChannelTouchType []*int32 `json:"ChannelTouchType,omitempty" xml:"ChannelTouchType,omitempty" type:"Repeated"`
	// This parameter is required.
	CheckFreqType *int32   `json:"CheckFreqType,omitempty" xml:"CheckFreqType,omitempty"`
	DepList       []*int64 `json:"DepList,omitempty" xml:"DepList,omitempty" type:"Repeated"`
	GroupList     []*int64 `json:"GroupList,omitempty" xml:"GroupList,omitempty" type:"Repeated"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// This parameter is required.
	ScopeType      *int32    `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
	ServicerList   []*string `json:"ServicerList,omitempty" xml:"ServicerList,omitempty" type:"Repeated"`
	TimeRangeEnd   *string   `json:"TimeRangeEnd,omitempty" xml:"TimeRangeEnd,omitempty"`
	TimeRangeStart *string   `json:"TimeRangeStart,omitempty" xml:"TimeRangeStart,omitempty"`
}

func (s CreateQualityProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateQualityProjectRequest) GetAnalysisIds() []*int64 {
	return s.AnalysisIds
}

func (s *CreateQualityProjectRequest) GetChannelTouchType() []*int32 {
	return s.ChannelTouchType
}

func (s *CreateQualityProjectRequest) GetCheckFreqType() *int32 {
	return s.CheckFreqType
}

func (s *CreateQualityProjectRequest) GetDepList() []*int64 {
	return s.DepList
}

func (s *CreateQualityProjectRequest) GetGroupList() []*int64 {
	return s.GroupList
}

func (s *CreateQualityProjectRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateQualityProjectRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateQualityProjectRequest) GetScopeType() *int32 {
	return s.ScopeType
}

func (s *CreateQualityProjectRequest) GetServicerList() []*string {
	return s.ServicerList
}

func (s *CreateQualityProjectRequest) GetTimeRangeEnd() *string {
	return s.TimeRangeEnd
}

func (s *CreateQualityProjectRequest) GetTimeRangeStart() *string {
	return s.TimeRangeStart
}

func (s *CreateQualityProjectRequest) SetAnalysisIds(v []*int64) *CreateQualityProjectRequest {
	s.AnalysisIds = v
	return s
}

func (s *CreateQualityProjectRequest) SetChannelTouchType(v []*int32) *CreateQualityProjectRequest {
	s.ChannelTouchType = v
	return s
}

func (s *CreateQualityProjectRequest) SetCheckFreqType(v int32) *CreateQualityProjectRequest {
	s.CheckFreqType = &v
	return s
}

func (s *CreateQualityProjectRequest) SetDepList(v []*int64) *CreateQualityProjectRequest {
	s.DepList = v
	return s
}

func (s *CreateQualityProjectRequest) SetGroupList(v []*int64) *CreateQualityProjectRequest {
	s.GroupList = v
	return s
}

func (s *CreateQualityProjectRequest) SetInstanceId(v string) *CreateQualityProjectRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateQualityProjectRequest) SetProjectName(v string) *CreateQualityProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateQualityProjectRequest) SetScopeType(v int32) *CreateQualityProjectRequest {
	s.ScopeType = &v
	return s
}

func (s *CreateQualityProjectRequest) SetServicerList(v []*string) *CreateQualityProjectRequest {
	s.ServicerList = v
	return s
}

func (s *CreateQualityProjectRequest) SetTimeRangeEnd(v string) *CreateQualityProjectRequest {
	s.TimeRangeEnd = &v
	return s
}

func (s *CreateQualityProjectRequest) SetTimeRangeStart(v string) *CreateQualityProjectRequest {
	s.TimeRangeStart = &v
	return s
}

func (s *CreateQualityProjectRequest) Validate() error {
	return dara.Validate(s)
}
