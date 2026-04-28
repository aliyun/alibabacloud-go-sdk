// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUncompressedFileInfo interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *UncompressedFileInfo
	GetDriveId() *string
	SetFileId(v string) *UncompressedFileInfo
	GetFileId() *string
	SetIsFolder(v bool) *UncompressedFileInfo
	GetIsFolder() *bool
	SetItems(v []*UncompressedFileInfo) *UncompressedFileInfo
	GetItems() []*UncompressedFileInfo
	SetName(v string) *UncompressedFileInfo
	GetName() *string
	SetSize(v int64) *UncompressedFileInfo
	GetSize() *int64
	SetUpdatedAt(v int64) *UncompressedFileInfo
	GetUpdatedAt() *int64
}

type UncompressedFileInfo struct {
	// The drive ID.
	//
	// example:
	//
	// 5060
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 66972349b2b12fa309a143fb9db29647b2ddabfd
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// Whether it is a folder.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsFolder *bool `json:"is_folder,omitempty" xml:"is_folder,omitempty"`
	// Subfiles
	Items []*UncompressedFileInfo `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The name of the file.
	//
	// example:
	//
	// 1.mov
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The size of the file.
	//
	// example:
	//
	// 218052
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// Update time.
	//
	// example:
	//
	// 1721368014000
	UpdatedAt *int64 `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s UncompressedFileInfo) String() string {
	return dara.Prettify(s)
}

func (s UncompressedFileInfo) GoString() string {
	return s.String()
}

func (s *UncompressedFileInfo) GetDriveId() *string {
	return s.DriveId
}

func (s *UncompressedFileInfo) GetFileId() *string {
	return s.FileId
}

func (s *UncompressedFileInfo) GetIsFolder() *bool {
	return s.IsFolder
}

func (s *UncompressedFileInfo) GetItems() []*UncompressedFileInfo {
	return s.Items
}

func (s *UncompressedFileInfo) GetName() *string {
	return s.Name
}

func (s *UncompressedFileInfo) GetSize() *int64 {
	return s.Size
}

func (s *UncompressedFileInfo) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *UncompressedFileInfo) SetDriveId(v string) *UncompressedFileInfo {
	s.DriveId = &v
	return s
}

func (s *UncompressedFileInfo) SetFileId(v string) *UncompressedFileInfo {
	s.FileId = &v
	return s
}

func (s *UncompressedFileInfo) SetIsFolder(v bool) *UncompressedFileInfo {
	s.IsFolder = &v
	return s
}

func (s *UncompressedFileInfo) SetItems(v []*UncompressedFileInfo) *UncompressedFileInfo {
	s.Items = v
	return s
}

func (s *UncompressedFileInfo) SetName(v string) *UncompressedFileInfo {
	s.Name = &v
	return s
}

func (s *UncompressedFileInfo) SetSize(v int64) *UncompressedFileInfo {
	s.Size = &v
	return s
}

func (s *UncompressedFileInfo) SetUpdatedAt(v int64) *UncompressedFileInfo {
	s.UpdatedAt = &v
	return s
}

func (s *UncompressedFileInfo) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
