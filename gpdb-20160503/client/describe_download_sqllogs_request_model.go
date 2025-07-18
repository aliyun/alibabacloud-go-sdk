// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDownloadSQLLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDownloadSQLLogsRequest
	GetDBInstanceId() *string
}

type DescribeDownloadSQLLogsRequest struct {
	// Instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) API to view details of all AnalyticDB for PostgreSQL instances in the target region, including the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeDownloadSQLLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadSQLLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDownloadSQLLogsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDownloadSQLLogsRequest) SetDBInstanceId(v string) *DescribeDownloadSQLLogsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDownloadSQLLogsRequest) Validate() error {
	return dara.Validate(s)
}
