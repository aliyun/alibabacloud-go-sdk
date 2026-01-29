// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateCostUnitResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromUnitId(v int64) *AllocateCostUnitResourceRequest
	GetFromUnitId() *int64
	SetFromUnitUserId(v int64) *AllocateCostUnitResourceRequest
	GetFromUnitUserId() *int64
	SetResourceInstanceList(v []*AllocateCostUnitResourceRequestResourceInstanceList) *AllocateCostUnitResourceRequest
	GetResourceInstanceList() []*AllocateCostUnitResourceRequestResourceInstanceList
	SetToUnitId(v int64) *AllocateCostUnitResourceRequest
	GetToUnitId() *int64
	SetToUnitUserId(v int64) *AllocateCostUnitResourceRequest
	GetToUnitUserId() *int64
}

type AllocateCostUnitResourceRequest struct {
	// The ID of the source cost center.
	//
	// 	- A value of 0 indicates that the resources to be transferred have not been allocated to a cost center.
	//
	// 	- A value greater than 0 indicates the ID of an existing cost center.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	FromUnitId *int64 `json:"FromUnitId,omitempty" xml:"FromUnitId,omitempty"`
	// The user ID of the owner of the source cost center.
	//
	// This parameter is required.
	//
	// example:
	//
	// 273394581313325532
	FromUnitUserId *int64 `json:"FromUnitUserId,omitempty" xml:"FromUnitUserId,omitempty"`
	// The resource instances to be transferred.
	//
	// This parameter is required.
	ResourceInstanceList []*AllocateCostUnitResourceRequestResourceInstanceList `json:"ResourceInstanceList,omitempty" xml:"ResourceInstanceList,omitempty" type:"Repeated"`
	// The ID of the destination cost center.
	//
	// 	- A value of -1 indicates that the allocated resources are changed to unallocated.
	//
	// 	- A value greater than 0 indicates the ID of an existing cost center.
	//
	// This parameter is required.
	//
	// example:
	//
	// 186419
	ToUnitId *int64 `json:"ToUnitId,omitempty" xml:"ToUnitId,omitempty"`
	// The user ID of the owner of the destination cost center.
	//
	// This parameter is required.
	//
	// example:
	//
	// 241021678450941490
	ToUnitUserId *int64 `json:"ToUnitUserId,omitempty" xml:"ToUnitUserId,omitempty"`
}

func (s AllocateCostUnitResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateCostUnitResourceRequest) GoString() string {
	return s.String()
}

func (s *AllocateCostUnitResourceRequest) GetFromUnitId() *int64 {
	return s.FromUnitId
}

func (s *AllocateCostUnitResourceRequest) GetFromUnitUserId() *int64 {
	return s.FromUnitUserId
}

func (s *AllocateCostUnitResourceRequest) GetResourceInstanceList() []*AllocateCostUnitResourceRequestResourceInstanceList {
	return s.ResourceInstanceList
}

func (s *AllocateCostUnitResourceRequest) GetToUnitId() *int64 {
	return s.ToUnitId
}

func (s *AllocateCostUnitResourceRequest) GetToUnitUserId() *int64 {
	return s.ToUnitUserId
}

func (s *AllocateCostUnitResourceRequest) SetFromUnitId(v int64) *AllocateCostUnitResourceRequest {
	s.FromUnitId = &v
	return s
}

func (s *AllocateCostUnitResourceRequest) SetFromUnitUserId(v int64) *AllocateCostUnitResourceRequest {
	s.FromUnitUserId = &v
	return s
}

func (s *AllocateCostUnitResourceRequest) SetResourceInstanceList(v []*AllocateCostUnitResourceRequestResourceInstanceList) *AllocateCostUnitResourceRequest {
	s.ResourceInstanceList = v
	return s
}

func (s *AllocateCostUnitResourceRequest) SetToUnitId(v int64) *AllocateCostUnitResourceRequest {
	s.ToUnitId = &v
	return s
}

func (s *AllocateCostUnitResourceRequest) SetToUnitUserId(v int64) *AllocateCostUnitResourceRequest {
	s.ToUnitUserId = &v
	return s
}

func (s *AllocateCostUnitResourceRequest) Validate() error {
	if s.ResourceInstanceList != nil {
		for _, item := range s.ResourceInstanceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AllocateCostUnitResourceRequestResourceInstanceList struct {
	// The split item of the shared instance. This parameter is required only for shared instances.
	//
	// 	- Eight cloud services support bill splitting. The commodity codes of the eight services are oss, dcdn, snapshot, vod, cdn, live, and cbwp.
	//
	// 	- You can obtain the split item of a shared instance by calling QueryCostUnitResource operation to obtain all resource instances within a cost center.
	//
	// example:
	//
	// qwer1-cn-beijing
	ApportionCode *string `json:"ApportionCode,omitempty" xml:"ApportionCode,omitempty"`
	// The commodity code of the resource instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The ID of the resource instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou;standard
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The user ID of the resource instance owner.
	//
	// This parameter is required.
	//
	// example:
	//
	// 273394581313325532
	ResourceUserId *int64 `json:"ResourceUserId,omitempty" xml:"ResourceUserId,omitempty"`
}

func (s AllocateCostUnitResourceRequestResourceInstanceList) String() string {
	return dara.Prettify(s)
}

func (s AllocateCostUnitResourceRequestResourceInstanceList) GoString() string {
	return s.String()
}

func (s *AllocateCostUnitResourceRequestResourceInstanceList) GetApportionCode() *string {
	return s.ApportionCode
}

func (s *AllocateCostUnitResourceRequestResourceInstanceList) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *AllocateCostUnitResourceRequestResourceInstanceList) GetResourceId() *string {
	return s.ResourceId
}

func (s *AllocateCostUnitResourceRequestResourceInstanceList) GetResourceUserId() *int64 {
	return s.ResourceUserId
}

func (s *AllocateCostUnitResourceRequestResourceInstanceList) SetApportionCode(v string) *AllocateCostUnitResourceRequestResourceInstanceList {
	s.ApportionCode = &v
	return s
}

func (s *AllocateCostUnitResourceRequestResourceInstanceList) SetCommodityCode(v string) *AllocateCostUnitResourceRequestResourceInstanceList {
	s.CommodityCode = &v
	return s
}

func (s *AllocateCostUnitResourceRequestResourceInstanceList) SetResourceId(v string) *AllocateCostUnitResourceRequestResourceInstanceList {
	s.ResourceId = &v
	return s
}

func (s *AllocateCostUnitResourceRequestResourceInstanceList) SetResourceUserId(v int64) *AllocateCostUnitResourceRequestResourceInstanceList {
	s.ResourceUserId = &v
	return s
}

func (s *AllocateCostUnitResourceRequestResourceInstanceList) Validate() error {
	return dara.Validate(s)
}
