// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationDatabaseAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseAccountName(v string) *ListOperationDatabaseAccountsRequest
	GetDatabaseAccountName() *string
	SetDatabaseId(v string) *ListOperationDatabaseAccountsRequest
	GetDatabaseId() *string
	SetInstanceId(v string) *ListOperationDatabaseAccountsRequest
	GetInstanceId() *string
	SetPageNumber(v string) *ListOperationDatabaseAccountsRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListOperationDatabaseAccountsRequest
	GetPageSize() *string
	SetRegionId(v string) *ListOperationDatabaseAccountsRequest
	GetRegionId() *string
}

type ListOperationDatabaseAccountsRequest struct {
	// The name of the database account. Exact match is supported.
	//
	// example:
	//
	// test
	DatabaseAccountName *string `json:"DatabaseAccountName,omitempty" xml:"DatabaseAccountName,omitempty"`
	// The database ID.
	//
	// >  You can call the [ListOperationDatabases](https://help.aliyun.com/document_detail/2758856.html) operation to query the database ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 56
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The ID of the bastion host.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-2r42t9cvf0i
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.\\
	//
	// Maximum value: 100. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
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
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListOperationDatabaseAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOperationDatabaseAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListOperationDatabaseAccountsRequest) GetDatabaseAccountName() *string {
	return s.DatabaseAccountName
}

func (s *ListOperationDatabaseAccountsRequest) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *ListOperationDatabaseAccountsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOperationDatabaseAccountsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListOperationDatabaseAccountsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListOperationDatabaseAccountsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListOperationDatabaseAccountsRequest) SetDatabaseAccountName(v string) *ListOperationDatabaseAccountsRequest {
	s.DatabaseAccountName = &v
	return s
}

func (s *ListOperationDatabaseAccountsRequest) SetDatabaseId(v string) *ListOperationDatabaseAccountsRequest {
	s.DatabaseId = &v
	return s
}

func (s *ListOperationDatabaseAccountsRequest) SetInstanceId(v string) *ListOperationDatabaseAccountsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOperationDatabaseAccountsRequest) SetPageNumber(v string) *ListOperationDatabaseAccountsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOperationDatabaseAccountsRequest) SetPageSize(v string) *ListOperationDatabaseAccountsRequest {
	s.PageSize = &v
	return s
}

func (s *ListOperationDatabaseAccountsRequest) SetRegionId(v string) *ListOperationDatabaseAccountsRequest {
	s.RegionId = &v
	return s
}

func (s *ListOperationDatabaseAccountsRequest) Validate() error {
	return dara.Validate(s)
}
