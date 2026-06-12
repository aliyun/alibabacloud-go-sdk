// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstanceLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListServiceInstanceLogsRequestFilter) *ListServiceInstanceLogsRequest
	GetFilter() []*ListServiceInstanceLogsRequestFilter
	SetLogSource(v string) *ListServiceInstanceLogsRequest
	GetLogSource() *string
	SetLogstore(v string) *ListServiceInstanceLogsRequest
	GetLogstore() *string
	SetMaxResults(v int32) *ListServiceInstanceLogsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceInstanceLogsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListServiceInstanceLogsRequest
	GetRegionId() *string
	SetServiceInstanceId(v string) *ListServiceInstanceLogsRequest
	GetServiceInstanceId() *string
	SetSortOrder(v string) *ListServiceInstanceLogsRequest
	GetSortOrder() *string
}

type ListServiceInstanceLogsRequest struct {
	// The filters.
	Filter []*ListServiceInstanceLogsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The source of the logs. Valid values:
	//
	// - application: Logs from the application.
	//
	// - computeNest: Logs from Compute Nest.
	//
	// - ros: Logs from Resource Orchestration Service (ROS).
	//
	// example:
	//
	// computeNest
	LogSource *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
	// The name of the Simple Log Service Logstore.
	//
	// example:
	//
	// livelog
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that specifies the next page of results to return. Set this parameter to the NextToken value from a previous response to retrieve the next page of results.
	//
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-70a3b15bb626435b****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The sort order. Valid values:
	//
	// - **Ascending**: Ascending order.
	//
	// - **Descending*	- (default): Descending order.
	//
	// example:
	//
	// Ascending
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListServiceInstanceLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceLogsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceLogsRequest) GetFilter() []*ListServiceInstanceLogsRequestFilter {
	return s.Filter
}

func (s *ListServiceInstanceLogsRequest) GetLogSource() *string {
	return s.LogSource
}

func (s *ListServiceInstanceLogsRequest) GetLogstore() *string {
	return s.Logstore
}

func (s *ListServiceInstanceLogsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceInstanceLogsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceInstanceLogsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListServiceInstanceLogsRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *ListServiceInstanceLogsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListServiceInstanceLogsRequest) SetFilter(v []*ListServiceInstanceLogsRequestFilter) *ListServiceInstanceLogsRequest {
	s.Filter = v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetLogSource(v string) *ListServiceInstanceLogsRequest {
	s.LogSource = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetLogstore(v string) *ListServiceInstanceLogsRequest {
	s.Logstore = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetMaxResults(v int32) *ListServiceInstanceLogsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetNextToken(v string) *ListServiceInstanceLogsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetRegionId(v string) *ListServiceInstanceLogsRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetServiceInstanceId(v string) *ListServiceInstanceLogsRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetSortOrder(v string) *ListServiceInstanceLogsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServiceInstanceLogsRequestFilter struct {
	// The name of the filter condition.
	//
	// Valid values:
	//
	// - StartTime
	//
	// - EndTime
	//
	// - ApplicationGroupName
	//
	// - ResourceName
	//
	// - EventName
	//
	// example:
	//
	// StartTime
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The values for the filter condition. You can specify up to 10 values.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceInstanceLogsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceLogsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceLogsRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListServiceInstanceLogsRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListServiceInstanceLogsRequestFilter) SetName(v string) *ListServiceInstanceLogsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServiceInstanceLogsRequestFilter) SetValue(v []*string) *ListServiceInstanceLogsRequestFilter {
	s.Value = v
	return s
}

func (s *ListServiceInstanceLogsRequestFilter) Validate() error {
	return dara.Validate(s)
}
