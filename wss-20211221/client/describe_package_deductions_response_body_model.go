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
	SetTotalUsedTimeDecimal(v string) *DescribePackageDeductionsResponseBody
	GetTotalUsedTimeDecimal() *string
}

type DescribePackageDeductionsResponseBody struct {
	// The list of deduction details.
	Deductions []*DescribePackageDeductionsResponseBodyDeductions `json:"Deductions,omitempty" xml:"Deductions,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page in the Settings for paging during a paged query.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 833C4D2C-09C7-5CE6-8159-06758B964970
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of core-hour package deduction details queried.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total deducted core-hours.
	//
	// example:
	//
	// 10.1
	TotalUsedCoreTime *float32 `json:"TotalUsedCoreTime,omitempty" xml:"TotalUsedCoreTime,omitempty"`
	// The total usage duration.
	//
	// example:
	//
	// 100000
	TotalUsedTime *int64 `json:"TotalUsedTime,omitempty" xml:"TotalUsedTime,omitempty"`
	// The total usage calculated by using a precision factor, rounded to four decimal places. This field uses the same metering criteria as UsedTimeDecimal in the details.
	//
	// example:
	//
	// 10.0000
	TotalUsedTimeDecimal *string `json:"TotalUsedTimeDecimal,omitempty" xml:"TotalUsedTimeDecimal,omitempty"`
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

func (s *DescribePackageDeductionsResponseBody) GetTotalUsedTimeDecimal() *string {
	return s.TotalUsedTimeDecimal
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

func (s *DescribePackageDeductionsResponseBody) SetTotalUsedTimeDecimal(v string) *DescribePackageDeductionsResponseBody {
	s.TotalUsedTimeDecimal = &v
	return s
}

func (s *DescribePackageDeductionsResponseBody) Validate() error {
	if s.Deductions != nil {
		for _, item := range s.Deductions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePackageDeductionsResponseBodyDeductions struct {
	// The agent type.
	//
	// example:
	//
	// OpenClaw
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 4
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The request ID.
	//
	// example:
	//
	// r-xxxx
	DeductionRequestId *string `json:"DeductionRequestId,omitempty" xml:"DeductionRequestId,omitempty"`
	// The deduction types.
	//
	// example:
	//
	// KNOWLEDGE_SPACE/CONVESATION
	DeductionType *string `json:"DeductionType,omitempty" xml:"DeductionType,omitempty"`
	// The Cloud Desktop ID.
	//
	// example:
	//
	// ecd-6wye9optu0kag****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The name of the cloud desktop.
	//
	// example:
	//
	// DemoComputer
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// The specifications of the cloud desktop.
	//
	// example:
	//
	// eds.enterprise_office.4c8g
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// The end time of the deduction record in UTC. The format is yyyy-MM-ddTHH:mm:ssZ. For records aggregated by hour, the format can be yyyy-MM-ddTHH:mmZ with seconds omitted.
	//
	// example:
	//
	// 2024-07-31T03:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The group resource type to which the deduction record belongs.
	//
	// example:
	//
	// GROUP_CORE_PACKAGE
	GroupResourceType *string `json:"GroupResourceType,omitempty" xml:"GroupResourceType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// A specific instance ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the core-hour package.
	//
	// example:
	//
	// Deleted
	InstanceState *string `json:"InstanceState,omitempty" xml:"InstanceState,omitempty"`
	// The instance specifications for the deduction.
	//
	// example:
	//
	// ec.agentbay-pc
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The memory. Unit: MB.
	//
	// example:
	//
	// 8192
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The model.
	//
	// example:
	//
	// qwen-3.7
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The operating system type of the cloud desktop.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The ID of the credit package.
	//
	// example:
	//
	// crp-xxx
	PackageId *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	// The region ID of the cloud desktop.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the first resource in the deduction item.
	//
	// example:
	//
	// ecd-xxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// Desktop
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the session associated with the deduction record.
	//
	// example:
	//
	// ai-123456****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The start time of the deduction record in UTC. The format is yyyy-MM-ddTHH:mm:ssZ. For records aggregated by hour, the format can be yyyy-MM-ddTHH:mmZ with seconds omitted.
	//
	// example:
	//
	// 2024-07-31T02:00:00Z
	StaTime *string `json:"StaTime,omitempty" xml:"StaTime,omitempty"`
	// The summary.
	//
	// example:
	//
	// Summary of the AI query
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// The number of cores deducted.
	//
	// example:
	//
	// 4.0
	UsedCoreTime *float32 `json:"UsedCoreTime,omitempty" xml:"UsedCoreTime,omitempty"`
	// The usage duration. Unit: seconds.
	//
	// example:
	//
	// 3600
	UsedTime *int64 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The usage calculated from UsedTime by using a precision factor, rounded to four decimal places. This field is returned only for applicable resource types and must be read based on the metering criteria of the corresponding resource.
	//
	// example:
	//
	// 1.0000
	UsedTimeDecimal *string `json:"UsedTimeDecimal,omitempty" xml:"UsedTimeDecimal,omitempty"`
	// The usage that includes the deduction coefficient. In scenarios where deduction is based on duration, the value is measured in seconds and returned as an integer.
	//
	// example:
	//
	// 3600
	UsedTimeWithScale *int64 `json:"UsedTimeWithScale,omitempty" xml:"UsedTimeWithScale,omitempty"`
}

func (s DescribePackageDeductionsResponseBodyDeductions) String() string {
	return dara.Prettify(s)
}

func (s DescribePackageDeductionsResponseBodyDeductions) GoString() string {
	return s.String()
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetAgentType() *string {
	return s.AgentType
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetDeductionRequestId() *string {
	return s.DeductionRequestId
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetDeductionType() *string {
	return s.DeductionType
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

func (s *DescribePackageDeductionsResponseBodyDeductions) GetModel() *string {
	return s.Model
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetOsType() *string {
	return s.OsType
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetPackageId() *string {
	return s.PackageId
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetResourceId() *string {
	return s.ResourceId
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

func (s *DescribePackageDeductionsResponseBodyDeductions) GetSummary() *string {
	return s.Summary
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetUsedCoreTime() *float32 {
	return s.UsedCoreTime
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetUsedTime() *int64 {
	return s.UsedTime
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetUsedTimeDecimal() *string {
	return s.UsedTimeDecimal
}

func (s *DescribePackageDeductionsResponseBodyDeductions) GetUsedTimeWithScale() *int64 {
	return s.UsedTimeWithScale
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetAgentType(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.AgentType = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetCpu(v int32) *DescribePackageDeductionsResponseBodyDeductions {
	s.Cpu = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetDeductionRequestId(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.DeductionRequestId = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetDeductionType(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.DeductionType = &v
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

func (s *DescribePackageDeductionsResponseBodyDeductions) SetModel(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.Model = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetOsType(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.OsType = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetPackageId(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.PackageId = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetRegionId(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.RegionId = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetResourceId(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.ResourceId = &v
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

func (s *DescribePackageDeductionsResponseBodyDeductions) SetSummary(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.Summary = &v
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

func (s *DescribePackageDeductionsResponseBodyDeductions) SetUsedTimeDecimal(v string) *DescribePackageDeductionsResponseBodyDeductions {
	s.UsedTimeDecimal = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) SetUsedTimeWithScale(v int64) *DescribePackageDeductionsResponseBodyDeductions {
	s.UsedTimeWithScale = &v
	return s
}

func (s *DescribePackageDeductionsResponseBodyDeductions) Validate() error {
	return dara.Validate(s)
}
