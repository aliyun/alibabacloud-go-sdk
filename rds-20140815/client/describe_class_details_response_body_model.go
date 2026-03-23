// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClassDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *DescribeClassDetailsResponseBody
	GetCategory() *string
	SetClassCode(v string) *DescribeClassDetailsResponseBody
	GetClassCode() *string
	SetClassGroup(v string) *DescribeClassDetailsResponseBody
	GetClassGroup() *string
	SetCpu(v string) *DescribeClassDetailsResponseBody
	GetCpu() *string
	SetDBInstanceStorageType(v string) *DescribeClassDetailsResponseBody
	GetDBInstanceStorageType() *string
	SetInstructionSetArch(v string) *DescribeClassDetailsResponseBody
	GetInstructionSetArch() *string
	SetMaxConnections(v string) *DescribeClassDetailsResponseBody
	GetMaxConnections() *string
	SetMaxIOMBPS(v string) *DescribeClassDetailsResponseBody
	GetMaxIOMBPS() *string
	SetMaxIOPS(v string) *DescribeClassDetailsResponseBody
	GetMaxIOPS() *string
	SetMemoryClass(v string) *DescribeClassDetailsResponseBody
	GetMemoryClass() *string
	SetReferencePrice(v string) *DescribeClassDetailsResponseBody
	GetReferencePrice() *string
	SetRequestId(v string) *DescribeClassDetailsResponseBody
	GetRequestId() *string
}

type DescribeClassDetailsResponseBody struct {
	Category              *string `json:"Category,omitempty" xml:"Category,omitempty"`
	ClassCode             *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	ClassGroup            *string `json:"ClassGroup,omitempty" xml:"ClassGroup,omitempty"`
	Cpu                   *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	InstructionSetArch    *string `json:"InstructionSetArch,omitempty" xml:"InstructionSetArch,omitempty"`
	MaxConnections        *string `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	MaxIOMBPS             *string `json:"MaxIOMBPS,omitempty" xml:"MaxIOMBPS,omitempty"`
	MaxIOPS               *string `json:"MaxIOPS,omitempty" xml:"MaxIOPS,omitempty"`
	MemoryClass           *string `json:"MemoryClass,omitempty" xml:"MemoryClass,omitempty"`
	ReferencePrice        *string `json:"ReferencePrice,omitempty" xml:"ReferencePrice,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClassDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClassDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClassDetailsResponseBody) GetCategory() *string {
	return s.Category
}

func (s *DescribeClassDetailsResponseBody) GetClassCode() *string {
	return s.ClassCode
}

func (s *DescribeClassDetailsResponseBody) GetClassGroup() *string {
	return s.ClassGroup
}

func (s *DescribeClassDetailsResponseBody) GetCpu() *string {
	return s.Cpu
}

func (s *DescribeClassDetailsResponseBody) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *DescribeClassDetailsResponseBody) GetInstructionSetArch() *string {
	return s.InstructionSetArch
}

func (s *DescribeClassDetailsResponseBody) GetMaxConnections() *string {
	return s.MaxConnections
}

func (s *DescribeClassDetailsResponseBody) GetMaxIOMBPS() *string {
	return s.MaxIOMBPS
}

func (s *DescribeClassDetailsResponseBody) GetMaxIOPS() *string {
	return s.MaxIOPS
}

func (s *DescribeClassDetailsResponseBody) GetMemoryClass() *string {
	return s.MemoryClass
}

func (s *DescribeClassDetailsResponseBody) GetReferencePrice() *string {
	return s.ReferencePrice
}

func (s *DescribeClassDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClassDetailsResponseBody) SetCategory(v string) *DescribeClassDetailsResponseBody {
	s.Category = &v
	return s
}

func (s *DescribeClassDetailsResponseBody) SetClassCode(v string) *DescribeClassDetailsResponseBody {
	s.ClassCode = &v
	return s
}

func (s *DescribeClassDetailsResponseBody) SetClassGroup(v string) *DescribeClassDetailsResponseBody {
	s.ClassGroup = &v
	return s
}

func (s *DescribeClassDetailsResponseBody) SetCpu(v string) *DescribeClassDetailsResponseBody {
	s.Cpu = &v
	return s
}

func (s *DescribeClassDetailsResponseBody) SetDBInstanceStorageType(v string) *DescribeClassDetailsResponseBody {
	s.DBInstanceStorageType = &v
	return s
}

func (s *DescribeClassDetailsResponseBody) SetInstructionSetArch(v string) *DescribeClassDetailsResponseBody {
	s.InstructionSetArch = &v
	return s
}

func (s *DescribeClassDetailsResponseBody) SetMaxConnections(v string) *DescribeClassDetailsResponseBody {
	s.MaxConnections = &v
	return s
}

func (s *DescribeClassDetailsResponseBody) SetMaxIOMBPS(v string) *DescribeClassDetailsResponseBody {
	s.MaxIOMBPS = &v
	return s
}

func (s *DescribeClassDetailsResponseBody) SetMaxIOPS(v string) *DescribeClassDetailsResponseBody {
	s.MaxIOPS = &v
	return s
}

func (s *DescribeClassDetailsResponseBody) SetMemoryClass(v string) *DescribeClassDetailsResponseBody {
	s.MemoryClass = &v
	return s
}

func (s *DescribeClassDetailsResponseBody) SetReferencePrice(v string) *DescribeClassDetailsResponseBody {
	s.ReferencePrice = &v
	return s
}

func (s *DescribeClassDetailsResponseBody) SetRequestId(v string) *DescribeClassDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClassDetailsResponseBody) Validate() error {
	return dara.Validate(s)
}
