// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppFileContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *GetAppFileContentRequest
	GetConversationId() *string
	SetFilePath(v string) *GetAppFileContentRequest
	GetFilePath() *string
}

type GetAppFileContentRequest struct {
	// example:
	//
	// 593fe1a2-d0b4-4fde-a2b0-78ad6a438d41
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// index.html
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
}

func (s GetAppFileContentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppFileContentRequest) GoString() string {
	return s.String()
}

func (s *GetAppFileContentRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *GetAppFileContentRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *GetAppFileContentRequest) SetConversationId(v string) *GetAppFileContentRequest {
	s.ConversationId = &v
	return s
}

func (s *GetAppFileContentRequest) SetFilePath(v string) *GetAppFileContentRequest {
	s.FilePath = &v
	return s
}

func (s *GetAppFileContentRequest) Validate() error {
	return dara.Validate(s)
}
