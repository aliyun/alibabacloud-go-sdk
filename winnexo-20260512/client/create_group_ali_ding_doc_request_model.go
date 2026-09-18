// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupAliDingDocRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateGroupAliDingDocRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateGroupAliDingDocRequest
	GetDirectoryId() *string
	SetFilePublicUrl(v string) *CreateGroupAliDingDocRequest
	GetFilePublicUrl() *string
	SetGroupId(v string) *CreateGroupAliDingDocRequest
	GetGroupId() *string
	SetName(v string) *CreateGroupAliDingDocRequest
	GetName() *string
	SetSourceTags(v string) *CreateGroupAliDingDocRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreateGroupAliDingDocRequest
	GetTenantId() *string
}

type CreateGroupAliDingDocRequest struct {
	// The description of the AI assistant.
	//
	// example:
	//
	// Sample description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The publicly accessible URL of the Alibaba DingTalk online document.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/winnexo/resource
	FilePublicUrl *string `json:"filePublicUrl,omitempty" xml:"filePublicUrl,omitempty"`
	// The project group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleGroupId
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The image name.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_paswd_104
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The resource tags. This parameter is optional. Specify a JSON string list, such as ["tagA","tagB"].
	//
	// example:
	//
	// ["Customer","GroupChat"]
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateGroupAliDingDocRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupAliDingDocRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupAliDingDocRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupAliDingDocRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupAliDingDocRequest) GetFilePublicUrl() *string {
	return s.FilePublicUrl
}

func (s *CreateGroupAliDingDocRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupAliDingDocRequest) GetName() *string {
	return s.Name
}

func (s *CreateGroupAliDingDocRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateGroupAliDingDocRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateGroupAliDingDocRequest) SetDescription(v string) *CreateGroupAliDingDocRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupAliDingDocRequest) SetDirectoryId(v string) *CreateGroupAliDingDocRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupAliDingDocRequest) SetFilePublicUrl(v string) *CreateGroupAliDingDocRequest {
	s.FilePublicUrl = &v
	return s
}

func (s *CreateGroupAliDingDocRequest) SetGroupId(v string) *CreateGroupAliDingDocRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupAliDingDocRequest) SetName(v string) *CreateGroupAliDingDocRequest {
	s.Name = &v
	return s
}

func (s *CreateGroupAliDingDocRequest) SetSourceTags(v string) *CreateGroupAliDingDocRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateGroupAliDingDocRequest) SetTenantId(v string) *CreateGroupAliDingDocRequest {
	s.TenantId = &v
	return s
}

func (s *CreateGroupAliDingDocRequest) Validate() error {
	return dara.Validate(s)
}
