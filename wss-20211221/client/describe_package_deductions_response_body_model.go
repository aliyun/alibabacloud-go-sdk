// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePackageDeductionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeductions(v []*DescribePackageDeductionsResponseBodyDeductions) *DescribePackageDeductionsResponseBody
	GetDeductions() []*DescribePackageDeductionsResponseBodyDeductions
	SetPageNum(v int32) *DescribePackageDeductionsResponseBody
	GetPageNum() *int32
	SetPageSize(v int32) *DescribePackageDeductionsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribePackageDeductionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribePackageDeductionsResponseBody
	GetTotalCount() *int64
	SetTotalUsedCoreTime(v float32) *DescribePackageDeductionsResponseBody
	GetTotalUsedCoreTime() *float32
	SetTotalUsedTime(v int64) *DescribePackageDeductionsResponseBody
	GetTotalUsedTime() *int64
}

type DescribePackageDeductionsResponseBody struct {
	Deductions []*DescribePackageDeductionsResponseBodyDeductions `json:"Deductions,omitempty" xml:"Deductions,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 833C4D2C-09C7-5CE6-8159-06758B964970
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount        *int64   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalUsedCoreTime *float32 `json:"TotalUsedCoreTime,omitempty" xml:"TotalUsedCoreTime,omitempty"`
	TotalUsedTime     *int64   `json:"TotalUsedTime,omitempty" xml:"TotalUsedTime,omitempty"`
}

func (s DescribePackageDeductionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePackageDeductionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePackageDeductionsResponseBody) GetDeductions() []*DescribePackageDeductionsResponseBodyDeductions {
	return s.Deductions
}

func (s *DescribePackageDeductionsResponseBody) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribePackageDeductionsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePackageDeductionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePackageDeductionsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePackageDeductionsResponseBody) GetTotalUsedCoreTime() *float32 {
	return s.TotalUsedCoreTime
}

func (s *DescribePackageDeductionsResponseBody) GetTotalUsedTime() *int64 {
	return s.TotalUsedTime
}

func (s *DescribePackageDeductionsResponseBody) SetDeductions(v []*DescribePackageDeductionsResponseBodyDeductions) *DescribePackageDeductionsResponseBody {
	s.Deductions = v
	return s
}

func (s *DescribePackageDeductionsResponseBody) SetPageNum(v int32) *DescribePackageDeductionsResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribePackageDeductionsResponseBody) SetPageSize(v int32) *DescribePackageDeductionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePackageDeductionsResponseBody) SetRequestId(v string) *DescribePackageDeductionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePackageDeductionsResponseBody) SetTotalCount(v int64) *DescribePackageDeductionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePackageDeductionsResponseBody) SetTotalUsedCoreTime(v float32) *DescribePackageDeductionsResponseBody {
	s.TotalUsedCoreTime = &v
	return s
}

func (s *DescribePackageDeductionsResponseBody) SetTotalUsedTime(v int64) *DescribePackageDeductionsResponseBody {
	s.TotalUsedTime = &v
	return s
}

func (s *DescribePackageDeductionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePackageDeductionsResponseBodyDeductions struct {
	// example:
	//
	// 4
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// ecd-6wye9optu0kag****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// example:
	//
	// DemoComputer
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// example:
	//
	// eds.enterprise_office.4c8g
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// example:
	//
	// 2024-07-31T03:00Z
	EndTime           *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GroupResourceType *string `json:"GroupResourceType,omitempty" xml:"GroupResourceType,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Deleted
	InstanceState *string `json:"InstanceState,omitempty" xml:"InstanceState,omitempty"`
	InstanceType  *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 8192
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SessionId    *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 2024-07-31T02:00Z
	StaTime *string `json:"StaTime,omitempty" xml:"StaTime,omitempty"`
	// example:
	//
	// 4.0
	UsedCoreTime *float32 `json:"UsedCoreTime,omitempty" xml:"UsedCoreTime,omitempty"`
	// example:
	//
	// 3600
	UsedTime          *int64 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	UsedTimeWithScale *int64 `json:"UsedTimeWithScale,omitempty" xml:"UsedTimeWithScale,omitempty"`
}

func (s DescribePackageDeductionsResponseBodyDeductions) String() string {
	return dara.Prettify(s)
}

func (s DescribePackageDeductionsResponseBodyDeductions) GoString() string {
	return s.String()
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetDesktopType() *string {
	return s.DesktopType
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetGroupResourceType() *string {
	return s.GroupResourceType
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetInstanceState() *string {
	return s.InstanceState
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetMemory() *int64 {
	return s.Memory
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetOsType() *string {
	return s.OsType
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetStaTime() *string {
	return s.StaTime
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetUsedCoreTime() *float32 {
	return s.UsedCoreTime
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetUsedTime() *int64 {
	return s.UsedTime
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetUsedTimeWithScale() *int64 {
	return s.UsedTimeWithScale
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetCpu(v int32) *DescribePackageDeductionsResponseBodyDeductions {
	s.Cpu = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetDesktopId(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.DesktopId = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetDesktopName(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.DesktopName = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetDesktopType(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.DesktopType = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetEndTime(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.EndTime = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetGroupResourceType(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.GroupResourceType = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetInstanceId(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.InstanceId = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetInstanceState(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.InstanceState = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetInstanceType(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.InstanceType = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetMemory(v int64) *DescribePackageDeductionsResponseBodyDeductions {
	s.Memory = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetOsType(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.OsType = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetRegionId(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.RegionId = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetResourceType(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.ResourceType = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetSessionId(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.SessionId = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetStaTime(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.StaTime = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetUsedCoreTime(v float32) *DescribePackageDeductionsResponseBodyDeductions {
	s.UsedCoreTime = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetUsedTime(v int64) *DescribePackageDeductionsResponseBodyDeductions {
	s.UsedTime = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetUsedTimeWithScale(v int64) *DescribePackageDeductionsResponseBodyDeductions {
	s.UsedTimeWithScale = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) Validate() error {
	return dara.Validate(s)
}
