// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRouteMapRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListRouteMapRulesRequest
	GetClientToken() *string
	SetDryRun(v bool) *ListRouteMapRulesRequest
	GetDryRun() *bool
	SetFilter(v []*ListRouteMapRulesRequestFilter) *ListRouteMapRulesRequest
	GetFilter() []*ListRouteMapRulesRequestFilter
	SetMaxResults(v int32) *ListRouteMapRulesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListRouteMapRulesRequest
	GetNextToken() *string
	SetRouteMapId(v string) *ListRouteMapRulesRequest
	GetRouteMapId() *string
	SetRouteMapRuleIds(v []*string) *ListRouteMapRulesRequest
	GetRouteMapRuleIds() []*string
}

type ListRouteMapRulesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate a parameter value on your client to ensure that the parameter value is unique across different requests. The ClientToken parameter supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// - **true**: sends a check request, and no route map rules are queried. Check items include whether required parameters are specified, the request format, and service limits. If the check fails, the corresponding error is returned. If the check succeeds, the error code DryRunOperation is returned.
	//
	// - **false*	- (default): sends a normal request, and route map rules are queried directly after the check succeeds.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The filter conditions.
	//
	// You can specify up to 5 filter conditions at a time.
	Filter []*ListRouteMapRulesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results. Valid values:
	//
	// - If this is the first query or no next page exists, you do not need to specify this parameter.
	//
	// - If a next page exists, set the value to the **NextToken*	- value returned from the previous API call.
	//
	// example:
	//
	// dd20****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the route map to query.
	//
	// example:
	//
	// rm-0a1b2c3d4e5f6a7b8
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
	// The list of route map rule IDs to query. This parameter and filter conditions can be used together for matching, and the relationship is AND.
	RouteMapRuleIds []*string `json:"RouteMapRuleIds,omitempty" xml:"RouteMapRuleIds,omitempty" type:"Repeated"`
}

func (s ListRouteMapRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRouteMapRulesRequest) GoString() string {
	return s.String()
}

func (s *ListRouteMapRulesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListRouteMapRulesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ListRouteMapRulesRequest) GetFilter() []*ListRouteMapRulesRequestFilter {
	return s.Filter
}

func (s *ListRouteMapRulesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRouteMapRulesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRouteMapRulesRequest) GetRouteMapId() *string {
	return s.RouteMapId
}

func (s *ListRouteMapRulesRequest) GetRouteMapRuleIds() []*string {
	return s.RouteMapRuleIds
}

func (s *ListRouteMapRulesRequest) SetClientToken(v string) *ListRouteMapRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListRouteMapRulesRequest) SetDryRun(v bool) *ListRouteMapRulesRequest {
	s.DryRun = &v
	return s
}

func (s *ListRouteMapRulesRequest) SetFilter(v []*ListRouteMapRulesRequestFilter) *ListRouteMapRulesRequest {
	s.Filter = v
	return s
}

func (s *ListRouteMapRulesRequest) SetMaxResults(v int32) *ListRouteMapRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRouteMapRulesRequest) SetNextToken(v string) *ListRouteMapRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListRouteMapRulesRequest) SetRouteMapId(v string) *ListRouteMapRulesRequest {
	s.RouteMapId = &v
	return s
}

func (s *ListRouteMapRulesRequest) SetRouteMapRuleIds(v []*string) *ListRouteMapRulesRequest {
	s.RouteMapRuleIds = v
	return s
}

func (s *ListRouteMapRulesRequest) Validate() error {
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

type ListRouteMapRulesRequestFilter struct {
	// The filter condition. Valid values:
	//
	// - **Name**: the name of the route map rule.
	//
	// - **Priority**: the priority of the route map rule.
	//
	// By default, multiple filter conditions are evaluated by using the AND logical operator. A rule matches the filter conditions only if it meets all the filter conditions.
	//
	// You can specify up to 5 filter conditions at a time.
	//
	// example:
	//
	// Name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter condition.
	//
	// Specify the filter value based on the Key. You can specify multiple filter values for a single Key. Multiple filter values are evaluated by using the OR logical operator. A rule matches the filter condition if it meets any of the filter values.
	//
	// You can specify up to 5 filter values for a single filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListRouteMapRulesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListRouteMapRulesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListRouteMapRulesRequestFilter) GetKey() *string {
	return s.Key
}

func (s *ListRouteMapRulesRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListRouteMapRulesRequestFilter) SetKey(v string) *ListRouteMapRulesRequestFilter {
	s.Key = &v
	return s
}

func (s *ListRouteMapRulesRequestFilter) SetValue(v []*string) *ListRouteMapRulesRequestFilter {
	s.Value = v
	return s
}

func (s *ListRouteMapRulesRequestFilter) Validate() error {
	return dara.Validate(s)
}
