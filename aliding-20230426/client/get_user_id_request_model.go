// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *GetUserIdRequestTenantContext) *GetUserIdRequest
	GetTenantContext() *GetUserIdRequestTenantContext
	SetUnionId(v string) *GetUserIdRequest
	GetUnionId() *string
}

type GetUserIdRequest struct {
	TenantContext *GetUserIdRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// unionId
	//
	// example:
	//
	// ****iE
	UnionId *string `json:"UnionId,omitempty" xml:"UnionId,omitempty"`
}

func (s GetUserIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdRequest) GoString() string {
	return s.String()
}

func (s *GetUserIdRequest) GetTenantContext() *GetUserIdRequestTenantContext {
	return s.TenantContext
}

func (s *GetUserIdRequest) GetUnionId() *string {
	return s.UnionId
}

func (s *GetUserIdRequest) SetTenantContext(v *GetUserIdRequestTenantContext) *GetUserIdRequest {
	s.TenantContext = v
	return s
}

func (s *GetUserIdRequest) SetUnionId(v string) *GetUserIdRequest {
	s.UnionId = &v
	return s
}

func (s *GetUserIdRequest) Validate() error {
	return dara.Validate(s)
}

type GetUserIdRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetUserIdRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetUserIdRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetUserIdRequestTenantContext) SetTenantId(v string) *GetUserIdRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetUserIdRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
