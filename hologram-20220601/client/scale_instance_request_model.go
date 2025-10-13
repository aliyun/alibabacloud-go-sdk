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
	// The infrequent access (IA) storage space of the instance. Unit: GB.
	//
	// > Ignore this parameter for pay-as-you-go instances.
	//
	// example:
	//
	// 1000G
	ColdStorageSize *int64 `json:"coldStorageSize,omitempty" xml:"coldStorageSize,omitempty"`
	// The specifications of the instance. Valid values:
	//
	// 	- 8-core 32GB (number of compute nodes: 1)
	//
	// 	- 16-core 64GB (number of compute nodes: 1)
	//
	// 	- 32-core 128GB (number of compute nodes: 2)
	//
	// 	- 64-core 256GB (number of compute nodes: 4)
	//
	// 	- 96-core 384GB (number of compute nodes: 6)
	//
	// 	- 128-core 512GB (number of compute nodes: 8)
	//
	// 	- Others
	//
	// >
	//
	// 	- Set this parameter to the number of cores.
	//
	// 	- If you want to set this parameter to specifications with more than 1,024 compute units (CUs), you must submit a ticket.
	//
	// 	- This parameter is invalid for Hologres Shared Cluster instances.
	//
	// 	- The specifications of 8-core 32GB (number of compute nodes: 1) are for trial use only and cannot be used for production.
	//
	// example:
	//
	// 128
	Cpu *int64 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// 是否开启ServerlessComputing
	//
	// example:
	//
	// true
	EnableServerlessComputing *bool `json:"enableServerlessComputing,omitempty" xml:"enableServerlessComputing,omitempty"`
	// The number of gateways. Valid values: 2 to 50.
	//
	// > This parameter is required only for virtual warehouse instances.
	//
	// example:
	//
	// 4
	GatewayCount *int64 `json:"gatewayCount,omitempty" xml:"gatewayCount,omitempty"`
	// The specification change type. Valid values:
	//
	// 	- UPGRADE
	//
	// 	- DOWNGRADE
	//
	// >
	//
	// 	- If you set this parameter to UPGRADE, the new specifications must be higher than the original specifications. You must configure at least one of the cpu, storageSize, and coldStorageSize parameters. If you leave a parameter empty, the related configuration remains unchanged.
	//
	// 	- If you set this parameter to DOWNGRADE, the new specifications must be lower than the original specifications. You must configure at least one of the cpu, storageSize, and coldStorageSize parameters. If you leave a parameter empty, the related configuration remains unchanged.
	//
	// This parameter is required.
	//
	// example:
	//
	// UPGRADE
	ScaleType *string `json:"scaleType,omitempty" xml:"scaleType,omitempty"`
	// The standard storage space of the instance. Unit: GB.
	//
	// > Ignore this parameter for pay-as-you-go instances.
	//
	// example:
	//
	// 1000G
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
