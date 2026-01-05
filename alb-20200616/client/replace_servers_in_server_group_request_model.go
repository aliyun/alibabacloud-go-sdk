// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceServersInServerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddedServers(v []*ReplaceServersInServerGroupRequestAddedServers) *ReplaceServersInServerGroupRequest
	GetAddedServers() []*ReplaceServersInServerGroupRequestAddedServers
	SetClientToken(v string) *ReplaceServersInServerGroupRequest
	GetClientToken() *string
	SetDryRun(v bool) *ReplaceServersInServerGroupRequest
	GetDryRun() *bool
	SetRemovedServers(v []*ReplaceServersInServerGroupRequestRemovedServers) *ReplaceServersInServerGroupRequest
	GetRemovedServers() []*ReplaceServersInServerGroupRequestRemovedServers
	SetServerGroupId(v string) *ReplaceServersInServerGroupRequest
	GetServerGroupId() *string
}

type ReplaceServersInServerGroupRequest struct {
	// The backend servers. You can specify at most 200 servers in each call.
	//
	// This parameter is required.
	AddedServers []*ReplaceServersInServerGroupRequestAddedServers `json:"AddedServers,omitempty" xml:"AddedServers,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 593B0448-D13E-4C56-AC0D-FDF0FDE0E9A3
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a `2xx` HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The backend servers that you want to remove.
	//
	// This parameter is required.
	RemovedServers []*ReplaceServersInServerGroupRequestRemovedServers `json:"RemovedServers,omitempty" xml:"RemovedServers,omitempty" type:"Repeated"`
	// The ID of the server group.
	//
	// > You cannot perform this operation on a server group of the Function type.
	//
	// This parameter is required.
	//
	// example:
	//
	// sgp-5114d593o96qxy****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s ReplaceServersInServerGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceServersInServerGroupRequest) GoString() string {
	return s.String()
}

func (s *ReplaceServersInServerGroupRequest) GetAddedServers() []*ReplaceServersInServerGroupRequestAddedServers {
	return s.AddedServers
}

func (s *ReplaceServersInServerGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ReplaceServersInServerGroupRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ReplaceServersInServerGroupRequest) GetRemovedServers() []*ReplaceServersInServerGroupRequestRemovedServers {
	return s.RemovedServers
}

func (s *ReplaceServersInServerGroupRequest) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *ReplaceServersInServerGroupRequest) SetAddedServers(v []*ReplaceServersInServerGroupRequestAddedServers) *ReplaceServersInServerGroupRequest {
	s.AddedServers = v
	return s
}

func (s *ReplaceServersInServerGroupRequest) SetClientToken(v string) *ReplaceServersInServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *ReplaceServersInServerGroupRequest) SetDryRun(v bool) *ReplaceServersInServerGroupRequest {
	s.DryRun = &v
	return s
}

func (s *ReplaceServersInServerGroupRequest) SetRemovedServers(v []*ReplaceServersInServerGroupRequestRemovedServers) *ReplaceServersInServerGroupRequest {
	s.RemovedServers = v
	return s
}

func (s *ReplaceServersInServerGroupRequest) SetServerGroupId(v string) *ReplaceServersInServerGroupRequest {
	s.ServerGroupId = &v
	return s
}

func (s *ReplaceServersInServerGroupRequest) Validate() error {
	if s.AddedServers != nil {
		for _, item := range s.AddedServers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RemovedServers != nil {
		for _, item := range s.RemovedServers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReplaceServersInServerGroupRequestAddedServers struct {
	// The description of the backend server. The description must be 2 to 256 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The port used by the backend server in the server group. Valid values: **1*	- to **65535**. You can specify at most 200 servers in each call.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the backend server. You can specify at most 200 servers in each call.
	//
	// 	- If the server group is of the **Instance*	- type, set ServerId to the ID of a resource of the **Ecs**, **Eni**, or **Eci*	- type.
	//
	// 	- If the server group is of the **Ip*	- type, set ServerId to IP addresses.
	//
	// >  You cannot perform this operation on a server group of the Function Compute type. You can call the [ListServerGroups](https://help.aliyun.com/document_detail/213627.html) operation to query the type of server groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1f9kdprbgy9uiu****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The IP address of the elastic network interface (ENI) in exclusive mode.
	//
	// example:
	//
	// 192.168.1.1
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The type of backend server. You can specify at most 200 servers in each call. Valid values:
	//
	// 	- **Ecs**: Elastic Compute Service (ECS) instance
	//
	// 	- **Eni**: ENI
	//
	// 	- **Eci**: elastic container instance
	//
	// example:
	//
	// Ecs
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// The weight of the backend server. You can specify at most 200 servers in each call.
	//
	// Valid values: **0*	- to **100**. Default value: **100**. If the value is set to **0**, no requests are forwarded to the server.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ReplaceServersInServerGroupRequestAddedServers) String() string {
	return dara.Prettify(s)
}

func (s ReplaceServersInServerGroupRequestAddedServers) GoString() string {
	return s.String()
}

func (s *ReplaceServersInServerGroupRequestAddedServers) GetDescription() *string {
	return s.Description
}

func (s *ReplaceServersInServerGroupRequestAddedServers) GetPort() *int32 {
	return s.Port
}

func (s *ReplaceServersInServerGroupRequestAddedServers) GetServerId() *string {
	return s.ServerId
}

func (s *ReplaceServersInServerGroupRequestAddedServers) GetServerIp() *string {
	return s.ServerIp
}

func (s *ReplaceServersInServerGroupRequestAddedServers) GetServerType() *string {
	return s.ServerType
}

func (s *ReplaceServersInServerGroupRequestAddedServers) GetWeight() *int32 {
	return s.Weight
}

func (s *ReplaceServersInServerGroupRequestAddedServers) SetDescription(v string) *ReplaceServersInServerGroupRequestAddedServers {
	s.Description = &v
	return s
}

func (s *ReplaceServersInServerGroupRequestAddedServers) SetPort(v int32) *ReplaceServersInServerGroupRequestAddedServers {
	s.Port = &v
	return s
}

func (s *ReplaceServersInServerGroupRequestAddedServers) SetServerId(v string) *ReplaceServersInServerGroupRequestAddedServers {
	s.ServerId = &v
	return s
}

func (s *ReplaceServersInServerGroupRequestAddedServers) SetServerIp(v string) *ReplaceServersInServerGroupRequestAddedServers {
	s.ServerIp = &v
	return s
}

func (s *ReplaceServersInServerGroupRequestAddedServers) SetServerType(v string) *ReplaceServersInServerGroupRequestAddedServers {
	s.ServerType = &v
	return s
}

func (s *ReplaceServersInServerGroupRequestAddedServers) SetWeight(v int32) *ReplaceServersInServerGroupRequestAddedServers {
	s.Weight = &v
	return s
}

func (s *ReplaceServersInServerGroupRequestAddedServers) Validate() error {
	return dara.Validate(s)
}

type ReplaceServersInServerGroupRequestRemovedServers struct {
	// The port that is used by the backend server. Valid values: **1*	- to **65535**. You can specify at most 200 servers in each call.
	//
	// example:
	//
	// 81
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the backend server. You can specify at most 200 servers in each call.
	//
	// 	- If the server group is of the **Instance*	- type, set ServerId to the ID of a resource of the **Ecs**, **Eni**, or **Eci*	- type.
	//
	// 	- If the server group is of the **Ip*	- type, set ServerId to IP addresses.
	//
	// >  You cannot perform this operation on a server group of the Function Compute type. You can call the [ListServerGroups](https://help.aliyun.com/document_detail/213627.html) operation to query the type of server groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs-bp1ac9uozods2uc****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The IP address of the ENI in exclusive mode.
	//
	// example:
	//
	// 192.168.1.12
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The type of backend server. You can specify at most 200 servers in each call. Valid values:
	//
	// 	- **Ecs**: ECS instance
	//
	// 	- **Eni**: ENI
	//
	// 	- **Eci**: elastic container instance
	//
	// example:
	//
	// ecs
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
}

func (s ReplaceServersInServerGroupRequestRemovedServers) String() string {
	return dara.Prettify(s)
}

func (s ReplaceServersInServerGroupRequestRemovedServers) GoString() string {
	return s.String()
}

func (s *ReplaceServersInServerGroupRequestRemovedServers) GetPort() *int32 {
	return s.Port
}

func (s *ReplaceServersInServerGroupRequestRemovedServers) GetServerId() *string {
	return s.ServerId
}

func (s *ReplaceServersInServerGroupRequestRemovedServers) GetServerIp() *string {
	return s.ServerIp
}

func (s *ReplaceServersInServerGroupRequestRemovedServers) GetServerType() *string {
	return s.ServerType
}

func (s *ReplaceServersInServerGroupRequestRemovedServers) SetPort(v int32) *ReplaceServersInServerGroupRequestRemovedServers {
	s.Port = &v
	return s
}

func (s *ReplaceServersInServerGroupRequestRemovedServers) SetServerId(v string) *ReplaceServersInServerGroupRequestRemovedServers {
	s.ServerId = &v
	return s
}

func (s *ReplaceServersInServerGroupRequestRemovedServers) SetServerIp(v string) *ReplaceServersInServerGroupRequestRemovedServers {
	s.ServerIp = &v
	return s
}

func (s *ReplaceServersInServerGroupRequestRemovedServers) SetServerType(v string) *ReplaceServersInServerGroupRequestRemovedServers {
	s.ServerType = &v
	return s
}

func (s *ReplaceServersInServerGroupRequestRemovedServers) Validate() error {
	return dara.Validate(s)
}
