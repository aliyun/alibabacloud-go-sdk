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
	SetProjectId(v int64) *ListHostAccountsForUserRequest
	GetProjectId() *int64
	SetRegionId(v string) *ListHostAccountsForUserRequest
	GetRegionId() *string
	SetUserId(v string) *ListHostAccountsForUserRequest
	GetUserId() *string
}

type ListHostAccountsForUserRequest struct {
	// The name of the host account. Exact match is supported.
	//
	// example:
	//
	// root
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// The ID of the host for which you want to query the host accounts.
	//
	// > You can call the [ListHosts](https://help.aliyun.com/document_detail/200665.html) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The instance ID of the Bastionhost instance to which the user whose host accounts with authorization you want to query belongs.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number for paging in a paged query. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries to return on each page for paging in a paged query.
	//
	// The maximum value of the PageSize parameter is 100. Default value: 20. If you leave this parameter empty, 20 entries are returned.
	//
	// > Specify a value for the PageSize parameter.
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
	// The region ID of the Bastionhost instance where the user whose authorized host accounts you want to query is located.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user whose authorized and unauthorized host accounts you want to query.
	//
	// > You can call the [ListUsers](https://help.aliyun.com/document_detail/204522.html) operation to obtain the value of this parameter.
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

func (s *ListHostAccountsForUserRequest) GetProjectId() *int64 {
	return s.ProjectId
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

func (s *ListHostAccountsForUserRequest) SetProjectId(v int64) *ListHostAccountsForUserRequest {
	s.ProjectId = &v
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
