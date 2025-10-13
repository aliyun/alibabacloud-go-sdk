// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWarehouseDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetWarehouseDetailResponseBody
	GetRequestId() *string
	SetWarehouseDetail(v *GetWarehouseDetailResponseBodyWarehouseDetail) *GetWarehouseDetailResponseBody
	GetWarehouseDetail() *GetWarehouseDetailResponseBodyWarehouseDetail
}

type GetWarehouseDetailResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D3AE84AB-0873-5FC7-A4C4-8CF869D2FA70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned values.
	WarehouseDetail *GetWarehouseDetailResponseBodyWarehouseDetail `json:"WarehouseDetail,omitempty" xml:"WarehouseDetail,omitempty" type:"Struct"`
}

func (s GetWarehouseDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWarehouseDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetWarehouseDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWarehouseDetailResponseBody) GetWarehouseDetail() *GetWarehouseDetailResponseBodyWarehouseDetail {
	return s.WarehouseDetail
}

func (s *GetWarehouseDetailResponseBody) SetRequestId(v string) *GetWarehouseDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWarehouseDetailResponseBody) SetWarehouseDetail(v *GetWarehouseDetailResponseBodyWarehouseDetail) *GetWarehouseDetailResponseBody {
	s.WarehouseDetail = v
	return s
}

func (s *GetWarehouseDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetWarehouseDetailResponseBodyWarehouseDetail struct {
	// The remaining unallocated computing resources of the virtual warehouse instance.
	//
	// example:
	//
	// 32
	RemainingCpu *string `json:"RemainingCpu,omitempty" xml:"RemainingCpu,omitempty"`
	// The reserved computing resources. The amount of computing resources in all running virtual warehouses in an instance cannot exceed the amount of reserved computing resources in the virtual warehouses.
	//
	// example:
	//
	// 64
	ReservedCpu     *string `json:"ReservedCpu,omitempty" xml:"ReservedCpu,omitempty"`
	TimedElasticCpu *string `json:"TimedElasticCpu,omitempty" xml:"TimedElasticCpu,omitempty"`
	// The list of virtual warehouses.
	WarehouseList []*GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList `json:"WarehouseList,omitempty" xml:"WarehouseList,omitempty" type:"Repeated"`
}

func (s GetWarehouseDetailResponseBodyWarehouseDetail) String() string {
	return dara.Prettify(s)
}

func (s GetWarehouseDetailResponseBodyWarehouseDetail) GoString() string {
	return s.String()
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetail) GetRemainingCpu() *string {
	return s.RemainingCpu
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetail) GetReservedCpu() *string {
	return s.ReservedCpu
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetail) GetTimedElasticCpu() *string {
	return s.TimedElasticCpu
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetail) GetWarehouseList() []*GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	return s.WarehouseList
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetail) SetRemainingCpu(v string) *GetWarehouseDetailResponseBodyWarehouseDetail {
	s.RemainingCpu = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetail) SetReservedCpu(v string) *GetWarehouseDetailResponseBodyWarehouseDetail {
	s.ReservedCpu = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetail) SetTimedElasticCpu(v string) *GetWarehouseDetailResponseBodyWarehouseDetail {
	s.TimedElasticCpu = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetail) SetWarehouseList(v []*GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) *GetWarehouseDetailResponseBodyWarehouseDetail {
	s.WarehouseList = v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetail) Validate() error {
	return dara.Validate(s)
}

type GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList struct {
	// The number of CPU cores.
	//
	// example:
	//
	// 32
	Cpu              *int64 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	DefaultWarehouse *bool  `json:"DefaultWarehouse,omitempty" xml:"DefaultWarehouse,omitempty"`
	ElasticCpu       *int64 `json:"ElasticCpu,omitempty" xml:"ElasticCpu,omitempty"`
	// The ID.
	//
	// example:
	//
	// 2
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
	NodeCount       *int64  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	RebalanceStatus *string `json:"RebalanceStatus,omitempty" xml:"RebalanceStatus,omitempty"`
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

func (s GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) String() string {
	return dara.Prettify(s)
}

func (s GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GoString() string {
	return s.String()
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GetCpu() *int64 {
	return s.Cpu
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GetDefaultWarehouse() *bool {
	return s.DefaultWarehouse
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GetElasticCpu() *int64 {
	return s.ElasticCpu
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GetId() *int64 {
	return s.Id
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GetMem() *int64 {
	return s.Mem
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GetName() *string {
	return s.Name
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GetNodeCount() *int64 {
	return s.NodeCount
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GetRebalanceStatus() *string {
	return s.RebalanceStatus
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GetStatus() *string {
	return s.Status
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetCpu(v int64) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.Cpu = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetDefaultWarehouse(v bool) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.DefaultWarehouse = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetElasticCpu(v int64) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.ElasticCpu = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetId(v int64) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.Id = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetMem(v int64) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.Mem = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetName(v string) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.Name = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetNodeCount(v int64) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.NodeCount = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetRebalanceStatus(v string) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.RebalanceStatus = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetStatus(v string) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.Status = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) Validate() error {
	return dara.Validate(s)
}
