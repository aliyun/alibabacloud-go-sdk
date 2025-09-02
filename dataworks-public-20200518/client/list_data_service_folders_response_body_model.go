// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceFoldersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFolderPagingResult(v *ListDataServiceFoldersResponseBodyFolderPagingResult) *ListDataServiceFoldersResponseBody
	GetFolderPagingResult() *ListDataServiceFoldersResponseBodyFolderPagingResult
	SetRequestId(v string) *ListDataServiceFoldersResponseBody
	GetRequestId() *string
}

type ListDataServiceFoldersResponseBody struct {
	// The paging result for the folders.
	FolderPagingResult *ListDataServiceFoldersResponseBodyFolderPagingResult `json:"FolderPagingResult,omitempty" xml:"FolderPagingResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataServiceFoldersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceFoldersResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataServiceFoldersResponseBody) GetFolderPagingResult() *ListDataServiceFoldersResponseBodyFolderPagingResult {
	return s.FolderPagingResult
}

func (s *ListDataServiceFoldersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataServiceFoldersResponseBody) SetFolderPagingResult(v *ListDataServiceFoldersResponseBodyFolderPagingResult) *ListDataServiceFoldersResponseBody {
	s.FolderPagingResult = v
	return s
}

func (s *ListDataServiceFoldersResponseBody) SetRequestId(v string) *ListDataServiceFoldersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataServiceFoldersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceFoldersResponseBodyFolderPagingResult struct {
	// The folders.
	Folders []*ListDataServiceFoldersResponseBodyFolderPagingResultFolders `json:"Folders,omitempty" xml:"Folders,omitempty" type:"Repeated"`
	// The page number. The value of this parameter is the same as that of the PageNumber parameter in the request.
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
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataServiceFoldersResponseBodyFolderPagingResult) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceFoldersResponseBodyFolderPagingResult) GoString() string {
	return s.String()
}

func (s *ListDataServiceFoldersResponseBodyFolderPagingResult) GetFolders() []*ListDataServiceFoldersResponseBodyFolderPagingResultFolders {
	return s.Folders
}

func (s *ListDataServiceFoldersResponseBodyFolderPagingResult) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataServiceFoldersResponseBodyFolderPagingResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataServiceFoldersResponseBodyFolderPagingResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataServiceFoldersResponseBodyFolderPagingResult) SetFolders(v []*ListDataServiceFoldersResponseBodyFolderPagingResultFolders) *ListDataServiceFoldersResponseBodyFolderPagingResult {
	s.Folders = v
	return s
}

func (s *ListDataServiceFoldersResponseBodyFolderPagingResult) SetPageNumber(v int32) *ListDataServiceFoldersResponseBodyFolderPagingResult {
	s.PageNumber = &v
	return s
}

func (s *ListDataServiceFoldersResponseBodyFolderPagingResult) SetPageSize(v int32) *ListDataServiceFoldersResponseBodyFolderPagingResult {
	s.PageSize = &v
	return s
}

func (s *ListDataServiceFoldersResponseBodyFolderPagingResult) SetTotalCount(v int32) *ListDataServiceFoldersResponseBodyFolderPagingResult {
	s.TotalCount = &v
	return s
}

func (s *ListDataServiceFoldersResponseBodyFolderPagingResult) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceFoldersResponseBodyFolderPagingResultFolders struct {
	// The time when the folder was created.
	//
	// example:
	//
	// 2020-09-24T18:37:51+0800
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// 11
	FolderId *int64 `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The name of the folder.
	//
	// example:
	//
	// test1
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
	// The ID of the business process to which the folder belongs.
	//
	// example:
	//
	// ds_1234
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The time when the folder was last modified.
	//
	// example:
	//
	// 2020-09-24T18:37:51+0800
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the parent folder. The ID of the root folder in a business process is 0, and the ID of a folder created by a user in a business process is greater than 0.
	//
	// example:
	//
	// 0
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 10001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10002
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListDataServiceFoldersResponseBodyFolderPagingResultFolders) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceFoldersResponseBodyFolderPagingResultFolders) GoString() string {
	return s.String()
}

func (s *ListDataServiceFoldersResponseBodyFolderPagingResultFolders) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListDataServiceFoldersResponseBodyFolderPagingResultFolders) GetFolderId() *int64 {
	return s.FolderId
}

func (s *ListDataServiceFoldersResponseBodyFolderPagingResultFolders) GetFolderName() *string {
	return s.FolderName
}

func (s *ListDataServiceFoldersResponseBodyFolderPagingResultFolders) GetGroupId() *string {
	return s.GroupId
}

func (s *ListDataServiceFoldersResponseBodyFolderPagingResultFolders) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListDataServiceFoldersResponseBodyFolderPagingResultFolders) GetParentId() *int64 {
	return s.ParentId
}

func (s *ListDataServiceFoldersResponseBodyFolderPagingResultFolders) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataServiceFoldersResponseBodyFolderPagingResultFolders) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListDataServiceFoldersResponseBodyFolderPagingResultFolders) SetCreatedTime(v string) *ListDataServiceFoldersResponseBodyFolderPagingResultFolders {
	s.CreatedTime = &v
	return s
}

func (s *ListDataServiceFoldersResponseBodyFolderPagingResultFolders) SetFolderId(v int64) *ListDataServiceFoldersResponseBodyFolderPagingResultFolders {
	s.FolderId = &v
	return s
}

func (s *ListDataServiceFoldersResponseBodyFolderPagingResultFolders) SetFolderName(v string) *ListDataServiceFoldersResponseBodyFolderPagingResultFolders {
	s.FolderName = &v
	return s
}

func (s *ListDataServiceFoldersResponseBodyFolderPagingResultFolders) SetGroupId(v string) *ListDataServiceFoldersResponseBodyFolderPagingResultFolders {
	s.GroupId = &v
	return s
}

func (s *ListDataServiceFoldersResponseBodyFolderPagingResultFolders) SetModifiedTime(v string) *ListDataServiceFoldersResponseBodyFolderPagingResultFolders {
	s.ModifiedTime = &v
	return s
}

func (s *ListDataServiceFoldersResponseBodyFolderPagingResultFolders) SetParentId(v int64) *ListDataServiceFoldersResponseBodyFolderPagingResultFolders {
	s.ParentId = &v
	return s
}

func (s *ListDataServiceFoldersResponseBodyFolderPagingResultFolders) SetProjectId(v int64) *ListDataServiceFoldersResponseBodyFolderPagingResultFolders {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceFoldersResponseBodyFolderPagingResultFolders) SetTenantId(v int64) *ListDataServiceFoldersResponseBodyFolderPagingResultFolders {
	s.TenantId = &v
	return s
}

func (s *ListDataServiceFoldersResponseBodyFolderPagingResultFolders) Validate() error {
	return dara.Validate(s)
}
