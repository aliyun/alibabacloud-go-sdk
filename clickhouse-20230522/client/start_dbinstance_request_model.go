// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *StartDBInstanceRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *StartDBInstanceRequest
	GetRegionId() *string
}

type StartDBInstanceRequest struct {
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

func (s StartDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *StartDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartDBInstanceRequest) SetDBInstanceId(v string) *StartDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *StartDBInstanceRequest) SetRegionId(v string) *StartDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StartDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
