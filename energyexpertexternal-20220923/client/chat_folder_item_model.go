// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatFolderItem interface {
	dara.Model
	String() string
	GoString() string
	SetFolderId(v string) *ChatFolderItem
	GetFolderId() *string
	SetFolderName(v string) *ChatFolderItem
	GetFolderName() *string
	SetSubFolders(v []*ChatItem) *ChatFolderItem
	GetSubFolders() []*ChatItem
}

type ChatFolderItem struct {
	FolderId   *string     `json:"folderId,omitempty" xml:"folderId,omitempty"`
	FolderName *string     `json:"folderName,omitempty" xml:"folderName,omitempty"`
	SubFolders []*ChatItem `json:"subFolders,omitempty" xml:"subFolders,omitempty" type:"Repeated"`
}

func (s ChatFolderItem) String() string {
	return dara.Prettify(s)
}

func (s ChatFolderItem) GoString() string {
	return s.String()
}

func (s *ChatFolderItem) GetFolderId() *string {
	return s.FolderId
}

func (s *ChatFolderItem) GetFolderName() *string {
	return s.FolderName
}

func (s *ChatFolderItem) GetSubFolders() []*ChatItem {
	return s.SubFolders
}

func (s *ChatFolderItem) SetFolderId(v string) *ChatFolderItem {
	s.FolderId = &v
	return s
}

func (s *ChatFolderItem) SetFolderName(v string) *ChatFolderItem {
	s.FolderName = &v
	return s
}

func (s *ChatFolderItem) SetSubFolders(v []*ChatItem) *ChatFolderItem {
	s.SubFolders = v
	return s
}

func (s *ChatFolderItem) Validate() error {
	return dara.Validate(s)
}
