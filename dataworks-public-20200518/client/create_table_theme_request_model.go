// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTableThemeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLevel(v int32) *CreateTableThemeRequest
	GetLevel() *int32
	SetName(v string) *CreateTableThemeRequest
	GetName() *string
	SetParentId(v int64) *CreateTableThemeRequest
	GetParentId() *int64
	SetProjectId(v int64) *CreateTableThemeRequest
	GetProjectId() *int64
}

type CreateTableThemeRequest struct {
	// The level of the table theme. Valid values: 1 and 2. The value 1 indicates the first level. The value 2 indicates the second level.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Level *int32 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The name of the table theme.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the level of the parent table theme.
	//
	// example:
	//
	// 122
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateTableThemeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTableThemeRequest) GoString() string {
	return s.String()
}

func (s *CreateTableThemeRequest) GetLevel() *int32 {
	return s.Level
}

func (s *CreateTableThemeRequest) GetName() *string {
	return s.Name
}

func (s *CreateTableThemeRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *CreateTableThemeRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateTableThemeRequest) SetLevel(v int32) *CreateTableThemeRequest {
	s.Level = &v
	return s
}

func (s *CreateTableThemeRequest) SetName(v string) *CreateTableThemeRequest {
	s.Name = &v
	return s
}

func (s *CreateTableThemeRequest) SetParentId(v int64) *CreateTableThemeRequest {
	s.ParentId = &v
	return s
}

func (s *CreateTableThemeRequest) SetProjectId(v int64) *CreateTableThemeRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateTableThemeRequest) Validate() error {
	return dara.Validate(s)
}
