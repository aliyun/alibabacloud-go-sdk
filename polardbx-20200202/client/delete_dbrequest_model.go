// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DeleteDBRequest
	GetDBInstanceName() *string
	SetDbName(v string) *DeleteDBRequest
	GetDbName() *string
	SetRegionId(v string) *DeleteDBRequest
	GetRegionId() *string
}

type DeleteDBRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testdb
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
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

func (s *DeleteDBRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DeleteDBRequest) GetDbName() *string {
	return s.DbName
}

func (s *DeleteDBRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDBRequest) SetDBInstanceName(v string) *DeleteDBRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DeleteDBRequest) SetDbName(v string) *DeleteDBRequest {
	s.DbName = &v
	return s
}

func (s *DeleteDBRequest) SetRegionId(v string) *DeleteDBRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDBRequest) Validate() error {
	return dara.Validate(s)
}
