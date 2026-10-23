// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryAutoExtractionSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClearExisting(v bool) *UpdateMemoryAutoExtractionSwitchRequest
	GetClearExisting() *bool
	SetEnabled(v bool) *UpdateMemoryAutoExtractionSwitchRequest
	GetEnabled() *bool
	SetGroupId(v string) *UpdateMemoryAutoExtractionSwitchRequest
	GetGroupId() *string
	SetOwnerType(v string) *UpdateMemoryAutoExtractionSwitchRequest
	GetOwnerType() *string
	SetTenantId(v string) *UpdateMemoryAutoExtractionSwitchRequest
	GetTenantId() *string
}

type UpdateMemoryAutoExtractionSwitchRequest struct {
	// Specifies whether to synchronously clear the automatically collected memories. This parameter can be used regardless of whether the feature is enabled or disabled.
	//
	// example:
	//
	// true
	ClearExisting *bool `json:"clearExisting,omitempty" xml:"clearExisting,omitempty"`
	// Specifies whether to enable automatic memory collection.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The collaborative workspace ID. This parameter is required when ownerType is set to group.
	//
	// example:
	//
	// cg_12345
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The owner type. A value of user indicates personal memories of the authenticated user. A value of group indicates memories of a specified collaborative workspace.
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

func (s UpdateMemoryAutoExtractionSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryAutoExtractionSwitchRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) GetClearExisting() *bool {
	return s.ClearExisting
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) GetOwnerType() *string {
	return s.OwnerType
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) SetClearExisting(v bool) *UpdateMemoryAutoExtractionSwitchRequest {
	s.ClearExisting = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) SetEnabled(v bool) *UpdateMemoryAutoExtractionSwitchRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) SetGroupId(v string) *UpdateMemoryAutoExtractionSwitchRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) SetOwnerType(v string) *UpdateMemoryAutoExtractionSwitchRequest {
	s.OwnerType = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) SetTenantId(v string) *UpdateMemoryAutoExtractionSwitchRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) Validate() error {
	return dara.Validate(s)
}
