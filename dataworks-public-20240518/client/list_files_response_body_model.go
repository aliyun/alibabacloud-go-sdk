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
	// The response details.
	Data *ListFilesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID. Use this ID to troubleshoot issues.
	//
	// example:
	//
	// 0000-ABCD-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call succeeded. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFilesResponseBodyData struct {
	// The file details.
	Files []*ListFilesResponseBodyDataFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
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
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFilesResponseBodyDataFiles struct {
	// The path to the folder where the file is located.
	//
	// example:
	//
	// Business_process/my_first_business_process/MaxCompute/ods_layer
	AbsoluteFolderPath *string `json:"AbsoluteFolderPath,omitempty" xml:"AbsoluteFolderPath,omitempty"`
	// Specifies whether automatic parsing is enabled for the file. Valid values:
	//
	// 	- true: The file automatically parses code.
	//
	// 	- false: The file does not automatically parse code.
	//
	// This parameter corresponds to Analyze Code when you set Dependencies to Same Cycle in the scheduling configuration of a Data Studio task in the [DataWorks console](https://workbench.data.aliyun.com/console).
	//
	// example:
	//
	// true
	AutoParsing *bool `json:"AutoParsing,omitempty" xml:"AutoParsing,omitempty"`
	// The ID of the workflow to which the file belongs. This parameter is deprecated. Use the BusinessId parameter instead.
	//
	// example:
	//
	// 300000
	BizId *int64 `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The ID of the workflow to which the file belongs.
	//
	// example:
	//
	// 300000
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// The current commit status of the file. Valid values: 0 (the latest code is not committed) and 1 (the latest code is committed).
	//
	// example:
	//
	// 1
	CommitStatus *int32 `json:"CommitStatus,omitempty" xml:"CommitStatus,omitempty"`
	// The data source name used by the task.
	//
	// example:
	//
	// odps_source
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// This parameter is deprecated. You can call the [GetFile](https://help.aliyun.com/document_detail/173954.html) operation to query this information.
	//
	// example:
	//
	// SHOW TABLES;
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The timestamp (in milliseconds) when the file was created.
	//
	// example:
	//
	// 1593950832000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The Alibaba Cloud account ID of the file creator.
	//
	// example:
	//
	// 382762****
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The latest version of the file.
	//
	// example:
	//
	// 2
	CurrentVersion *int32 `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	// The description of the file.
	//
	// example:
	//
	// my test datastudio file
	FileDescription *string `json:"FileDescription,omitempty" xml:"FileDescription,omitempty"`
	// The ID of the folder where the file is located.
	//
	// example:
	//
	// 2735c2****
	FileFolderId *string `json:"FileFolderId,omitempty" xml:"FileFolderId,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 10000001
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The file name.
	//
	// example:
	//
	// ods_user_info_d
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file type. Different file types have different code. For more information, see [DataWorks node types](https://help.aliyun.com/document_detail/600169.html).
	//
	// example:
	//
	// 10
	FileType *int32 `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// If the current file is a MaxCompute resource file, this parameter specifies whether the resource file needs to be uploaded to MaxCompute.
	//
	// You only need to configure this parameter when the file is a MaxCompute resource file.
	//
	// example:
	//
	// false
	IsMaxCompute *bool `json:"IsMaxCompute,omitempty" xml:"IsMaxCompute,omitempty"`
	// The timestamp (in milliseconds) when the file was last modified.
	//
	// example:
	//
	// 1593950832000
	LastEditTime *int64 `json:"LastEditTime,omitempty" xml:"LastEditTime,omitempty"`
	// The Alibaba Cloud account ID of the user who last updated the file.
	//
	// example:
	//
	// 382762****
	LastEditUser *string `json:"LastEditUser,omitempty" xml:"LastEditUser,omitempty"`
	// The ID of the scheduling task generated in the scheduling system after the file is committed.
	//
	// example:
	//
	// 300001
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The Alibaba Cloud account ID of the file owner.
	//
	// example:
	//
	// 3872572****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// If the current file is an internal file of a combined node, this parameter specifies the ID of the corresponding combined node file.
	//
	// example:
	//
	// -1
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
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
