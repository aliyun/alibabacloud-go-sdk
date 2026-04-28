// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotDriveFile interface {
	dara.Model
	String() string
	GoString() string
	SetActionCount(v int64) *HotDriveFile
	GetActionCount() *int64
	SetActionList(v []*string) *HotDriveFile
	GetActionList() []*string
	SetCategory(v string) *HotDriveFile
	GetCategory() *string
	SetCountAt(v int64) *HotDriveFile
	GetCountAt() *int64
	SetDriveId(v string) *HotDriveFile
	GetDriveId() *string
	SetFileId(v string) *HotDriveFile
	GetFileId() *string
	SetName(v string) *HotDriveFile
	GetName() *string
	SetRevisionId(v string) *HotDriveFile
	GetRevisionId() *string
}

type HotDriveFile struct {
	// example:
	//
	// 2
	ActionCount *int64    `json:"action_count,omitempty" xml:"action_count,omitempty"`
	ActionList  []*string `json:"action_list,omitempty" xml:"action_list,omitempty" type:"Repeated"`
	// example:
	//
	// doc
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// 1727059860000
	CountAt *int64 `json:"count_at,omitempty" xml:"count_at,omitempty"`
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// example:
	//
	// 666ff36c22278f023ec
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// example:
	//
	// a.jpg
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 666ff36c22278f023ec
	RevisionId *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
}

func (s HotDriveFile) String() string {
	return dara.Prettify(s)
}

func (s HotDriveFile) GoString() string {
	return s.String()
}

func (s *HotDriveFile) GetActionCount() *int64 {
	return s.ActionCount
}

func (s *HotDriveFile) GetActionList() []*string {
	return s.ActionList
}

func (s *HotDriveFile) GetCategory() *string {
	return s.Category
}

func (s *HotDriveFile) GetCountAt() *int64 {
	return s.CountAt
}

func (s *HotDriveFile) GetDriveId() *string {
	return s.DriveId
}

func (s *HotDriveFile) GetFileId() *string {
	return s.FileId
}

func (s *HotDriveFile) GetName() *string {
	return s.Name
}

func (s *HotDriveFile) GetRevisionId() *string {
	return s.RevisionId
}

func (s *HotDriveFile) SetActionCount(v int64) *HotDriveFile {
	s.ActionCount = &v
	return s
}

func (s *HotDriveFile) SetActionList(v []*string) *HotDriveFile {
	s.ActionList = v
	return s
}

func (s *HotDriveFile) SetCategory(v string) *HotDriveFile {
	s.Category = &v
	return s
}

func (s *HotDriveFile) SetCountAt(v int64) *HotDriveFile {
	s.CountAt = &v
	return s
}

func (s *HotDriveFile) SetDriveId(v string) *HotDriveFile {
	s.DriveId = &v
	return s
}

func (s *HotDriveFile) SetFileId(v string) *HotDriveFile {
	s.FileId = &v
	return s
}

func (s *HotDriveFile) SetName(v string) *HotDriveFile {
	s.Name = &v
	return s
}

func (s *HotDriveFile) SetRevisionId(v string) *HotDriveFile {
	s.RevisionId = &v
	return s
}

func (s *HotDriveFile) Validate() error {
	return dara.Validate(s)
}
