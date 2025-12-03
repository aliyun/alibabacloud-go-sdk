// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetBackendServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServers(v string) *SetBackendServersRequest
	GetBackendServers() *string
	SetLoadBalancerId(v string) *SetBackendServersRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *SetBackendServersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SetBackendServersRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetBackendServersRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *SetBackendServersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetBackendServersRequest
	GetResourceOwnerId() *int64
}

type SetBackendServersRequest struct {
	// The backend servers that you want to add. Configure the following parameters:
	//
	// 	- **ServerId**: Required. The ID of the backend server. Specify the value in a string. You can specify the ID of an Elastic Compute Service (ECS) instance, an elastic network interface (ENI), or an elastic container instance. If you set **ServerId*	- to the ID of an ENI or an elastic container instance, you must configure the **Type*	- parameter.
	//
	// 	- **Weight**: the weight of the backend server. Valid values: **0*	- to **100**. Default value: **100**. If you set the weight of a backend server to 0, no requests are forwarded to the backend server.
	//
	// 	- **Description**: Optional. The description of the backend server. Specify the value in a string. The description must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	//
	// 	- **Type**: the type of the backend server. Valid values:
	//
	//     	- **ecs*	- (default): ECS instance
	//
	//     	- **eni**: ENI
	//
	//     	- **eci**: elastic container instance
	//
	// >  You can specify ENIs and elastic container instances as backend servers only for high-performance CLB instances.
	//
	// 	- **ServerIp**: the IP address of the ENI or elastic container instance.
	//
	// 	- **Port**: the backend port.
	//
	// Examples:
	//
	// 	- ECS instance:
	//
	//     `[{ "ServerId": "i-xxxxxxxxx", "Weight": "100", "Type": "ecs", "Port":"80","Description":"test-112" }]`
	//
	// 	- ENI:
	//
	//     `[{ "ServerId": "eni-xxxxxxxxx", "Weight": "100", "Type": "eni", "ServerIp": "``192.168.**.**``", "Port":"80","Description":"test-112" }]`
	//
	// 	- ENI with multiple IP addresses:
	//
	//     `[{ "ServerId": "eni-xxxxxxxxx", "Weight": "100", "Type": "eni", "ServerIp": "``192.168.**.**``", "Port":"80","Description":"test-113" },{ "ServerId": "eni-xxxxxxxxx", "Weight": "100", "Type": "eni", "ServerIp": "``172.166.**.**``", "Port":"80","Description":"test-113" }]`
	//
	// 	- Elastic container instance:
	//
	//     `[{ "ServerId": "eci-xxxxxxxxx", "Weight": "100", "Type": "eci", "ServerIp": "``192.168.**.**``", "Port":"80","Description":"test-114" }]`
	//
	// >  You can add only running backend servers to a CLB instance. You can specify at most 20 backend servers in each call.
	//
	// example:
	//
	// [{ "ServerId": "ecs-******FmYAXG", "Weight": "100", "Type": "ecs",  "Port":"80","Description":"test-112" }]
	BackendServers *string `json:"BackendServers,omitempty" xml:"BackendServers,omitempty"`
	// The ID of the CLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1qjwo61pqz3a******
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Classic Load Balancer (CLB) instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SetBackendServersRequest) String() string {
	return dara.Prettify(s)
}

func (s SetBackendServersRequest) GoString() string {
	return s.String()
}

func (s *SetBackendServersRequest) GetBackendServers() *string {
	return s.BackendServers
}

func (s *SetBackendServersRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *SetBackendServersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetBackendServersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetBackendServersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetBackendServersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetBackendServersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetBackendServersRequest) SetBackendServers(v string) *SetBackendServersRequest {
	s.BackendServers = &v
	return s
}

func (s *SetBackendServersRequest) SetLoadBalancerId(v string) *SetBackendServersRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetBackendServersRequest) SetOwnerAccount(v string) *SetBackendServersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetBackendServersRequest) SetOwnerId(v int64) *SetBackendServersRequest {
	s.OwnerId = &v
	return s
}

func (s *SetBackendServersRequest) SetRegionId(v string) *SetBackendServersRequest {
	s.RegionId = &v
	return s
}

func (s *SetBackendServersRequest) SetResourceOwnerAccount(v string) *SetBackendServersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetBackendServersRequest) SetResourceOwnerId(v int64) *SetBackendServersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetBackendServersRequest) Validate() error {
	return dara.Validate(s)
}
