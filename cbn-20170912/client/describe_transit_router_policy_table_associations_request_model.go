// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransitRouterPolicyTableAssociationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*DescribeTransitRouterPolicyTableAssociationsRequestFilter) *DescribeTransitRouterPolicyTableAssociationsRequest
	GetFilter() []*DescribeTransitRouterPolicyTableAssociationsRequestFilter
	SetMaxResults(v int32) *DescribeTransitRouterPolicyTableAssociationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeTransitRouterPolicyTableAssociationsRequest
	GetNextToken() *string
}

type DescribeTransitRouterPolicyTableAssociationsRequest struct {
	// The filter conditions. You can specify up to five filter conditions at a time.
	Filter []*DescribeTransitRouterPolicyTableAssociationsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries to return per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. Valid values:
	//
	// - If **NextToken*	- is empty, no subsequent query is available.
	//
	// - If a value is returned for **NextToken**, the value is the token for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeTransitRouterPolicyTableAssociationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterPolicyTableAssociationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterPolicyTableAssociationsRequest) GetFilter() []*DescribeTransitRouterPolicyTableAssociationsRequestFilter {
	return s.Filter
}

func (s *DescribeTransitRouterPolicyTableAssociationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeTransitRouterPolicyTableAssociationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTransitRouterPolicyTableAssociationsRequest) SetFilter(v []*DescribeTransitRouterPolicyTableAssociationsRequestFilter) *DescribeTransitRouterPolicyTableAssociationsRequest {
	s.Filter = v
	return s
}

func (s *DescribeTransitRouterPolicyTableAssociationsRequest) SetMaxResults(v int32) *DescribeTransitRouterPolicyTableAssociationsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeTransitRouterPolicyTableAssociationsRequest) SetNextToken(v string) *DescribeTransitRouterPolicyTableAssociationsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeTransitRouterPolicyTableAssociationsRequest) Validate() error {
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

type DescribeTransitRouterPolicyTableAssociationsRequestFilter struct {
	// The filter condition. Valid values:
	//
	// - **PolicyTableId**: the policy routing table ID.
	//
	// - **AttachmentId**: the network connectivity instance ID.
	//
	// - **AssociationId**: the policy routing table associate relationship instance ID.
	//
	// By default, multiple filter conditions are connected by the **AND*	- logical operator. A result is returned only when all filter conditions are met.
	//
	// You can specify up to five filter conditions at a time.
	//
	// example:
	//
	// CenId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter condition.
	//
	// Specify the filter value based on the **Key*	- parameter. You can specify multiple filter values for a single **Key**. Multiple filter values are connected by the **OR*	- logical operator. A result is returned when any of the filter values is met.
	//
	// You can specify up to five filter values for a single filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeTransitRouterPolicyTableAssociationsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterPolicyTableAssociationsRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterPolicyTableAssociationsRequestFilter) GetKey() *string {
	return s.Key
}

func (s *DescribeTransitRouterPolicyTableAssociationsRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *DescribeTransitRouterPolicyTableAssociationsRequestFilter) SetKey(v string) *DescribeTransitRouterPolicyTableAssociationsRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeTransitRouterPolicyTableAssociationsRequestFilter) SetValue(v []*string) *DescribeTransitRouterPolicyTableAssociationsRequestFilter {
	s.Value = v
	return s
}

func (s *DescribeTransitRouterPolicyTableAssociationsRequestFilter) Validate() error {
	return dara.Validate(s)
}
