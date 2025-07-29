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
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// The event type. Valid values:
	//
	// 	- `cluster_create`: cluster creation.
	//
	// 	- `cluster_scaleout`: cluster scale-out.
	//
	// 	- `cluster_attach`: node addition.
	//
	// 	- `cluster_delete`: cluster deletion.
	//
	// 	- `cluster_upgrade`: cluster upgrades.
	//
	// 	- `cluster_migrate`: cluster migration.
	//
	// 	- `cluster_node_delete`: node removal.
	//
	// 	- `cluster_node_drain`: node draining.
	//
	// 	- `cluster_modify`: cluster modifications.
	//
	// 	- `cluster_configuration_modify`: modifications of control plane configurations.
	//
	// 	- `cluster_addon_install`: component installation.
	//
	// 	- `cluster_addon_upgrade`: component updates.
	//
	// 	- `cluster_addon_uninstall`: component uninstallation.
	//
	// 	- `runtime_upgrade`: runtime updates.
	//
	// 	- `nodepool_upgrade`: node pool upgrades.
	//
	// 	- `nodepool_update`: node pool updates.
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
