// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransitRouterPolicyTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*DescribeTransitRouterPolicyTablesRequestFilter) *DescribeTransitRouterPolicyTablesRequest
	GetFilter() []*DescribeTransitRouterPolicyTablesRequestFilter
	SetMaxResults(v int32) *DescribeTransitRouterPolicyTablesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeTransitRouterPolicyTablesRequest
	GetNextToken() *string
	SetTag(v []*DescribeTransitRouterPolicyTablesRequestTag) *DescribeTransitRouterPolicyTablesRequest
	GetTag() []*DescribeTransitRouterPolicyTablesRequestTag
	SetTransitRouterId(v string) *DescribeTransitRouterPolicyTablesRequest
	GetTransitRouterId() *string
}

type DescribeTransitRouterPolicyTablesRequest struct {
	// The filter conditions. You can specify up to 5 filter conditions at a time.
	Filter []*DescribeTransitRouterPolicyTablesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries to return on each page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results. Valid values:
	//
	// If this is your first query or no next query is to be sent, ignore this parameter.
	//
	// If a next query is to be sent, set the value to the NextToken value that is returned from the last API invocation.
	//
	// example:
	//
	// dd20****
	NextToken *string                                        `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Tag       []*DescribeTransitRouterPolicyTablesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the Enterprise Edition transit router.
	//
	// example:
	//
	// tr-uf654ttymmljlvh2x****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s DescribeTransitRouterPolicyTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterPolicyTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterPolicyTablesRequest) GetFilter() []*DescribeTransitRouterPolicyTablesRequestFilter {
	return s.Filter
}

func (s *DescribeTransitRouterPolicyTablesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeTransitRouterPolicyTablesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTransitRouterPolicyTablesRequest) GetTag() []*DescribeTransitRouterPolicyTablesRequestTag {
	return s.Tag
}

func (s *DescribeTransitRouterPolicyTablesRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *DescribeTransitRouterPolicyTablesRequest) SetFilter(v []*DescribeTransitRouterPolicyTablesRequestFilter) *DescribeTransitRouterPolicyTablesRequest {
	s.Filter = v
	return s
}

func (s *DescribeTransitRouterPolicyTablesRequest) SetMaxResults(v int32) *DescribeTransitRouterPolicyTablesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeTransitRouterPolicyTablesRequest) SetNextToken(v string) *DescribeTransitRouterPolicyTablesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeTransitRouterPolicyTablesRequest) SetTag(v []*DescribeTransitRouterPolicyTablesRequestTag) *DescribeTransitRouterPolicyTablesRequest {
	s.Tag = v
	return s
}

func (s *DescribeTransitRouterPolicyTablesRequest) SetTransitRouterId(v string) *DescribeTransitRouterPolicyTablesRequest {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeTransitRouterPolicyTablesRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTransitRouterPolicyTablesRequestFilter struct {
	// The filter condition. Valid values:
	//
	// PolicyTableId: the ID of the Cloud Enterprise Network (CEN) instance.
	//
	// By default, multiple filter conditions are evaluated by using the AND operator. A result is returned only if all filter conditions are met.
	//
	// You can specify up to 5 filter conditions at a time.
	//
	// example:
	//
	// PolicyTableId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter condition.
	//
	// Specify the value based on the key. You can specify multiple values for a key. Multiple values are evaluated by using the OR operator. If one of the values is met, the filter condition is met.
	//
	// You can specify up to 5 values for each filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeTransitRouterPolicyTablesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterPolicyTablesRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterPolicyTablesRequestFilter) GetKey() *string {
	return s.Key
}

func (s *DescribeTransitRouterPolicyTablesRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *DescribeTransitRouterPolicyTablesRequestFilter) SetKey(v string) *DescribeTransitRouterPolicyTablesRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeTransitRouterPolicyTablesRequestFilter) SetValue(v []*string) *DescribeTransitRouterPolicyTablesRequestFilter {
	s.Value = v
	return s
}

func (s *DescribeTransitRouterPolicyTablesRequestFilter) Validate() error {
	return dara.Validate(s)
}

type DescribeTransitRouterPolicyTablesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeTransitRouterPolicyTablesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterPolicyTablesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterPolicyTablesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeTransitRouterPolicyTablesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeTransitRouterPolicyTablesRequestTag) SetKey(v string) *DescribeTransitRouterPolicyTablesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeTransitRouterPolicyTablesRequestTag) SetValue(v string) *DescribeTransitRouterPolicyTablesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeTransitRouterPolicyTablesRequestTag) Validate() error {
	return dara.Validate(s)
}
