// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSupabaseNetTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DeleteSupabaseNetTypeRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DeleteSupabaseNetTypeRequest
	GetRegionId() *string
}

type DeleteSupabaseNetTypeRequest struct {
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

func (s DeleteSupabaseNetTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSupabaseNetTypeRequest) GoString() string {
	return s.String()
}

func (s *DeleteSupabaseNetTypeRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DeleteSupabaseNetTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSupabaseNetTypeRequest) SetDBInstanceName(v string) *DeleteSupabaseNetTypeRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DeleteSupabaseNetTypeRequest) SetRegionId(v string) *DeleteSupabaseNetTypeRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSupabaseNetTypeRequest) Validate() error {
	return dara.Validate(s)
}
