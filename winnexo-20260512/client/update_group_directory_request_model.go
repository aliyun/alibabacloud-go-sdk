// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateGroupDirectoryRequest
	GetDescription() *string
	SetDirectoryId(v string) *UpdateGroupDirectoryRequest
	GetDirectoryId() *string
	SetGroupId(v string) *UpdateGroupDirectoryRequest
	GetGroupId() *string
	SetName(v string) *UpdateGroupDirectoryRequest
	GetName() *string
	SetTenantId(v string) *UpdateGroupDirectoryRequest
	GetTenantId() *string
}

type UpdateGroupDirectoryRequest struct {
	// The new description. If this parameter is set to an empty string, the description is cleared. If this parameter is omitted or set to null, the description remains unchanged. At least one of name or description must be non-null.
	//
	// example:
	//
	// Project description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ID of the physical subfolder in the current space. The internal root folder and reference folders are not allowed.
	//
	// This parameter is required.
	//
	// example:
	//
	// dir_example
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The ID of the collaborative share.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The new name. If this parameter is omitted or set to null, the name remains unchanged.
	//
	// example:
	//
	// Project Materials
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateGroupDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupDirectoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupDirectoryRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateGroupDirectoryRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateGroupDirectoryRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateGroupDirectoryRequest) GetName() *string {
	return s.Name
}

func (s *UpdateGroupDirectoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateGroupDirectoryRequest) SetDescription(v string) *UpdateGroupDirectoryRequest {
	s.Description = &v
	return s
}

func (s *UpdateGroupDirectoryRequest) SetDirectoryId(v string) *UpdateGroupDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateGroupDirectoryRequest) SetGroupId(v string) *UpdateGroupDirectoryRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateGroupDirectoryRequest) SetName(v string) *UpdateGroupDirectoryRequest {
	s.Name = &v
	return s
}

func (s *UpdateGroupDirectoryRequest) SetTenantId(v string) *UpdateGroupDirectoryRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateGroupDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
