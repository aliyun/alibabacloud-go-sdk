// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMasterSlaveServerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancerId(v string) *CreateMasterSlaveServerGroupRequest
	GetLoadBalancerId() *string
	SetMasterSlaveBackendServers(v string) *CreateMasterSlaveServerGroupRequest
	GetMasterSlaveBackendServers() *string
	SetMasterSlaveServerGroupName(v string) *CreateMasterSlaveServerGroupRequest
	GetMasterSlaveServerGroupName() *string
	SetOwnerAccount(v string) *CreateMasterSlaveServerGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateMasterSlaveServerGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateMasterSlaveServerGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateMasterSlaveServerGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateMasterSlaveServerGroupRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateMasterSlaveServerGroupRequestTag) *CreateMasterSlaveServerGroupRequest
	GetTag() []*CreateMasterSlaveServerGroupRequestTag
}

type CreateMasterSlaveServerGroupRequest struct {
	// The CLB instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1hv944r69al4j******
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The backend servers in the primary/secondary server group. Each primary/secondary server group consists of two backend servers.
	//
	// Configure the following parameters:
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
	// >  You can specify ENIs and elastic container instances as backend servers only for high-performance CLB instances.
	//
	// 	- **ServerIp**: the IP address of the ENI or elastic container instance.
	//
	// 	- **Port**: the backend port.
	//
	// 	- **ServerType**: Specify the primary and secondary backend servers in a string. Valid values:
	//
	//     	- **Master**: primary server
	//
	//     	- **Slave**: secondary server
	//
	// example:
	//
	// [{ "ServerId": "i-xxxxxxxxx", "Weight": "100", "Type": "ecs",  "Port":"82","ServerType":"Master","Description":"test-112" },  { "ServerId": "i-xxxxxxxxx", "Weight": "100", "Type": "ecs",  "Port":"84","ServerType":"Slave","Description":"test-112" }]
	MasterSlaveBackendServers *string `json:"MasterSlaveBackendServers,omitempty" xml:"MasterSlaveBackendServers,omitempty"`
	// The name of the primary/secondary server group.
	//
	// example:
	//
	// Group1
	MasterSlaveServerGroupName *string `json:"MasterSlaveServerGroupName,omitempty" xml:"MasterSlaveServerGroupName,omitempty"`
	OwnerAccount               *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Classic Load Balancer (CLB) instance.
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
	Tag []*CreateMasterSlaveServerGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateMasterSlaveServerGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMasterSlaveServerGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateMasterSlaveServerGroupRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateMasterSlaveServerGroupRequest) GetMasterSlaveBackendServers() *string {
	return s.MasterSlaveBackendServers
}

func (s *CreateMasterSlaveServerGroupRequest) GetMasterSlaveServerGroupName() *string {
	return s.MasterSlaveServerGroupName
}

func (s *CreateMasterSlaveServerGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateMasterSlaveServerGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateMasterSlaveServerGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateMasterSlaveServerGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateMasterSlaveServerGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateMasterSlaveServerGroupRequest) GetTag() []*CreateMasterSlaveServerGroupRequestTag {
	return s.Tag
}

func (s *CreateMasterSlaveServerGroupRequest) SetLoadBalancerId(v string) *CreateMasterSlaveServerGroupRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateMasterSlaveServerGroupRequest) SetMasterSlaveBackendServers(v string) *CreateMasterSlaveServerGroupRequest {
	s.MasterSlaveBackendServers = &v
	return s
}

func (s *CreateMasterSlaveServerGroupRequest) SetMasterSlaveServerGroupName(v string) *CreateMasterSlaveServerGroupRequest {
	s.MasterSlaveServerGroupName = &v
	return s
}

func (s *CreateMasterSlaveServerGroupRequest) SetOwnerAccount(v string) *CreateMasterSlaveServerGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateMasterSlaveServerGroupRequest) SetOwnerId(v int64) *CreateMasterSlaveServerGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMasterSlaveServerGroupRequest) SetRegionId(v string) *CreateMasterSlaveServerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMasterSlaveServerGroupRequest) SetResourceOwnerAccount(v string) *CreateMasterSlaveServerGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateMasterSlaveServerGroupRequest) SetResourceOwnerId(v int64) *CreateMasterSlaveServerGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateMasterSlaveServerGroupRequest) SetTag(v []*CreateMasterSlaveServerGroupRequestTag) *CreateMasterSlaveServerGroupRequest {
	s.Tag = v
	return s
}

func (s *CreateMasterSlaveServerGroupRequest) Validate() error {
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

type CreateMasterSlaveServerGroupRequestTag struct {
	// The key of tag N. Valid values of N: **1*	- to **20**. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length, and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N. Valid values of N: **1 to 20**. The tag value can be an empty string. The tag value can be up to 128 characters in length, and cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateMasterSlaveServerGroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateMasterSlaveServerGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateMasterSlaveServerGroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateMasterSlaveServerGroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateMasterSlaveServerGroupRequestTag) SetKey(v string) *CreateMasterSlaveServerGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateMasterSlaveServerGroupRequestTag) SetValue(v string) *CreateMasterSlaveServerGroupRequestTag {
	s.Value = &v
	return s
}

func (s *CreateMasterSlaveServerGroupRequestTag) Validate() error {
	return dara.Validate(s)
}
