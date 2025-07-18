// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWaitingSQLInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeWaitingSQLInfoRequest
	GetDBInstanceId() *string
	SetDatabase(v string) *DescribeWaitingSQLInfoRequest
	GetDatabase() *string
	SetPID(v string) *DescribeWaitingSQLInfoRequest
	GetPID() *string
}

type DescribeWaitingSQLInfoRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the instance IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The ID of the process that uniquely identifies the query.
	//
	// >  You can call the [DescribeWaitingSQLRecords](https://help.aliyun.com/document_detail/461735.html) operation to obtain the process IDs of lock-waiting queries.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	PID *string `json:"PID,omitempty" xml:"PID,omitempty"`
}

func (s DescribeWaitingSQLInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWaitingSQLInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeWaitingSQLInfoRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeWaitingSQLInfoRequest) GetDatabase() *string {
	return s.Database
}

func (s *DescribeWaitingSQLInfoRequest) GetPID() *string {
	return s.PID
}

func (s *DescribeWaitingSQLInfoRequest) SetDBInstanceId(v string) *DescribeWaitingSQLInfoRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeWaitingSQLInfoRequest) SetDatabase(v string) *DescribeWaitingSQLInfoRequest {
	s.Database = &v
	return s
}

func (s *DescribeWaitingSQLInfoRequest) SetPID(v string) *DescribeWaitingSQLInfoRequest {
	s.PID = &v
	return s
}

func (s *DescribeWaitingSQLInfoRequest) Validate() error {
	return dara.Validate(s)
}
