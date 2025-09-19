// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeStrategyResponseBody
	GetRequestId() *string
	SetStrategies(v []*DescribeStrategyResponseBodyStrategies) *DescribeStrategyResponseBody
	GetStrategies() []*DescribeStrategyResponseBodyStrategies
}

type DescribeStrategyResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 75C127E6-76CD-59A7-B6E4-1CBBDC98F2EB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the baseline check policies.
	Strategies []*DescribeStrategyResponseBodyStrategies `json:"Strategies,omitempty" xml:"Strategies,omitempty" type:"Repeated"`
}

func (s DescribeStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStrategyResponseBody) GetStrategies() []*DescribeStrategyResponseBodyStrategies {
	return s.Strategies
}

func (s *DescribeStrategyResponseBody) SetRequestId(v string) *DescribeStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStrategyResponseBody) SetStrategies(v []*DescribeStrategyResponseBodyStrategies) *DescribeStrategyResponseBody {
	s.Strategies = v
	return s
}

func (s *DescribeStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeStrategyResponseBodyStrategies struct {
	// The details of the assets to which the baseline check policy is applied.
	ConfigTargets []*DescribeStrategyResponseBodyStrategiesConfigTargets `json:"ConfigTargets,omitempty" xml:"ConfigTargets,omitempty" type:"Repeated"`
	// The type of the baseline check policy. Valid values:
	//
	// 	- **common**
	//
	// 	- **custom**
	//
	// example:
	//
	// custom
	CustomType *string `json:"CustomType,omitempty" xml:"CustomType,omitempty"`
	// The cycle of the baseline check. Valid values:
	//
	// 	- **1**: every 2 days
	//
	// 	- **3**: every 4 days
	//
	// 	- **7**: every 8 days
	//
	// 	- 30: every 31 days
	//
	// example:
	//
	// 1
	CycleDays *int32 `json:"CycleDays,omitempty" xml:"CycleDays,omitempty"`
	// The time when the baseline check starts. Valid values:
	//
	// 	- **0**: The baseline check starts within the time range from 00:00 to 06:00.
	//
	// 	- **6**: The baseline check starts within the time range from 06:00 to 12:00.
	//
	// 	- **12**: The baseline check starts within the time range from 12:00 to 18:00.
	//
	// 	- **18**: The baseline check starts within the time range from 18:00 to 24:00.
	//
	// example:
	//
	// 0
	CycleStartTime *int32 `json:"CycleStartTime,omitempty" xml:"CycleStartTime,omitempty"`
	// The number of the assets to which the baseline check policy is applied.
	//
	// example:
	//
	// 50
	EcsCount *int32 `json:"EcsCount,omitempty" xml:"EcsCount,omitempty"`
	// The end time of the baseline check policy.
	//
	// example:
	//
	// 03:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The status of the baseline check policy. Valid values:
	//
	// 	- **1**: not executed
	//
	// 	- **2**: executing
	//
	// example:
	//
	// 1
	ExecStatus    *int32  `json:"ExecStatus,omitempty" xml:"ExecStatus,omitempty"`
	ExecutionType *string `json:"ExecutionType,omitempty" xml:"ExecutionType,omitempty"`
	// The ID of the baseline check policy.
	//
	// example:
	//
	// 8164248
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the baseline check policy.
	//
	// example:
	//
	// text2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The proportion of risky baselines in the baseline check result.
	//
	// example:
	//
	// 0
	PassRate *int32 `json:"PassRate,omitempty" xml:"PassRate,omitempty"`
	// The progress of the baseline check by using the baseline. This parameter is returned only if the value of the ExecStatus parameter is 2.
	//
	// example:
	//
	// 50%
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of the assets on which the baseline check is complete.
	//
	// example:
	//
	// 20
	ProcessRate *int32 `json:"ProcessRate,omitempty" xml:"ProcessRate,omitempty"`
	// The number of baseline check items in the baseline check policy.
	//
	// example:
	//
	// 23
	RiskCount *int32 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// The start time of the baseline check policy.
	//
	// example:
	//
	// 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The source type of the baseline check policy. Valid values:
	//
	// 	- **1**: indicates a built-in policy provided and performed by Security Center by default.
	//
	// 	- **2**: indicates a user-defined policy. It can be a standard or custom baseline check policy.
	//
	// example:
	//
	// 2
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the baseline check policy was last modified.
	//
	// example:
	//
	// 2025-01-07 10:46:43
	UserModifyTime *int64 `json:"UserModifyTime,omitempty" xml:"UserModifyTime,omitempty"`
}

func (s DescribeStrategyResponseBodyStrategies) String() string {
	return dara.Prettify(s)
}

func (s DescribeStrategyResponseBodyStrategies) GoString() string {
	return s.String()
}

func (s *DescribeStrategyResponseBodyStrategies) GetConfigTargets() []*DescribeStrategyResponseBodyStrategiesConfigTargets {
	return s.ConfigTargets
}

func (s *DescribeStrategyResponseBodyStrategies) GetCustomType() *string {
	return s.CustomType
}

func (s *DescribeStrategyResponseBodyStrategies) GetCycleDays() *int32 {
	return s.CycleDays
}

func (s *DescribeStrategyResponseBodyStrategies) GetCycleStartTime() *int32 {
	return s.CycleStartTime
}

func (s *DescribeStrategyResponseBodyStrategies) GetEcsCount() *int32 {
	return s.EcsCount
}

func (s *DescribeStrategyResponseBodyStrategies) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeStrategyResponseBodyStrategies) GetExecStatus() *int32 {
	return s.ExecStatus
}

func (s *DescribeStrategyResponseBodyStrategies) GetExecutionType() *string {
	return s.ExecutionType
}

func (s *DescribeStrategyResponseBodyStrategies) GetId() *int32 {
	return s.Id
}

func (s *DescribeStrategyResponseBodyStrategies) GetName() *string {
	return s.Name
}

func (s *DescribeStrategyResponseBodyStrategies) GetPassRate() *int32 {
	return s.PassRate
}

func (s *DescribeStrategyResponseBodyStrategies) GetPercent() *string {
	return s.Percent
}

func (s *DescribeStrategyResponseBodyStrategies) GetProcessRate() *int32 {
	return s.ProcessRate
}

func (s *DescribeStrategyResponseBodyStrategies) GetRiskCount() *int32 {
	return s.RiskCount
}

func (s *DescribeStrategyResponseBodyStrategies) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeStrategyResponseBodyStrategies) GetType() *int32 {
	return s.Type
}

func (s *DescribeStrategyResponseBodyStrategies) GetUserModifyTime() *int64 {
	return s.UserModifyTime
}

func (s *DescribeStrategyResponseBodyStrategies) SetConfigTargets(v []*DescribeStrategyResponseBodyStrategiesConfigTargets) *DescribeStrategyResponseBodyStrategies {
	s.ConfigTargets = v
	return s
}

func (s *DescribeStrategyResponseBodyStrategies) SetCustomType(v string) *DescribeStrategyResponseBodyStrategies {
	s.CustomType = &v
	return s
}

func (s *DescribeStrategyResponseBodyStrategies) SetCycleDays(v int32) *DescribeStrategyResponseBodyStrategies {
	s.CycleDays = &v
	return s
}

func (s *DescribeStrategyResponseBodyStrategies) SetCycleStartTime(v int32) *DescribeStrategyResponseBodyStrategies {
	s.CycleStartTime = &v
	return s
}

func (s *DescribeStrategyResponseBodyStrategies) SetEcsCount(v int32) *DescribeStrategyResponseBodyStrategies {
	s.EcsCount = &v
	return s
}

func (s *DescribeStrategyResponseBodyStrategies) SetEndTime(v string) *DescribeStrategyResponseBodyStrategies {
	s.EndTime = &v
	return s
}

func (s *DescribeStrategyResponseBodyStrategies) SetExecStatus(v int32) *DescribeStrategyResponseBodyStrategies {
	s.ExecStatus = &v
	return s
}

func (s *DescribeStrategyResponseBodyStrategies) SetExecutionType(v string) *DescribeStrategyResponseBodyStrategies {
	s.ExecutionType = &v
	return s
}

func (s *DescribeStrategyResponseBodyStrategies) SetId(v int32) *DescribeStrategyResponseBodyStrategies {
	s.Id = &v
	return s
}

func (s *DescribeStrategyResponseBodyStrategies) SetName(v string) *DescribeStrategyResponseBodyStrategies {
	s.Name = &v
	return s
}

func (s *DescribeStrategyResponseBodyStrategies) SetPassRate(v int32) *DescribeStrategyResponseBodyStrategies {
	s.PassRate = &v
	return s
}

func (s *DescribeStrategyResponseBodyStrategies) SetPercent(v string) *DescribeStrategyResponseBodyStrategies {
	s.Percent = &v
	return s
}

func (s *DescribeStrategyResponseBodyStrategies) SetProcessRate(v int32) *DescribeStrategyResponseBodyStrategies {
	s.ProcessRate = &v
	return s
}

func (s *DescribeStrategyResponseBodyStrategies) SetRiskCount(v int32) *DescribeStrategyResponseBodyStrategies {
	s.RiskCount = &v
	return s
}

func (s *DescribeStrategyResponseBodyStrategies) SetStartTime(v string) *DescribeStrategyResponseBodyStrategies {
	s.StartTime = &v
	return s
}

func (s *DescribeStrategyResponseBodyStrategies) SetType(v int32) *DescribeStrategyResponseBodyStrategies {
	s.Type = &v
	return s
}

func (s *DescribeStrategyResponseBodyStrategies) SetUserModifyTime(v int64) *DescribeStrategyResponseBodyStrategies {
	s.UserModifyTime = &v
	return s
}

func (s *DescribeStrategyResponseBodyStrategies) Validate() error {
	return dara.Validate(s)
}

type DescribeStrategyResponseBodyStrategiesConfigTargets struct {
	// Indicates whether the baseline check policy is applied to the asset group. Valid values:
	//
	// 	- **add**: The baseline check policy is applied to the asset group.
	//
	// 	- **del**: the baseline check policy is not applied to the asset group.
	//
	// example:
	//
	// add
	Flag *string `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// The asset group ID or UUID of the asset to which the baseline check policy is applied.
	//
	// example:
	//
	// 10099713
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The condition by which the baseline check policy is applied to the asset. Valid values:
	//
	// 	- **groupId**: the ID of the asset group
	//
	// 	- **uuid**: the UUID of the asset
	//
	// example:
	//
	// groupId
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s DescribeStrategyResponseBodyStrategiesConfigTargets) String() string {
	return dara.Prettify(s)
}

func (s DescribeStrategyResponseBodyStrategiesConfigTargets) GoString() string {
	return s.String()
}

func (s *DescribeStrategyResponseBodyStrategiesConfigTargets) GetFlag() *string {
	return s.Flag
}

func (s *DescribeStrategyResponseBodyStrategiesConfigTargets) GetTarget() *string {
	return s.Target
}

func (s *DescribeStrategyResponseBodyStrategiesConfigTargets) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeStrategyResponseBodyStrategiesConfigTargets) SetFlag(v string) *DescribeStrategyResponseBodyStrategiesConfigTargets {
	s.Flag = &v
	return s
}

func (s *DescribeStrategyResponseBodyStrategiesConfigTargets) SetTarget(v string) *DescribeStrategyResponseBodyStrategiesConfigTargets {
	s.Target = &v
	return s
}

func (s *DescribeStrategyResponseBodyStrategiesConfigTargets) SetTargetType(v string) *DescribeStrategyResponseBodyStrategiesConfigTargets {
	s.TargetType = &v
	return s
}

func (s *DescribeStrategyResponseBodyStrategiesConfigTargets) Validate() error {
	return dara.Validate(s)
}
