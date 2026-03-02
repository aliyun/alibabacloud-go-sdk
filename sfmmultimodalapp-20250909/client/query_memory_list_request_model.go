// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMemoryListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryMemoryListRequest
	GetAppId() *string
	SetPageNumber(v int32) *QueryMemoryListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryMemoryListRequest
	GetPageSize() *int32
	SetProjectId(v string) *QueryMemoryListRequest
	GetProjectId() *string
	SetUserDefinedId(v string) *QueryMemoryListRequest
	GetUserDefinedId() *string
	SetWorkspaceId(v string) *QueryMemoryListRequest
	GetWorkspaceId() *string
}

type QueryMemoryListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_bfaf7e110b6d4359977d1686a3f8
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// profile_project
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// b6d4359977d16
	UserDefinedId *string `json:"UserDefinedId,omitempty" xml:"UserDefinedId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-jb5sabg80b4ts71g
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryMemoryListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMemoryListRequest) GoString() string {
	return s.String()
}

func (s *QueryMemoryListRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryMemoryListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryMemoryListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMemoryListRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *QueryMemoryListRequest) GetUserDefinedId() *string {
	return s.UserDefinedId
}

func (s *QueryMemoryListRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryMemoryListRequest) SetAppId(v string) *QueryMemoryListRequest {
	s.AppId = &v
	return s
}

func (s *QueryMemoryListRequest) SetPageNumber(v int32) *QueryMemoryListRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryMemoryListRequest) SetPageSize(v int32) *QueryMemoryListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMemoryListRequest) SetProjectId(v string) *QueryMemoryListRequest {
	s.ProjectId = &v
	return s
}

func (s *QueryMemoryListRequest) SetUserDefinedId(v string) *QueryMemoryListRequest {
	s.UserDefinedId = &v
	return s
}

func (s *QueryMemoryListRequest) SetWorkspaceId(v string) *QueryMemoryListRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryMemoryListRequest) Validate() error {
	return dara.Validate(s)
}
