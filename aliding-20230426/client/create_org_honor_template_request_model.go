// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrgHonorTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *CreateOrgHonorTemplateRequestTenantContext) *CreateOrgHonorTemplateRequest
	GetTenantContext() *CreateOrgHonorTemplateRequestTenantContext
	SetAvatarFrameMediaId(v string) *CreateOrgHonorTemplateRequest
	GetAvatarFrameMediaId() *string
	SetDefaultBgColor(v string) *CreateOrgHonorTemplateRequest
	GetDefaultBgColor() *string
	SetMedalDesc(v string) *CreateOrgHonorTemplateRequest
	GetMedalDesc() *string
	SetMedalMediaId(v string) *CreateOrgHonorTemplateRequest
	GetMedalMediaId() *string
	SetMedalName(v string) *CreateOrgHonorTemplateRequest
	GetMedalName() *string
	SetOrgId(v int64) *CreateOrgHonorTemplateRequest
	GetOrgId() *int64
	SetUserId(v string) *CreateOrgHonorTemplateRequest
	GetUserId() *string
}

type CreateOrgHonorTemplateRequest struct {
	TenantContext *CreateOrgHonorTemplateRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
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

func (s CreateOrgHonorTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrgHonorTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateOrgHonorTemplateRequest) GetTenantContext() *CreateOrgHonorTemplateRequestTenantContext {
	return s.TenantContext
}

func (s *CreateOrgHonorTemplateRequest) GetAvatarFrameMediaId() *string {
	return s.AvatarFrameMediaId
}

func (s *CreateOrgHonorTemplateRequest) GetDefaultBgColor() *string {
	return s.DefaultBgColor
}

func (s *CreateOrgHonorTemplateRequest) GetMedalDesc() *string {
	return s.MedalDesc
}

func (s *CreateOrgHonorTemplateRequest) GetMedalMediaId() *string {
	return s.MedalMediaId
}

func (s *CreateOrgHonorTemplateRequest) GetMedalName() *string {
	return s.MedalName
}

func (s *CreateOrgHonorTemplateRequest) GetOrgId() *int64 {
	return s.OrgId
}

func (s *CreateOrgHonorTemplateRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateOrgHonorTemplateRequest) SetTenantContext(v *CreateOrgHonorTemplateRequestTenantContext) *CreateOrgHonorTemplateRequest {
	s.TenantContext = v
	return s
}

func (s *CreateOrgHonorTemplateRequest) SetAvatarFrameMediaId(v string) *CreateOrgHonorTemplateRequest {
	s.AvatarFrameMediaId = &v
	return s
}

func (s *CreateOrgHonorTemplateRequest) SetDefaultBgColor(v string) *CreateOrgHonorTemplateRequest {
	s.DefaultBgColor = &v
	return s
}

func (s *CreateOrgHonorTemplateRequest) SetMedalDesc(v string) *CreateOrgHonorTemplateRequest {
	s.MedalDesc = &v
	return s
}

func (s *CreateOrgHonorTemplateRequest) SetMedalMediaId(v string) *CreateOrgHonorTemplateRequest {
	s.MedalMediaId = &v
	return s
}

func (s *CreateOrgHonorTemplateRequest) SetMedalName(v string) *CreateOrgHonorTemplateRequest {
	s.MedalName = &v
	return s
}

func (s *CreateOrgHonorTemplateRequest) SetOrgId(v int64) *CreateOrgHonorTemplateRequest {
	s.OrgId = &v
	return s
}

func (s *CreateOrgHonorTemplateRequest) SetUserId(v string) *CreateOrgHonorTemplateRequest {
	s.UserId = &v
	return s
}

func (s *CreateOrgHonorTemplateRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOrgHonorTemplateRequestTenantContext struct {
	// example:
	//
	// 123456
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateOrgHonorTemplateRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s CreateOrgHonorTemplateRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CreateOrgHonorTemplateRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateOrgHonorTemplateRequestTenantContext) SetTenantId(v string) *CreateOrgHonorTemplateRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *CreateOrgHonorTemplateRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
