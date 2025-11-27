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
	// The RDS edition of the instance. Valid values:
	//
	// 	- **Basic**: RDS Basic Edition
	//
	// 	- **HighAvailability**: RDS High-availability Edition
	//
	// 	- **AlwaysOn**: RDS Cluster Edition
	//
	// 	- **Finance**: RDS Enterprise Edition
	//
	// example:
	//
	// Basic
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The code of the instance type.
	//
	// example:
	//
	// mysql.n2.medium.1
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	// The instance family of the instance.
	//
	// example:
	//
	// x
	ClassGroup *string `json:"ClassGroup,omitempty" xml:"ClassGroup,omitempty"`
	// The number of CPU cores that are supported by the instance type. Unit: cores.
	//
	// example:
	//
	// 4
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The storage type of the instance. Valid values:
	//
	// 	- **local_ssd**: local SSDs
	//
	// 	- **cloud_ssd**: standard SSDs
	//
	// 	- **cloud_essd**: enhanced SSDs (ESSDs) of performance level 1 (PL1)
	//
	// 	- **cloud_essd2**: ESSDs of PL2
	//
	// 	- **cloud_essd3**: ESSD of PL3
	//
	// example:
	//
	// local_ssd
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// The architecture of the instance.
	//
	// example:
	//
	// x86
	InstructionSetArch *string `json:"InstructionSetArch,omitempty" xml:"InstructionSetArch,omitempty"`
	// The maximum number of connections.
	//
	// example:
	//
	// 4000
	MaxConnections *string `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	// The maximum I/O bandwidth that is supported by the instance type. Unit: Mbit/s.
	//
	// example:
	//
	// 1024
	MaxIOMBPS *string `json:"MaxIOMBPS,omitempty" xml:"MaxIOMBPS,omitempty"`
	// The maximum input/output operations per second (IOPS) that is supported by the instance type. Unit: operations per second.
	//
	// example:
	//
	// N/A
	MaxIOPS *string `json:"MaxIOPS,omitempty" xml:"MaxIOPS,omitempty"`
	// The memory size. Unit: GB.
	//
	// example:
	//
	// 2GB
	MemoryClass *string `json:"MemoryClass,omitempty" xml:"MemoryClass,omitempty"`
	// The price.
	//
	// Unit: cents (US dollars).
	//
	// > 	- If you set the CommodityCode parameter to a value that indicates the pay-as-you-go billing method, the ReferencePrice parameter specifies the hourly fee that you must pay.
	//
	// > 	- If you set the CommodityCode parameter to a value that indicates the subscription billing method, the ReferencePrice parameter specifies the monthly fee that you must pay.
	//
	// example:
	//
	// 13400
	ReferencePrice *string `json:"ReferencePrice,omitempty" xml:"ReferencePrice,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E9DD55F4-1A5F-48CA-BA57-DFB3CA8C4C34
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
