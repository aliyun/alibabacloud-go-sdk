// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKgRelationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateKgRelationShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateKgRelationShrinkRequest
	GetUpdateCommandShrink() *string
	SetWorkspaceId(v string) *UpdateKgRelationShrinkRequest
	GetWorkspaceId() *string
}

type UpdateKgRelationShrinkRequest struct {
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

func (s UpdateKgRelationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKgRelationShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateKgRelationShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateKgRelationShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateKgRelationShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateKgRelationShrinkRequest) SetOpTenantId(v int64) *UpdateKgRelationShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateKgRelationShrinkRequest) SetUpdateCommandShrink(v string) *UpdateKgRelationShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateKgRelationShrinkRequest) SetWorkspaceId(v string) *UpdateKgRelationShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateKgRelationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
