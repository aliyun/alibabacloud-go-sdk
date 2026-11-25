// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeEventsRequest
	GetClusterId() *string
	SetMaxResults(v int32) *DescribeEventsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeEventsRequest
	GetNextToken() *string
	SetPageNumber(v int64) *DescribeEventsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeEventsRequest
	GetPageSize() *int64
	SetType(v string) *DescribeEventsRequest
	GetType() *string
}

type DescribeEventsRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// cf62854ac2130470897be7a27ed1f****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The maximum number of results to return for a single request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"max_results,omitempty" xml:"max_results,omitempty"`
	// The query token, which is the value of the NextToken parameter returned by the previous API call.
	//
	// example:
	//
	// 5c0a1c0f91c14c6****
	NextToken *string `json:"next_token,omitempty" xml:"next_token,omitempty"`
	// The page number for paging.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The maximum number of results to return per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 50.
	//
	// example:
	//
	// 50
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// The event type. If you do not specify this parameter, events of all types are queried. Valid values:
	//
	// - `cluster_create`: cluster creation
	//
	// - `cluster_scaleout`: cluster scale-out
	//
	// - `cluster_attach`: adding existing nodes
	//
	// - `cluster_delete`: cluster deletion
	//
	// - `cluster_upgrade`: cluster upgrade
	//
	// - `cluster_migrate`: cluster migration
	//
	// - `cluster_node_delete`: node removal
	//
	// - `cluster_node_drain`: node draining
	//
	// - `cluster_modify`: cluster modification
	//
	// - `cluster_configuration_modify`: cluster management configuration modification
	//
	// - `cluster_addon_install`: add-on installation
	//
	// - `cluster_addon_upgrade`: add-on upgrade
	//
	// - `cluster_addon_uninstall`: add-on uninstallation
	//
	// - `runtime_upgrade`: runtime upgrade
	//
	// - `nodepool_upgrade`: node pool upgrade
	//
	// - `nodepool_update`: node pool update
	//
	// example:
	//
	// nodepool_upgrade
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeEventsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeEventsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeEventsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeEventsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeEventsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeEventsRequest) SetClusterId(v string) *DescribeEventsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeEventsRequest) SetMaxResults(v int32) *DescribeEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeEventsRequest) SetNextToken(v string) *DescribeEventsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeEventsRequest) SetPageNumber(v int64) *DescribeEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEventsRequest) SetPageSize(v int64) *DescribeEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEventsRequest) SetType(v string) *DescribeEventsRequest {
	s.Type = &v
	return s
}

func (s *DescribeEventsRequest) Validate() error {
	return dara.Validate(s)
}
