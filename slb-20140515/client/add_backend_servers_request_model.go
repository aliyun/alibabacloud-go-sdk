// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBackendServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServers(v string) *AddBackendServersRequest
	GetBackendServers() *string
	SetLoadBalancerId(v string) *AddBackendServersRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *AddBackendServersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddBackendServersRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddBackendServersRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AddBackendServersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddBackendServersRequest
	GetResourceOwnerId() *int64
}

type AddBackendServersRequest struct {
	// The list of backend servers that you want to add. Set the following parameters:
	//
	// 	- **ServerId**: Required. This value must be a string. Enter the ID of an ECS instance, elastic network interface (ENI), or elastic container instance. If **ServerId*	- is set to the ID of an ENI or elastic container instance, **Type*	- is required.
	//
	// 	- **Weight**: the weight of the backend server. Valid values: **0*	- to **100**. Default value: **100**.
	//
	//     If the value is set to 0, no requests are forwarded to the backend server.
	//
	// 	- **Description**: Optional. The description of the backend server. This value must be a string. The description must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	//
	// 	- **Type**: the type of the backend server. Valid values:
	//
	//     	- **ecs*	- (default): an ECS instance
	//
	//     	- **eni**: an ENI
	//
	//     	- **eci**: an elastic container instance
	//
	// >  You can specify ENIs and elastic container instances as the backend servers only for high-performance CLB instances.
	//
	// 	- **ServerIp**: the IP address of the ECS instance, ENI, or elastic container instance
	//
	// 	- **Port**: the backend port
	//
	// Examples:
	//
	// 	- ECS instance: `[{ "ServerId": "i-xxxxxxxxx", "Weight": "100", "Type": "ecs", "Port":"80","Description":"test-112" }]`
	//
	// 	- ENI: `[{ "ServerId": "eni-xxxxxxxxx", "Weight": "100", "Type": "eni", "ServerIp": "``192.168.**.**``", "Port":"80","Description":"test-112" }]`
	//
	// 	- ENI with multiple IP addresses: `[{ "ServerId": "eni-xxxxxxxxx", "Weight": "100", "Type": "eni", "ServerIp": "``192.168.**.**``", "Port":"80","Description":"test-113" },{ "ServerId": "eni-xxxxxxxxx", "Weight": "100", "Type": "eni", "ServerIp": "``172.166.**.**``", "Port":"80","Description":"test-113" }]`
	//
	// 	- Elastic container instance: `[{ "ServerId": "eci-xxxxxxxxx", "Weight": "100", "Type": "eci", "ServerIp": "``192.168.**.**``", "Port":"80","Description":"test-114" }]`
	//
	// >  The backend servers that you add to a CLB instance must be in the Running state. You can add at most 20 backend servers to a CLB instance in each request.
	//
	// example:
	//
	// [{ "ServerId": "eni-xxxxxxxxx", "Weight": "100", "Type": "eni", "ServerIp": "``192.168.**.**``", "Port":"80","Description":"test-112" },{ "ServerId": "eni-xxxxxxxxx", "Weight": "100", "Type": "eni", "ServerIp": "``172.166.**.**``", "Port":"80","Description":"test-113" }]
	BackendServers *string `json:"BackendServers,omitempty" xml:"BackendServers,omitempty"`
	// The ID of the CLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-2ze7o5h52g02kkzz******
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the Classic Load Balancer (CLB) instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AddBackendServersRequest) String() string {
	return dara.Prettify(s)
}

func (s AddBackendServersRequest) GoString() string {
	return s.String()
}

func (s *AddBackendServersRequest) GetBackendServers() *string {
	return s.BackendServers
}

func (s *AddBackendServersRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *AddBackendServersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddBackendServersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddBackendServersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddBackendServersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddBackendServersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddBackendServersRequest) SetBackendServers(v string) *AddBackendServersRequest {
	s.BackendServers = &v
	return s
}

func (s *AddBackendServersRequest) SetLoadBalancerId(v string) *AddBackendServersRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *AddBackendServersRequest) SetOwnerAccount(v string) *AddBackendServersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddBackendServersRequest) SetOwnerId(v int64) *AddBackendServersRequest {
	s.OwnerId = &v
	return s
}

func (s *AddBackendServersRequest) SetRegionId(v string) *AddBackendServersRequest {
	s.RegionId = &v
	return s
}

func (s *AddBackendServersRequest) SetResourceOwnerAccount(v string) *AddBackendServersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddBackendServersRequest) SetResourceOwnerId(v int64) *AddBackendServersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddBackendServersRequest) Validate() error {
	return dara.Validate(s)
}
