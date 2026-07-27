// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKgRelationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateKgRelationShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateKgRelationShrinkRequest
	GetOpTenantId() *int64
	SetWorkspaceId(v string) *CreateKgRelationShrinkRequest
	GetWorkspaceId() *string
}

type CreateKgRelationShrinkRequest struct {
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
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f1d4559a4db044158305e2d89bccf81f
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateKgRelationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKgRelationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateKgRelationShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateKgRelationShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateKgRelationShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateKgRelationShrinkRequest) SetCreateCommandShrink(v string) *CreateKgRelationShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateKgRelationShrinkRequest) SetOpTenantId(v int64) *CreateKgRelationShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateKgRelationShrinkRequest) SetWorkspaceId(v string) *CreateKgRelationShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateKgRelationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
