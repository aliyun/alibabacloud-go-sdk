// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNetworkDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *ModifyNetworkDomainRequest
	GetComment() *string
	SetInstanceId(v string) *ModifyNetworkDomainRequest
	GetInstanceId() *string
	SetNetworkDomainId(v string) *ModifyNetworkDomainRequest
	GetNetworkDomainId() *string
	SetNetworkDomainName(v string) *ModifyNetworkDomainRequest
	GetNetworkDomainName() *string
	SetNetworkDomainType(v string) *ModifyNetworkDomainRequest
	GetNetworkDomainType() *string
	SetProxies(v []*ModifyNetworkDomainRequestProxies) *ModifyNetworkDomainRequest
	GetProxies() []*ModifyNetworkDomainRequestProxies
	SetRegionId(v string) *ModifyNetworkDomainRequest
	GetRegionId() *string
}

type ModifyNetworkDomainRequest struct {
	// The new remarks of the network domain.
	//
	// example:
	//
	// xxx
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the bastion host to which the network domain to modify belongs.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-x0r3hyr3f09
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the network domain to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	NetworkDomainId *string `json:"NetworkDomainId,omitempty" xml:"NetworkDomainId,omitempty"`
	// The new name of the network domain.
	//
	// example:
	//
	// test
	NetworkDomainName *string `json:"NetworkDomainName,omitempty" xml:"NetworkDomainName,omitempty"`
	// The new connection mode of the network domain. Valid values:
	//
	// 	- **Direct**
	//
	// 	- **Proxy**
	//
	// example:
	//
	// Proxy
	NetworkDomainType *string `json:"NetworkDomainType,omitempty" xml:"NetworkDomainType,omitempty"`
	// The information about the proxy servers in the network domain.
	Proxies []*ModifyNetworkDomainRequestProxies `json:"Proxies,omitempty" xml:"Proxies,omitempty" type:"Repeated"`
	// The region ID of the bastion host to which the network domain to modify belongs.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyNetworkDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkDomainRequest) GoString() string {
	return s.String()
}

func (s *ModifyNetworkDomainRequest) GetComment() *string {
	return s.Comment
}

func (s *ModifyNetworkDomainRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyNetworkDomainRequest) GetNetworkDomainId() *string {
	return s.NetworkDomainId
}

func (s *ModifyNetworkDomainRequest) GetNetworkDomainName() *string {
	return s.NetworkDomainName
}

func (s *ModifyNetworkDomainRequest) GetNetworkDomainType() *string {
	return s.NetworkDomainType
}

func (s *ModifyNetworkDomainRequest) GetProxies() []*ModifyNetworkDomainRequestProxies {
	return s.Proxies
}

func (s *ModifyNetworkDomainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyNetworkDomainRequest) SetComment(v string) *ModifyNetworkDomainRequest {
	s.Comment = &v
	return s
}

func (s *ModifyNetworkDomainRequest) SetInstanceId(v string) *ModifyNetworkDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyNetworkDomainRequest) SetNetworkDomainId(v string) *ModifyNetworkDomainRequest {
	s.NetworkDomainId = &v
	return s
}

func (s *ModifyNetworkDomainRequest) SetNetworkDomainName(v string) *ModifyNetworkDomainRequest {
	s.NetworkDomainName = &v
	return s
}

func (s *ModifyNetworkDomainRequest) SetNetworkDomainType(v string) *ModifyNetworkDomainRequest {
	s.NetworkDomainType = &v
	return s
}

func (s *ModifyNetworkDomainRequest) SetProxies(v []*ModifyNetworkDomainRequestProxies) *ModifyNetworkDomainRequest {
	s.Proxies = v
	return s
}

func (s *ModifyNetworkDomainRequest) SetRegionId(v string) *ModifyNetworkDomainRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyNetworkDomainRequest) Validate() error {
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

type ModifyNetworkDomainRequestProxies struct {
	// The new IP address of the proxy server.
	//
	// example:
	//
	// 114.21**.**
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The node type of the proxy server to modify. Valid values:
	//
	// 	- **Master**: primary proxy server.
	//
	// 	- **Slave**: secondary proxy server.
	//
	// example:
	//
	// Slave
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The new password of the proxy server account.
	//
	// example:
	//
	// ***
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The new port of the proxy server.
	//
	// example:
	//
	// 22
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The new proxy mode. Valid values:
	//
	// 	- **SSHProxy**
	//
	// 	- **HTTPProxy**
	//
	// 	- **Socks5Proxy**
	//
	// example:
	//
	// HTTPProxy
	ProxyType *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	// The new username of the proxy server account.
	//
	// example:
	//
	// test
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s ModifyNetworkDomainRequestProxies) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkDomainRequestProxies) GoString() string {
	return s.String()
}

func (s *ModifyNetworkDomainRequestProxies) GetAddress() *string {
	return s.Address
}

func (s *ModifyNetworkDomainRequestProxies) GetNodeType() *string {
	return s.NodeType
}

func (s *ModifyNetworkDomainRequestProxies) GetPassword() *string {
	return s.Password
}

func (s *ModifyNetworkDomainRequestProxies) GetPort() *int32 {
	return s.Port
}

func (s *ModifyNetworkDomainRequestProxies) GetProxyType() *string {
	return s.ProxyType
}

func (s *ModifyNetworkDomainRequestProxies) GetUser() *string {
	return s.User
}

func (s *ModifyNetworkDomainRequestProxies) SetAddress(v string) *ModifyNetworkDomainRequestProxies {
	s.Address = &v
	return s
}

func (s *ModifyNetworkDomainRequestProxies) SetNodeType(v string) *ModifyNetworkDomainRequestProxies {
	s.NodeType = &v
	return s
}

func (s *ModifyNetworkDomainRequestProxies) SetPassword(v string) *ModifyNetworkDomainRequestProxies {
	s.Password = &v
	return s
}

func (s *ModifyNetworkDomainRequestProxies) SetPort(v int32) *ModifyNetworkDomainRequestProxies {
	s.Port = &v
	return s
}

func (s *ModifyNetworkDomainRequestProxies) SetProxyType(v string) *ModifyNetworkDomainRequestProxies {
	s.ProxyType = &v
	return s
}

func (s *ModifyNetworkDomainRequestProxies) SetUser(v string) *ModifyNetworkDomainRequestProxies {
	s.User = &v
	return s
}

func (s *ModifyNetworkDomainRequestProxies) Validate() error {
	return dara.Validate(s)
}
