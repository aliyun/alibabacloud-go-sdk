// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartSupabaseInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *RestartSupabaseInstanceRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *RestartSupabaseInstanceRequest
	GetRegionId() *string
}

type RestartSupabaseInstanceRequest struct {
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxsp-*********
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

func (s RestartSupabaseInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartSupabaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartSupabaseInstanceRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *RestartSupabaseInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestartSupabaseInstanceRequest) SetDBInstanceName(v string) *RestartSupabaseInstanceRequest {
	s.DBInstanceName = &v
	return s
}

func (s *RestartSupabaseInstanceRequest) SetRegionId(v string) *RestartSupabaseInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RestartSupabaseInstanceRequest) Validate() error {
	return dara.Validate(s)
}
