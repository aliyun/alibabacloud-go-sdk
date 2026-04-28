// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlbum interface {
	dara.Model
	String() string
	GoString() string
	SetAlbumId(v string) *Album
	GetAlbumId() *string
	SetBaseFaceFile(v *File) *Album
	GetBaseFaceFile() *File
	SetBaseFaceGroupId(v string) *Album
	GetBaseFaceGroupId() *string
	SetCoverFile(v *File) *Album
	GetCoverFile() *File
	SetCreatedAt(v string) *Album
	GetCreatedAt() *string
	SetDescription(v string) *Album
	GetDescription() *string
	SetFileCount(v int64) *Album
	GetFileCount() *int64
	SetName(v string) *Album
	GetName() *string
	SetOwner(v string) *Album
	GetOwner() *string
	SetUpdatedAt(v string) *Album
	GetUpdatedAt() *string
	SetUserTags(v map[string]*string) *Album
	GetUserTags() map[string]*string
}

type Album struct {
	AlbumId         *string            `json:"album_id,omitempty" xml:"album_id,omitempty"`
	BaseFaceFile    *File              `json:"base_face_file,omitempty" xml:"base_face_file,omitempty"`
	BaseFaceGroupId *string            `json:"base_face_group_id,omitempty" xml:"base_face_group_id,omitempty"`
	CoverFile       *File              `json:"cover_file,omitempty" xml:"cover_file,omitempty"`
	CreatedAt       *string            `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Description     *string            `json:"description,omitempty" xml:"description,omitempty"`
	FileCount       *int64             `json:"file_count,omitempty" xml:"file_count,omitempty"`
	Name            *string            `json:"name,omitempty" xml:"name,omitempty"`
	Owner           *string            `json:"owner,omitempty" xml:"owner,omitempty"`
	UpdatedAt       *string            `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	UserTags        map[string]*string `json:"user_tags,omitempty" xml:"user_tags,omitempty"`
}

func (s Album) String() string {
	return dara.Prettify(s)
}

func (s Album) GoString() string {
	return s.String()
}

func (s *Album) GetAlbumId() *string {
	return s.AlbumId
}

func (s *Album) GetBaseFaceFile() *File {
	return s.BaseFaceFile
}

func (s *Album) GetBaseFaceGroupId() *string {
	return s.BaseFaceGroupId
}

func (s *Album) GetCoverFile() *File {
	return s.CoverFile
}

func (s *Album) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *Album) GetDescription() *string {
	return s.Description
}

func (s *Album) GetFileCount() *int64 {
	return s.FileCount
}

func (s *Album) GetName() *string {
	return s.Name
}

func (s *Album) GetOwner() *string {
	return s.Owner
}

func (s *Album) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *Album) GetUserTags() map[string]*string {
	return s.UserTags
}

func (s *Album) SetAlbumId(v string) *Album {
	s.AlbumId = &v
	return s
}

func (s *Album) SetBaseFaceFile(v *File) *Album {
	s.BaseFaceFile = v
	return s
}

func (s *Album) SetBaseFaceGroupId(v string) *Album {
	s.BaseFaceGroupId = &v
	return s
}

func (s *Album) SetCoverFile(v *File) *Album {
	s.CoverFile = v
	return s
}

func (s *Album) SetCreatedAt(v string) *Album {
	s.CreatedAt = &v
	return s
}

func (s *Album) SetDescription(v string) *Album {
	s.Description = &v
	return s
}

func (s *Album) SetFileCount(v int64) *Album {
	s.FileCount = &v
	return s
}

func (s *Album) SetName(v string) *Album {
	s.Name = &v
	return s
}

func (s *Album) SetOwner(v string) *Album {
	s.Owner = &v
	return s
}

func (s *Album) SetUpdatedAt(v string) *Album {
	s.UpdatedAt = &v
	return s
}

func (s *Album) SetUserTags(v map[string]*string) *Album {
	s.UserTags = v
	return s
}

func (s *Album) Validate() error {
	if s.BaseFaceFile != nil {
		if err := s.BaseFaceFile.Validate(); err != nil {
			return err
		}
	}
	if s.CoverFile != nil {
		if err := s.CoverFile.Validate(); err != nil {
			return err
		}
	}
	return nil
}
