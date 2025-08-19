// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFolderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFolder(v *UpdateFolderResponseBodyFolder) *UpdateFolderResponseBody
	GetFolder() *UpdateFolderResponseBodyFolder
	SetRequestId(v string) *UpdateFolderResponseBody
	GetRequestId() *string
}

type UpdateFolderResponseBody struct {
	// The information about the folder.
	Folder *UpdateFolderResponseBodyFolder `json:"Folder,omitempty" xml:"Folder,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// C2CBCA30-C8DD-423E-B4AD-4FB694C9180C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFolderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFolderResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFolderResponseBody) GetFolder() *UpdateFolderResponseBodyFolder {
	return s.Folder
}

func (s *UpdateFolderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFolderResponseBody) SetFolder(v *UpdateFolderResponseBodyFolder) *UpdateFolderResponseBody {
	s.Folder = v
	return s
}

func (s *UpdateFolderResponseBody) SetRequestId(v string) *UpdateFolderResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFolderResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateFolderResponseBodyFolder struct {
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

func (s UpdateFolderResponseBodyFolder) String() string {
	return dara.Prettify(s)
}

func (s UpdateFolderResponseBodyFolder) GoString() string {
	return s.String()
}

func (s *UpdateFolderResponseBodyFolder) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateFolderResponseBodyFolder) GetFolderId() *string {
	return s.FolderId
}

func (s *UpdateFolderResponseBodyFolder) GetFolderName() *string {
	return s.FolderName
}

func (s *UpdateFolderResponseBodyFolder) GetParentFolderId() *string {
	return s.ParentFolderId
}

func (s *UpdateFolderResponseBodyFolder) SetCreateTime(v string) *UpdateFolderResponseBodyFolder {
	s.CreateTime = &v
	return s
}

func (s *UpdateFolderResponseBodyFolder) SetFolderId(v string) *UpdateFolderResponseBodyFolder {
	s.FolderId = &v
	return s
}

func (s *UpdateFolderResponseBodyFolder) SetFolderName(v string) *UpdateFolderResponseBodyFolder {
	s.FolderName = &v
	return s
}

func (s *UpdateFolderResponseBodyFolder) SetParentFolderId(v string) *UpdateFolderResponseBodyFolder {
	s.ParentFolderId = &v
	return s
}

func (s *UpdateFolderResponseBodyFolder) Validate() error {
	return dara.Validate(s)
}
