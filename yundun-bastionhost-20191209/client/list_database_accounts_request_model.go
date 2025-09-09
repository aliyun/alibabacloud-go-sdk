// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabaseAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseAccountName(v string) *ListDatabaseAccountsRequest
	GetDatabaseAccountName() *string
	SetDatabaseId(v string) *ListDatabaseAccountsRequest
	GetDatabaseId() *string
	SetInstanceId(v string) *ListDatabaseAccountsRequest
	GetInstanceId() *string
	SetPageNumber(v string) *ListDatabaseAccountsRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListDatabaseAccountsRequest
	GetPageSize() *string
	SetRegionId(v string) *ListDatabaseAccountsRequest
	GetRegionId() *string
}

type ListDatabaseAccountsRequest struct {
	// The name of the database account to query. The name can be up to 128 characters in length. Only exact match is supported.
	//
	// example:
	//
	// test
	DatabaseAccountName *string `json:"DatabaseAccountName,omitempty" xml:"DatabaseAccountName,omitempty"`
	// The ID of the database whose database accounts you want to query.
	//
	// >  You can call the [ListDatabases](https://help.aliyun.com/document_detail/2758822.html) operation to query the database ID.
	//
	// example:
	//
	// 3
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The bastion host ID.
	//
	// > You can call the DescribeInstances operation to query the bastion host ID.[](~~153281~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-7mz2za0ro06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Default value: 1.
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
}

func (s ListDatabaseAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatabaseAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListDatabaseAccountsRequest) GetDatabaseAccountName() *string {
	return s.DatabaseAccountName
}

func (s *ListDatabaseAccountsRequest) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *ListDatabaseAccountsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDatabaseAccountsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListDatabaseAccountsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDatabaseAccountsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDatabaseAccountsRequest) SetDatabaseAccountName(v string) *ListDatabaseAccountsRequest {
	s.DatabaseAccountName = &v
	return s
}

func (s *ListDatabaseAccountsRequest) SetDatabaseId(v string) *ListDatabaseAccountsRequest {
	s.DatabaseId = &v
	return s
}

func (s *ListDatabaseAccountsRequest) SetInstanceId(v string) *ListDatabaseAccountsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDatabaseAccountsRequest) SetPageNumber(v string) *ListDatabaseAccountsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatabaseAccountsRequest) SetPageSize(v string) *ListDatabaseAccountsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatabaseAccountsRequest) SetRegionId(v string) *ListDatabaseAccountsRequest {
	s.RegionId = &v
	return s
}

func (s *ListDatabaseAccountsRequest) Validate() error {
	return dara.Validate(s)
}
