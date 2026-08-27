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
	// Specifies whether to return the complete file tree. Default value: False. This avoids large response payloads.
	//
	// example:
	//
	// false
	IncludeSkillFiles *bool `json:"includeSkillFiles,omitempty" xml:"includeSkillFiles,omitempty"`
	// The skill code. This parameter has a value when type is set to skill.
	//
	// example:
	//
	// string_value
	SkillCode *string `json:"skillCode,omitempty" xml:"skillCode,omitempty"`
	// The skill name.
	//
	// example:
	//
	// string_value
	SkillName *string `json:"skillName,omitempty" xml:"skillName,omitempty"`
	// The tenant ID to which the task belongs.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The view mode. Valid values: draft (draft/editing view) or published (published view, default).
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
