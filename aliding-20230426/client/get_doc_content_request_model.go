// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryUuid(v string) *GetDocContentRequest
	GetDentryUuid() *string
	SetTargetFormat(v string) *GetDocContentRequest
	GetTargetFormat() *string
	SetTenantContext(v *GetDocContentRequestTenantContext) *GetDocContentRequest
	GetTenantContext() *GetDocContentRequestTenantContext
	SetUserToken(v string) *GetDocContentRequest
	GetUserToken() *string
}

type GetDocContentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dentry_uuid
	DentryUuid *string `json:"DentryUuid,omitempty" xml:"DentryUuid,omitempty"`
	// example:
	//
	// markdown
	TargetFormat  *string                            `json:"TargetFormat,omitempty" xml:"TargetFormat,omitempty"`
	TenantContext *GetDocContentRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	UserToken *string `json:"userToken,omitempty" xml:"userToken,omitempty"`
}

func (s GetDocContentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocContentRequest) GoString() string {
	return s.String()
}

func (s *GetDocContentRequest) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *GetDocContentRequest) GetTargetFormat() *string {
	return s.TargetFormat
}

func (s *GetDocContentRequest) GetTenantContext() *GetDocContentRequestTenantContext {
	return s.TenantContext
}

func (s *GetDocContentRequest) GetUserToken() *string {
	return s.UserToken
}

func (s *GetDocContentRequest) SetDentryUuid(v string) *GetDocContentRequest {
	s.DentryUuid = &v
	return s
}

func (s *GetDocContentRequest) SetTargetFormat(v string) *GetDocContentRequest {
	s.TargetFormat = &v
	return s
}

func (s *GetDocContentRequest) SetTenantContext(v *GetDocContentRequestTenantContext) *GetDocContentRequest {
	s.TenantContext = v
	return s
}

func (s *GetDocContentRequest) SetUserToken(v string) *GetDocContentRequest {
	s.UserToken = &v
	return s
}

func (s *GetDocContentRequest) Validate() error {
	return dara.Validate(s)
}

type GetDocContentRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetDocContentRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetDocContentRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetDocContentRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetDocContentRequestTenantContext) SetTenantId(v string) *GetDocContentRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetDocContentRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
