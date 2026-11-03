// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceInstanceSpecShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateServiceInstanceSpecShrinkRequest
	GetClientToken() *string
	SetCommodity(v *UpdateServiceInstanceSpecShrinkRequestCommodity) *UpdateServiceInstanceSpecShrinkRequest
	GetCommodity() *UpdateServiceInstanceSpecShrinkRequestCommodity
	SetDryRun(v bool) *UpdateServiceInstanceSpecShrinkRequest
	GetDryRun() *bool
	SetEnableUserPrometheus(v bool) *UpdateServiceInstanceSpecShrinkRequest
	GetEnableUserPrometheus() *bool
	SetOperationName(v string) *UpdateServiceInstanceSpecShrinkRequest
	GetOperationName() *string
	SetParametersShrink(v string) *UpdateServiceInstanceSpecShrinkRequest
	GetParametersShrink() *string
	SetPredefinedParametersName(v string) *UpdateServiceInstanceSpecShrinkRequest
	GetPredefinedParametersName() *string
	SetServiceInstanceId(v string) *UpdateServiceInstanceSpecShrinkRequest
	GetServiceInstanceId() *string
}

type UpdateServiceInstanceSpecShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can generate the value from your client, but you must ensure that the value is unique among different requests. The ClientToken parameter supports only ASCII characters and can be up to 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The purchase order information for Alibaba Cloud Marketplace. This parameter is not required if the service is not listed in Alibaba Cloud Marketplace or uses the pay-as-you-go billing method.
	Commodity *UpdateServiceInstanceSpecShrinkRequestCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// Specifies whether to perform a dry run on the request, including checks on permissions and instance status.
	//
	// Valid values:
	//
	// - true: sends the request but does not create a service instance.
	//
	// - false: sends the request and creates the service instance if the check is passed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to enable Prometheus monitoring on the user side.
	//
	// Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	EnableUserPrometheus *bool `json:"EnableUserPrometheus,omitempty" xml:"EnableUserPrometheus,omitempty"`
	// The name of the specification modification operation.
	//
	// You can obtain the name and content of the specification modification operation configured for the service by viewing the **ModifyParametersConfig*	- property of the **OperationMetadata*	- response parameter in the [GetService - Get service details](https://help.aliyun.com/document_detail/2340828.html) operation.
	//
	// example:
	//
	// Plan configuration change
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// The modification parameters of the service instance.
	//
	// You can specify this parameter if the service provider set **Change Type*	- to **Parameter Change*	- when adding a specification modification operation.
	//
	// > - You can obtain the modification parameters configured for the service by viewing the **ModifyParametersConfig*	- property of the **OperationMetadata*	- response parameter in the [GetService - Get service details](https://help.aliyun.com/document_detail/2340828.html) operation.
	//
	// > - You can also view the modification parameters configured for the service on the **Service Instance Modification*	- page in the [Compute Nest console](https://computenest.console.aliyun.com/service/instance/cn-hangzhou).
	//
	// For example, if the service is configured to upgrade the ECS instance type, you must specify an instance type that is higher than the original instance type of the service instance.
	//
	// example:
	//
	// {
	//
	//   "InstanceType": "ecs.g8ise.2xlarge"
	//
	// }
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The plan name of the service instance.
	//
	// You can specify this parameter if the service provider set **Change Type*	- to **Plan Change*	- when adding a specification modification operation.
	//
	// You can obtain the plan name configured for the service by viewing the **PredefinedParameters*	- property of the **DeployMetadata*	- response parameter in the [GetService - Get service details](https://help.aliyun.com/document_detail/2340828.html) operation.
	//
	// example:
	//
	// Package 1
	PredefinedParametersName *string `json:"PredefinedParametersName,omitempty" xml:"PredefinedParametersName,omitempty"`
	// The ID of the service instance.
	//
	// You can call the [ListServiceInstances - List service instances](https://help.aliyun.com/document_detail/396200.html) operation to obtain the service instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s UpdateServiceInstanceSpecShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceInstanceSpecShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateServiceInstanceSpecShrinkRequest) GetCommodity() *UpdateServiceInstanceSpecShrinkRequestCommodity {
	return s.Commodity
}

func (s *UpdateServiceInstanceSpecShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateServiceInstanceSpecShrinkRequest) GetEnableUserPrometheus() *bool {
	return s.EnableUserPrometheus
}

func (s *UpdateServiceInstanceSpecShrinkRequest) GetOperationName() *string {
	return s.OperationName
}

func (s *UpdateServiceInstanceSpecShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *UpdateServiceInstanceSpecShrinkRequest) GetPredefinedParametersName() *string {
	return s.PredefinedParametersName
}

func (s *UpdateServiceInstanceSpecShrinkRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetClientToken(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetCommodity(v *UpdateServiceInstanceSpecShrinkRequestCommodity) *UpdateServiceInstanceSpecShrinkRequest {
	s.Commodity = v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetDryRun(v bool) *UpdateServiceInstanceSpecShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetEnableUserPrometheus(v bool) *UpdateServiceInstanceSpecShrinkRequest {
	s.EnableUserPrometheus = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetOperationName(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.OperationName = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetParametersShrink(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetPredefinedParametersName(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.PredefinedParametersName = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetServiceInstanceId(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) Validate() error {
	if s.Commodity != nil {
		if err := s.Commodity.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateServiceInstanceSpecShrinkRequestCommodity struct {
	// Specifies whether to enable automatic payment.
	//
	// Valid values:
	//
	// - **true*	- (default): enables automatic payment. You must ensure that your account balance is sufficient.
	//
	// - **false**: generates an order but does not deduct the fee. If the balance of your payment method is insufficient, you can set the AutoPay parameter to false. In this case, an unpaid order is generated, and you can log on to the [ApsaraDB RDS console](https://rdsnext.console.aliyun.com/dashboard/cn-beijing) to pay the order.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The coupon.
	//
	// example:
	//
	// sxs-xxx
	CouponId *string `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
}

func (s UpdateServiceInstanceSpecShrinkRequestCommodity) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceInstanceSpecShrinkRequestCommodity) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecShrinkRequestCommodity) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *UpdateServiceInstanceSpecShrinkRequestCommodity) GetCouponId() *string {
	return s.CouponId
}

func (s *UpdateServiceInstanceSpecShrinkRequestCommodity) SetAutoPay(v bool) *UpdateServiceInstanceSpecShrinkRequestCommodity {
	s.AutoPay = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequestCommodity) SetCouponId(v string) *UpdateServiceInstanceSpecShrinkRequestCommodity {
	s.CouponId = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequestCommodity) Validate() error {
	return dara.Validate(s)
}
