// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotKnowledgeBaseFile interface {
	dara.Model
	String() string
	GoString() string
	SetActionCount(v int64) *HotKnowledgeBaseFile
	GetActionCount() *int64
	SetActionList(v []*string) *HotKnowledgeBaseFile
	GetActionList() []*string
	SetCategory(v string) *HotKnowledgeBaseFile
	GetCategory() *string
	SetCountAt(v int64) *HotKnowledgeBaseFile
	GetCountAt() *int64
	SetDriveId(v string) *HotKnowledgeBaseFile
	GetDriveId() *string
	SetFileId(v string) *HotKnowledgeBaseFile
	GetFileId() *string
	SetKnowledgeBaseId(v string) *HotKnowledgeBaseFile
	GetKnowledgeBaseId() *string
	SetName(v string) *HotKnowledgeBaseFile
	GetName() *string
	SetRevisionId(v string) *HotKnowledgeBaseFile
	GetRevisionId() *string
}

type HotKnowledgeBaseFile struct {
	// example:
	//
	// 1
	ActionCount *int64    `json:"action_count,omitempty" xml:"action_count,omitempty"`
	ActionList  []*string `json:"action_list,omitempty" xml:"action_list,omitempty" type:"Repeated"`
	// example:
	//
	// image
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// 1727578860000
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
	// 4jTsp3AgW
	KnowledgeBaseId *string `json:"knowledge_base_id,omitempty" xml:"knowledge_base_id,omitempty"`
	// example:
	//
	// a.jpg
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 666ff36c22278f023ec
	RevisionId *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
}

func (s HotKnowledgeBaseFile) String() string {
	return dara.Prettify(s)
}

func (s HotKnowledgeBaseFile) GoString() string {
	return s.String()
}

func (s *HotKnowledgeBaseFile) GetActionCount() *int64 {
	return s.ActionCount
}

func (s *HotKnowledgeBaseFile) GetActionList() []*string {
	return s.ActionList
}

func (s *HotKnowledgeBaseFile) GetCategory() *string {
	return s.Category
}

func (s *HotKnowledgeBaseFile) GetCountAt() *int64 {
	return s.CountAt
}

func (s *HotKnowledgeBaseFile) GetDriveId() *string {
	return s.DriveId
}

func (s *HotKnowledgeBaseFile) GetFileId() *string {
	return s.FileId
}

func (s *HotKnowledgeBaseFile) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *HotKnowledgeBaseFile) GetName() *string {
	return s.Name
}

func (s *HotKnowledgeBaseFile) GetRevisionId() *string {
	return s.RevisionId
}

func (s *HotKnowledgeBaseFile) SetActionCount(v int64) *HotKnowledgeBaseFile {
	s.ActionCount = &v
	return s
}

func (s *HotKnowledgeBaseFile) SetActionList(v []*string) *HotKnowledgeBaseFile {
	s.ActionList = v
	return s
}

func (s *HotKnowledgeBaseFile) SetCategory(v string) *HotKnowledgeBaseFile {
	s.Category = &v
	return s
}

func (s *HotKnowledgeBaseFile) SetCountAt(v int64) *HotKnowledgeBaseFile {
	s.CountAt = &v
	return s
}

func (s *HotKnowledgeBaseFile) SetDriveId(v string) *HotKnowledgeBaseFile {
	s.DriveId = &v
	return s
}

func (s *HotKnowledgeBaseFile) SetFileId(v string) *HotKnowledgeBaseFile {
	s.FileId = &v
	return s
}

func (s *HotKnowledgeBaseFile) SetKnowledgeBaseId(v string) *HotKnowledgeBaseFile {
	s.KnowledgeBaseId = &v
	return s
}

func (s *HotKnowledgeBaseFile) SetName(v string) *HotKnowledgeBaseFile {
	s.Name = &v
	return s
}

func (s *HotKnowledgeBaseFile) SetRevisionId(v string) *HotKnowledgeBaseFile {
	s.RevisionId = &v
	return s
}

func (s *HotKnowledgeBaseFile) Validate() error {
	return dara.Validate(s)
}
