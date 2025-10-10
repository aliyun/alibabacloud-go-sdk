// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServerGroupServersAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateServerGroupServersAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateServerGroupServersAttributeRequest
	GetDryRun() *bool
	SetServerGroupId(v string) *UpdateServerGroupServersAttributeRequest
	GetServerGroupId() *string
	SetServers(v []*UpdateServerGroupServersAttributeRequestServers) *UpdateServerGroupServersAttributeRequest
	GetServers() []*UpdateServerGroupServersAttributeRequestServers
}

type UpdateServerGroupServersAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
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
	// The server groups. You can specify at most 40 server groups in each call.
	//
	// This parameter is required.
	Servers []*UpdateServerGroupServersAttributeRequestServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
}

func (s UpdateServerGroupServersAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServerGroupServersAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupServersAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateServerGroupServersAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateServerGroupServersAttributeRequest) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *UpdateServerGroupServersAttributeRequest) GetServers() []*UpdateServerGroupServersAttributeRequestServers {
	return s.Servers
}

func (s *UpdateServerGroupServersAttributeRequest) SetClientToken(v string) *UpdateServerGroupServersAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequest) SetDryRun(v bool) *UpdateServerGroupServersAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequest) SetServerGroupId(v string) *UpdateServerGroupServersAttributeRequest {
	s.ServerGroupId = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequest) SetServers(v []*UpdateServerGroupServersAttributeRequestServers) *UpdateServerGroupServersAttributeRequest {
	s.Servers = v
	return s
}

func (s *UpdateServerGroupServersAttributeRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateServerGroupServersAttributeRequestServers struct {
	// The description of the backend server. The description must be 2 to 256 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The port that is used by the backend server. Valid values: **1*	- to **65535**.
	//
	// > You do not need to set this parameter if **ServerType*	- is set to **Fc**.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the backend server.
	//
	// 	- Specify the ID of an Elastic Compute Service (ECS) instance, an elastic network interface (ENI), or an elastic container instance if you set **ServerType*	- to **Ecs**, **Eni**, or **Eci**.
	//
	// 	- Specify an IP address if you set **ServerType*	- to **Ip**.
	//
	// 	- Specify the Alibaba Cloud Resource Name (ARN) of a Function Compute function if you set **ServerType*	- to **Fc**.
	//
	// example:
	//
	// i-bp1f9kdprbgy9uiu****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The IP address of the backend server.
	//
	// example:
	//
	// 192.168.1.1
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The type of the backend server. Valid values:
	//
	// 	- **Ecs**: ECS instance
	//
	// 	- **Eni**: ENI
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
	// The weight of the backend server. Valid values: **0*	- to **100**. Default value: **100**. If the value is set to **0**, no requests are forwarded to the server. You can specify up to 40 servers in each call.
	//
	// > You do not need to set this parameter if **ServerType*	- is set to **Fc**.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateServerGroupServersAttributeRequestServers) String() string {
	return dara.Prettify(s)
}

func (s UpdateServerGroupServersAttributeRequestServers) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupServersAttributeRequestServers) GetDescription() *string {
	return s.Description
}

func (s *UpdateServerGroupServersAttributeRequestServers) GetPort() *int32 {
	return s.Port
}

func (s *UpdateServerGroupServersAttributeRequestServers) GetServerId() *string {
	return s.ServerId
}

func (s *UpdateServerGroupServersAttributeRequestServers) GetServerIp() *string {
	return s.ServerIp
}

func (s *UpdateServerGroupServersAttributeRequestServers) GetServerType() *string {
	return s.ServerType
}

func (s *UpdateServerGroupServersAttributeRequestServers) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateServerGroupServersAttributeRequestServers) SetDescription(v string) *UpdateServerGroupServersAttributeRequestServers {
	s.Description = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequestServers) SetPort(v int32) *UpdateServerGroupServersAttributeRequestServers {
	s.Port = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequestServers) SetServerId(v string) *UpdateServerGroupServersAttributeRequestServers {
	s.ServerId = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequestServers) SetServerIp(v string) *UpdateServerGroupServersAttributeRequestServers {
	s.ServerIp = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequestServers) SetServerType(v string) *UpdateServerGroupServersAttributeRequestServers {
	s.ServerType = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequestServers) SetWeight(v int32) *UpdateServerGroupServersAttributeRequestServers {
	s.Weight = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequestServers) Validate() error {
	return dara.Validate(s)
}
