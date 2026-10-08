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
	InstanceTypes *DescribeRCInstanceTypesResponseBodyInstanceTypes `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Struct"`
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CpuCoreCount *int32 `json:"CpuCoreCount,omitempty" xml:"CpuCoreCount,omitempty"`
	// example:
	//
	// 16
	DiskQuantity       *int64  `json:"DiskQuantity,omitempty" xml:"DiskQuantity,omitempty"`
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	InstanceTypeId     *string `json:"InstanceTypeId,omitempty" xml:"InstanceTypeId,omitempty"`
	JumboFrameSupport  *bool   `json:"JumboFrameSupport,omitempty" xml:"JumboFrameSupport,omitempty"`
	MemorySize         *int32  `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
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

func (s *DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType) GetJumboFrameSupport() *bool {
	return s.JumboFrameSupport
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

func (s *DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType) SetJumboFrameSupport(v bool) *DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.JumboFrameSupport = &v
	return s
}

func (s *DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType) SetMemorySize(v int32) *DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.MemorySize = &v
	return s
}

func (s *DescribeRCInstanceTypesResponseBodyInstanceTypesInstanceType) Validate() error {
	return dara.Validate(s)
}
