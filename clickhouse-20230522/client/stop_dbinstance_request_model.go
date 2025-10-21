// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *StopDBInstanceRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *StopDBInstanceRequest
	GetRegionId() *string
}

type StopDBInstanceRequest struct {
	// The cluster ID.
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

func (s StopDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StopDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *StopDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopDBInstanceRequest) SetDBInstanceId(v string) *StopDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *StopDBInstanceRequest) SetRegionId(v string) *StopDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StopDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
