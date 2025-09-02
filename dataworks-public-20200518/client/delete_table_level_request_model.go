// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTableLevelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLevelId(v int64) *DeleteTableLevelRequest
	GetLevelId() *int64
	SetProjectId(v int64) *DeleteTableLevelRequest
	GetProjectId() *int64
}

type DeleteTableLevelRequest struct {
	// The ID of the table level that you want to delete. You can call the ListTableLevel operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	LevelId *int64 `json:"LevelId,omitempty" xml:"LevelId,omitempty"`
	// The DataWorks workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteTableLevelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTableLevelRequest) GoString() string {
	return s.String()
}

func (s *DeleteTableLevelRequest) GetLevelId() *int64 {
	return s.LevelId
}

func (s *DeleteTableLevelRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteTableLevelRequest) SetLevelId(v int64) *DeleteTableLevelRequest {
	s.LevelId = &v
	return s
}

func (s *DeleteTableLevelRequest) SetProjectId(v int64) *DeleteTableLevelRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteTableLevelRequest) Validate() error {
	return dara.Validate(s)
}
