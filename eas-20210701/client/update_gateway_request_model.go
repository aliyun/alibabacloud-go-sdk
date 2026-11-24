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
	SetIdleTimeout(v int32) *UpdateGatewayRequest
	GetIdleTimeout() *int32
	SetInstanceType(v string) *UpdateGatewayRequest
	GetInstanceType() *string
	SetIsDefault(v bool) *UpdateGatewayRequest
	GetIsDefault() *bool
	SetName(v string) *UpdateGatewayRequest
	GetName() *string
	SetReplicas(v int32) *UpdateGatewayRequest
	GetReplicas() *int32
	SetRequestTimeout(v int32) *UpdateGatewayRequest
	GetRequestTimeout() *int32
	SetVSwitchIds(v []*string) *UpdateGatewayRequest
	GetVSwitchIds() []*string
	SetVpcId(v string) *UpdateGatewayRequest
	GetVpcId() *string
	SetWorkspaceId(v string) *UpdateGatewayRequest
	GetWorkspaceId() *string
}

type UpdateGatewayRequest struct {
	// Specifies whether to enable public network access. Default value: false.
	//
	// example:
	//
	// false
	EnableInternet *bool `json:"EnableInternet,omitempty" xml:"EnableInternet,omitempty"`
	// Specifies whether to enable internal network access. Default value: true.
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
	// The idle connection timeout period. Unit: seconds. Default value: 600. Valid values: 1 to 3600.
	//
	// example:
	//
	// 600
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The instance type used by the private gateway. Valid values:
	//
	// - 2c4g
	//
	// - 4c8g
	//
	// - 8c16g
	//
	// - 16c32g
	//
	// example:
	//
	// 2c4g
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Specifies whether the private gateway is the default private gateway.
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
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The request timeout period. Unit: seconds. Default value: 600. Valid values: 1 to 3600.
	//
	// example:
	//
	// 600
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// The list of vSwitches. This parameter applies only to dedicated gateways for applications.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The ID of the VPC to which the gateway belongs. This parameter applies only to dedicated gateways for applications.
	//
	// example:
	//
	// vpc-bp1jkde2******3mew
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 123456
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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

func (s *UpdateGatewayRequest) GetIdleTimeout() *int32 {
	return s.IdleTimeout
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

func (s *UpdateGatewayRequest) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *UpdateGatewayRequest) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *UpdateGatewayRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateGatewayRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
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

func (s *UpdateGatewayRequest) SetIdleTimeout(v int32) *UpdateGatewayRequest {
	s.IdleTimeout = &v
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

func (s *UpdateGatewayRequest) SetRequestTimeout(v int32) *UpdateGatewayRequest {
	s.RequestTimeout = &v
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

func (s *UpdateGatewayRequest) SetWorkspaceId(v string) *UpdateGatewayRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateGatewayRequest) Validate() error {
	return dara.Validate(s)
}
