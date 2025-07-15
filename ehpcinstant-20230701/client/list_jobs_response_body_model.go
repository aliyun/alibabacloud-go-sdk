// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobList(v []*ListJobsResponseBodyJobList) *ListJobsResponseBody
	GetJobList() []*ListJobsResponseBodyJobList
	SetPageNumber(v int32) *ListJobsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListJobsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListJobsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListJobsResponseBody
	GetTotalCount() *int32
}

type ListJobsResponseBody struct {
	JobList []*ListJobsResponseBodyJobList `json:"JobList,omitempty" xml:"JobList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) GetJobList() []*ListJobsResponseBodyJobList {
	return s.JobList
}

func (s *ListJobsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJobsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListJobsResponseBody) SetJobList(v []*ListJobsResponseBodyJobList) *ListJobsResponseBody {
	s.JobList = v
	return s
}

func (s *ListJobsResponseBody) SetPageNumber(v int32) *ListJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListJobsResponseBody) SetPageSize(v int32) *ListJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListJobsResponseBody) SetRequestId(v string) *ListJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobsResponseBody) SetTotalCount(v int32) *ListJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListJobsResponseBodyJobList struct {
	AppExtraInfo *string `json:"AppExtraInfo,omitempty" xml:"AppExtraInfo,omitempty"`
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 2024-01-25 12:29:21
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2024-01-25 12:35:23
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	ExecutorCount *int32 `json:"ExecutorCount,omitempty" xml:"ExecutorCount,omitempty"`
	// example:
	//
	// Demo
	JobDescription *string `json:"JobDescription,omitempty" xml:"JobDescription,omitempty"`
	// example:
	//
	// job-xxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// testJob
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// 129**********
	OwnerUid *string `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// example:
	//
	// 2024-01-25 12:29:23
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Running
	Status *string                            `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*ListJobsResponseBodyJobListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TaskCount *int32 `json:"TaskCount,omitempty" xml:"TaskCount,omitempty"`
	// example:
	//
	// true
	TaskSustainable *bool `json:"TaskSustainable,omitempty" xml:"TaskSustainable,omitempty"`
}

func (s ListJobsResponseBodyJobList) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyJobList) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobList) GetAppExtraInfo() *string {
	return s.AppExtraInfo
}

func (s *ListJobsResponseBodyJobList) GetAppName() *string {
	return s.AppName
}

func (s *ListJobsResponseBodyJobList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListJobsResponseBodyJobList) GetEndTime() *string {
	return s.EndTime
}

func (s *ListJobsResponseBodyJobList) GetExecutorCount() *int32 {
	return s.ExecutorCount
}

func (s *ListJobsResponseBodyJobList) GetJobDescription() *string {
	return s.JobDescription
}

func (s *ListJobsResponseBodyJobList) GetJobId() *string {
	return s.JobId
}

func (s *ListJobsResponseBodyJobList) GetJobName() *string {
	return s.JobName
}

func (s *ListJobsResponseBodyJobList) GetOwnerUid() *string {
	return s.OwnerUid
}

func (s *ListJobsResponseBodyJobList) GetStartTime() *string {
	return s.StartTime
}

func (s *ListJobsResponseBodyJobList) GetStatus() *string {
	return s.Status
}

func (s *ListJobsResponseBodyJobList) GetTags() []*ListJobsResponseBodyJobListTags {
	return s.Tags
}

func (s *ListJobsResponseBodyJobList) GetTaskCount() *int32 {
	return s.TaskCount
}

func (s *ListJobsResponseBodyJobList) GetTaskSustainable() *bool {
	return s.TaskSustainable
}

func (s *ListJobsResponseBodyJobList) SetAppExtraInfo(v string) *ListJobsResponseBodyJobList {
	s.AppExtraInfo = &v
	return s
}

func (s *ListJobsResponseBodyJobList) SetAppName(v string) *ListJobsResponseBodyJobList {
	s.AppName = &v
	return s
}

func (s *ListJobsResponseBodyJobList) SetCreateTime(v string) *ListJobsResponseBodyJobList {
	s.CreateTime = &v
	return s
}

func (s *ListJobsResponseBodyJobList) SetEndTime(v string) *ListJobsResponseBodyJobList {
	s.EndTime = &v
	return s
}

func (s *ListJobsResponseBodyJobList) SetExecutorCount(v int32) *ListJobsResponseBodyJobList {
	s.ExecutorCount = &v
	return s
}

func (s *ListJobsResponseBodyJobList) SetJobDescription(v string) *ListJobsResponseBodyJobList {
	s.JobDescription = &v
	return s
}

func (s *ListJobsResponseBodyJobList) SetJobId(v string) *ListJobsResponseBodyJobList {
	s.JobId = &v
	return s
}

func (s *ListJobsResponseBodyJobList) SetJobName(v string) *ListJobsResponseBodyJobList {
	s.JobName = &v
	return s
}

func (s *ListJobsResponseBodyJobList) SetOwnerUid(v string) *ListJobsResponseBodyJobList {
	s.OwnerUid = &v
	return s
}

func (s *ListJobsResponseBodyJobList) SetStartTime(v string) *ListJobsResponseBodyJobList {
	s.StartTime = &v
	return s
}

func (s *ListJobsResponseBodyJobList) SetStatus(v string) *ListJobsResponseBodyJobList {
	s.Status = &v
	return s
}

func (s *ListJobsResponseBodyJobList) SetTags(v []*ListJobsResponseBodyJobListTags) *ListJobsResponseBodyJobList {
	s.Tags = v
	return s
}

func (s *ListJobsResponseBodyJobList) SetTaskCount(v int32) *ListJobsResponseBodyJobList {
	s.TaskCount = &v
	return s
}

func (s *ListJobsResponseBodyJobList) SetTaskSustainable(v bool) *ListJobsResponseBodyJobList {
	s.TaskSustainable = &v
	return s
}

func (s *ListJobsResponseBodyJobList) Validate() error {
	return dara.Validate(s)
}

type ListJobsResponseBodyJobListTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListJobsResponseBodyJobListTags) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyJobListTags) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobListTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListJobsResponseBodyJobListTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListJobsResponseBodyJobListTags) SetTagKey(v string) *ListJobsResponseBodyJobListTags {
	s.TagKey = &v
	return s
}

func (s *ListJobsResponseBodyJobListTags) SetTagValue(v string) *ListJobsResponseBodyJobListTags {
	s.TagValue = &v
	return s
}

func (s *ListJobsResponseBodyJobListTags) Validate() error {
	return dara.Validate(s)
}
