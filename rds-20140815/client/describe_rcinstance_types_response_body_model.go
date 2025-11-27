// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstanceTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceTypes(v *DescribeRCInstanceTypesResponseBodyInstanceTypes) *DescribeRCInstanceTypesResponseBody
	GetInstanceTypes() *DescribeRCInstanceTypesResponseBodyInstanceTypes
	SetRequestId(v string) *DescribeRCInstanceTypesResponseBody
	GetRequestId() *string
}

type DescribeRCInstanceTypesResponseBody struct {
	// The information about the instance types.
	InstanceTypes *DescribeRCInstanceTypesResponseBodyInstanceTypes `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F2911788-25E8-42E5-A3A3-1B38D263F01E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRCInstanceTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceTypesResponseBody) GetInstanceTypes() *DescribeRCInstanceTypesResponseBodyInstanceTypes {
	return s.InstanceTypes
}

func (s *DescribeRCInstanceTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCInstanceTypesResponseBody) SetInstanceTypes(v *DescribeRCInstanceTypesResponseBodyInstanceTypes) *DescribeRCInstanceTypesResponseBody {
	s.InstanceTypes = v
	return s
}

func (s *DescribeRCInstanceTypesResponseBody) SetRequestId(v string) *DescribeRCInstanceTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCInstanceTypesResponseBody) Validate() error {
	if s.InstanceTypes != nil {
		if err := s.InstanceTypes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRCInstanceTypesResponseBodyInstanceTypes struct {
	// The instance types.
	InstanceType []*DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" type:"Repeated"`
}

func (s DescribeRCInstanceTypesResponseBodyInstanceTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceTypesResponseBodyInstanceTypes) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceTypesResponseBodyInstanceTypes) GetInstanceType() []*DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType {
	return s.InstanceType
}

func (s *DescribeRCInstanceTypesResponseBodyInstanceTypes) SetInstanceType(v []*DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType) *DescribeRCInstanceTypesResponseBodyInstanceTypes {
	s.InstanceType = v
	return s
}

func (s *DescribeRCInstanceTypesResponseBodyInstanceTypes) Validate() error {
	if s.InstanceType != nil {
		for _, item := range s.InstanceType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType struct {
	// The maximum number of CPU cores.
	//
	// example:
	//
	// 32
	CpuCoreCount *int32 `json:"CpuCoreCount,omitempty" xml:"CpuCoreCount,omitempty"`
	// example:
	//
	// 16
	DiskQuantity *int64 `json:"DiskQuantity,omitempty" xml:"DiskQuantity,omitempty"`
	// The ID of the instance family.
	//
	// example:
	//
	// gn8.cm
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	// The instance type of the instance.
	//
	// example:
	//
	// rds.gna8.2xlarge.8cm
	InstanceTypeId *string `json:"InstanceTypeId,omitempty" xml:"InstanceTypeId,omitempty"`
	// The memory size of the instance type. Unit: GiB.
	//
	// example:
	//
	// 256
	MemorySize *int32 `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
}

func (s DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType) GetCpuCoreCount() *int32 {
	return s.CpuCoreCount
}

func (s *DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType) GetDiskQuantity() *int64 {
	return s.DiskQuantity
}

func (s *DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType) GetInstanceTypeId() *string {
	return s.InstanceTypeId
}

func (s *DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType) GetMemorySize() *int32 {
	return s.MemorySize
}

func (s *DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType) SetCpuCoreCount(v int32) *DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.CpuCoreCount = &v
	return s
}

func (s *DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType) SetDiskQuantity(v int64) *DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.DiskQuantity = &v
	return s
}

func (s *DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType) SetInstanceTypeFamily(v string) *DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType) SetInstanceTypeId(v string) *DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.InstanceTypeId = &v
	return s
}

func (s *DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType) SetMemorySize(v int32) *DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.MemorySize = &v
	return s
}

func (s *DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType) Validate() error {
	return dara.Validate(s)
}
