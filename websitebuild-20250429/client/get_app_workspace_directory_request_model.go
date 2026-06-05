// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppWorkspaceDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *GetAppWorkspaceDirectoryRequest
	GetConversationId() *string
	SetDeep(v int32) *GetAppWorkspaceDirectoryRequest
	GetDeep() *int32
	SetFilePath(v string) *GetAppWorkspaceDirectoryRequest
	GetFilePath() *string
}

type GetAppWorkspaceDirectoryRequest struct {
	// example:
	//
	// 5b7105a2-2999-430b-ba23-ba09149d5434
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// 1
	Deep *int32 `json:"Deep,omitempty" xml:"Deep,omitempty"`
	// example:
	//
	// 1
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
}

func (s GetAppWorkspaceDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppWorkspaceDirectoryRequest) GoString() string {
	return s.String()
}

func (s *GetAppWorkspaceDirectoryRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *GetAppWorkspaceDirectoryRequest) GetDeep() *int32 {
	return s.Deep
}

func (s *GetAppWorkspaceDirectoryRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *GetAppWorkspaceDirectoryRequest) SetConversationId(v string) *GetAppWorkspaceDirectoryRequest {
	s.ConversationId = &v
	return s
}

func (s *GetAppWorkspaceDirectoryRequest) SetDeep(v int32) *GetAppWorkspaceDirectoryRequest {
	s.Deep = &v
	return s
}

func (s *GetAppWorkspaceDirectoryRequest) SetFilePath(v string) *GetAppWorkspaceDirectoryRequest {
	s.FilePath = &v
	return s
}

func (s *GetAppWorkspaceDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
