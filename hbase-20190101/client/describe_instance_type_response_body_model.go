// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceTypeSpecList(v *DescribeInstanceTypeResponseBodyInstanceTypeSpecList) *DescribeInstanceTypeResponseBody
	GetInstanceTypeSpecList() *DescribeInstanceTypeResponseBodyInstanceTypeSpecList
	SetRequestId(v string) *DescribeInstanceTypeResponseBody
	GetRequestId() *string
}

type DescribeInstanceTypeResponseBody struct {
	InstanceTypeSpecList *DescribeInstanceTypeResponseBodyInstanceTypeSpecList `json:"InstanceTypeSpecList,omitempty" xml:"InstanceTypeSpecList,omitempty" type:"Struct"`
	// example:
	//
	// DD23BBB4-64C2-42A4-B2E2-7E56C7AA815A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypeResponseBody) GetInstanceTypeSpecList() *DescribeInstanceTypeResponseBodyInstanceTypeSpecList {
	return s.InstanceTypeSpecList
}

func (s *DescribeInstanceTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceTypeResponseBody) SetInstanceTypeSpecList(v *DescribeInstanceTypeResponseBodyInstanceTypeSpecList) *DescribeInstanceTypeResponseBody {
	s.InstanceTypeSpecList = v
	return s
}

func (s *DescribeInstanceTypeResponseBody) SetRequestId(v string) *DescribeInstanceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceTypeResponseBody) Validate() error {
	if s.InstanceTypeSpecList != nil {
		if err := s.InstanceTypeSpecList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceTypeResponseBodyInstanceTypeSpecList struct {
	InstanceTypeSpec []*DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec `json:"InstanceTypeSpec,omitempty" xml:"InstanceTypeSpec,omitempty" type:"Repeated"`
}

func (s DescribeInstanceTypeResponseBodyInstanceTypeSpecList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypeResponseBodyInstanceTypeSpecList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypeResponseBodyInstanceTypeSpecList) GetInstanceTypeSpec() []*DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec {
	return s.InstanceTypeSpec
}

func (s *DescribeInstanceTypeResponseBodyInstanceTypeSpecList) SetInstanceTypeSpec(v []*DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec) *DescribeInstanceTypeResponseBodyInstanceTypeSpecList {
	s.InstanceTypeSpec = v
	return s
}

func (s *DescribeInstanceTypeResponseBodyInstanceTypeSpecList) Validate() error {
	if s.InstanceTypeSpec != nil {
		for _, item := range s.InstanceTypeSpec {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec struct {
	// example:
	//
	// 8
	CpuSize *int64 `json:"CpuSize,omitempty" xml:"CpuSize,omitempty"`
	// example:
	//
	// hbase.n2.4xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 16
	MemSize *int64 `json:"MemSize,omitempty" xml:"MemSize,omitempty"`
}

func (s DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec) GetCpuSize() *int64 {
	return s.CpuSize
}

func (s *DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec) GetMemSize() *int64 {
	return s.MemSize
}

func (s *DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec) SetCpuSize(v int64) *DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec {
	s.CpuSize = &v
	return s
}

func (s *DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec) SetInstanceType(v string) *DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec) SetMemSize(v int64) *DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec {
	s.MemSize = &v
	return s
}

func (s *DescribeInstanceTypeResponseBodyInstanceTypeSpecListInstanceTypeSpec) Validate() error {
	return dara.Validate(s)
}
