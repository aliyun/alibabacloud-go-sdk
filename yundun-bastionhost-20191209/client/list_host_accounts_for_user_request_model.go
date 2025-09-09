// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostAccountsForUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostAccountName(v string) *ListHostAccountsForUserRequest
	GetHostAccountName() *string
	SetHostId(v string) *ListHostAccountsForUserRequest
	GetHostId() *string
	SetInstanceId(v string) *ListHostAccountsForUserRequest
	GetInstanceId() *string
	SetPageNumber(v string) *ListHostAccountsForUserRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListHostAccountsForUserRequest
	GetPageSize() *string
	SetRegionId(v string) *ListHostAccountsForUserRequest
	GetRegionId() *string
	SetUserId(v string) *ListHostAccountsForUserRequest
	GetUserId() *string
}

type ListHostAccountsForUserRequest struct {
	// The name of the host account that you want to query. Exact match is supported.
	//
	// example:
	//
	// root
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// The ID of the host to query.
	//
	// > You can call the [ListHosts](https://help.aliyun.com/document_detail/200665.html) operation to query the ID of the host.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The ID of the bastion host on which you want to perform the query. The host accounts that the specified user is authorized to manage on the specified host are queried.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
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
	// The number of entries to return on each page.\\
	//
	// Maximum value: 100. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host on which you want to perform the query. The host accounts that the specified user is authorized to manage on the specified host are queried.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user for which you want to query authorized host accounts.
	//
	// > You can call the [ListUsers](https://help.aliyun.com/document_detail/204522.html) operation to query the ID of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListHostAccountsForUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHostAccountsForUserRequest) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForUserRequest) GetHostAccountName() *string {
	return s.HostAccountName
}

func (s *ListHostAccountsForUserRequest) GetHostId() *string {
	return s.HostId
}

func (s *ListHostAccountsForUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListHostAccountsForUserRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListHostAccountsForUserRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListHostAccountsForUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListHostAccountsForUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListHostAccountsForUserRequest) SetHostAccountName(v string) *ListHostAccountsForUserRequest {
	s.HostAccountName = &v
	return s
}

func (s *ListHostAccountsForUserRequest) SetHostId(v string) *ListHostAccountsForUserRequest {
	s.HostId = &v
	return s
}

func (s *ListHostAccountsForUserRequest) SetInstanceId(v string) *ListHostAccountsForUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostAccountsForUserRequest) SetPageNumber(v string) *ListHostAccountsForUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHostAccountsForUserRequest) SetPageSize(v string) *ListHostAccountsForUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListHostAccountsForUserRequest) SetRegionId(v string) *ListHostAccountsForUserRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostAccountsForUserRequest) SetUserId(v string) *ListHostAccountsForUserRequest {
	s.UserId = &v
	return s
}

func (s *ListHostAccountsForUserRequest) Validate() error {
	return dara.Validate(s)
}
