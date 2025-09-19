// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUniBackupDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeUniBackupDatabaseRequest
	GetCurrentPage() *int32
	SetDatabaseType(v string) *DescribeUniBackupDatabaseRequest
	GetDatabaseType() *string
	SetInstanceName(v string) *DescribeUniBackupDatabaseRequest
	GetInstanceName() *string
	SetPageSize(v int32) *DescribeUniBackupDatabaseRequest
	GetPageSize() *int32
	SetQueryType(v string) *DescribeUniBackupDatabaseRequest
	GetQueryType() *string
	SetUniRegionId(v string) *DescribeUniBackupDatabaseRequest
	GetUniRegionId() *string
}

type DescribeUniBackupDatabaseRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The type of the database. Valid values:
	//
	// 	- **MYSQL**
	//
	// 	- **MSSQL**
	//
	// 	- **Oracle**
	//
	// example:
	//
	// MYSQL
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// The name of the Elastic Compute Service (ECS) instance.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The condition that is used to query the database. Valid values:
	//
	// 	- **create**: newly created
	//
	// 	- **restore**: restored
	//
	// example:
	//
	// create
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The region ID of the server that hosts the database.
	//
	// example:
	//
	// cn-hongkong
	UniRegionId *string `json:"UniRegionId,omitempty" xml:"UniRegionId,omitempty"`
}

func (s DescribeUniBackupDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUniBackupDatabaseRequest) GoString() string {
	return s.String()
}

func (s *DescribeUniBackupDatabaseRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeUniBackupDatabaseRequest) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *DescribeUniBackupDatabaseRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeUniBackupDatabaseRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeUniBackupDatabaseRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *DescribeUniBackupDatabaseRequest) GetUniRegionId() *string {
	return s.UniRegionId
}

func (s *DescribeUniBackupDatabaseRequest) SetCurrentPage(v int32) *DescribeUniBackupDatabaseRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeUniBackupDatabaseRequest) SetDatabaseType(v string) *DescribeUniBackupDatabaseRequest {
	s.DatabaseType = &v
	return s
}

func (s *DescribeUniBackupDatabaseRequest) SetInstanceName(v string) *DescribeUniBackupDatabaseRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeUniBackupDatabaseRequest) SetPageSize(v int32) *DescribeUniBackupDatabaseRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUniBackupDatabaseRequest) SetQueryType(v string) *DescribeUniBackupDatabaseRequest {
	s.QueryType = &v
	return s
}

func (s *DescribeUniBackupDatabaseRequest) SetUniRegionId(v string) *DescribeUniBackupDatabaseRequest {
	s.UniRegionId = &v
	return s
}

func (s *DescribeUniBackupDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
