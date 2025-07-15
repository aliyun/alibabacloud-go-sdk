// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInventoryEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListInventoryEntriesRequestFilter) *ListInventoryEntriesRequest
	GetFilter() []*ListInventoryEntriesRequestFilter
	SetInstanceId(v string) *ListInventoryEntriesRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListInventoryEntriesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListInventoryEntriesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListInventoryEntriesRequest
	GetRegionId() *string
	SetTypeName(v string) *ListInventoryEntriesRequest
	GetTypeName() *string
}

type ListInventoryEntriesRequest struct {
	// The filter rules for the component.
	Filter []*ListInventoryEntriesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1cpoxxxwxxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzA
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the component. Valid values:
	//
	// 	- ACS:InstanceInformation
	//
	// 	- ACS:Application
	//
	// 	- ACS:File
	//
	// 	- ACS:Network
	//
	// 	- ACS:WindowsRole
	//
	// 	- ACS:Service
	//
	// 	- ACS:WindowsRegistry
	//
	// 	- ACS:WindowsUpdate
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS:InstanceInformation
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s ListInventoryEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInventoryEntriesRequest) GoString() string {
	return s.String()
}

func (s *ListInventoryEntriesRequest) GetFilter() []*ListInventoryEntriesRequestFilter {
	return s.Filter
}

func (s *ListInventoryEntriesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInventoryEntriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInventoryEntriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInventoryEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInventoryEntriesRequest) GetTypeName() *string {
	return s.TypeName
}

func (s *ListInventoryEntriesRequest) SetFilter(v []*ListInventoryEntriesRequestFilter) *ListInventoryEntriesRequest {
	s.Filter = v
	return s
}

func (s *ListInventoryEntriesRequest) SetInstanceId(v string) *ListInventoryEntriesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInventoryEntriesRequest) SetMaxResults(v int32) *ListInventoryEntriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInventoryEntriesRequest) SetNextToken(v string) *ListInventoryEntriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListInventoryEntriesRequest) SetRegionId(v string) *ListInventoryEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *ListInventoryEntriesRequest) SetTypeName(v string) *ListInventoryEntriesRequest {
	s.TypeName = &v
	return s
}

func (s *ListInventoryEntriesRequest) Validate() error {
	return dara.Validate(s)
}

type ListInventoryEntriesRequestFilter struct {
	// The name of the component property. Valid values of N: 1 to 5.
	//
	// example:
	//
	// PlatformName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The comparison operator that is used to filter property values. Valid values of N: 1 to 5. Valid values:
	//
	// 	- Equal
	//
	// 	- NotEqual
	//
	// 	- BeginWith
	//
	// 	- LessThan
	//
	// 	- GreaterThan
	//
	// example:
	//
	// Equal
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The values of properties. Valid values of the first N: 1 to 5. Valid values of the second N: 1 to 20.
	//
	// example:
	//
	// test
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListInventoryEntriesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListInventoryEntriesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListInventoryEntriesRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListInventoryEntriesRequestFilter) GetOperator() *string {
	return s.Operator
}

func (s *ListInventoryEntriesRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListInventoryEntriesRequestFilter) SetName(v string) *ListInventoryEntriesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListInventoryEntriesRequestFilter) SetOperator(v string) *ListInventoryEntriesRequestFilter {
	s.Operator = &v
	return s
}

func (s *ListInventoryEntriesRequestFilter) SetValue(v []*string) *ListInventoryEntriesRequestFilter {
	s.Value = v
	return s
}

func (s *ListInventoryEntriesRequestFilter) Validate() error {
	return dara.Validate(s)
}
