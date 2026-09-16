// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *GetSkillRequest
	GetLanguage() *string
	SetSkillId(v string) *GetSkillRequest
	GetSkillId() *string
	SetWorkspaceId(v string) *GetSkillRequest
	GetWorkspaceId() *string
}

type GetSkillRequest struct {
	// The supported languages. Valid values:
	//
	// 	- zh-CN: Simplified Chinese
	//
	// 	- zh-TW: Traditional Chinese
	//
	// 	- en-US: English
	//
	// 	- ja-JP: Japanese
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The unique identifier of the Skill.
	//
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-44665544****
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// The ContextDB workspace ID.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillRequest) GoString() string {
	return s.String()
}

func (s *GetSkillRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetSkillRequest) GetSkillId() *string {
	return s.SkillId
}

func (s *GetSkillRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetSkillRequest) SetLanguage(v string) *GetSkillRequest {
	s.Language = &v
	return s
}

func (s *GetSkillRequest) SetSkillId(v string) *GetSkillRequest {
	s.SkillId = &v
	return s
}

func (s *GetSkillRequest) SetWorkspaceId(v string) *GetSkillRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetSkillRequest) Validate() error {
	return dara.Validate(s)
}
