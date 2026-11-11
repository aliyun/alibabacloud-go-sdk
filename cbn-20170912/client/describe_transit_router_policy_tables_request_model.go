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
	SetTransitRouterId(v string) *DescribeTransitRouterPolicyTablesRequest
	GetTransitRouterId() *string
}

type DescribeTransitRouterPolicyTablesRequest struct {
	// The filter conditions.
	//
	// You can specify up to 5 filter conditions at a time.
	Filter []*DescribeTransitRouterPolicyTablesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries to return on each page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. Valid values:
	//
	// - If this is the first query or no subsequent query is required, you do not need to specify this parameter.
	//
	// - If a subsequent query is required, set the value to the NextToken value returned from the previous API call.
	//
	// example:
	//
	// dd20****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	return nil
}

type DescribeTransitRouterPolicyTablesRequestFilter struct {
	// The filter condition. Valid values:
	//
	// PolicyTableId: the ID of the Cloud Enterprise Network (CEN) instance.
	//
	// By default, multiple filter conditions are combined with the AND operator. A resource is considered a match only if it meets all the filter conditions.
	//
	// You can specify up to 5 filter conditions at a time.
	//
	// example:
	//
	// PolicyTableId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter condition.
	//
	// Specify the filter value based on the key. You can specify multiple values for a single key. Multiple values are combined with the OR operator. A resource is considered a match if it meets any of the values.
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
