// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetVServerGroupAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServers(v string) *SetVServerGroupAttributeRequest
	GetBackendServers() *string
	SetOwnerAccount(v string) *SetVServerGroupAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SetVServerGroupAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetVServerGroupAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *SetVServerGroupAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetVServerGroupAttributeRequest
	GetResourceOwnerId() *int64
	SetVServerGroupId(v string) *SetVServerGroupAttributeRequest
	GetVServerGroupId() *string
	SetVServerGroupName(v string) *SetVServerGroupAttributeRequest
	GetVServerGroupName() *string
}

type SetVServerGroupAttributeRequest struct {
	// The backend servers. This operation only can be used to modify the weights of backend servers and names of vServer groups. Configure the following parameters:
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
	// >  You can add only running backend servers to SLB instances. You can specify at most 20 backend servers in each call.
	//
	// example:
	//
	// [{ "ServerId": "eni-xxxxxxxxx", "Weight": "100", "Type": "eni", "ServerIp": "192.XX.XX.6", "Port":"80","Description":"test-112" },{ "ServerId": "eni-xxxxxxxxx", "Weight": "100", "Type": "eni", "ServerIp": "172.XX.XX.6", "Port":"80","Description":"test-113" }]
	BackendServers *string `json:"BackendServers,omitempty" xml:"BackendServers,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Server Load Balancer (SLB) instance, which cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The vServer group ID, which cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// rsp-cige6****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	// The vServer group name. You can specify a name.
	//
	// example:
	//
	// Group1
	VServerGroupName *string `json:"VServerGroupName,omitempty" xml:"VServerGroupName,omitempty"`
}

func (s SetVServerGroupAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s SetVServerGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *SetVServerGroupAttributeRequest) GetBackendServers() *string {
	return s.BackendServers
}

func (s *SetVServerGroupAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetVServerGroupAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetVServerGroupAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetVServerGroupAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetVServerGroupAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetVServerGroupAttributeRequest) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *SetVServerGroupAttributeRequest) GetVServerGroupName() *string {
	return s.VServerGroupName
}

func (s *SetVServerGroupAttributeRequest) SetBackendServers(v string) *SetVServerGroupAttributeRequest {
	s.BackendServers = &v
	return s
}

func (s *SetVServerGroupAttributeRequest) SetOwnerAccount(v string) *SetVServerGroupAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetVServerGroupAttributeRequest) SetOwnerId(v int64) *SetVServerGroupAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *SetVServerGroupAttributeRequest) SetRegionId(v string) *SetVServerGroupAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *SetVServerGroupAttributeRequest) SetResourceOwnerAccount(v string) *SetVServerGroupAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetVServerGroupAttributeRequest) SetResourceOwnerId(v int64) *SetVServerGroupAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetVServerGroupAttributeRequest) SetVServerGroupId(v string) *SetVServerGroupAttributeRequest {
	s.VServerGroupId = &v
	return s
}

func (s *SetVServerGroupAttributeRequest) SetVServerGroupName(v string) *SetVServerGroupAttributeRequest {
	s.VServerGroupName = &v
	return s
}

func (s *SetVServerGroupAttributeRequest) Validate() error {
	return dara.Validate(s)
}
