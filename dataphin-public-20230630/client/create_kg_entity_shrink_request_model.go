// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKgEntityShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateKgEntityShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateKgEntityShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CreateKgEntityShrinkRequest
	GetOpUserId() *string
	SetWorkspaceId(v string) *CreateKgEntityShrinkRequest
	GetWorkspaceId() *string
}

type CreateKgEntityShrinkRequest struct {
	// The create command.
	//
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
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
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f1d4559a4db044158305e2d89bccf81f
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateKgEntityShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKgEntityShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateKgEntityShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateKgEntityShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateKgEntityShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CreateKgEntityShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateKgEntityShrinkRequest) SetCreateCommandShrink(v string) *CreateKgEntityShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateKgEntityShrinkRequest) SetOpTenantId(v int64) *CreateKgEntityShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateKgEntityShrinkRequest) SetOpUserId(v string) *CreateKgEntityShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateKgEntityShrinkRequest) SetWorkspaceId(v string) *CreateKgEntityShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateKgEntityShrinkRequest) Validate() error {
	return dara.Validate(s)
}
