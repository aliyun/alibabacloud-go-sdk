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
	// The pagination token. Set this parameter to the NextToken value returned from the previous API call.
	//
	// example:
	//
	// 5c0a1c0f91c14c6****
	NextToken *string `json:"next_token,omitempty" xml:"next_token,omitempty"`
	// The page number for paging in the paged query.
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
	// - `cluster_create`: Cluster creation.
	//
	// - `cluster_scaleout`: Cluster scale-out.
	//
	// - `cluster_attach`: Addition of existing nodes.
	//
	// - `cluster_delete`: Cluster deletion.
	//
	// - `cluster_upgrade`: Cluster upgrade.
	//
	// - `cluster_migrate`: Cluster migration.
	//
	// - `cluster_node_delete`: Node removal.
	//
	// - `cluster_node_drain`: Node draining.
	//
	// - `cluster_modify`: Cluster modification.
	//
	// - `cluster_configuration_modify`: Cluster control configuration modification.
	//
	// - `cluster_addon_install`: Add-on installation.
	//
	// - `cluster_addon_upgrade`: Add-on upgrade.
	//
	// - `cluster_addon_uninstall`: Add-on uninstallation.
	//
	// - `runtime_upgrade`: Runtime upgrade.
	//
	// - `nodepool_upgrade`: Node pool upgrade.
	//
	// - `nodepool_update`: Node pool update.
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
