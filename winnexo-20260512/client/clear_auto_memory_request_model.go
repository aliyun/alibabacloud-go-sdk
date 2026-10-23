// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearAutoMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ClearAutoMemoryRequest
	GetGroupId() *string
	SetOwnerType(v string) *ClearAutoMemoryRequest
	GetOwnerType() *string
	SetTenantId(v string) *ClearAutoMemoryRequest
	GetTenantId() *string
	SetType(v string) *ClearAutoMemoryRequest
	GetType() *string
}

type ClearAutoMemoryRequest struct {
	// The workspace ID. This parameter is required when ownerType is set to group.
	//
	// example:
	//
	// cg_12345
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The owner type. Valid values:
	//
	// - user: the personal memory of the authenticated user.
	//
	// - group: the memory of a specified workspace.
	//
	// example:
	//
	// group
	OwnerType *string `json:"ownerType,omitempty" xml:"ownerType,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10001
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The memory type. Valid values: user, feedback, reference, project, and all. A value of all does not clear manual memory.
	//
	// This parameter is required.
	//
	// example:
	//
	// project
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ClearAutoMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ClearAutoMemoryRequest) GoString() string {
	return s.String()
}

func (s *ClearAutoMemoryRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ClearAutoMemoryRequest) GetOwnerType() *string {
	return s.OwnerType
}

func (s *ClearAutoMemoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ClearAutoMemoryRequest) GetType() *string {
	return s.Type
}

func (s *ClearAutoMemoryRequest) SetGroupId(v string) *ClearAutoMemoryRequest {
	s.GroupId = &v
	return s
}

func (s *ClearAutoMemoryRequest) SetOwnerType(v string) *ClearAutoMemoryRequest {
	s.OwnerType = &v
	return s
}

func (s *ClearAutoMemoryRequest) SetTenantId(v string) *ClearAutoMemoryRequest {
	s.TenantId = &v
	return s
}

func (s *ClearAutoMemoryRequest) SetType(v string) *ClearAutoMemoryRequest {
	s.Type = &v
	return s
}

func (s *ClearAutoMemoryRequest) Validate() error {
	return dara.Validate(s)
}
