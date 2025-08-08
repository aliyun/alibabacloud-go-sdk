// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelFilePreview interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *ModelFilePreview
	GetContent() *string
	SetHash(v string) *ModelFilePreview
	GetHash() *string
	SetIsCompressedImage(v bool) *ModelFilePreview
	GetIsCompressedImage() *bool
	SetIsDir(v bool) *ModelFilePreview
	GetIsDir() *bool
	SetName(v string) *ModelFilePreview
	GetName() *string
	SetPath(v string) *ModelFilePreview
	GetPath() *string
	SetSize(v int64) *ModelFilePreview
	GetSize() *int64
	SetUnpreviewable(v bool) *ModelFilePreview
	GetUnpreviewable() *bool
}

type ModelFilePreview struct {
	Content           *string `json:"content,omitempty" xml:"content,omitempty"`
	Hash              *string `json:"hash,omitempty" xml:"hash,omitempty"`
	IsCompressedImage *bool   `json:"isCompressedImage,omitempty" xml:"isCompressedImage,omitempty"`
	IsDir             *bool   `json:"isDir,omitempty" xml:"isDir,omitempty"`
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
	Path              *string `json:"path,omitempty" xml:"path,omitempty"`
	Size              *int64  `json:"size,omitempty" xml:"size,omitempty"`
	Unpreviewable     *bool   `json:"unpreviewable,omitempty" xml:"unpreviewable,omitempty"`
}

func (s ModelFilePreview) String() string {
	return dara.Prettify(s)
}

func (s ModelFilePreview) GoString() string {
	return s.String()
}

func (s *ModelFilePreview) GetContent() *string {
	return s.Content
}

func (s *ModelFilePreview) GetHash() *string {
	return s.Hash
}

func (s *ModelFilePreview) GetIsCompressedImage() *bool {
	return s.IsCompressedImage
}

func (s *ModelFilePreview) GetIsDir() *bool {
	return s.IsDir
}

func (s *ModelFilePreview) GetName() *string {
	return s.Name
}

func (s *ModelFilePreview) GetPath() *string {
	return s.Path
}

func (s *ModelFilePreview) GetSize() *int64 {
	return s.Size
}

func (s *ModelFilePreview) GetUnpreviewable() *bool {
	return s.Unpreviewable
}

func (s *ModelFilePreview) SetContent(v string) *ModelFilePreview {
	s.Content = &v
	return s
}

func (s *ModelFilePreview) SetHash(v string) *ModelFilePreview {
	s.Hash = &v
	return s
}

func (s *ModelFilePreview) SetIsCompressedImage(v bool) *ModelFilePreview {
	s.IsCompressedImage = &v
	return s
}

func (s *ModelFilePreview) SetIsDir(v bool) *ModelFilePreview {
	s.IsDir = &v
	return s
}

func (s *ModelFilePreview) SetName(v string) *ModelFilePreview {
	s.Name = &v
	return s
}

func (s *ModelFilePreview) SetPath(v string) *ModelFilePreview {
	s.Path = &v
	return s
}

func (s *ModelFilePreview) SetSize(v int64) *ModelFilePreview {
	s.Size = &v
	return s
}

func (s *ModelFilePreview) SetUnpreviewable(v bool) *ModelFilePreview {
	s.Unpreviewable = &v
	return s
}

func (s *ModelFilePreview) Validate() error {
	return dara.Validate(s)
}
