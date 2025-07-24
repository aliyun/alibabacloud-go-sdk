// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIds(v []*string) *ListDoctorApplicationsRequest
	GetAppIds() []*string
	SetClusterId(v string) *ListDoctorApplicationsRequest
	GetClusterId() *string
	SetDateTime(v string) *ListDoctorApplicationsRequest
	GetDateTime() *string
	SetMaxResults(v int32) *ListDoctorApplicationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDoctorApplicationsRequest
	GetNextToken() *string
	SetOrderBy(v string) *ListDoctorApplicationsRequest
	GetOrderBy() *string
	SetOrderType(v string) *ListDoctorApplicationsRequest
	GetOrderType() *string
	SetQueues(v []*string) *ListDoctorApplicationsRequest
	GetQueues() []*string
	SetRegionId(v string) *ListDoctorApplicationsRequest
	GetRegionId() *string
	SetTypes(v []*string) *ListDoctorApplicationsRequest
	GetTypes() []*string
	SetUsers(v []*string) *ListDoctorApplicationsRequest
	GetUsers() []*string
}

type ListDoctorApplicationsRequest struct {
	// The IDs of jobs that are submitted to YARN.
	//
	// example:
	//
	// null
	AppIds []*string `json:"AppIds,omitempty" xml:"AppIds,omitempty" type:"Repeated"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specify the date in the ISO 8601 standard. For example, 2023-01-01 represents January 1, 2023.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-01
	DateTime *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the request to retrieve a new page of results.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The field that you use to sort the query results. Valid values:
	//
	// 1.  startTime: the time when the job starts
	//
	// 2.  endTime: the time when the job ends
	//
	// 3.  vcoreUtilization: the vCPU utilization of the job
	//
	// 4.  memUtilization: the memory usage of the job
	//
	// 5.  vcoreSeconds: the aggregated number of vCPUs that are allocated to the job multiplied by the number of seconds the job has been running
	//
	// 6.  memSeconds: the aggregated amount of memory that is allocated to the job multiplied by the number of seconds the job has been running
	//
	// 7.  score: the score of the job
	//
	// example:
	//
	// score
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The order in which you want to sort the query results. Valid values:
	//
	// 	- ASC: the ascending order
	//
	// 	- DESC: the descending order
	//
	// example:
	//
	// ASC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The YARN queues to which the jobs are submitted.
	//
	// example:
	//
	// null
	Queues []*string `json:"Queues,omitempty" xml:"Queues,omitempty" type:"Repeated"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The YARN engines to which the jobs are submitted.
	//
	// example:
	//
	// null
	Types []*string `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
	// The users who submit the jobs.
	//
	// example:
	//
	// null
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListDoctorApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListDoctorApplicationsRequest) GetAppIds() []*string {
	return s.AppIds
}

func (s *ListDoctorApplicationsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListDoctorApplicationsRequest) GetDateTime() *string {
	return s.DateTime
}

func (s *ListDoctorApplicationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDoctorApplicationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDoctorApplicationsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListDoctorApplicationsRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListDoctorApplicationsRequest) GetQueues() []*string {
	return s.Queues
}

func (s *ListDoctorApplicationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDoctorApplicationsRequest) GetTypes() []*string {
	return s.Types
}

func (s *ListDoctorApplicationsRequest) GetUsers() []*string {
	return s.Users
}

func (s *ListDoctorApplicationsRequest) SetAppIds(v []*string) *ListDoctorApplicationsRequest {
	s.AppIds = v
	return s
}

func (s *ListDoctorApplicationsRequest) SetClusterId(v string) *ListDoctorApplicationsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListDoctorApplicationsRequest) SetDateTime(v string) *ListDoctorApplicationsRequest {
	s.DateTime = &v
	return s
}

func (s *ListDoctorApplicationsRequest) SetMaxResults(v int32) *ListDoctorApplicationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDoctorApplicationsRequest) SetNextToken(v string) *ListDoctorApplicationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDoctorApplicationsRequest) SetOrderBy(v string) *ListDoctorApplicationsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListDoctorApplicationsRequest) SetOrderType(v string) *ListDoctorApplicationsRequest {
	s.OrderType = &v
	return s
}

func (s *ListDoctorApplicationsRequest) SetQueues(v []*string) *ListDoctorApplicationsRequest {
	s.Queues = v
	return s
}

func (s *ListDoctorApplicationsRequest) SetRegionId(v string) *ListDoctorApplicationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListDoctorApplicationsRequest) SetTypes(v []*string) *ListDoctorApplicationsRequest {
	s.Types = v
	return s
}

func (s *ListDoctorApplicationsRequest) SetUsers(v []*string) *ListDoctorApplicationsRequest {
	s.Users = v
	return s
}

func (s *ListDoctorApplicationsRequest) Validate() error {
	return dara.Validate(s)
}
