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
	SetAutoUpdate(v bool) *CreateMemoryRequest
	GetAutoUpdate() *bool
	SetContent(v string) *CreateMemoryRequest
	GetContent() *string
	SetExpirationTime(v int32) *CreateMemoryRequest
	GetExpirationTime() *int32
	SetMessagesJson(v string) *CreateMemoryRequest
	GetMessagesJson() *string
	SetMetaData(v map[string]*string) *CreateMemoryRequest
	GetMetaData() map[string]*string
	SetProjectId(v string) *CreateMemoryRequest
	GetProjectId() *string
	SetPrompt(v string) *CreateMemoryRequest
	GetPrompt() *string
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
	AppId          *string            `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AutoUpdate     *bool              `json:"AutoUpdate,omitempty" xml:"AutoUpdate,omitempty"`
	Content        *string            `json:"Content,omitempty" xml:"Content,omitempty"`
	ExpirationTime *int32             `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	MessagesJson   *string            `json:"MessagesJson,omitempty" xml:"MessagesJson,omitempty"`
	MetaData       map[string]*string `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
	// example:
	//
	// profile_project
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Prompt    *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
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

func (s *CreateMemoryRequest) GetAutoUpdate() *bool {
	return s.AutoUpdate
}

func (s *CreateMemoryRequest) GetContent() *string {
	return s.Content
}

func (s *CreateMemoryRequest) GetExpirationTime() *int32 {
	return s.ExpirationTime
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

func (s *CreateMemoryRequest) GetPrompt() *string {
	return s.Prompt
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

func (s *CreateMemoryRequest) SetAutoUpdate(v bool) *CreateMemoryRequest {
	s.AutoUpdate = &v
	return s
}

func (s *CreateMemoryRequest) SetContent(v string) *CreateMemoryRequest {
	s.Content = &v
	return s
}

func (s *CreateMemoryRequest) SetExpirationTime(v int32) *CreateMemoryRequest {
	s.ExpirationTime = &v
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

func (s *CreateMemoryRequest) SetPrompt(v string) *CreateMemoryRequest {
	s.Prompt = &v
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
