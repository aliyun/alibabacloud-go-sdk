// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePxfsUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DeletePxfsUserRequest
	GetDBInstanceName() *string
	SetIdentityName(v string) *DeletePxfsUserRequest
	GetIdentityName() *string
	SetRegionId(v string) *DeletePxfsUserRequest
	GetRegionId() *string
	SetTenantId(v string) *DeletePxfsUserRequest
	GetTenantId() *string
}

type DeletePxfsUserRequest struct {
	// The PolarDB-X instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-xxxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The identity name.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent-runtime
	IdentityName *string `json:"IdentityName,omitempty" xml:"IdentityName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// tenant-001
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DeletePxfsUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePxfsUserRequest) GoString() string {
	return s.String()
}

func (s *DeletePxfsUserRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DeletePxfsUserRequest) GetIdentityName() *string {
	return s.IdentityName
}

func (s *DeletePxfsUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeletePxfsUserRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeletePxfsUserRequest) SetDBInstanceName(v string) *DeletePxfsUserRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DeletePxfsUserRequest) SetIdentityName(v string) *DeletePxfsUserRequest {
	s.IdentityName = &v
	return s
}

func (s *DeletePxfsUserRequest) SetRegionId(v string) *DeletePxfsUserRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePxfsUserRequest) SetTenantId(v string) *DeletePxfsUserRequest {
	s.TenantId = &v
	return s
}

func (s *DeletePxfsUserRequest) Validate() error {
	return dara.Validate(s)
}
