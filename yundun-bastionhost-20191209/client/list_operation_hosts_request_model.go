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
	SetOperationProjectId(v string) *ListOperationHostsRequest
	GetOperationProjectId() *string
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
	// The address of the host to query. You can use a domain name or an IP address. Fuzzy match is not supported. Only exact match is supported.
	//
	// example:
	//
	// 10.162.XX.XX
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The name of the host to query. Fuzzy match is not supported. Only exact match is supported.
	//
	// example:
	//
	// abc
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The instance ID of the Bastionhost instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-09k22avmw0q
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The operating system of the host to query. Valid values:
	//
	// - **Linux**
	//
	// - **Windows**
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
	OperationProjectId *string `json:"OperationProjectId,omitempty" xml:"OperationProjectId,omitempty"`
	// The page number of the current page for a paged query. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries to return on each page for a paged query.
	//
	// Maximum value of PageSize: 100. Default value: 20. If you leave PageSize empty, 20 entries are returned by default.
	//
	// > Do not leave PageSize empty.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For the mappings between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The source of the host to query. Valid values:
	//
	// - **Local**: local host
	//
	// - **Ecs**: ECS instance
	//
	// example:
	//
	// Local
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ECS instance ID. Exact search.
	//
	// example:
	//
	// i-bp19ienyt0yax748****
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	// The status of the host to query. Valid values:
	//
	// - **Normal**: Normal
	//
	// - **Release**: Released
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

func (s *ListOperationHostsRequest) GetOperationProjectId() *string {
	return s.OperationProjectId
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

func (s *ListOperationHostsRequest) SetOperationProjectId(v string) *ListOperationHostsRequest {
	s.OperationProjectId = &v
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
