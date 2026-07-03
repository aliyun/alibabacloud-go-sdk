// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSupabaseNetTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionString(v string) *CreateSupabaseNetTypeRequest
	GetConnectionString() *string
	SetDBInstanceName(v string) *CreateSupabaseNetTypeRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *CreateSupabaseNetTypeRequest
	GetRegionId() *string
}

type CreateSupabaseNetTypeRequest struct {
	// The custom endpoint prefix.
	//
	// example:
	//
	// pxsp-*********
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
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

func (s CreateSupabaseNetTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSupabaseNetTypeRequest) GoString() string {
	return s.String()
}

func (s *CreateSupabaseNetTypeRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *CreateSupabaseNetTypeRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateSupabaseNetTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSupabaseNetTypeRequest) SetConnectionString(v string) *CreateSupabaseNetTypeRequest {
	s.ConnectionString = &v
	return s
}

func (s *CreateSupabaseNetTypeRequest) SetDBInstanceName(v string) *CreateSupabaseNetTypeRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateSupabaseNetTypeRequest) SetRegionId(v string) *CreateSupabaseNetTypeRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSupabaseNetTypeRequest) Validate() error {
	return dara.Validate(s)
}
