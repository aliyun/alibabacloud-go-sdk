// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostAddress(v string) *ListHostsRequest
	GetHostAddress() *string
	SetHostGroupId(v string) *ListHostsRequest
	GetHostGroupId() *string
	SetHostName(v string) *ListHostsRequest
	GetHostName() *string
	SetInstanceId(v string) *ListHostsRequest
	GetInstanceId() *string
	SetOSType(v string) *ListHostsRequest
	GetOSType() *string
	SetPageNumber(v string) *ListHostsRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListHostsRequest
	GetPageSize() *string
	SetRegionId(v string) *ListHostsRequest
	GetRegionId() *string
	SetSource(v string) *ListHostsRequest
	GetSource() *string
	SetSourceInstanceId(v string) *ListHostsRequest
	GetSourceInstanceId() *string
	SetSourceInstanceState(v string) *ListHostsRequest
	GetSourceInstanceState() *string
}

type ListHostsRequest struct {
	// The address of the host that you want to query. You can set this parameter to a domain name or an IP address. Only exact match is supported.
	//
	// example:
	//
	// 1.1.XX.XX
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The ID of the host group to which the host to be queried belongs.
	//
	// > You can call the [ListHostGroups](https://help.aliyun.com/document_detail/201307.html) operation to query the ID of the host group.
	//
	// example:
	//
	// 1
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The name of the host that you want to query. Only exact match is supported.
	//
	// example:
	//
	// host
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the bastion host on which you want to query hosts.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The operating system of the host that you want to query. Valid values:
	//
	// 	- **Linux**
	//
	// 	- **Windows**
	//
	// example:
	//
	// Linux
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// > We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host on which you want to query hosts.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The source of the host that you want to query. Valid values:
	//
	// 	- **Local**: a host in a data center
	//
	// 	- **Ecs**: an Elastic Compute Service (ECS) instance
	//
	// 	- **Rds**: a host in an ApsaraDB MyBase dedicated cluster
	//
	// example:
	//
	// Local
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ID of the ECS instance or the host in an ApsaraDB MyBase dedicated cluster that you want to query. Only exact match is supported.
	//
	// example:
	//
	// i-bp19ienyt0yax748****
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	// The status of the host that you want to query. Valid values:
	//
	// 	- **Normal**: normal
	//
	// 	- **Release**: released
	//
	// example:
	//
	// Normal
	SourceInstanceState *string `json:"SourceInstanceState,omitempty" xml:"SourceInstanceState,omitempty"`
}

func (s ListHostsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHostsRequest) GoString() string {
	return s.String()
}

func (s *ListHostsRequest) GetHostAddress() *string {
	return s.HostAddress
}

func (s *ListHostsRequest) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *ListHostsRequest) GetHostName() *string {
	return s.HostName
}

func (s *ListHostsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListHostsRequest) GetOSType() *string {
	return s.OSType
}

func (s *ListHostsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListHostsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListHostsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListHostsRequest) GetSource() *string {
	return s.Source
}

func (s *ListHostsRequest) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *ListHostsRequest) GetSourceInstanceState() *string {
	return s.SourceInstanceState
}

func (s *ListHostsRequest) SetHostAddress(v string) *ListHostsRequest {
	s.HostAddress = &v
	return s
}

func (s *ListHostsRequest) SetHostGroupId(v string) *ListHostsRequest {
	s.HostGroupId = &v
	return s
}

func (s *ListHostsRequest) SetHostName(v string) *ListHostsRequest {
	s.HostName = &v
	return s
}

func (s *ListHostsRequest) SetInstanceId(v string) *ListHostsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostsRequest) SetOSType(v string) *ListHostsRequest {
	s.OSType = &v
	return s
}

func (s *ListHostsRequest) SetPageNumber(v string) *ListHostsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHostsRequest) SetPageSize(v string) *ListHostsRequest {
	s.PageSize = &v
	return s
}

func (s *ListHostsRequest) SetRegionId(v string) *ListHostsRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostsRequest) SetSource(v string) *ListHostsRequest {
	s.Source = &v
	return s
}

func (s *ListHostsRequest) SetSourceInstanceId(v string) *ListHostsRequest {
	s.SourceInstanceId = &v
	return s
}

func (s *ListHostsRequest) SetSourceInstanceState(v string) *ListHostsRequest {
	s.SourceInstanceState = &v
	return s
}

func (s *ListHostsRequest) Validate() error {
	return dara.Validate(s)
}
