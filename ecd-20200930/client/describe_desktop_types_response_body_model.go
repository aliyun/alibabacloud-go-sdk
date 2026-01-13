// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopTypes(v []*DescribeDesktopTypesResponseBodyDesktopTypes) *DescribeDesktopTypesResponseBody
	GetDesktopTypes() []*DescribeDesktopTypesResponseBodyDesktopTypes
	SetRequestId(v string) *DescribeDesktopTypesResponseBody
	GetRequestId() *string
}

type DescribeDesktopTypesResponseBody struct {
	// The specifications.
	DesktopTypes []*DescribeDesktopTypesResponseBodyDesktopTypes `json:"DesktopTypes,omitempty" xml:"DesktopTypes,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDesktopTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopTypesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDesktopTypesResponseBody) GetDesktopTypes() []*DescribeDesktopTypesResponseBodyDesktopTypes {
	return s.DesktopTypes
}

func (s *DescribeDesktopTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDesktopTypesResponseBody) SetDesktopTypes(v []*DescribeDesktopTypesResponseBodyDesktopTypes) *DescribeDesktopTypesResponseBody {
	s.DesktopTypes = v
	return s
}

func (s *DescribeDesktopTypesResponseBody) SetRequestId(v string) *DescribeDesktopTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopTypesResponseBody) Validate() error {
	if s.DesktopTypes != nil {
		for _, item := range s.DesktopTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDesktopTypesResponseBodyDesktopTypes struct {
	// The number of vCPUs.
	//
	// example:
	//
	// 2
	CpuCount *string `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	// The size of the data disk. Unit: GiB.
	//
	// example:
	//
	// 150
	DataDiskSize *string `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the cloud desktop type.
	//
	// example:
	//
	// ecd.graphics.xlarge
	DesktopTypeId *string `json:"DesktopTypeId,omitempty" xml:"DesktopTypeId,omitempty"`
	// The status of the cloud desktop type. If SUFFICIENT is returned, the number of cloud desktops of the type is sufficient.
	//
	// example:
	//
	// SUFFICIENT
	DesktopTypeStatus *string `json:"DesktopTypeStatus,omitempty" xml:"DesktopTypeStatus,omitempty"`
	EnvId             *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	EnvType           *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The number of GPUs.
	//
	// example:
	//
	// 1
	GpuCount *float32 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	// The GPU memory size. For GPU-accelerated cloud computers, this return value is significant. Unit: MB.
	//
	// example:
	//
	// 2048
	GpuMemory *int32 `json:"GpuMemory,omitempty" xml:"GpuMemory,omitempty"`
	// The GPU memory.
	//
	// example:
	//
	// 16 GiB
	GpuSpec *string `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty"`
	// The family of the cloud desktop type.
	//
	// example:
	//
	// ecd.graphics
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	// The number of sessions supported by the specification.
	//
	// example:
	//
	// 4
	MaxSessionCount *int32 `json:"MaxSessionCount,omitempty" xml:"MaxSessionCount,omitempty"`
	// The memory size. Unit: MiB.
	//
	// example:
	//
	// 23552
	MemorySize *string `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	// The sales modes of the specifications.
	Scopes []*string `json:"Scopes,omitempty" xml:"Scopes,omitempty" type:"Repeated"`
	// The inventory status of the specification.
	//
	// Valid values:
	//
	// 	- Insufficient
	//
	// 	- Sufficient
	//
	// example:
	//
	// Sufficient
	StockState *string `json:"StockState,omitempty" xml:"StockState,omitempty"`
	// The size of the system disk. Unit: GiB.
	//
	// example:
	//
	// 150
	SystemDiskSize *string `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribeDesktopTypesResponseBodyDesktopTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopTypesResponseBodyDesktopTypes) GoString() string {
	return s.String()
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) GetCpuCount() *string {
	return s.CpuCount
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) GetDataDiskSize() *string {
	return s.DataDiskSize
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) GetDescription() *string {
	return s.Description
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) GetDesktopTypeId() *string {
	return s.DesktopTypeId
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) GetDesktopTypeStatus() *string {
	return s.DesktopTypeStatus
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) GetEnvId() *string {
	return s.EnvId
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) GetEnvType() *string {
	return s.EnvType
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) GetGpuCount() *float32 {
	return s.GpuCount
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) GetGpuMemory() *int32 {
	return s.GpuMemory
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) GetGpuSpec() *string {
	return s.GpuSpec
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) GetMaxSessionCount() *int32 {
	return s.MaxSessionCount
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) GetMemorySize() *string {
	return s.MemorySize
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) GetScopes() []*string {
	return s.Scopes
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) GetStockState() *string {
	return s.StockState
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) GetSystemDiskSize() *string {
	return s.SystemDiskSize
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetCpuCount(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.CpuCount = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetDataDiskSize(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetDescription(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.Description = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetDesktopTypeId(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.DesktopTypeId = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetDesktopTypeStatus(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.DesktopTypeStatus = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetEnvId(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.EnvId = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetEnvType(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.EnvType = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetGpuCount(v float32) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.GpuCount = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetGpuMemory(v int32) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.GpuMemory = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetGpuSpec(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.GpuSpec = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetInstanceTypeFamily(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetMaxSessionCount(v int32) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.MaxSessionCount = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetMemorySize(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.MemorySize = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetScopes(v []*string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.Scopes = v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetStockState(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.StockState = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetSystemDiskSize(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) Validate() error {
	return dara.Validate(s)
}
