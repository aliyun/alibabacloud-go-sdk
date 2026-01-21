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
	if s.Servers != nil {
		for _, item := range s.Servers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddServersToServerGroupRequestServers struct {
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 80
	Port            *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	RemoteIpEnabled *bool  `json:"RemoteIpEnabled,omitempty" xml:"RemoteIpEnabled,omitempty"`
	// example:
	//
	// ecs-123
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// example:
	//
	// 192.168.1.1
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ecs
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
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
