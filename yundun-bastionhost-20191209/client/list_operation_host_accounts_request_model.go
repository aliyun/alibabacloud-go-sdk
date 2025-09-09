// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationHostAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostAccountName(v string) *ListOperationHostAccountsRequest
	GetHostAccountName() *string
	SetHostId(v string) *ListOperationHostAccountsRequest
	GetHostId() *string
	SetInstanceId(v string) *ListOperationHostAccountsRequest
	GetInstanceId() *string
	SetPageNumber(v string) *ListOperationHostAccountsRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListOperationHostAccountsRequest
	GetPageSize() *string
	SetRegionId(v string) *ListOperationHostAccountsRequest
	GetRegionId() *string
}

type ListOperationHostAccountsRequest struct {
	// The name of the host account to query. Only exact match is supported.
	//
	// example:
	//
	// root
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// The ID of the host whose accounts you want to query.
	//
	// >  You can call the [ListOperationHosts](https://help.aliyun.com/document_detail/2758857.html) operation to query the host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The bastion host ID.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.\\
	//
	// Valid values: 1 to 100. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
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
}

func (s ListOperationHostAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOperationHostAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListOperationHostAccountsRequest) GetHostAccountName() *string {
	return s.HostAccountName
}

func (s *ListOperationHostAccountsRequest) GetHostId() *string {
	return s.HostId
}

func (s *ListOperationHostAccountsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOperationHostAccountsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListOperationHostAccountsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListOperationHostAccountsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListOperationHostAccountsRequest) SetHostAccountName(v string) *ListOperationHostAccountsRequest {
	s.HostAccountName = &v
	return s
}

func (s *ListOperationHostAccountsRequest) SetHostId(v string) *ListOperationHostAccountsRequest {
	s.HostId = &v
	return s
}

func (s *ListOperationHostAccountsRequest) SetInstanceId(v string) *ListOperationHostAccountsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOperationHostAccountsRequest) SetPageNumber(v string) *ListOperationHostAccountsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOperationHostAccountsRequest) SetPageSize(v string) *ListOperationHostAccountsRequest {
	s.PageSize = &v
	return s
}

func (s *ListOperationHostAccountsRequest) SetRegionId(v string) *ListOperationHostAccountsRequest {
	s.RegionId = &v
	return s
}

func (s *ListOperationHostAccountsRequest) Validate() error {
	return dara.Validate(s)
}
