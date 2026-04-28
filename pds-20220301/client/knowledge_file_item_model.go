// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKnowledgeFileItem interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *KnowledgeFileItem
	GetDriveId() *string
	SetFileId(v string) *KnowledgeFileItem
	GetFileId() *string
}

type KnowledgeFileItem struct {
	// This parameter is required.
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// This parameter is required.
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

func (s KnowledgeFileItem) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeFileItem) GoString() string {
	return s.String()
}

func (s *KnowledgeFileItem) GetDriveId() *string {
	return s.DriveId
}

func (s *KnowledgeFileItem) GetFileId() *string {
	return s.FileId
}

func (s *KnowledgeFileItem) SetDriveId(v string) *KnowledgeFileItem {
	s.DriveId = &v
	return s
}

func (s *KnowledgeFileItem) SetFileId(v string) *KnowledgeFileItem {
	s.FileId = &v
	return s
}

func (s *KnowledgeFileItem) Validate() error {
	return dara.Validate(s)
}
