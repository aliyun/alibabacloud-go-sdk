// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateGroupTextRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateGroupTextRequest
	GetDirectoryId() *string
	SetGroupId(v string) *CreateGroupTextRequest
	GetGroupId() *string
	SetName(v string) *CreateGroupTextRequest
	GetName() *string
	SetSourceTags(v string) *CreateGroupTextRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreateGroupTextRequest
	GetTenantId() *string
	SetTextContent(v string) *CreateGroupTextRequest
	GetTextContent() *string
}

type CreateGroupTextRequest struct {
	// The description of the AI assistant.
	//
	// example:
	//
	// example
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// dir_example
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The project group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The image name.
	//
	// This parameter is required.
	//
	// example:
	//
	// ProjectResources
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The source tags.
	//
	// example:
	//
	// example
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The message content for text messages.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	TextContent *string `json:"textContent,omitempty" xml:"textContent,omitempty"`
}

func (s CreateGroupTextRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupTextRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupTextRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupTextRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupTextRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupTextRequest) GetName() *string {
	return s.Name
}

func (s *CreateGroupTextRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateGroupTextRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateGroupTextRequest) GetTextContent() *string {
	return s.TextContent
}

func (s *CreateGroupTextRequest) SetDescription(v string) *CreateGroupTextRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupTextRequest) SetDirectoryId(v string) *CreateGroupTextRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupTextRequest) SetGroupId(v string) *CreateGroupTextRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupTextRequest) SetName(v string) *CreateGroupTextRequest {
	s.Name = &v
	return s
}

func (s *CreateGroupTextRequest) SetSourceTags(v string) *CreateGroupTextRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateGroupTextRequest) SetTenantId(v string) *CreateGroupTextRequest {
	s.TenantId = &v
	return s
}

func (s *CreateGroupTextRequest) SetTextContent(v string) *CreateGroupTextRequest {
	s.TextContent = &v
	return s
}

func (s *CreateGroupTextRequest) Validate() error {
	return dara.Validate(s)
}
