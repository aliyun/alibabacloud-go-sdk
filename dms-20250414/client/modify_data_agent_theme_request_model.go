// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataAgentThemeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyDataAgentThemeRequest
	GetDescription() *string
	SetThemeId(v string) *ModifyDataAgentThemeRequest
	GetThemeId() *string
	SetThemeName(v string) *ModifyDataAgentThemeRequest
	GetThemeName() *string
}

type ModifyDataAgentThemeRequest struct {
	// The description of the theme. Maximum length: 255 characters. A value of null indicates that the field is not modified. An empty string clears the field.
	//
	// example:
	//
	// weekly report
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The business identifier of the theme.
	//
	// example:
	//
	// 0f8b2c1d-****-****-****-9a3e5f7b1c2d
	ThemeId *string `json:"ThemeId,omitempty" xml:"ThemeId,omitempty"`
	// The display name of the theme. Maximum length: 64 characters. A value of null indicates that the field is not modified. An empty string clears the field.
	//
	// example:
	//
	// weekly report
	ThemeName *string `json:"ThemeName,omitempty" xml:"ThemeName,omitempty"`
}

func (s ModifyDataAgentThemeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataAgentThemeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDataAgentThemeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyDataAgentThemeRequest) GetThemeId() *string {
	return s.ThemeId
}

func (s *ModifyDataAgentThemeRequest) GetThemeName() *string {
	return s.ThemeName
}

func (s *ModifyDataAgentThemeRequest) SetDescription(v string) *ModifyDataAgentThemeRequest {
	s.Description = &v
	return s
}

func (s *ModifyDataAgentThemeRequest) SetThemeId(v string) *ModifyDataAgentThemeRequest {
	s.ThemeId = &v
	return s
}

func (s *ModifyDataAgentThemeRequest) SetThemeName(v string) *ModifyDataAgentThemeRequest {
	s.ThemeName = &v
	return s
}

func (s *ModifyDataAgentThemeRequest) Validate() error {
	return dara.Validate(s)
}
