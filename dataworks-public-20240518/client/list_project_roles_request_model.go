// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCodes(v []*string) *ListProjectRolesRequest
	GetCodes() []*string
	SetNames(v []*string) *ListProjectRolesRequest
	GetNames() []*string
	SetPageNumber(v int32) *ListProjectRolesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProjectRolesRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListProjectRolesRequest
	GetProjectId() *int64
	SetType(v string) *ListProjectRolesRequest
	GetType() *string
}

type ListProjectRolesRequest struct {
	// An array of workspace role codes.
	Codes []*string `json:"Codes,omitempty" xml:"Codes,omitempty" type:"Repeated"`
	// An array of workspace role names.
	Names []*string `json:"Names,omitempty" xml:"Names,omitempty" type:"Repeated"`
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

func (s ListProjectRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectRolesRequest) GoString() string {
	return s.String()
}

func (s *ListProjectRolesRequest) GetCodes() []*string {
	return s.Codes
}

func (s *ListProjectRolesRequest) GetNames() []*string {
	return s.Names
}

func (s *ListProjectRolesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProjectRolesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProjectRolesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListProjectRolesRequest) GetType() *string {
	return s.Type
}

func (s *ListProjectRolesRequest) SetCodes(v []*string) *ListProjectRolesRequest {
	s.Codes = v
	return s
}

func (s *ListProjectRolesRequest) SetNames(v []*string) *ListProjectRolesRequest {
	s.Names = v
	return s
}

func (s *ListProjectRolesRequest) SetPageNumber(v int32) *ListProjectRolesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectRolesRequest) SetPageSize(v int32) *ListProjectRolesRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectRolesRequest) SetProjectId(v int64) *ListProjectRolesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListProjectRolesRequest) SetType(v string) *ListProjectRolesRequest {
	s.Type = &v
	return s
}

func (s *ListProjectRolesRequest) Validate() error {
	return dara.Validate(s)
}
