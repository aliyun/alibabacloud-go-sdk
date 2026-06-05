// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UpdateAppFileRequest
	GetContent() *string
	SetConversationId(v string) *UpdateAppFileRequest
	GetConversationId() *string
	SetFilePath(v string) *UpdateAppFileRequest
	GetFilePath() *string
}

type UpdateAppFileRequest struct {
	// example:
	//
	// verify_46630893e2b5efde444c82b4e3707adb
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 593fe1a2-d0b4-4fde-a2b0-78ad6a438d41
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// pt3/01/31/pengpeixin.png
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
}

func (s UpdateAppFileRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppFileRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppFileRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateAppFileRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *UpdateAppFileRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *UpdateAppFileRequest) SetContent(v string) *UpdateAppFileRequest {
	s.Content = &v
	return s
}

func (s *UpdateAppFileRequest) SetConversationId(v string) *UpdateAppFileRequest {
	s.ConversationId = &v
	return s
}

func (s *UpdateAppFileRequest) SetFilePath(v string) *UpdateAppFileRequest {
	s.FilePath = &v
	return s
}

func (s *UpdateAppFileRequest) Validate() error {
	return dara.Validate(s)
}
