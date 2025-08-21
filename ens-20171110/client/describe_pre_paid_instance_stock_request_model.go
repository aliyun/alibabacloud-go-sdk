// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrePaidInstanceStockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataDiskSize(v int32) *DescribePrePaidInstanceStockRequest
	GetDataDiskSize() *int32
	SetEnsRegionId(v string) *DescribePrePaidInstanceStockRequest
	GetEnsRegionId() *string
	SetInstanceSpec(v string) *DescribePrePaidInstanceStockRequest
	GetInstanceSpec() *string
	SetSystemDiskSize(v int32) *DescribePrePaidInstanceStockRequest
	GetSystemDiskSize() *int32
}

type DescribePrePaidInstanceStockRequest struct {
	// The size of the data disk. Unit: GB.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	DataDiskSize *int32 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	// The ID of the edge node.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-suzhou-telecom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The specification of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ens.sn1.stiny
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// The size of the system disk. Unit: GB.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribePrePaidInstanceStockRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePrePaidInstanceStockRequest) GoString() string {
	return s.String()
}

func (s *DescribePrePaidInstanceStockRequest) GetDataDiskSize() *int32 {
	return s.DataDiskSize
}

func (s *DescribePrePaidInstanceStockRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribePrePaidInstanceStockRequest) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *DescribePrePaidInstanceStockRequest) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *DescribePrePaidInstanceStockRequest) SetDataDiskSize(v int32) *DescribePrePaidInstanceStockRequest {
	s.DataDiskSize = &v
	return s
}

func (s *DescribePrePaidInstanceStockRequest) SetEnsRegionId(v string) *DescribePrePaidInstanceStockRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribePrePaidInstanceStockRequest) SetInstanceSpec(v string) *DescribePrePaidInstanceStockRequest {
	s.InstanceSpec = &v
	return s
}

func (s *DescribePrePaidInstanceStockRequest) SetSystemDiskSize(v int32) *DescribePrePaidInstanceStockRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribePrePaidInstanceStockRequest) Validate() error {
	return dara.Validate(s)
}
