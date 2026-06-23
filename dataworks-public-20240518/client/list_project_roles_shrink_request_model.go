// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectRolesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCodesShrink(v string) *ListProjectRolesShrinkRequest
	GetCodesShrink() *string
	SetNamesShrink(v string) *ListProjectRolesShrinkRequest
	GetNamesShrink() *string
	SetPageNumber(v int32) *ListProjectRolesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProjectRolesShrinkRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListProjectRolesShrinkRequest
	GetProjectId() *int64
	SetType(v string) *ListProjectRolesShrinkRequest
	GetType() *string
}

type ListProjectRolesShrinkRequest struct {
	// An array of workspace role codes.
	CodesShrink *string `json:"Codes,omitempty" xml:"Codes,omitempty"`
	// An array of workspace role names.
	NamesShrink *string `json:"Names,omitempty" xml:"Names,omitempty"`
	// The page number to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the DataWorks workspace. You can find the ID on the Workspace Management page in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// This parameter specifies the DataWorks workspace for which you want to list roles.
	//
	// This parameter is required.
	//
	// example:
	//
	// 21229
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The type of the workspace role.
	//
	// - `UserCustom`: Custom Role
	//
	// - `System`: System Role
	//
	// example:
	//
	// System
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListProjectRolesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectRolesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListProjectRolesShrinkRequest) GetCodesShrink() *string {
	return s.CodesShrink
}

func (s *ListProjectRolesShrinkRequest) GetNamesShrink() *string {
	return s.NamesShrink
}

func (s *ListProjectRolesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProjectRolesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProjectRolesShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListProjectRolesShrinkRequest) GetType() *string {
	return s.Type
}

func (s *ListProjectRolesShrinkRequest) SetCodesShrink(v string) *ListProjectRolesShrinkRequest {
	s.CodesShrink = &v
	return s
}

func (s *ListProjectRolesShrinkRequest) SetNamesShrink(v string) *ListProjectRolesShrinkRequest {
	s.NamesShrink = &v
	return s
}

func (s *ListProjectRolesShrinkRequest) SetPageNumber(v int32) *ListProjectRolesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectRolesShrinkRequest) SetPageSize(v int32) *ListProjectRolesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectRolesShrinkRequest) SetProjectId(v int64) *ListProjectRolesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListProjectRolesShrinkRequest) SetType(v string) *ListProjectRolesShrinkRequest {
	s.Type = &v
	return s
}

func (s *ListProjectRolesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
