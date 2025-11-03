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
	// The list of failover tests.
	FailoverTestJobList []*DescribeFailoverTestJobsResponseBodyFailoverTestJobList `json:"FailoverTestJobList,omitempty" xml:"FailoverTestJobList,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: **1 to 100**. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If no value is returned for **NextToken**, no next queries are sent.
	//
	// 	- If a value is returned for **NextToken**, the value is used to retrieve a new page of results.
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
	// The number of entries returned.
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
	// The description of the failover test.
	//
	// The description must be 0 to 256 characters in length and cannot start with \\*\\*http:// **or*	- https://\\*\\*.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The duration of the failover test. Unit: minutes. Valid values: **1 to 4320**.
	//
	// example:
	//
	// 60
	JobDuration *string `json:"JobDuration,omitempty" xml:"JobDuration,omitempty"`
	// The ID of the failover test.
	//
	// example:
	//
	// ftj-xxxxxxxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Indicates whether the failover test is performed immediately. Valid values:
	//
	// 	- **StartNow**
	//
	// 	- **StartLater**
	//
	// example:
	//
	// StartNow
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The name of the failover test.
	//
	// The name must be 0 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The IDs of the failover test resources.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the failover test resource. Only **PHYSICALCONNECTION*	- is returned.
	//
	// example:
	//
	// PHYSICALCONNECTION
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The beginning of the fault drill task. The time must be in UTC. Specify the time in the ISO 8601 standard in `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// example:
	//
	// 2023-11-21T14:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the failover test. Valid values:
	//
	// 	- **Init**
	//
	// 	- **Starting**
	//
	// 	- **Testing**
	//
	// 	- **Stopping**
	//
	// 	- **Stopped**
	//
	// example:
	//
	// Init
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The end of the fault drill task. The time must be in UTC. Specify the time in the ISO 8601 standard in `YYYY-MM-DDThh:mm:ssZ` format.
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
