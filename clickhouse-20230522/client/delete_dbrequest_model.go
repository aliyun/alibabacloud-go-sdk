// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteDBRequest
	GetDBInstanceId() *string
	SetDBName(v string) *DeleteDBRequest
	GetDBName() *string
	SetRegionId(v string) *DeleteDBRequest
	GetRegionId() *string
}

type DeleteDBRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the destination database.
	//
	// This parameter is required.
	//
	// example:
	//
	// testdb001
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDBRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteDBRequest) GetDBName() *string {
	return s.DBName
}

func (s *DeleteDBRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDBRequest) SetDBInstanceId(v string) *DeleteDBRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBRequest) SetDBName(v string) *DeleteDBRequest {
	s.DBName = &v
	return s
}

func (s *DeleteDBRequest) SetRegionId(v string) *DeleteDBRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDBRequest) Validate() error {
	return dara.Validate(s)
}
