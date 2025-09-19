// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableInternet(v bool) *UpdateGatewayRequest
	GetEnableInternet() *bool
	SetEnableIntranet(v bool) *UpdateGatewayRequest
	GetEnableIntranet() *bool
	SetEnableSSLRedirection(v bool) *UpdateGatewayRequest
	GetEnableSSLRedirection() *bool
	SetInstanceType(v string) *UpdateGatewayRequest
	GetInstanceType() *string
	SetIsDefault(v bool) *UpdateGatewayRequest
	GetIsDefault() *bool
	SetName(v string) *UpdateGatewayRequest
	GetName() *string
	SetReplicas(v int32) *UpdateGatewayRequest
	GetReplicas() *int32
	SetVSwitchIds(v []*string) *UpdateGatewayRequest
	GetVSwitchIds() []*string
	SetVpcId(v string) *UpdateGatewayRequest
	GetVpcId() *string
}

type UpdateGatewayRequest struct {
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
	// Specifies whether to enable HTTP to HTTPS redirection. Default value: false.
	//
	// example:
	//
	// false
	EnableSSLRedirection *bool `json:"EnableSSLRedirection,omitempty" xml:"EnableSSLRedirection,omitempty"`
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
	// example:
	//
	// ecs.c6.4xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Specifies whether it is the default private gateway.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
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
	Replicas   *int32    `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	VpcId      *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UpdateGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRequest) GetEnableInternet() *bool {
	return s.EnableInternet
}

func (s *UpdateGatewayRequest) GetEnableIntranet() *bool {
	return s.EnableIntranet
}

func (s *UpdateGatewayRequest) GetEnableSSLRedirection() *bool {
	return s.EnableSSLRedirection
}

func (s *UpdateGatewayRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *UpdateGatewayRequest) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *UpdateGatewayRequest) GetName() *string {
	return s.Name
}

func (s *UpdateGatewayRequest) GetReplicas() *int32 {
	return s.Replicas
}

func (s *UpdateGatewayRequest) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *UpdateGatewayRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateGatewayRequest) SetEnableInternet(v bool) *UpdateGatewayRequest {
	s.EnableInternet = &v
	return s
}

func (s *UpdateGatewayRequest) SetEnableIntranet(v bool) *UpdateGatewayRequest {
	s.EnableIntranet = &v
	return s
}

func (s *UpdateGatewayRequest) SetEnableSSLRedirection(v bool) *UpdateGatewayRequest {
	s.EnableSSLRedirection = &v
	return s
}

func (s *UpdateGatewayRequest) SetInstanceType(v string) *UpdateGatewayRequest {
	s.InstanceType = &v
	return s
}

func (s *UpdateGatewayRequest) SetIsDefault(v bool) *UpdateGatewayRequest {
	s.IsDefault = &v
	return s
}

func (s *UpdateGatewayRequest) SetName(v string) *UpdateGatewayRequest {
	s.Name = &v
	return s
}

func (s *UpdateGatewayRequest) SetReplicas(v int32) *UpdateGatewayRequest {
	s.Replicas = &v
	return s
}

func (s *UpdateGatewayRequest) SetVSwitchIds(v []*string) *UpdateGatewayRequest {
	s.VSwitchIds = v
	return s
}

func (s *UpdateGatewayRequest) SetVpcId(v string) *UpdateGatewayRequest {
	s.VpcId = &v
	return s
}

func (s *UpdateGatewayRequest) Validate() error {
	return dara.Validate(s)
}
