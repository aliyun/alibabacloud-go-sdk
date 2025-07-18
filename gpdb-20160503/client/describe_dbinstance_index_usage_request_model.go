// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceIndexUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceIndexUsageRequest
	GetDBInstanceId() *string
	SetDatabase(v string) *DescribeDBInstanceIndexUsageRequest
	GetDatabase() *string
	SetOrderBy(v string) *DescribeDBInstanceIndexUsageRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *DescribeDBInstanceIndexUsageRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBInstanceIndexUsageRequest
	GetPageSize() *int32
}

type DescribeDBInstanceIndexUsageRequest struct {
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// test
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// order by search condition
	//
	// example:
	//
	// {Field: TableName, Type: Desc}
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number. Pages start from page 1. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// Default value: **30**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDBInstanceIndexUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceIndexUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIndexUsageRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceIndexUsageRequest) GetDatabase() *string {
	return s.Database
}

func (s *DescribeDBInstanceIndexUsageRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeDBInstanceIndexUsageRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstanceIndexUsageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBInstanceIndexUsageRequest) SetDBInstanceId(v string) *DescribeDBInstanceIndexUsageRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageRequest) SetDatabase(v string) *DescribeDBInstanceIndexUsageRequest {
	s.Database = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageRequest) SetOrderBy(v string) *DescribeDBInstanceIndexUsageRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageRequest) SetPageNumber(v int32) *DescribeDBInstanceIndexUsageRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageRequest) SetPageSize(v int32) *DescribeDBInstanceIndexUsageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageRequest) Validate() error {
	return dara.Validate(s)
}
