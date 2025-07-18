// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceDataSkewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceDataSkewRequest
	GetDBInstanceId() *string
	SetDatabase(v string) *DescribeDBInstanceDataSkewRequest
	GetDatabase() *string
	SetOrderBy(v string) *DescribeDBInstanceDataSkewRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *DescribeDBInstanceDataSkewRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBInstanceDataSkewRequest
	GetPageSize() *int32
}

type DescribeDBInstanceDataSkewRequest struct {
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
	// order by condition
	//
	// example:
	//
	// {Field: TableSkew, Type: Desc}
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: **1**.
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

func (s DescribeDBInstanceDataSkewRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceDataSkewRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataSkewRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceDataSkewRequest) GetDatabase() *string {
	return s.Database
}

func (s *DescribeDBInstanceDataSkewRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeDBInstanceDataSkewRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstanceDataSkewRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBInstanceDataSkewRequest) SetDBInstanceId(v string) *DescribeDBInstanceDataSkewRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceDataSkewRequest) SetDatabase(v string) *DescribeDBInstanceDataSkewRequest {
	s.Database = &v
	return s
}

func (s *DescribeDBInstanceDataSkewRequest) SetOrderBy(v string) *DescribeDBInstanceDataSkewRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeDBInstanceDataSkewRequest) SetPageNumber(v int32) *DescribeDBInstanceDataSkewRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceDataSkewRequest) SetPageSize(v int32) *DescribeDBInstanceDataSkewRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstanceDataSkewRequest) Validate() error {
	return dara.Validate(s)
}
