// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListActionPlanActivitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActionPlanActivities(v []*ListActionPlanActivitiesResponseBodyActionPlanActivities) *ListActionPlanActivitiesResponseBody
	GetActionPlanActivities() []*ListActionPlanActivitiesResponseBodyActionPlanActivities
	SetMaxResults(v int32) *ListActionPlanActivitiesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListActionPlanActivitiesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListActionPlanActivitiesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListActionPlanActivitiesResponseBody
	GetTotalCount() *int32
}

type ListActionPlanActivitiesResponseBody struct {
	ActionPlanActivities []*ListActionPlanActivitiesResponseBodyActionPlanActivities `json:"ActionPlanActivities,omitempty" xml:"ActionPlanActivities,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1d2db86scXXXXXXXXXX
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 40
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListActionPlanActivitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListActionPlanActivitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListActionPlanActivitiesResponseBody) GetActionPlanActivities() []*ListActionPlanActivitiesResponseBodyActionPlanActivities {
	return s.ActionPlanActivities
}

func (s *ListActionPlanActivitiesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListActionPlanActivitiesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListActionPlanActivitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListActionPlanActivitiesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListActionPlanActivitiesResponseBody) SetActionPlanActivities(v []*ListActionPlanActivitiesResponseBodyActionPlanActivities) *ListActionPlanActivitiesResponseBody {
	s.ActionPlanActivities = v
	return s
}

func (s *ListActionPlanActivitiesResponseBody) SetMaxResults(v int32) *ListActionPlanActivitiesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListActionPlanActivitiesResponseBody) SetNextToken(v string) *ListActionPlanActivitiesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListActionPlanActivitiesResponseBody) SetRequestId(v string) *ListActionPlanActivitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListActionPlanActivitiesResponseBody) SetTotalCount(v int32) *ListActionPlanActivitiesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListActionPlanActivitiesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListActionPlanActivitiesResponseBodyActionPlanActivities struct {
	// example:
	//
	// et-4119e3f60eb34fc4
	ActionPlanActivityId *string `json:"ActionPlanActivityId,omitempty" xml:"ActionPlanActivityId,omitempty"`
	// example:
	//
	// 100
	CreatedCapacity *float32 `json:"CreatedCapacity,omitempty" xml:"CreatedCapacity,omitempty"`
	// example:
	//
	// 0
	DestroyCapacity *float32 `json:"DestroyCapacity,omitempty" xml:"DestroyCapacity,omitempty"`
	// example:
	//
	// 2025-08-10 18:28:05
	EndTime *string                                                         `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Jobs    []*ListActionPlanActivitiesResponseBodyActionPlanActivitiesJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// example:
	//
	// 2025-08-10 18:28:05
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// InProcess
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListActionPlanActivitiesResponseBodyActionPlanActivities) String() string {
	return dara.Prettify(s)
}

func (s ListActionPlanActivitiesResponseBodyActionPlanActivities) GoString() string {
	return s.String()
}

func (s *ListActionPlanActivitiesResponseBodyActionPlanActivities) GetActionPlanActivityId() *string {
	return s.ActionPlanActivityId
}

func (s *ListActionPlanActivitiesResponseBodyActionPlanActivities) GetCreatedCapacity() *float32 {
	return s.CreatedCapacity
}

func (s *ListActionPlanActivitiesResponseBodyActionPlanActivities) GetDestroyCapacity() *float32 {
	return s.DestroyCapacity
}

func (s *ListActionPlanActivitiesResponseBodyActionPlanActivities) GetEndTime() *string {
	return s.EndTime
}

func (s *ListActionPlanActivitiesResponseBodyActionPlanActivities) GetJobs() []*ListActionPlanActivitiesResponseBodyActionPlanActivitiesJobs {
	return s.Jobs
}

func (s *ListActionPlanActivitiesResponseBodyActionPlanActivities) GetStartTime() *string {
	return s.StartTime
}

func (s *ListActionPlanActivitiesResponseBodyActionPlanActivities) GetStatus() *string {
	return s.Status
}

func (s *ListActionPlanActivitiesResponseBodyActionPlanActivities) SetActionPlanActivityId(v string) *ListActionPlanActivitiesResponseBodyActionPlanActivities {
	s.ActionPlanActivityId = &v
	return s
}

func (s *ListActionPlanActivitiesResponseBodyActionPlanActivities) SetCreatedCapacity(v float32) *ListActionPlanActivitiesResponseBodyActionPlanActivities {
	s.CreatedCapacity = &v
	return s
}

func (s *ListActionPlanActivitiesResponseBodyActionPlanActivities) SetDestroyCapacity(v float32) *ListActionPlanActivitiesResponseBodyActionPlanActivities {
	s.DestroyCapacity = &v
	return s
}

func (s *ListActionPlanActivitiesResponseBodyActionPlanActivities) SetEndTime(v string) *ListActionPlanActivitiesResponseBodyActionPlanActivities {
	s.EndTime = &v
	return s
}

func (s *ListActionPlanActivitiesResponseBodyActionPlanActivities) SetJobs(v []*ListActionPlanActivitiesResponseBodyActionPlanActivitiesJobs) *ListActionPlanActivitiesResponseBodyActionPlanActivities {
	s.Jobs = v
	return s
}

func (s *ListActionPlanActivitiesResponseBodyActionPlanActivities) SetStartTime(v string) *ListActionPlanActivitiesResponseBodyActionPlanActivities {
	s.StartTime = &v
	return s
}

func (s *ListActionPlanActivitiesResponseBodyActionPlanActivities) SetStatus(v string) *ListActionPlanActivitiesResponseBodyActionPlanActivities {
	s.Status = &v
	return s
}

func (s *ListActionPlanActivitiesResponseBodyActionPlanActivities) Validate() error {
	return dara.Validate(s)
}

type ListActionPlanActivitiesResponseBodyActionPlanActivitiesJobs struct {
	// example:
	//
	// job-hz12dqq8y3ormo1hz49h
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// Create
	JobOperationType *string `json:"JobOperationType,omitempty" xml:"JobOperationType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListActionPlanActivitiesResponseBodyActionPlanActivitiesJobs) String() string {
	return dara.Prettify(s)
}

func (s ListActionPlanActivitiesResponseBodyActionPlanActivitiesJobs) GoString() string {
	return s.String()
}

func (s *ListActionPlanActivitiesResponseBodyActionPlanActivitiesJobs) GetJobId() *string {
	return s.JobId
}

func (s *ListActionPlanActivitiesResponseBodyActionPlanActivitiesJobs) GetJobOperationType() *string {
	return s.JobOperationType
}

func (s *ListActionPlanActivitiesResponseBodyActionPlanActivitiesJobs) GetRegionId() *string {
	return s.RegionId
}

func (s *ListActionPlanActivitiesResponseBodyActionPlanActivitiesJobs) SetJobId(v string) *ListActionPlanActivitiesResponseBodyActionPlanActivitiesJobs {
	s.JobId = &v
	return s
}

func (s *ListActionPlanActivitiesResponseBodyActionPlanActivitiesJobs) SetJobOperationType(v string) *ListActionPlanActivitiesResponseBodyActionPlanActivitiesJobs {
	s.JobOperationType = &v
	return s
}

func (s *ListActionPlanActivitiesResponseBodyActionPlanActivitiesJobs) SetRegionId(v string) *ListActionPlanActivitiesResponseBodyActionPlanActivitiesJobs {
	s.RegionId = &v
	return s
}

func (s *ListActionPlanActivitiesResponseBodyActionPlanActivitiesJobs) Validate() error {
	return dara.Validate(s)
}
