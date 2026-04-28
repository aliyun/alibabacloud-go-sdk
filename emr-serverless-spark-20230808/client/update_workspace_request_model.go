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
	// example:
	//
	// 5000
	Cu *int32 `json:"cu,omitempty" xml:"cu,omitempty"`
	// example:
	//
	// 100
	Gpu     *int32    `json:"gpu,omitempty" xml:"gpu,omitempty"`
	GpuSpec []*string `json:"gpuSpec,omitempty" xml:"gpuSpec,omitempty" type:"Repeated"`
	// example:
	//
	// rg-acfmwpi66knkxny
	ResourceGroupId *string                             `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	Subscription    *UpdateWorkspaceRequestSubscription `json:"subscription,omitempty" xml:"subscription,omitempty" type:"Struct"`
	// example:
	//
	// w-975bcfda9625****
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// example:
	//
	// default
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
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
	if s.Subscription != nil {
		if err := s.Subscription.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWorkspaceRequestSubscription struct {
	// example:
	//
	// true
	AutoRenew *string `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	// example:
	//
	// 1
	AutoRenewPeriod *string `json:"autoRenewPeriod,omitempty" xml:"autoRenewPeriod,omitempty"`
	// example:
	//
	// MONTH
	AutoRenewPeriodUnit *string `json:"autoRenewPeriodUnit,omitempty" xml:"autoRenewPeriodUnit,omitempty"`
	// example:
	//
	// my-token-asxkxxxxxxx
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// 1799
	Duration *string `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// MONTH
	PaymentDurationUnit *string   `json:"paymentDurationUnit,omitempty" xml:"paymentDurationUnit,omitempty"`
	Queue               []*string `json:"queue,omitempty" xml:"queue,omitempty" type:"Repeated"`
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
