// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteDatabaseRequest
	GetDBInstanceId() *string
	SetDatabaseName(v string) *DeleteDatabaseRequest
	GetDatabaseName() *string
}

type DeleteDatabaseRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the IDs of all AnalyticDB for PostgreSQL instances in the specified region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the queried database.
	//
	// This parameter is required.
	//
	// example:
	//
	// test01
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
}

func (s DeleteDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatabaseRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteDatabaseRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DeleteDatabaseRequest) SetDBInstanceId(v string) *DeleteDatabaseRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDatabaseRequest) SetDatabaseName(v string) *DeleteDatabaseRequest {
	s.DatabaseName = &v
	return s
}

func (s *DeleteDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
