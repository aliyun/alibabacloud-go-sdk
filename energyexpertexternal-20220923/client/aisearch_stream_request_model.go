// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAISearchStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFolderId(v string) *AISearchStreamRequest
	GetFolderId() *string
	SetMessage(v []*AISearchMessageItem) *AISearchStreamRequest
	GetMessage() []*AISearchMessageItem
	SetQuestion(v string) *AISearchStreamRequest
	GetQuestion() *string
	SetResourceTypeNeeded(v []*string) *AISearchStreamRequest
	GetResourceTypeNeeded() []*string
}

type AISearchStreamRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1a851c4a-1d65-11ef-99a7-ssfsfdd
	FolderId *string                `json:"folderId,omitempty" xml:"folderId,omitempty"`
	Message  []*AISearchMessageItem `json:"message,omitempty" xml:"message,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// “my net is slow, what can I do?”
	Question           *string   `json:"question,omitempty" xml:"question,omitempty"`
	ResourceTypeNeeded []*string `json:"resourceTypeNeeded,omitempty" xml:"resourceTypeNeeded,omitempty" type:"Repeated"`
}

func (s AISearchStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s AISearchStreamRequest) GoString() string {
	return s.String()
}

func (s *AISearchStreamRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *AISearchStreamRequest) GetMessage() []*AISearchMessageItem {
	return s.Message
}

func (s *AISearchStreamRequest) GetQuestion() *string {
	return s.Question
}

func (s *AISearchStreamRequest) GetResourceTypeNeeded() []*string {
	return s.ResourceTypeNeeded
}

func (s *AISearchStreamRequest) SetFolderId(v string) *AISearchStreamRequest {
	s.FolderId = &v
	return s
}

func (s *AISearchStreamRequest) SetMessage(v []*AISearchMessageItem) *AISearchStreamRequest {
	s.Message = v
	return s
}

func (s *AISearchStreamRequest) SetQuestion(v string) *AISearchStreamRequest {
	s.Question = &v
	return s
}

func (s *AISearchStreamRequest) SetResourceTypeNeeded(v []*string) *AISearchStreamRequest {
	s.ResourceTypeNeeded = v
	return s
}

func (s *AISearchStreamRequest) Validate() error {
	if s.Message != nil {
		for _, item := range s.Message {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
