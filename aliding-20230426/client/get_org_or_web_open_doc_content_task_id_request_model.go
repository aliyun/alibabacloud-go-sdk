// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrgOrWebOpenDocContentTaskIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryUuid(v string) *GetOrgOrWebOpenDocContentTaskIdRequest
	GetDentryUuid() *string
	SetGenerateCp(v bool) *GetOrgOrWebOpenDocContentTaskIdRequest
	GetGenerateCp() *bool
	SetScopeType(v int32) *GetOrgOrWebOpenDocContentTaskIdRequest
	GetScopeType() *int32
	SetTargetFormat(v string) *GetOrgOrWebOpenDocContentTaskIdRequest
	GetTargetFormat() *string
	SetTenantContext(v *GetOrgOrWebOpenDocContentTaskIdRequestTenantContext) *GetOrgOrWebOpenDocContentTaskIdRequest
	GetTenantContext() *GetOrgOrWebOpenDocContentTaskIdRequestTenantContext
}

type GetOrgOrWebOpenDocContentTaskIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20eMKjyp810mMdK4Hz4B5BA6JxAZB1Gv
	DentryUuid *string `json:"DentryUuid,omitempty" xml:"DentryUuid,omitempty"`
	GenerateCp *bool   `json:"GenerateCp,omitempty" xml:"GenerateCp,omitempty"`
	// example:
	//
	// 0
	ScopeType *int32 `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
	// example:
	//
	// markdown
	TargetFormat  *string                                              `json:"TargetFormat,omitempty" xml:"TargetFormat,omitempty"`
	TenantContext *GetOrgOrWebOpenDocContentTaskIdRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetOrgOrWebOpenDocContentTaskIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOrgOrWebOpenDocContentTaskIdRequest) GoString() string {
	return s.String()
}

func (s *GetOrgOrWebOpenDocContentTaskIdRequest) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *GetOrgOrWebOpenDocContentTaskIdRequest) GetGenerateCp() *bool {
	return s.GenerateCp
}

func (s *GetOrgOrWebOpenDocContentTaskIdRequest) GetScopeType() *int32 {
	return s.ScopeType
}

func (s *GetOrgOrWebOpenDocContentTaskIdRequest) GetTargetFormat() *string {
	return s.TargetFormat
}

func (s *GetOrgOrWebOpenDocContentTaskIdRequest) GetTenantContext() *GetOrgOrWebOpenDocContentTaskIdRequestTenantContext {
	return s.TenantContext
}

func (s *GetOrgOrWebOpenDocContentTaskIdRequest) SetDentryUuid(v string) *GetOrgOrWebOpenDocContentTaskIdRequest {
	s.DentryUuid = &v
	return s
}

func (s *GetOrgOrWebOpenDocContentTaskIdRequest) SetGenerateCp(v bool) *GetOrgOrWebOpenDocContentTaskIdRequest {
	s.GenerateCp = &v
	return s
}

func (s *GetOrgOrWebOpenDocContentTaskIdRequest) SetScopeType(v int32) *GetOrgOrWebOpenDocContentTaskIdRequest {
	s.ScopeType = &v
	return s
}

func (s *GetOrgOrWebOpenDocContentTaskIdRequest) SetTargetFormat(v string) *GetOrgOrWebOpenDocContentTaskIdRequest {
	s.TargetFormat = &v
	return s
}

func (s *GetOrgOrWebOpenDocContentTaskIdRequest) SetTenantContext(v *GetOrgOrWebOpenDocContentTaskIdRequestTenantContext) *GetOrgOrWebOpenDocContentTaskIdRequest {
	s.TenantContext = v
	return s
}

func (s *GetOrgOrWebOpenDocContentTaskIdRequest) Validate() error {
	return dara.Validate(s)
}

type GetOrgOrWebOpenDocContentTaskIdRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetOrgOrWebOpenDocContentTaskIdRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetOrgOrWebOpenDocContentTaskIdRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetOrgOrWebOpenDocContentTaskIdRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetOrgOrWebOpenDocContentTaskIdRequestTenantContext) SetTenantId(v string) *GetOrgOrWebOpenDocContentTaskIdRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetOrgOrWebOpenDocContentTaskIdRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
