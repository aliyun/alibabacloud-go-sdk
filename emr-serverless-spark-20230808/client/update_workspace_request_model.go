// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCu(v int32) *UpdateWorkspaceRequest
	GetCu() *int32
	SetGpu(v int32) *UpdateWorkspaceRequest
	GetGpu() *int32
	SetGpuSpec(v []*string) *UpdateWorkspaceRequest
	GetGpuSpec() []*string
	SetGpuSubscription(v *UpdateWorkspaceRequestGpuSubscription) *UpdateWorkspaceRequest
	GetGpuSubscription() *UpdateWorkspaceRequestGpuSubscription
	SetIpWhiteList(v []*string) *UpdateWorkspaceRequest
	GetIpWhiteList() []*string
	SetResourceGroupId(v string) *UpdateWorkspaceRequest
	GetResourceGroupId() *string
	SetSubscription(v *UpdateWorkspaceRequestSubscription) *UpdateWorkspaceRequest
	GetSubscription() *UpdateWorkspaceRequestSubscription
	SetWorkspaceId(v string) *UpdateWorkspaceRequest
	GetWorkspaceId() *string
	SetWorkspaceName(v string) *UpdateWorkspaceRequest
	GetWorkspaceName() *string
	SetRegionId(v string) *UpdateWorkspaceRequest
	GetRegionId() *string
}

type UpdateWorkspaceRequest struct {
	// The upper limit of workspace resources.
	//
	// example:
	//
	// 5000
	Cu *int32 `json:"cu,omitempty" xml:"cu,omitempty"`
	// The number of GPU cards.
	//
	// example:
	//
	// 100
	Gpu *int32 `json:"gpu,omitempty" xml:"gpu,omitempty"`
	// The GPU instance type.
	GpuSpec         []*string                              `json:"gpuSpec,omitempty" xml:"gpuSpec,omitempty" type:"Repeated"`
	GpuSubscription *UpdateWorkspaceRequestGpuSubscription `json:"gpuSubscription,omitempty" xml:"gpuSubscription,omitempty" type:"Struct"`
	IpWhiteList     []*string                              `json:"ipWhiteList,omitempty" xml:"ipWhiteList,omitempty" type:"Repeated"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmwpi66knkxny
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The information for converting from pay-as-you-go to subscription.
	Subscription *UpdateWorkspaceRequestSubscription `json:"subscription,omitempty" xml:"subscription,omitempty" type:"Struct"`
	// The workspace ID.
	//
	// example:
	//
	// w-975bcfda9625****
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// The workspace name.
	//
	// example:
	//
	// default
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s UpdateWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceRequest) GetCu() *int32 {
	return s.Cu
}

func (s *UpdateWorkspaceRequest) GetGpu() *int32 {
	return s.Gpu
}

func (s *UpdateWorkspaceRequest) GetGpuSpec() []*string {
	return s.GpuSpec
}

func (s *UpdateWorkspaceRequest) GetGpuSubscription() *UpdateWorkspaceRequestGpuSubscription {
	return s.GpuSubscription
}

func (s *UpdateWorkspaceRequest) GetIpWhiteList() []*string {
	return s.IpWhiteList
}

func (s *UpdateWorkspaceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateWorkspaceRequest) GetSubscription() *UpdateWorkspaceRequestSubscription {
	return s.Subscription
}

func (s *UpdateWorkspaceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateWorkspaceRequest) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *UpdateWorkspaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateWorkspaceRequest) SetCu(v int32) *UpdateWorkspaceRequest {
	s.Cu = &v
	return s
}

func (s *UpdateWorkspaceRequest) SetGpu(v int32) *UpdateWorkspaceRequest {
	s.Gpu = &v
	return s
}

func (s *UpdateWorkspaceRequest) SetGpuSpec(v []*string) *UpdateWorkspaceRequest {
	s.GpuSpec = v
	return s
}

func (s *UpdateWorkspaceRequest) SetGpuSubscription(v *UpdateWorkspaceRequestGpuSubscription) *UpdateWorkspaceRequest {
	s.GpuSubscription = v
	return s
}

func (s *UpdateWorkspaceRequest) SetIpWhiteList(v []*string) *UpdateWorkspaceRequest {
	s.IpWhiteList = v
	return s
}

func (s *UpdateWorkspaceRequest) SetResourceGroupId(v string) *UpdateWorkspaceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateWorkspaceRequest) SetSubscription(v *UpdateWorkspaceRequestSubscription) *UpdateWorkspaceRequest {
	s.Subscription = v
	return s
}

func (s *UpdateWorkspaceRequest) SetWorkspaceId(v string) *UpdateWorkspaceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateWorkspaceRequest) SetWorkspaceName(v string) *UpdateWorkspaceRequest {
	s.WorkspaceName = &v
	return s
}

func (s *UpdateWorkspaceRequest) SetRegionId(v string) *UpdateWorkspaceRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateWorkspaceRequest) Validate() error {
	if s.GpuSubscription != nil {
		if err := s.GpuSubscription.Validate(); err != nil {
			return err
		}
	}
	if s.Subscription != nil {
		if err := s.Subscription.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWorkspaceRequestGpuSubscription struct {
	AutoRenew *bool  `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	Duration  *int32 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 8
	GpuMachineNum *int32  `json:"gpuMachineNum,omitempty" xml:"gpuMachineNum,omitempty"`
	InstanceId    *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// ecs.gn7i-c8g1.2xlarge
	InstanceTypeId *string `json:"instanceTypeId,omitempty" xml:"instanceTypeId,omitempty"`
	// example:
	//
	// BUY
	Operation           *string `json:"operation,omitempty" xml:"operation,omitempty"`
	PaymentDurationUnit *string `json:"paymentDurationUnit,omitempty" xml:"paymentDurationUnit,omitempty"`
}

func (s UpdateWorkspaceRequestGpuSubscription) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceRequestGpuSubscription) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceRequestGpuSubscription) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *UpdateWorkspaceRequestGpuSubscription) GetDuration() *int32 {
	return s.Duration
}

func (s *UpdateWorkspaceRequestGpuSubscription) GetGpuMachineNum() *int32 {
	return s.GpuMachineNum
}

func (s *UpdateWorkspaceRequestGpuSubscription) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateWorkspaceRequestGpuSubscription) GetInstanceTypeId() *string {
	return s.InstanceTypeId
}

func (s *UpdateWorkspaceRequestGpuSubscription) GetOperation() *string {
	return s.Operation
}

func (s *UpdateWorkspaceRequestGpuSubscription) GetPaymentDurationUnit() *string {
	return s.PaymentDurationUnit
}

func (s *UpdateWorkspaceRequestGpuSubscription) SetAutoRenew(v bool) *UpdateWorkspaceRequestGpuSubscription {
	s.AutoRenew = &v
	return s
}

func (s *UpdateWorkspaceRequestGpuSubscription) SetDuration(v int32) *UpdateWorkspaceRequestGpuSubscription {
	s.Duration = &v
	return s
}

func (s *UpdateWorkspaceRequestGpuSubscription) SetGpuMachineNum(v int32) *UpdateWorkspaceRequestGpuSubscription {
	s.GpuMachineNum = &v
	return s
}

func (s *UpdateWorkspaceRequestGpuSubscription) SetInstanceId(v string) *UpdateWorkspaceRequestGpuSubscription {
	s.InstanceId = &v
	return s
}

func (s *UpdateWorkspaceRequestGpuSubscription) SetInstanceTypeId(v string) *UpdateWorkspaceRequestGpuSubscription {
	s.InstanceTypeId = &v
	return s
}

func (s *UpdateWorkspaceRequestGpuSubscription) SetOperation(v string) *UpdateWorkspaceRequestGpuSubscription {
	s.Operation = &v
	return s
}

func (s *UpdateWorkspaceRequestGpuSubscription) SetPaymentDurationUnit(v string) *UpdateWorkspaceRequestGpuSubscription {
	s.PaymentDurationUnit = &v
	return s
}

func (s *UpdateWorkspaceRequestGpuSubscription) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkspaceRequestSubscription struct {
	// Specifies whether to enable auto-renewal. This parameter is required for the pre-paid billing type.
	//
	// example:
	//
	// true
	AutoRenew *string `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	// The auto-renewal duration. This parameter is required for the pre-paid billing type.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *string `json:"autoRenewPeriod,omitempty" xml:"autoRenewPeriod,omitempty"`
	// The auto-renewal period unit. This parameter is required for the pre-paid billing type.
	//
	// example:
	//
	// MONTH
	AutoRenewPeriodUnit *string `json:"autoRenewPeriodUnit,omitempty" xml:"autoRenewPeriodUnit,omitempty"`
	// The idempotency token.
	//
	// example:
	//
	// my-token-asxkxxxxxxx
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The number of subscription periods. This parameter is required for the pre-paid billing type.
	//
	// example:
	//
	// 1799
	Duration *string `json:"duration,omitempty" xml:"duration,omitempty"`
	// The subscription period unit.
	//
	// example:
	//
	// MONTH
	PaymentDurationUnit *string `json:"paymentDurationUnit,omitempty" xml:"paymentDurationUnit,omitempty"`
	// The list of running queues to be converted.
	Queue []*string `json:"queue,omitempty" xml:"queue,omitempty" type:"Repeated"`
}

func (s UpdateWorkspaceRequestSubscription) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceRequestSubscription) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceRequestSubscription) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *UpdateWorkspaceRequestSubscription) GetAutoRenewPeriod() *string {
	return s.AutoRenewPeriod
}

func (s *UpdateWorkspaceRequestSubscription) GetAutoRenewPeriodUnit() *string {
	return s.AutoRenewPeriodUnit
}

func (s *UpdateWorkspaceRequestSubscription) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateWorkspaceRequestSubscription) GetDuration() *string {
	return s.Duration
}

func (s *UpdateWorkspaceRequestSubscription) GetPaymentDurationUnit() *string {
	return s.PaymentDurationUnit
}

func (s *UpdateWorkspaceRequestSubscription) GetQueue() []*string {
	return s.Queue
}

func (s *UpdateWorkspaceRequestSubscription) SetAutoRenew(v string) *UpdateWorkspaceRequestSubscription {
	s.AutoRenew = &v
	return s
}

func (s *UpdateWorkspaceRequestSubscription) SetAutoRenewPeriod(v string) *UpdateWorkspaceRequestSubscription {
	s.AutoRenewPeriod = &v
	return s
}

func (s *UpdateWorkspaceRequestSubscription) SetAutoRenewPeriodUnit(v string) *UpdateWorkspaceRequestSubscription {
	s.AutoRenewPeriodUnit = &v
	return s
}

func (s *UpdateWorkspaceRequestSubscription) SetClientToken(v string) *UpdateWorkspaceRequestSubscription {
	s.ClientToken = &v
	return s
}

func (s *UpdateWorkspaceRequestSubscription) SetDuration(v string) *UpdateWorkspaceRequestSubscription {
	s.Duration = &v
	return s
}

func (s *UpdateWorkspaceRequestSubscription) SetPaymentDurationUnit(v string) *UpdateWorkspaceRequestSubscription {
	s.PaymentDurationUnit = &v
	return s
}

func (s *UpdateWorkspaceRequestSubscription) SetQueue(v []*string) *UpdateWorkspaceRequestSubscription {
	s.Queue = v
	return s
}

func (s *UpdateWorkspaceRequestSubscription) Validate() error {
	return dara.Validate(s)
}
