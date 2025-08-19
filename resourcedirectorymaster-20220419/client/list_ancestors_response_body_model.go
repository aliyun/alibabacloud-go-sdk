// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAncestorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFolders(v *ListAncestorsResponseBodyFolders) *ListAncestorsResponseBody
	GetFolders() *ListAncestorsResponseBodyFolders
	SetRequestId(v string) *ListAncestorsResponseBody
	GetRequestId() *string
}

type ListAncestorsResponseBody struct {
	// The information of the folders.
	Folders *ListAncestorsResponseBodyFolders `json:"Folders,omitempty" xml:"Folders,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 83AFBEB6-DC03-406E-9686-867461FF6698
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAncestorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAncestorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAncestorsResponseBody) GetFolders() *ListAncestorsResponseBodyFolders {
	return s.Folders
}

func (s *ListAncestorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAncestorsResponseBody) SetFolders(v *ListAncestorsResponseBodyFolders) *ListAncestorsResponseBody {
	s.Folders = v
	return s
}

func (s *ListAncestorsResponseBody) SetRequestId(v string) *ListAncestorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAncestorsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAncestorsResponseBodyFolders struct {
	Folder []*ListAncestorsResponseBodyFoldersFolder `json:"Folder,omitempty" xml:"Folder,omitempty" type:"Repeated"`
}

func (s ListAncestorsResponseBodyFolders) String() string {
	return dara.Prettify(s)
}

func (s ListAncestorsResponseBodyFolders) GoString() string {
	return s.String()
}

func (s *ListAncestorsResponseBodyFolders) GetFolder() []*ListAncestorsResponseBodyFoldersFolder {
	return s.Folder
}

func (s *ListAncestorsResponseBodyFolders) SetFolder(v []*ListAncestorsResponseBodyFoldersFolder) *ListAncestorsResponseBodyFolders {
	s.Folder = v
	return s
}

func (s *ListAncestorsResponseBodyFolders) Validate() error {
	return dara.Validate(s)
}

type ListAncestorsResponseBodyFoldersFolder struct {
	// The time when the folder was created.
	//
	// example:
	//
	// 2019-01-18T10:03:35.217Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the folder.
	//
	// example:
	//
	// r-b1****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The name of the folder.
	//
	// example:
	//
	// root
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
}

func (s ListAncestorsResponseBodyFoldersFolder) String() string {
	return dara.Prettify(s)
}

func (s ListAncestorsResponseBodyFoldersFolder) GoString() string {
	return s.String()
}

func (s *ListAncestorsResponseBodyFoldersFolder) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAncestorsResponseBodyFoldersFolder) GetFolderId() *string {
	return s.FolderId
}

func (s *ListAncestorsResponseBodyFoldersFolder) GetFolderName() *string {
	return s.FolderName
}

func (s *ListAncestorsResponseBodyFoldersFolder) SetCreateTime(v string) *ListAncestorsResponseBodyFoldersFolder {
	s.CreateTime = &v
	return s
}

func (s *ListAncestorsResponseBodyFoldersFolder) SetFolderId(v string) *ListAncestorsResponseBodyFoldersFolder {
	s.FolderId = &v
	return s
}

func (s *ListAncestorsResponseBodyFoldersFolder) SetFolderName(v string) *ListAncestorsResponseBodyFoldersFolder {
	s.FolderName = &v
	return s
}

func (s *ListAncestorsResponseBodyFoldersFolder) Validate() error {
	return dara.Validate(s)
}
