// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveBackendServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServers(v string) *RemoveBackendServersRequest
	GetBackendServers() *string
	SetLoadBalancerId(v string) *RemoveBackendServersRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *RemoveBackendServersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RemoveBackendServersRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RemoveBackendServersRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RemoveBackendServersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RemoveBackendServersRequest
	GetResourceOwnerId() *int64
}

type RemoveBackendServersRequest struct {
	// The backend servers that you want to remove.
	//
	// 	- **ServerId**: The IDs of the backend servers. Set the value to a string. This parameter is required.
	//
	// 	- **Type**: the type of the backend server. Valid values:
	//
	//     	- **ecs*	- (default): Elastic Compute Service (ECS) instance
	//
	//     	- **eni**: elastic network interface (ENI)
	//
	//     	- **eci**: elastic container instance
	//
	// 	- **Weight**: the weight of the backend server. Valid values: **0*	- to **100**. Set the value to an integer.
	//
	// You can specify at most 20 backend servers in each call. Examples:
	//
	// 	- Remove ECS instances:
	//
	// `[{"ServerId":"i-bp1fq61enf4loa5i****", "Type": "ecs","Weight":"100"}]`
	//
	// 	- Remove ENIs:
	//
	// `[{"ServerId":"eni-2ze1sdp5****","Type": "eni","Weight":"100"}]`
	//
	// 	- Remove elastic container instances:
	//
	// `[{"ServerId":"eci-2ze1sdp5****","Type": "eci","Weight":"100"}]`
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"ServerId":"i-bp1fq61enf4loa5i****", "Type": "ecs","Weight":"100"}]
	BackendServers *string `json:"BackendServers,omitempty" xml:"BackendServers,omitempty"`
	// The ID of the CLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp15lbk8uja8rvm4a****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the CLB instance is deployed.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RemoveBackendServersRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveBackendServersRequest) GoString() string {
	return s.String()
}

func (s *RemoveBackendServersRequest) GetBackendServers() *string {
	return s.BackendServers
}

func (s *RemoveBackendServersRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *RemoveBackendServersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RemoveBackendServersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RemoveBackendServersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveBackendServersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RemoveBackendServersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RemoveBackendServersRequest) SetBackendServers(v string) *RemoveBackendServersRequest {
	s.BackendServers = &v
	return s
}

func (s *RemoveBackendServersRequest) SetLoadBalancerId(v string) *RemoveBackendServersRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *RemoveBackendServersRequest) SetOwnerAccount(v string) *RemoveBackendServersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveBackendServersRequest) SetOwnerId(v int64) *RemoveBackendServersRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveBackendServersRequest) SetRegionId(v string) *RemoveBackendServersRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveBackendServersRequest) SetResourceOwnerAccount(v string) *RemoveBackendServersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveBackendServersRequest) SetResourceOwnerId(v int64) *RemoveBackendServersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveBackendServersRequest) Validate() error {
	return dara.Validate(s)
}
