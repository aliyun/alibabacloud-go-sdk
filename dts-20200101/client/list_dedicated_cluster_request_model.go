// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDedicatedClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderColumn(v string) *ListDedicatedClusterRequest
	GetOrderColumn() *string
	SetOrderDirection(v string) *ListDedicatedClusterRequest
	GetOrderDirection() *string
	SetOwnerId(v string) *ListDedicatedClusterRequest
	GetOwnerId() *string
	SetPageNumber(v int32) *ListDedicatedClusterRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDedicatedClusterRequest
	GetPageSize() *int32
	SetParams(v string) *ListDedicatedClusterRequest
	GetParams() *string
	SetRegionId(v string) *ListDedicatedClusterRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListDedicatedClusterRequest
	GetResourceGroupId() *string
	SetState(v string) *ListDedicatedClusterRequest
	GetState() *string
	SetType(v string) *ListDedicatedClusterRequest
	GetType() *string
}

type ListDedicatedClusterRequest struct {
	// The basis on which the retrieved entries are sorted if multiple DTS dedicated clusters are returned. Valid values:
	//
	// 	- **gmtCreated**: the time when a cluster was created.
	//
	// 	- **orderCount**: the number of nodes in a cluster.
	//
	// example:
	//
	// gmtCreated
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// The order in which you want to sort the retrieved entries. Valid values:
	//
	// 	- asc: sorts the retrieved entries in ascending order. This is the default value.
	//
	// 	- desc: sorts the retrieved entries in descending order.
	//
	// example:
	//
	// asc
	OrderDirection *string `json:"OrderDirection,omitempty" xml:"OrderDirection,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. The value of this parameter must be an integer that is greater than 0. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of clusters to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The content of the query condition.
	//
	// >  You must set the **Type parameter*	- to specify the type of the query condition.
	//
	// example:
	//
	// dtspk3f13r731m****
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfntftbiobqyky
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the cluster. Valid values:
	//
	// 	- **init**: The cluster is being initialized.
	//
	// 	- **schedule**: The cluster is pending scheduling.
	//
	// 	- **running**: The cluster is running.
	//
	// 	- **upgrade**: The cluster is being upgraded.
	//
	// 	- **downgrade**: The cluster is being downgraded.
	//
	// 	- **locked**: The cluster is locked.
	//
	// 	- **releasing**: The cluster is being released.
	//
	// 	- **released**: The cluster is released.
	//
	// example:
	//
	// init
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The type of the query condition. Valid values:
	//
	// 	- **NAME**: the name of the cluster.
	//
	// 	- **INSTANCE**: the ID of a cluster instance.
	//
	// 	- **DEDICAETEDCLUSTERID**: the ID of a dedicated cluster.
	//
	// >  You must specify the query condition by using the **Params*	- parameter.
	//
	// example:
	//
	// NAME
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDedicatedClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDedicatedClusterRequest) GoString() string {
	return s.String()
}

func (s *ListDedicatedClusterRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *ListDedicatedClusterRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListDedicatedClusterRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListDedicatedClusterRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDedicatedClusterRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDedicatedClusterRequest) GetParams() *string {
	return s.Params
}

func (s *ListDedicatedClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDedicatedClusterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListDedicatedClusterRequest) GetState() *string {
	return s.State
}

func (s *ListDedicatedClusterRequest) GetType() *string {
	return s.Type
}

func (s *ListDedicatedClusterRequest) SetOrderColumn(v string) *ListDedicatedClusterRequest {
	s.OrderColumn = &v
	return s
}

func (s *ListDedicatedClusterRequest) SetOrderDirection(v string) *ListDedicatedClusterRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListDedicatedClusterRequest) SetOwnerId(v string) *ListDedicatedClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *ListDedicatedClusterRequest) SetPageNumber(v int32) *ListDedicatedClusterRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDedicatedClusterRequest) SetPageSize(v int32) *ListDedicatedClusterRequest {
	s.PageSize = &v
	return s
}

func (s *ListDedicatedClusterRequest) SetParams(v string) *ListDedicatedClusterRequest {
	s.Params = &v
	return s
}

func (s *ListDedicatedClusterRequest) SetRegionId(v string) *ListDedicatedClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ListDedicatedClusterRequest) SetResourceGroupId(v string) *ListDedicatedClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListDedicatedClusterRequest) SetState(v string) *ListDedicatedClusterRequest {
	s.State = &v
	return s
}

func (s *ListDedicatedClusterRequest) SetType(v string) *ListDedicatedClusterRequest {
	s.Type = &v
	return s
}

func (s *ListDedicatedClusterRequest) Validate() error {
	return dara.Validate(s)
}
