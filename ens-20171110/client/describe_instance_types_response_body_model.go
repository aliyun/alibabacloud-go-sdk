// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeInstanceTypesResponseBody
	GetCode() *int32
	SetInstanceTypes(v *DescribeInstanceTypesResponseBodyInstanceTypes) *DescribeInstanceTypesResponseBody
	GetInstanceTypes() *DescribeInstanceTypesResponseBodyInstanceTypes
	SetRequestId(v string) *DescribeInstanceTypesResponseBody
	GetRequestId() *string
}

type DescribeInstanceTypesResponseBody struct {
	// The business code. A value of 0 indicates success. An error code is returned upon failure.
	//
	// example:
	//
	// 0
	Code          *int32                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	InstanceTypes *DescribeInstanceTypesResponseBodyInstanceTypes `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D452D190-BADF-5D09-910D-599B96D42AAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeInstanceTypesResponseBody) GetInstanceTypes() *DescribeInstanceTypesResponseBodyInstanceTypes {
	return s.InstanceTypes
}

func (s *DescribeInstanceTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceTypesResponseBody) SetCode(v int32) *DescribeInstanceTypesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstanceTypesResponseBody) SetInstanceTypes(v *DescribeInstanceTypesResponseBodyInstanceTypes) *DescribeInstanceTypesResponseBody {
	s.InstanceTypes = v
	return s
}

func (s *DescribeInstanceTypesResponseBody) SetRequestId(v string) *DescribeInstanceTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceTypesResponseBody) Validate() error {
	if s.InstanceTypes != nil {
		if err := s.InstanceTypes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceTypesResponseBodyInstanceTypes struct {
	InstanceType []*DescribeInstanceTypesResponseBodyInstanceTypesInstanceType `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" type:"Repeated"`
}

func (s DescribeInstanceTypesResponseBodyInstanceTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypesResponseBodyInstanceTypes) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypes) GetInstanceType() []*DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	return s.InstanceType
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypes) SetInstanceType(v []*DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) *DescribeInstanceTypesResponseBodyInstanceTypes {
	s.InstanceType = v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypes) Validate() error {
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

type DescribeInstanceTypesResponseBodyInstanceTypesInstanceType struct {
	CpuCoreCount        *int32  `json:"CpuCoreCount,omitempty" xml:"CpuCoreCount,omitempty"`
	InstanceBandwidthRx *int64  `json:"InstanceBandwidthRx,omitempty" xml:"InstanceBandwidthRx,omitempty"`
	InstanceBandwidthTx *int64  `json:"InstanceBandwidthTx,omitempty" xml:"InstanceBandwidthTx,omitempty"`
	InstanceType        *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceTypeId      *string `json:"InstanceTypeId,omitempty" xml:"InstanceTypeId,omitempty"`
	InstanceTypeName    *string `json:"InstanceTypeName,omitempty" xml:"InstanceTypeName,omitempty"`
	MemorySize          *int32  `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetCpuCoreCount() *int32 {
	return s.CpuCoreCount
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetInstanceBandwidthRx() *int64 {
	return s.InstanceBandwidthRx
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetInstanceBandwidthTx() *int64 {
	return s.InstanceBandwidthTx
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetInstanceTypeId() *string {
	return s.InstanceTypeId
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetInstanceTypeName() *string {
	return s.InstanceTypeName
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetMemorySize() *int32 {
	return s.MemorySize
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetCpuCoreCount(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.CpuCoreCount = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetInstanceBandwidthRx(v int64) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.InstanceBandwidthRx = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetInstanceBandwidthTx(v int64) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.InstanceBandwidthTx = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetInstanceType(v string) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetInstanceTypeId(v string) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.InstanceTypeId = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetInstanceTypeName(v string) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.InstanceTypeName = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetMemorySize(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.MemorySize = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) Validate() error {
	return dara.Validate(s)
}
