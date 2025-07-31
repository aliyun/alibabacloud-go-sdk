// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceSpecInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCPU(v string) *DescribeDBInstanceSpecInfoResponseBody
	GetCPU() *string
	SetDescription(v string) *DescribeDBInstanceSpecInfoResponseBody
	GetDescription() *string
	SetMemory(v string) *DescribeDBInstanceSpecInfoResponseBody
	GetMemory() *string
	SetRequestId(v string) *DescribeDBInstanceSpecInfoResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceSpecInfoResponseBody struct {
	// example:
	//
	// 2
	CPU         *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 8
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// 6B82A9EF-9961-5A31-A19F-009B709xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceSpecInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSpecInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSpecInfoResponseBody) GetCPU() *string {
	return s.CPU
}

func (s *DescribeDBInstanceSpecInfoResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeDBInstanceSpecInfoResponseBody) GetMemory() *string {
	return s.Memory
}

func (s *DescribeDBInstanceSpecInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceSpecInfoResponseBody) SetCPU(v string) *DescribeDBInstanceSpecInfoResponseBody {
	s.CPU = &v
	return s
}

func (s *DescribeDBInstanceSpecInfoResponseBody) SetDescription(v string) *DescribeDBInstanceSpecInfoResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeDBInstanceSpecInfoResponseBody) SetMemory(v string) *DescribeDBInstanceSpecInfoResponseBody {
	s.Memory = &v
	return s
}

func (s *DescribeDBInstanceSpecInfoResponseBody) SetRequestId(v string) *DescribeDBInstanceSpecInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceSpecInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
