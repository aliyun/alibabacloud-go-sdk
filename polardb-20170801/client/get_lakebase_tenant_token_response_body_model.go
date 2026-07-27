// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLakebaseTenantTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *GetLakebaseTenantTokenResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *GetLakebaseTenantTokenResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetLakebaseTenantTokenResponseBody
	GetStatus() *string
	SetSubdir(v string) *GetLakebaseTenantTokenResponseBody
	GetSubdir() *string
	SetTenant(v string) *GetLakebaseTenantTokenResponseBody
	GetTenant() *string
	SetToken(v string) *GetLakebaseTenantTokenResponseBody
	GetToken() *string
}

type GetLakebaseTenantTokenResponseBody struct {
	// The associated PolarDB instance ID.
	//
	// example:
	//
	// pc-xxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CD35F3-F3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status.
	//
	// example:
	//
	// ACTIVATION
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The mount subdirectory.
	//
	// example:
	//
	// /home/project
	Subdir *string `json:"Subdir,omitempty" xml:"Subdir,omitempty"`
	// The tenant identifier.
	//
	// example:
	//
	// admin
	Tenant *string `json:"Tenant,omitempty" xml:"Tenant,omitempty"`
	// The tenant token.
	//
	// example:
	//
	// ***
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetLakebaseTenantTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLakebaseTenantTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetLakebaseTenantTokenResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetLakebaseTenantTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLakebaseTenantTokenResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetLakebaseTenantTokenResponseBody) GetSubdir() *string {
	return s.Subdir
}

func (s *GetLakebaseTenantTokenResponseBody) GetTenant() *string {
	return s.Tenant
}

func (s *GetLakebaseTenantTokenResponseBody) GetToken() *string {
	return s.Token
}

func (s *GetLakebaseTenantTokenResponseBody) SetDBClusterId(v string) *GetLakebaseTenantTokenResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *GetLakebaseTenantTokenResponseBody) SetRequestId(v string) *GetLakebaseTenantTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLakebaseTenantTokenResponseBody) SetStatus(v string) *GetLakebaseTenantTokenResponseBody {
	s.Status = &v
	return s
}

func (s *GetLakebaseTenantTokenResponseBody) SetSubdir(v string) *GetLakebaseTenantTokenResponseBody {
	s.Subdir = &v
	return s
}

func (s *GetLakebaseTenantTokenResponseBody) SetTenant(v string) *GetLakebaseTenantTokenResponseBody {
	s.Tenant = &v
	return s
}

func (s *GetLakebaseTenantTokenResponseBody) SetToken(v string) *GetLakebaseTenantTokenResponseBody {
	s.Token = &v
	return s
}

func (s *GetLakebaseTenantTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
