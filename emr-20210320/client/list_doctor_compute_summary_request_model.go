// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorComputeSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListDoctorComputeSummaryRequest
	GetClusterId() *string
	SetComponentTypes(v []*string) *ListDoctorComputeSummaryRequest
	GetComponentTypes() []*string
	SetDateTime(v string) *ListDoctorComputeSummaryRequest
	GetDateTime() *string
	SetMaxResults(v int32) *ListDoctorComputeSummaryRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDoctorComputeSummaryRequest
	GetNextToken() *string
	SetOrderBy(v string) *ListDoctorComputeSummaryRequest
	GetOrderBy() *string
	SetOrderType(v string) *ListDoctorComputeSummaryRequest
	GetOrderType() *string
	SetRegionId(v string) *ListDoctorComputeSummaryRequest
	GetRegionId() *string
}

type ListDoctorComputeSummaryRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The resource types, which are used to filter query results. Valid values:
	//
	// 	- engine: filters results by engine.
	//
	// 	- queue: filters results by queue.
	//
	// 	- cluster: displays the results at the cluster level.
	//
	// If you do not specify this parameter, the information at the cluster level is displayed by default. Currently, only one resource type is supported. If you specify multiple resource types, the first resource type is used by default.
	//
	// example:
	//
	// null
	ComponentTypes []*string `json:"ComponentTypes,omitempty" xml:"ComponentTypes,omitempty" type:"Repeated"`
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
	// The basis on which you want to sort the query results. Valid values:
	//
	// 1.  vcoreSeconds: the total CPU consumption over time in seconds.
	//
	// 2.  memSeconds: the total memory consumption over time in seconds.
	//
	// 3.  vcoreUtilization: the average CPU utilization. The meaning is the same as the %CPU parameter in the output of the top command in Linux.
	//
	// 4.  memUtilization: the average memory usage.
	//
	// 5.  vcoreSecondsDayGrowthRatio: the day-to-day growth rate of the total CPU consumption over time in seconds.
	//
	// 6.  memSecondsDayGrowthRatio: the day-to-day growth rate of the total memory consumption over time in seconds.
	//
	// 7.  readSize: the total amount of data read from the file system.
	//
	// 8.  writeSize: the total amount of data written to the file system.
	//
	// 9.  healthyJobCount: the total number of healthy jobs.
	//
	// 10. subHealthyJobCount: the total number of sub-healthy jobs.
	//
	// 11. unhealthyJobCount: the total number of unhealthy jobs.
	//
	// 12. needAttentionJobCount: the total number of jobs that require attention.
	//
	// 13. score: the score for jobs.
	//
	// 14. scoreDayGrowthRatio: the day-to-day growth rate of the score for jobs.
	//
	// example:
	//
	// score
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The order in which you want to sort the query results. Valid values:
	//
	// 	- ASC: in ascending order.
	//
	// 	- DESC: in descending order.
	//
	// example:
	//
	// ASC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListDoctorComputeSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorComputeSummaryRequest) GoString() string {
	return s.String()
}

func (s *ListDoctorComputeSummaryRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListDoctorComputeSummaryRequest) GetComponentTypes() []*string {
	return s.ComponentTypes
}

func (s *ListDoctorComputeSummaryRequest) GetDateTime() *string {
	return s.DateTime
}

func (s *ListDoctorComputeSummaryRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDoctorComputeSummaryRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDoctorComputeSummaryRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListDoctorComputeSummaryRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListDoctorComputeSummaryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDoctorComputeSummaryRequest) SetClusterId(v string) *ListDoctorComputeSummaryRequest {
	s.ClusterId = &v
	return s
}

func (s *ListDoctorComputeSummaryRequest) SetComponentTypes(v []*string) *ListDoctorComputeSummaryRequest {
	s.ComponentTypes = v
	return s
}

func (s *ListDoctorComputeSummaryRequest) SetDateTime(v string) *ListDoctorComputeSummaryRequest {
	s.DateTime = &v
	return s
}

func (s *ListDoctorComputeSummaryRequest) SetMaxResults(v int32) *ListDoctorComputeSummaryRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDoctorComputeSummaryRequest) SetNextToken(v string) *ListDoctorComputeSummaryRequest {
	s.NextToken = &v
	return s
}

func (s *ListDoctorComputeSummaryRequest) SetOrderBy(v string) *ListDoctorComputeSummaryRequest {
	s.OrderBy = &v
	return s
}

func (s *ListDoctorComputeSummaryRequest) SetOrderType(v string) *ListDoctorComputeSummaryRequest {
	s.OrderType = &v
	return s
}

func (s *ListDoctorComputeSummaryRequest) SetRegionId(v string) *ListDoctorComputeSummaryRequest {
	s.RegionId = &v
	return s
}

func (s *ListDoctorComputeSummaryRequest) Validate() error {
	return dara.Validate(s)
}
