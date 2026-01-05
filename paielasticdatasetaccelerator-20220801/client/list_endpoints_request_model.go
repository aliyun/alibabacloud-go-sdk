// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointIds(v string) *ListEndpointsRequest
	GetEndpointIds() *string
	SetInstanceIds(v string) *ListEndpointsRequest
	GetInstanceIds() *string
	SetName(v string) *ListEndpointsRequest
	GetName() *string
	SetOrder(v string) *ListEndpointsRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListEndpointsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEndpointsRequest
	GetPageSize() *int32
	SetSlotIds(v string) *ListEndpointsRequest
	GetSlotIds() *string
	SetSortBy(v string) *ListEndpointsRequest
	GetSortBy() *string
	SetType(v string) *ListEndpointsRequest
	GetType() *string
}

type ListEndpointsRequest struct {
	// example:
	//
	// end-my1tk3jggooi5uwks5,end-n69468yvjz8d12as7d,end-tga4omjrepklkg1onn
	EndpointIds *string `json:"EndpointIds,omitempty" xml:"EndpointIds,omitempty"`
	// 所属加速实例的ID。
	//
	// example:
	//
	// inst-ivrq92qhbyrg4jctih
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// example:
	//
	// acc_instance_slot_mount_1
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
	// slot-my1tk3jggooi5uwks5
	SlotIds *string `json:"SlotIds,omitempty" xml:"SlotIds,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// VPC
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEndpointsRequest) GoString() string {
	return s.String()
}

func (s *ListEndpointsRequest) GetEndpointIds() *string {
	return s.EndpointIds
}

func (s *ListEndpointsRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *ListEndpointsRequest) GetName() *string {
	return s.Name
}

func (s *ListEndpointsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListEndpointsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEndpointsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEndpointsRequest) GetSlotIds() *string {
	return s.SlotIds
}

func (s *ListEndpointsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListEndpointsRequest) GetType() *string {
	return s.Type
}

func (s *ListEndpointsRequest) SetEndpointIds(v string) *ListEndpointsRequest {
	s.EndpointIds = &v
	return s
}

func (s *ListEndpointsRequest) SetInstanceIds(v string) *ListEndpointsRequest {
	s.InstanceIds = &v
	return s
}

func (s *ListEndpointsRequest) SetName(v string) *ListEndpointsRequest {
	s.Name = &v
	return s
}

func (s *ListEndpointsRequest) SetOrder(v string) *ListEndpointsRequest {
	s.Order = &v
	return s
}

func (s *ListEndpointsRequest) SetPageNumber(v int32) *ListEndpointsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEndpointsRequest) SetPageSize(v int32) *ListEndpointsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEndpointsRequest) SetSlotIds(v string) *ListEndpointsRequest {
	s.SlotIds = &v
	return s
}

func (s *ListEndpointsRequest) SetSortBy(v string) *ListEndpointsRequest {
	s.SortBy = &v
	return s
}

func (s *ListEndpointsRequest) SetType(v string) *ListEndpointsRequest {
	s.Type = &v
	return s
}

func (s *ListEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
