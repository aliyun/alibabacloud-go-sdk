// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllDBInstanceClassResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClassCodeList(v []*DescribeAllDBInstanceClassResponseBodyClassCodeList) *DescribeAllDBInstanceClassResponseBody
	GetClassCodeList() []*DescribeAllDBInstanceClassResponseBodyClassCodeList
	SetRequestId(v string) *DescribeAllDBInstanceClassResponseBody
	GetRequestId() *string
}

type DescribeAllDBInstanceClassResponseBody struct {
	// The instance specifications.
	ClassCodeList []*DescribeAllDBInstanceClassResponseBodyClassCodeList `json:"ClassCodeList,omitempty" xml:"ClassCodeList,omitempty" type:"Repeated"`
	// example:
	//
	// 4773E4EC-025D-509F-AEA9-D53123FDFB0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAllDBInstanceClassResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllDBInstanceClassResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAllDBInstanceClassResponseBody) GetClassCodeList() []*DescribeAllDBInstanceClassResponseBodyClassCodeList {
	return s.ClassCodeList
}

func (s *DescribeAllDBInstanceClassResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAllDBInstanceClassResponseBody) SetClassCodeList(v []*DescribeAllDBInstanceClassResponseBodyClassCodeList) *DescribeAllDBInstanceClassResponseBody {
	s.ClassCodeList = v
	return s
}

func (s *DescribeAllDBInstanceClassResponseBody) SetRequestId(v string) *DescribeAllDBInstanceClassResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAllDBInstanceClassResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAllDBInstanceClassResponseBodyClassCodeList struct {
	// example:
	//
	// selectdb.xlarge
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	// example:
	//
	// 4
	CpuCores *int64 `json:"CpuCores,omitempty" xml:"CpuCores,omitempty"`
	// example:
	//
	// 200
	DefaultStorageInGB *int64 `json:"DefaultStorageInGB,omitempty" xml:"DefaultStorageInGB,omitempty"`
	// example:
	//
	// 2000
	MaxStorageInGB *int64 `json:"MaxStorageInGB,omitempty" xml:"MaxStorageInGB,omitempty"`
	// The memory size.
	//
	// example:
	//
	// 32
	MemoryInGB *int64 `json:"MemoryInGB,omitempty" xml:"MemoryInGB,omitempty"`
	// example:
	//
	// 100
	MinStorageInGB *int64 `json:"MinStorageInGB,omitempty" xml:"MinStorageInGB,omitempty"`
	// example:
	//
	// 100
	StepStorageInGB *int64 `json:"StepStorageInGB,omitempty" xml:"StepStorageInGB,omitempty"`
}

func (s DescribeAllDBInstanceClassResponseBodyClassCodeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllDBInstanceClassResponseBodyClassCodeList) GoString() string {
	return s.String()
}

func (s *DescribeAllDBInstanceClassResponseBodyClassCodeList) GetClassCode() *string {
	return s.ClassCode
}

func (s *DescribeAllDBInstanceClassResponseBodyClassCodeList) GetCpuCores() *int64 {
	return s.CpuCores
}

func (s *DescribeAllDBInstanceClassResponseBodyClassCodeList) GetDefaultStorageInGB() *int64 {
	return s.DefaultStorageInGB
}

func (s *DescribeAllDBInstanceClassResponseBodyClassCodeList) GetMaxStorageInGB() *int64 {
	return s.MaxStorageInGB
}

func (s *DescribeAllDBInstanceClassResponseBodyClassCodeList) GetMemoryInGB() *int64 {
	return s.MemoryInGB
}

func (s *DescribeAllDBInstanceClassResponseBodyClassCodeList) GetMinStorageInGB() *int64 {
	return s.MinStorageInGB
}

func (s *DescribeAllDBInstanceClassResponseBodyClassCodeList) GetStepStorageInGB() *int64 {
	return s.StepStorageInGB
}

func (s *DescribeAllDBInstanceClassResponseBodyClassCodeList) SetClassCode(v string) *DescribeAllDBInstanceClassResponseBodyClassCodeList {
	s.ClassCode = &v
	return s
}

func (s *DescribeAllDBInstanceClassResponseBodyClassCodeList) SetCpuCores(v int64) *DescribeAllDBInstanceClassResponseBodyClassCodeList {
	s.CpuCores = &v
	return s
}

func (s *DescribeAllDBInstanceClassResponseBodyClassCodeList) SetDefaultStorageInGB(v int64) *DescribeAllDBInstanceClassResponseBodyClassCodeList {
	s.DefaultStorageInGB = &v
	return s
}

func (s *DescribeAllDBInstanceClassResponseBodyClassCodeList) SetMaxStorageInGB(v int64) *DescribeAllDBInstanceClassResponseBodyClassCodeList {
	s.MaxStorageInGB = &v
	return s
}

func (s *DescribeAllDBInstanceClassResponseBodyClassCodeList) SetMemoryInGB(v int64) *DescribeAllDBInstanceClassResponseBodyClassCodeList {
	s.MemoryInGB = &v
	return s
}

func (s *DescribeAllDBInstanceClassResponseBodyClassCodeList) SetMinStorageInGB(v int64) *DescribeAllDBInstanceClassResponseBodyClassCodeList {
	s.MinStorageInGB = &v
	return s
}

func (s *DescribeAllDBInstanceClassResponseBodyClassCodeList) SetStepStorageInGB(v int64) *DescribeAllDBInstanceClassResponseBodyClassCodeList {
	s.StepStorageInGB = &v
	return s
}

func (s *DescribeAllDBInstanceClassResponseBodyClassCodeList) Validate() error {
	return dara.Validate(s)
}
