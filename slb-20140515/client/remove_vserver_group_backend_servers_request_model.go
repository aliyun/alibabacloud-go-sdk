// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveVServerGroupBackendServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServers(v string) *RemoveVServerGroupBackendServersRequest
	GetBackendServers() *string
	SetOwnerAccount(v string) *RemoveVServerGroupBackendServersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RemoveVServerGroupBackendServersRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RemoveVServerGroupBackendServersRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RemoveVServerGroupBackendServersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RemoveVServerGroupBackendServersRequest
	GetResourceOwnerId() *int64
	SetVServerGroupId(v string) *RemoveVServerGroupBackendServersRequest
	GetVServerGroupId() *string
}

type RemoveVServerGroupBackendServersRequest struct {
	// The backend servers that you want to remove. Configure the following parameters:
	//
	// 	- **ServerId**: Required. The ID of the backend server. Specify the value in a string. You can specify the ID of an Elastic Compute Service (ECS) instance, an elastic network interface (ENI), or an elastic container instance. If you set **ServerId*	- to the ID of an ENI or an elastic container instance, you must configure the **Type*	- parameter.
	//
	// 	- **Weight**: the weight of the backend server. Valid values: **0*	- to **100**. Default value: **100**. If you set the weight of a backend server to 0, no requests are forwarded to the backend server.
	//
	// 	- **Description**: Optional. The description of the backend server. Specify the value in a string. The description must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	//
	// 	- **Type**: the type of the backend server. Valid values:
	//
	//     	- **ecs**: ECS instance
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
	// 	- Add ECS instances:
	//
	//     `[{ "ServerId": "i-xxxxxxxxx", "Weight": "100", "Type": "ecs", "Port":"80","Description":"test-112" }]`
	//
	// 	- Add ENIs:
	//
	//     `[{ "ServerId": "eni-xxxxxxxxx", "Weight": "100", "Type": "eni", "ServerIp": "``192.168.**.**``", "Port":"80","Description":"test-112" }]`
	//
	// 	- Add ENIs with multiple IP addresses:
	//
	//     `[{ "ServerId": "eni-xxxxxxxxx", "Weight": "100", "Type": "eni", "ServerIp": "``192.168.**.**``", "Port":"80","Description":"test-113" },{ "ServerId": "eni-xxxxxxxxx", "Weight": "100", "Type": "eni", "ServerIp": "``172.166.**.**``", "Port":"80","Description":"test-113" }]`
	//
	// 	- Add elastic container instances:
	//
	//     `[{ "ServerId": "eci-xxxxxxxxx", "Weight": "100", "Type": "eci", "ServerIp": "``192.168.**.**``", "Port":"80","Description":"test-114" }]`
	//
	// >  You can add only running backend servers to SLB instances. You can specify at most 20 backend servers.
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
	// rsp-cige6****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s RemoveVServerGroupBackendServersRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveVServerGroupBackendServersRequest) GoString() string {
	return s.String()
}

func (s *RemoveVServerGroupBackendServersRequest) GetBackendServers() *string {
	return s.BackendServers
}

func (s *RemoveVServerGroupBackendServersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RemoveVServerGroupBackendServersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RemoveVServerGroupBackendServersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveVServerGroupBackendServersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RemoveVServerGroupBackendServersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RemoveVServerGroupBackendServersRequest) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *RemoveVServerGroupBackendServersRequest) SetBackendServers(v string) *RemoveVServerGroupBackendServersRequest {
	s.BackendServers = &v
	return s
}

func (s *RemoveVServerGroupBackendServersRequest) SetOwnerAccount(v string) *RemoveVServerGroupBackendServersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveVServerGroupBackendServersRequest) SetOwnerId(v int64) *RemoveVServerGroupBackendServersRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveVServerGroupBackendServersRequest) SetRegionId(v string) *RemoveVServerGroupBackendServersRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveVServerGroupBackendServersRequest) SetResourceOwnerAccount(v string) *RemoveVServerGroupBackendServersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveVServerGroupBackendServersRequest) SetResourceOwnerId(v int64) *RemoveVServerGroupBackendServersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveVServerGroupBackendServersRequest) SetVServerGroupId(v string) *RemoveVServerGroupBackendServersRequest {
	s.VServerGroupId = &v
	return s
}

func (s *RemoveVServerGroupBackendServersRequest) Validate() error {
	return dara.Validate(s)
}
