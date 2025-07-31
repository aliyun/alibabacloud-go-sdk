// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOnlineTestResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAudioData(v *DescribeOnlineTestResultResponseBodyAudioData) *DescribeOnlineTestResultResponseBody
	GetAudioData() *DescribeOnlineTestResultResponseBodyAudioData
	SetFrameData(v *DescribeOnlineTestResultResponseBodyFrameData) *DescribeOnlineTestResultResponseBody
	GetFrameData() *DescribeOnlineTestResultResponseBodyFrameData
	SetModerationTime(v string) *DescribeOnlineTestResultResponseBody
	GetModerationTime() *string
	SetRequestId(v string) *DescribeOnlineTestResultResponseBody
	GetRequestId() *string
	SetRiskLevel(v string) *DescribeOnlineTestResultResponseBody
	GetRiskLevel() *string
	SetServiceCode(v string) *DescribeOnlineTestResultResponseBody
	GetServiceCode() *string
	SetSummaryList(v []*DescribeOnlineTestResultResponseBodySummaryList) *DescribeOnlineTestResultResponseBody
	GetSummaryList() []*DescribeOnlineTestResultResponseBodySummaryList
	SetTaskId(v string) *DescribeOnlineTestResultResponseBody
	GetTaskId() *string
	SetTaskStatus(v string) *DescribeOnlineTestResultResponseBody
	GetTaskStatus() *string
	SetUrl(v string) *DescribeOnlineTestResultResponseBody
	GetUrl() *string
}

type DescribeOnlineTestResultResponseBody struct {
	AudioData *DescribeOnlineTestResultResponseBodyAudioData `json:"AudioData,omitempty" xml:"AudioData,omitempty" type:"Struct"`
	FrameData *DescribeOnlineTestResultResponseBodyFrameData `json:"FrameData,omitempty" xml:"FrameData,omitempty" type:"Struct"`
	// example:
	//
	// 1725761005419
	ModerationTime *string `json:"ModerationTime,omitempty" xml:"ModerationTime,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// VideoModeration
	ServiceCode *string                                            `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	SummaryList []*DescribeOnlineTestResultResponseBodySummaryList `json:"SummaryList,omitempty" xml:"SummaryList,omitempty" type:"Repeated"`
	// example:
	//
	// xxxxx-xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// SUCCESS
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// https://xxxxxxxxx.com/data/data.png
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeOnlineTestResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOnlineTestResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOnlineTestResultResponseBody) GetAudioData() *DescribeOnlineTestResultResponseBodyAudioData {
	return s.AudioData
}

func (s *DescribeOnlineTestResultResponseBody) GetFrameData() *DescribeOnlineTestResultResponseBodyFrameData {
	return s.FrameData
}

func (s *DescribeOnlineTestResultResponseBody) GetModerationTime() *string {
	return s.ModerationTime
}

func (s *DescribeOnlineTestResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOnlineTestResultResponseBody) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeOnlineTestResultResponseBody) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *DescribeOnlineTestResultResponseBody) GetSummaryList() []*DescribeOnlineTestResultResponseBodySummaryList {
	return s.SummaryList
}

func (s *DescribeOnlineTestResultResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeOnlineTestResultResponseBody) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeOnlineTestResultResponseBody) GetUrl() *string {
	return s.Url
}

func (s *DescribeOnlineTestResultResponseBody) SetAudioData(v *DescribeOnlineTestResultResponseBodyAudioData) *DescribeOnlineTestResultResponseBody {
	s.AudioData = v
	return s
}

func (s *DescribeOnlineTestResultResponseBody) SetFrameData(v *DescribeOnlineTestResultResponseBodyFrameData) *DescribeOnlineTestResultResponseBody {
	s.FrameData = v
	return s
}

func (s *DescribeOnlineTestResultResponseBody) SetModerationTime(v string) *DescribeOnlineTestResultResponseBody {
	s.ModerationTime = &v
	return s
}

func (s *DescribeOnlineTestResultResponseBody) SetRequestId(v string) *DescribeOnlineTestResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOnlineTestResultResponseBody) SetRiskLevel(v string) *DescribeOnlineTestResultResponseBody {
	s.RiskLevel = &v
	return s
}

func (s *DescribeOnlineTestResultResponseBody) SetServiceCode(v string) *DescribeOnlineTestResultResponseBody {
	s.ServiceCode = &v
	return s
}

func (s *DescribeOnlineTestResultResponseBody) SetSummaryList(v []*DescribeOnlineTestResultResponseBodySummaryList) *DescribeOnlineTestResultResponseBody {
	s.SummaryList = v
	return s
}

func (s *DescribeOnlineTestResultResponseBody) SetTaskId(v string) *DescribeOnlineTestResultResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeOnlineTestResultResponseBody) SetTaskStatus(v string) *DescribeOnlineTestResultResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *DescribeOnlineTestResultResponseBody) SetUrl(v string) *DescribeOnlineTestResultResponseBody {
	s.Url = &v
	return s
}

func (s *DescribeOnlineTestResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeOnlineTestResultResponseBodyAudioData struct {
	// example:
	//
	// 1724378510396
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeOnlineTestResultResponseBodyAudioData) String() string {
	return dara.Prettify(s)
}

func (s DescribeOnlineTestResultResponseBodyAudioData) GoString() string {
	return s.String()
}

func (s *DescribeOnlineTestResultResponseBodyAudioData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeOnlineTestResultResponseBodyAudioData) SetTimeStamp(v string) *DescribeOnlineTestResultResponseBodyAudioData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeOnlineTestResultResponseBodyAudioData) Validate() error {
	return dara.Validate(s)
}

type DescribeOnlineTestResultResponseBodyFrameData struct {
	// example:
	//
	// 1725761005419
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// example:
	//
	// https://xxxxxxxxx.com/data/data.mp4
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeOnlineTestResultResponseBodyFrameData) String() string {
	return dara.Prettify(s)
}

func (s DescribeOnlineTestResultResponseBodyFrameData) GoString() string {
	return s.String()
}

func (s *DescribeOnlineTestResultResponseBodyFrameData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeOnlineTestResultResponseBodyFrameData) GetUrl() *string {
	return s.Url
}

func (s *DescribeOnlineTestResultResponseBodyFrameData) SetTimeStamp(v string) *DescribeOnlineTestResultResponseBodyFrameData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeOnlineTestResultResponseBodyFrameData) SetUrl(v string) *DescribeOnlineTestResultResponseBodyFrameData {
	s.Url = &v
	return s
}

func (s *DescribeOnlineTestResultResponseBodyFrameData) Validate() error {
	return dara.Validate(s)
}

type DescribeOnlineTestResultResponseBodySummaryList struct {
	// example:
	//
	// video
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// high
	RiskLevel        *string           `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	RiskLevelSummary map[string]*int64 `json:"RiskLevelSummary,omitempty" xml:"RiskLevelSummary,omitempty"`
	// example:
	//
	// 10
	SliceCount *int32 `json:"SliceCount,omitempty" xml:"SliceCount,omitempty"`
}

func (s DescribeOnlineTestResultResponseBodySummaryList) String() string {
	return dara.Prettify(s)
}

func (s DescribeOnlineTestResultResponseBodySummaryList) GoString() string {
	return s.String()
}

func (s *DescribeOnlineTestResultResponseBodySummaryList) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeOnlineTestResultResponseBodySummaryList) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeOnlineTestResultResponseBodySummaryList) GetRiskLevelSummary() map[string]*int64 {
	return s.RiskLevelSummary
}

func (s *DescribeOnlineTestResultResponseBodySummaryList) GetSliceCount() *int32 {
	return s.SliceCount
}

func (s *DescribeOnlineTestResultResponseBodySummaryList) SetResourceType(v string) *DescribeOnlineTestResultResponseBodySummaryList {
	s.ResourceType = &v
	return s
}

func (s *DescribeOnlineTestResultResponseBodySummaryList) SetRiskLevel(v string) *DescribeOnlineTestResultResponseBodySummaryList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeOnlineTestResultResponseBodySummaryList) SetRiskLevelSummary(v map[string]*int64) *DescribeOnlineTestResultResponseBodySummaryList {
	s.RiskLevelSummary = v
	return s
}

func (s *DescribeOnlineTestResultResponseBodySummaryList) SetSliceCount(v int32) *DescribeOnlineTestResultResponseBodySummaryList {
	s.SliceCount = &v
	return s
}

func (s *DescribeOnlineTestResultResponseBodySummaryList) Validate() error {
	return dara.Validate(s)
}
