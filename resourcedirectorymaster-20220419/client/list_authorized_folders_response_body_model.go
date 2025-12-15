// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedFoldersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFolders(v *ListAuthorizedFoldersResponseBodyFolders) *ListAuthorizedFoldersResponseBody
	GetFolders() *ListAuthorizedFoldersResponseBodyFolders
	SetPageNumber(v int32) *ListAuthorizedFoldersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAuthorizedFoldersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAuthorizedFoldersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAuthorizedFoldersResponseBody
	GetTotalCount() *int32
}

type ListAuthorizedFoldersResponseBody struct {
	// The folders.
	Folders *ListAuthorizedFoldersResponseBodyFolders `json:"Folders,omitempty" xml:"Folders,omitempty" type:"Struct"`
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
	// The request ID.
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

func (s ListAuthorizedFoldersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedFoldersResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthorizedFoldersResponseBody) GetFolders() *ListAuthorizedFoldersResponseBodyFolders {
	return s.Folders
}

func (s *ListAuthorizedFoldersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAuthorizedFoldersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAuthorizedFoldersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuthorizedFoldersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAuthorizedFoldersResponseBody) SetFolders(v *ListAuthorizedFoldersResponseBodyFolders) *ListAuthorizedFoldersResponseBody {
	s.Folders = v
	return s
}

func (s *ListAuthorizedFoldersResponseBody) SetPageNumber(v int32) *ListAuthorizedFoldersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAuthorizedFoldersResponseBody) SetPageSize(v int32) *ListAuthorizedFoldersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAuthorizedFoldersResponseBody) SetRequestId(v string) *ListAuthorizedFoldersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthorizedFoldersResponseBody) SetTotalCount(v int32) *ListAuthorizedFoldersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAuthorizedFoldersResponseBody) Validate() error {
	if s.Folders != nil {
		if err := s.Folders.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAuthorizedFoldersResponseBodyFolders struct {
	Folder []*ListAuthorizedFoldersResponseBodyFoldersFolder `json:"Folder,omitempty" xml:"Folder,omitempty" type:"Repeated"`
}

func (s ListAuthorizedFoldersResponseBodyFolders) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedFoldersResponseBodyFolders) GoString() string {
	return s.String()
}

func (s *ListAuthorizedFoldersResponseBodyFolders) GetFolder() []*ListAuthorizedFoldersResponseBodyFoldersFolder {
	return s.Folder
}

func (s *ListAuthorizedFoldersResponseBodyFolders) SetFolder(v []*ListAuthorizedFoldersResponseBodyFoldersFolder) *ListAuthorizedFoldersResponseBodyFolders {
	s.Folder = v
	return s
}

func (s *ListAuthorizedFoldersResponseBodyFolders) Validate() error {
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

type ListAuthorizedFoldersResponseBodyFoldersFolder struct {
	// The ID of the folder.
	//
	// example:
	//
	// fd-bVaRIG****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The name of the folder.
	//
	// example:
	//
	// project-1
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
	// The RDPath of the folder.
	ResourceDirectoryPath *string `json:"ResourceDirectoryPath,omitempty" xml:"ResourceDirectoryPath,omitempty"`
}

func (s ListAuthorizedFoldersResponseBodyFoldersFolder) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedFoldersResponseBodyFoldersFolder) GoString() string {
	return s.String()
}

func (s *ListAuthorizedFoldersResponseBodyFoldersFolder) GetFolderId() *string {
	return s.FolderId
}

func (s *ListAuthorizedFoldersResponseBodyFoldersFolder) GetFolderName() *string {
	return s.FolderName
}

func (s *ListAuthorizedFoldersResponseBodyFoldersFolder) GetResourceDirectoryPath() *string {
	return s.ResourceDirectoryPath
}

func (s *ListAuthorizedFoldersResponseBodyFoldersFolder) SetFolderId(v string) *ListAuthorizedFoldersResponseBodyFoldersFolder {
	s.FolderId = &v
	return s
}

func (s *ListAuthorizedFoldersResponseBodyFoldersFolder) SetFolderName(v string) *ListAuthorizedFoldersResponseBodyFoldersFolder {
	s.FolderName = &v
	return s
}

func (s *ListAuthorizedFoldersResponseBodyFoldersFolder) SetResourceDirectoryPath(v string) *ListAuthorizedFoldersResponseBodyFoldersFolder {
	s.ResourceDirectoryPath = &v
	return s
}

func (s *ListAuthorizedFoldersResponseBodyFoldersFolder) Validate() error {
	return dara.Validate(s)
}
