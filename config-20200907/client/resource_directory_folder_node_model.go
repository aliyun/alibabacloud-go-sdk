// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceDirectoryFolderNode interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *ResourceDirectoryFolderNode
	GetAccountId() *string
	SetChildren(v []*ResourceDirectoryFolderNode) *ResourceDirectoryFolderNode
	GetChildren() []*ResourceDirectoryFolderNode
	SetDisplayName(v string) *ResourceDirectoryFolderNode
	GetDisplayName() *string
	SetFolderId(v string) *ResourceDirectoryFolderNode
	GetFolderId() *string
	SetFolderName(v string) *ResourceDirectoryFolderNode
	GetFolderName() *string
	SetParentFolderId(v string) *ResourceDirectoryFolderNode
	GetParentFolderId() *string
}

type ResourceDirectoryFolderNode struct {
	AccountId      *string                        `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	Children       []*ResourceDirectoryFolderNode `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	DisplayName    *string                        `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	FolderId       *string                        `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	FolderName     *string                        `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
	ParentFolderId *string                        `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
}

func (s ResourceDirectoryFolderNode) String() string {
	return dara.Prettify(s)
}

func (s ResourceDirectoryFolderNode) GoString() string {
	return s.String()
}

func (s *ResourceDirectoryFolderNode) GetAccountId() *string {
	return s.AccountId
}

func (s *ResourceDirectoryFolderNode) GetChildren() []*ResourceDirectoryFolderNode {
	return s.Children
}

func (s *ResourceDirectoryFolderNode) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ResourceDirectoryFolderNode) GetFolderId() *string {
	return s.FolderId
}

func (s *ResourceDirectoryFolderNode) GetFolderName() *string {
	return s.FolderName
}

func (s *ResourceDirectoryFolderNode) GetParentFolderId() *string {
	return s.ParentFolderId
}

func (s *ResourceDirectoryFolderNode) SetAccountId(v string) *ResourceDirectoryFolderNode {
	s.AccountId = &v
	return s
}

func (s *ResourceDirectoryFolderNode) SetChildren(v []*ResourceDirectoryFolderNode) *ResourceDirectoryFolderNode {
	s.Children = v
	return s
}

func (s *ResourceDirectoryFolderNode) SetDisplayName(v string) *ResourceDirectoryFolderNode {
	s.DisplayName = &v
	return s
}

func (s *ResourceDirectoryFolderNode) SetFolderId(v string) *ResourceDirectoryFolderNode {
	s.FolderId = &v
	return s
}

func (s *ResourceDirectoryFolderNode) SetFolderName(v string) *ResourceDirectoryFolderNode {
	s.FolderName = &v
	return s
}

func (s *ResourceDirectoryFolderNode) SetParentFolderId(v string) *ResourceDirectoryFolderNode {
	s.ParentFolderId = &v
	return s
}

func (s *ResourceDirectoryFolderNode) Validate() error {
	if s.Children != nil {
		for _, item := range s.Children {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
