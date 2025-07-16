// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrgHonorTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *CreateOrgHonorTemplateShrinkRequest
	GetTenantContextShrink() *string
	SetAvatarFrameMediaId(v string) *CreateOrgHonorTemplateShrinkRequest
	GetAvatarFrameMediaId() *string
	SetDefaultBgColor(v string) *CreateOrgHonorTemplateShrinkRequest
	GetDefaultBgColor() *string
	SetMedalDesc(v string) *CreateOrgHonorTemplateShrinkRequest
	GetMedalDesc() *string
	SetMedalMediaId(v string) *CreateOrgHonorTemplateShrinkRequest
	GetMedalMediaId() *string
	SetMedalName(v string) *CreateOrgHonorTemplateShrinkRequest
	GetMedalName() *string
	SetOrgId(v int64) *CreateOrgHonorTemplateShrinkRequest
	GetOrgId() *int64
	SetUserId(v string) *CreateOrgHonorTemplateShrinkRequest
	GetUserId() *string
}

type CreateOrgHonorTemplateShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fsdfasdjf132342d
	AvatarFrameMediaId *string `json:"avatarFrameMediaId,omitempty" xml:"avatarFrameMediaId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// #FFFBB4
	DefaultBgColor *string `json:"defaultBgColor,omitempty" xml:"defaultBgColor,omitempty"`
	// This parameter is required.
	MedalDesc *string `json:"medalDesc,omitempty" xml:"medalDesc,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1273adf23
	MedalMediaId *string `json:"medalMediaId,omitempty" xml:"medalMediaId,omitempty"`
	// This parameter is required.
	MedalName *string `json:"medalName,omitempty" xml:"medalName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	OrgId *int64 `json:"orgId,omitempty" xml:"orgId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 363784
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateOrgHonorTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrgHonorTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateOrgHonorTemplateShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *CreateOrgHonorTemplateShrinkRequest) GetAvatarFrameMediaId() *string {
	return s.AvatarFrameMediaId
}

func (s *CreateOrgHonorTemplateShrinkRequest) GetDefaultBgColor() *string {
	return s.DefaultBgColor
}

func (s *CreateOrgHonorTemplateShrinkRequest) GetMedalDesc() *string {
	return s.MedalDesc
}

func (s *CreateOrgHonorTemplateShrinkRequest) GetMedalMediaId() *string {
	return s.MedalMediaId
}

func (s *CreateOrgHonorTemplateShrinkRequest) GetMedalName() *string {
	return s.MedalName
}

func (s *CreateOrgHonorTemplateShrinkRequest) GetOrgId() *int64 {
	return s.OrgId
}

func (s *CreateOrgHonorTemplateShrinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateOrgHonorTemplateShrinkRequest) SetTenantContextShrink(v string) *CreateOrgHonorTemplateShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CreateOrgHonorTemplateShrinkRequest) SetAvatarFrameMediaId(v string) *CreateOrgHonorTemplateShrinkRequest {
	s.AvatarFrameMediaId = &v
	return s
}

func (s *CreateOrgHonorTemplateShrinkRequest) SetDefaultBgColor(v string) *CreateOrgHonorTemplateShrinkRequest {
	s.DefaultBgColor = &v
	return s
}

func (s *CreateOrgHonorTemplateShrinkRequest) SetMedalDesc(v string) *CreateOrgHonorTemplateShrinkRequest {
	s.MedalDesc = &v
	return s
}

func (s *CreateOrgHonorTemplateShrinkRequest) SetMedalMediaId(v string) *CreateOrgHonorTemplateShrinkRequest {
	s.MedalMediaId = &v
	return s
}

func (s *CreateOrgHonorTemplateShrinkRequest) SetMedalName(v string) *CreateOrgHonorTemplateShrinkRequest {
	s.MedalName = &v
	return s
}

func (s *CreateOrgHonorTemplateShrinkRequest) SetOrgId(v int64) *CreateOrgHonorTemplateShrinkRequest {
	s.OrgId = &v
	return s
}

func (s *CreateOrgHonorTemplateShrinkRequest) SetUserId(v string) *CreateOrgHonorTemplateShrinkRequest {
	s.UserId = &v
	return s
}

func (s *CreateOrgHonorTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
