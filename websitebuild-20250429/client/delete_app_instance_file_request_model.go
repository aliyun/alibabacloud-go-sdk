// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppInstanceFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *DeleteAppInstanceFileRequest
	GetConversationId() *string
	SetFilePath(v string) *DeleteAppInstanceFileRequest
	GetFilePath() *string
}

type DeleteAppInstanceFileRequest struct {
	// example:
	//
	// 81bc5a34-1d8d-4ef7-a208-7401c51b054b
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// sdms-test/static/
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
}

func (s DeleteAppInstanceFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppInstanceFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppInstanceFileRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *DeleteAppInstanceFileRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *DeleteAppInstanceFileRequest) SetConversationId(v string) *DeleteAppInstanceFileRequest {
	s.ConversationId = &v
	return s
}

func (s *DeleteAppInstanceFileRequest) SetFilePath(v string) *DeleteAppInstanceFileRequest {
	s.FilePath = &v
	return s
}

func (s *DeleteAppInstanceFileRequest) Validate() error {
	return dara.Validate(s)
}
