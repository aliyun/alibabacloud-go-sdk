// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListFilesResponseBodyData) *ListFilesResponseBody
	GetData() *ListFilesResponseBodyData
	SetErrorCode(v string) *ListFilesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListFilesResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListFilesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListFilesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListFilesResponseBody
	GetSuccess() *bool
}

type ListFilesResponseBody struct {
	Data *ListFilesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 0000-ABCD-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFilesResponseBody) GetData() *ListFilesResponseBodyData {
	return s.Data
}

func (s *ListFilesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListFilesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListFilesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFilesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListFilesResponseBody) SetData(v *ListFilesResponseBodyData) *ListFilesResponseBody {
	s.Data = v
	return s
}

func (s *ListFilesResponseBody) SetErrorCode(v string) *ListFilesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListFilesResponseBody) SetErrorMessage(v string) *ListFilesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListFilesResponseBody) SetHttpStatusCode(v int32) *ListFilesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListFilesResponseBody) SetRequestId(v string) *ListFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFilesResponseBody) SetSuccess(v bool) *ListFilesResponseBody {
	s.Success = &v
	return s
}

func (s *ListFilesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListFilesResponseBodyData struct {
	Files []*ListFilesResponseBodyDataFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
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
	// 13
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFilesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFilesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFilesResponseBodyData) GetFiles() []*ListFilesResponseBodyDataFiles {
	return s.Files
}

func (s *ListFilesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFilesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFilesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListFilesResponseBodyData) SetFiles(v []*ListFilesResponseBodyDataFiles) *ListFilesResponseBodyData {
	s.Files = v
	return s
}

func (s *ListFilesResponseBodyData) SetPageNumber(v int32) *ListFilesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListFilesResponseBodyData) SetPageSize(v int32) *ListFilesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListFilesResponseBodyData) SetTotalCount(v int32) *ListFilesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListFilesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListFilesResponseBodyDataFiles struct {
	// example:
	//
	// Business_process/my_first_business_process/MaxCompute/ods_layer
	AbsoluteFolderPath *string `json:"AbsoluteFolderPath,omitempty" xml:"AbsoluteFolderPath,omitempty"`
	// example:
	//
	// true
	AutoParsing *bool `json:"AutoParsing,omitempty" xml:"AutoParsing,omitempty"`
	// example:
	//
	// 300000
	BizId *int64 `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// 300000
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// example:
	//
	// 1
	CommitStatus *int32 `json:"CommitStatus,omitempty" xml:"CommitStatus,omitempty"`
	// example:
	//
	// odps_source
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// example:
	//
	// SHOW TABLES;
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1593950832000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 382762****
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// 2
	CurrentVersion *int32 `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	// example:
	//
	// my test datastudio file
	FileDescription *string `json:"FileDescription,omitempty" xml:"FileDescription,omitempty"`
	// example:
	//
	// 2735c2****
	FileFolderId *string `json:"FileFolderId,omitempty" xml:"FileFolderId,omitempty"`
	// example:
	//
	// 10000001
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// ods_user_info_d
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 10
	FileType *int32 `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// example:
	//
	// false
	IsMaxCompute *bool `json:"IsMaxCompute,omitempty" xml:"IsMaxCompute,omitempty"`
	// example:
	//
	// 1593950832000
	LastEditTime *int64 `json:"LastEditTime,omitempty" xml:"LastEditTime,omitempty"`
	// example:
	//
	// 382762****
	LastEditUser *string `json:"LastEditUser,omitempty" xml:"LastEditUser,omitempty"`
	// example:
	//
	// 300001
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// 3872572****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// -1
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// NORMAL
	UseType *string `json:"UseType,omitempty" xml:"UseType,omitempty"`
}

func (s ListFilesResponseBodyDataFiles) String() string {
	return dara.Prettify(s)
}

func (s ListFilesResponseBodyDataFiles) GoString() string {
	return s.String()
}

func (s *ListFilesResponseBodyDataFiles) GetAbsoluteFolderPath() *string {
	return s.AbsoluteFolderPath
}

func (s *ListFilesResponseBodyDataFiles) GetAutoParsing() *bool {
	return s.AutoParsing
}

func (s *ListFilesResponseBodyDataFiles) GetBizId() *int64 {
	return s.BizId
}

func (s *ListFilesResponseBodyDataFiles) GetBusinessId() *int64 {
	return s.BusinessId
}

func (s *ListFilesResponseBodyDataFiles) GetCommitStatus() *int32 {
	return s.CommitStatus
}

func (s *ListFilesResponseBodyDataFiles) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *ListFilesResponseBodyDataFiles) GetContent() *string {
	return s.Content
}

func (s *ListFilesResponseBodyDataFiles) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListFilesResponseBodyDataFiles) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListFilesResponseBodyDataFiles) GetCurrentVersion() *int32 {
	return s.CurrentVersion
}

func (s *ListFilesResponseBodyDataFiles) GetFileDescription() *string {
	return s.FileDescription
}

func (s *ListFilesResponseBodyDataFiles) GetFileFolderId() *string {
	return s.FileFolderId
}

func (s *ListFilesResponseBodyDataFiles) GetFileId() *int64 {
	return s.FileId
}

func (s *ListFilesResponseBodyDataFiles) GetFileName() *string {
	return s.FileName
}

func (s *ListFilesResponseBodyDataFiles) GetFileType() *int32 {
	return s.FileType
}

func (s *ListFilesResponseBodyDataFiles) GetIsMaxCompute() *bool {
	return s.IsMaxCompute
}

func (s *ListFilesResponseBodyDataFiles) GetLastEditTime() *int64 {
	return s.LastEditTime
}

func (s *ListFilesResponseBodyDataFiles) GetLastEditUser() *string {
	return s.LastEditUser
}

func (s *ListFilesResponseBodyDataFiles) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListFilesResponseBodyDataFiles) GetOwner() *string {
	return s.Owner
}

func (s *ListFilesResponseBodyDataFiles) GetParentId() *int64 {
	return s.ParentId
}

func (s *ListFilesResponseBodyDataFiles) GetUseType() *string {
	return s.UseType
}

func (s *ListFilesResponseBodyDataFiles) SetAbsoluteFolderPath(v string) *ListFilesResponseBodyDataFiles {
	s.AbsoluteFolderPath = &v
	return s
}

func (s *ListFilesResponseBodyDataFiles) SetAutoParsing(v bool) *ListFilesResponseBodyDataFiles {
	s.AutoParsing = &v
	return s
}

func (s *ListFilesResponseBodyDataFiles) SetBizId(v int64) *ListFilesResponseBodyDataFiles {
	s.BizId = &v
	return s
}

func (s *ListFilesResponseBodyDataFiles) SetBusinessId(v int64) *ListFilesResponseBodyDataFiles {
	s.BusinessId = &v
	return s
}

func (s *ListFilesResponseBodyDataFiles) SetCommitStatus(v int32) *ListFilesResponseBodyDataFiles {
	s.CommitStatus = &v
	return s
}

func (s *ListFilesResponseBodyDataFiles) SetConnectionName(v string) *ListFilesResponseBodyDataFiles {
	s.ConnectionName = &v
	return s
}

func (s *ListFilesResponseBodyDataFiles) SetContent(v string) *ListFilesResponseBodyDataFiles {
	s.Content = &v
	return s
}

func (s *ListFilesResponseBodyDataFiles) SetCreateTime(v int64) *ListFilesResponseBodyDataFiles {
	s.CreateTime = &v
	return s
}

func (s *ListFilesResponseBodyDataFiles) SetCreateUser(v string) *ListFilesResponseBodyDataFiles {
	s.CreateUser = &v
	return s
}

func (s *ListFilesResponseBodyDataFiles) SetCurrentVersion(v int32) *ListFilesResponseBodyDataFiles {
	s.CurrentVersion = &v
	return s
}

func (s *ListFilesResponseBodyDataFiles) SetFileDescription(v string) *ListFilesResponseBodyDataFiles {
	s.FileDescription = &v
	return s
}

func (s *ListFilesResponseBodyDataFiles) SetFileFolderId(v string) *ListFilesResponseBodyDataFiles {
	s.FileFolderId = &v
	return s
}

func (s *ListFilesResponseBodyDataFiles) SetFileId(v int64) *ListFilesResponseBodyDataFiles {
	s.FileId = &v
	return s
}

func (s *ListFilesResponseBodyDataFiles) SetFileName(v string) *ListFilesResponseBodyDataFiles {
	s.FileName = &v
	return s
}

func (s *ListFilesResponseBodyDataFiles) SetFileType(v int32) *ListFilesResponseBodyDataFiles {
	s.FileType = &v
	return s
}

func (s *ListFilesResponseBodyDataFiles) SetIsMaxCompute(v bool) *ListFilesResponseBodyDataFiles {
	s.IsMaxCompute = &v
	return s
}

func (s *ListFilesResponseBodyDataFiles) SetLastEditTime(v int64) *ListFilesResponseBodyDataFiles {
	s.LastEditTime = &v
	return s
}

func (s *ListFilesResponseBodyDataFiles) SetLastEditUser(v string) *ListFilesResponseBodyDataFiles {
	s.LastEditUser = &v
	return s
}

func (s *ListFilesResponseBodyDataFiles) SetNodeId(v int64) *ListFilesResponseBodyDataFiles {
	s.NodeId = &v
	return s
}

func (s *ListFilesResponseBodyDataFiles) SetOwner(v string) *ListFilesResponseBodyDataFiles {
	s.Owner = &v
	return s
}

func (s *ListFilesResponseBodyDataFiles) SetParentId(v int64) *ListFilesResponseBodyDataFiles {
	s.ParentId = &v
	return s
}

func (s *ListFilesResponseBodyDataFiles) SetUseType(v string) *ListFilesResponseBodyDataFiles {
	s.UseType = &v
	return s
}

func (s *ListFilesResponseBodyDataFiles) Validate() error {
	return dara.Validate(s)
}
