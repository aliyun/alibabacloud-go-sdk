// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFolder interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v int64) *Folder
	GetCreatedAt() *int64
	SetFolderId(v string) *Folder
	GetFolderId() *string
	SetModifiedAt(v int64) *Folder
	GetModifiedAt() *int64
	SetName(v string) *Folder
	GetName() *string
	SetNamespace(v string) *Folder
	GetNamespace() *string
	SetParentId(v string) *Folder
	GetParentId() *string
	SetSubFolder(v []*SubFolder) *Folder
	GetSubFolder() []*SubFolder
	SetWorkspace(v string) *Folder
	GetWorkspace() *string
}

type Folder struct {
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 00000000-0000-0000-0000-0000012312****
	FolderId   *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
	ModifiedAt *int64  `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// default-namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// 00000000-0000-0000-0000-0000012390****
	ParentId  *string      `json:"parentId,omitempty" xml:"parentId,omitempty"`
	SubFolder []*SubFolder `json:"subFolder,omitempty" xml:"subFolder,omitempty" type:"Repeated"`
	// example:
	//
	// edcef******b4f
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s Folder) String() string {
	return dara.Prettify(s)
}

func (s Folder) GoString() string {
	return s.String()
}

func (s *Folder) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *Folder) GetFolderId() *string {
	return s.FolderId
}

func (s *Folder) GetModifiedAt() *int64 {
	return s.ModifiedAt
}

func (s *Folder) GetName() *string {
	return s.Name
}

func (s *Folder) GetNamespace() *string {
	return s.Namespace
}

func (s *Folder) GetParentId() *string {
	return s.ParentId
}

func (s *Folder) GetSubFolder() []*SubFolder {
	return s.SubFolder
}

func (s *Folder) GetWorkspace() *string {
	return s.Workspace
}

func (s *Folder) SetCreatedAt(v int64) *Folder {
	s.CreatedAt = &v
	return s
}

func (s *Folder) SetFolderId(v string) *Folder {
	s.FolderId = &v
	return s
}

func (s *Folder) SetModifiedAt(v int64) *Folder {
	s.ModifiedAt = &v
	return s
}

func (s *Folder) SetName(v string) *Folder {
	s.Name = &v
	return s
}

func (s *Folder) SetNamespace(v string) *Folder {
	s.Namespace = &v
	return s
}

func (s *Folder) SetParentId(v string) *Folder {
	s.ParentId = &v
	return s
}

func (s *Folder) SetSubFolder(v []*SubFolder) *Folder {
	s.SubFolder = v
	return s
}

func (s *Folder) SetWorkspace(v string) *Folder {
	s.Workspace = &v
	return s
}

func (s *Folder) Validate() error {
	if s.SubFolder != nil {
		for _, item := range s.SubFolder {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
