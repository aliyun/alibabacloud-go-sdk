// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVServerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServers(v string) *CreateVServerGroupRequest
	GetBackendServers() *string
	SetLoadBalancerId(v string) *CreateVServerGroupRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *CreateVServerGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateVServerGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateVServerGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateVServerGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateVServerGroupRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateVServerGroupRequestTag) *CreateVServerGroupRequest
	GetTag() []*CreateVServerGroupRequestTag
	SetVServerGroupName(v string) *CreateVServerGroupRequest
	GetVServerGroupName() *string
}

type CreateVServerGroupRequest struct {
	// The backend servers that you want to add. Configure the following parameters:
	//
	// 	- **ServerId**:  required. The ID of the backend server. Specify the ID in a string. You can specify the ID of an Elastic Compute Service (ECS) instance, an elastic network interface (ENI), or an elastic container instance. If you set ServerId to the ID of an ENI or an elastic container instance, you must configure the Type parameter.
	//
	// 	- **Weight**: the weight of the backend server. Valid values: 0 to 100. Default value: 100. If you set the weight of a backend server to 0, no requests are forwarded to the backend server.
	//
	// 	- **Description**: optional. The description of the backend server. Specify the description in a string. The description must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	//
	// 	- **Type**: the type of the backend server. Valid values:
	//
	//     	- **ecs (default)**: ECS instance
	//
	//     	- **eni**: ENI.
	//
	//     	- **eni**: elastic container instance.
	//
	// > You can specify ENIs and elastic container instances as backend servers only for high-performance SLB instances.
	//
	// 	- **ServerIp**: The IP address of the ECS instance or ENI.
	//
	// 	- **Port**: the backend port.
	//
	// Examples:
	//
	// - Add an ECS instance:
	//
	//   `[{ "ServerId": "i-xxxxxxxxx", "Weight": "100", "Type": "ecs", "Port":"80","Description":"test-112" }]`
	//
	// - Add an ENI:
	//
	//   ` [{ "ServerId": "eni-xxxxxxxxx", "Weight": "100", "Type": "eni", "ServerIp": "``192.168.**.**``", "Port":"80","Description":"test-112" }]`
	//
	// - Add an ENI with multiple IP addresses:
	//
	//    `[{ "ServerId": "eni-xxxxxxxxx", "Weight": "100", "Type": "eni", "ServerIp": "``192.168.**.**``", "Port":"80","Description":"test-113" },{ "ServerId": "eni-xxxxxxxxx", "Weight": "100", "Type": "eni", "ServerIp": "``172.166.**.**``", "Port":"80","Description":"test-113" }]`
	//
	// - Add an elastic container instance:
	//
	//   ` [{ "ServerId": "eci-xxxxxxxxx", "Weight": "100", "Type": "eci", "ServerIp": "``192.168.**.**``", "Port":"80","Description":"test-114" }]`
	//
	// > You can add only running backend servers to SLB instances. You can specify at most 20 backend servers.
	//
	// example:
	//
	// [{ "ServerId": "eni-xxxxxxxxx", "Weight": "100", "Type": "eni", "ServerIp": "``192.168.**.**``", "Port":"80","Description":"test-112" },{ "ServerId": "eni-xxxxxxxxx", "Weight": "100", "Type": "eni", "ServerIp": "``172.166.**.**``", "Port":"80","Description":"test-113" }]
	BackendServers *string `json:"BackendServers,omitempty" xml:"BackendServers,omitempty"`
	// The ID of the Server Load Balancer (SLB) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1qjwo61pqz3ahl******
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SLB instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	Tag []*CreateVServerGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the vServer group.
	//
	// The name must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.),and underscores (_).
	//
	// example:
	//
	// Group1
	VServerGroupName *string `json:"VServerGroupName,omitempty" xml:"VServerGroupName,omitempty"`
}

func (s CreateVServerGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVServerGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateVServerGroupRequest) GetBackendServers() *string {
	return s.BackendServers
}

func (s *CreateVServerGroupRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateVServerGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateVServerGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateVServerGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVServerGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateVServerGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateVServerGroupRequest) GetTag() []*CreateVServerGroupRequestTag {
	return s.Tag
}

func (s *CreateVServerGroupRequest) GetVServerGroupName() *string {
	return s.VServerGroupName
}

func (s *CreateVServerGroupRequest) SetBackendServers(v string) *CreateVServerGroupRequest {
	s.BackendServers = &v
	return s
}

func (s *CreateVServerGroupRequest) SetLoadBalancerId(v string) *CreateVServerGroupRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateVServerGroupRequest) SetOwnerAccount(v string) *CreateVServerGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateVServerGroupRequest) SetOwnerId(v int64) *CreateVServerGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVServerGroupRequest) SetRegionId(v string) *CreateVServerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVServerGroupRequest) SetResourceOwnerAccount(v string) *CreateVServerGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateVServerGroupRequest) SetResourceOwnerId(v int64) *CreateVServerGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateVServerGroupRequest) SetTag(v []*CreateVServerGroupRequestTag) *CreateVServerGroupRequest {
	s.Tag = v
	return s
}

func (s *CreateVServerGroupRequest) SetVServerGroupName(v string) *CreateVServerGroupRequest {
	s.VServerGroupName = &v
	return s
}

func (s *CreateVServerGroupRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateVServerGroupRequestTag struct {
	// The key of tag N. Valid values of N: **1 to 20**. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length, and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. Valid values of N: **1 to 20**. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag value cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVServerGroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateVServerGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateVServerGroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateVServerGroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateVServerGroupRequestTag) SetKey(v string) *CreateVServerGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateVServerGroupRequestTag) SetValue(v string) *CreateVServerGroupRequestTag {
	s.Value = &v
	return s
}

func (s *CreateVServerGroupRequestTag) Validate() error {
	return dara.Validate(s)
}
