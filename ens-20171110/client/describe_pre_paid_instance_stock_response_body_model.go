// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrePaidInstanceStockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvaliableCount(v int32) *DescribePrePaidInstanceStockResponseBody
	GetAvaliableCount() *int32
	SetCores(v int32) *DescribePrePaidInstanceStockResponseBody
	GetCores() *int32
	SetDataDiskSize(v int32) *DescribePrePaidInstanceStockResponseBody
	GetDataDiskSize() *int32
	SetEnsRegionId(v string) *DescribePrePaidInstanceStockResponseBody
	GetEnsRegionId() *string
	SetInstanceSpec(v string) *DescribePrePaidInstanceStockResponseBody
	GetInstanceSpec() *string
	SetMemory(v int32) *DescribePrePaidInstanceStockResponseBody
	GetMemory() *int32
	SetRequestId(v string) *DescribePrePaidInstanceStockResponseBody
	GetRequestId() *string
	SetResourceGap(v string) *DescribePrePaidInstanceStockResponseBody
	GetResourceGap() *string
	SetSystemDiskSize(v int32) *DescribePrePaidInstanceStockResponseBody
	GetSystemDiskSize() *int32
}

type DescribePrePaidInstanceStockResponseBody struct {
	// The number of resources that you can purchase.
	//
	// example:
	//
	// 84
	AvaliableCount *int32 `json:"AvaliableCount,omitempty" xml:"AvaliableCount,omitempty"`
	// The number of CPU cores.
	//
	// example:
	//
	// 1
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The size of the data disk.
	//
	// example:
	//
	// 20
	DataDiskSize *int32 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// cn-suzhou-telecom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The specification of the instance.
	//
	// example:
	//
	// ens.sn1.stiny
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// The memory size. Unit: GB.
	//
	// example:
	//
	// 2048
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 007833C8-E819-4122-B636-0D48D7BF6DFB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The reason why resources are insufficient.
	//
	// example:
	//
	// StockNotEnough
	ResourceGap *string `json:"ResourceGap,omitempty" xml:"ResourceGap,omitempty"`
	// The size of the system disk.
	//
	// example:
	//
	// 20
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribePrePaidInstanceStockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePrePaidInstanceStockResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePrePaidInstanceStockResponseBody) GetAvaliableCount() *int32 {
	return s.AvaliableCount
}

func (s *DescribePrePaidInstanceStockResponseBody) GetCores() *int32 {
	return s.Cores
}

func (s *DescribePrePaidInstanceStockResponseBody) GetDataDiskSize() *int32 {
	return s.DataDiskSize
}

func (s *DescribePrePaidInstanceStockResponseBody) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribePrePaidInstanceStockResponseBody) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *DescribePrePaidInstanceStockResponseBody) GetMemory() *int32 {
	return s.Memory
}

func (s *DescribePrePaidInstanceStockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePrePaidInstanceStockResponseBody) GetResourceGap() *string {
	return s.ResourceGap
}

func (s *DescribePrePaidInstanceStockResponseBody) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *DescribePrePaidInstanceStockResponseBody) SetAvaliableCount(v int32) *DescribePrePaidInstanceStockResponseBody {
	s.AvaliableCount = &v
	return s
}

func (s *DescribePrePaidInstanceStockResponseBody) SetCores(v int32) *DescribePrePaidInstanceStockResponseBody {
	s.Cores = &v
	return s
}

func (s *DescribePrePaidInstanceStockResponseBody) SetDataDiskSize(v int32) *DescribePrePaidInstanceStockResponseBody {
	s.DataDiskSize = &v
	return s
}

func (s *DescribePrePaidInstanceStockResponseBody) SetEnsRegionId(v string) *DescribePrePaidInstanceStockResponseBody {
	s.EnsRegionId = &v
	return s
}

func (s *DescribePrePaidInstanceStockResponseBody) SetInstanceSpec(v string) *DescribePrePaidInstanceStockResponseBody {
	s.InstanceSpec = &v
	return s
}

func (s *DescribePrePaidInstanceStockResponseBody) SetMemory(v int32) *DescribePrePaidInstanceStockResponseBody {
	s.Memory = &v
	return s
}

func (s *DescribePrePaidInstanceStockResponseBody) SetRequestId(v string) *DescribePrePaidInstanceStockResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePrePaidInstanceStockResponseBody) SetResourceGap(v string) *DescribePrePaidInstanceStockResponseBody {
	s.ResourceGap = &v
	return s
}

func (s *DescribePrePaidInstanceStockResponseBody) SetSystemDiskSize(v int32) *DescribePrePaidInstanceStockResponseBody {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribePrePaidInstanceStockResponseBody) Validate() error {
	return dara.Validate(s)
}
