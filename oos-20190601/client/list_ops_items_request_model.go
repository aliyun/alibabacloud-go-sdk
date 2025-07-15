// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOpsItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListOpsItemsRequestFilter) *ListOpsItemsRequest
	GetFilter() []*ListOpsItemsRequestFilter
	SetMaxResults(v int32) *ListOpsItemsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListOpsItemsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListOpsItemsRequest
	GetRegionId() *string
	SetResourceTags(v map[string]interface{}) *ListOpsItemsRequest
	GetResourceTags() map[string]interface{}
	SetTags(v map[string]interface{}) *ListOpsItemsRequest
	GetTags() map[string]interface{}
}

type ListOpsItemsRequest struct {
	// The filter rules for the component.
	Filter []*ListOpsItemsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries to return on each page. Valid values: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ3NjFENDdB
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The information about resource tags.
	//
	// example:
	//
	// {
	//
	//       "k1": "v1",
	//
	//       "k2": "v2"
	//
	// }
	ResourceTags map[string]interface{} `json:"ResourceTags,omitempty" xml:"ResourceTags,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListOpsItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOpsItemsRequest) GoString() string {
	return s.String()
}

func (s *ListOpsItemsRequest) GetFilter() []*ListOpsItemsRequestFilter {
	return s.Filter
}

func (s *ListOpsItemsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListOpsItemsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOpsItemsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListOpsItemsRequest) GetResourceTags() map[string]interface{} {
	return s.ResourceTags
}

func (s *ListOpsItemsRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *ListOpsItemsRequest) SetFilter(v []*ListOpsItemsRequestFilter) *ListOpsItemsRequest {
	s.Filter = v
	return s
}

func (s *ListOpsItemsRequest) SetMaxResults(v int32) *ListOpsItemsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListOpsItemsRequest) SetNextToken(v string) *ListOpsItemsRequest {
	s.NextToken = &v
	return s
}

func (s *ListOpsItemsRequest) SetRegionId(v string) *ListOpsItemsRequest {
	s.RegionId = &v
	return s
}

func (s *ListOpsItemsRequest) SetResourceTags(v map[string]interface{}) *ListOpsItemsRequest {
	s.ResourceTags = v
	return s
}

func (s *ListOpsItemsRequest) SetTags(v map[string]interface{}) *ListOpsItemsRequest {
	s.Tags = v
	return s
}

func (s *ListOpsItemsRequest) Validate() error {
	return dara.Validate(s)
}

type ListOpsItemsRequestFilter struct {
	// The parameter name of the filter.
	//
	// example:
	//
	// Status
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The comparison operator that is used to filter property values.
	//
	// example:
	//
	// Equal
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The parameter values of the filter.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListOpsItemsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListOpsItemsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListOpsItemsRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListOpsItemsRequestFilter) GetOperator() *string {
	return s.Operator
}

func (s *ListOpsItemsRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListOpsItemsRequestFilter) SetName(v string) *ListOpsItemsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListOpsItemsRequestFilter) SetOperator(v string) *ListOpsItemsRequestFilter {
	s.Operator = &v
	return s
}

func (s *ListOpsItemsRequestFilter) SetValue(v []*string) *ListOpsItemsRequestFilter {
	s.Value = v
	return s
}

func (s *ListOpsItemsRequestFilter) Validate() error {
	return dara.Validate(s)
}
