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
	// The list of workspace role codes.
	CodesShrink *string `json:"Codes,omitempty" xml:"Codes,omitempty"`
	// The list of workspace role names.
	NamesShrink *string `json:"Names,omitempty" xml:"Names,omitempty"`
	// The page number. Used for paging.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the workspace management page to obtain the ID.
	//
	// This parameter specifies the DataWorks workspace for this API invoke operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 21229
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The type of the workspace role. Valid values:
	//
	// - UserCustom: user-defined role.
	//
	// - System: system role.
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
