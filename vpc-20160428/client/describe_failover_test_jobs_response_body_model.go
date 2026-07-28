// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFailoverTestJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeFailoverTestJobsResponseBody
	GetCount() *int32
	SetFailoverTestJobList(v []*DescribeFailoverTestJobsResponseBodyFailoverTestJobList) *DescribeFailoverTestJobsResponseBody
	GetFailoverTestJobList() []*DescribeFailoverTestJobsResponseBodyFailoverTestJobList
	SetMaxResults(v int32) *DescribeFailoverTestJobsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeFailoverTestJobsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeFailoverTestJobsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeFailoverTestJobsResponseBody
	GetTotalCount() *int32
}

type DescribeFailoverTestJobsResponseBody struct {
	// The number of entries on the current page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The list of failover test jobs.
	FailoverTestJobList []*DescribeFailoverTestJobsResponseBodyFailoverTestJobList `json:"FailoverTestJobList,omitempty" xml:"FailoverTestJobList,omitempty" type:"Repeated"`
	// The number of entries per page for paginated queries. Valid values: **1 to 100**. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. Valid values:
	//
	// - Leave this parameter empty for the first query or if no next query exists.
	//
	// - If a next query exists, set this parameter to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries in the list.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFailoverTestJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFailoverTestJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFailoverTestJobsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeFailoverTestJobsResponseBody) GetFailoverTestJobList() []*DescribeFailoverTestJobsResponseBodyFailoverTestJobList {
	return s.FailoverTestJobList
}

func (s *DescribeFailoverTestJobsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeFailoverTestJobsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeFailoverTestJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFailoverTestJobsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeFailoverTestJobsResponseBody) SetCount(v int32) *DescribeFailoverTestJobsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeFailoverTestJobsResponseBody) SetFailoverTestJobList(v []*DescribeFailoverTestJobsResponseBodyFailoverTestJobList) *DescribeFailoverTestJobsResponseBody {
	s.FailoverTestJobList = v
	return s
}

func (s *DescribeFailoverTestJobsResponseBody) SetMaxResults(v int32) *DescribeFailoverTestJobsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeFailoverTestJobsResponseBody) SetNextToken(v string) *DescribeFailoverTestJobsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeFailoverTestJobsResponseBody) SetRequestId(v string) *DescribeFailoverTestJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFailoverTestJobsResponseBody) SetTotalCount(v int32) *DescribeFailoverTestJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeFailoverTestJobsResponseBody) Validate() error {
	if s.FailoverTestJobList != nil {
		for _, item := range s.FailoverTestJobList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFailoverTestJobsResponseBodyFailoverTestJobList struct {
	// The description of the failover test job.
	//
	// The description is 0 to 256 characters in length and cannot start with **http://*	- or **https://**.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The test duration. Unit: minutes. Valid values: **1 to 4320**.
	//
	// example:
	//
	// 60
	JobDuration *string `json:"JobDuration,omitempty" xml:"JobDuration,omitempty"`
	// The failover test job ID.
	//
	// example:
	//
	// ftj-bp1yh6mvi13aq3g8w****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The failover test type. Valid values:
	//
	// - **StartNow**: The test starts immediately after the failover test job is created.
	//
	// - **StartLater**: Only the test job is created. The test is not started.
	//
	// example:
	//
	// StartNow
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The name of the failover test job.
	//
	// The name is 0 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of failover test resource IDs.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The failover test resource type. Valid values: **PHYSICALCONNECTION**: Express Connect circuit.
	//
	// example:
	//
	// PHYSICALCONNECTION
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The start time of the failover test job. The time is displayed in UTC in the YYYY-MM-DDThh:mm:ssZ format based on the ISO 8601 standard.
	//
	// example:
	//
	// 2023-11-21T14:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the failover test job. Valid values:
	//
	// - **Init**: Pending.
	//
	// - **Starting**: Starting.
	//
	// - **Testing**: In progress.
	//
	// - **Stopping**: Stopping.
	//
	// - **Stopped**: Completed.
	//
	// example:
	//
	// Init
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The end time of the failover test job. The time is displayed in UTC in the YYYY-MM-DDThh:mm:ssZ format based on the ISO 8601 standard.
	//
	// example:
	//
	// 2023-11-21T15:00:00Z
	StopTime *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
}

func (s DescribeFailoverTestJobsResponseBodyFailoverTestJobList) String() string {
	return dara.Prettify(s)
}

func (s DescribeFailoverTestJobsResponseBodyFailoverTestJobList) GoString() string {
	return s.String()
}

func (s *DescribeFailoverTestJobsResponseBodyFailoverTestJobList) GetDescription() *string {
	return s.Description
}

func (s *DescribeFailoverTestJobsResponseBodyFailoverTestJobList) GetJobDuration() *string {
	return s.JobDuration
}

func (s *DescribeFailoverTestJobsResponseBodyFailoverTestJobList) GetJobId() *string {
	return s.JobId
}

func (s *DescribeFailoverTestJobsResponseBodyFailoverTestJobList) GetJobType() *string {
	return s.JobType
}

func (s *DescribeFailoverTestJobsResponseBodyFailoverTestJobList) GetName() *string {
	return s.Name
}

func (s *DescribeFailoverTestJobsResponseBodyFailoverTestJobList) GetResourceId() []*string {
	return s.ResourceId
}

func (s *DescribeFailoverTestJobsResponseBodyFailoverTestJobList) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeFailoverTestJobsResponseBodyFailoverTestJobList) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeFailoverTestJobsResponseBodyFailoverTestJobList) GetStatus() *string {
	return s.Status
}

func (s *DescribeFailoverTestJobsResponseBodyFailoverTestJobList) GetStopTime() *string {
	return s.StopTime
}

func (s *DescribeFailoverTestJobsResponseBodyFailoverTestJobList) SetDescription(v string) *DescribeFailoverTestJobsResponseBodyFailoverTestJobList {
	s.Description = &v
	return s
}

func (s *DescribeFailoverTestJobsResponseBodyFailoverTestJobList) SetJobDuration(v string) *DescribeFailoverTestJobsResponseBodyFailoverTestJobList {
	s.JobDuration = &v
	return s
}

func (s *DescribeFailoverTestJobsResponseBodyFailoverTestJobList) SetJobId(v string) *DescribeFailoverTestJobsResponseBodyFailoverTestJobList {
	s.JobId = &v
	return s
}

func (s *DescribeFailoverTestJobsResponseBodyFailoverTestJobList) SetJobType(v string) *DescribeFailoverTestJobsResponseBodyFailoverTestJobList {
	s.JobType = &v
	return s
}

func (s *DescribeFailoverTestJobsResponseBodyFailoverTestJobList) SetName(v string) *DescribeFailoverTestJobsResponseBodyFailoverTestJobList {
	s.Name = &v
	return s
}

func (s *DescribeFailoverTestJobsResponseBodyFailoverTestJobList) SetResourceId(v []*string) *DescribeFailoverTestJobsResponseBodyFailoverTestJobList {
	s.ResourceId = v
	return s
}

func (s *DescribeFailoverTestJobsResponseBodyFailoverTestJobList) SetResourceType(v string) *DescribeFailoverTestJobsResponseBodyFailoverTestJobList {
	s.ResourceType = &v
	return s
}

func (s *DescribeFailoverTestJobsResponseBodyFailoverTestJobList) SetStartTime(v string) *DescribeFailoverTestJobsResponseBodyFailoverTestJobList {
	s.StartTime = &v
	return s
}

func (s *DescribeFailoverTestJobsResponseBodyFailoverTestJobList) SetStatus(v string) *DescribeFailoverTestJobsResponseBodyFailoverTestJobList {
	s.Status = &v
	return s
}

func (s *DescribeFailoverTestJobsResponseBodyFailoverTestJobList) SetStopTime(v string) *DescribeFailoverTestJobsResponseBodyFailoverTestJobList {
	s.StopTime = &v
	return s
}

func (s *DescribeFailoverTestJobsResponseBodyFailoverTestJobList) Validate() error {
	return dara.Validate(s)
}
