// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColdStorageSize(v int64) *ScaleInstanceRequest
	GetColdStorageSize() *int64
	SetCpu(v int64) *ScaleInstanceRequest
	GetCpu() *int64
	SetEnableServerlessComputing(v bool) *ScaleInstanceRequest
	GetEnableServerlessComputing() *bool
	SetGatewayCount(v int64) *ScaleInstanceRequest
	GetGatewayCount() *int64
	SetScaleType(v string) *ScaleInstanceRequest
	GetScaleType() *string
	SetStorageSize(v int64) *ScaleInstanceRequest
	GetStorageSize() *int64
}

type ScaleInstanceRequest struct {
	// The infrequent access storage space of the instance. Unit: GB.
	//
	// > This parameter is ignored for pay-as-you-go instances.
	//
	// example:
	//
	// 1000
	ColdStorageSize *int64 `json:"coldStorageSize,omitempty" xml:"coldStorageSize,omitempty"`
	// The instance specifications. Valid values:
	//
	// - 16 cores, 64 GB (Number of compute nodes: 1)
	//
	// - 32 cores, 128 GB (Number of compute nodes: 2)
	//
	// - 64 cores, 256 GB (Number of compute nodes: 4)
	//
	// - 96 cores, 384 GB (Number of compute nodes: 6)
	//
	// - 128 cores, 512 GB (Number of compute nodes: 8)
	//
	// - More
	//
	// >
	//
	//  >- Specify only the number of cores.
	//
	//  >- To purchase specifications with more than 1,024 cores, submit a ticket.
	//
	//  >- You do not need to specify the specifications for shared instances.
	//
	// example:
	//
	// 128
	Cpu *int64 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// Specifies whether to enable serverless computing.
	//
	// example:
	//
	// true
	EnableServerlessComputing *bool `json:"enableServerlessComputing,omitempty" xml:"enableServerlessComputing,omitempty"`
	// The number of gateways. Valid values: 2 to 50.
	//
	// > This parameter is required only for the Warehouse type.
	//
	// example:
	//
	// 4
	GatewayCount *int64 `json:"gatewayCount,omitempty" xml:"gatewayCount,omitempty"`
	// The type of the specification change. Valid values:
	//
	// - UPGRADE: upgrades the specifications.
	//
	// - DOWNGRADE: downgrades the specifications.
	//
	// >
	//
	// > - When you upgrade the specifications, the new specifications cannot be lower than the original specifications. An empty field indicates that the original specifications remain unchanged. At least one specification must be higher than the original specification.
	//
	// > - When you downgrade the specifications, the new specifications cannot be higher than the original specifications. An empty field indicates that the original specifications remain unchanged. At least one specification must be lower than the original specification.
	//
	// This parameter is required.
	//
	// example:
	//
	// UPGRADE
	ScaleType *string `json:"scaleType,omitempty" xml:"scaleType,omitempty"`
	// The standard storage space of the instance. Unit: GB.
	//
	// > This parameter is ignored for pay-as-you-go instances.
	//
	// example:
	//
	// 1000
	StorageSize *int64 `json:"storageSize,omitempty" xml:"storageSize,omitempty"`
}

func (s ScaleInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ScaleInstanceRequest) GoString() string {
	return s.String()
}

func (s *ScaleInstanceRequest) GetColdStorageSize() *int64 {
	return s.ColdStorageSize
}

func (s *ScaleInstanceRequest) GetCpu() *int64 {
	return s.Cpu
}

func (s *ScaleInstanceRequest) GetEnableServerlessComputing() *bool {
	return s.EnableServerlessComputing
}

func (s *ScaleInstanceRequest) GetGatewayCount() *int64 {
	return s.GatewayCount
}

func (s *ScaleInstanceRequest) GetScaleType() *string {
	return s.ScaleType
}

func (s *ScaleInstanceRequest) GetStorageSize() *int64 {
	return s.StorageSize
}

func (s *ScaleInstanceRequest) SetColdStorageSize(v int64) *ScaleInstanceRequest {
	s.ColdStorageSize = &v
	return s
}

func (s *ScaleInstanceRequest) SetCpu(v int64) *ScaleInstanceRequest {
	s.Cpu = &v
	return s
}

func (s *ScaleInstanceRequest) SetEnableServerlessComputing(v bool) *ScaleInstanceRequest {
	s.EnableServerlessComputing = &v
	return s
}

func (s *ScaleInstanceRequest) SetGatewayCount(v int64) *ScaleInstanceRequest {
	s.GatewayCount = &v
	return s
}

func (s *ScaleInstanceRequest) SetScaleType(v string) *ScaleInstanceRequest {
	s.ScaleType = &v
	return s
}

func (s *ScaleInstanceRequest) SetStorageSize(v int64) *ScaleInstanceRequest {
	s.StorageSize = &v
	return s
}

func (s *ScaleInstanceRequest) Validate() error {
	return dara.Validate(s)
}
