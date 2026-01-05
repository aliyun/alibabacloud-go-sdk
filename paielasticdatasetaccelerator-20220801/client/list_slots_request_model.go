// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSlotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointIds(v string) *ListSlotsRequest
	GetEndpointIds() *string
	SetInstanceIds(v string) *ListSlotsRequest
	GetInstanceIds() *string
	SetName(v string) *ListSlotsRequest
	GetName() *string
	SetOrder(v string) *ListSlotsRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListSlotsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSlotsRequest
	GetPageSize() *int32
	SetPhase(v string) *ListSlotsRequest
	GetPhase() *string
	SetSlotIds(v string) *ListSlotsRequest
	GetSlotIds() *string
	SetSortBy(v string) *ListSlotsRequest
	GetSortBy() *string
	SetStorageType(v string) *ListSlotsRequest
	GetStorageType() *string
	SetStorageUri(v string) *ListSlotsRequest
	GetStorageUri() *string
}

type ListSlotsRequest struct {
	// 加速槽所对应的数据集加速挂载点的唯一标识符。
	//
	// example:
	//
	// endp-my1tk3jggooi5uwks5
	EndpointIds *string `json:"EndpointIds,omitempty" xml:"EndpointIds,omitempty"`
	// example:
	//
	// inst-my1tk3jggooi5uwks5
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// example:
	//
	// acc_instance_slot_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// Running
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// slot-my1tk3jggooi5uwks5,slot-n69468yvjz8d12as7d,slot-tga4omjrepklkg1onn
	SlotIds *string `json:"SlotIds,omitempty" xml:"SlotIds,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// OSS
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// 数据集加速槽的数据存储路径（URI）。
	StorageUri *string `json:"StorageUri,omitempty" xml:"StorageUri,omitempty"`
}

func (s ListSlotsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSlotsRequest) GoString() string {
	return s.String()
}

func (s *ListSlotsRequest) GetEndpointIds() *string {
	return s.EndpointIds
}

func (s *ListSlotsRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *ListSlotsRequest) GetName() *string {
	return s.Name
}

func (s *ListSlotsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListSlotsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSlotsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSlotsRequest) GetPhase() *string {
	return s.Phase
}

func (s *ListSlotsRequest) GetSlotIds() *string {
	return s.SlotIds
}

func (s *ListSlotsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListSlotsRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *ListSlotsRequest) GetStorageUri() *string {
	return s.StorageUri
}

func (s *ListSlotsRequest) SetEndpointIds(v string) *ListSlotsRequest {
	s.EndpointIds = &v
	return s
}

func (s *ListSlotsRequest) SetInstanceIds(v string) *ListSlotsRequest {
	s.InstanceIds = &v
	return s
}

func (s *ListSlotsRequest) SetName(v string) *ListSlotsRequest {
	s.Name = &v
	return s
}

func (s *ListSlotsRequest) SetOrder(v string) *ListSlotsRequest {
	s.Order = &v
	return s
}

func (s *ListSlotsRequest) SetPageNumber(v int32) *ListSlotsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSlotsRequest) SetPageSize(v int32) *ListSlotsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSlotsRequest) SetPhase(v string) *ListSlotsRequest {
	s.Phase = &v
	return s
}

func (s *ListSlotsRequest) SetSlotIds(v string) *ListSlotsRequest {
	s.SlotIds = &v
	return s
}

func (s *ListSlotsRequest) SetSortBy(v string) *ListSlotsRequest {
	s.SortBy = &v
	return s
}

func (s *ListSlotsRequest) SetStorageType(v string) *ListSlotsRequest {
	s.StorageType = &v
	return s
}

func (s *ListSlotsRequest) SetStorageUri(v string) *ListSlotsRequest {
	s.StorageUri = &v
	return s
}

func (s *ListSlotsRequest) Validate() error {
	return dara.Validate(s)
}
