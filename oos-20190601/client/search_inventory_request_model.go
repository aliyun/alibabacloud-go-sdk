// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchInventoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregator(v []*string) *SearchInventoryRequest
	GetAggregator() []*string
	SetFilter(v []*SearchInventoryRequestFilter) *SearchInventoryRequest
	GetFilter() []*SearchInventoryRequestFilter
	SetMaxResults(v int32) *SearchInventoryRequest
	GetMaxResults() *int32
	SetNextToken(v string) *SearchInventoryRequest
	GetNextToken() *string
	SetRegionId(v string) *SearchInventoryRequest
	GetRegionId() *string
}

type SearchInventoryRequest struct {
	// The information about aggregators. You can use one or more aggregators to query the aggregate information of an instance. Valid values:
	//
	// 	- ACS:Application.Name
	//
	// 	- ACS:Application.Version
	//
	// example:
	//
	// ACS:Application.Name
	Aggregator []*string `json:"Aggregator,omitempty" xml:"Aggregator,omitempty" type:"Repeated"`
	// The filter rules for the component.
	Filter []*SearchInventoryRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// gAAAAABfTgv5ewUWmNdJ3g7JVLvX70sPH90GZOVGC
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SearchInventoryRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchInventoryRequest) GoString() string {
	return s.String()
}

func (s *SearchInventoryRequest) GetAggregator() []*string {
	return s.Aggregator
}

func (s *SearchInventoryRequest) GetFilter() []*SearchInventoryRequestFilter {
	return s.Filter
}

func (s *SearchInventoryRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SearchInventoryRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchInventoryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchInventoryRequest) SetAggregator(v []*string) *SearchInventoryRequest {
	s.Aggregator = v
	return s
}

func (s *SearchInventoryRequest) SetFilter(v []*SearchInventoryRequestFilter) *SearchInventoryRequest {
	s.Filter = v
	return s
}

func (s *SearchInventoryRequest) SetMaxResults(v int32) *SearchInventoryRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchInventoryRequest) SetNextToken(v string) *SearchInventoryRequest {
	s.NextToken = &v
	return s
}

func (s *SearchInventoryRequest) SetRegionId(v string) *SearchInventoryRequest {
	s.RegionId = &v
	return s
}

func (s *SearchInventoryRequest) Validate() error {
	return dara.Validate(s)
}

type SearchInventoryRequestFilter struct {
	// The name of the component property. Valid values of N: 1 to 5. Different components have different property names. You can call the [GetInventorySchema](https://api.aliyun.com/#/?product=oos\\&version=2019-06-01\\&api=GetInventorySchema) operation to query the property names of different components. For example, the ACS:InstanceInformation component has the InstanceId property. Therefore, you can set this parameter to ACS:InstanceInformation.InstanceId.
	//
	// example:
	//
	// ACS:InstanceInformation.InstanceId
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
	// The property values to query.
	//
	// example:
	//
	// i-bp1cpoxxxxxxxxxxxxxx
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s SearchInventoryRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s SearchInventoryRequestFilter) GoString() string {
	return s.String()
}

func (s *SearchInventoryRequestFilter) GetName() *string {
	return s.Name
}

func (s *SearchInventoryRequestFilter) GetOperator() *string {
	return s.Operator
}

func (s *SearchInventoryRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *SearchInventoryRequestFilter) SetName(v string) *SearchInventoryRequestFilter {
	s.Name = &v
	return s
}

func (s *SearchInventoryRequestFilter) SetOperator(v string) *SearchInventoryRequestFilter {
	s.Operator = &v
	return s
}

func (s *SearchInventoryRequestFilter) SetValue(v []*string) *SearchInventoryRequestFilter {
	s.Value = v
	return s
}

func (s *SearchInventoryRequestFilter) Validate() error {
	return dara.Validate(s)
}
