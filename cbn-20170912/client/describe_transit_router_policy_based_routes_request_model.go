// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransitRouterPolicyBasedRoutesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*DescribeTransitRouterPolicyBasedRoutesRequestFilter) *DescribeTransitRouterPolicyBasedRoutesRequest
	GetFilter() []*DescribeTransitRouterPolicyBasedRoutesRequestFilter
	SetMaxResults(v int32) *DescribeTransitRouterPolicyBasedRoutesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeTransitRouterPolicyBasedRoutesRequest
	GetNextToken() *string
	SetPolicyTableId(v string) *DescribeTransitRouterPolicyBasedRoutesRequest
	GetPolicyTableId() *string
}

type DescribeTransitRouterPolicyBasedRoutesRequest struct {
	// The filter conditions.
	//
	// You can specify up to 5 filter conditions at a time.
	Filter []*DescribeTransitRouterPolicyBasedRoutesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries to return on each page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results. Valid values:
	//
	// - If **NextToken*	- is empty, no more results are available.
	//
	// - If **NextToken*	- returns a value, the value is the token used to start the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The instance ID of the policy routing table.
	//
	// example:
	//
	// ptbl-abc123xyz
	PolicyTableId *string `json:"PolicyTableId,omitempty" xml:"PolicyTableId,omitempty"`
}

func (s DescribeTransitRouterPolicyBasedRoutesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterPolicyBasedRoutesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterPolicyBasedRoutesRequest) GetFilter() []*DescribeTransitRouterPolicyBasedRoutesRequestFilter {
	return s.Filter
}

func (s *DescribeTransitRouterPolicyBasedRoutesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeTransitRouterPolicyBasedRoutesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTransitRouterPolicyBasedRoutesRequest) GetPolicyTableId() *string {
	return s.PolicyTableId
}

func (s *DescribeTransitRouterPolicyBasedRoutesRequest) SetFilter(v []*DescribeTransitRouterPolicyBasedRoutesRequestFilter) *DescribeTransitRouterPolicyBasedRoutesRequest {
	s.Filter = v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesRequest) SetMaxResults(v int32) *DescribeTransitRouterPolicyBasedRoutesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesRequest) SetNextToken(v string) *DescribeTransitRouterPolicyBasedRoutesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesRequest) SetPolicyTableId(v string) *DescribeTransitRouterPolicyBasedRoutesRequest {
	s.PolicyTableId = &v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesRequest) Validate() error {
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

type DescribeTransitRouterPolicyBasedRoutesRequestFilter struct {
	// The filter condition. Valid values:
	//
	// - **PolicyBasedRouteId**: the instance ID of the policy-based routing rule.
	//
	// By default, multiple filter conditions are evaluated by using the **AND*	- operator. A resource matches the filter only if it meets all the specified conditions.
	//
	// You can specify up to 5 filter conditions at a time.
	//
	// example:
	//
	// PolicyBasedRouteId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter condition.
	//
	// Specify the filter value based on the **Key**. You can specify multiple values for a single **Key**. Multiple values are evaluated by using the **OR*	- operator. A resource matches the filter condition if it meets at least one of the specified values.
	//
	// You can specify up to 5 values for each filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeTransitRouterPolicyBasedRoutesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterPolicyBasedRoutesRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterPolicyBasedRoutesRequestFilter) GetKey() *string {
	return s.Key
}

func (s *DescribeTransitRouterPolicyBasedRoutesRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *DescribeTransitRouterPolicyBasedRoutesRequestFilter) SetKey(v string) *DescribeTransitRouterPolicyBasedRoutesRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesRequestFilter) SetValue(v []*string) *DescribeTransitRouterPolicyBasedRoutesRequestFilter {
	s.Value = v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesRequestFilter) Validate() error {
	return dara.Validate(s)
}
