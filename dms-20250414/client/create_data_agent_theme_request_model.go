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
}

type CreateDataAgentThemeRequest struct {
	// The application scenario, which affects filtering when viewing the theme list in the console. Valid values:
	//
	// - (Recommended) custom: a user-uploaded custom theme with no preset style or information organization structure.
	//
	// - report: a web report that conforms to the DataAgent information organization structure.
	//
	// - (Not supported) infographic: an infographic that conforms to the DataAgent information organization structure.
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
	// - upload: uploaded through OSS.
	//
	// - (Not supported) public_url: provided through an OSS URL that allows public network access.
	//
	// - (Not supported) user_oss: provided through a user OSS URL.
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
	// The display name of the theme. The value can be up to 64 characters in length. This parameter is required during creation.
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

func (s *CreateDataAgentThemeRequest) Validate() error {
	return dara.Validate(s)
}
