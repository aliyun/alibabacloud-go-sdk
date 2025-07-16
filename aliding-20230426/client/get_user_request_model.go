// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *GetUserRequestTenantContext) *GetUserRequest
	GetTenantContext() *GetUserRequestTenantContext
	SetLanguage(v string) *GetUserRequest
	GetLanguage() *string
}

type GetUserRequest struct {
	TenantContext *GetUserRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s GetUserRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) GetTenantContext() *GetUserRequestTenantContext {
	return s.TenantContext
}

func (s *GetUserRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetUserRequest) SetTenantContext(v *GetUserRequestTenantContext) *GetUserRequest {
	s.TenantContext = v
	return s
}

func (s *GetUserRequest) SetLanguage(v string) *GetUserRequest {
	s.Language = &v
	return s
}

func (s *GetUserRequest) Validate() error {
	return dara.Validate(s)
}

type GetUserRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetUserRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetUserRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetUserRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetUserRequestTenantContext) SetTenantId(v string) *GetUserRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetUserRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
