// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJDBCDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeJDBCDataSourceRequest
	GetDBInstanceId() *string
	SetDataSourceId(v string) *DescribeJDBCDataSourceRequest
	GetDataSourceId() *string
}

type DescribeJDBCDataSourceRequest struct {
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
}

func (s DescribeJDBCDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeJDBCDataSourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeJDBCDataSourceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeJDBCDataSourceRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *DescribeJDBCDataSourceRequest) SetDBInstanceId(v string) *DescribeJDBCDataSourceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeJDBCDataSourceRequest) SetDataSourceId(v string) *DescribeJDBCDataSourceRequest {
	s.DataSourceId = &v
	return s
}

func (s *DescribeJDBCDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
