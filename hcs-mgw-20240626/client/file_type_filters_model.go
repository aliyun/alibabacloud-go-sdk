// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileTypeFilters interface {
	dara.Model
	String() string
	GoString() string
	SetExcludeDir(v bool) *FileTypeFilters
	GetExcludeDir() *bool
	SetExcludeSymlink(v bool) *FileTypeFilters
	GetExcludeSymlink() *bool
}

type FileTypeFilters struct {
	// example:
	//
	// fasle
	ExcludeDir *bool `json:"ExcludeDir,omitempty" xml:"ExcludeDir,omitempty"`
	// example:
	//
	// fasle
	ExcludeSymlink *bool `json:"ExcludeSymlink,omitempty" xml:"ExcludeSymlink,omitempty"`
}

func (s FileTypeFilters) String() string {
	return dara.Prettify(s)
}

func (s FileTypeFilters) GoString() string {
	return s.String()
}

func (s *FileTypeFilters) GetExcludeDir() *bool {
	return s.ExcludeDir
}

func (s *FileTypeFilters) GetExcludeSymlink() *bool {
	return s.ExcludeSymlink
}

func (s *FileTypeFilters) SetExcludeDir(v bool) *FileTypeFilters {
	s.ExcludeDir = &v
	return s
}

func (s *FileTypeFilters) SetExcludeSymlink(v bool) *FileTypeFilters {
	s.ExcludeSymlink = &v
	return s
}

func (s *FileTypeFilters) Validate() error {
	return dara.Validate(s)
}
