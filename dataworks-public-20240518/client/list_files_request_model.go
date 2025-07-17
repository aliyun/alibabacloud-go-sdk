// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommitStatus(v int32) *ListFilesRequest
	GetCommitStatus() *int32
	SetExactFileName(v string) *ListFilesRequest
	GetExactFileName() *string
	SetFileFolderPath(v string) *ListFilesRequest
	GetFileFolderPath() *string
	SetFileIdIn(v string) *ListFilesRequest
	GetFileIdIn() *string
	SetFileTypes(v string) *ListFilesRequest
	GetFileTypes() *string
	SetKeyword(v string) *ListFilesRequest
	GetKeyword() *string
	SetLastEditUser(v string) *ListFilesRequest
	GetLastEditUser() *string
	SetNeedAbsoluteFolderPath(v bool) *ListFilesRequest
	GetNeedAbsoluteFolderPath() *bool
	SetNeedContent(v bool) *ListFilesRequest
	GetNeedContent() *bool
	SetNodeId(v int64) *ListFilesRequest
	GetNodeId() *int64
	SetOwner(v string) *ListFilesRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListFilesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListFilesRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListFilesRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *ListFilesRequest
	GetProjectIdentifier() *string
	SetUseType(v string) *ListFilesRequest
	GetUseType() *string
}

type ListFilesRequest struct {
	// example:
	//
	// 1
	CommitStatus *int32 `json:"CommitStatus,omitempty" xml:"CommitStatus,omitempty"`
	// example:
	//
	// ods_create.sql
	ExactFileName *string `json:"ExactFileName,omitempty" xml:"ExactFileName,omitempty"`
	// example:
	//
	// Business_process/my_first_business_process/MaxCompute/ods_layer
	FileFolderPath *string `json:"FileFolderPath,omitempty" xml:"FileFolderPath,omitempty"`
	// example:
	//
	// 78237,816123
	FileIdIn *string `json:"FileIdIn,omitempty" xml:"FileIdIn,omitempty"`
	// example:
	//
	// 10,23
	FileTypes *string `json:"FileTypes,omitempty" xml:"FileTypes,omitempty"`
	// example:
	//
	// ods
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 11233***
	LastEditUser *string `json:"LastEditUser,omitempty" xml:"LastEditUser,omitempty"`
	// example:
	//
	// false
	NeedAbsoluteFolderPath *bool `json:"NeedAbsoluteFolderPath,omitempty" xml:"NeedAbsoluteFolderPath,omitempty"`
	// example:
	//
	// false
	NeedContent *bool `json:"NeedContent,omitempty" xml:"NeedContent,omitempty"`
	// example:
	//
	// 123541234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// 3726346****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
	// example:
	//
	// NORMAL
	UseType *string `json:"UseType,omitempty" xml:"UseType,omitempty"`
}

func (s ListFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFilesRequest) GoString() string {
	return s.String()
}

func (s *ListFilesRequest) GetCommitStatus() *int32 {
	return s.CommitStatus
}

func (s *ListFilesRequest) GetExactFileName() *string {
	return s.ExactFileName
}

func (s *ListFilesRequest) GetFileFolderPath() *string {
	return s.FileFolderPath
}

func (s *ListFilesRequest) GetFileIdIn() *string {
	return s.FileIdIn
}

func (s *ListFilesRequest) GetFileTypes() *string {
	return s.FileTypes
}

func (s *ListFilesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListFilesRequest) GetLastEditUser() *string {
	return s.LastEditUser
}

func (s *ListFilesRequest) GetNeedAbsoluteFolderPath() *bool {
	return s.NeedAbsoluteFolderPath
}

func (s *ListFilesRequest) GetNeedContent() *bool {
	return s.NeedContent
}

func (s *ListFilesRequest) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListFilesRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListFilesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFilesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFilesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListFilesRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *ListFilesRequest) GetUseType() *string {
	return s.UseType
}

func (s *ListFilesRequest) SetCommitStatus(v int32) *ListFilesRequest {
	s.CommitStatus = &v
	return s
}

func (s *ListFilesRequest) SetExactFileName(v string) *ListFilesRequest {
	s.ExactFileName = &v
	return s
}

func (s *ListFilesRequest) SetFileFolderPath(v string) *ListFilesRequest {
	s.FileFolderPath = &v
	return s
}

func (s *ListFilesRequest) SetFileIdIn(v string) *ListFilesRequest {
	s.FileIdIn = &v
	return s
}

func (s *ListFilesRequest) SetFileTypes(v string) *ListFilesRequest {
	s.FileTypes = &v
	return s
}

func (s *ListFilesRequest) SetKeyword(v string) *ListFilesRequest {
	s.Keyword = &v
	return s
}

func (s *ListFilesRequest) SetLastEditUser(v string) *ListFilesRequest {
	s.LastEditUser = &v
	return s
}

func (s *ListFilesRequest) SetNeedAbsoluteFolderPath(v bool) *ListFilesRequest {
	s.NeedAbsoluteFolderPath = &v
	return s
}

func (s *ListFilesRequest) SetNeedContent(v bool) *ListFilesRequest {
	s.NeedContent = &v
	return s
}

func (s *ListFilesRequest) SetNodeId(v int64) *ListFilesRequest {
	s.NodeId = &v
	return s
}

func (s *ListFilesRequest) SetOwner(v string) *ListFilesRequest {
	s.Owner = &v
	return s
}

func (s *ListFilesRequest) SetPageNumber(v int32) *ListFilesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFilesRequest) SetPageSize(v int32) *ListFilesRequest {
	s.PageSize = &v
	return s
}

func (s *ListFilesRequest) SetProjectId(v int64) *ListFilesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFilesRequest) SetProjectIdentifier(v string) *ListFilesRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *ListFilesRequest) SetUseType(v string) *ListFilesRequest {
	s.UseType = &v
	return s
}

func (s *ListFilesRequest) Validate() error {
	return dara.Validate(s)
}
