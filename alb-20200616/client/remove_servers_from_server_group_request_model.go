// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveServersFromServerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RemoveServersFromServerGroupRequest
	GetClientToken() *string
	SetDryRun(v bool) *RemoveServersFromServerGroupRequest
	GetDryRun() *bool
	SetServerGroupId(v string) *RemoveServersFromServerGroupRequest
	GetServerGroupId() *string
	SetServers(v []*RemoveServersFromServerGroupRequestServers) *RemoveServersFromServerGroupRequest
	GetServers() []*RemoveServersFromServerGroupRequestServers
}

type RemoveServersFromServerGroupRequest struct {
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
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
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
	// The backend servers to be removed. You can specify up to 200 backend servers.
	//
	// This parameter is required.
	Servers []*RemoveServersFromServerGroupRequestServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
}

func (s RemoveServersFromServerGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveServersFromServerGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveServersFromServerGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RemoveServersFromServerGroupRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RemoveServersFromServerGroupRequest) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *RemoveServersFromServerGroupRequest) GetServers() []*RemoveServersFromServerGroupRequestServers {
	return s.Servers
}

func (s *RemoveServersFromServerGroupRequest) SetClientToken(v string) *RemoveServersFromServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveServersFromServerGroupRequest) SetDryRun(v bool) *RemoveServersFromServerGroupRequest {
	s.DryRun = &v
	return s
}

func (s *RemoveServersFromServerGroupRequest) SetServerGroupId(v string) *RemoveServersFromServerGroupRequest {
	s.ServerGroupId = &v
	return s
}

func (s *RemoveServersFromServerGroupRequest) SetServers(v []*RemoveServersFromServerGroupRequestServers) *RemoveServersFromServerGroupRequest {
	s.Servers = v
	return s
}

func (s *RemoveServersFromServerGroupRequest) Validate() error {
	return dara.Validate(s)
}

type RemoveServersFromServerGroupRequestServers struct {
	// The port that is used by the backend server. Valid values: **1*	- to **65535**.
	//
	// >  This parameter is required when you set **ServerType*	- to **Ecs**, **Eni**, **Eci**, or **Ip**.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The backend server ID.
	//
	// 	- If the server group is of the **Instance*	- type, set ServerId to the ID of a resource of the **Ecs**, **Eni**, or **Eci*	- type.
	//
	// 	- If the server group is of the **Ip*	- type, set this parameter to IP addresses.
	//
	// 	- If the server group is of the **Fc*	- type, set ServerId to the Alibaba Cloud Resource Name (ARN) of a function.
	//
	// >  You can call the [ListServerGroups](https://help.aliyun.com/document_detail/2254862.html) operation to query information about the server group type so that you can set ServerId to a proper value.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1f9kdprbgy9uiu****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The IP address of the elastic network interface (ENI) in inclusive mode.
	//
	// example:
	//
	// 192.168.1.1
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The type of the backend server. Valid values:
	//
	// 	- **Ecs**: Elastic Compute Service (ECS) instance
	//
	// 	- **Eni**: ENI
	//
	// 	- **Eci**: elastic container instance
	//
	// 	- **Ip**: IP address
	//
	// 	- **Fc**: Function Compute instance
	//
	// This parameter is required.
	//
	// example:
	//
	// Ecs
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
}

func (s RemoveServersFromServerGroupRequestServers) String() string {
	return dara.Prettify(s)
}

func (s RemoveServersFromServerGroupRequestServers) GoString() string {
	return s.String()
}

func (s *RemoveServersFromServerGroupRequestServers) GetPort() *int32 {
	return s.Port
}

func (s *RemoveServersFromServerGroupRequestServers) GetServerId() *string {
	return s.ServerId
}

func (s *RemoveServersFromServerGroupRequestServers) GetServerIp() *string {
	return s.ServerIp
}

func (s *RemoveServersFromServerGroupRequestServers) GetServerType() *string {
	return s.ServerType
}

func (s *RemoveServersFromServerGroupRequestServers) SetPort(v int32) *RemoveServersFromServerGroupRequestServers {
	s.Port = &v
	return s
}

func (s *RemoveServersFromServerGroupRequestServers) SetServerId(v string) *RemoveServersFromServerGroupRequestServers {
	s.ServerId = &v
	return s
}

func (s *RemoveServersFromServerGroupRequestServers) SetServerIp(v string) *RemoveServersFromServerGroupRequestServers {
	s.ServerIp = &v
	return s
}

func (s *RemoveServersFromServerGroupRequestServers) SetServerType(v string) *RemoveServersFromServerGroupRequestServers {
	s.ServerType = &v
	return s
}

func (s *RemoveServersFromServerGroupRequestServers) Validate() error {
	return dara.Validate(s)
}
