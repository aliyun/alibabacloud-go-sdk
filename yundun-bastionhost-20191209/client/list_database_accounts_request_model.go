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
	SetProjectId(v int64) *ListDatabaseAccountsRequest
	GetProjectId() *int64
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
	// The ID of the database instance for which you want to query database accounts.
	//
	// > You can call the [ListDatabases](https://help.aliyun.com/document_detail/2758822.html) operation to query the database instance ID.
	//
	// example:
	//
	// 3
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The ID of the Bastionhost instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the Bastionhost instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-7mz2za0ro06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number of the current page in a paging query. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries to return on each page in a paging query.
	//
	// The maximum value of the PageSize parameter is 100. Default value: 20. If this parameter is left empty, 20 entries are returned by default.
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
	// The region ID of the Bastionhost instance.
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

func (s *ListDatabaseAccountsRequest) GetProjectId() *int64 {
	return s.ProjectId
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

func (s *ListDatabaseAccountsRequest) SetProjectId(v int64) *ListDatabaseAccountsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDatabaseAccountsRequest) SetRegionId(v string) *ListDatabaseAccountsRequest {
	s.RegionId = &v
	return s
}

func (s *ListDatabaseAccountsRequest) Validate() error {
	return dara.Validate(s)
}
