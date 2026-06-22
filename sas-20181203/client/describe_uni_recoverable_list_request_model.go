// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUniRecoverableListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeUniRecoverableListRequest
	GetCurrentPage() *int32
	SetDatabase(v string) *DescribeUniRecoverableListRequest
	GetDatabase() *string
	SetPageSize(v int32) *DescribeUniRecoverableListRequest
	GetPageSize() *int32
	SetPolicyId(v int64) *DescribeUniRecoverableListRequest
	GetPolicyId() *int64
}

type DescribeUniRecoverableListRequest struct {
	// The page number of the page to return. Default value: **1**, which indicates the first page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The database name.
	//
	// example:
	//
	// msdb
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The maximum number of entries per page when using paging. Default value: 20. If you leave this parameter empty, 20 entries are returned per page by default.
	//
	// > Do not leave PageSize empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the anti-ransomware backup policy for the database.
	//
	// >You can call the [DescribeUniBackupPolicies](~~DescribeUniBackupPolicies~~) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s DescribeUniRecoverableListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUniRecoverableListRequest) GoString() string {
	return s.String()
}

func (s *DescribeUniRecoverableListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeUniRecoverableListRequest) GetDatabase() *string {
	return s.Database
}

func (s *DescribeUniRecoverableListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeUniRecoverableListRequest) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *DescribeUniRecoverableListRequest) SetCurrentPage(v int32) *DescribeUniRecoverableListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeUniRecoverableListRequest) SetDatabase(v string) *DescribeUniRecoverableListRequest {
	s.Database = &v
	return s
}

func (s *DescribeUniRecoverableListRequest) SetPageSize(v int32) *DescribeUniRecoverableListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUniRecoverableListRequest) SetPolicyId(v int64) *DescribeUniRecoverableListRequest {
	s.PolicyId = &v
	return s
}

func (s *DescribeUniRecoverableListRequest) Validate() error {
	return dara.Validate(s)
}
