// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateMemoryShrinkRequest
	GetAppId() *string
	SetContent(v string) *CreateMemoryShrinkRequest
	GetContent() *string
	SetMessagesJson(v string) *CreateMemoryShrinkRequest
	GetMessagesJson() *string
	SetMetaDataShrink(v string) *CreateMemoryShrinkRequest
	GetMetaDataShrink() *string
	SetProjectId(v string) *CreateMemoryShrinkRequest
	GetProjectId() *string
	SetUserDefinedId(v string) *CreateMemoryShrinkRequest
	GetUserDefinedId() *string
	SetWorkspaceId(v string) *CreateMemoryShrinkRequest
	GetWorkspaceId() *string
}

type CreateMemoryShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_bfaf7e110b6d4359977d1686a3f8
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Content        *string `json:"Content,omitempty" xml:"Content,omitempty"`
	MessagesJson   *string `json:"MessagesJson,omitempty" xml:"MessagesJson,omitempty"`
	MetaDataShrink *string `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
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

func (s CreateMemoryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMemoryShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateMemoryShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *CreateMemoryShrinkRequest) GetMessagesJson() *string {
	return s.MessagesJson
}

func (s *CreateMemoryShrinkRequest) GetMetaDataShrink() *string {
	return s.MetaDataShrink
}

func (s *CreateMemoryShrinkRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateMemoryShrinkRequest) GetUserDefinedId() *string {
	return s.UserDefinedId
}

func (s *CreateMemoryShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateMemoryShrinkRequest) SetAppId(v string) *CreateMemoryShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateMemoryShrinkRequest) SetContent(v string) *CreateMemoryShrinkRequest {
	s.Content = &v
	return s
}

func (s *CreateMemoryShrinkRequest) SetMessagesJson(v string) *CreateMemoryShrinkRequest {
	s.MessagesJson = &v
	return s
}

func (s *CreateMemoryShrinkRequest) SetMetaDataShrink(v string) *CreateMemoryShrinkRequest {
	s.MetaDataShrink = &v
	return s
}

func (s *CreateMemoryShrinkRequest) SetProjectId(v string) *CreateMemoryShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateMemoryShrinkRequest) SetUserDefinedId(v string) *CreateMemoryShrinkRequest {
	s.UserDefinedId = &v
	return s
}

func (s *CreateMemoryShrinkRequest) SetWorkspaceId(v string) *CreateMemoryShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateMemoryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
