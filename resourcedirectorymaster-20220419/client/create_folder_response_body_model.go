// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFolderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFolder(v *CreateFolderResponseBodyFolder) *CreateFolderResponseBody
	GetFolder() *CreateFolderResponseBodyFolder
	SetRequestId(v string) *CreateFolderResponseBody
	GetRequestId() *string
}

type CreateFolderResponseBody struct {
	// The information about the folder.
	Folder *CreateFolderResponseBodyFolder `json:"Folder,omitempty" xml:"Folder,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// C2CBCA30-C8DD-423E-B4AD-4FB694C9180C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFolderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFolderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFolderResponseBody) GetFolder() *CreateFolderResponseBodyFolder {
	return s.Folder
}

func (s *CreateFolderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFolderResponseBody) SetFolder(v *CreateFolderResponseBodyFolder) *CreateFolderResponseBody {
	s.Folder = v
	return s
}

func (s *CreateFolderResponseBody) SetRequestId(v string) *CreateFolderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFolderResponseBody) Validate() error {
	if s.Folder != nil {
		if err := s.Folder.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateFolderResponseBodyFolder struct {
	// The time when the folder was created.
	//
	// example:
	//
	// 2019-02-19T09:34:50.757Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the folder.
	//
	// example:
	//
	// fd-u8B321****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The name of the folder.
	//
	// example:
	//
	// rdFolder
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
	// The ID of the parent folder.
	//
	// example:
	//
	// r-b1****
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
}

func (s CreateFolderResponseBodyFolder) String() string {
	return dara.Prettify(s)
}

func (s CreateFolderResponseBodyFolder) GoString() string {
	return s.String()
}

func (s *CreateFolderResponseBodyFolder) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateFolderResponseBodyFolder) GetFolderId() *string {
	return s.FolderId
}

func (s *CreateFolderResponseBodyFolder) GetFolderName() *string {
	return s.FolderName
}

func (s *CreateFolderResponseBodyFolder) GetParentFolderId() *string {
	return s.ParentFolderId
}

func (s *CreateFolderResponseBodyFolder) SetCreateTime(v string) *CreateFolderResponseBodyFolder {
	s.CreateTime = &v
	return s
}

func (s *CreateFolderResponseBodyFolder) SetFolderId(v string) *CreateFolderResponseBodyFolder {
	s.FolderId = &v
	return s
}

func (s *CreateFolderResponseBodyFolder) SetFolderName(v string) *CreateFolderResponseBodyFolder {
	s.FolderName = &v
	return s
}

func (s *CreateFolderResponseBodyFolder) SetParentFolderId(v string) *CreateFolderResponseBodyFolder {
	s.ParentFolderId = &v
	return s
}

func (s *CreateFolderResponseBodyFolder) Validate() error {
	return dara.Validate(s)
}
