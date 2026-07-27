// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLakebaseTenantTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *GetLakebaseTenantTokenRequest
	GetDBClusterId() *string
	SetPolarFsInstanceId(v string) *GetLakebaseTenantTokenRequest
	GetPolarFsInstanceId() *string
	SetSubdir(v string) *GetLakebaseTenantTokenRequest
	GetSubdir() *string
	SetTenant(v string) *GetLakebaseTenantTokenRequest
	GetTenant() *string
}

type GetLakebaseTenantTokenRequest struct {
	// The associated PolarDB instance ID.
	//
	// example:
	//
	// pc-xxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The PolarFS instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pfs-xxx
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// The mount subdirectory. Specify an absolute path.
	//
	// This parameter is required.
	//
	// example:
	//
	// /home/project/p1
	Subdir *string `json:"Subdir,omitempty" xml:"Subdir,omitempty"`
	// The tenant identifier.
	//
	// example:
	//
	// tenant-xxx
	Tenant *string `json:"Tenant,omitempty" xml:"Tenant,omitempty"`
}

func (s GetLakebaseTenantTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLakebaseTenantTokenRequest) GoString() string {
	return s.String()
}

func (s *GetLakebaseTenantTokenRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetLakebaseTenantTokenRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *GetLakebaseTenantTokenRequest) GetSubdir() *string {
	return s.Subdir
}

func (s *GetLakebaseTenantTokenRequest) GetTenant() *string {
	return s.Tenant
}

func (s *GetLakebaseTenantTokenRequest) SetDBClusterId(v string) *GetLakebaseTenantTokenRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetLakebaseTenantTokenRequest) SetPolarFsInstanceId(v string) *GetLakebaseTenantTokenRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *GetLakebaseTenantTokenRequest) SetSubdir(v string) *GetLakebaseTenantTokenRequest {
	s.Subdir = &v
	return s
}

func (s *GetLakebaseTenantTokenRequest) SetTenant(v string) *GetLakebaseTenantTokenRequest {
	s.Tenant = &v
	return s
}

func (s *GetLakebaseTenantTokenRequest) Validate() error {
	return dara.Validate(s)
}
