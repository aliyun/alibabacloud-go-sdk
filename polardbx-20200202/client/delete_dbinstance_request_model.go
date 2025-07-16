// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DeleteDBInstanceRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DeleteDBInstanceRequest
	GetRegionId() *string
}

type DeleteDBInstanceRequest struct {
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
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DeleteDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDBInstanceRequest) SetDBInstanceName(v string) *DeleteDBInstanceRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetRegionId(v string) *DeleteDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
