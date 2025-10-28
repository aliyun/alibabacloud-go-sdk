// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScalingRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetScalingRulesResponseBody
	GetCode() *int32
	SetData(v *GetScalingRulesResponseBodyData) *GetScalingRulesResponseBody
	GetData() *GetScalingRulesResponseBodyData
	SetMessage(v string) *GetScalingRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetScalingRulesResponseBody
	GetRequestId() *string
	SetUpdateTime(v int64) *GetScalingRulesResponseBody
	GetUpdateTime() *int64
}

type GetScalingRulesResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data that is returned.
	Data *GetScalingRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D16979DC-4D42-***********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when the scaling rule was last updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1574251601785
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetScalingRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScalingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *GetScalingRulesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetScalingRulesResponseBody) GetData() *GetScalingRulesResponseBodyData {
	return s.Data
}

func (s *GetScalingRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetScalingRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScalingRulesResponseBody) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetScalingRulesResponseBody) SetCode(v int32) *GetScalingRulesResponseBody {
	s.Code = &v
	return s
}

func (s *GetScalingRulesResponseBody) SetData(v *GetScalingRulesResponseBodyData) *GetScalingRulesResponseBody {
	s.Data = v
	return s
}

func (s *GetScalingRulesResponseBody) SetMessage(v string) *GetScalingRulesResponseBody {
	s.Message = &v
	return s
}

func (s *GetScalingRulesResponseBody) SetRequestId(v string) *GetScalingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScalingRulesResponseBody) SetUpdateTime(v int64) *GetScalingRulesResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetScalingRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScalingRulesResponseBodyData struct {
	// The type of the cluster. Valid values:
	//
	// 	- 0: regular Docker cluster
	//
	// 	- 1: Swarm cluster (deprecated)
	//
	// 	- 2: Elastic Compute Service (ECS) cluster
	//
	// 	- 3: self-managed Kubernetes cluster in EDAS
	//
	// 	- 4: cluster in which Pandora automatically registers applications
	//
	// 	- 5: Container Service for Kubernetes (ACK) clusters
	//
	// example:
	//
	// 2
	ClusterType *int32 `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The overcommit ratio supported by a Docker cluster. Valid values:
	//
	// 	- 1: 1:1, which means that resources are not overcommitted.
	//
	// 	- 2: 1:2, which means that resources are overcommitted by 1:2.
	//
	// 	- 4: 1:4, which means that resources are overcommitted by 1:4.
	//
	// 	- 8: 1:8, which means that resources are overcommitted by 1:8.
	//
	// example:
	//
	// 1
	OversoldFactor *int32 `json:"OversoldFactor,omitempty" xml:"OversoldFactor,omitempty"`
	// The array data of the scaling rule.
	RuleList *GetScalingRulesResponseBodyDataRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Struct"`
	// The time when the scaling rule was last updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1574251601785
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-wz9b246z******
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetScalingRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetScalingRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetScalingRulesResponseBodyData) GetClusterType() *int32 {
	return s.ClusterType
}

func (s *GetScalingRulesResponseBodyData) GetOversoldFactor() *int32 {
	return s.OversoldFactor
}

func (s *GetScalingRulesResponseBodyData) GetRuleList() *GetScalingRulesResponseBodyDataRuleList {
	return s.RuleList
}

func (s *GetScalingRulesResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetScalingRulesResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *GetScalingRulesResponseBodyData) SetClusterType(v int32) *GetScalingRulesResponseBodyData {
	s.ClusterType = &v
	return s
}

func (s *GetScalingRulesResponseBodyData) SetOversoldFactor(v int32) *GetScalingRulesResponseBodyData {
	s.OversoldFactor = &v
	return s
}

func (s *GetScalingRulesResponseBodyData) SetRuleList(v *GetScalingRulesResponseBodyDataRuleList) *GetScalingRulesResponseBodyData {
	s.RuleList = v
	return s
}

func (s *GetScalingRulesResponseBodyData) SetUpdateTime(v int64) *GetScalingRulesResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetScalingRulesResponseBodyData) SetVpcId(v string) *GetScalingRulesResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *GetScalingRulesResponseBodyData) Validate() error {
	if s.RuleList != nil {
		if err := s.RuleList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScalingRulesResponseBodyDataRuleList struct {
	Rule []*GetScalingRulesResponseBodyDataRuleListRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s GetScalingRulesResponseBodyDataRuleList) String() string {
	return dara.Prettify(s)
}

func (s GetScalingRulesResponseBodyDataRuleList) GoString() string {
	return s.String()
}

func (s *GetScalingRulesResponseBodyDataRuleList) GetRule() []*GetScalingRulesResponseBodyDataRuleListRule {
	return s.Rule
}

func (s *GetScalingRulesResponseBodyDataRuleList) SetRule(v []*GetScalingRulesResponseBodyDataRuleListRule) *GetScalingRulesResponseBodyDataRuleList {
	s.Rule = v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleList) Validate() error {
	if s.Rule != nil {
		for _, item := range s.Rule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetScalingRulesResponseBodyDataRuleListRule struct {
	// The ID of the application.
	//
	// example:
	//
	// 33e39be9-3e5f-*********
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The relationship among the conditions that trigger the scaling rule.
	//
	// 	- OR: one of the conditions
	//
	// 	- AND: all conditions
	//
	// example:
	//
	// OR
	Cond *string `json:"Cond,omitempty" xml:"Cond,omitempty"`
	// The minimum CPU utilization that triggers the scaling rule.
	//
	// example:
	//
	// 1
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The time when the scaling rule was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1574251601801
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The duration of the scaling rule. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1574251601
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Indicates whether scale-ins or scale-outs are allowed. Valid values:
	//
	// 	- true: Scale-ins or scale-outs are allowed.
	//
	// 	- false: Scale-ins or scale-outs are disallowed.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The ID of the instance group to which the application is deployed.
	//
	// example:
	//
	// d8bb9d60-91b5-4cdf-****-************
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The maximum number of instances in the group when a scale-out is performed, or the minimum number of instances in the group when a scale-in is performed.
	//
	// example:
	//
	// 2
	InstNum *int32 `json:"InstNum,omitempty" xml:"InstNum,omitempty"`
	// The system load that triggers the scaling rule. The system load is evaluated based on the number of processes that are being executed by CPUs and the number of processes that wait to be executed by CPUs.
	//
	// example:
	//
	// 1
	LoadNum *int32 `json:"LoadNum,omitempty" xml:"LoadNum,omitempty"`
	// The type of the metric.
	//
	// example:
	//
	// HSF
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The type of the scaling rule. Valid values:
	//
	// 	- SCALE_IN: scale-in rules
	//
	// 	- SCALE_OUT: scale-out rules
	//
	// example:
	//
	// SCALE_OUT
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The policy of auto scaling across multiple zones. Valid values:
	//
	// 	- PRIORITY: The vSwitch that is first selected has the highest priority.
	//
	// 	- BALANCE: This policy evenly distributes instances across zones in which the vSwitches reside.
	//
	// example:
	//
	// PRIORITY
	MultiAzPolicy *string `json:"MultiAzPolicy,omitempty" xml:"MultiAzPolicy,omitempty"`
	// The source of the instance that you want to add during a scale-out. Valid values:
	//
	// 	- NEW: Elastic resources are used.
	//
	// 	- AVAILABLE: The existing resources are used.
	//
	// 	- AVAILABLE_FIRST: The existing resources are used first.
	//
	// example:
	//
	// AVAILABLE
	ResourceFrom *string `json:"ResourceFrom,omitempty" xml:"ResourceFrom,omitempty"`
	// The service latency that triggers the scaling rule. Unit: milliseconds.
	//
	// example:
	//
	// 1
	Rt *int32 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	// The ID of the specification.
	//
	// example:
	//
	// 03f493c0-xxxx-xxxx-xxxx-12e85cadeb41
	SpecId *string `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
	// The number of instances that are added during each scale-out or removed during each scale-in.
	//
	// example:
	//
	// 1
	Step *int32 `json:"Step,omitempty" xml:"Step,omitempty"`
	// The ID of the launch template.
	//
	// example:
	//
	// lt-bp1xxxxn73pxxxxf83l
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The version of the launch template.
	//
	// example:
	//
	// 1143542
	TemplateVersion *int32 `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The time when the scaling rule was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1574251601785
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The IDs of the vSwitches. The IDs of multiple vSwitches are separated by commas (,).
	//
	// example:
	//
	// vsw-mxxxxkxxxx4xxxxwbionj
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-wz9b246z******
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetScalingRulesResponseBodyDataRuleListRule) String() string {
	return dara.Prettify(s)
}

func (s GetScalingRulesResponseBodyDataRuleListRule) GoString() string {
	return s.String()
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) GetAppId() *string {
	return s.AppId
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) GetCond() *string {
	return s.Cond
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) GetCpu() *int32 {
	return s.Cpu
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) GetDuration() *int32 {
	return s.Duration
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) GetEnable() *bool {
	return s.Enable
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) GetGroupId() *string {
	return s.GroupId
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) GetInstNum() *int32 {
	return s.InstNum
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) GetLoadNum() *int32 {
	return s.LoadNum
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) GetMetricType() *string {
	return s.MetricType
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) GetMode() *string {
	return s.Mode
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) GetMultiAzPolicy() *string {
	return s.MultiAzPolicy
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) GetResourceFrom() *string {
	return s.ResourceFrom
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) GetRt() *int32 {
	return s.Rt
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) GetSpecId() *string {
	return s.SpecId
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) GetStep() *int32 {
	return s.Step
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) GetTemplateVersion() *int32 {
	return s.TemplateVersion
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) GetVpcId() *string {
	return s.VpcId
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetAppId(v string) *GetScalingRulesResponseBodyDataRuleListRule {
	s.AppId = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetCond(v string) *GetScalingRulesResponseBodyDataRuleListRule {
	s.Cond = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetCpu(v int32) *GetScalingRulesResponseBodyDataRuleListRule {
	s.Cpu = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetCreateTime(v int64) *GetScalingRulesResponseBodyDataRuleListRule {
	s.CreateTime = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetDuration(v int32) *GetScalingRulesResponseBodyDataRuleListRule {
	s.Duration = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetEnable(v bool) *GetScalingRulesResponseBodyDataRuleListRule {
	s.Enable = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetGroupId(v string) *GetScalingRulesResponseBodyDataRuleListRule {
	s.GroupId = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetInstNum(v int32) *GetScalingRulesResponseBodyDataRuleListRule {
	s.InstNum = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetLoadNum(v int32) *GetScalingRulesResponseBodyDataRuleListRule {
	s.LoadNum = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetMetricType(v string) *GetScalingRulesResponseBodyDataRuleListRule {
	s.MetricType = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetMode(v string) *GetScalingRulesResponseBodyDataRuleListRule {
	s.Mode = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetMultiAzPolicy(v string) *GetScalingRulesResponseBodyDataRuleListRule {
	s.MultiAzPolicy = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetResourceFrom(v string) *GetScalingRulesResponseBodyDataRuleListRule {
	s.ResourceFrom = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetRt(v int32) *GetScalingRulesResponseBodyDataRuleListRule {
	s.Rt = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetSpecId(v string) *GetScalingRulesResponseBodyDataRuleListRule {
	s.SpecId = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetStep(v int32) *GetScalingRulesResponseBodyDataRuleListRule {
	s.Step = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetTemplateId(v string) *GetScalingRulesResponseBodyDataRuleListRule {
	s.TemplateId = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetTemplateVersion(v int32) *GetScalingRulesResponseBodyDataRuleListRule {
	s.TemplateVersion = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetUpdateTime(v int64) *GetScalingRulesResponseBodyDataRuleListRule {
	s.UpdateTime = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetVSwitchIds(v string) *GetScalingRulesResponseBodyDataRuleListRule {
	s.VSwitchIds = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetVpcId(v string) *GetScalingRulesResponseBodyDataRuleListRule {
	s.VpcId = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) Validate() error {
	return dara.Validate(s)
}
