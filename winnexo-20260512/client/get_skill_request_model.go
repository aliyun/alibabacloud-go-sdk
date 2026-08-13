// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncludeSkillFiles(v bool) *GetSkillRequest
	GetIncludeSkillFiles() *bool
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
	// 是否返回完整文件树（默认 False，避免大体积响应）
	//
	// example:
	//
	// false
	IncludeSkillFiles *bool `json:"includeSkillFiles,omitempty" xml:"includeSkillFiles,omitempty"`
	// 技能编码（全局唯一），优先级高于 skillName
	//
	// example:
	//
	// string_value
	SkillCode *string `json:"skillCode,omitempty" xml:"skillCode,omitempty"`
	// 技能名称，未传 skillCode 时使用；租户范围内必须唯一
	//
	// example:
	//
	// string_value
	SkillName *string `json:"skillName,omitempty" xml:"skillName,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 视角：draft（草稿/编辑视角）或 published（已发布视角，默认）
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

func (s *GetSkillRequest) GetIncludeSkillFiles() *bool {
	return s.IncludeSkillFiles
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

func (s *GetSkillRequest) SetIncludeSkillFiles(v bool) *GetSkillRequest {
	s.IncludeSkillFiles = &v
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
