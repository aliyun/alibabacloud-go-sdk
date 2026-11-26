// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetSkillRequest
	GetAppToken() *string
	SetIncludeSkillFiles(v bool) *GetSkillRequest
	GetIncludeSkillFiles() *bool
	SetOpsUserCode(v string) *GetSkillRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetSkillRequest
	GetOpsUserId() *string
	SetSkillCode(v string) *GetSkillRequest
	GetSkillCode() *string
	SetSkillName(v string) *GetSkillRequest
	GetSkillName() *string
	SetTenantId(v string) *GetSkillRequest
	GetTenantId() *string
	SetViewMode(v string) *GetSkillRequest
	GetViewMode() *string
}

type GetSkillRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// Specifies whether to return the complete file tree. Default value: false. This prevents large response payloads.
	//
	// example:
	//
	// false
	IncludeSkillFiles *bool `json:"includeSkillFiles,omitempty" xml:"includeSkillFiles,omitempty"`
	// 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The skill code. This parameter has a value when type is set to skill.
	//
	// example:
	//
	// string_value
	SkillCode *string `json:"skillCode,omitempty" xml:"skillCode,omitempty"`
	// The name of the skill.
	//
	// example:
	//
	// string_value
	SkillName *string `json:"skillName,omitempty" xml:"skillName,omitempty"`
	// The ID of the tenant to which the task belongs.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The view mode. Valid values: draft (draft or edit view) and published (published view, default).
	//
	// example:
	//
	// draft
	ViewMode *string `json:"viewMode,omitempty" xml:"viewMode,omitempty"`
}

func (s GetSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillRequest) GoString() string {
	return s.String()
}

func (s *GetSkillRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetSkillRequest) GetIncludeSkillFiles() *bool {
	return s.IncludeSkillFiles
}

func (s *GetSkillRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetSkillRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetSkillRequest) GetSkillCode() *string {
	return s.SkillCode
}

func (s *GetSkillRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *GetSkillRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetSkillRequest) GetViewMode() *string {
	return s.ViewMode
}

func (s *GetSkillRequest) SetAppToken(v string) *GetSkillRequest {
	s.AppToken = &v
	return s
}

func (s *GetSkillRequest) SetIncludeSkillFiles(v bool) *GetSkillRequest {
	s.IncludeSkillFiles = &v
	return s
}

func (s *GetSkillRequest) SetOpsUserCode(v string) *GetSkillRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetSkillRequest) SetOpsUserId(v string) *GetSkillRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetSkillRequest) SetSkillCode(v string) *GetSkillRequest {
	s.SkillCode = &v
	return s
}

func (s *GetSkillRequest) SetSkillName(v string) *GetSkillRequest {
	s.SkillName = &v
	return s
}

func (s *GetSkillRequest) SetTenantId(v string) *GetSkillRequest {
	s.TenantId = &v
	return s
}

func (s *GetSkillRequest) SetViewMode(v string) *GetSkillRequest {
	s.ViewMode = &v
	return s
}

func (s *GetSkillRequest) Validate() error {
	return dara.Validate(s)
}
