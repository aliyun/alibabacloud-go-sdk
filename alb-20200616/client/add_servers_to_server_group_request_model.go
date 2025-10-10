// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddServersToServerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AddServersToServerGroupRequest
	GetClientToken() *string
	SetDryRun(v bool) *AddServersToServerGroupRequest
	GetDryRun() *bool
	SetServerGroupId(v string) *AddServersToServerGroupRequest
	GetServerGroupId() *string
	SetServers(v []*AddServersToServerGroupRequestServers) *AddServersToServerGroupRequest
	GetServers() []*AddServersToServerGroupRequestServers
}

type AddServersToServerGroupRequest struct {
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
	// 	- **true**: prechecks the request, but does not add a backend server to a server group. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The server group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sgp-atstuj3rtop****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The backend servers. You can specify at most 200 servers in each call.
	//
	// This parameter is required.
	Servers []*AddServersToServerGroupRequestServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
}

func (s AddServersToServerGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddServersToServerGroupRequest) GoString() string {
	return s.String()
}

func (s *AddServersToServerGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddServersToServerGroupRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AddServersToServerGroupRequest) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *AddServersToServerGroupRequest) GetServers() []*AddServersToServerGroupRequestServers {
	return s.Servers
}

func (s *AddServersToServerGroupRequest) SetClientToken(v string) *AddServersToServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *AddServersToServerGroupRequest) SetDryRun(v bool) *AddServersToServerGroupRequest {
	s.DryRun = &v
	return s
}

func (s *AddServersToServerGroupRequest) SetServerGroupId(v string) *AddServersToServerGroupRequest {
	s.ServerGroupId = &v
	return s
}

func (s *AddServersToServerGroupRequest) SetServers(v []*AddServersToServerGroupRequestServers) *AddServersToServerGroupRequest {
	s.Servers = v
	return s
}

func (s *AddServersToServerGroupRequest) Validate() error {
	return dara.Validate(s)
}

type AddServersToServerGroupRequestServers struct {
	// The description of the backend server. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The port that is used by the backend server. Valid values: **1*	- to **65535**. You can specify at most 200 servers in each call.
	//
	// >  This parameter is required if you set **ServerType*	- to **Ecs**, **Eni**, **Eci**, or **Ip**. You do not need to set this parameter if **ServerType*	- is set to **Fc**.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// Specifies whether to enable the remote IP feature. You can specify at most 200 servers in each call. Default values:
	//
	// 	- **true**: enables the feature.
	//
	// 	- **false**: disables the feature.
	//
	// >  This parameter takes effect only when **ServerType*	- is set to **Ip**.
	//
	// example:
	//
	// false
	RemoteIpEnabled *bool `json:"RemoteIpEnabled,omitempty" xml:"RemoteIpEnabled,omitempty"`
	// The ID of the server group. You can specify at most 200 servers in each call.
	//
	// 	- If the server group is of the **Instance*	- type, set ServerId to the ID of a resource of the **Ecs**, **Eni**, or **Eci*	- type.
	//
	// 	- If the server group is of the **Ip*	- type, set this parameter to IP addresses.
	//
	// 	- If the server group is of the **Fc*	- type, set ServerId to an Alibaba Cloud Resource Name (ARN).
	//
	// example:
	//
	// ecs-bp67acfmxazb4p****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The IP address of the backend server. You can specify at most 200 servers in each call.
	//
	// >  You do not need to set this parameter if you set **ServerType*	- to **Fc**.
	//
	// example:
	//
	// 192.168.1.1
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The type of the backend server. You can specify at most 200 servers in each call. Default values:
	//
	// 	- **Ecs**: Elastic Compute Service (ECS) instance
	//
	// 	- **Eni**: elastic network interface (ENI)
	//
	// 	- **Eci**: elastic container instance
	//
	// 	- **Ip**: IP address
	//
	// 	- **Fc**: Function Compute
	//
	// This parameter is required.
	//
	// example:
	//
	// Ecs
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// The weight of the backend server. Valid values: **0*	- to **100**. Default value: **100**. If the value is set to **0**, no requests are forwarded to the server. You can specify at most 200 servers in each call.
	//
	// >  You do not need to set this parameter if you set **ServerType*	- to **Fc**.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s AddServersToServerGroupRequestServers) String() string {
	return dara.Prettify(s)
}

func (s AddServersToServerGroupRequestServers) GoString() string {
	return s.String()
}

func (s *AddServersToServerGroupRequestServers) GetDescription() *string {
	return s.Description
}

func (s *AddServersToServerGroupRequestServers) GetPort() *int32 {
	return s.Port
}

func (s *AddServersToServerGroupRequestServers) GetRemoteIpEnabled() *bool {
	return s.RemoteIpEnabled
}

func (s *AddServersToServerGroupRequestServers) GetServerId() *string {
	return s.ServerId
}

func (s *AddServersToServerGroupRequestServers) GetServerIp() *string {
	return s.ServerIp
}

func (s *AddServersToServerGroupRequestServers) GetServerType() *string {
	return s.ServerType
}

func (s *AddServersToServerGroupRequestServers) GetWeight() *int32 {
	return s.Weight
}

func (s *AddServersToServerGroupRequestServers) SetDescription(v string) *AddServersToServerGroupRequestServers {
	s.Description = &v
	return s
}

func (s *AddServersToServerGroupRequestServers) SetPort(v int32) *AddServersToServerGroupRequestServers {
	s.Port = &v
	return s
}

func (s *AddServersToServerGroupRequestServers) SetRemoteIpEnabled(v bool) *AddServersToServerGroupRequestServers {
	s.RemoteIpEnabled = &v
	return s
}

func (s *AddServersToServerGroupRequestServers) SetServerId(v string) *AddServersToServerGroupRequestServers {
	s.ServerId = &v
	return s
}

func (s *AddServersToServerGroupRequestServers) SetServerIp(v string) *AddServersToServerGroupRequestServers {
	s.ServerIp = &v
	return s
}

func (s *AddServersToServerGroupRequestServers) SetServerType(v string) *AddServersToServerGroupRequestServers {
	s.ServerType = &v
	return s
}

func (s *AddServersToServerGroupRequestServers) SetWeight(v int32) *AddServersToServerGroupRequestServers {
	s.Weight = &v
	return s
}

func (s *AddServersToServerGroupRequestServers) Validate() error {
	return dara.Validate(s)
}
