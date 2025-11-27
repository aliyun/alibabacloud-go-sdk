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
	SetDBName(v string) *DeleteDatabaseRequest
	GetDBName() *string
	SetResourceOwnerId(v int64) *DeleteDatabaseRequest
	GetResourceOwnerId() *int64
}

type DeleteDatabaseRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5*****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// testdb01
	DBName          *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

func (s *DeleteDatabaseRequest) GetDBName() *string {
	return s.DBName
}

func (s *DeleteDatabaseRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDatabaseRequest) SetDBInstanceId(v string) *DeleteDatabaseRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDatabaseRequest) SetDBName(v string) *DeleteDatabaseRequest {
	s.DBName = &v
	return s
}

func (s *DeleteDatabaseRequest) SetResourceOwnerId(v int64) *DeleteDatabaseRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
