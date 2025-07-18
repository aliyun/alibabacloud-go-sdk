// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosisSQLInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDiagnosisSQLInfoRequest
	GetDBInstanceId() *string
	SetDatabase(v string) *DescribeDiagnosisSQLInfoRequest
	GetDatabase() *string
	SetQueryID(v string) *DescribeDiagnosisSQLInfoRequest
	GetQueryID() *string
}

type DescribeDiagnosisSQLInfoRequest struct {
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query details about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
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
	// adbtest
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The query ID. It is a unique identifier of the query.
	//
	// > You can call the [DescribeDiagnosisRecords](https://help.aliyun.com/document_detail/450511.html) operation to obtain query IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 71403480878****
	QueryID *string `json:"QueryID,omitempty" xml:"QueryID,omitempty"`
}

func (s DescribeDiagnosisSQLInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisSQLInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSQLInfoRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDiagnosisSQLInfoRequest) GetDatabase() *string {
	return s.Database
}

func (s *DescribeDiagnosisSQLInfoRequest) GetQueryID() *string {
	return s.QueryID
}

func (s *DescribeDiagnosisSQLInfoRequest) SetDBInstanceId(v string) *DescribeDiagnosisSQLInfoRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoRequest) SetDatabase(v string) *DescribeDiagnosisSQLInfoRequest {
	s.Database = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoRequest) SetQueryID(v string) *DescribeDiagnosisSQLInfoRequest {
	s.QueryID = &v
	return s
}

func (s *DescribeDiagnosisSQLInfoRequest) Validate() error {
	return dara.Validate(s)
}
