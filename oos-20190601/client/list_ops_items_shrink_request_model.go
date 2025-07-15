// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOpsItemsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListOpsItemsShrinkRequestFilter) *ListOpsItemsShrinkRequest
	GetFilter() []*ListOpsItemsShrinkRequestFilter
	SetMaxResults(v int32) *ListOpsItemsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListOpsItemsShrinkRequest
	GetNextToken() *string
	SetRegionId(v string) *ListOpsItemsShrinkRequest
	GetRegionId() *string
	SetResourceTagsShrink(v string) *ListOpsItemsShrinkRequest
	GetResourceTagsShrink() *string
	SetTagsShrink(v string) *ListOpsItemsShrinkRequest
	GetTagsShrink() *string
}

type ListOpsItemsShrinkRequest struct {
	// The filter rules for the component.
	Filter []*ListOpsItemsShrinkRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
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
	ResourceTagsShrink *string `json:"ResourceTags,omitempty" xml:"ResourceTags,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListOpsItemsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOpsItemsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListOpsItemsShrinkRequest) GetFilter() []*ListOpsItemsShrinkRequestFilter {
	return s.Filter
}

func (s *ListOpsItemsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListOpsItemsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOpsItemsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListOpsItemsShrinkRequest) GetResourceTagsShrink() *string {
	return s.ResourceTagsShrink
}

func (s *ListOpsItemsShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListOpsItemsShrinkRequest) SetFilter(v []*ListOpsItemsShrinkRequestFilter) *ListOpsItemsShrinkRequest {
	s.Filter = v
	return s
}

func (s *ListOpsItemsShrinkRequest) SetMaxResults(v int32) *ListOpsItemsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListOpsItemsShrinkRequest) SetNextToken(v string) *ListOpsItemsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListOpsItemsShrinkRequest) SetRegionId(v string) *ListOpsItemsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListOpsItemsShrinkRequest) SetResourceTagsShrink(v string) *ListOpsItemsShrinkRequest {
	s.ResourceTagsShrink = &v
	return s
}

func (s *ListOpsItemsShrinkRequest) SetTagsShrink(v string) *ListOpsItemsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListOpsItemsShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type ListOpsItemsShrinkRequestFilter struct {
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

func (s ListOpsItemsShrinkRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListOpsItemsShrinkRequestFilter) GoString() string {
	return s.String()
}

func (s *ListOpsItemsShrinkRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListOpsItemsShrinkRequestFilter) GetOperator() *string {
	return s.Operator
}

func (s *ListOpsItemsShrinkRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListOpsItemsShrinkRequestFilter) SetName(v string) *ListOpsItemsShrinkRequestFilter {
	s.Name = &v
	return s
}

func (s *ListOpsItemsShrinkRequestFilter) SetOperator(v string) *ListOpsItemsShrinkRequestFilter {
	s.Operator = &v
	return s
}

func (s *ListOpsItemsShrinkRequestFilter) SetValue(v []*string) *ListOpsItemsShrinkRequestFilter {
	s.Value = v
	return s
}

func (s *ListOpsItemsShrinkRequestFilter) Validate() error {
	return dara.Validate(s)
}
