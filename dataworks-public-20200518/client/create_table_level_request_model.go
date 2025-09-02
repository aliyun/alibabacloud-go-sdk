// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTableLevelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateTableLevelRequest
	GetDescription() *string
	SetLevelType(v int32) *CreateTableLevelRequest
	GetLevelType() *int32
	SetName(v string) *CreateTableLevelRequest
	GetName() *string
	SetProjectId(v int64) *CreateTableLevelRequest
	GetProjectId() *int64
}

type CreateTableLevelRequest struct {
	// The description of the table level.
	//
	// example:
	//
	// The HTTP status code returned.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the table level. Valid values: 1 and 2. The value 1 indicates the logical level. The value 2 indicates the physical level.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	LevelType *int32 `json:"LevelType,omitempty" xml:"LevelType,omitempty"`
	// The name of the table level.
	//
	// This parameter is required.
	//
	// example:
	//
	// The description of the table level.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateTableLevelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTableLevelRequest) GoString() string {
	return s.String()
}

func (s *CreateTableLevelRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTableLevelRequest) GetLevelType() *int32 {
	return s.LevelType
}

func (s *CreateTableLevelRequest) GetName() *string {
	return s.Name
}

func (s *CreateTableLevelRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateTableLevelRequest) SetDescription(v string) *CreateTableLevelRequest {
	s.Description = &v
	return s
}

func (s *CreateTableLevelRequest) SetLevelType(v int32) *CreateTableLevelRequest {
	s.LevelType = &v
	return s
}

func (s *CreateTableLevelRequest) SetName(v string) *CreateTableLevelRequest {
	s.Name = &v
	return s
}

func (s *CreateTableLevelRequest) SetProjectId(v int64) *CreateTableLevelRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateTableLevelRequest) Validate() error {
	return dara.Validate(s)
}
