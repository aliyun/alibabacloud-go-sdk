// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteDBInstanceRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DeleteDBInstanceRequest
	GetRegionId() *string
}

type DeleteDBInstanceRequest struct {
	// The ID of the destination cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID.
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

func (s *DeleteDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDBInstanceRequest) SetDBInstanceId(v string) *DeleteDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetRegionId(v string) *DeleteDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
