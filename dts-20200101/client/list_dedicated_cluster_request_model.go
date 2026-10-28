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
	// The sort field when the response contains multiple DTS dedicated cluster instances. Valid values:
	//
	// - **gmtCreated**: Creation time.
	//
	// - **orderCount**: Number of nodes.
	//
	// example:
	//
	// gmtCreated
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// The sort order. Valid values:
	//
	// - **asc**: Ascending order. This is the default value.
	//
	// - **desc**: Descending order.
	//
	// example:
	//
	// asc
	OrderDirection *string `json:"OrderDirection,omitempty" xml:"OrderDirection,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be an integer greater than 0 and less than or equal to the maximum value of the Integer data type. Default value: **1**.
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
	// The specific content of the query condition.
	//
	// > You must specify **Type*	- to define the query key.
	//
	// example:
	//
	// dtspk3f13r731m****
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The region ID.
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
	// The cluster status. Valid values:
	//
	// - **init**: Initializing.
	//
	// - **schedule**: Pending scheduling.
	//
	// - **running**: Running.
	//
	// - **upgrade**: Upgrading.
	//
	// - **downgrade**: Downgrading.
	//
	// - **locked**: Locked.
	//
	// - **releasing**: Releasing.
	//
	// - **released**: Released.
	//
	// example:
	//
	// init
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The query key. Valid values:
	//
	// - **NAME**: The cluster name.
	//
	// - **INSTANCE**: The cluster instance ID.
	//
	// - **DEDICATEDCLUSTERID**: The dedicated cluster ID.
	//
	// > You must specify **Params*	- to define the query condition.
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
