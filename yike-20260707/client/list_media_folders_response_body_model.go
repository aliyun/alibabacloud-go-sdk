// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaFoldersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFolderList(v []*ListMediaFoldersResponseBodyFolderList) *ListMediaFoldersResponseBody
	GetFolderList() []*ListMediaFoldersResponseBodyFolderList
	SetRequestId(v string) *ListMediaFoldersResponseBody
	GetRequestId() *string
}

type ListMediaFoldersResponseBody struct {
	// The list of folders.
	FolderList []*ListMediaFoldersResponseBodyFolderList `json:"FolderList,omitempty" xml:"FolderList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4B6524F7-*****-A55F-B7DEA1EC628A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMediaFoldersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMediaFoldersResponseBody) GoString() string {
	return s.String()
}

func (s *ListMediaFoldersResponseBody) GetFolderList() []*ListMediaFoldersResponseBodyFolderList {
	return s.FolderList
}

func (s *ListMediaFoldersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMediaFoldersResponseBody) SetFolderList(v []*ListMediaFoldersResponseBodyFolderList) *ListMediaFoldersResponseBody {
	s.FolderList = v
	return s
}

func (s *ListMediaFoldersResponseBody) SetRequestId(v string) *ListMediaFoldersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMediaFoldersResponseBody) Validate() error {
	if s.FolderList != nil {
		for _, item := range s.FolderList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMediaFoldersResponseBodyFolderList struct {
	// The creation time.
	//
	// example:
	//
	// -
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The folder ID. The ID of the default folder is fixed to `fd_{ProductionId}_default`.
	//
	// example:
	//
	// fd_pd_1585***_default
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The folder name.
	//
	// example:
	//
	// Default folder
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
	// Indicates whether the folder is the default folder. Each project has one and only one default folder.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The modification time.
	//
	// example:
	//
	// -
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the project to which the folder belongs.
	//
	// example:
	//
	// pd_1585***
	ProductionId *string `json:"ProductionId,omitempty" xml:"ProductionId,omitempty"`
}

func (s ListMediaFoldersResponseBodyFolderList) String() string {
	return dara.Prettify(s)
}

func (s ListMediaFoldersResponseBodyFolderList) GoString() string {
	return s.String()
}

func (s *ListMediaFoldersResponseBodyFolderList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListMediaFoldersResponseBodyFolderList) GetFolderId() *string {
	return s.FolderId
}

func (s *ListMediaFoldersResponseBodyFolderList) GetFolderName() *string {
	return s.FolderName
}

func (s *ListMediaFoldersResponseBodyFolderList) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListMediaFoldersResponseBodyFolderList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListMediaFoldersResponseBodyFolderList) GetProductionId() *string {
	return s.ProductionId
}

func (s *ListMediaFoldersResponseBodyFolderList) SetCreateTime(v string) *ListMediaFoldersResponseBodyFolderList {
	s.CreateTime = &v
	return s
}

func (s *ListMediaFoldersResponseBodyFolderList) SetFolderId(v string) *ListMediaFoldersResponseBodyFolderList {
	s.FolderId = &v
	return s
}

func (s *ListMediaFoldersResponseBodyFolderList) SetFolderName(v string) *ListMediaFoldersResponseBodyFolderList {
	s.FolderName = &v
	return s
}

func (s *ListMediaFoldersResponseBodyFolderList) SetIsDefault(v bool) *ListMediaFoldersResponseBodyFolderList {
	s.IsDefault = &v
	return s
}

func (s *ListMediaFoldersResponseBodyFolderList) SetModifiedTime(v string) *ListMediaFoldersResponseBodyFolderList {
	s.ModifiedTime = &v
	return s
}

func (s *ListMediaFoldersResponseBodyFolderList) SetProductionId(v string) *ListMediaFoldersResponseBodyFolderList {
	s.ProductionId = &v
	return s
}

func (s *ListMediaFoldersResponseBodyFolderList) Validate() error {
	return dara.Validate(s)
}
