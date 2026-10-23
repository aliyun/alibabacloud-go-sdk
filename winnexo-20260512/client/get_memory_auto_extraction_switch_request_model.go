// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryAutoExtractionSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *GetMemoryAutoExtractionSwitchRequest
	GetGroupId() *string
	SetOwnerType(v string) *GetMemoryAutoExtractionSwitchRequest
	GetOwnerType() *string
	SetTenantId(v string) *GetMemoryAutoExtractionSwitchRequest
	GetTenantId() *string
}

type GetMemoryAutoExtractionSwitchRequest struct {
	// The collaboration workspace ID. This parameter is required when ownerType is set to group.
	//
	// example:
	//
	// cg_12345
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The owner type. A value of user indicates the personal memory of the authenticated user, and a value of group indicates the memory of the specified collaboration workspace.
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
}

func (s GetMemoryAutoExtractionSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryAutoExtractionSwitchRequest) GoString() string {
	return s.String()
}

func (s *GetMemoryAutoExtractionSwitchRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetMemoryAutoExtractionSwitchRequest) GetOwnerType() *string {
	return s.OwnerType
}

func (s *GetMemoryAutoExtractionSwitchRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetMemoryAutoExtractionSwitchRequest) SetGroupId(v string) *GetMemoryAutoExtractionSwitchRequest {
	s.GroupId = &v
	return s
}

func (s *GetMemoryAutoExtractionSwitchRequest) SetOwnerType(v string) *GetMemoryAutoExtractionSwitchRequest {
	s.OwnerType = &v
	return s
}

func (s *GetMemoryAutoExtractionSwitchRequest) SetTenantId(v string) *GetMemoryAutoExtractionSwitchRequest {
	s.TenantId = &v
	return s
}

func (s *GetMemoryAutoExtractionSwitchRequest) Validate() error {
	return dara.Validate(s)
}
