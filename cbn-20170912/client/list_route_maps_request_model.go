// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRouteMapsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListRouteMapsRequest
	GetClientToken() *string
	SetDryRun(v bool) *ListRouteMapsRequest
	GetDryRun() *bool
	SetFilter(v []*ListRouteMapsRequestFilter) *ListRouteMapsRequest
	GetFilter() []*ListRouteMapsRequestFilter
	SetMaxResults(v int32) *ListRouteMapsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListRouteMapsRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListRouteMapsRequest
	GetResourceGroupId() *string
	SetRouteMapIds(v []*string) *ListRouteMapsRequest
	GetRouteMapIds() []*string
}

type ListRouteMapsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate a value for this parameter on your client to ensure that the value is unique across different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the client token. The RequestId of each API request may be different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// - **true**: sends a check request without performing the query. The system checks whether required parameters are specified, the request format, and service limits. If the check fails, the corresponding error is returned. If the check succeeds, the error code `DryRunOperation` is returned.
	//
	// - **false*	- (default): sends a normal request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The filter conditions.
	//
	// You can specify up to 5 filter conditions at a time.
	Filter []*ListRouteMapsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries to return on each page. Valid values: 1 to 100.
	//
	// If you do not specify the MaxResults parameter, pagination is not used. After you call the operation, the value of MaxResults indicates the total number of entries. If you specify the MaxResults parameter, pagination is used. After you call the operation, the value of MaxResults indicates the number of entries on the current page. We recommend that you set MaxResults to 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. Valid values:
	//
	// - If this is the first query or no next query is available, you do not need to specify this parameter.
	//
	// - If a next query is available, set the value to the **NextToken*	- value returned from the previous API call.
	//
	// example:
	//
	// dd20****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-0a1b2c3d4e5f6a7b8
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The IDs of the route maps to query. If you do not specify this parameter, all route maps are queried.
	RouteMapIds []*string `json:"RouteMapIds,omitempty" xml:"RouteMapIds,omitempty" type:"Repeated"`
}

func (s ListRouteMapsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRouteMapsRequest) GoString() string {
	return s.String()
}

func (s *ListRouteMapsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListRouteMapsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ListRouteMapsRequest) GetFilter() []*ListRouteMapsRequestFilter {
	return s.Filter
}

func (s *ListRouteMapsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRouteMapsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRouteMapsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListRouteMapsRequest) GetRouteMapIds() []*string {
	return s.RouteMapIds
}

func (s *ListRouteMapsRequest) SetClientToken(v string) *ListRouteMapsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListRouteMapsRequest) SetDryRun(v bool) *ListRouteMapsRequest {
	s.DryRun = &v
	return s
}

func (s *ListRouteMapsRequest) SetFilter(v []*ListRouteMapsRequestFilter) *ListRouteMapsRequest {
	s.Filter = v
	return s
}

func (s *ListRouteMapsRequest) SetMaxResults(v int32) *ListRouteMapsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRouteMapsRequest) SetNextToken(v string) *ListRouteMapsRequest {
	s.NextToken = &v
	return s
}

func (s *ListRouteMapsRequest) SetResourceGroupId(v string) *ListRouteMapsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListRouteMapsRequest) SetRouteMapIds(v []*string) *ListRouteMapsRequest {
	s.RouteMapIds = v
	return s
}

func (s *ListRouteMapsRequest) Validate() error {
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

type ListRouteMapsRequestFilter struct {
	// The filter condition. Valid values:
	//
	// - **Direction**: the direction of the route map.
	//
	// - **Name**: the name of the route map.
	//
	// - **Priority**: the priority of the route map.
	//
	// By default, multiple filter conditions are connected by an AND logical operator. A route map must meet all filter conditions to be considered a match.
	//
	// You can specify up to 5 filter conditions at a time.
	//
	// example:
	//
	// Priority
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter condition.
	//
	// Specify the filter value based on the **Key**. You can specify multiple filter values for a single **Key**. Multiple filter values are connected by an OR logical operator. A route map that meets any of the filter values is considered a match for the filter condition.
	//
	// You can specify up to 5 filter values for a single filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListRouteMapsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListRouteMapsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListRouteMapsRequestFilter) GetKey() *string {
	return s.Key
}

func (s *ListRouteMapsRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListRouteMapsRequestFilter) SetKey(v string) *ListRouteMapsRequestFilter {
	s.Key = &v
	return s
}

func (s *ListRouteMapsRequestFilter) SetValue(v []*string) *ListRouteMapsRequestFilter {
	s.Value = v
	return s
}

func (s *ListRouteMapsRequestFilter) Validate() error {
	return dara.Validate(s)
}
