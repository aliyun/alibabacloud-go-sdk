// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceBuildLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListServiceBuildLogsRequestFilter) *ListServiceBuildLogsRequest
	GetFilter() []*ListServiceBuildLogsRequestFilter
	SetMaxResults(v int32) *ListServiceBuildLogsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceBuildLogsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListServiceBuildLogsRequest
	GetRegionId() *string
	SetServiceId(v string) *ListServiceBuildLogsRequest
	GetServiceId() *string
	SetSortOrder(v string) *ListServiceBuildLogsRequest
	GetSortOrder() *string
}

type ListServiceBuildLogsRequest struct {
	Filter []*ListServiceBuildLogsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAYChudnQUoBH+mGWFpb6oP0=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// service-39f4f251e94843xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// Ascending
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListServiceBuildLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceBuildLogsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceBuildLogsRequest) GetFilter() []*ListServiceBuildLogsRequestFilter {
	return s.Filter
}

func (s *ListServiceBuildLogsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceBuildLogsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceBuildLogsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListServiceBuildLogsRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListServiceBuildLogsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListServiceBuildLogsRequest) SetFilter(v []*ListServiceBuildLogsRequestFilter) *ListServiceBuildLogsRequest {
	s.Filter = v
	return s
}

func (s *ListServiceBuildLogsRequest) SetMaxResults(v int32) *ListServiceBuildLogsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceBuildLogsRequest) SetNextToken(v string) *ListServiceBuildLogsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceBuildLogsRequest) SetRegionId(v string) *ListServiceBuildLogsRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceBuildLogsRequest) SetServiceId(v string) *ListServiceBuildLogsRequest {
	s.ServiceId = &v
	return s
}

func (s *ListServiceBuildLogsRequest) SetSortOrder(v string) *ListServiceBuildLogsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListServiceBuildLogsRequest) Validate() error {
	return dara.Validate(s)
}

type ListServiceBuildLogsRequestFilter struct {
	// example:
	//
	// Name
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceBuildLogsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListServiceBuildLogsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceBuildLogsRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListServiceBuildLogsRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListServiceBuildLogsRequestFilter) SetName(v string) *ListServiceBuildLogsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServiceBuildLogsRequestFilter) SetValue(v []*string) *ListServiceBuildLogsRequestFilter {
	s.Value = v
	return s
}

func (s *ListServiceBuildLogsRequestFilter) Validate() error {
	return dara.Validate(s)
}
