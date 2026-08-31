// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDepartmentUserShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *SyncDepartmentUserShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *SyncDepartmentUserShrinkRequest
	GetOpUserId() *string
	SetSyncDepartmentUserCommandShrink(v string) *SyncDepartmentUserShrinkRequest
	GetSyncDepartmentUserCommandShrink() *string
}

type SyncDepartmentUserShrinkRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The request command.
	//
	// This parameter is required.
	SyncDepartmentUserCommandShrink *string `json:"SyncDepartmentUserCommand,omitempty" xml:"SyncDepartmentUserCommand,omitempty"`
}

func (s SyncDepartmentUserShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncDepartmentUserShrinkRequest) GoString() string {
	return s.String()
}

func (s *SyncDepartmentUserShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *SyncDepartmentUserShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *SyncDepartmentUserShrinkRequest) GetSyncDepartmentUserCommandShrink() *string {
	return s.SyncDepartmentUserCommandShrink
}

func (s *SyncDepartmentUserShrinkRequest) SetOpTenantId(v int64) *SyncDepartmentUserShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *SyncDepartmentUserShrinkRequest) SetOpUserId(v string) *SyncDepartmentUserShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *SyncDepartmentUserShrinkRequest) SetSyncDepartmentUserCommandShrink(v string) *SyncDepartmentUserShrinkRequest {
	s.SyncDepartmentUserCommandShrink = &v
	return s
}

func (s *SyncDepartmentUserShrinkRequest) Validate() error {
	return dara.Validate(s)
}
