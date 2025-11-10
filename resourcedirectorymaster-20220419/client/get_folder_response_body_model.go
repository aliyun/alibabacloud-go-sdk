// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFolderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFolder(v *GetFolderResponseBodyFolder) *GetFolderResponseBody
	GetFolder() *GetFolderResponseBodyFolder
	SetRequestId(v string) *GetFolderResponseBody
	GetRequestId() *string
}

type GetFolderResponseBody struct {
	// The information about the folder.
	Folder *GetFolderResponseBodyFolder `json:"Folder,omitempty" xml:"Folder,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// C2CBCA30-C8DD-423E-B4AD-4FB694C9180C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFolderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFolderResponseBody) GoString() string {
	return s.String()
}

func (s *GetFolderResponseBody) GetFolder() *GetFolderResponseBodyFolder {
	return s.Folder
}

func (s *GetFolderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFolderResponseBody) SetFolder(v *GetFolderResponseBodyFolder) *GetFolderResponseBody {
	s.Folder = v
	return s
}

func (s *GetFolderResponseBody) SetRequestId(v string) *GetFolderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFolderResponseBody) Validate() error {
	if s.Folder != nil {
		if err := s.Folder.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFolderResponseBodyFolder struct {
	// The time when the folder was created.
	//
	// example:
	//
	// 2021-06-15T06:39:08.521Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the folder.
	//
	// example:
	//
	// fd-Jyl5U7****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The name of the folder.
	//
	// example:
	//
	// Applications
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
	// The ID of the parent folder.
	//
	// example:
	//
	// r-Wm****
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	// The path of the folder in the resource directory.
	ResourceDirectoryPath *string `json:"ResourceDirectoryPath,omitempty" xml:"ResourceDirectoryPath,omitempty"`
}

func (s GetFolderResponseBodyFolder) String() string {
	return dara.Prettify(s)
}

func (s GetFolderResponseBodyFolder) GoString() string {
	return s.String()
}

func (s *GetFolderResponseBodyFolder) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetFolderResponseBodyFolder) GetFolderId() *string {
	return s.FolderId
}

func (s *GetFolderResponseBodyFolder) GetFolderName() *string {
	return s.FolderName
}

func (s *GetFolderResponseBodyFolder) GetParentFolderId() *string {
	return s.ParentFolderId
}

func (s *GetFolderResponseBodyFolder) GetResourceDirectoryPath() *string {
	return s.ResourceDirectoryPath
}

func (s *GetFolderResponseBodyFolder) SetCreateTime(v string) *GetFolderResponseBodyFolder {
	s.CreateTime = &v
	return s
}

func (s *GetFolderResponseBodyFolder) SetFolderId(v string) *GetFolderResponseBodyFolder {
	s.FolderId = &v
	return s
}

func (s *GetFolderResponseBodyFolder) SetFolderName(v string) *GetFolderResponseBodyFolder {
	s.FolderName = &v
	return s
}

func (s *GetFolderResponseBodyFolder) SetParentFolderId(v string) *GetFolderResponseBodyFolder {
	s.ParentFolderId = &v
	return s
}

func (s *GetFolderResponseBodyFolder) SetResourceDirectoryPath(v string) *GetFolderResponseBodyFolder {
	s.ResourceDirectoryPath = &v
	return s
}

func (s *GetFolderResponseBodyFolder) Validate() error {
	return dara.Validate(s)
}
