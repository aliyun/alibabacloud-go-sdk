// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActiveAddressType(v string) *CreateHostRequest
	GetActiveAddressType() *string
	SetComment(v string) *CreateHostRequest
	GetComment() *string
	SetHostName(v string) *CreateHostRequest
	GetHostName() *string
	SetHostPrivateAddress(v string) *CreateHostRequest
	GetHostPrivateAddress() *string
	SetHostPublicAddress(v string) *CreateHostRequest
	GetHostPublicAddress() *string
	SetInstanceId(v string) *CreateHostRequest
	GetInstanceId() *string
	SetInstanceMemberId(v int64) *CreateHostRequest
	GetInstanceMemberId() *int64
	SetInstanceRegionId(v string) *CreateHostRequest
	GetInstanceRegionId() *string
	SetNetworkDomainId(v string) *CreateHostRequest
	GetNetworkDomainId() *string
	SetOSType(v string) *CreateHostRequest
	GetOSType() *string
	SetProjectId(v int64) *CreateHostRequest
	GetProjectId() *int64
	SetRegionId(v string) *CreateHostRequest
	GetRegionId() *string
	SetSource(v string) *CreateHostRequest
	GetSource() *string
	SetSourceInstanceId(v string) *CreateHostRequest
	GetSourceInstanceId() *string
}

type CreateHostRequest struct {
	// The address type of the host to create. Valid values:
	//
	// - **Public**: public address.
	//
	// - **Private**: private network address.
	//
	// This parameter is required.
	//
	// example:
	//
	// Public
	ActiveAddressType *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	// The description of the host. The description can be up to 500 characters in length.
	//
	// example:
	//
	// Local Host
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The name of the host to create. The name can be up to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// host01
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The private network address of the host to create. You can set this parameter to a domain name or an IP address.
	//
	// > This parameter is required if **ActiveAddressType*	- is set to **Private**.
	//
	// example:
	//
	// 192.168.XX.XX
	HostPrivateAddress *string `json:"HostPrivateAddress,omitempty" xml:"HostPrivateAddress,omitempty"`
	// The public address of the host to create. You can set this parameter to a domain name or an IP address.
	//
	// > This parameter is required if **ActiveAddressType*	- is set to **Public**.
	//
	// example:
	//
	// 172.16.XX.XX
	HostPublicAddress *string `json:"HostPublicAddress,omitempty" xml:"HostPublicAddress,omitempty"`
	// The ID of the Bastionhost instance where you want to create the host.
	//
	// > You can call [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the ECS instance belongs.
	//
	// example:
	//
	// 1605494xxxx
	InstanceMemberId *int64 `json:"InstanceMemberId,omitempty" xml:"InstanceMemberId,omitempty"`
	// The region ID of the ECS instance or the host in the dedicated cluster.
	//
	// > This parameter is required if **Source*	- is set to **Ecs*	- or **Rds**.
	//
	// example:
	//
	// cn-hangzhou
	InstanceRegionId *string `json:"InstanceRegionId,omitempty" xml:"InstanceRegionId,omitempty"`
	// The ID of the network domain to which the host to create belongs.
	//
	// > You can call [ListNetworkDomains](https://help.aliyun.com/document_detail/2758827.html) to query the network domain ID.
	//
	// example:
	//
	// 1
	NetworkDomainId *string `json:"NetworkDomainId,omitempty" xml:"NetworkDomainId,omitempty"`
	// The operating system of the host to create. Valid values:
	//
	// - **Linux**
	//
	// - **Windows**
	//
	// This parameter is required.
	//
	// example:
	//
	// Linux
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the Bastionhost instance where you want to create the host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The source of the host to create. Valid values:
	//
	// - **Local**: on-premises host.
	//
	// - **Ecs**: ECS instance.
	//
	// - **Rds**: host in an ApsaraDB RDS dedicated cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// Local
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The instance ID of the ECS instance or the host in the dedicated cluster.
	//
	// > This parameter is required if **Source*	- is set to **Ecs*	- or **Rds**.
	//
	// example:
	//
	// i-dfabfda
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
}

func (s CreateHostRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHostRequest) GoString() string {
	return s.String()
}

func (s *CreateHostRequest) GetActiveAddressType() *string {
	return s.ActiveAddressType
}

func (s *CreateHostRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateHostRequest) GetHostName() *string {
	return s.HostName
}

func (s *CreateHostRequest) GetHostPrivateAddress() *string {
	return s.HostPrivateAddress
}

func (s *CreateHostRequest) GetHostPublicAddress() *string {
	return s.HostPublicAddress
}

func (s *CreateHostRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateHostRequest) GetInstanceMemberId() *int64 {
	return s.InstanceMemberId
}

func (s *CreateHostRequest) GetInstanceRegionId() *string {
	return s.InstanceRegionId
}

func (s *CreateHostRequest) GetNetworkDomainId() *string {
	return s.NetworkDomainId
}

func (s *CreateHostRequest) GetOSType() *string {
	return s.OSType
}

func (s *CreateHostRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateHostRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateHostRequest) GetSource() *string {
	return s.Source
}

func (s *CreateHostRequest) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *CreateHostRequest) SetActiveAddressType(v string) *CreateHostRequest {
	s.ActiveAddressType = &v
	return s
}

func (s *CreateHostRequest) SetComment(v string) *CreateHostRequest {
	s.Comment = &v
	return s
}

func (s *CreateHostRequest) SetHostName(v string) *CreateHostRequest {
	s.HostName = &v
	return s
}

func (s *CreateHostRequest) SetHostPrivateAddress(v string) *CreateHostRequest {
	s.HostPrivateAddress = &v
	return s
}

func (s *CreateHostRequest) SetHostPublicAddress(v string) *CreateHostRequest {
	s.HostPublicAddress = &v
	return s
}

func (s *CreateHostRequest) SetInstanceId(v string) *CreateHostRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateHostRequest) SetInstanceMemberId(v int64) *CreateHostRequest {
	s.InstanceMemberId = &v
	return s
}

func (s *CreateHostRequest) SetInstanceRegionId(v string) *CreateHostRequest {
	s.InstanceRegionId = &v
	return s
}

func (s *CreateHostRequest) SetNetworkDomainId(v string) *CreateHostRequest {
	s.NetworkDomainId = &v
	return s
}

func (s *CreateHostRequest) SetOSType(v string) *CreateHostRequest {
	s.OSType = &v
	return s
}

func (s *CreateHostRequest) SetProjectId(v int64) *CreateHostRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateHostRequest) SetRegionId(v string) *CreateHostRequest {
	s.RegionId = &v
	return s
}

func (s *CreateHostRequest) SetSource(v string) *CreateHostRequest {
	s.Source = &v
	return s
}

func (s *CreateHostRequest) SetSourceInstanceId(v string) *CreateHostRequest {
	s.SourceInstanceId = &v
	return s
}

func (s *CreateHostRequest) Validate() error {
	return dara.Validate(s)
}
