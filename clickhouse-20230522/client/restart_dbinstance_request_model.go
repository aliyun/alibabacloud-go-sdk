// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *RestartDBInstanceRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *RestartDBInstanceRequest
	GetRegionId() *string
}

type RestartDBInstanceRequest struct {
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

func (s RestartDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RestartDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestartDBInstanceRequest) SetDBInstanceId(v string) *RestartDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RestartDBInstanceRequest) SetRegionId(v string) *RestartDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RestartDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
