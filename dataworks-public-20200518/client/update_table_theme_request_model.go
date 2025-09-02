// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableThemeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *UpdateTableThemeRequest
	GetName() *string
	SetProjectId(v int64) *UpdateTableThemeRequest
	GetProjectId() *int64
	SetThemeId(v int64) *UpdateTableThemeRequest
	GetThemeId() *int64
}

type UpdateTableThemeRequest struct {
	// The name of the theme.
	//
	// This parameter is required.
	//
	// example:
	//
	// table folder name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The ID of the theme.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ThemeId *int64 `json:"ThemeId,omitempty" xml:"ThemeId,omitempty"`
}

func (s UpdateTableThemeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableThemeRequest) GoString() string {
	return s.String()
}

func (s *UpdateTableThemeRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTableThemeRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateTableThemeRequest) GetThemeId() *int64 {
	return s.ThemeId
}

func (s *UpdateTableThemeRequest) SetName(v string) *UpdateTableThemeRequest {
	s.Name = &v
	return s
}

func (s *UpdateTableThemeRequest) SetProjectId(v int64) *UpdateTableThemeRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateTableThemeRequest) SetThemeId(v int64) *UpdateTableThemeRequest {
	s.ThemeId = &v
	return s
}

func (s *UpdateTableThemeRequest) Validate() error {
	return dara.Validate(s)
}
