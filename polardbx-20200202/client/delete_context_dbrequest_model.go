// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContextDBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DeleteContextDBRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DeleteContextDBRequest
	GetRegionId() *string
}

type DeleteContextDBRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region in which the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteContextDBRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextDBRequest) GoString() string {
	return s.String()
}

func (s *DeleteContextDBRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DeleteContextDBRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteContextDBRequest) SetDBInstanceName(v string) *DeleteContextDBRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DeleteContextDBRequest) SetRegionId(v string) *DeleteContextDBRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteContextDBRequest) Validate() error {
	return dara.Validate(s)
}
