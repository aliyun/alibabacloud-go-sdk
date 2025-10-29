// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunResourceGroupId(v string) *CreateResourceGroupRequest
	GetAliyunResourceGroupId() *string
	SetAliyunResourceTags(v []*CreateResourceGroupRequestAliyunResourceTags) *CreateResourceGroupRequest
	GetAliyunResourceTags() []*CreateResourceGroupRequestAliyunResourceTags
	SetAutoRenewEnabled(v bool) *CreateResourceGroupRequest
	GetAutoRenewEnabled() *bool
	SetClientToken(v string) *CreateResourceGroupRequest
	GetClientToken() *string
	SetName(v string) *CreateResourceGroupRequest
	GetName() *string
	SetPaymentDuration(v int32) *CreateResourceGroupRequest
	GetPaymentDuration() *int32
	SetPaymentDurationUnit(v string) *CreateResourceGroupRequest
	GetPaymentDurationUnit() *string
	SetPaymentType(v string) *CreateResourceGroupRequest
	GetPaymentType() *string
	SetRemark(v string) *CreateResourceGroupRequest
	GetRemark() *string
	SetSpec(v int32) *CreateResourceGroupRequest
	GetSpec() *int32
	SetVpcId(v string) *CreateResourceGroupRequest
	GetVpcId() *string
	SetVswitchId(v string) *CreateResourceGroupRequest
	GetVswitchId() *string
}

type CreateResourceGroupRequest struct {
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-aek2kqofrgXXXXX
	AliyunResourceGroupId *string `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	// The tags.
	AliyunResourceTags []*CreateResourceGroupRequestAliyunResourceTags `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty" type:"Repeated"`
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

func (s CreateResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupRequest) GetAliyunResourceGroupId() *string {
	return s.AliyunResourceGroupId
}

func (s *CreateResourceGroupRequest) GetAliyunResourceTags() []*CreateResourceGroupRequestAliyunResourceTags {
	return s.AliyunResourceTags
}

func (s *CreateResourceGroupRequest) GetAutoRenewEnabled() *bool {
	return s.AutoRenewEnabled
}

func (s *CreateResourceGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateResourceGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateResourceGroupRequest) GetPaymentDuration() *int32 {
	return s.PaymentDuration
}

func (s *CreateResourceGroupRequest) GetPaymentDurationUnit() *string {
	return s.PaymentDurationUnit
}

func (s *CreateResourceGroupRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CreateResourceGroupRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateResourceGroupRequest) GetSpec() *int32 {
	return s.Spec
}

func (s *CreateResourceGroupRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateResourceGroupRequest) GetVswitchId() *string {
	return s.VswitchId
}

func (s *CreateResourceGroupRequest) SetAliyunResourceGroupId(v string) *CreateResourceGroupRequest {
	s.AliyunResourceGroupId = &v
	return s
}

func (s *CreateResourceGroupRequest) SetAliyunResourceTags(v []*CreateResourceGroupRequestAliyunResourceTags) *CreateResourceGroupRequest {
	s.AliyunResourceTags = v
	return s
}

func (s *CreateResourceGroupRequest) SetAutoRenewEnabled(v bool) *CreateResourceGroupRequest {
	s.AutoRenewEnabled = &v
	return s
}

func (s *CreateResourceGroupRequest) SetClientToken(v string) *CreateResourceGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateResourceGroupRequest) SetName(v string) *CreateResourceGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateResourceGroupRequest) SetPaymentDuration(v int32) *CreateResourceGroupRequest {
	s.PaymentDuration = &v
	return s
}

func (s *CreateResourceGroupRequest) SetPaymentDurationUnit(v string) *CreateResourceGroupRequest {
	s.PaymentDurationUnit = &v
	return s
}

func (s *CreateResourceGroupRequest) SetPaymentType(v string) *CreateResourceGroupRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateResourceGroupRequest) SetRemark(v string) *CreateResourceGroupRequest {
	s.Remark = &v
	return s
}

func (s *CreateResourceGroupRequest) SetSpec(v int32) *CreateResourceGroupRequest {
	s.Spec = &v
	return s
}

func (s *CreateResourceGroupRequest) SetVpcId(v string) *CreateResourceGroupRequest {
	s.VpcId = &v
	return s
}

func (s *CreateResourceGroupRequest) SetVswitchId(v string) *CreateResourceGroupRequest {
	s.VswitchId = &v
	return s
}

func (s *CreateResourceGroupRequest) Validate() error {
	if s.AliyunResourceTags != nil {
		for _, item := range s.AliyunResourceTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateResourceGroupRequestAliyunResourceTags struct {
	// The tag key.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateResourceGroupRequestAliyunResourceTags) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceGroupRequestAliyunResourceTags) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupRequestAliyunResourceTags) GetKey() *string {
	return s.Key
}

func (s *CreateResourceGroupRequestAliyunResourceTags) GetValue() *string {
	return s.Value
}

func (s *CreateResourceGroupRequestAliyunResourceTags) SetKey(v string) *CreateResourceGroupRequestAliyunResourceTags {
	s.Key = &v
	return s
}

func (s *CreateResourceGroupRequestAliyunResourceTags) SetValue(v string) *CreateResourceGroupRequestAliyunResourceTags {
	s.Value = &v
	return s
}

func (s *CreateResourceGroupRequestAliyunResourceTags) Validate() error {
	return dara.Validate(s)
}
