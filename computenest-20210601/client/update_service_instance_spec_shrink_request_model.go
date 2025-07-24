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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The information about the order placed in Alibaba Cloud Marketplace. You do not need to specify this parameter if the service is not published in Alibaba Cloud Marketplace or uses the pay-as-you-go billing method.
	Commodity *UpdateServiceInstanceSpecShrinkRequestCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// Specifies whether to perform only a dry run, without performing the actual request. A dry run includes checks on the permissions and instance state.
	//
	// Valid values:
	//
	// 	- true: performs a dry run but does not create a service instance.
	//
	// 	- false: performs a dry run and creates a service instance if the request passes the dry run.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to enable Prometheus monitoring on the user side.
	//
	// Valid values:
	//
	// true
	//
	// false
	//
	// example:
	//
	// true
	EnableUserPrometheus *bool `json:"EnableUserPrometheus,omitempty" xml:"EnableUserPrometheus,omitempty"`
	// The name of the configuration change operation.
	//
	// To obtain the names and content of the configuration change operations of the service, you can call the [GetService](https://help.aliyun.com/document_detail/2340828.html) operation. In the response, check the value of **ModifyParametersConfig*	- in the value of **OperationMetadata**.
	//
	// example:
	//
	// package modify
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// The configuration parameter.
	//
	// This parameter is available if the service provider set **Method*	- to **Change Parameter*	- when configuring configuration change operations.
	//
	// >
	//
	// 	- To obtain the parameters of the service that support configuration change, you can call the [GetService](https://help.aliyun.com/document_detail/2340828.html) operation. In the response, check the value of **ModifyParametersConfig*	- in the value of **OperationMetadata**.
	//
	// 	- You can also view the parameters of the service that support configuration change in the **configuration change*	- dialog box in the [Compute Nest console](https://computenest.console.aliyun.com/service/instance/cn-hangzhou).
	//
	// For example, if the service supports Elastic Compute Service (ECS) instance type upgrade, you must specify an instance type that has higher specifications than the current one.
	//
	// example:
	//
	// {
	//
	//   "InstanceType": "ecs.g8ise.2xlarge"
	//
	// }
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The name of the configuration plan.
	//
	// This parameter is available if the service provider set **Method*	- to **Change Plan*	- when configuring configuration change operations.
	//
	// To obtain the configuration plan names of the service, you can call the [GetService](https://help.aliyun.com/document_detail/2340828.html) operation. In the response, check the value of **PredefinedParameters*	- in the value of **DeployMetadata**.
	//
	// example:
	//
	// package One
	PredefinedParametersName *string `json:"PredefinedParametersName,omitempty" xml:"PredefinedParametersName,omitempty"`
	// The ID of the service instance.
	//
	// You can call the [ListServiceInstances](https://help.aliyun.com/document_detail/396200.html) operation to obtain the ID of the service instance.
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
	return dara.Validate(s)
}

type UpdateServiceInstanceSpecShrinkRequestCommodity struct {
	// Specifies whether to enable automatic payment.
	//
	// Valid values:
	//
	// 	- **true (default)**: automatically completes the payment. You must make sure that your account balance is sufficient.
	//
	// 	- **false**: does not automatically complete the payment. An unpaid order is generated. If your account balance is insufficient, you can set AutoPay to false. In this case, an unpaid order is generated. You can complete the payment in the Expenses and Costs console.[](https://rdsnext.console.aliyun.com/dashboard/cn-beijing)
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
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

func (s *UpdateServiceInstanceSpecShrinkRequestCommodity) SetAutoPay(v bool) *UpdateServiceInstanceSpecShrinkRequestCommodity {
	s.AutoPay = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequestCommodity) Validate() error {
	return dara.Validate(s)
}
