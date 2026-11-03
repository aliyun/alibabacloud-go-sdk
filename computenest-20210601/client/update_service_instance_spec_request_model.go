// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceInstanceSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateServiceInstanceSpecRequest
	GetClientToken() *string
	SetCommodity(v *UpdateServiceInstanceSpecRequestCommodity) *UpdateServiceInstanceSpecRequest
	GetCommodity() *UpdateServiceInstanceSpecRequestCommodity
	SetDryRun(v bool) *UpdateServiceInstanceSpecRequest
	GetDryRun() *bool
	SetEnableUserPrometheus(v bool) *UpdateServiceInstanceSpecRequest
	GetEnableUserPrometheus() *bool
	SetOperationName(v string) *UpdateServiceInstanceSpecRequest
	GetOperationName() *string
	SetParameters(v map[string]interface{}) *UpdateServiceInstanceSpecRequest
	GetParameters() map[string]interface{}
	SetPredefinedParametersName(v string) *UpdateServiceInstanceSpecRequest
	GetPredefinedParametersName() *string
	SetServiceInstanceId(v string) *UpdateServiceInstanceSpecRequest
	GetServiceInstanceId() *string
}

type UpdateServiceInstanceSpecRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can generate the value from your client, but you must ensure that the value is unique among different requests. The ClientToken parameter supports only ASCII characters and can be up to 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The purchase order information for Alibaba Cloud Marketplace. This parameter is not required if the service is not listed in Alibaba Cloud Marketplace or uses the pay-as-you-go billing method.
	Commodity *UpdateServiceInstanceSpecRequestCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
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
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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

func (s UpdateServiceInstanceSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateServiceInstanceSpecRequest) GetCommodity() *UpdateServiceInstanceSpecRequestCommodity {
	return s.Commodity
}

func (s *UpdateServiceInstanceSpecRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateServiceInstanceSpecRequest) GetEnableUserPrometheus() *bool {
	return s.EnableUserPrometheus
}

func (s *UpdateServiceInstanceSpecRequest) GetOperationName() *string {
	return s.OperationName
}

func (s *UpdateServiceInstanceSpecRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *UpdateServiceInstanceSpecRequest) GetPredefinedParametersName() *string {
	return s.PredefinedParametersName
}

func (s *UpdateServiceInstanceSpecRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *UpdateServiceInstanceSpecRequest) SetClientToken(v string) *UpdateServiceInstanceSpecRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetCommodity(v *UpdateServiceInstanceSpecRequestCommodity) *UpdateServiceInstanceSpecRequest {
	s.Commodity = v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetDryRun(v bool) *UpdateServiceInstanceSpecRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetEnableUserPrometheus(v bool) *UpdateServiceInstanceSpecRequest {
	s.EnableUserPrometheus = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetOperationName(v string) *UpdateServiceInstanceSpecRequest {
	s.OperationName = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetParameters(v map[string]interface{}) *UpdateServiceInstanceSpecRequest {
	s.Parameters = v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetPredefinedParametersName(v string) *UpdateServiceInstanceSpecRequest {
	s.PredefinedParametersName = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetServiceInstanceId(v string) *UpdateServiceInstanceSpecRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) Validate() error {
	if s.Commodity != nil {
		if err := s.Commodity.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateServiceInstanceSpecRequestCommodity struct {
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

func (s UpdateServiceInstanceSpecRequestCommodity) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceInstanceSpecRequestCommodity) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecRequestCommodity) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *UpdateServiceInstanceSpecRequestCommodity) GetCouponId() *string {
	return s.CouponId
}

func (s *UpdateServiceInstanceSpecRequestCommodity) SetAutoPay(v bool) *UpdateServiceInstanceSpecRequestCommodity {
	s.AutoPay = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequestCommodity) SetCouponId(v string) *UpdateServiceInstanceSpecRequestCommodity {
	s.CouponId = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequestCommodity) Validate() error {
	return dara.Validate(s)
}
