// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScalingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptEULA(v bool) *ModifyScalingRuleRequest
	GetAcceptEULA() *bool
	SetAppId(v string) *ModifyScalingRuleRequest
	GetAppId() *string
	SetGroupId(v string) *ModifyScalingRuleRequest
	GetGroupId() *string
	SetInCondition(v string) *ModifyScalingRuleRequest
	GetInCondition() *string
	SetInCpu(v int32) *ModifyScalingRuleRequest
	GetInCpu() *int32
	SetInDuration(v int32) *ModifyScalingRuleRequest
	GetInDuration() *int32
	SetInEnable(v bool) *ModifyScalingRuleRequest
	GetInEnable() *bool
	SetInInstanceNum(v int32) *ModifyScalingRuleRequest
	GetInInstanceNum() *int32
	SetInLoad(v int32) *ModifyScalingRuleRequest
	GetInLoad() *int32
	SetInRT(v int32) *ModifyScalingRuleRequest
	GetInRT() *int32
	SetInStep(v int32) *ModifyScalingRuleRequest
	GetInStep() *int32
	SetKeyPairName(v string) *ModifyScalingRuleRequest
	GetKeyPairName() *string
	SetMultiAzPolicy(v string) *ModifyScalingRuleRequest
	GetMultiAzPolicy() *string
	SetOutCPU(v int32) *ModifyScalingRuleRequest
	GetOutCPU() *int32
	SetOutCondition(v string) *ModifyScalingRuleRequest
	GetOutCondition() *string
	SetOutDuration(v int32) *ModifyScalingRuleRequest
	GetOutDuration() *int32
	SetOutEnable(v bool) *ModifyScalingRuleRequest
	GetOutEnable() *bool
	SetOutInstanceNum(v int32) *ModifyScalingRuleRequest
	GetOutInstanceNum() *int32
	SetOutLoad(v int32) *ModifyScalingRuleRequest
	GetOutLoad() *int32
	SetOutRT(v int32) *ModifyScalingRuleRequest
	GetOutRT() *int32
	SetOutStep(v int32) *ModifyScalingRuleRequest
	GetOutStep() *int32
	SetPassword(v string) *ModifyScalingRuleRequest
	GetPassword() *string
	SetResourceFrom(v string) *ModifyScalingRuleRequest
	GetResourceFrom() *string
	SetScalingPolicy(v string) *ModifyScalingRuleRequest
	GetScalingPolicy() *string
	SetTemplateId(v string) *ModifyScalingRuleRequest
	GetTemplateId() *string
	SetTemplateInstanceId(v string) *ModifyScalingRuleRequest
	GetTemplateInstanceId() *string
	SetTemplateInstanceName(v string) *ModifyScalingRuleRequest
	GetTemplateInstanceName() *string
	SetTemplateVersion(v int32) *ModifyScalingRuleRequest
	GetTemplateVersion() *int32
	SetVSwitchIds(v string) *ModifyScalingRuleRequest
	GetVSwitchIds() *string
	SetVpcId(v string) *ModifyScalingRuleRequest
	GetVpcId() *string
}

type ModifyScalingRuleRequest struct {
	// Set the value to true if scale-outs are allowed.
	//
	// example:
	//
	// true
	AcceptEULA *bool `json:"AcceptEULA,omitempty" xml:"AcceptEULA,omitempty"`
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 74ee****-db65-4322-a1f6-bcb60e5b****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the instance group to which the application is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8123db90-880f-486f-****-************
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The relationship among the conditions that trigger a scale-in.
	//
	// 	- OR: one of the conditions
	//
	// 	- AND: all conditions
	//
	// example:
	//
	// "OR"
	InCondition *string `json:"InCondition,omitempty" xml:"InCondition,omitempty"`
	// The CPU utilization that triggers a scale-in.
	//
	// example:
	//
	// 50
	InCpu *int32 `json:"InCpu,omitempty" xml:"InCpu,omitempty"`
	// The duration in which the metric threshold is exceeded. Unit: minutes.
	//
	// example:
	//
	// 50
	InDuration *int32 `json:"InDuration,omitempty" xml:"InDuration,omitempty"`
	// Specifies whether to allow scale-ins.
	//
	// 	- true: allows scale-ins.
	//
	// 	- false: does not allow scale-ins.
	//
	// example:
	//
	// true
	InEnable *bool `json:"InEnable,omitempty" xml:"InEnable,omitempty"`
	// The minimum number of instances that must be retained in each group when a scale-in is performed.
	//
	// example:
	//
	// 3
	InInstanceNum *int32 `json:"InInstanceNum,omitempty" xml:"InInstanceNum,omitempty"`
	// The system load that triggers a scale-in.
	//
	// example:
	//
	// 50
	InLoad *int32 `json:"InLoad,omitempty" xml:"InLoad,omitempty"`
	// The minimum service latency that triggers a scale-in. The lower limit is 0. Unit: milliseconds.
	//
	// example:
	//
	// 50
	InRT *int32 `json:"InRT,omitempty" xml:"InRT,omitempty"`
	// The number of instances that are removed during each scale-in.
	//
	// example:
	//
	// 1
	InStep *int32 `json:"InStep,omitempty" xml:"InStep,omitempty"`
	// The key pair that is used to log on to the instance. This parameter takes effect only if you choose to create instances based on the specifications of an existing instance during a scale-out.
	//
	// example:
	//
	// "tdy218"
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The multi-zone scaling policy. Valid values:
	//
	// 	- PRIORITY: The vSwitch that is first selected has the highest priority.
	//
	// 	- BALANCE: This policy evenly distributes instances across zones in which the vSwitches reside.
	//
	// example:
	//
	// "PRIORITY"
	MultiAzPolicy *string `json:"MultiAzPolicy,omitempty" xml:"MultiAzPolicy,omitempty"`
	// The CPU utilization that triggers a scale-out.
	//
	// example:
	//
	// 50
	OutCPU *int32 `json:"OutCPU,omitempty" xml:"OutCPU,omitempty"`
	// The relationship among the conditions that trigger a scale-out.
	//
	// 	- OR: one of the conditions
	//
	// 	- AND: all conditions
	//
	// example:
	//
	// "OR"
	OutCondition *string `json:"OutCondition,omitempty" xml:"OutCondition,omitempty"`
	// The duration in which the metric threshold is exceeded. Unit: minutes.
	//
	// example:
	//
	// 50
	OutDuration *int32 `json:"OutDuration,omitempty" xml:"OutDuration,omitempty"`
	// Specifies whether to allow scale-outs.
	//
	// example:
	//
	// true
	OutEnable *bool `json:"OutEnable,omitempty" xml:"OutEnable,omitempty"`
	// The maximum number of instances in each group when a scale-out is performed.
	//
	// example:
	//
	// 10
	OutInstanceNum *int32 `json:"OutInstanceNum,omitempty" xml:"OutInstanceNum,omitempty"`
	// The system load that triggers a scale-out.
	//
	// example:
	//
	// 50
	OutLoad *int32 `json:"OutLoad,omitempty" xml:"OutLoad,omitempty"`
	// The minimum service latency that triggers a scale-out. The lower limit is 0. Unit: milliseconds.
	//
	// example:
	//
	// 0
	OutRT *int32 `json:"OutRT,omitempty" xml:"OutRT,omitempty"`
	// The number of instances that are added during each scale-out.
	//
	// example:
	//
	// 0
	OutStep *int32 `json:"OutStep,omitempty" xml:"OutStep,omitempty"`
	// The password that is used to log on to the instance. This parameter takes effect only if you choose to create instances based on the specifications of an existing instance during a scale-out.
	//
	// example:
	//
	// "Pwd*****"
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The source of the instance to be added during a scale-out. Valid values:
	//
	// 	- NEW: elastic resources
	//
	// 	- AVAILABLE: existing resources If you prefer existing resources to elastic resources, set this parameter to AVAILABLE_FIRST.
	//
	// If you set this parameter to NEW or AVAILABLE_FIRST, you must specify the auto-scaling parameters. If you set this parameter to NEW, instances are created based on a launch template or the specifications of an existing instance.
	//
	// example:
	//
	// "AVAILABLE"
	ResourceFrom *string `json:"ResourceFrom,omitempty" xml:"ResourceFrom,omitempty"`
	// The instance handling mode during a scale-in. Valid values:
	//
	// 	- release: When a scale-in is performed, instances that are no longer used are released.
	//
	// 	- recycle: When a scale-in is performed, instances that are no longer used are stopped and reclaimed.
	//
	// example:
	//
	// "release"
	ScalingPolicy *string `json:"ScalingPolicy,omitempty" xml:"ScalingPolicy,omitempty"`
	// The ID of the launch template that is used to create instances during a scale-out. This parameter takes effect only if you set the OutEnable parameter to true. This parameter takes precedence over the TemplateInstanceId parameter.
	//
	// example:
	//
	// "lt-wz9hkhn8wp*****"
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The ID of the instance whose specifications are used to create instances during a scale-out. This parameter is valid only when you set the OutEnable parameter to true.
	//
	// example:
	//
	// "1"
	TemplateInstanceId *string `json:"TemplateInstanceId,omitempty" xml:"TemplateInstanceId,omitempty"`
	// The name of the instance whose specifications are used to create instances during a scale-out. This parameter takes effect only if you specify the TemplateInstanceId parameter.
	//
	// example:
	//
	// "tpl-tdy218"
	TemplateInstanceName *string `json:"TemplateInstanceName,omitempty" xml:"TemplateInstanceName,omitempty"`
	// The version of the launch template that is used to create instances during a scale-out. This parameter takes effect only if you set the OutEnable parameter to true. To use the default template version, set this parameter to `-1`. Otherwise, set this parameter to the version that you want to use.
	//
	// example:
	//
	// -1
	TemplateVersion *int32 `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The IDs of the vSwitches that are associated with the VPC. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// "vsw-bp1ldxs3d4fd*****"
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The ID of the virtual private cloud (VPC) that is associated with the instances created based on a launch template or the specifications of an existing instance.
	//
	// example:
	//
	// "vpc-bp1j55oz3bg*****"
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ModifyScalingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyScalingRuleRequest) GetAcceptEULA() *bool {
	return s.AcceptEULA
}

func (s *ModifyScalingRuleRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyScalingRuleRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ModifyScalingRuleRequest) GetInCondition() *string {
	return s.InCondition
}

func (s *ModifyScalingRuleRequest) GetInCpu() *int32 {
	return s.InCpu
}

func (s *ModifyScalingRuleRequest) GetInDuration() *int32 {
	return s.InDuration
}

func (s *ModifyScalingRuleRequest) GetInEnable() *bool {
	return s.InEnable
}

func (s *ModifyScalingRuleRequest) GetInInstanceNum() *int32 {
	return s.InInstanceNum
}

func (s *ModifyScalingRuleRequest) GetInLoad() *int32 {
	return s.InLoad
}

func (s *ModifyScalingRuleRequest) GetInRT() *int32 {
	return s.InRT
}

func (s *ModifyScalingRuleRequest) GetInStep() *int32 {
	return s.InStep
}

func (s *ModifyScalingRuleRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *ModifyScalingRuleRequest) GetMultiAzPolicy() *string {
	return s.MultiAzPolicy
}

func (s *ModifyScalingRuleRequest) GetOutCPU() *int32 {
	return s.OutCPU
}

func (s *ModifyScalingRuleRequest) GetOutCondition() *string {
	return s.OutCondition
}

func (s *ModifyScalingRuleRequest) GetOutDuration() *int32 {
	return s.OutDuration
}

func (s *ModifyScalingRuleRequest) GetOutEnable() *bool {
	return s.OutEnable
}

func (s *ModifyScalingRuleRequest) GetOutInstanceNum() *int32 {
	return s.OutInstanceNum
}

func (s *ModifyScalingRuleRequest) GetOutLoad() *int32 {
	return s.OutLoad
}

func (s *ModifyScalingRuleRequest) GetOutRT() *int32 {
	return s.OutRT
}

func (s *ModifyScalingRuleRequest) GetOutStep() *int32 {
	return s.OutStep
}

func (s *ModifyScalingRuleRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyScalingRuleRequest) GetResourceFrom() *string {
	return s.ResourceFrom
}

func (s *ModifyScalingRuleRequest) GetScalingPolicy() *string {
	return s.ScalingPolicy
}

func (s *ModifyScalingRuleRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ModifyScalingRuleRequest) GetTemplateInstanceId() *string {
	return s.TemplateInstanceId
}

func (s *ModifyScalingRuleRequest) GetTemplateInstanceName() *string {
	return s.TemplateInstanceName
}

func (s *ModifyScalingRuleRequest) GetTemplateVersion() *int32 {
	return s.TemplateVersion
}

func (s *ModifyScalingRuleRequest) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *ModifyScalingRuleRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ModifyScalingRuleRequest) SetAcceptEULA(v bool) *ModifyScalingRuleRequest {
	s.AcceptEULA = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetAppId(v string) *ModifyScalingRuleRequest {
	s.AppId = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetGroupId(v string) *ModifyScalingRuleRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetInCondition(v string) *ModifyScalingRuleRequest {
	s.InCondition = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetInCpu(v int32) *ModifyScalingRuleRequest {
	s.InCpu = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetInDuration(v int32) *ModifyScalingRuleRequest {
	s.InDuration = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetInEnable(v bool) *ModifyScalingRuleRequest {
	s.InEnable = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetInInstanceNum(v int32) *ModifyScalingRuleRequest {
	s.InInstanceNum = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetInLoad(v int32) *ModifyScalingRuleRequest {
	s.InLoad = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetInRT(v int32) *ModifyScalingRuleRequest {
	s.InRT = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetInStep(v int32) *ModifyScalingRuleRequest {
	s.InStep = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetKeyPairName(v string) *ModifyScalingRuleRequest {
	s.KeyPairName = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetMultiAzPolicy(v string) *ModifyScalingRuleRequest {
	s.MultiAzPolicy = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetOutCPU(v int32) *ModifyScalingRuleRequest {
	s.OutCPU = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetOutCondition(v string) *ModifyScalingRuleRequest {
	s.OutCondition = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetOutDuration(v int32) *ModifyScalingRuleRequest {
	s.OutDuration = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetOutEnable(v bool) *ModifyScalingRuleRequest {
	s.OutEnable = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetOutInstanceNum(v int32) *ModifyScalingRuleRequest {
	s.OutInstanceNum = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetOutLoad(v int32) *ModifyScalingRuleRequest {
	s.OutLoad = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetOutRT(v int32) *ModifyScalingRuleRequest {
	s.OutRT = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetOutStep(v int32) *ModifyScalingRuleRequest {
	s.OutStep = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetPassword(v string) *ModifyScalingRuleRequest {
	s.Password = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetResourceFrom(v string) *ModifyScalingRuleRequest {
	s.ResourceFrom = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetScalingPolicy(v string) *ModifyScalingRuleRequest {
	s.ScalingPolicy = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetTemplateId(v string) *ModifyScalingRuleRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetTemplateInstanceId(v string) *ModifyScalingRuleRequest {
	s.TemplateInstanceId = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetTemplateInstanceName(v string) *ModifyScalingRuleRequest {
	s.TemplateInstanceName = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetTemplateVersion(v int32) *ModifyScalingRuleRequest {
	s.TemplateVersion = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetVSwitchIds(v string) *ModifyScalingRuleRequest {
	s.VSwitchIds = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetVpcId(v string) *ModifyScalingRuleRequest {
	s.VpcId = &v
	return s
}

func (s *ModifyScalingRuleRequest) Validate() error {
	return dara.Validate(s)
}
