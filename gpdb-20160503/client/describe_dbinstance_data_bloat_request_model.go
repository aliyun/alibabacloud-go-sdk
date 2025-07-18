// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceDataBloatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceDataBloatRequest
	GetDBInstanceId() *string
	SetDatabase(v string) *DescribeDBInstanceDataBloatRequest
	GetDatabase() *string
	SetOrderBy(v string) *DescribeDBInstanceDataBloatRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *DescribeDBInstanceDataBloatRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBInstanceDataBloatRequest
	GetPageSize() *int32
}

type DescribeDBInstanceDataBloatRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
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
	// The sorting order.
	//
	// example:
	//
	// {Field: TableName, Type: Desc}
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
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
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDBInstanceDataBloatRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceDataBloatRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataBloatRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceDataBloatRequest) GetDatabase() *string {
	return s.Database
}

func (s *DescribeDBInstanceDataBloatRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeDBInstanceDataBloatRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstanceDataBloatRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBInstanceDataBloatRequest) SetDBInstanceId(v string) *DescribeDBInstanceDataBloatRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceDataBloatRequest) SetDatabase(v string) *DescribeDBInstanceDataBloatRequest {
	s.Database = &v
	return s
}

func (s *DescribeDBInstanceDataBloatRequest) SetOrderBy(v string) *DescribeDBInstanceDataBloatRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeDBInstanceDataBloatRequest) SetPageNumber(v int32) *DescribeDBInstanceDataBloatRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceDataBloatRequest) SetPageSize(v int32) *DescribeDBInstanceDataBloatRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstanceDataBloatRequest) Validate() error {
	return dara.Validate(s)
}
