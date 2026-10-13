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
	// Must be set to true if auto scaling is used.
	//
	// example:
	//
	// true
	AcceptEULA *bool `json:"AcceptEULA,omitempty" xml:"AcceptEULA,omitempty"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 74ee****-db65-4322-a1f6-bcb60e5b****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The deployment group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8123db90-880f-486f-****-************
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The trigger condition. Valid values:
	//
	// - OR: Any condition.
	//
	// - AND: All conditions.
	//
	// example:
	//
	// "OR"
	InCondition *string `json:"InCondition,omitempty" xml:"InCondition,omitempty"`
	// The trigger metric: CPU utilization.
	//
	// example:
	//
	// 50
	InCpu *int32 `json:"InCpu,omitempty" xml:"InCpu,omitempty"`
	// The duration, in minutes.
	//
	// example:
	//
	// 50
	InDuration *int32 `json:"InDuration,omitempty" xml:"InDuration,omitempty"`
	// Specifies whether to include the scale-in rule. Valid values:
	//
	// - true: Included.
	//
	// - false: Not included.
	//
	// example:
	//
	// true
	InEnable *bool `json:"InEnable,omitempty" xml:"InEnable,omitempty"`
	// The minimum number of instances to retain when the group scales in.
	//
	// example:
	//
	// 3
	InInstanceNum *int32 `json:"InInstanceNum,omitempty" xml:"InInstanceNum,omitempty"`
	// The trigger metric: system load.
	//
	// example:
	//
	// 50
	InLoad *int32 `json:"InLoad,omitempty" xml:"InLoad,omitempty"`
	// The trigger metric: service latency. Unit: milliseconds. Minimum value: 0.
	//
	// example:
	//
	// 50
	InRT *int32 `json:"InRT,omitempty" xml:"InRT,omitempty"`
	// The number of instances to scale in at a time.
	//
	// example:
	//
	// 1
	InStep *int32 `json:"InStep,omitempty" xml:"InStep,omitempty"`
	// The name of the logon key pair for the instance. This parameter is valid only when the existing instance rule is used for auto scaling.
	//
	// example:
	//
	// "tdy218"
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The multi-zone scaling policy for auto scaling. Valid values:
	//
	// - PRIORITY: Priority policy.
	//
	// - BALANCE: Balanced distribution policy.
	//
	// example:
	//
	// "PRIORITY"
	MultiAzPolicy *string `json:"MultiAzPolicy,omitempty" xml:"MultiAzPolicy,omitempty"`
	// The trigger metric: CPU utilization.
	//
	// example:
	//
	// 50
	OutCPU *int32 `json:"OutCPU,omitempty" xml:"OutCPU,omitempty"`
	// The trigger condition. Valid values:
	//
	// - OR: Any condition.
	//
	// - AND: All conditions.
	//
	// example:
	//
	// "OR"
	OutCondition *string `json:"OutCondition,omitempty" xml:"OutCondition,omitempty"`
	// The duration, in minutes.
	//
	// example:
	//
	// 50
	OutDuration *int32 `json:"OutDuration,omitempty" xml:"OutDuration,omitempty"`
	// Specifies whether to include the scale-out rule.
	//
	// example:
	//
	// true
	OutEnable *bool `json:"OutEnable,omitempty" xml:"OutEnable,omitempty"`
	// The maximum number of instances when the group scales out.
	//
	// example:
	//
	// 10
	OutInstanceNum *int32 `json:"OutInstanceNum,omitempty" xml:"OutInstanceNum,omitempty"`
	// The trigger metric: system load.
	//
	// example:
	//
	// 50
	OutLoad *int32 `json:"OutLoad,omitempty" xml:"OutLoad,omitempty"`
	// The trigger metric: service latency. Unit: milliseconds. Minimum value: 0.
	//
	// example:
	//
	// 0
	OutRT *int32 `json:"OutRT,omitempty" xml:"OutRT,omitempty"`
	// The number of instances to scale out at a time.
	//
	// example:
	//
	// 0
	OutStep *int32 `json:"OutStep,omitempty" xml:"OutStep,omitempty"`
	// The logon password of the instance. This parameter is valid only when the existing instance rule is used for auto scaling.
	//
	// example:
	//
	// "Pwd*****"
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The source of instances for scale-out. Valid values:
	//
	// - NEW: Elastic resources.
	//
	// - AVAILABLE: Existing resources. To prioritize existing resources, use AVAILABLE_FIRST.
	//
	// If you set this parameter to NEW or AVAILABLE_FIRST, you must specify auto scaling parameters. Auto scaling includes two modes: launch template and existing instance rule.
	//
	// example:
	//
	// "AVAILABLE"
	ResourceFrom *string `json:"ResourceFrom,omitempty" xml:"ResourceFrom,omitempty"`
	// The reclamation mode for scale-in. Valid values:
	//
	// - release: Release mode.
	//
	// - recycle: Stop and reclaim mode.
	//
	// example:
	//
	// "release"
	ScalingPolicy *string `json:"ScalingPolicy,omitempty" xml:"ScalingPolicy,omitempty"`
	// The launch template ID for auto scaling. This parameter is valid only when OutEnable is set to true, and it has a higher priority than TemplateInstanceId.
	//
	// example:
	//
	// "lt-wz9hkhn8wp*****"
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The instance ID for the existing instance rule of auto scaling. This parameter is valid only when OutEnable is set to true.
	//
	// example:
	//
	// "1"
	TemplateInstanceId *string `json:"TemplateInstanceId,omitempty" xml:"TemplateInstanceId,omitempty"`
	// The instance name for the existing instance rule of auto scaling. This parameter is valid only when TemplateInstanceId is specified.
	//
	// example:
	//
	// "tpl-tdy218"
	TemplateInstanceName *string `json:"TemplateInstanceName,omitempty" xml:"TemplateInstanceName,omitempty"`
	// The launch template version for auto scaling. This parameter is valid only when OutEnable is set to true. Set this parameter to `-1` to use the default template version, or set it to a specific version number.
	//
	// example:
	//
	// -1
	TemplateVersion *int32 `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The list of vSwitch IDs in the VPC. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// "vsw-bp1ldxs3d4fd*****"
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The VPC ID associated with the launch template or existing instance rule for auto scaling.
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
