// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *DeleteAppFileRequest
	GetConversationId() *string
	SetFilePath(v string) *DeleteAppFileRequest
	GetFilePath() *string
}

type DeleteAppFileRequest struct {
	// example:
	//
	// 593fe1a2-d0b4-4fde-a2b0-78ad6a438d41
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// pt3/01/31/pengpeixin.png
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
}

func (s DeleteAppFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppFileRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *DeleteAppFileRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *DeleteAppFileRequest) SetConversationId(v string) *DeleteAppFileRequest {
	s.ConversationId = &v
	return s
}

func (s *DeleteAppFileRequest) SetFilePath(v string) *DeleteAppFileRequest {
	s.FilePath = &v
	return s
}

func (s *DeleteAppFileRequest) Validate() error {
	return dara.Validate(s)
}
