// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *UpdateGroupDirectoryRequest
	GetAppToken() *string
	SetDescription(v string) *UpdateGroupDirectoryRequest
	GetDescription() *string
	SetDirectoryId(v string) *UpdateGroupDirectoryRequest
	GetDirectoryId() *string
	SetGroupId(v string) *UpdateGroupDirectoryRequest
	GetGroupId() *string
	SetName(v string) *UpdateGroupDirectoryRequest
	GetName() *string
	SetOpsUserCode(v string) *UpdateGroupDirectoryRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *UpdateGroupDirectoryRequest
	GetOpsUserId() *string
	SetTenantId(v string) *UpdateGroupDirectoryRequest
	GetTenantId() *string
}

type UpdateGroupDirectoryRequest struct {
	// The platform application token issued by the console. If specified, the operation is executed as the user identified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The new description. An empty string clears the description. If omitted or set to null, the description is not modified. At least one of the name or description parameters must be non-null.
	//
	// example:
	//
	// Project Description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ID of the physical subdirectory in the current workspace. Internal root directories and referenced directories are not allowed.
	//
	// This parameter is required.
	//
	// example:
	//
	// dir_example
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The ID of the collaboration workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The new name. If omitted or set to null, the name is not modified.
	//
	// example:
	//
	// Project Materials
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The proxy user code, which is unique within the tenant. This parameter is used only when the application credential is configured as source_uid.
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// The proxy user ID, which is the WINNEXO platform user ID. This parameter is used only when the application credential is configured as winnexo_uid.
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The tenant ID. This is a common parameter. If omitted, the default tenant of the caller is used.
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

func (s *UpdateGroupDirectoryRequest) GetAppToken() *string {
	return s.AppToken
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

func (s *UpdateGroupDirectoryRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *UpdateGroupDirectoryRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *UpdateGroupDirectoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateGroupDirectoryRequest) SetAppToken(v string) *UpdateGroupDirectoryRequest {
	s.AppToken = &v
	return s
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

func (s *UpdateGroupDirectoryRequest) SetOpsUserCode(v string) *UpdateGroupDirectoryRequest {
	s.OpsUserCode = &v
	return s
}

func (s *UpdateGroupDirectoryRequest) SetOpsUserId(v string) *UpdateGroupDirectoryRequest {
	s.OpsUserId = &v
	return s
}

func (s *UpdateGroupDirectoryRequest) SetTenantId(v string) *UpdateGroupDirectoryRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateGroupDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
