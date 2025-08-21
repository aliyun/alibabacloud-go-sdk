// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthLimit(v int32) *DescribeInstanceSpecResponseBody
	GetBandwidthLimit() *int32
	SetCode(v int32) *DescribeInstanceSpecResponseBody
	GetCode() *int32
	SetDataDiskMaxSize(v int32) *DescribeInstanceSpecResponseBody
	GetDataDiskMaxSize() *int32
	SetDataDiskMinSize(v int32) *DescribeInstanceSpecResponseBody
	GetDataDiskMinSize() *int32
	SetInstanceSpecs(v *DescribeInstanceSpecResponseBodyInstanceSpecs) *DescribeInstanceSpecResponseBody
	GetInstanceSpecs() *DescribeInstanceSpecResponseBodyInstanceSpecs
	SetRequestId(v string) *DescribeInstanceSpecResponseBody
	GetRequestId() *string
	SetSystemDiskMaxSize(v int32) *DescribeInstanceSpecResponseBody
	GetSystemDiskMaxSize() *int32
}

type DescribeInstanceSpecResponseBody struct {
	// The bandwidth limit for a single instance. Unit: Mbit/s.
	//
	// example:
	//
	// 0
	BandwidthLimit *int32 `json:"BandwidthLimit,omitempty" xml:"BandwidthLimit,omitempty"`
	// The returned service code. A value of 0 indicates that the operation was successful.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The maximum capacity of a data disk. Unit: GB.
	//
	// example:
	//
	// 20015
	DataDiskMaxSize *int32 `json:"DataDiskMaxSize,omitempty" xml:"DataDiskMaxSize,omitempty"`
	// The minimum capacity of a data disk. Unit: GB.
	//
	// example:
	//
	// 0
	DataDiskMinSize *int32 `json:"DataDiskMinSize,omitempty" xml:"DataDiskMinSize,omitempty"`
	// The information about instance specifications.
	InstanceSpecs *DescribeInstanceSpecResponseBodyInstanceSpecs `json:"InstanceSpecs,omitempty" xml:"InstanceSpecs,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1ECC937A-AE0E-4626-BE51-DED1D6D1C888
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The maximum capacity of the system disk. Unit: GiB.
	//
	// example:
	//
	// 80
	SystemDiskMaxSize *int32 `json:"SystemDiskMaxSize,omitempty" xml:"SystemDiskMaxSize,omitempty"`
}

func (s DescribeInstanceSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecResponseBody) GetBandwidthLimit() *int32 {
	return s.BandwidthLimit
}

func (s *DescribeInstanceSpecResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeInstanceSpecResponseBody) GetDataDiskMaxSize() *int32 {
	return s.DataDiskMaxSize
}

func (s *DescribeInstanceSpecResponseBody) GetDataDiskMinSize() *int32 {
	return s.DataDiskMinSize
}

func (s *DescribeInstanceSpecResponseBody) GetInstanceSpecs() *DescribeInstanceSpecResponseBodyInstanceSpecs {
	return s.InstanceSpecs
}

func (s *DescribeInstanceSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceSpecResponseBody) GetSystemDiskMaxSize() *int32 {
	return s.SystemDiskMaxSize
}

func (s *DescribeInstanceSpecResponseBody) SetBandwidthLimit(v int32) *DescribeInstanceSpecResponseBody {
	s.BandwidthLimit = &v
	return s
}

func (s *DescribeInstanceSpecResponseBody) SetCode(v int32) *DescribeInstanceSpecResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstanceSpecResponseBody) SetDataDiskMaxSize(v int32) *DescribeInstanceSpecResponseBody {
	s.DataDiskMaxSize = &v
	return s
}

func (s *DescribeInstanceSpecResponseBody) SetDataDiskMinSize(v int32) *DescribeInstanceSpecResponseBody {
	s.DataDiskMinSize = &v
	return s
}

func (s *DescribeInstanceSpecResponseBody) SetInstanceSpecs(v *DescribeInstanceSpecResponseBodyInstanceSpecs) *DescribeInstanceSpecResponseBody {
	s.InstanceSpecs = v
	return s
}

func (s *DescribeInstanceSpecResponseBody) SetRequestId(v string) *DescribeInstanceSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSpecResponseBody) SetSystemDiskMaxSize(v int32) *DescribeInstanceSpecResponseBody {
	s.SystemDiskMaxSize = &v
	return s
}

func (s *DescribeInstanceSpecResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceSpecResponseBodyInstanceSpecs struct {
	InstanceSpec []*DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty" type:"Repeated"`
}

func (s DescribeInstanceSpecResponseBodyInstanceSpecs) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSpecResponseBodyInstanceSpecs) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecResponseBodyInstanceSpecs) GetInstanceSpec() []*DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec {
	return s.InstanceSpec
}

func (s *DescribeInstanceSpecResponseBodyInstanceSpecs) SetInstanceSpec(v []*DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec) *DescribeInstanceSpecResponseBodyInstanceSpecs {
	s.InstanceSpec = v
	return s
}

func (s *DescribeInstanceSpecResponseBodyInstanceSpecs) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec struct {
	// The number of CPU cores.
	//
	// example:
	//
	// 1
	Core *string `json:"Core,omitempty" xml:"Core,omitempty"`
	// The display name of the instance type.
	//
	// example:
	//
	// Computational 1C2G
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The type of the instance.
	//
	// example:
	//
	// ens.sn1.stiny
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The memory size. Unit: MB.
	//
	// example:
	//
	// 2048
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec) GetCore() *string {
	return s.Core
}

func (s *DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec) GetMemory() *string {
	return s.Memory
}

func (s *DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec) SetCore(v string) *DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec {
	s.Core = &v
	return s
}

func (s *DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec) SetDisplayName(v string) *DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec {
	s.DisplayName = &v
	return s
}

func (s *DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec) SetInstanceType(v string) *DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec) SetMemory(v string) *DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec {
	s.Memory = &v
	return s
}

func (s *DescribeInstanceSpecResponseBodyInstanceSpecsInstanceSpec) Validate() error {
	return dara.Validate(s)
}
