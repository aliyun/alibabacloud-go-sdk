// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceTestTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListServiceTestTasksRequestFilter) *ListServiceTestTasksRequest
	GetFilter() []*ListServiceTestTasksRequestFilter
	SetMaxResults(v int32) *ListServiceTestTasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceTestTasksRequest
	GetNextToken() *string
	SetRegionId(v string) *ListServiceTestTasksRequest
	GetRegionId() *string
	SetServiceId(v string) *ListServiceTestTasksRequest
	GetServiceId() *string
	SetServiceVersion(v string) *ListServiceTestTasksRequest
	GetServiceVersion() *string
}

type ListServiceTestTasksRequest struct {
	// One or more filters for the query.
	Filter []*ListServiceTestTasksRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries to return on each page. The maximum value is 100. The default value is 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results. Set this parameter to the value of NextToken returned from the previous API call.
	//
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-062ae8e13b394dd5b63c
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// draft
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s ListServiceTestTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceTestTasksRequest) GoString() string {
	return s.String()
}

func (s *ListServiceTestTasksRequest) GetFilter() []*ListServiceTestTasksRequestFilter {
	return s.Filter
}

func (s *ListServiceTestTasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceTestTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceTestTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListServiceTestTasksRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListServiceTestTasksRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *ListServiceTestTasksRequest) SetFilter(v []*ListServiceTestTasksRequestFilter) *ListServiceTestTasksRequest {
	s.Filter = v
	return s
}

func (s *ListServiceTestTasksRequest) SetMaxResults(v int32) *ListServiceTestTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceTestTasksRequest) SetNextToken(v string) *ListServiceTestTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceTestTasksRequest) SetRegionId(v string) *ListServiceTestTasksRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceTestTasksRequest) SetServiceId(v string) *ListServiceTestTasksRequest {
	s.ServiceId = &v
	return s
}

func (s *ListServiceTestTasksRequest) SetServiceVersion(v string) *ListServiceTestTasksRequest {
	s.ServiceVersion = &v
	return s
}

func (s *ListServiceTestTasksRequest) Validate() error {
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

type ListServiceTestTasksRequestFilter struct {
	// The property to filter by. Valid values:
	//
	// - Status: The task status.
	//
	// - TaskId: The task ID.
	//
	// example:
	//
	// Status
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// A list of filter values.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceTestTasksRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListServiceTestTasksRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceTestTasksRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListServiceTestTasksRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListServiceTestTasksRequestFilter) SetName(v string) *ListServiceTestTasksRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServiceTestTasksRequestFilter) SetValue(v []*string) *ListServiceTestTasksRequestFilter {
	s.Value = v
	return s
}

func (s *ListServiceTestTasksRequestFilter) Validate() error {
	return dara.Validate(s)
}
