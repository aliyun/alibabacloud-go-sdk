// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolardbxSupabaseInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DeletePolardbxSupabaseInstanceRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DeletePolardbxSupabaseInstanceRequest
	GetRegionId() *string
}

type DeletePolardbxSupabaseInstanceRequest struct {
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

func (s DeletePolardbxSupabaseInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolardbxSupabaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeletePolardbxSupabaseInstanceRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DeletePolardbxSupabaseInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeletePolardbxSupabaseInstanceRequest) SetDBInstanceName(v string) *DeletePolardbxSupabaseInstanceRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DeletePolardbxSupabaseInstanceRequest) SetRegionId(v string) *DeletePolardbxSupabaseInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePolardbxSupabaseInstanceRequest) Validate() error {
	return dara.Validate(s)
}
