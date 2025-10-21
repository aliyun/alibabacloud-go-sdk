// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubFolder interface {
	dara.Model
	String() string
	GoString() string
	SetFolderId(v string) *SubFolder
	GetFolderId() *string
	SetName(v string) *SubFolder
	GetName() *string
	SetParentId(v string) *SubFolder
	GetParentId() *string
}

type SubFolder struct {
	// example:
	//
	// a579aec9-1d5e-3382-9d65-9887ff6cfaff
	FolderId *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 95c0787c-408f-4e1f-88ba-ef0a84a2c2ee
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
}

func (s SubFolder) String() string {
	return dara.Prettify(s)
}

func (s SubFolder) GoString() string {
	return s.String()
}

func (s *SubFolder) GetFolderId() *string {
	return s.FolderId
}

func (s *SubFolder) GetName() *string {
	return s.Name
}

func (s *SubFolder) GetParentId() *string {
	return s.ParentId
}

func (s *SubFolder) SetFolderId(v string) *SubFolder {
	s.FolderId = &v
	return s
}

func (s *SubFolder) SetName(v string) *SubFolder {
	s.Name = &v
	return s
}

func (s *SubFolder) SetParentId(v string) *SubFolder {
	s.ParentId = &v
	return s
}

func (s *SubFolder) Validate() error {
	return dara.Validate(s)
}
