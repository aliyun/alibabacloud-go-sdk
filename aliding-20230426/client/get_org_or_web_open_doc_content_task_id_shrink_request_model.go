// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrgOrWebOpenDocContentTaskIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryUuid(v string) *GetOrgOrWebOpenDocContentTaskIdShrinkRequest
	GetDentryUuid() *string
	SetGenerateCp(v bool) *GetOrgOrWebOpenDocContentTaskIdShrinkRequest
	GetGenerateCp() *bool
	SetScopeType(v int32) *GetOrgOrWebOpenDocContentTaskIdShrinkRequest
	GetScopeType() *int32
	SetTargetFormat(v string) *GetOrgOrWebOpenDocContentTaskIdShrinkRequest
	GetTargetFormat() *string
	SetTenantContextShrink(v string) *GetOrgOrWebOpenDocContentTaskIdShrinkRequest
	GetTenantContextShrink() *string
}

type GetOrgOrWebOpenDocContentTaskIdShrinkRequest struct {
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
	TargetFormat        *string `json:"TargetFormat,omitempty" xml:"TargetFormat,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetOrgOrWebOpenDocContentTaskIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOrgOrWebOpenDocContentTaskIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetOrgOrWebOpenDocContentTaskIdShrinkRequest) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *GetOrgOrWebOpenDocContentTaskIdShrinkRequest) GetGenerateCp() *bool {
	return s.GenerateCp
}

func (s *GetOrgOrWebOpenDocContentTaskIdShrinkRequest) GetScopeType() *int32 {
	return s.ScopeType
}

func (s *GetOrgOrWebOpenDocContentTaskIdShrinkRequest) GetTargetFormat() *string {
	return s.TargetFormat
}

func (s *GetOrgOrWebOpenDocContentTaskIdShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetOrgOrWebOpenDocContentTaskIdShrinkRequest) SetDentryUuid(v string) *GetOrgOrWebOpenDocContentTaskIdShrinkRequest {
	s.DentryUuid = &v
	return s
}

func (s *GetOrgOrWebOpenDocContentTaskIdShrinkRequest) SetGenerateCp(v bool) *GetOrgOrWebOpenDocContentTaskIdShrinkRequest {
	s.GenerateCp = &v
	return s
}

func (s *GetOrgOrWebOpenDocContentTaskIdShrinkRequest) SetScopeType(v int32) *GetOrgOrWebOpenDocContentTaskIdShrinkRequest {
	s.ScopeType = &v
	return s
}

func (s *GetOrgOrWebOpenDocContentTaskIdShrinkRequest) SetTargetFormat(v string) *GetOrgOrWebOpenDocContentTaskIdShrinkRequest {
	s.TargetFormat = &v
	return s
}

func (s *GetOrgOrWebOpenDocContentTaskIdShrinkRequest) SetTenantContextShrink(v string) *GetOrgOrWebOpenDocContentTaskIdShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetOrgOrWebOpenDocContentTaskIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
