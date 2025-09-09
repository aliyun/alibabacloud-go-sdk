// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationHostsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostAddress(v string) *ListOperationHostsRequest
	GetHostAddress() *string
	SetHostName(v string) *ListOperationHostsRequest
	GetHostName() *string
	SetInstanceId(v string) *ListOperationHostsRequest
	GetInstanceId() *string
	SetOSType(v string) *ListOperationHostsRequest
	GetOSType() *string
	SetPageNumber(v string) *ListOperationHostsRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListOperationHostsRequest
	GetPageSize() *string
	SetRegionId(v string) *ListOperationHostsRequest
	GetRegionId() *string
	SetSource(v string) *ListOperationHostsRequest
	GetSource() *string
	SetSourceInstanceId(v string) *ListOperationHostsRequest
	GetSourceInstanceId() *string
	SetSourceInstanceState(v string) *ListOperationHostsRequest
	GetSourceInstanceState() *string
}

type ListOperationHostsRequest struct {
	// The address of the host that you want to query. You can set this parameter to a domain name or an IP address. Only exact match is supported.
	//
	// example:
	//
	// 10.162.172.132
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The name of the host that you want to query. Only exact match is supported.
	//
	// example:
	//
	// abc
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the bastion host.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-09k22avmw0q
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
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.\\
	//
	// Maximum value: 100. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The source of the host that you want to query. Valid values:
	//
	// 	- **Local**
	//
	// 	- **Ecs**
	//
	// example:
	//
	// Local
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ID of the Elastic Compute Service (ECS) instance. Exact match is supported.
	//
	// example:
	//
	// i-bp19ienyt0yax748****
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	// The status of the host that you want to query. Valid values:
	//
	// 	- **Normal**
	//
	// 	- **Release**
	//
	// example:
	//
	// Normal
	SourceInstanceState *string `json:"SourceInstanceState,omitempty" xml:"SourceInstanceState,omitempty"`
}

func (s ListOperationHostsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOperationHostsRequest) GoString() string {
	return s.String()
}

func (s *ListOperationHostsRequest) GetHostAddress() *string {
	return s.HostAddress
}

func (s *ListOperationHostsRequest) GetHostName() *string {
	return s.HostName
}

func (s *ListOperationHostsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOperationHostsRequest) GetOSType() *string {
	return s.OSType
}

func (s *ListOperationHostsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListOperationHostsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListOperationHostsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListOperationHostsRequest) GetSource() *string {
	return s.Source
}

func (s *ListOperationHostsRequest) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *ListOperationHostsRequest) GetSourceInstanceState() *string {
	return s.SourceInstanceState
}

func (s *ListOperationHostsRequest) SetHostAddress(v string) *ListOperationHostsRequest {
	s.HostAddress = &v
	return s
}

func (s *ListOperationHostsRequest) SetHostName(v string) *ListOperationHostsRequest {
	s.HostName = &v
	return s
}

func (s *ListOperationHostsRequest) SetInstanceId(v string) *ListOperationHostsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOperationHostsRequest) SetOSType(v string) *ListOperationHostsRequest {
	s.OSType = &v
	return s
}

func (s *ListOperationHostsRequest) SetPageNumber(v string) *ListOperationHostsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOperationHostsRequest) SetPageSize(v string) *ListOperationHostsRequest {
	s.PageSize = &v
	return s
}

func (s *ListOperationHostsRequest) SetRegionId(v string) *ListOperationHostsRequest {
	s.RegionId = &v
	return s
}

func (s *ListOperationHostsRequest) SetSource(v string) *ListOperationHostsRequest {
	s.Source = &v
	return s
}

func (s *ListOperationHostsRequest) SetSourceInstanceId(v string) *ListOperationHostsRequest {
	s.SourceInstanceId = &v
	return s
}

func (s *ListOperationHostsRequest) SetSourceInstanceState(v string) *ListOperationHostsRequest {
	s.SourceInstanceState = &v
	return s
}

func (s *ListOperationHostsRequest) Validate() error {
	return dara.Validate(s)
}
