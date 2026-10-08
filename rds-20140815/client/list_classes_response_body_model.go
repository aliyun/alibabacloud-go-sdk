// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClassesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*ListClassesResponseBodyItems) *ListClassesResponseBody
	GetItems() []*ListClassesResponseBodyItems
	SetRegionId(v string) *ListClassesResponseBody
	GetRegionId() *string
	SetRequestId(v string) *ListClassesResponseBody
	GetRequestId() *string
}

type ListClassesResponseBody struct {
	// The list of instance type information.
	Items []*ListClassesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CF8D35BF-263D-4F7B-883A-1163B79A9EC6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClassesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClassesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClassesResponseBody) GetItems() []*ListClassesResponseBodyItems {
	return s.Items
}

func (s *ListClassesResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *ListClassesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClassesResponseBody) SetItems(v []*ListClassesResponseBodyItems) *ListClassesResponseBody {
	s.Items = v
	return s
}

func (s *ListClassesResponseBody) SetRegionId(v string) *ListClassesResponseBody {
	s.RegionId = &v
	return s
}

func (s *ListClassesResponseBody) SetRequestId(v string) *ListClassesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClassesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClassesResponseBodyItems struct {
	// The instance type code. For more information, see [Primary instance types](https://help.aliyun.com/document_detail/26312.html) and [Read-only instance types](https://help.aliyun.com/document_detail/145759.html).
	//
	// example:
	//
	// mysql.n1.micro.1
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	// The instance family. For more information, see [Instance families](https://help.aliyun.com/document_detail/57184.html).
	//
	// example:
	//
	// general-purpose
	ClassGroup *string `json:"ClassGroup,omitempty" xml:"ClassGroup,omitempty"`
	// The number of CPU cores for the instance type. Unit: cores.
	//
	// example:
	//
	// 1
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The encrypted memory size for the security-enhanced instance family. Unit: GB.
	//
	// example:
	//
	// 4
	EncryptedMemory *string `json:"EncryptedMemory,omitempty" xml:"EncryptedMemory,omitempty"`
	// The architecture type of the instance type. Valid values:
	//
	// - If the instance uses the **x86*	- architecture, this parameter is empty by default.
	//
	// - If the instance uses the **arm*	- architecture, **arm*	- is returned.
	//
	// example:
	//
	// arm
	InstructionSetArch *string `json:"InstructionSetArch,omitempty" xml:"InstructionSetArch,omitempty"`
	// The maximum number of connections for the instance type.
	//
	// example:
	//
	// 2000
	MaxConnections *string `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	// The maximum I/O bandwidth for the instance type. Unit: Mbit/s.
	//
	// example:
	//
	// 1024Mbps
	MaxIOMBPS *string `json:"MaxIOMBPS,omitempty" xml:"MaxIOMBPS,omitempty"`
	// The maximum IOPS for the instance type.
	//
	// example:
	//
	// 10000
	MaxIOPS *string `json:"MaxIOPS,omitempty" xml:"MaxIOPS,omitempty"`
	// The memory size for the instance type. Unit: GB.
	//
	// example:
	//
	// 1GB
	MemoryClass *string `json:"MemoryClass,omitempty" xml:"MemoryClass,omitempty"`
	// The price for the instance type.
	//
	// <props="china">
	//
	// 	- Unit: cents (CNY).
	//
	// <props="intl">
	//
	// 	- Unit: cents (USD).
	//
	//
	// > 	- If you set the **CommodityCode*	- parameter to a pay-as-you-go commodity code, this parameter indicates the hourly price.
	//
	// > 	- If you set the **CommodityCode*	- parameter to a subscription commodity code, this parameter indicates the monthly price.
	//
	// example:
	//
	// 2500
	ReferencePrice *string `json:"ReferencePrice,omitempty" xml:"ReferencePrice,omitempty"`
	// The instance edition. Valid values:
	//
	// 	- Regular instances
	//
	//     	- **Basic**: Basic Edition.
	//
	//     	- **HighAvailability**: High availability series.
	//
	//     	- **cluster**: MySQL or PostgreSQL Cluster Edition.
	//
	//     	- **AlwaysOn**: SQL Server Cluster Edition.
	//
	//     	- **Finance**: RDS Enterprise Edition.
	//
	// 	- Serverless instances
	//
	//     	- **serverless_basic**: Serverless Basic Edition. (Applicable only to MySQL and PostgreSQL)
	//
	//     	- **serverless_standard**: Serverless high availability series. (Applicable only to MySQL and PostgreSQL)
	//
	//     	- **serverless_ha**: SQL Server Serverless high availability series.
	//
	// example:
	//
	// Basic
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The instance storage type.
	//
	// example:
	//
	// cloud_essd
	StorageType *string `json:"storageType,omitempty" xml:"storageType,omitempty"`
}

func (s ListClassesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListClassesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListClassesResponseBodyItems) GetClassCode() *string {
	return s.ClassCode
}

func (s *ListClassesResponseBodyItems) GetClassGroup() *string {
	return s.ClassGroup
}

func (s *ListClassesResponseBodyItems) GetCpu() *string {
	return s.Cpu
}

func (s *ListClassesResponseBodyItems) GetEncryptedMemory() *string {
	return s.EncryptedMemory
}

func (s *ListClassesResponseBodyItems) GetInstructionSetArch() *string {
	return s.InstructionSetArch
}

func (s *ListClassesResponseBodyItems) GetMaxConnections() *string {
	return s.MaxConnections
}

func (s *ListClassesResponseBodyItems) GetMaxIOMBPS() *string {
	return s.MaxIOMBPS
}

func (s *ListClassesResponseBodyItems) GetMaxIOPS() *string {
	return s.MaxIOPS
}

func (s *ListClassesResponseBodyItems) GetMemoryClass() *string {
	return s.MemoryClass
}

func (s *ListClassesResponseBodyItems) GetReferencePrice() *string {
	return s.ReferencePrice
}

func (s *ListClassesResponseBodyItems) GetCategory() *string {
	return s.Category
}

func (s *ListClassesResponseBodyItems) GetStorageType() *string {
	return s.StorageType
}

func (s *ListClassesResponseBodyItems) SetClassCode(v string) *ListClassesResponseBodyItems {
	s.ClassCode = &v
	return s
}

func (s *ListClassesResponseBodyItems) SetClassGroup(v string) *ListClassesResponseBodyItems {
	s.ClassGroup = &v
	return s
}

func (s *ListClassesResponseBodyItems) SetCpu(v string) *ListClassesResponseBodyItems {
	s.Cpu = &v
	return s
}

func (s *ListClassesResponseBodyItems) SetEncryptedMemory(v string) *ListClassesResponseBodyItems {
	s.EncryptedMemory = &v
	return s
}

func (s *ListClassesResponseBodyItems) SetInstructionSetArch(v string) *ListClassesResponseBodyItems {
	s.InstructionSetArch = &v
	return s
}

func (s *ListClassesResponseBodyItems) SetMaxConnections(v string) *ListClassesResponseBodyItems {
	s.MaxConnections = &v
	return s
}

func (s *ListClassesResponseBodyItems) SetMaxIOMBPS(v string) *ListClassesResponseBodyItems {
	s.MaxIOMBPS = &v
	return s
}

func (s *ListClassesResponseBodyItems) SetMaxIOPS(v string) *ListClassesResponseBodyItems {
	s.MaxIOPS = &v
	return s
}

func (s *ListClassesResponseBodyItems) SetMemoryClass(v string) *ListClassesResponseBodyItems {
	s.MemoryClass = &v
	return s
}

func (s *ListClassesResponseBodyItems) SetReferencePrice(v string) *ListClassesResponseBodyItems {
	s.ReferencePrice = &v
	return s
}

func (s *ListClassesResponseBodyItems) SetCategory(v string) *ListClassesResponseBodyItems {
	s.Category = &v
	return s
}

func (s *ListClassesResponseBodyItems) SetStorageType(v string) *ListClassesResponseBodyItems {
	s.StorageType = &v
	return s
}

func (s *ListClassesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
