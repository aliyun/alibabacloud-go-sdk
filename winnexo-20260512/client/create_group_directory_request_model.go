// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreateGroupDirectoryRequest
	GetAppToken() *string
	SetDescription(v string) *CreateGroupDirectoryRequest
	GetDescription() *string
	SetGroupId(v string) *CreateGroupDirectoryRequest
	GetGroupId() *string
	SetName(v string) *CreateGroupDirectoryRequest
	GetName() *string
	SetOpsUserCode(v string) *CreateGroupDirectoryRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateGroupDirectoryRequest
	GetOpsUserId() *string
	SetParentDirectoryId(v string) *CreateGroupDirectoryRequest
	GetParentDirectoryId() *string
	SetTenantId(v string) *CreateGroupDirectoryRequest
	GetTenantId() *string
}

type CreateGroupDirectoryRequest struct {
	// The platform application token issued by the console. If specified, the operation is performed as the user identified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The workspace description.
	//
	// example:
	//
	// Project Description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The project group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The updated name of the filtered view.
	//
	// This parameter is required.
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
	// The directory ID.
	//
	// example:
	//
	// dir_parent
	ParentDirectoryId *string `json:"parentDirectoryId,omitempty" xml:"parentDirectoryId,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
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

func (s *CreateGroupDirectoryRequest) GetAppToken() *string {
	return s.AppToken
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

func (s *CreateGroupDirectoryRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateGroupDirectoryRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreateGroupDirectoryRequest) GetParentDirectoryId() *string {
	return s.ParentDirectoryId
}

func (s *CreateGroupDirectoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateGroupDirectoryRequest) SetAppToken(v string) *CreateGroupDirectoryRequest {
	s.AppToken = &v
	return s
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

func (s *CreateGroupDirectoryRequest) SetOpsUserCode(v string) *CreateGroupDirectoryRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateGroupDirectoryRequest) SetOpsUserId(v string) *CreateGroupDirectoryRequest {
	s.OpsUserId = &v
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
