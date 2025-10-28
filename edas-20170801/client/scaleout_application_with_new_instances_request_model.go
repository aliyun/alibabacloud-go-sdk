// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleoutApplicationWithNewInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ScaleoutApplicationWithNewInstancesRequest
	GetAppId() *string
	SetAutoRenew(v bool) *ScaleoutApplicationWithNewInstancesRequest
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v int32) *ScaleoutApplicationWithNewInstancesRequest
	GetAutoRenewPeriod() *int32
	SetClusterId(v string) *ScaleoutApplicationWithNewInstancesRequest
	GetClusterId() *string
	SetGroupId(v string) *ScaleoutApplicationWithNewInstancesRequest
	GetGroupId() *string
	SetInstanceChargePeriod(v int32) *ScaleoutApplicationWithNewInstancesRequest
	GetInstanceChargePeriod() *int32
	SetInstanceChargePeriodUnit(v string) *ScaleoutApplicationWithNewInstancesRequest
	GetInstanceChargePeriodUnit() *string
	SetInstanceChargeType(v string) *ScaleoutApplicationWithNewInstancesRequest
	GetInstanceChargeType() *string
	SetScalingNum(v int32) *ScaleoutApplicationWithNewInstancesRequest
	GetScalingNum() *int32
	SetScalingPolicy(v string) *ScaleoutApplicationWithNewInstancesRequest
	GetScalingPolicy() *string
	SetTemplateId(v string) *ScaleoutApplicationWithNewInstancesRequest
	GetTemplateId() *string
	SetTemplateInstanceId(v string) *ScaleoutApplicationWithNewInstancesRequest
	GetTemplateInstanceId() *string
	SetTemplateVersion(v string) *ScaleoutApplicationWithNewInstancesRequest
	GetTemplateVersion() *string
}

type ScaleoutApplicationWithNewInstancesRequest struct {
	// The ID of the application that you want to scale out. You can call the ListApplication operation to query the application ID. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// example:
	//
	// e370c17f-*****-3df0721a327
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Specifies whether to enable auto-renewal. This parameter takes effect only when the InstanceChargeType parameter is set to PrePaid. Valid values:
	//
	// 	- true: enables auto-renewal.
	//
	// 	- false: does not enable auto-renewal. This is the default value.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal period. Valid values:
	//
	// 	- If the InstanceChargePeriodUnit parameter is set to Week, the valid values of the AutoRenewPeriod parameter are 1, 2, and 3.
	//
	// 	- If the InstanceChargePeriodUnit parameter is set to Month, the valid values of the AutoRenewPeriod parameter are 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The ID of the cluster to which you want to add ECS instances. If the application and application instance group for the scale-out are specified, this parameter is ignored.
	//
	// example:
	//
	// e37**********-33df0721a327
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the instance group that you want to scale out. You can call the ListDeployGroup operation to query the group ID. For more information, see [ListDeployGroup](https://help.aliyun.com/document_detail/62077.html).
	//
	// example:
	//
	// e37**********-33df0721a327
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The duration of the subscription. The unit of the subscription duration is specified by the InstanceChargePeriodUnit parameter. This parameter takes effect only when the InstanceChargeType parameter is set to PrePaid.
	//
	// 	- If the InstanceChargePeriodUnit parameter is set to Week, the valid values of the InstanceChargePeriod parameter are 1, 2, 3, and 4.
	//
	// 	- If the InstanceChargePeriodUnit parameter is set to Month, the valid values of the InstanceChargePeriod parameter are 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	//
	// example:
	//
	// 1
	InstanceChargePeriod *int32 `json:"InstanceChargePeriod,omitempty" xml:"InstanceChargePeriod,omitempty"`
	// The unit of the subscription period. Valid values:
	//
	// 	- Week: billed on a weekly basis.
	//
	// 	- Month: billed on a monthly basis. This is the default value.
	//
	// example:
	//
	// Month
	InstanceChargePeriodUnit *string `json:"InstanceChargePeriodUnit,omitempty" xml:"InstanceChargePeriodUnit,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- PrePaid: subscription.
	//
	// 	- PostPaid: pay-as-you-go. This is the default value.
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The number of instances to be added for the scale-out.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ScalingNum *int32 `json:"ScalingNum,omitempty" xml:"ScalingNum,omitempty"`
	// The instance reclaim mode of the scaling group. Valid values:
	//
	// 	- recycle: economical mode
	//
	// 	- release: release mode
	//
	// For more information about how to remove instances from a specified scaling group, see [RemoveInstances](https://help.aliyun.com/document_detail/25955.html).
	//
	// example:
	//
	// release
	ScalingPolicy *string `json:"ScalingPolicy,omitempty" xml:"ScalingPolicy,omitempty"`
	// The ID of the ECS instance launch template. You can call the DescribeLaunchTemplates operation to query the launch template ID. For more information, see [DescribeLaunchTemplates](https://help.aliyun.com/document_detail/73759.html).
	//
	// example:
	//
	// lt-****hy9s2
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The ID of the existing ECS instance used for the scale-out. If this parameter is specified, the specifications and configurations of the specified ECS instance are used as a template to purchase new instances.
	//
	// example:
	//
	// i-28wt4****
	TemplateInstanceId *string `json:"TemplateInstanceId,omitempty" xml:"TemplateInstanceId,omitempty"`
	// The version of the ECS instance launch template. You can call the DescribeLaunchTemplateVersions operation to query the launch template version. For more information, see [DescribeLaunchTemplateVersions](https://help.aliyun.com/document_detail/73761.html).
	//
	// > If you set this parameter to `-1`, the default launch template version is used.
	//
	// example:
	//
	// -1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s ScaleoutApplicationWithNewInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ScaleoutApplicationWithNewInstancesRequest) GoString() string {
	return s.String()
}

func (s *ScaleoutApplicationWithNewInstancesRequest) GetAppId() *string {
	return s.AppId
}

func (s *ScaleoutApplicationWithNewInstancesRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *ScaleoutApplicationWithNewInstancesRequest) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *ScaleoutApplicationWithNewInstancesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ScaleoutApplicationWithNewInstancesRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ScaleoutApplicationWithNewInstancesRequest) GetInstanceChargePeriod() *int32 {
	return s.InstanceChargePeriod
}

func (s *ScaleoutApplicationWithNewInstancesRequest) GetInstanceChargePeriodUnit() *string {
	return s.InstanceChargePeriodUnit
}

func (s *ScaleoutApplicationWithNewInstancesRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *ScaleoutApplicationWithNewInstancesRequest) GetScalingNum() *int32 {
	return s.ScalingNum
}

func (s *ScaleoutApplicationWithNewInstancesRequest) GetScalingPolicy() *string {
	return s.ScalingPolicy
}

func (s *ScaleoutApplicationWithNewInstancesRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ScaleoutApplicationWithNewInstancesRequest) GetTemplateInstanceId() *string {
	return s.TemplateInstanceId
}

func (s *ScaleoutApplicationWithNewInstancesRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *ScaleoutApplicationWithNewInstancesRequest) SetAppId(v string) *ScaleoutApplicationWithNewInstancesRequest {
	s.AppId = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesRequest) SetAutoRenew(v bool) *ScaleoutApplicationWithNewInstancesRequest {
	s.AutoRenew = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesRequest) SetAutoRenewPeriod(v int32) *ScaleoutApplicationWithNewInstancesRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesRequest) SetClusterId(v string) *ScaleoutApplicationWithNewInstancesRequest {
	s.ClusterId = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesRequest) SetGroupId(v string) *ScaleoutApplicationWithNewInstancesRequest {
	s.GroupId = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesRequest) SetInstanceChargePeriod(v int32) *ScaleoutApplicationWithNewInstancesRequest {
	s.InstanceChargePeriod = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesRequest) SetInstanceChargePeriodUnit(v string) *ScaleoutApplicationWithNewInstancesRequest {
	s.InstanceChargePeriodUnit = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesRequest) SetInstanceChargeType(v string) *ScaleoutApplicationWithNewInstancesRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesRequest) SetScalingNum(v int32) *ScaleoutApplicationWithNewInstancesRequest {
	s.ScalingNum = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesRequest) SetScalingPolicy(v string) *ScaleoutApplicationWithNewInstancesRequest {
	s.ScalingPolicy = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesRequest) SetTemplateId(v string) *ScaleoutApplicationWithNewInstancesRequest {
	s.TemplateId = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesRequest) SetTemplateInstanceId(v string) *ScaleoutApplicationWithNewInstancesRequest {
	s.TemplateInstanceId = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesRequest) SetTemplateVersion(v string) *ScaleoutApplicationWithNewInstancesRequest {
	s.TemplateVersion = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesRequest) Validate() error {
	return dara.Validate(s)
}
