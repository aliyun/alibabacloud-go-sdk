// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceGroupSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeResourceGroupSpecResponseBody
	GetRequestId() *string
	SetSpecs(v []*DescribeResourceGroupSpecResponseBodySpecs) *DescribeResourceGroupSpecResponseBody
	GetSpecs() []*DescribeResourceGroupSpecResponseBodySpecs
}

type DescribeResourceGroupSpecResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 10226189-4391-5B10-97AF-5CA5XXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried specifications.
	Specs []*DescribeResourceGroupSpecResponseBodySpecs `json:"Specs,omitempty" xml:"Specs,omitempty" type:"Repeated"`
}

func (s DescribeResourceGroupSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceGroupSpecResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceGroupSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourceGroupSpecResponseBody) GetSpecs() []*DescribeResourceGroupSpecResponseBodySpecs {
	return s.Specs
}

func (s *DescribeResourceGroupSpecResponseBody) SetRequestId(v string) *DescribeResourceGroupSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceGroupSpecResponseBody) SetSpecs(v []*DescribeResourceGroupSpecResponseBodySpecs) *DescribeResourceGroupSpecResponseBody {
	s.Specs = v
	return s
}

func (s *DescribeResourceGroupSpecResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceGroupSpecResponseBodySpecs struct {
	// The allocation units supported by this specification.
	AllocateUnits []*string `json:"AllocateUnits,omitempty" xml:"AllocateUnits,omitempty" type:"Repeated"`
	// The maximum number of resource groups that can be used with this specification.
	//
	// example:
	//
	// 8
	MaxQuantity *int32 `json:"MaxQuantity,omitempty" xml:"MaxQuantity,omitempty"`
	// The name of the specification.
	//
	// example:
	//
	// xlarge
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The resource type.
	//
	// example:
	//
	// GPU
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeResourceGroupSpecResponseBodySpecs) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceGroupSpecResponseBodySpecs) GoString() string {
	return s.String()
}

func (s *DescribeResourceGroupSpecResponseBodySpecs) GetAllocateUnits() []*string {
	return s.AllocateUnits
}

func (s *DescribeResourceGroupSpecResponseBodySpecs) GetMaxQuantity() *int32 {
	return s.MaxQuantity
}

func (s *DescribeResourceGroupSpecResponseBodySpecs) GetName() *string {
	return s.Name
}

func (s *DescribeResourceGroupSpecResponseBodySpecs) GetType() *string {
	return s.Type
}

func (s *DescribeResourceGroupSpecResponseBodySpecs) SetAllocateUnits(v []*string) *DescribeResourceGroupSpecResponseBodySpecs {
	s.AllocateUnits = v
	return s
}

func (s *DescribeResourceGroupSpecResponseBodySpecs) SetMaxQuantity(v int32) *DescribeResourceGroupSpecResponseBodySpecs {
	s.MaxQuantity = &v
	return s
}

func (s *DescribeResourceGroupSpecResponseBodySpecs) SetName(v string) *DescribeResourceGroupSpecResponseBodySpecs {
	s.Name = &v
	return s
}

func (s *DescribeResourceGroupSpecResponseBodySpecs) SetType(v string) *DescribeResourceGroupSpecResponseBodySpecs {
	s.Type = &v
	return s
}

func (s *DescribeResourceGroupSpecResponseBodySpecs) Validate() error {
	return dara.Validate(s)
}
