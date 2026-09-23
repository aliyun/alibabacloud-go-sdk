// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServerVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *GetServerVersionRequest
	GetEnv() *string
	SetOpTenantId(v int64) *GetServerVersionRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetServerVersionRequest
	GetOpUserId() *string
}

type GetServerVersionRequest struct {
	// **[Deprecated]*	- The environment identifier. Valid values:
	//
	// - DEV: Development environment.
	//
	// - PROD (default): Production environment.
	//
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operation user.
	//
	// example:
	//
	// 30010012
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s GetServerVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServerVersionRequest) GoString() string {
	return s.String()
}

func (s *GetServerVersionRequest) GetEnv() *string {
	return s.Env
}

func (s *GetServerVersionRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetServerVersionRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetServerVersionRequest) SetEnv(v string) *GetServerVersionRequest {
	s.Env = &v
	return s
}

func (s *GetServerVersionRequest) SetOpTenantId(v int64) *GetServerVersionRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetServerVersionRequest) SetOpUserId(v string) *GetServerVersionRequest {
	s.OpUserId = &v
	return s
}

func (s *GetServerVersionRequest) Validate() error {
	return dara.Validate(s)
}
