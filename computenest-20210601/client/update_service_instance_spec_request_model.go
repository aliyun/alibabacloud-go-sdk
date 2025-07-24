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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The information about the order placed in Alibaba Cloud Marketplace. You do not need to specify this parameter if the service is not published in Alibaba Cloud Marketplace or uses the pay-as-you-go billing method.
	Commodity *UpdateServiceInstanceSpecRequestCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
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
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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
	return dara.Validate(s)
}

type UpdateServiceInstanceSpecRequestCommodity struct {
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

func (s UpdateServiceInstanceSpecRequestCommodity) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceInstanceSpecRequestCommodity) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecRequestCommodity) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *UpdateServiceInstanceSpecRequestCommodity) SetAutoPay(v bool) *UpdateServiceInstanceSpecRequestCommodity {
	s.AutoPay = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequestCommodity) Validate() error {
	return dara.Validate(s)
}
