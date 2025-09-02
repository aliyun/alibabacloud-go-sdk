// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTableLevelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLevelType(v int32) *ListTableLevelRequest
	GetLevelType() *int32
	SetPageNum(v int32) *ListTableLevelRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListTableLevelRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListTableLevelRequest
	GetProjectId() *int64
}

type ListTableLevelRequest struct {
	// The table level type. Valid values: 1 and 2. The value 1 indicates the logical level. The value 2 indicates the physical level.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	LevelType *int32 `json:"LevelType,omitempty" xml:"LevelType,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the DataWorks console to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListTableLevelRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTableLevelRequest) GoString() string {
	return s.String()
}

func (s *ListTableLevelRequest) GetLevelType() *int32 {
	return s.LevelType
}

func (s *ListTableLevelRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListTableLevelRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTableLevelRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListTableLevelRequest) SetLevelType(v int32) *ListTableLevelRequest {
	s.LevelType = &v
	return s
}

func (s *ListTableLevelRequest) SetPageNum(v int32) *ListTableLevelRequest {
	s.PageNum = &v
	return s
}

func (s *ListTableLevelRequest) SetPageSize(v int32) *ListTableLevelRequest {
	s.PageSize = &v
	return s
}

func (s *ListTableLevelRequest) SetProjectId(v int64) *ListTableLevelRequest {
	s.ProjectId = &v
	return s
}

func (s *ListTableLevelRequest) Validate() error {
	return dara.Validate(s)
}
