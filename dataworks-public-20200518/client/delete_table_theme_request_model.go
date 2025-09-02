// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTableThemeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *DeleteTableThemeRequest
	GetProjectId() *int64
	SetThemeId(v int64) *DeleteTableThemeRequest
	GetThemeId() *int64
}

type DeleteTableThemeRequest struct {
	// The ID of the DataWorks workspace.
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

func (s DeleteTableThemeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTableThemeRequest) GoString() string {
	return s.String()
}

func (s *DeleteTableThemeRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteTableThemeRequest) GetThemeId() *int64 {
	return s.ThemeId
}

func (s *DeleteTableThemeRequest) SetProjectId(v int64) *DeleteTableThemeRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteTableThemeRequest) SetThemeId(v int64) *DeleteTableThemeRequest {
	s.ThemeId = &v
	return s
}

func (s *DeleteTableThemeRequest) Validate() error {
	return dara.Validate(s)
}
