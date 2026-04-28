// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileLogDetail interface {
	dara.Model
	String() string
	GoString() string
	SetDecompressFileList(v []*string) *FileLogDetail
	GetDecompressFileList() []*string
	SetNewName(v string) *FileLogDetail
	GetNewName() *string
	SetParentPath(v string) *FileLogDetail
	GetParentPath() *string
	SetRevVersion(v int64) *FileLogDetail
	GetRevVersion() *int64
	SetRevisionId(v string) *FileLogDetail
	GetRevisionId() *string
	SetSize(v int64) *FileLogDetail
	GetSize() *int64
	SetToParentPath(v string) *FileLogDetail
	GetToParentPath() *string
	SetToParentPathType(v string) *FileLogDetail
	GetToParentPathType() *string
	SetType(v string) *FileLogDetail
	GetType() *string
}

type FileLogDetail struct {
	DecompressFileList []*string `json:"decompress_file_list,omitempty" xml:"decompress_file_list,omitempty" type:"Repeated"`
	NewName            *string   `json:"new_name,omitempty" xml:"new_name,omitempty"`
	ParentPath         *string   `json:"parent_path,omitempty" xml:"parent_path,omitempty"`
	RevVersion         *int64    `json:"rev_version,omitempty" xml:"rev_version,omitempty"`
	RevisionId         *string   `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
	Size               *int64    `json:"size,omitempty" xml:"size,omitempty"`
	ToParentPath       *string   `json:"to_parent_path,omitempty" xml:"to_parent_path,omitempty"`
	ToParentPathType   *string   `json:"to_parent_path_type,omitempty" xml:"to_parent_path_type,omitempty"`
	Type               *string   `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FileLogDetail) String() string {
	return dara.Prettify(s)
}

func (s FileLogDetail) GoString() string {
	return s.String()
}

func (s *FileLogDetail) GetDecompressFileList() []*string {
	return s.DecompressFileList
}

func (s *FileLogDetail) GetNewName() *string {
	return s.NewName
}

func (s *FileLogDetail) GetParentPath() *string {
	return s.ParentPath
}

func (s *FileLogDetail) GetRevVersion() *int64 {
	return s.RevVersion
}

func (s *FileLogDetail) GetRevisionId() *string {
	return s.RevisionId
}

func (s *FileLogDetail) GetSize() *int64 {
	return s.Size
}

func (s *FileLogDetail) GetToParentPath() *string {
	return s.ToParentPath
}

func (s *FileLogDetail) GetToParentPathType() *string {
	return s.ToParentPathType
}

func (s *FileLogDetail) GetType() *string {
	return s.Type
}

func (s *FileLogDetail) SetDecompressFileList(v []*string) *FileLogDetail {
	s.DecompressFileList = v
	return s
}

func (s *FileLogDetail) SetNewName(v string) *FileLogDetail {
	s.NewName = &v
	return s
}

func (s *FileLogDetail) SetParentPath(v string) *FileLogDetail {
	s.ParentPath = &v
	return s
}

func (s *FileLogDetail) SetRevVersion(v int64) *FileLogDetail {
	s.RevVersion = &v
	return s
}

func (s *FileLogDetail) SetRevisionId(v string) *FileLogDetail {
	s.RevisionId = &v
	return s
}

func (s *FileLogDetail) SetSize(v int64) *FileLogDetail {
	s.Size = &v
	return s
}

func (s *FileLogDetail) SetToParentPath(v string) *FileLogDetail {
	s.ToParentPath = &v
	return s
}

func (s *FileLogDetail) SetToParentPathType(v string) *FileLogDetail {
	s.ToParentPathType = &v
	return s
}

func (s *FileLogDetail) SetType(v string) *FileLogDetail {
	s.Type = &v
	return s
}

func (s *FileLogDetail) Validate() error {
	return dara.Validate(s)
}
