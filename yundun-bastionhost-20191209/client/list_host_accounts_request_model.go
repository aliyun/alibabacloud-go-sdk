// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostAccountName(v string) *ListHostAccountsRequest
	GetHostAccountName() *string
	SetHostId(v string) *ListHostAccountsRequest
	GetHostId() *string
	SetHostIds(v string) *ListHostAccountsRequest
	GetHostIds() *string
	SetInstanceId(v string) *ListHostAccountsRequest
	GetInstanceId() *string
	SetPageNumber(v string) *ListHostAccountsRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListHostAccountsRequest
	GetPageSize() *string
	SetProtocolName(v string) *ListHostAccountsRequest
	GetProtocolName() *string
	SetRegionId(v string) *ListHostAccountsRequest
	GetRegionId() *string
}

type ListHostAccountsRequest struct {
	// The name of the host account that you want to query. The name can be up to 128 characters in length. This parameter supports only term queries.
	//
	// example:
	//
	// abc
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// The ID of the host for which you want to query host accounts.
	//
	// > You can call the [ListHosts](https://help.aliyun.com/document_detail/200665.html) operation to obtain the host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The array of host IDs for which you want to query host accounts.
	//
	// > This parameter takes effect only when the value of the HostId parameter is 0. If the HostId parameter is specified with a non-zero value, this parameter is ignored.
	//
	// example:
	//
	// ["2","3"]
	HostIds *string `json:"HostIds,omitempty" xml:"HostIds,omitempty"`
	// The ID of the Bastionhost instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.<br> The maximum value of the PageSize parameter is 100. The default value is 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The protocol of the host account that you want to query.<br> Valid values:
	//
	// - SSH
	//
	// - RDP
	//
	// example:
	//
	// SSH
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListHostAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHostAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListHostAccountsRequest) GetHostAccountName() *string {
	return s.HostAccountName
}

func (s *ListHostAccountsRequest) GetHostId() *string {
	return s.HostId
}

func (s *ListHostAccountsRequest) GetHostIds() *string {
	return s.HostIds
}

func (s *ListHostAccountsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListHostAccountsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListHostAccountsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListHostAccountsRequest) GetProtocolName() *string {
	return s.ProtocolName
}

func (s *ListHostAccountsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListHostAccountsRequest) SetHostAccountName(v string) *ListHostAccountsRequest {
	s.HostAccountName = &v
	return s
}

func (s *ListHostAccountsRequest) SetHostId(v string) *ListHostAccountsRequest {
	s.HostId = &v
	return s
}

func (s *ListHostAccountsRequest) SetHostIds(v string) *ListHostAccountsRequest {
	s.HostIds = &v
	return s
}

func (s *ListHostAccountsRequest) SetInstanceId(v string) *ListHostAccountsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostAccountsRequest) SetPageNumber(v string) *ListHostAccountsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHostAccountsRequest) SetPageSize(v string) *ListHostAccountsRequest {
	s.PageSize = &v
	return s
}

func (s *ListHostAccountsRequest) SetProtocolName(v string) *ListHostAccountsRequest {
	s.ProtocolName = &v
	return s
}

func (s *ListHostAccountsRequest) SetRegionId(v string) *ListHostAccountsRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostAccountsRequest) Validate() error {
	return dara.Validate(s)
}
