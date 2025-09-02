// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableLevelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateTableLevelRequest
	GetDescription() *string
	SetLevelId(v int64) *UpdateTableLevelRequest
	GetLevelId() *int64
	SetLevelType(v int32) *UpdateTableLevelRequest
	GetLevelType() *int32
	SetName(v string) *UpdateTableLevelRequest
	GetName() *string
	SetProjectId(v int64) *UpdateTableLevelRequest
	GetProjectId() *int64
}

type UpdateTableLevelRequest struct {
	// The description of the table level.
	//
	// example:
	//
	// level description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the table level. You can call the ListTableLevel operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	LevelId *int64 `json:"LevelId,omitempty" xml:"LevelId,omitempty"`
	// The table level type. Valid values: 1 and 2. The value 1 indicates the logical level. The value 2 indicates the physical level.
	//
	// example:
	//
	// 1
	LevelType *int32 `json:"LevelType,omitempty" xml:"LevelType,omitempty"`
	// The name of the table level.
	//
	// example:
	//
	// level name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the DataWorks workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateTableLevelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableLevelRequest) GoString() string {
	return s.String()
}

func (s *UpdateTableLevelRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateTableLevelRequest) GetLevelId() *int64 {
	return s.LevelId
}

func (s *UpdateTableLevelRequest) GetLevelType() *int32 {
	return s.LevelType
}

func (s *UpdateTableLevelRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTableLevelRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateTableLevelRequest) SetDescription(v string) *UpdateTableLevelRequest {
	s.Description = &v
	return s
}

func (s *UpdateTableLevelRequest) SetLevelId(v int64) *UpdateTableLevelRequest {
	s.LevelId = &v
	return s
}

func (s *UpdateTableLevelRequest) SetLevelType(v int32) *UpdateTableLevelRequest {
	s.LevelType = &v
	return s
}

func (s *UpdateTableLevelRequest) SetName(v string) *UpdateTableLevelRequest {
	s.Name = &v
	return s
}

func (s *UpdateTableLevelRequest) SetProjectId(v int64) *UpdateTableLevelRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateTableLevelRequest) Validate() error {
	return dara.Validate(s)
}
