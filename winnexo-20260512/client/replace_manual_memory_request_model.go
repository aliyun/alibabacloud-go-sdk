// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceManualMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *ReplaceManualMemoryRequest
	GetContent() *string
	SetGroupId(v string) *ReplaceManualMemoryRequest
	GetGroupId() *string
	SetOwnerType(v string) *ReplaceManualMemoryRequest
	GetOwnerType() *string
	SetTenantId(v string) *ReplaceManualMemoryRequest
	GetTenantId() *string
}

type ReplaceManualMemoryRequest struct {
	// The complete content of profile.md. An empty string indicates that the memory is cleared. This field is required, but an empty value is allowed.
	//
	// This parameter is required.
	//
	// example:
	//
	// Prefers concise answers
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
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

func (s ReplaceManualMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceManualMemoryRequest) GoString() string {
	return s.String()
}

func (s *ReplaceManualMemoryRequest) GetContent() *string {
	return s.Content
}

func (s *ReplaceManualMemoryRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ReplaceManualMemoryRequest) GetOwnerType() *string {
	return s.OwnerType
}

func (s *ReplaceManualMemoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ReplaceManualMemoryRequest) SetContent(v string) *ReplaceManualMemoryRequest {
	s.Content = &v
	return s
}

func (s *ReplaceManualMemoryRequest) SetGroupId(v string) *ReplaceManualMemoryRequest {
	s.GroupId = &v
	return s
}

func (s *ReplaceManualMemoryRequest) SetOwnerType(v string) *ReplaceManualMemoryRequest {
	s.OwnerType = &v
	return s
}

func (s *ReplaceManualMemoryRequest) SetTenantId(v string) *ReplaceManualMemoryRequest {
	s.TenantId = &v
	return s
}

func (s *ReplaceManualMemoryRequest) Validate() error {
	return dara.Validate(s)
}
