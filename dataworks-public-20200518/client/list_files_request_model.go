// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFilesRequest interface {
	dara.Model
	String() string
	GoString() string
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
	// The exact matching file name. The file name of the query result is exactly the same as this parameter.
	//
	// example:
	//
	// ods_create.sql
	ExactFileName *string `json:"ExactFileName,omitempty" xml:"ExactFileName,omitempty"`
	// The path of the folder to which files belong.
	//
	// example:
	//
	// Business_process/my_first_business_process/MaxCompute/ods_layer
	FileFolderPath *string `json:"FileFolderPath,omitempty" xml:"FileFolderPath,omitempty"`
	// The file ID list. The File ID set of the query result can only be a subset of the list. You can specify up to 50 fileids at a time.
	//
	// example:
	//
	// 78237,816123
	FileIdIn *string `json:"FileIdIn,omitempty" xml:"FileIdIn,omitempty"`
	// The types of the code in the files.
	//
	// Valid values: 6 (Shell), 10 (ODPS SQL), 11 (ODPS MR), 23 (Data Integration), 24 (ODPS Script), 97 (PAI), 98 (node group), 99 (zero load), 221 (PyODPS 2), 225 (ODPS Spark), 227 (EMR Hive), 228 (EMR Spark), 229 (EMR Spark SQL), 230 (EMR MR), 239 (OSS object inspection), 257 (EMR Shell), 258 (EMR Spark Shell), 259 (EMR Presto), 260 (EMR Impala), 900 (real-time synchronization), 1002 (PAI inner node), 1089 (cross-tenant collaboration), 1091 (Hologres development), 1093 (Hologres SQL), 1100 (assignment), 1106 (for-each), and 1221 (PyODPS 3).
	//
	// example:
	//
	// 10,23
	FileTypes *string `json:"FileTypes,omitempty" xml:"FileTypes,omitempty"`
	// The keyword in the file names. The keyword is used to perform a fuzzy match. You can specify a keyword to query all files whose names contain the keyword.
	//
	// example:
	//
	// ods
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The ID of the Alibaba Cloud account that is used to last modify the file.
	//
	// example:
	//
	// 38748246285727
	LastEditUser *string `json:"LastEditUser,omitempty" xml:"LastEditUser,omitempty"`
	// Whether the query result contains the path of the folder where the file is located.
	//
	// example:
	//
	// false
	NeedAbsoluteFolderPath *bool `json:"NeedAbsoluteFolderPath,omitempty" xml:"NeedAbsoluteFolderPath,omitempty"`
	// Whether the query results contain file content (for files with more content, there may be a long network transmission delay).
	//
	// example:
	//
	// false
	NeedContent *bool `json:"NeedContent,omitempty" xml:"NeedContent,omitempty"`
	// The ID of the node that is scheduled. You can call the [ListNodes](https://help.aliyun.com/document_detail/173979.html) operation to query the ID of the node.
	//
	// example:
	//
	// 123541234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The owner of the files.
	//
	// example:
	//
	// 3726346****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number.
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
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace Management page to obtain the workspace ID.
	//
	// You must configure either the ProjectId or ProjectIdentifier parameter to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace Management page to obtain the workspace name.
	//
	// You must configure either the ProjectId or ProjectIdentifier parameter to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
	// The module to which the files belong. Valid values:
	//
	// 	- NORMAL: The files are used for DataStudio.
	//
	// 	- MANUAL: The files are used for manually triggered nodes.
	//
	// 	- MANUAL_BIZ: The files are used for manually triggered workflows.
	//
	// 	- SKIP: The files are used for dry-run nodes in DataStudio.
	//
	// 	- ADHOCQUERY: The files are used for ad hoc queries.
	//
	// 	- COMPONENT: The files are used for snippets.
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
