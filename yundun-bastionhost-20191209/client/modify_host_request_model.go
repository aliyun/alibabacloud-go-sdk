// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *ModifyHostRequest
	GetComment() *string
	SetHostId(v string) *ModifyHostRequest
	GetHostId() *string
	SetHostName(v string) *ModifyHostRequest
	GetHostName() *string
	SetHostPrivateAddress(v string) *ModifyHostRequest
	GetHostPrivateAddress() *string
	SetHostPublicAddress(v string) *ModifyHostRequest
	GetHostPublicAddress() *string
	SetInstanceId(v string) *ModifyHostRequest
	GetInstanceId() *string
	SetNetworkDomainId(v string) *ModifyHostRequest
	GetNetworkDomainId() *string
	SetOSType(v string) *ModifyHostRequest
	GetOSType() *string
	SetPrefKex(v string) *ModifyHostRequest
	GetPrefKex() *string
	SetRegionId(v string) *ModifyHostRequest
	GetRegionId() *string
}

type ModifyHostRequest struct {
	// The new description of the host. The description can be up to 500 characters in length.
	//
	// example:
	//
	// Host for test.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the host.
	//
	// > You can call the [ListHosts](https://help.aliyun.com/document_detail/200665.html) operation to query the ID of the host.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The new name of the host. The name can be up to 128 characters.
	//
	// example:
	//
	// TestHost
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The new internal endpoint of the host. You can set this parameter to a domain name or an IP address.
	//
	// example:
	//
	// 193.168.XX.XX
	HostPrivateAddress *string `json:"HostPrivateAddress,omitempty" xml:"HostPrivateAddress,omitempty"`
	// The new public endpoint of the host. You can set this parameter to a domain name or an IP address.
	//
	// example:
	//
	// 200.1.XX.XX
	HostPublicAddress *string `json:"HostPublicAddress,omitempty" xml:"HostPublicAddress,omitempty"`
	// The ID of the bastion host on which you want to modify the information about the host.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the new network domain to which the host belongs.
	//
	// > You can call the [ListNetworkDomains](https://help.aliyun.com/document_detail/2758827.html) operation to query the network domain ID.
	//
	// example:
	//
	// 1
	NetworkDomainId *string `json:"NetworkDomainId,omitempty" xml:"NetworkDomainId,omitempty"`
	// The new operating system of the host. Valid values:
	//
	// 	- **Linux**
	//
	// 	- **Windows**
	//
	// example:
	//
	// Linux
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// The preferred key exchange algorithm of the host. If you set OSType to Linux, you can modify this parameter. Valid values:
	//
	// 	- **default**
	//
	// 	- **diffie-hellman-group1-sha1**
	//
	// 	- **diffie-hellman-group14-sha1**
	//
	// 	- **diffie-hellman-group-exchange-sha1**
	//
	// example:
	//
	// default
	PrefKex *string `json:"PrefKex,omitempty" xml:"PrefKex,omitempty"`
	// The region ID of the bastion host on which you want to modify the information about the host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyHostRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostRequest) GoString() string {
	return s.String()
}

func (s *ModifyHostRequest) GetComment() *string {
	return s.Comment
}

func (s *ModifyHostRequest) GetHostId() *string {
	return s.HostId
}

func (s *ModifyHostRequest) GetHostName() *string {
	return s.HostName
}

func (s *ModifyHostRequest) GetHostPrivateAddress() *string {
	return s.HostPrivateAddress
}

func (s *ModifyHostRequest) GetHostPublicAddress() *string {
	return s.HostPublicAddress
}

func (s *ModifyHostRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyHostRequest) GetNetworkDomainId() *string {
	return s.NetworkDomainId
}

func (s *ModifyHostRequest) GetOSType() *string {
	return s.OSType
}

func (s *ModifyHostRequest) GetPrefKex() *string {
	return s.PrefKex
}

func (s *ModifyHostRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyHostRequest) SetComment(v string) *ModifyHostRequest {
	s.Comment = &v
	return s
}

func (s *ModifyHostRequest) SetHostId(v string) *ModifyHostRequest {
	s.HostId = &v
	return s
}

func (s *ModifyHostRequest) SetHostName(v string) *ModifyHostRequest {
	s.HostName = &v
	return s
}

func (s *ModifyHostRequest) SetHostPrivateAddress(v string) *ModifyHostRequest {
	s.HostPrivateAddress = &v
	return s
}

func (s *ModifyHostRequest) SetHostPublicAddress(v string) *ModifyHostRequest {
	s.HostPublicAddress = &v
	return s
}

func (s *ModifyHostRequest) SetInstanceId(v string) *ModifyHostRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHostRequest) SetNetworkDomainId(v string) *ModifyHostRequest {
	s.NetworkDomainId = &v
	return s
}

func (s *ModifyHostRequest) SetOSType(v string) *ModifyHostRequest {
	s.OSType = &v
	return s
}

func (s *ModifyHostRequest) SetPrefKex(v string) *ModifyHostRequest {
	s.PrefKex = &v
	return s
}

func (s *ModifyHostRequest) SetRegionId(v string) *ModifyHostRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHostRequest) Validate() error {
	return dara.Validate(s)
}
