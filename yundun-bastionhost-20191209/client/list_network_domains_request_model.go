// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListNetworkDomainsRequest
	GetInstanceId() *string
	SetNetworkDomainName(v string) *ListNetworkDomainsRequest
	GetNetworkDomainName() *string
	SetNetworkDomainType(v string) *ListNetworkDomainsRequest
	GetNetworkDomainType() *string
	SetPageNumber(v string) *ListNetworkDomainsRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListNetworkDomainsRequest
	GetPageSize() *string
	SetProjectId(v int64) *ListNetworkDomainsRequest
	GetProjectId() *int64
	SetRegionId(v string) *ListNetworkDomainsRequest
	GetRegionId() *string
}

type ListNetworkDomainsRequest struct {
	// The instance ID of the Bastionhost instance.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-tl329pvu70x
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the network domain.
	//
	// example:
	//
	// SSH Proxy
	NetworkDomainName *string `json:"NetworkDomainName,omitempty" xml:"NetworkDomainName,omitempty"`
	// The type of the network domain. Valid values:
	//
	// - **Direct**: direct connection
	//
	// - **Proxy**: proxy
	//
	// example:
	//
	// Proxy
	NetworkDomainType *string `json:"NetworkDomainType,omitempty" xml:"NetworkDomainType,omitempty"`
	// The page number of the current page for paging in a paged query. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries to return on each page for paging in a paged query.
	//
	// Maximum value: 100. Default value: 20. If this parameter is left empty, 20 entries are returned by default.
	//
	// > Do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListNetworkDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListNetworkDomainsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListNetworkDomainsRequest) GetNetworkDomainName() *string {
	return s.NetworkDomainName
}

func (s *ListNetworkDomainsRequest) GetNetworkDomainType() *string {
	return s.NetworkDomainType
}

func (s *ListNetworkDomainsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListNetworkDomainsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListNetworkDomainsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListNetworkDomainsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListNetworkDomainsRequest) SetInstanceId(v string) *ListNetworkDomainsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListNetworkDomainsRequest) SetNetworkDomainName(v string) *ListNetworkDomainsRequest {
	s.NetworkDomainName = &v
	return s
}

func (s *ListNetworkDomainsRequest) SetNetworkDomainType(v string) *ListNetworkDomainsRequest {
	s.NetworkDomainType = &v
	return s
}

func (s *ListNetworkDomainsRequest) SetPageNumber(v string) *ListNetworkDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNetworkDomainsRequest) SetPageSize(v string) *ListNetworkDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *ListNetworkDomainsRequest) SetProjectId(v int64) *ListNetworkDomainsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListNetworkDomainsRequest) SetRegionId(v string) *ListNetworkDomainsRequest {
	s.RegionId = &v
	return s
}

func (s *ListNetworkDomainsRequest) Validate() error {
	return dara.Validate(s)
}
