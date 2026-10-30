// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePxfsUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreatePxfsUserRequest
	GetDBInstanceName() *string
	SetIdentityName(v string) *CreatePxfsUserRequest
	GetIdentityName() *string
	SetPosixGid(v int64) *CreatePxfsUserRequest
	GetPosixGid() *int64
	SetPosixUid(v int64) *CreatePxfsUserRequest
	GetPosixUid() *int64
	SetRegionId(v string) *CreatePxfsUserRequest
	GetRegionId() *string
	SetRole(v string) *CreatePxfsUserRequest
	GetRole() *string
	SetTenantId(v string) *CreatePxfsUserRequest
	GetTenantId() *string
}

type CreatePxfsUserRequest struct {
	// The PolarDB-X instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-xxxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The name of the identity.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent-runtime
	IdentityName *string `json:"IdentityName,omitempty" xml:"IdentityName,omitempty"`
	// POSIX GID
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	PosixGid *int64 `json:"PosixGid,omitempty" xml:"PosixGid,omitempty"`
	// POSIX UID
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	PosixUid *int64 `json:"PosixUid,omitempty" xml:"PosixUid,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The role of the identity.
	//
	// This parameter is required.
	//
	// example:
	//
	// MEMBER
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The ID of the tenant.
	//
	// This parameter is required.
	//
	// example:
	//
	// tenant-001
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s CreatePxfsUserRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePxfsUserRequest) GoString() string {
	return s.String()
}

func (s *CreatePxfsUserRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreatePxfsUserRequest) GetIdentityName() *string {
	return s.IdentityName
}

func (s *CreatePxfsUserRequest) GetPosixGid() *int64 {
	return s.PosixGid
}

func (s *CreatePxfsUserRequest) GetPosixUid() *int64 {
	return s.PosixUid
}

func (s *CreatePxfsUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePxfsUserRequest) GetRole() *string {
	return s.Role
}

func (s *CreatePxfsUserRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePxfsUserRequest) SetDBInstanceName(v string) *CreatePxfsUserRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreatePxfsUserRequest) SetIdentityName(v string) *CreatePxfsUserRequest {
	s.IdentityName = &v
	return s
}

func (s *CreatePxfsUserRequest) SetPosixGid(v int64) *CreatePxfsUserRequest {
	s.PosixGid = &v
	return s
}

func (s *CreatePxfsUserRequest) SetPosixUid(v int64) *CreatePxfsUserRequest {
	s.PosixUid = &v
	return s
}

func (s *CreatePxfsUserRequest) SetRegionId(v string) *CreatePxfsUserRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePxfsUserRequest) SetRole(v string) *CreatePxfsUserRequest {
	s.Role = &v
	return s
}

func (s *CreatePxfsUserRequest) SetTenantId(v string) *CreatePxfsUserRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePxfsUserRequest) Validate() error {
	return dara.Validate(s)
}
