// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKgEntityShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateKgEntityShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateKgEntityShrinkRequest
	GetUpdateCommandShrink() *string
	SetWorkspaceId(v string) *UpdateKgEntityShrinkRequest
	GetWorkspaceId() *string
}

type UpdateKgEntityShrinkRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The update command.
	//
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f1d4559a4db044158305e2d89bccf81f
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateKgEntityShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKgEntityShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateKgEntityShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateKgEntityShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateKgEntityShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateKgEntityShrinkRequest) SetOpTenantId(v int64) *UpdateKgEntityShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateKgEntityShrinkRequest) SetUpdateCommandShrink(v string) *UpdateKgEntityShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateKgEntityShrinkRequest) SetWorkspaceId(v string) *UpdateKgEntityShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateKgEntityShrinkRequest) Validate() error {
	return dara.Validate(s)
}
