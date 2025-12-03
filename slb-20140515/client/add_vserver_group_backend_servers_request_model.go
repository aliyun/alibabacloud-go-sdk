// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVServerGroupBackendServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServers(v string) *AddVServerGroupBackendServersRequest
	GetBackendServers() *string
	SetOwnerAccount(v string) *AddVServerGroupBackendServersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddVServerGroupBackendServersRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddVServerGroupBackendServersRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AddVServerGroupBackendServersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddVServerGroupBackendServersRequest
	GetResourceOwnerId() *int64
	SetVServerGroupId(v string) *AddVServerGroupBackendServersRequest
	GetVServerGroupId() *string
}

type AddVServerGroupBackendServersRequest struct {
	// The backend servers that you want to add. Configure the following parameters:
	//
	// 	- **ServerId**: Required. The ID of the backend server. Specify the ID in a string. You can specify the ID of an Elastic Compute Service (ECS) instance, an elastic network interface (ENI), and an elastic container instance. If you set **ServerId*	- to the ID of an ENI or an elastic container instance, you must configure the **Type*	- parameter.
	//
	// 	- **Weight**: the weight of the backend server. Valid values: **0*	- to **100**. Default value: **100**. If you set the weight of a backend server to 0, no requests are forwarded to the backend server.
	//
	// 	- **Description**: Optional. The description of the backend server. Specify the description in a string. The description must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	//
	// 	- **Type**: the type of the backend server. Valid values:
	//
	//     	- **ecs*	- (default): ECS instance
	//
	//     	- **eni**: ENI
	//
	//     	- **eci**: elastic container instance
	//
	// >  You can specify ENIs and elastic container instances as backend servers only for high-performance SLB instances.
	//
	// 	- **ServerIp**: the IP address of an ENI or an elastic container instance.
	//
	// 	- **Port**: the backend port.
	//
	// Examples:
	//
	// 	- Add an ECS instance:
	//
	//     `[{ "ServerId": "i-xxxxxxxxx", "Weight": "100", "Type": "ecs", "Port":"80","Description":"test-112" }]`
	//
	// 	- Add an ENI:
	//
	//     `[{ "ServerId": "eni-xxxxxxxxx", "Weight": "100", "Type": "eni", "ServerIp": "``192.168.**.**``", "Port":"80","Description":"test-112" }]`
	//
	// 	- Add an ENI with multiple IP addresses:
	//
	//     `[{ "ServerId": "eni-xxxxxxxxx", "Weight": "100", "Type": "eni", "ServerIp": "``192.168.**.**``", "Port":"80","Description":"test-113" },{ "ServerId": "eni-xxxxxxxxx", "Weight": "100", "Type": "eni", "ServerIp": "``172.166.**.**``", "Port":"80","Description":"test-113" }]`
	//
	// 	- Add an elastic container instance:
	//
	//     `[{ "ServerId": "eci-xxxxxxxxx", "Weight": "100", "Type": "eci", "ServerIp": "``192.168.**.**``", "Port":"80","Description":"test-114" }]`
	//
	// >  You can add only running backend servers to SLB instances. You can specify at most 20 backend servers in each call.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{ "ServerId": "eni-xxxxxxxxx", "Weight": "100", "Type": "eni", "ServerIp": "``192.168.**.**``", "Port":"80","Description":"test-112" },{ "ServerId": "eni-xxxxxxxxx", "Weight": "100", "Type": "eni", "ServerIp": "``172.166.**.**``", "Port":"80","Description":"test-113" }]
	BackendServers *string `json:"BackendServers,omitempty" xml:"BackendServers,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Server Load Balancer (SLB) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the vServer group.
	//
	// This parameter is required.
	//
	// example:
	//
	// rsp-cige6******
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s AddVServerGroupBackendServersRequest) String() string {
	return dara.Prettify(s)
}

func (s AddVServerGroupBackendServersRequest) GoString() string {
	return s.String()
}

func (s *AddVServerGroupBackendServersRequest) GetBackendServers() *string {
	return s.BackendServers
}

func (s *AddVServerGroupBackendServersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddVServerGroupBackendServersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddVServerGroupBackendServersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddVServerGroupBackendServersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddVServerGroupBackendServersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddVServerGroupBackendServersRequest) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *AddVServerGroupBackendServersRequest) SetBackendServers(v string) *AddVServerGroupBackendServersRequest {
	s.BackendServers = &v
	return s
}

func (s *AddVServerGroupBackendServersRequest) SetOwnerAccount(v string) *AddVServerGroupBackendServersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddVServerGroupBackendServersRequest) SetOwnerId(v int64) *AddVServerGroupBackendServersRequest {
	s.OwnerId = &v
	return s
}

func (s *AddVServerGroupBackendServersRequest) SetRegionId(v string) *AddVServerGroupBackendServersRequest {
	s.RegionId = &v
	return s
}

func (s *AddVServerGroupBackendServersRequest) SetResourceOwnerAccount(v string) *AddVServerGroupBackendServersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddVServerGroupBackendServersRequest) SetResourceOwnerId(v int64) *AddVServerGroupBackendServersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddVServerGroupBackendServersRequest) SetVServerGroupId(v string) *AddVServerGroupBackendServersRequest {
	s.VServerGroupId = &v
	return s
}

func (s *AddVServerGroupBackendServersRequest) Validate() error {
	return dara.Validate(s)
}
