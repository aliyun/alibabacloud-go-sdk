// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetIMVMonitorDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ResetIMVMonitorDataRequest
	GetDBInstanceId() *string
	SetDatabase(v string) *ResetIMVMonitorDataRequest
	GetDatabase() *string
}

type ResetIMVMonitorDataRequest struct {
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp152460513z****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// testdb
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
}

func (s ResetIMVMonitorDataRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetIMVMonitorDataRequest) GoString() string {
	return s.String()
}

func (s *ResetIMVMonitorDataRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ResetIMVMonitorDataRequest) GetDatabase() *string {
	return s.Database
}

func (s *ResetIMVMonitorDataRequest) SetDBInstanceId(v string) *ResetIMVMonitorDataRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ResetIMVMonitorDataRequest) SetDatabase(v string) *ResetIMVMonitorDataRequest {
	s.Database = &v
	return s
}

func (s *ResetIMVMonitorDataRequest) Validate() error {
	return dara.Validate(s)
}
