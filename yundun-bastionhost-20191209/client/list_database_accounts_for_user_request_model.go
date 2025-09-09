// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabaseAccountsForUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseAccountName(v string) *ListDatabaseAccountsForUserRequest
	GetDatabaseAccountName() *string
	SetDatabaseId(v string) *ListDatabaseAccountsForUserRequest
	GetDatabaseId() *string
	SetInstanceId(v string) *ListDatabaseAccountsForUserRequest
	GetInstanceId() *string
	SetPageNumber(v string) *ListDatabaseAccountsForUserRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListDatabaseAccountsForUserRequest
	GetPageSize() *string
	SetRegionId(v string) *ListDatabaseAccountsForUserRequest
	GetRegionId() *string
	SetUserId(v string) *ListDatabaseAccountsForUserRequest
	GetUserId() *string
}

type ListDatabaseAccountsForUserRequest struct {
	// The name of the database account to query. The name can be up to 128 characters in length. Only exact match is supported.
	//
	// example:
	//
	// test
	DatabaseAccountName *string `json:"DatabaseAccountName,omitempty" xml:"DatabaseAccountName,omitempty"`
	// The ID of the database whose accounts you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 89
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The bastion host ID.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-zz42zoqql01
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
	// > We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user to query. This operation returns whether the user is authorized to manage each database account.
	//
	// > You can call the ListUsers operation to query the ID of the user.[](~~204522~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListDatabaseAccountsForUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatabaseAccountsForUserRequest) GoString() string {
	return s.String()
}

func (s *ListDatabaseAccountsForUserRequest) GetDatabaseAccountName() *string {
	return s.DatabaseAccountName
}

func (s *ListDatabaseAccountsForUserRequest) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *ListDatabaseAccountsForUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDatabaseAccountsForUserRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListDatabaseAccountsForUserRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDatabaseAccountsForUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDatabaseAccountsForUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListDatabaseAccountsForUserRequest) SetDatabaseAccountName(v string) *ListDatabaseAccountsForUserRequest {
	s.DatabaseAccountName = &v
	return s
}

func (s *ListDatabaseAccountsForUserRequest) SetDatabaseId(v string) *ListDatabaseAccountsForUserRequest {
	s.DatabaseId = &v
	return s
}

func (s *ListDatabaseAccountsForUserRequest) SetInstanceId(v string) *ListDatabaseAccountsForUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDatabaseAccountsForUserRequest) SetPageNumber(v string) *ListDatabaseAccountsForUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatabaseAccountsForUserRequest) SetPageSize(v string) *ListDatabaseAccountsForUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatabaseAccountsForUserRequest) SetRegionId(v string) *ListDatabaseAccountsForUserRequest {
	s.RegionId = &v
	return s
}

func (s *ListDatabaseAccountsForUserRequest) SetUserId(v string) *ListDatabaseAccountsForUserRequest {
	s.UserId = &v
	return s
}

func (s *ListDatabaseAccountsForUserRequest) Validate() error {
	return dara.Validate(s)
}
