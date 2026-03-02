// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateMemoryRequest
	GetAppId() *string
	SetContent(v string) *CreateMemoryRequest
	GetContent() *string
	SetMessagesJson(v string) *CreateMemoryRequest
	GetMessagesJson() *string
	SetMetaData(v map[string]*string) *CreateMemoryRequest
	GetMetaData() map[string]*string
	SetProjectId(v string) *CreateMemoryRequest
	GetProjectId() *string
	SetUserDefinedId(v string) *CreateMemoryRequest
	GetUserDefinedId() *string
	SetWorkspaceId(v string) *CreateMemoryRequest
	GetWorkspaceId() *string
}

type CreateMemoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_bfaf7e110b6d4359977d1686a3f8
	AppId        *string            `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Content      *string            `json:"Content,omitempty" xml:"Content,omitempty"`
	MessagesJson *string            `json:"MessagesJson,omitempty" xml:"MessagesJson,omitempty"`
	MetaData     map[string]*string `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
	// example:
	//
	// profile_project
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CD51C0ED-4666-13DA-BC7D-C0D21CBE16C9
	UserDefinedId *string `json:"UserDefinedId,omitempty" xml:"UserDefinedId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-jb5sabg80b4ts71g
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryRequest) GoString() string {
	return s.String()
}

func (s *CreateMemoryRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateMemoryRequest) GetContent() *string {
	return s.Content
}

func (s *CreateMemoryRequest) GetMessagesJson() *string {
	return s.MessagesJson
}

func (s *CreateMemoryRequest) GetMetaData() map[string]*string {
	return s.MetaData
}

func (s *CreateMemoryRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateMemoryRequest) GetUserDefinedId() *string {
	return s.UserDefinedId
}

func (s *CreateMemoryRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateMemoryRequest) SetAppId(v string) *CreateMemoryRequest {
	s.AppId = &v
	return s
}

func (s *CreateMemoryRequest) SetContent(v string) *CreateMemoryRequest {
	s.Content = &v
	return s
}

func (s *CreateMemoryRequest) SetMessagesJson(v string) *CreateMemoryRequest {
	s.MessagesJson = &v
	return s
}

func (s *CreateMemoryRequest) SetMetaData(v map[string]*string) *CreateMemoryRequest {
	s.MetaData = v
	return s
}

func (s *CreateMemoryRequest) SetProjectId(v string) *CreateMemoryRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateMemoryRequest) SetUserDefinedId(v string) *CreateMemoryRequest {
	s.UserDefinedId = &v
	return s
}

func (s *CreateMemoryRequest) SetWorkspaceId(v string) *CreateMemoryRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateMemoryRequest) Validate() error {
	return dara.Validate(s)
}
