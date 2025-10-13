// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWarehousesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetWarehouseList(v []*ListWarehousesResponseBodyWarehouseList) *ListWarehousesResponseBody
	GetWarehouseList() []*ListWarehousesResponseBodyWarehouseList
	SetRequestId(v string) *ListWarehousesResponseBody
	GetRequestId() *string
}

type ListWarehousesResponseBody struct {
	// The list of virtual warehouse instances.
	WarehouseList []*ListWarehousesResponseBodyWarehouseList `json:"WarehouseList,omitempty" xml:"WarehouseList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 819A7F0F-2951-540F-BD94-6A41ECF0281F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListWarehousesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWarehousesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWarehousesResponseBody) GetWarehouseList() []*ListWarehousesResponseBodyWarehouseList {
	return s.WarehouseList
}

func (s *ListWarehousesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWarehousesResponseBody) SetWarehouseList(v []*ListWarehousesResponseBodyWarehouseList) *ListWarehousesResponseBody {
	s.WarehouseList = v
	return s
}

func (s *ListWarehousesResponseBody) SetRequestId(v string) *ListWarehousesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWarehousesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListWarehousesResponseBodyWarehouseList struct {
	// The number of CPU cores.
	//
	// example:
	//
	// 32
	Cpu *int64 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The ID.
	//
	// example:
	//
	// 3
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The memory capacity.
	//
	// example:
	//
	// 128
	Mem *int64 `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// The name of the virtual warehouse instance.
	//
	// example:
	//
	// MyWarehouse
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of compute nodes.
	//
	// example:
	//
	// 2
	NodeCount *int64 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The status.
	//
	// Valid values:
	//
	// 	- kRunning
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- kSuspended
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- kInit
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- kFailed
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- kAllocating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// kRunning
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListWarehousesResponseBodyWarehouseList) String() string {
	return dara.Prettify(s)
}

func (s ListWarehousesResponseBodyWarehouseList) GoString() string {
	return s.String()
}

func (s *ListWarehousesResponseBodyWarehouseList) GetCpu() *int64 {
	return s.Cpu
}

func (s *ListWarehousesResponseBodyWarehouseList) GetId() *int64 {
	return s.Id
}

func (s *ListWarehousesResponseBodyWarehouseList) GetMem() *int64 {
	return s.Mem
}

func (s *ListWarehousesResponseBodyWarehouseList) GetName() *string {
	return s.Name
}

func (s *ListWarehousesResponseBodyWarehouseList) GetNodeCount() *int64 {
	return s.NodeCount
}

func (s *ListWarehousesResponseBodyWarehouseList) GetStatus() *string {
	return s.Status
}

func (s *ListWarehousesResponseBodyWarehouseList) SetCpu(v int64) *ListWarehousesResponseBodyWarehouseList {
	s.Cpu = &v
	return s
}

func (s *ListWarehousesResponseBodyWarehouseList) SetId(v int64) *ListWarehousesResponseBodyWarehouseList {
	s.Id = &v
	return s
}

func (s *ListWarehousesResponseBodyWarehouseList) SetMem(v int64) *ListWarehousesResponseBodyWarehouseList {
	s.Mem = &v
	return s
}

func (s *ListWarehousesResponseBodyWarehouseList) SetName(v string) *ListWarehousesResponseBodyWarehouseList {
	s.Name = &v
	return s
}

func (s *ListWarehousesResponseBodyWarehouseList) SetNodeCount(v int64) *ListWarehousesResponseBodyWarehouseList {
	s.NodeCount = &v
	return s
}

func (s *ListWarehousesResponseBodyWarehouseList) SetStatus(v string) *ListWarehousesResponseBodyWarehouseList {
	s.Status = &v
	return s
}

func (s *ListWarehousesResponseBodyWarehouseList) Validate() error {
	return dara.Validate(s)
}
