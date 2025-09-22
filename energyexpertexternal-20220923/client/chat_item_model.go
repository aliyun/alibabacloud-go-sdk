// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatItem interface {
	dara.Model
	String() string
	GoString() string
	SetAnswer(v string) *ChatItem
	GetAnswer() *string
	SetCreateTime(v int64) *ChatItem
	GetCreateTime() *int64
	SetFolderId(v string) *ChatItem
	GetFolderId() *string
	SetFolderName(v string) *ChatItem
	GetFolderName() *string
	SetQuestion(v string) *ChatItem
	GetQuestion() *string
	SetRefDocList(v []*ChatRefDocItem) *ChatItem
	GetRefDocList() []*ChatRefDocItem
}

type ChatItem struct {
	Answer     *string           `json:"answer,omitempty" xml:"answer,omitempty"`
	CreateTime *int64            `json:"createTime,omitempty" xml:"createTime,omitempty"`
	FolderId   *string           `json:"folderId,omitempty" xml:"folderId,omitempty"`
	FolderName *string           `json:"folderName,omitempty" xml:"folderName,omitempty"`
	Question   *string           `json:"question,omitempty" xml:"question,omitempty"`
	RefDocList []*ChatRefDocItem `json:"refDocList,omitempty" xml:"refDocList,omitempty" type:"Repeated"`
}

func (s ChatItem) String() string {
	return dara.Prettify(s)
}

func (s ChatItem) GoString() string {
	return s.String()
}

func (s *ChatItem) GetAnswer() *string {
	return s.Answer
}

func (s *ChatItem) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ChatItem) GetFolderId() *string {
	return s.FolderId
}

func (s *ChatItem) GetFolderName() *string {
	return s.FolderName
}

func (s *ChatItem) GetQuestion() *string {
	return s.Question
}

func (s *ChatItem) GetRefDocList() []*ChatRefDocItem {
	return s.RefDocList
}

func (s *ChatItem) SetAnswer(v string) *ChatItem {
	s.Answer = &v
	return s
}

func (s *ChatItem) SetCreateTime(v int64) *ChatItem {
	s.CreateTime = &v
	return s
}

func (s *ChatItem) SetFolderId(v string) *ChatItem {
	s.FolderId = &v
	return s
}

func (s *ChatItem) SetFolderName(v string) *ChatItem {
	s.FolderName = &v
	return s
}

func (s *ChatItem) SetQuestion(v string) *ChatItem {
	s.Question = &v
	return s
}

func (s *ChatItem) SetRefDocList(v []*ChatRefDocItem) *ChatItem {
	s.RefDocList = v
	return s
}

func (s *ChatItem) Validate() error {
	return dara.Validate(s)
}
