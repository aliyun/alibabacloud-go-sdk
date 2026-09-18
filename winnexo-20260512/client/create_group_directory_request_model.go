// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateGroupDirectoryRequest
	GetDescription() *string
	SetGroupId(v string) *CreateGroupDirectoryRequest
	GetGroupId() *string
	SetName(v string) *CreateGroupDirectoryRequest
	GetName() *string
	SetParentDirectoryId(v string) *CreateGroupDirectoryRequest
	GetParentDirectoryId() *string
	SetTenantId(v string) *CreateGroupDirectoryRequest
	GetTenantId() *string
}

type CreateGroupDirectoryRequest struct {
	// The workspace description.
	//
	// example:
	//
	// ProjectDescription
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The project group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The updated name of the filter view.
	//
	// This parameter is required.
	//
	// example:
	//
	// ProjectFiles
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// dir_parent
	ParentDirectoryId *string `json:"parentDirectoryId,omitempty" xml:"parentDirectoryId,omitempty"`
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateGroupDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupDirectoryRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupDirectoryRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupDirectoryRequest) GetName() *string {
	return s.Name
}

func (s *CreateGroupDirectoryRequest) GetParentDirectoryId() *string {
	return s.ParentDirectoryId
}

func (s *CreateGroupDirectoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateGroupDirectoryRequest) SetDescription(v string) *CreateGroupDirectoryRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupDirectoryRequest) SetGroupId(v string) *CreateGroupDirectoryRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupDirectoryRequest) SetName(v string) *CreateGroupDirectoryRequest {
	s.Name = &v
	return s
}

func (s *CreateGroupDirectoryRequest) SetParentDirectoryId(v string) *CreateGroupDirectoryRequest {
	s.ParentDirectoryId = &v
	return s
}

func (s *CreateGroupDirectoryRequest) SetTenantId(v string) *CreateGroupDirectoryRequest {
	s.TenantId = &v
	return s
}

func (s *CreateGroupDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
