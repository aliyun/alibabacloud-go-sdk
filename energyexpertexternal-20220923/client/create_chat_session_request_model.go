// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFolderId(v string) *CreateChatSessionRequest
	GetFolderId() *string
	SetName(v string) *CreateChatSessionRequest
	GetName() *string
	SetUserId(v string) *CreateChatSessionRequest
	GetUserId() *string
}

type CreateChatSessionRequest struct {
	// Folder ID, to search for Q&A content within the current folder and its subfolders.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1a851c4a-1d65-11ef-99a7-ssfsfdd
	FolderId *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
	// Name of the current session.
	//
	// This parameter is required.
	//
	// example:
	//
	// analyzer_1744684195
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The unique identifier of the user. If not provided, the SDK calling account will be used as the user ID by default.
	//
	// example:
	//
	// 1233333
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateChatSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateChatSessionRequest) GoString() string {
	return s.String()
}

func (s *CreateChatSessionRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *CreateChatSessionRequest) GetName() *string {
	return s.Name
}

func (s *CreateChatSessionRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateChatSessionRequest) SetFolderId(v string) *CreateChatSessionRequest {
	s.FolderId = &v
	return s
}

func (s *CreateChatSessionRequest) SetName(v string) *CreateChatSessionRequest {
	s.Name = &v
	return s
}

func (s *CreateChatSessionRequest) SetUserId(v string) *CreateChatSessionRequest {
	s.UserId = &v
	return s
}

func (s *CreateChatSessionRequest) Validate() error {
	return dara.Validate(s)
}
