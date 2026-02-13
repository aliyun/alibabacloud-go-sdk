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
	// The content of the answer to the question.
	//
	// example:
	//
	// No related content found.
	Answer *string `json:"answer,omitempty" xml:"answer,omitempty"`
	// Q&A time, in milliseconds timestamp.
	//
	// example:
	//
	// 1747280158000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Folder selected for the Q&A.
	//
	// example:
	//
	// 7708fddb-21dc-4403-a4ea-5b94eccce4c3
	FolderId *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
	// Name of the folder selected for the Q&A.
	//
	// example:
	//
	// deafult folder
	FolderName *string `json:"folderName,omitempty" xml:"folderName,omitempty"`
	// Question content.
	//
	// example:
	//
	// How to use the knowledge base.
	Question *string `json:"question,omitempty" xml:"question,omitempty"`
	// - List of files related to the Q&A.
	//
	// - If streaming question answering is used, only the first shard contains data.
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
	if s.RefDocList != nil {
		for _, item := range s.RefDocList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
