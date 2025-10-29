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
	// The current commit status of the file. Valid values: 0 (the latest code is not committed) and 1 (the latest code is committed).
	//
	// example:
	//
	// 1
	CommitStatus *int32 `json:"CommitStatus,omitempty" xml:"CommitStatus,omitempty"`
	// The exact file name. The file name in the query result must exactly match this parameter.
	//
	// example:
	//
	// ods_create.sql
	ExactFileName *string `json:"ExactFileName,omitempty" xml:"ExactFileName,omitempty"`
	// The path to the folder where the file is located.
	//
	// example:
	//
	// Business_process/my_first_business_process/MaxCompute/ods_layer
	FileFolderPath *string `json:"FileFolderPath,omitempty" xml:"FileFolderPath,omitempty"`
	// The list of file IDs. The file IDs in the query result must be a subset of this list. You can specify up to 50 file IDs at a time.
	//
	// example:
	//
	// 78237,816123
	FileIdIn *string `json:"FileIdIn,omitempty" xml:"FileIdIn,omitempty"`
	// The code type of the file.
	//
	// The code type of the file. Common code types and their corresponding file types include: 6 (Shell), 10 (ODPS SQL), 11 (ODPS MR), 23 (Data Integration), 24 (ODPS Script), 97 (PAI), 98 (Combined node), 99 (Virtual node), 221 (PyODPS 2), 225 (ODPS Spark), 227 (EMR Hive), 228 (EMR Spark), 229 (EMR Spark SQL), 230 (EMR MR), 239 (OSS object inspection), 257 (EMR Shell), 258 (EMR Spark Shell), 259 (EMR Presto), 260 (EMR Impala), 900 (Real-time sync), 1002 (PAI internal node), 1089 (Cross-tenant node), 1091 (Hologres development), 1093 (Hologres SQL), 1100 (Assignment node), 1106 (ForEach node), 1221 (PyODPS 3).
	//
	// example:
	//
	// 10,23
	FileTypes *string `json:"FileTypes,omitempty" xml:"FileTypes,omitempty"`
	// The keyword for the file name. Fuzzy match is supported. You can enter a keyword to query all files that contain the keyword.
	//
	// example:
	//
	// ods
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The Alibaba Cloud account ID of the user who last updated the file.
	//
	// example:
	//
	// 11233***
	LastEditUser *string `json:"LastEditUser,omitempty" xml:"LastEditUser,omitempty"`
	// Specifies whether the query result includes the path to the folder where the file is located.
	//
	// example:
	//
	// false
	NeedAbsoluteFolderPath *bool `json:"NeedAbsoluteFolderPath,omitempty" xml:"NeedAbsoluteFolderPath,omitempty"`
	// Specifies whether the query result includes the file content. For files with large content, network transmission delays may occur.
	//
	// example:
	//
	// false
	NeedContent *bool `json:"NeedContent,omitempty" xml:"NeedContent,omitempty"`
	// The ID of the scheduling node. You can call the [ListNodes](https://help.aliyun.com/document_detail/173979.html) operation to obtain the node ID.
	//
	// example:
	//
	// 123541234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The ID of the file owner.
	//
	// example:
	//
	// 3726346****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number for pagination.
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
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the ID.
	//
	// You must configure either this parameter or the ProjectIdentifier parameter to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The DataWorks workspace name. To obtain the workspace name, log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and navigate to the workspace configuration page.
	//
	// You must specify either this parameter or ProjectId to identify the target DataWorks workspace for this API call.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
	// The functional module to which the file belongs. Valid values:
	//
	// 	- NORMAL: Data Studio
	//
	// 	- MANUAL: Manually triggered node
	//
	// 	- MANUAL_BIZ: Manually triggered workflow
	//
	// 	- SKIP: Dry-run scheduling in Data Studio
	//
	// 	- ADHOCQUERY: Ad hoc query
	//
	// 	- COMPONENT: Component management
	//
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
