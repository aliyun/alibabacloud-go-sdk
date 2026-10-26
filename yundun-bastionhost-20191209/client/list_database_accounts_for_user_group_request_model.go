// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabaseAccountsForUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseAccountName(v string) *ListDatabaseAccountsForUserGroupRequest
	GetDatabaseAccountName() *string
	SetDatabaseId(v string) *ListDatabaseAccountsForUserGroupRequest
	GetDatabaseId() *string
	SetInstanceId(v string) *ListDatabaseAccountsForUserGroupRequest
	GetInstanceId() *string
	SetPageNumber(v string) *ListDatabaseAccountsForUserGroupRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListDatabaseAccountsForUserGroupRequest
	GetPageSize() *string
	SetProjectId(v int64) *ListDatabaseAccountsForUserGroupRequest
	GetProjectId() *int64
	SetRegionId(v string) *ListDatabaseAccountsForUserGroupRequest
	GetRegionId() *string
	SetUserGroupId(v string) *ListDatabaseAccountsForUserGroupRequest
	GetUserGroupId() *string
}

type ListDatabaseAccountsForUserGroupRequest struct {
	// The name of the database account to query. The name can be up to 128 characters in length. Only exact match is supported.
	//
	// example:
	//
	// test
	DatabaseAccountName *string `json:"DatabaseAccountName,omitempty" xml:"DatabaseAccountName,omitempty"`
	// The instance ID of the database for which you want to query database accounts.
	//
	// > You can invoke the [ListDatabaseAccounts](https://help.aliyun.com/document_detail/2758839.html) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 36
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The ID of the Bastionhost instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-pe334a03o0h
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number of the current page for paging in a paged query. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries to return on each page for paging in a paged query.
	//
	// The maximum value of the PageSize parameter is 100. The default number of entries returned on each page is 20. If the PageSize parameter is left empty, 20 entries are returned by default.
	//
	// > Do not leave the PageSize parameter empty.
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
	// The ID of the user group for which you want to query the authorized database accounts.
	//
	// > You can call the [ListUserGroups](https://help.aliyun.com/document_detail/204509.html) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListDatabaseAccountsForUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatabaseAccountsForUserGroupRequest) GoString() string {
	return s.String()
}

func (s *ListDatabaseAccountsForUserGroupRequest) GetDatabaseAccountName() *string {
	return s.DatabaseAccountName
}

func (s *ListDatabaseAccountsForUserGroupRequest) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *ListDatabaseAccountsForUserGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDatabaseAccountsForUserGroupRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListDatabaseAccountsForUserGroupRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDatabaseAccountsForUserGroupRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDatabaseAccountsForUserGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDatabaseAccountsForUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListDatabaseAccountsForUserGroupRequest) SetDatabaseAccountName(v string) *ListDatabaseAccountsForUserGroupRequest {
	s.DatabaseAccountName = &v
	return s
}

func (s *ListDatabaseAccountsForUserGroupRequest) SetDatabaseId(v string) *ListDatabaseAccountsForUserGroupRequest {
	s.DatabaseId = &v
	return s
}

func (s *ListDatabaseAccountsForUserGroupRequest) SetInstanceId(v string) *ListDatabaseAccountsForUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDatabaseAccountsForUserGroupRequest) SetPageNumber(v string) *ListDatabaseAccountsForUserGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatabaseAccountsForUserGroupRequest) SetPageSize(v string) *ListDatabaseAccountsForUserGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatabaseAccountsForUserGroupRequest) SetProjectId(v int64) *ListDatabaseAccountsForUserGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDatabaseAccountsForUserGroupRequest) SetRegionId(v string) *ListDatabaseAccountsForUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ListDatabaseAccountsForUserGroupRequest) SetUserGroupId(v string) *ListDatabaseAccountsForUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *ListDatabaseAccountsForUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
