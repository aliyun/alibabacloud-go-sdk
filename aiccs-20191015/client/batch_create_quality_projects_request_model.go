// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateQualityProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnalysisIds(v []*int64) *BatchCreateQualityProjectsRequest
	GetAnalysisIds() []*int64
	SetChannelTouchType(v []*int32) *BatchCreateQualityProjectsRequest
	GetChannelTouchType() []*int32
	SetCheckFreqType(v int32) *BatchCreateQualityProjectsRequest
	GetCheckFreqType() *int32
	SetInstanceList(v []*string) *BatchCreateQualityProjectsRequest
	GetInstanceList() []*string
	SetProjectName(v string) *BatchCreateQualityProjectsRequest
	GetProjectName() *string
	SetTimeRangeEnd(v string) *BatchCreateQualityProjectsRequest
	GetTimeRangeEnd() *string
	SetTimeRangeStart(v string) *BatchCreateQualityProjectsRequest
	GetTimeRangeStart() *string
}

type BatchCreateQualityProjectsRequest struct {
	// This parameter is required.
	AnalysisIds      []*int64 `json:"AnalysisIds,omitempty" xml:"AnalysisIds,omitempty" type:"Repeated"`
	ChannelTouchType []*int32 `json:"ChannelTouchType,omitempty" xml:"ChannelTouchType,omitempty" type:"Repeated"`
	// This parameter is required.
	CheckFreqType *int32 `json:"CheckFreqType,omitempty" xml:"CheckFreqType,omitempty"`
	// This parameter is required.
	InstanceList []*string `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// This parameter is required.
	ProjectName    *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	TimeRangeEnd   *string `json:"TimeRangeEnd,omitempty" xml:"TimeRangeEnd,omitempty"`
	TimeRangeStart *string `json:"TimeRangeStart,omitempty" xml:"TimeRangeStart,omitempty"`
}

func (s BatchCreateQualityProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateQualityProjectsRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateQualityProjectsRequest) GetAnalysisIds() []*int64 {
	return s.AnalysisIds
}

func (s *BatchCreateQualityProjectsRequest) GetChannelTouchType() []*int32 {
	return s.ChannelTouchType
}

func (s *BatchCreateQualityProjectsRequest) GetCheckFreqType() *int32 {
	return s.CheckFreqType
}

func (s *BatchCreateQualityProjectsRequest) GetInstanceList() []*string {
	return s.InstanceList
}

func (s *BatchCreateQualityProjectsRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *BatchCreateQualityProjectsRequest) GetTimeRangeEnd() *string {
	return s.TimeRangeEnd
}

func (s *BatchCreateQualityProjectsRequest) GetTimeRangeStart() *string {
	return s.TimeRangeStart
}

func (s *BatchCreateQualityProjectsRequest) SetAnalysisIds(v []*int64) *BatchCreateQualityProjectsRequest {
	s.AnalysisIds = v
	return s
}

func (s *BatchCreateQualityProjectsRequest) SetChannelTouchType(v []*int32) *BatchCreateQualityProjectsRequest {
	s.ChannelTouchType = v
	return s
}

func (s *BatchCreateQualityProjectsRequest) SetCheckFreqType(v int32) *BatchCreateQualityProjectsRequest {
	s.CheckFreqType = &v
	return s
}

func (s *BatchCreateQualityProjectsRequest) SetInstanceList(v []*string) *BatchCreateQualityProjectsRequest {
	s.InstanceList = v
	return s
}

func (s *BatchCreateQualityProjectsRequest) SetProjectName(v string) *BatchCreateQualityProjectsRequest {
	s.ProjectName = &v
	return s
}

func (s *BatchCreateQualityProjectsRequest) SetTimeRangeEnd(v string) *BatchCreateQualityProjectsRequest {
	s.TimeRangeEnd = &v
	return s
}

func (s *BatchCreateQualityProjectsRequest) SetTimeRangeStart(v string) *BatchCreateQualityProjectsRequest {
	s.TimeRangeStart = &v
	return s
}

func (s *BatchCreateQualityProjectsRequest) Validate() error {
	return dara.Validate(s)
}
