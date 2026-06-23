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
	SetRegionId(v string) *ListNetworkDomainsRequest
	GetRegionId() *string
}

type ListNetworkDomainsRequest struct {
	// The ID of the Bastionhost instance.
	//
	// > Call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain the instance ID.
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
	// - **Proxy**: proxy connection
	//
	// example:
	//
	// Proxy
	NetworkDomainType *string `json:"NetworkDomainType,omitempty" xml:"NetworkDomainType,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.<br> The maximum value is 100. The default value is 20. If you do not specify this parameter, 20 entries are returned on each page.<br>
	//
	// > Do not leave the PageSize value empty.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For more information about regions and zones, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
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

func (s *ListNetworkDomainsRequest) SetRegionId(v string) *ListNetworkDomainsRequest {
	s.RegionId = &v
	return s
}

func (s *ListNetworkDomainsRequest) Validate() error {
	return dara.Validate(s)
}
