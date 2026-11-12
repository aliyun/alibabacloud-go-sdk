// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceNetInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionString(v string) *DescribeDBInstanceNetInfoRequest
	GetConnectionString() *string
	SetDBInstanceId(v string) *DescribeDBInstanceNetInfoRequest
	GetDBInstanceId() *string
}

type DescribeDBInstanceNetInfoRequest struct {
	// The database endpoint.
	//
	// > If you do not specify this parameter, the information about all database endpoints is returned.
	//
	// example:
	//
	// gp-0xin9q82c33xc****-master.gpdb.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-0xin9q82c33xc****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeDBInstanceNetInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBInstanceNetInfoRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceNetInfoRequest) SetConnectionString(v string) *DescribeDBInstanceNetInfoRequest {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) SetDBInstanceId(v string) *DescribeDBInstanceNetInfoRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) Validate() error {
	return dara.Validate(s)
}
