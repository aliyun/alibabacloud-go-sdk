// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataAgentThemeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *CreateDataAgentThemeRequest
	GetCategory() *string
	SetDescription(v string) *CreateDataAgentThemeRequest
	GetDescription() *string
	SetFileFrom(v string) *CreateDataAgentThemeRequest
	GetFileFrom() *string
	SetThemeId(v string) *CreateDataAgentThemeRequest
	GetThemeId() *string
	SetThemeName(v string) *CreateDataAgentThemeRequest
	GetThemeName() *string
	SetThemeType(v string) *CreateDataAgentThemeRequest
	GetThemeType() *string
	SetWorkspaceId(v string) *CreateDataAgentThemeRequest
	GetWorkspaceId() *string
}

type CreateDataAgentThemeRequest struct {
	// The scenario, which affects the filtering when you view the theme list in the console. Valid values:
	//
	// - (Recommended) custom: A user-uploaded custom theme with no preset style or information organization structure.
	//
	// - report: A web report that conforms to the DataAgent information organization structure.
	//
	// - (Not supported) infographic: An infographic that conforms to the DataAgent information organization structure.
	//
	// example:
	//
	// custom
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The description. The value can be up to 255 characters in length.
	//
	// example:
	//
	// weekly report
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The file source, which affects the backend logic for determining whether the theme is valid. Valid values:
	//
	// - upload: The file is uploaded through OSS.
	//
	// - (Not supported) public_url: The file is provided through a public network access OSS URL.
	//
	// - (Not supported) user_oss: The file is provided through a user OSS URL.
	//
	// example:
	//
	// upload
	FileFrom *string `json:"FileFrom,omitempty" xml:"FileFrom,omitempty"`
	// The UUID of the theme. The value must be returned by GetDataAgentThemeUploadSignature, and the file must have been uploaded. If the UUID is forged or the file has not been uploaded, the creation fails.
	//
	// example:
	//
	// 0f8b2c1d************9a3e5f7b1c2d
	ThemeId *string `json:"ThemeId,omitempty" xml:"ThemeId,omitempty"`
	// The display name of the theme. The value can be up to 64 characters in length. This parameter is required when you create a theme.
	//
	// example:
	//
	// weekly report
	ThemeName *string `json:"ThemeName,omitempty" xml:"ThemeName,omitempty"`
	// The type of the custom theme. Valid values:
	//
	// - (Default) template: The theme is a template.
	//
	// - (Not supported) design: The theme is a DESIGN.md file.
	//
	// example:
	//
	// template
	ThemeType *string `json:"ThemeType,omitempty" xml:"ThemeType,omitempty"`
	// The workspace to which the theme belongs. If this parameter is not specified or is set to personal, the personal workspace is used. You can also specify a collaboration workspace ID.
	//
	// example:
	//
	// 99fad******qg6c0l4nlacu
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDataAgentThemeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentThemeRequest) GoString() string {
	return s.String()
}

func (s *CreateDataAgentThemeRequest) GetCategory() *string {
	return s.Category
}

func (s *CreateDataAgentThemeRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDataAgentThemeRequest) GetFileFrom() *string {
	return s.FileFrom
}

func (s *CreateDataAgentThemeRequest) GetThemeId() *string {
	return s.ThemeId
}

func (s *CreateDataAgentThemeRequest) GetThemeName() *string {
	return s.ThemeName
}

func (s *CreateDataAgentThemeRequest) GetThemeType() *string {
	return s.ThemeType
}

func (s *CreateDataAgentThemeRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateDataAgentThemeRequest) SetCategory(v string) *CreateDataAgentThemeRequest {
	s.Category = &v
	return s
}

func (s *CreateDataAgentThemeRequest) SetDescription(v string) *CreateDataAgentThemeRequest {
	s.Description = &v
	return s
}

func (s *CreateDataAgentThemeRequest) SetFileFrom(v string) *CreateDataAgentThemeRequest {
	s.FileFrom = &v
	return s
}

func (s *CreateDataAgentThemeRequest) SetThemeId(v string) *CreateDataAgentThemeRequest {
	s.ThemeId = &v
	return s
}

func (s *CreateDataAgentThemeRequest) SetThemeName(v string) *CreateDataAgentThemeRequest {
	s.ThemeName = &v
	return s
}

func (s *CreateDataAgentThemeRequest) SetThemeType(v string) *CreateDataAgentThemeRequest {
	s.ThemeType = &v
	return s
}

func (s *CreateDataAgentThemeRequest) SetWorkspaceId(v string) *CreateDataAgentThemeRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDataAgentThemeRequest) Validate() error {
	return dara.Validate(s)
}
