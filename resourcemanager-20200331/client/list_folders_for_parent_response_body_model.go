// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFoldersForParentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFolders(v *ListFoldersForParentResponseBodyFolders) *ListFoldersForParentResponseBody
	GetFolders() *ListFoldersForParentResponseBodyFolders
	SetPageNumber(v int32) *ListFoldersForParentResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListFoldersForParentResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListFoldersForParentResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListFoldersForParentResponseBody
	GetTotalCount() *int32
}

type ListFoldersForParentResponseBody struct {
	// The information of the folders.
	Folders *ListFoldersForParentResponseBodyFolders `json:"Folders,omitempty" xml:"Folders,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7B8A4E7D-6CFF-471D-84DF-195A7A241ECB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFoldersForParentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFoldersForParentResponseBody) GoString() string {
	return s.String()
}

func (s *ListFoldersForParentResponseBody) GetFolders() *ListFoldersForParentResponseBodyFolders {
	return s.Folders
}

func (s *ListFoldersForParentResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFoldersForParentResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFoldersForParentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFoldersForParentResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListFoldersForParentResponseBody) SetFolders(v *ListFoldersForParentResponseBodyFolders) *ListFoldersForParentResponseBody {
	s.Folders = v
	return s
}

func (s *ListFoldersForParentResponseBody) SetPageNumber(v int32) *ListFoldersForParentResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFoldersForParentResponseBody) SetPageSize(v int32) *ListFoldersForParentResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFoldersForParentResponseBody) SetRequestId(v string) *ListFoldersForParentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFoldersForParentResponseBody) SetTotalCount(v int32) *ListFoldersForParentResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFoldersForParentResponseBody) Validate() error {
	if s.Folders != nil {
		if err := s.Folders.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFoldersForParentResponseBodyFolders struct {
	Folder []*ListFoldersForParentResponseBodyFoldersFolder `json:"Folder,omitempty" xml:"Folder,omitempty" type:"Repeated"`
}

func (s ListFoldersForParentResponseBodyFolders) String() string {
	return dara.Prettify(s)
}

func (s ListFoldersForParentResponseBodyFolders) GoString() string {
	return s.String()
}

func (s *ListFoldersForParentResponseBodyFolders) GetFolder() []*ListFoldersForParentResponseBodyFoldersFolder {
	return s.Folder
}

func (s *ListFoldersForParentResponseBodyFolders) SetFolder(v []*ListFoldersForParentResponseBodyFoldersFolder) *ListFoldersForParentResponseBodyFolders {
	s.Folder = v
	return s
}

func (s *ListFoldersForParentResponseBodyFolders) Validate() error {
	if s.Folder != nil {
		for _, item := range s.Folder {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFoldersForParentResponseBodyFoldersFolder struct {
	// The time when the folder was created.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the folder.
	//
	// example:
	//
	// rd-evic31****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The name of the folder.
	//
	// example:
	//
	// project-1
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
}

func (s ListFoldersForParentResponseBodyFoldersFolder) String() string {
	return dara.Prettify(s)
}

func (s ListFoldersForParentResponseBodyFoldersFolder) GoString() string {
	return s.String()
}

func (s *ListFoldersForParentResponseBodyFoldersFolder) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListFoldersForParentResponseBodyFoldersFolder) GetFolderId() *string {
	return s.FolderId
}

func (s *ListFoldersForParentResponseBodyFoldersFolder) GetFolderName() *string {
	return s.FolderName
}

func (s *ListFoldersForParentResponseBodyFoldersFolder) SetCreateTime(v string) *ListFoldersForParentResponseBodyFoldersFolder {
	s.CreateTime = &v
	return s
}

func (s *ListFoldersForParentResponseBodyFoldersFolder) SetFolderId(v string) *ListFoldersForParentResponseBodyFoldersFolder {
	s.FolderId = &v
	return s
}

func (s *ListFoldersForParentResponseBodyFoldersFolder) SetFolderName(v string) *ListFoldersForParentResponseBodyFoldersFolder {
	s.FolderName = &v
	return s
}

func (s *ListFoldersForParentResponseBodyFoldersFolder) Validate() error {
	return dara.Validate(s)
}
