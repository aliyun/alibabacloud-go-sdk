// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateNetworkDomainRequest
	GetComment() *string
	SetInstanceId(v string) *CreateNetworkDomainRequest
	GetInstanceId() *string
	SetNetworkDomainName(v string) *CreateNetworkDomainRequest
	GetNetworkDomainName() *string
	SetNetworkDomainType(v string) *CreateNetworkDomainRequest
	GetNetworkDomainType() *string
	SetProjectId(v int64) *CreateNetworkDomainRequest
	GetProjectId() *int64
	SetProxies(v []*CreateNetworkDomainRequestProxies) *CreateNetworkDomainRequest
	GetProxies() []*CreateNetworkDomainRequestProxies
	SetRegionId(v string) *CreateNetworkDomainRequest
	GetRegionId() *string
}

type CreateNetworkDomainRequest struct {
	// The description of the network domain. The description can be up to 500 characters in length.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the Bastionhost instance for which you want to create a network domain.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-lbj3bw4ma02
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the network domain to create. The name can be up to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// SSH Proxy
	NetworkDomainName *string `json:"NetworkDomainName,omitempty" xml:"NetworkDomainName,omitempty"`
	// The type of the network domain to create. Valid values:
	//
	// - Direct: The bastion host directly connects to the asset network without going through an intermediate proxy server.
	//
	// - Proxy: When the asset network cannot communicate with the bastion host network, a proxy server forwards network requests to enable O&M on assets in different network environments.
	//
	// This parameter is required.
	//
	// example:
	//
	// Proxy
	NetworkDomainType *string `json:"NetworkDomainType,omitempty" xml:"NetworkDomainType,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The proxy server information.
	Proxies []*CreateNetworkDomainRequestProxies `json:"Proxies,omitempty" xml:"Proxies,omitempty" type:"Repeated"`
	// The region ID of the Bastionhost instance for which you want to create a network domain.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateNetworkDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkDomainRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateNetworkDomainRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateNetworkDomainRequest) GetNetworkDomainName() *string {
	return s.NetworkDomainName
}

func (s *CreateNetworkDomainRequest) GetNetworkDomainType() *string {
	return s.NetworkDomainType
}

func (s *CreateNetworkDomainRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateNetworkDomainRequest) GetProxies() []*CreateNetworkDomainRequestProxies {
	return s.Proxies
}

func (s *CreateNetworkDomainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNetworkDomainRequest) SetComment(v string) *CreateNetworkDomainRequest {
	s.Comment = &v
	return s
}

func (s *CreateNetworkDomainRequest) SetInstanceId(v string) *CreateNetworkDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateNetworkDomainRequest) SetNetworkDomainName(v string) *CreateNetworkDomainRequest {
	s.NetworkDomainName = &v
	return s
}

func (s *CreateNetworkDomainRequest) SetNetworkDomainType(v string) *CreateNetworkDomainRequest {
	s.NetworkDomainType = &v
	return s
}

func (s *CreateNetworkDomainRequest) SetProjectId(v int64) *CreateNetworkDomainRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateNetworkDomainRequest) SetProxies(v []*CreateNetworkDomainRequestProxies) *CreateNetworkDomainRequest {
	s.Proxies = v
	return s
}

func (s *CreateNetworkDomainRequest) SetRegionId(v string) *CreateNetworkDomainRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNetworkDomainRequest) Validate() error {
	if s.Proxies != nil {
		for _, item := range s.Proxies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateNetworkDomainRequestProxies struct {
	// The address of the proxy server.
	//
	// example:
	//
	// ``47.104.**.**``
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The node type of the proxy server. Valid values:
	//
	// - Master: primary proxy server.
	//
	// - Slave: secondary proxy server.
	//
	// example:
	//
	// Master
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The Base64-encoded password of the proxy server account.
	//
	// example:
	//
	// UWdi******Ng==
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The Server Port of the proxy server.
	//
	// example:
	//
	// 22
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The proxy type. Valid values:
	//
	// - SSHProxy: SSH proxy.
	//
	// - HTTPProxy: HTTP proxy.
	//
	// - Socks5Proxy: SOCKS proxy.
	//
	// example:
	//
	// SSHProxy
	ProxyType *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	// The account of the proxy server.
	//
	// example:
	//
	// root
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s CreateNetworkDomainRequestProxies) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkDomainRequestProxies) GoString() string {
	return s.String()
}

func (s *CreateNetworkDomainRequestProxies) GetAddress() *string {
	return s.Address
}

func (s *CreateNetworkDomainRequestProxies) GetNodeType() *string {
	return s.NodeType
}

func (s *CreateNetworkDomainRequestProxies) GetPassword() *string {
	return s.Password
}

func (s *CreateNetworkDomainRequestProxies) GetPort() *int32 {
	return s.Port
}

func (s *CreateNetworkDomainRequestProxies) GetProxyType() *string {
	return s.ProxyType
}

func (s *CreateNetworkDomainRequestProxies) GetUser() *string {
	return s.User
}

func (s *CreateNetworkDomainRequestProxies) SetAddress(v string) *CreateNetworkDomainRequestProxies {
	s.Address = &v
	return s
}

func (s *CreateNetworkDomainRequestProxies) SetNodeType(v string) *CreateNetworkDomainRequestProxies {
	s.NodeType = &v
	return s
}

func (s *CreateNetworkDomainRequestProxies) SetPassword(v string) *CreateNetworkDomainRequestProxies {
	s.Password = &v
	return s
}

func (s *CreateNetworkDomainRequestProxies) SetPort(v int32) *CreateNetworkDomainRequestProxies {
	s.Port = &v
	return s
}

func (s *CreateNetworkDomainRequestProxies) SetProxyType(v string) *CreateNetworkDomainRequestProxies {
	s.ProxyType = &v
	return s
}

func (s *CreateNetworkDomainRequestProxies) SetUser(v string) *CreateNetworkDomainRequestProxies {
	s.User = &v
	return s
}

func (s *CreateNetworkDomainRequestProxies) Validate() error {
	return dara.Validate(s)
}
