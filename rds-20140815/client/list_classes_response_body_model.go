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
	Items     []*ListClassesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	RegionId  *string                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ClassCode          *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	ClassGroup         *string `json:"ClassGroup,omitempty" xml:"ClassGroup,omitempty"`
	Cpu                *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	EncryptedMemory    *string `json:"EncryptedMemory,omitempty" xml:"EncryptedMemory,omitempty"`
	InstructionSetArch *string `json:"InstructionSetArch,omitempty" xml:"InstructionSetArch,omitempty"`
	MaxConnections     *string `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	MaxIOMBPS          *string `json:"MaxIOMBPS,omitempty" xml:"MaxIOMBPS,omitempty"`
	MaxIOPS            *string `json:"MaxIOPS,omitempty" xml:"MaxIOPS,omitempty"`
	MemoryClass        *string `json:"MemoryClass,omitempty" xml:"MemoryClass,omitempty"`
	ReferencePrice     *string `json:"ReferencePrice,omitempty" xml:"ReferencePrice,omitempty"`
	Category           *string `json:"category,omitempty" xml:"category,omitempty"`
	StorageType        *string `json:"storageType,omitempty" xml:"storageType,omitempty"`
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
