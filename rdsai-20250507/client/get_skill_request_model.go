// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDownload(v bool) *GetSkillRequest
	GetDownload() *bool
	SetLanguage(v string) *GetSkillRequest
	GetLanguage() *string
	SetPath(v string) *GetSkillRequest
	GetPath() *string
	SetSkillId(v string) *GetSkillRequest
	GetSkillId() *string
	SetVersionId(v string) *GetSkillRequest
	GetVersionId() *string
	SetWorkspaceId(v string) *GetSkillRequest
	GetWorkspaceId() *string
}

type GetSkillRequest struct {
	// Specifies whether to obtain a temporary download URL for the specified version. Default value: false.
	Download *bool `json:"Download,omitempty" xml:"Download,omitempty"`
	// The supported language. Valid values:
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
	// The file path within the package, such as SKILL.md. The body is in File.text, and the full Markdown is in SkillMarkdown.
	//
	// example:
	//
	// SKILL.md
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The unique identifier of the skill.
	//
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-44665544****
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// The version UUID (Versions[].Id), not a version number such as 1.0.0. If omitted, the currently active version is read.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000002
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
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

func (s *GetSkillRequest) GetDownload() *bool {
	return s.Download
}

func (s *GetSkillRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetSkillRequest) GetPath() *string {
	return s.Path
}

func (s *GetSkillRequest) GetSkillId() *string {
	return s.SkillId
}

func (s *GetSkillRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *GetSkillRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetSkillRequest) SetDownload(v bool) *GetSkillRequest {
	s.Download = &v
	return s
}

func (s *GetSkillRequest) SetLanguage(v string) *GetSkillRequest {
	s.Language = &v
	return s
}

func (s *GetSkillRequest) SetPath(v string) *GetSkillRequest {
	s.Path = &v
	return s
}

func (s *GetSkillRequest) SetSkillId(v string) *GetSkillRequest {
	s.SkillId = &v
	return s
}

func (s *GetSkillRequest) SetVersionId(v string) *GetSkillRequest {
	s.VersionId = &v
	return s
}

func (s *GetSkillRequest) SetWorkspaceId(v string) *GetSkillRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetSkillRequest) Validate() error {
	return dara.Validate(s)
}
