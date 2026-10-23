// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *GetMemoryDetailRequest
	GetGroupId() *string
	SetOwnerType(v string) *GetMemoryDetailRequest
	GetOwnerType() *string
	SetTenantId(v string) *GetMemoryDetailRequest
	GetTenantId() *string
}

type GetMemoryDetailRequest struct {
	// The workspace ID. This parameter is required when ownerType is set to group.
	//
	// example:
	//
	// cg_12345
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The owner type. A value of user indicates the personal memory of the authenticated user. A value of group indicates the memory of a specified workspace.
	//
	// example:
	//
	// group
	OwnerType *string `json:"ownerType,omitempty" xml:"ownerType,omitempty"`
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10001
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetMemoryDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryDetailRequest) GoString() string {
	return s.String()
}

func (s *GetMemoryDetailRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetMemoryDetailRequest) GetOwnerType() *string {
	return s.OwnerType
}

func (s *GetMemoryDetailRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetMemoryDetailRequest) SetGroupId(v string) *GetMemoryDetailRequest {
	s.GroupId = &v
	return s
}

func (s *GetMemoryDetailRequest) SetOwnerType(v string) *GetMemoryDetailRequest {
	s.OwnerType = &v
	return s
}

func (s *GetMemoryDetailRequest) SetTenantId(v string) *GetMemoryDetailRequest {
	s.TenantId = &v
	return s
}

func (s *GetMemoryDetailRequest) Validate() error {
	return dara.Validate(s)
}
