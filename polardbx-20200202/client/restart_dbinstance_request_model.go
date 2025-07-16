// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *RestartDBInstanceRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *RestartDBInstanceRequest
	GetRegionId() *string
}

type RestartDBInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
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

func (s *RestartDBInstanceRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *RestartDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestartDBInstanceRequest) SetDBInstanceName(v string) *RestartDBInstanceRequest {
	s.DBInstanceName = &v
	return s
}

func (s *RestartDBInstanceRequest) SetRegionId(v string) *RestartDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RestartDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
