// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceFolderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFolder(v *GetDataServiceFolderResponseBodyFolder) *GetDataServiceFolderResponseBody
	GetFolder() *GetDataServiceFolderResponseBodyFolder
	SetRequestId(v string) *GetDataServiceFolderResponseBody
	GetRequestId() *string
}

type GetDataServiceFolderResponseBody struct {
	// The details of the folder.
	Folder *GetDataServiceFolderResponseBodyFolder `json:"Folder,omitempty" xml:"Folder,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataServiceFolderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceFolderResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataServiceFolderResponseBody) GetFolder() *GetDataServiceFolderResponseBodyFolder {
	return s.Folder
}

func (s *GetDataServiceFolderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataServiceFolderResponseBody) SetFolder(v *GetDataServiceFolderResponseBodyFolder) *GetDataServiceFolderResponseBody {
	s.Folder = v
	return s
}

func (s *GetDataServiceFolderResponseBody) SetRequestId(v string) *GetDataServiceFolderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataServiceFolderResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataServiceFolderResponseBodyFolder struct {
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
	// The parent folder ID. The ID of the root folder in a business process is 0, and the ID of a folder created by a user in a business process is greater than 0.
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

func (s GetDataServiceFolderResponseBodyFolder) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceFolderResponseBodyFolder) GoString() string {
	return s.String()
}

func (s *GetDataServiceFolderResponseBodyFolder) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *GetDataServiceFolderResponseBodyFolder) GetFolderId() *int64 {
	return s.FolderId
}

func (s *GetDataServiceFolderResponseBodyFolder) GetFolderName() *string {
	return s.FolderName
}

func (s *GetDataServiceFolderResponseBodyFolder) GetGroupId() *string {
	return s.GroupId
}

func (s *GetDataServiceFolderResponseBodyFolder) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetDataServiceFolderResponseBodyFolder) GetParentId() *int64 {
	return s.ParentId
}

func (s *GetDataServiceFolderResponseBodyFolder) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDataServiceFolderResponseBodyFolder) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetDataServiceFolderResponseBodyFolder) SetCreatedTime(v string) *GetDataServiceFolderResponseBodyFolder {
	s.CreatedTime = &v
	return s
}

func (s *GetDataServiceFolderResponseBodyFolder) SetFolderId(v int64) *GetDataServiceFolderResponseBodyFolder {
	s.FolderId = &v
	return s
}

func (s *GetDataServiceFolderResponseBodyFolder) SetFolderName(v string) *GetDataServiceFolderResponseBodyFolder {
	s.FolderName = &v
	return s
}

func (s *GetDataServiceFolderResponseBodyFolder) SetGroupId(v string) *GetDataServiceFolderResponseBodyFolder {
	s.GroupId = &v
	return s
}

func (s *GetDataServiceFolderResponseBodyFolder) SetModifiedTime(v string) *GetDataServiceFolderResponseBodyFolder {
	s.ModifiedTime = &v
	return s
}

func (s *GetDataServiceFolderResponseBodyFolder) SetParentId(v int64) *GetDataServiceFolderResponseBodyFolder {
	s.ParentId = &v
	return s
}

func (s *GetDataServiceFolderResponseBodyFolder) SetProjectId(v int64) *GetDataServiceFolderResponseBodyFolder {
	s.ProjectId = &v
	return s
}

func (s *GetDataServiceFolderResponseBodyFolder) SetTenantId(v int64) *GetDataServiceFolderResponseBodyFolder {
	s.TenantId = &v
	return s
}

func (s *GetDataServiceFolderResponseBodyFolder) Validate() error {
	return dara.Validate(s)
}
