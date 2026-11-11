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
	SetRegionId(v string) *AddServersToServerGroupRequest
	GetRegionId() *string
	SetServerGroupId(v string) *AddServersToServerGroupRequest
	GetServerGroupId() *string
	SetServers(v []*AddServersToServerGroupRequestServers) *AddServersToServerGroupRequest
	GetServers() []*AddServersToServerGroupRequestServers
}

type AddServersToServerGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// - **true**: prechecks the request without adding the servers. The system checks whether required parameters are specified, the request format is valid, and service limits are met. If the request fails the precheck, an error is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends the request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the region where the Network Load Balancer (NLB) instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/443657.html) operation to query region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the server group.
	//
	// This parameter is required.
	//
	// example:
	//
	// sgp-atstuj3rtoptyui****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The list of backend servers.
	//
	// > You can specify up to 200 servers in a single call.
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

func (s *AddServersToServerGroupRequest) GetRegionId() *string {
	return s.RegionId
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

func (s *AddServersToServerGroupRequest) SetRegionId(v string) *AddServersToServerGroupRequest {
	s.RegionId = &v
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
	// The description of the server.
	//
	// The description must be 2 to 256 characters in length, and can contain letters, digits, Chinese characters, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
	//
	// example:
	//
	// ECS
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The port used by the backend server. Valid values: **0*	- to **65535**. If you do not specify this parameter, the default value **0*	- is used.
	//
	// If the all-port forwarding feature is enabled for the server group, you do not need to specify a port when you add a backend server. In this case, the default value 0 is used, and NLB forwards traffic to the backend server based on the frontend request port. You can check whether the all-port forwarding feature is enabled by querying the **AnyPortEnabled*	- field returned by the [ListServerGroups](https://help.aliyun.com/document_detail/445895.html) operation.
	//
	// example:
	//
	// 443
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the backend server.
	//
	// - If the server group is of the **Instance*	- type, this parameter specifies the resource ID of an **Ecs**, **Eni**, or **Eci*	- resource.
	//
	// - If the server group is of the **Ip*	- type, this parameter specifies an IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The IP address of the server. This parameter is required when the server group is of the **Ip*	- type.
	//
	// example:
	//
	// 192.168.6.6
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The backend server type. Valid values:
	//
	// - **Ecs**: an ECS instance.
	//
	// - **Eni**: an elastic network interface (ENI).
	//
	// - **Eci**: an Elastic Container Instance (ECI).
	//
	// - **Ip**: an IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// Ecs
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// The weight of the backend server. Valid values: **0*	- to **100**. Default value: **100**. If the weight is set to **0**, no requests are forwarded to the backend server.
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
