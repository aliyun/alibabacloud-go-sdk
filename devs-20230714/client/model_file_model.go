// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelFile interface {
	dara.Model
	String() string
	GoString() string
	SetIsDir(v bool) *ModelFile
	GetIsDir() *bool
	SetModeTime(v int64) *ModelFile
	GetModeTime() *int64
	SetName(v string) *ModelFile
	GetName() *string
	SetPath(v string) *ModelFile
	GetPath() *string
	SetSize(v int64) *ModelFile
	GetSize() *int64
}

type ModelFile struct {
	IsDir    *bool   `json:"isDir,omitempty" xml:"isDir,omitempty"`
	ModeTime *int64  `json:"modeTime,omitempty" xml:"modeTime,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Path     *string `json:"path,omitempty" xml:"path,omitempty"`
	Size     *int64  `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ModelFile) String() string {
	return dara.Prettify(s)
}

func (s ModelFile) GoString() string {
	return s.String()
}

func (s *ModelFile) GetIsDir() *bool {
	return s.IsDir
}

func (s *ModelFile) GetModeTime() *int64 {
	return s.ModeTime
}

func (s *ModelFile) GetName() *string {
	return s.Name
}

func (s *ModelFile) GetPath() *string {
	return s.Path
}

func (s *ModelFile) GetSize() *int64 {
	return s.Size
}

func (s *ModelFile) SetIsDir(v bool) *ModelFile {
	s.IsDir = &v
	return s
}

func (s *ModelFile) SetModeTime(v int64) *ModelFile {
	s.ModeTime = &v
	return s
}

func (s *ModelFile) SetName(v string) *ModelFile {
	s.Name = &v
	return s
}

func (s *ModelFile) SetPath(v string) *ModelFile {
	s.Path = &v
	return s
}

func (s *ModelFile) SetSize(v int64) *ModelFile {
	s.Size = &v
	return s
}

func (s *ModelFile) Validate() error {
	return dara.Validate(s)
}
