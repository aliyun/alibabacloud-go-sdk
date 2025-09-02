// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTableThemeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int32) *ListTableThemeRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListTableThemeRequest
	GetPageSize() *int32
	SetParentId(v int64) *ListTableThemeRequest
	GetParentId() *int64
	SetProjectId(v int64) *ListTableThemeRequest
	GetProjectId() *int64
}

type ListTableThemeRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the parent table theme.
	//
	// example:
	//
	// 121
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListTableThemeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTableThemeRequest) GoString() string {
	return s.String()
}

func (s *ListTableThemeRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListTableThemeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTableThemeRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *ListTableThemeRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListTableThemeRequest) SetPageNum(v int32) *ListTableThemeRequest {
	s.PageNum = &v
	return s
}

func (s *ListTableThemeRequest) SetPageSize(v int32) *ListTableThemeRequest {
	s.PageSize = &v
	return s
}

func (s *ListTableThemeRequest) SetParentId(v int64) *ListTableThemeRequest {
	s.ParentId = &v
	return s
}

func (s *ListTableThemeRequest) SetProjectId(v int64) *ListTableThemeRequest {
	s.ProjectId = &v
	return s
}

func (s *ListTableThemeRequest) Validate() error {
	return dara.Validate(s)
}
