// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceDatabasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ListInstanceDatabasesRequest
	GetDBInstanceId() *string
	SetPageNumber(v int32) *ListInstanceDatabasesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstanceDatabasesRequest
	GetPageSize() *int32
}

type ListInstanceDatabasesRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **20**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// Default value: **20**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInstanceDatabasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceDatabasesRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceDatabasesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListInstanceDatabasesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstanceDatabasesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstanceDatabasesRequest) SetDBInstanceId(v string) *ListInstanceDatabasesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListInstanceDatabasesRequest) SetPageNumber(v int32) *ListInstanceDatabasesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceDatabasesRequest) SetPageSize(v int32) *ListInstanceDatabasesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceDatabasesRequest) Validate() error {
	return dara.Validate(s)
}
