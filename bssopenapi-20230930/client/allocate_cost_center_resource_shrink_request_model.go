// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateCostCenterResourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromCostCenterId(v int64) *AllocateCostCenterResourceShrinkRequest
	GetFromCostCenterId() *int64
	SetFromOwnerAccountId(v int64) *AllocateCostCenterResourceShrinkRequest
	GetFromOwnerAccountId() *int64
	SetNbid(v string) *AllocateCostCenterResourceShrinkRequest
	GetNbid() *string
	SetResourceInstanceListShrink(v string) *AllocateCostCenterResourceShrinkRequest
	GetResourceInstanceListShrink() *string
	SetToCostCenterId(v int64) *AllocateCostCenterResourceShrinkRequest
	GetToCostCenterId() *int64
}

type AllocateCostCenterResourceShrinkRequest struct {
	// The ID of the source cost center. This parameter is required.
	//
	// - 0 indicates that the cost center is unallocated.
	//
	// - A value greater than 0 indicates an allocated cost center ID.
	//
	// example:
	//
	// 637180
	FromCostCenterId *int64 `json:"FromCostCenterId,omitempty" xml:"FromCostCenterId,omitempty"`
	// The ID of the owner of the source cost center.
	//
	// example:
	//
	// 1529600453335198
	FromOwnerAccountId *int64 `json:"FromOwnerAccountId,omitempty" xml:"FromOwnerAccountId,omitempty"`
	// The primary sales channel ID. If this parameter is left empty, the sales channel ID of the current user is used by default.
	//
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// The list of resource instances.
	//
	// This parameter is required.
	ResourceInstanceListShrink *string `json:"ResourceInstanceList,omitempty" xml:"ResourceInstanceList,omitempty"`
	// The ID of the destination cost center. Valid values:
	//
	// - -1: moves the allocated resource to the unallocated state.
	//
	// - A value greater than 0: allocates the resource to the specified cost center.
	//
	// example:
	//
	// 638288
	ToCostCenterId *int64 `json:"ToCostCenterId,omitempty" xml:"ToCostCenterId,omitempty"`
}

func (s AllocateCostCenterResourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateCostCenterResourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *AllocateCostCenterResourceShrinkRequest) GetFromCostCenterId() *int64 {
	return s.FromCostCenterId
}

func (s *AllocateCostCenterResourceShrinkRequest) GetFromOwnerAccountId() *int64 {
	return s.FromOwnerAccountId
}

func (s *AllocateCostCenterResourceShrinkRequest) GetNbid() *string {
	return s.Nbid
}

func (s *AllocateCostCenterResourceShrinkRequest) GetResourceInstanceListShrink() *string {
	return s.ResourceInstanceListShrink
}

func (s *AllocateCostCenterResourceShrinkRequest) GetToCostCenterId() *int64 {
	return s.ToCostCenterId
}

func (s *AllocateCostCenterResourceShrinkRequest) SetFromCostCenterId(v int64) *AllocateCostCenterResourceShrinkRequest {
	s.FromCostCenterId = &v
	return s
}

func (s *AllocateCostCenterResourceShrinkRequest) SetFromOwnerAccountId(v int64) *AllocateCostCenterResourceShrinkRequest {
	s.FromOwnerAccountId = &v
	return s
}

func (s *AllocateCostCenterResourceShrinkRequest) SetNbid(v string) *AllocateCostCenterResourceShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *AllocateCostCenterResourceShrinkRequest) SetResourceInstanceListShrink(v string) *AllocateCostCenterResourceShrinkRequest {
	s.ResourceInstanceListShrink = &v
	return s
}

func (s *AllocateCostCenterResourceShrinkRequest) SetToCostCenterId(v int64) *AllocateCostCenterResourceShrinkRequest {
	s.ToCostCenterId = &v
	return s
}

func (s *AllocateCostCenterResourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
