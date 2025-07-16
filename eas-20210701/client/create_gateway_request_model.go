// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceName(v string) *CreateGatewayRequest
	GetResourceName() *string
	SetAutoRenewal(v bool) *CreateGatewayRequest
	GetAutoRenewal() *bool
	SetChargeType(v string) *CreateGatewayRequest
	GetChargeType() *string
	SetEnableInternet(v bool) *CreateGatewayRequest
	GetEnableInternet() *bool
	SetEnableIntranet(v bool) *CreateGatewayRequest
	GetEnableIntranet() *bool
	SetInstanceType(v string) *CreateGatewayRequest
	GetInstanceType() *string
	SetName(v string) *CreateGatewayRequest
	GetName() *string
	SetReplicas(v int32) *CreateGatewayRequest
	GetReplicas() *int32
}

type CreateGatewayRequest struct {
	// The resource group ID. To obtain a resource group ID, see the ResourceId field in the response of the [ListResources](https://help.aliyun.com/document_detail/412133.html) operation.
	//
	// example:
	//
	// eas-r-4gt8twzwllfo******
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// 	- false (default)
	//
	// 	- true
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// The billing method. Valid values:
	//
	// 	- PrePaid: subscription.
	//
	// 	- PostPaid: pay-as-you-go.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// Specifies whether to enable Internet access. Default value: false.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	EnableInternet *bool `json:"EnableInternet,omitempty" xml:"EnableInternet,omitempty"`
	// Specifies whether to enable private access. Default value: true.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableIntranet *bool `json:"EnableIntranet,omitempty" xml:"EnableIntranet,omitempty"`
	// The instance type used by the private gateway. Valid values:
	//
	// 	- 2c4g
	//
	// 	- 4c8g
	//
	// 	- 8c16g
	//
	// 	- 16c32g
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs.c6.4xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The alias of the private gateway.
	//
	// example:
	//
	// mygateway1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of nodes in the private gateway.
	//
	// example:
	//
	// 2
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
}

func (s CreateGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *CreateGatewayRequest) GetAutoRenewal() *bool {
	return s.AutoRenewal
}

func (s *CreateGatewayRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateGatewayRequest) GetEnableInternet() *bool {
	return s.EnableInternet
}

func (s *CreateGatewayRequest) GetEnableIntranet() *bool {
	return s.EnableIntranet
}

func (s *CreateGatewayRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateGatewayRequest) GetName() *string {
	return s.Name
}

func (s *CreateGatewayRequest) GetReplicas() *int32 {
	return s.Replicas
}

func (s *CreateGatewayRequest) SetResourceName(v string) *CreateGatewayRequest {
	s.ResourceName = &v
	return s
}

func (s *CreateGatewayRequest) SetAutoRenewal(v bool) *CreateGatewayRequest {
	s.AutoRenewal = &v
	return s
}

func (s *CreateGatewayRequest) SetChargeType(v string) *CreateGatewayRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateGatewayRequest) SetEnableInternet(v bool) *CreateGatewayRequest {
	s.EnableInternet = &v
	return s
}

func (s *CreateGatewayRequest) SetEnableIntranet(v bool) *CreateGatewayRequest {
	s.EnableIntranet = &v
	return s
}

func (s *CreateGatewayRequest) SetInstanceType(v string) *CreateGatewayRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateGatewayRequest) SetName(v string) *CreateGatewayRequest {
	s.Name = &v
	return s
}

func (s *CreateGatewayRequest) SetReplicas(v int32) *CreateGatewayRequest {
	s.Replicas = &v
	return s
}

func (s *CreateGatewayRequest) Validate() error {
	return dara.Validate(s)
}
