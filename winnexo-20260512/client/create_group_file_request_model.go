// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateGroupFileRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateGroupFileRequest
	GetDirectoryId() *string
	SetFileRecordId(v string) *CreateGroupFileRequest
	GetFileRecordId() *string
	SetGroupId(v string) *CreateGroupFileRequest
	GetGroupId() *string
	SetName(v string) *CreateGroupFileRequest
	GetName() *string
	SetSourceTags(v string) *CreateGroupFileRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreateGroupFileRequest
	GetTenantId() *string
}

type CreateGroupFileRequest struct {
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
	// The file record ID. This parameter is optional and corresponds to settings.file_record_id.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	FileRecordId *string `json:"fileRecordId,omitempty" xml:"fileRecordId,omitempty"`
	// The project group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Project Files
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The source tags.
	//
	// example:
	//
	// example
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateGroupFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupFileRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupFileRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupFileRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupFileRequest) GetFileRecordId() *string {
	return s.FileRecordId
}

func (s *CreateGroupFileRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupFileRequest) GetName() *string {
	return s.Name
}

func (s *CreateGroupFileRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateGroupFileRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateGroupFileRequest) SetDescription(v string) *CreateGroupFileRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupFileRequest) SetDirectoryId(v string) *CreateGroupFileRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupFileRequest) SetFileRecordId(v string) *CreateGroupFileRequest {
	s.FileRecordId = &v
	return s
}

func (s *CreateGroupFileRequest) SetGroupId(v string) *CreateGroupFileRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupFileRequest) SetName(v string) *CreateGroupFileRequest {
	s.Name = &v
	return s
}

func (s *CreateGroupFileRequest) SetSourceTags(v string) *CreateGroupFileRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateGroupFileRequest) SetTenantId(v string) *CreateGroupFileRequest {
	s.TenantId = &v
	return s
}

func (s *CreateGroupFileRequest) Validate() error {
	return dara.Validate(s)
}
