// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstanceDeployDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCycleTimeZone(v string) *ListServiceInstanceDeployDetailsRequest
	GetCycleTimeZone() *string
	SetCycleType(v string) *ListServiceInstanceDeployDetailsRequest
	GetCycleType() *string
	SetDimension(v []*string) *ListServiceInstanceDeployDetailsRequest
	GetDimension() []*string
	SetEndTime(v string) *ListServiceInstanceDeployDetailsRequest
	GetEndTime() *string
	SetFilter(v []*ListServiceInstanceDeployDetailsRequestFilter) *ListServiceInstanceDeployDetailsRequest
	GetFilter() []*ListServiceInstanceDeployDetailsRequestFilter
	SetMaxResults(v int32) *ListServiceInstanceDeployDetailsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceInstanceDeployDetailsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListServiceInstanceDeployDetailsRequest
	GetRegionId() *string
	SetStartTime(v string) *ListServiceInstanceDeployDetailsRequest
	GetStartTime() *string
}

type ListServiceInstanceDeployDetailsRequest struct {
	// The time zone.
	//
	// Reference Format: "+08:00"
	//
	// Valid Range: "-12:59" to "+13:00"
	//
	// example:
	//
	// +08:00
	CycleTimeZone *string `json:"CycleTimeZone,omitempty" xml:"CycleTimeZone,omitempty"`
	// Determines the time period over which data is aggregated. If no aggregation dimension is specified, the query defaults to providing detailed, unaggregated results.
	//
	// Optional Values:
	//
	// - Year
	//
	// - Month
	//
	// - Day
	//
	// - All
	//
	// example:
	//
	// Month
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// The dimension names. (Equivalent to SQL\\"s GROUP BY Clause)
	//
	// Optional Values:
	//
	// - UserId
	//
	// - ServiceId
	//
	// - ServiceVersion
	//
	// - ServiceInstanceId
	//
	// - DeploySucceeded
	//
	// - ErrorType
	//
	// - ErrorCode
	Dimension []*string `json:"Dimension,omitempty" xml:"Dimension,omitempty" type:"Repeated"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2024-12-31T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The filter.
	Filter []*ListServiceInstanceDeployDetailsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAZbOYA+x9UgM6xrgcMqFUjk=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2024-08-25T02:23:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListServiceInstanceDeployDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceDeployDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceDeployDetailsRequest) GetCycleTimeZone() *string {
	return s.CycleTimeZone
}

func (s *ListServiceInstanceDeployDetailsRequest) GetCycleType() *string {
	return s.CycleType
}

func (s *ListServiceInstanceDeployDetailsRequest) GetDimension() []*string {
	return s.Dimension
}

func (s *ListServiceInstanceDeployDetailsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListServiceInstanceDeployDetailsRequest) GetFilter() []*ListServiceInstanceDeployDetailsRequestFilter {
	return s.Filter
}

func (s *ListServiceInstanceDeployDetailsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceInstanceDeployDetailsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceInstanceDeployDetailsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListServiceInstanceDeployDetailsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListServiceInstanceDeployDetailsRequest) SetCycleTimeZone(v string) *ListServiceInstanceDeployDetailsRequest {
	s.CycleTimeZone = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsRequest) SetCycleType(v string) *ListServiceInstanceDeployDetailsRequest {
	s.CycleType = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsRequest) SetDimension(v []*string) *ListServiceInstanceDeployDetailsRequest {
	s.Dimension = v
	return s
}

func (s *ListServiceInstanceDeployDetailsRequest) SetEndTime(v string) *ListServiceInstanceDeployDetailsRequest {
	s.EndTime = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsRequest) SetFilter(v []*ListServiceInstanceDeployDetailsRequestFilter) *ListServiceInstanceDeployDetailsRequest {
	s.Filter = v
	return s
}

func (s *ListServiceInstanceDeployDetailsRequest) SetMaxResults(v int32) *ListServiceInstanceDeployDetailsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsRequest) SetNextToken(v string) *ListServiceInstanceDeployDetailsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsRequest) SetRegionId(v string) *ListServiceInstanceDeployDetailsRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsRequest) SetStartTime(v string) *ListServiceInstanceDeployDetailsRequest {
	s.StartTime = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsRequest) Validate() error {
	return dara.Validate(s)
}

type ListServiceInstanceDeployDetailsRequestFilter struct {
	// Filter Value Names (Equivalent to SQL\\"s WHERE Clause)
	//
	// Available Options:
	//
	// - UserId
	//
	// - ServiceId
	//
	// - ServiceVersion
	//
	// - ServiceInstanceId
	//
	// - DeploySucceeded (Accepts True or False and case-insensitive)
	//
	// - ErrorType
	//
	// - ErrorCode
	//
	// example:
	//
	// ServiceId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// A value of the filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceInstanceDeployDetailsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceDeployDetailsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceDeployDetailsRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListServiceInstanceDeployDetailsRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListServiceInstanceDeployDetailsRequestFilter) SetName(v string) *ListServiceInstanceDeployDetailsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsRequestFilter) SetValue(v []*string) *ListServiceInstanceDeployDetailsRequestFilter {
	s.Value = v
	return s
}

func (s *ListServiceInstanceDeployDetailsRequestFilter) Validate() error {
	return dara.Validate(s)
}
