// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunResourceGroupId(v string) *CreateResourceGroupShrinkRequest
	GetAliyunResourceGroupId() *string
	SetAliyunResourceTagsShrink(v string) *CreateResourceGroupShrinkRequest
	GetAliyunResourceTagsShrink() *string
	SetAutoRenewEnabled(v bool) *CreateResourceGroupShrinkRequest
	GetAutoRenewEnabled() *bool
	SetClientToken(v string) *CreateResourceGroupShrinkRequest
	GetClientToken() *string
	SetName(v string) *CreateResourceGroupShrinkRequest
	GetName() *string
	SetPaymentDuration(v int32) *CreateResourceGroupShrinkRequest
	GetPaymentDuration() *int32
	SetPaymentDurationUnit(v string) *CreateResourceGroupShrinkRequest
	GetPaymentDurationUnit() *string
	SetPaymentType(v string) *CreateResourceGroupShrinkRequest
	GetPaymentType() *string
	SetRemark(v string) *CreateResourceGroupShrinkRequest
	GetRemark() *string
	SetSpec(v int32) *CreateResourceGroupShrinkRequest
	GetSpec() *int32
	SetVpcId(v string) *CreateResourceGroupShrinkRequest
	GetVpcId() *string
	SetVswitchId(v string) *CreateResourceGroupShrinkRequest
	GetVswitchId() *string
}

type CreateResourceGroupShrinkRequest struct {
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-aek2kqofrgXXXXX
	AliyunResourceGroupId *string `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	// The tags.
	AliyunResourceTagsShrink *string `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty"`
	// Specifies whether to enable auto-renewal.
	AutoRenewEnabled *bool `json:"AutoRenewEnabled,omitempty" xml:"AutoRenewEnabled,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// eb870033-74c8-4b1b-9664-04c4e7cc3465
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the serverless resource group. The name can be a maximum of 128 characters in length and can contain letters, digits, and underscores (_). The name must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// common_resource_group
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The subscription duration.
	//
	// example:
	//
	// 1
	PaymentDuration *int32 `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	// The unit of the subscription duration. Valid values: Month and Year.
	//
	// example:
	//
	// Month
	PaymentDurationUnit *string `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	// The billing method of the serverless resource group. Valid values: PrePaid and PostPaid. The value PrePaid indicates the subscription billing method, and the value PostPaid indicates the pay-as-you-go billing method.
	//
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The description of the serverless resource group. The description can be a maximum of 128 characters in length and can contain letters, digits, and underscores (_).
	//
	// example:
	//
	// Create a serverless resource group for common tasks
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The specifications of the serverless resource group. Unit: CU. This parameter is required only if you set the PaymentType parameter to PrePaid.
	//
	// example:
	//
	// 2
	Spec *int32 `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The ID of the virtual private cloud (VPC) with which the serverless resource group is associated by default.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-m2et4f3oc8msfbccXXXXX
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the vSwitch with which the serverless resource group is associated by default.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-uf8usrhs7hjd9amsXXXXX
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s CreateResourceGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupShrinkRequest) GetAliyunResourceGroupId() *string {
	return s.AliyunResourceGroupId
}

func (s *CreateResourceGroupShrinkRequest) GetAliyunResourceTagsShrink() *string {
	return s.AliyunResourceTagsShrink
}

func (s *CreateResourceGroupShrinkRequest) GetAutoRenewEnabled() *bool {
	return s.AutoRenewEnabled
}

func (s *CreateResourceGroupShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateResourceGroupShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateResourceGroupShrinkRequest) GetPaymentDuration() *int32 {
	return s.PaymentDuration
}

func (s *CreateResourceGroupShrinkRequest) GetPaymentDurationUnit() *string {
	return s.PaymentDurationUnit
}

func (s *CreateResourceGroupShrinkRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CreateResourceGroupShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateResourceGroupShrinkRequest) GetSpec() *int32 {
	return s.Spec
}

func (s *CreateResourceGroupShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateResourceGroupShrinkRequest) GetVswitchId() *string {
	return s.VswitchId
}

func (s *CreateResourceGroupShrinkRequest) SetAliyunResourceGroupId(v string) *CreateResourceGroupShrinkRequest {
	s.AliyunResourceGroupId = &v
	return s
}

func (s *CreateResourceGroupShrinkRequest) SetAliyunResourceTagsShrink(v string) *CreateResourceGroupShrinkRequest {
	s.AliyunResourceTagsShrink = &v
	return s
}

func (s *CreateResourceGroupShrinkRequest) SetAutoRenewEnabled(v bool) *CreateResourceGroupShrinkRequest {
	s.AutoRenewEnabled = &v
	return s
}

func (s *CreateResourceGroupShrinkRequest) SetClientToken(v string) *CreateResourceGroupShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateResourceGroupShrinkRequest) SetName(v string) *CreateResourceGroupShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateResourceGroupShrinkRequest) SetPaymentDuration(v int32) *CreateResourceGroupShrinkRequest {
	s.PaymentDuration = &v
	return s
}

func (s *CreateResourceGroupShrinkRequest) SetPaymentDurationUnit(v string) *CreateResourceGroupShrinkRequest {
	s.PaymentDurationUnit = &v
	return s
}

func (s *CreateResourceGroupShrinkRequest) SetPaymentType(v string) *CreateResourceGroupShrinkRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateResourceGroupShrinkRequest) SetRemark(v string) *CreateResourceGroupShrinkRequest {
	s.Remark = &v
	return s
}

func (s *CreateResourceGroupShrinkRequest) SetSpec(v int32) *CreateResourceGroupShrinkRequest {
	s.Spec = &v
	return s
}

func (s *CreateResourceGroupShrinkRequest) SetVpcId(v string) *CreateResourceGroupShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateResourceGroupShrinkRequest) SetVswitchId(v string) *CreateResourceGroupShrinkRequest {
	s.VswitchId = &v
	return s
}

func (s *CreateResourceGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
