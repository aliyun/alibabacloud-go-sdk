// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeliveryHistoryJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryHistoryJobs(v []*ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) *ListDeliveryHistoryJobsResponseBody
	GetDeliveryHistoryJobs() []*ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs
	SetPageNumber(v int32) *ListDeliveryHistoryJobsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDeliveryHistoryJobsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDeliveryHistoryJobsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDeliveryHistoryJobsResponseBody
	GetTotalCount() *int32
}

type ListDeliveryHistoryJobsResponseBody struct {
	// The list of historical event delivery tasks.
	DeliveryHistoryJobs []*ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs `json:"DeliveryHistoryJobs,omitempty" xml:"DeliveryHistoryJobs,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B190816C-6DCA-4DC5-9B8E-EE0367B57CFF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of historical event delivery tasks returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDeliveryHistoryJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeliveryHistoryJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeliveryHistoryJobsResponseBody) GetDeliveryHistoryJobs() []*ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs {
	return s.DeliveryHistoryJobs
}

func (s *ListDeliveryHistoryJobsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDeliveryHistoryJobsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDeliveryHistoryJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeliveryHistoryJobsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDeliveryHistoryJobsResponseBody) SetDeliveryHistoryJobs(v []*ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) *ListDeliveryHistoryJobsResponseBody {
	s.DeliveryHistoryJobs = v
	return s
}

func (s *ListDeliveryHistoryJobsResponseBody) SetPageNumber(v int32) *ListDeliveryHistoryJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponseBody) SetPageSize(v int32) *ListDeliveryHistoryJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponseBody) SetRequestId(v string) *ListDeliveryHistoryJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponseBody) SetTotalCount(v int32) *ListDeliveryHistoryJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs struct {
	// The time when the task was created.
	//
	// example:
	//
	// 2021-04-26T03:17:04Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The time when the task ended.
	//
	// example:
	//
	// 2021-04-26T03:22:04Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The home region of the trail.
	//
	// example:
	//
	// cn-hangzhou
	HomeRegion *string `json:"HomeRegion,omitempty" xml:"HomeRegion,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 16602
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The task status. Valid values:
	//
	// 	- 0: The task is initializing.
	//
	// 	- 1: The task is delivering historical events.
	//
	// 	- 2: The task is complete.
	//
	// 	- 3: The task fails.
	//
	// example:
	//
	// 2
	JobStatus *int32 `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// The time when the task started.
	//
	// example:
	//
	// 2021-01-26T03:17:04Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the trail.
	//
	// example:
	//
	// trail-name
	TrailName *string `json:"TrailName,omitempty" xml:"TrailName,omitempty"`
	// The time when the task was updated.
	//
	// example:
	//
	// 2021-04-26T03:20:08Z
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) String() string {
	return dara.Prettify(s)
}

func (s ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) GoString() string {
	return s.String()
}

func (s *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) GetEndTime() *string {
	return s.EndTime
}

func (s *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) GetHomeRegion() *string {
	return s.HomeRegion
}

func (s *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) GetJobId() *int64 {
	return s.JobId
}

func (s *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) GetJobStatus() *int32 {
	return s.JobStatus
}

func (s *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) GetStartTime() *string {
	return s.StartTime
}

func (s *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) GetTrailName() *string {
	return s.TrailName
}

func (s *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) GetUpdatedTime() *string {
	return s.UpdatedTime
}

func (s *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) SetCreatedTime(v string) *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs {
	s.CreatedTime = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) SetEndTime(v string) *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs {
	s.EndTime = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) SetHomeRegion(v string) *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs {
	s.HomeRegion = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) SetJobId(v int64) *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs {
	s.JobId = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) SetJobStatus(v int32) *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs {
	s.JobStatus = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) SetStartTime(v string) *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs {
	s.StartTime = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) SetTrailName(v string) *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs {
	s.TrailName = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) SetUpdatedTime(v string) *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs {
	s.UpdatedTime = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) Validate() error {
	return dara.Validate(s)
}
