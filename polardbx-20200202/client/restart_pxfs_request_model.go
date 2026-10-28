// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartPxfsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *RestartPxfsRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *RestartPxfsRequest
	GetRegionId() *string
}

type RestartPxfsRequest struct {
	// The PolarDB-X instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-xxxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RestartPxfsRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartPxfsRequest) GoString() string {
	return s.String()
}

func (s *RestartPxfsRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *RestartPxfsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestartPxfsRequest) SetDBInstanceName(v string) *RestartPxfsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *RestartPxfsRequest) SetRegionId(v string) *RestartPxfsRequest {
	s.RegionId = &v
	return s
}

func (s *RestartPxfsRequest) Validate() error {
	return dara.Validate(s)
}
