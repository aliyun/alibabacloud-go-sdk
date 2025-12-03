// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVServerGroupBackendServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNewBackendServers(v string) *ModifyVServerGroupBackendServersRequest
	GetNewBackendServers() *string
	SetOldBackendServers(v string) *ModifyVServerGroupBackendServersRequest
	GetOldBackendServers() *string
	SetOwnerAccount(v string) *ModifyVServerGroupBackendServersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyVServerGroupBackendServersRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyVServerGroupBackendServersRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyVServerGroupBackendServersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyVServerGroupBackendServersRequest
	GetResourceOwnerId() *int64
	SetVServerGroupId(v string) *ModifyVServerGroupBackendServersRequest
	GetVServerGroupId() *string
}

type ModifyVServerGroupBackendServersRequest struct {
	// The backend servers that you want to add to the vServer group. Configure the following parameters:
	//
	// 	- **ServerId**: required. The IDs of the backend servers. Specify the IDs in a string. You can specify the IDs of ECS instances, ENIs, and elastic container instances. If you set **ServerId*	- to the IDs of ENIs or elastic container instances, you must configure the **Type*	- parameter.
	//
	// 	- **Weight**: the weight of the backend server. Valid values: **0*	- to **100**. Default value: **100**. If you set the weight of a backend server to 0, no requests are forwarded to the backend server.
	//
	// 	- **Description**: optional. The description of the backend servers. Specify the description in a string. The description must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/). periods (.), and underscores (_).
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
	// 	- **ServerIp**: the IP address of the ENI or elastic container instance.
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
	// 	- Add an elastic container instance
	//
	//     `[{ "ServerId": "eci-xxxxxxxxx", "Weight": "100", "Type": "eci", "ServerIp": "``192.168.**.**``", "Port":"80","Description":"test-114" }]`
	//
	// >  You can add only running backend servers to SLB instances. You can specify at most 20 backend servers in each call.
	//
	// example:
	//
	// [{ "ServerId": "eni-xxxxxxxxx", "Weight": "100", "Type": "eni", "ServerIp": "``192.168.**.**``", "Port":"80","Description":"test-112" },{ "ServerId": "eni-xxxxxxxxx", "Weight": "100", "Type": "eni", "ServerIp": "``172.166.**.**``", "Port":"80","Description":"test-113" }]
	NewBackendServers *string `json:"NewBackendServers,omitempty" xml:"NewBackendServers,omitempty"`
	// The backend servers that you want to replace. Configure the following parameters:
	//
	// 	- **ServerId**: required. The IDs of the backend servers. Specify the IDs in a string. You can specify the IDs of Elastic Compute Service (ECS) instances, elastic network interfaces (ENIs), and elastic container instances. If you set **ServerId*	- to the IDs of ENIs or elastic container instances, you must configure the **Type*	- parameter.
	//
	// 	- **Weight**: the weight of the backend server. Valid values: **0*	- to **100**. Default value: **100**. If you set the weight of a backend server to 0, no requests are forwarded to the backend server.
	//
	// 	- **Description**: optional. The description of the backend servers. Specify the description in a string. The description must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/). periods (.), and underscores (_).
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
	// 	- **ServerIp**: the IP address of the ENI or elastic container instance.
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
	// 	- Add an elastic container instance
	//
	//     `[{ "ServerId": "eci-xxxxxxxxx", "Weight": "100", "Type": "eci", "ServerIp": "``192.168.**.**``", "Port":"80","Description":"test-114" }]`
	//
	// >  You can add only running backend servers to SLB instances. You can specify at most 20 backend servers in each call.
	//
	// example:
	//
	// [{ "ServerId": "i-xxxxxxxxx", "Weight": "100", "Type": "ecs",  "Port":"80","Description":"test-112" }]
	OldBackendServers *string `json:"OldBackendServers,omitempty" xml:"OldBackendServers,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the Classic Load Balancer (CLB) instance is deployed.
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
	// rsp-cige6j****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s ModifyVServerGroupBackendServersRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVServerGroupBackendServersRequest) GoString() string {
	return s.String()
}

func (s *ModifyVServerGroupBackendServersRequest) GetNewBackendServers() *string {
	return s.NewBackendServers
}

func (s *ModifyVServerGroupBackendServersRequest) GetOldBackendServers() *string {
	return s.OldBackendServers
}

func (s *ModifyVServerGroupBackendServersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyVServerGroupBackendServersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyVServerGroupBackendServersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyVServerGroupBackendServersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyVServerGroupBackendServersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyVServerGroupBackendServersRequest) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *ModifyVServerGroupBackendServersRequest) SetNewBackendServers(v string) *ModifyVServerGroupBackendServersRequest {
	s.NewBackendServers = &v
	return s
}

func (s *ModifyVServerGroupBackendServersRequest) SetOldBackendServers(v string) *ModifyVServerGroupBackendServersRequest {
	s.OldBackendServers = &v
	return s
}

func (s *ModifyVServerGroupBackendServersRequest) SetOwnerAccount(v string) *ModifyVServerGroupBackendServersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyVServerGroupBackendServersRequest) SetOwnerId(v int64) *ModifyVServerGroupBackendServersRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyVServerGroupBackendServersRequest) SetRegionId(v string) *ModifyVServerGroupBackendServersRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyVServerGroupBackendServersRequest) SetResourceOwnerAccount(v string) *ModifyVServerGroupBackendServersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyVServerGroupBackendServersRequest) SetResourceOwnerId(v int64) *ModifyVServerGroupBackendServersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyVServerGroupBackendServersRequest) SetVServerGroupId(v string) *ModifyVServerGroupBackendServersRequest {
	s.VServerGroupId = &v
	return s
}

func (s *ModifyVServerGroupBackendServersRequest) Validate() error {
	return dara.Validate(s)
}
