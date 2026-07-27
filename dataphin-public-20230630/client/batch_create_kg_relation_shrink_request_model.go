// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateKgRelationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *BatchCreateKgRelationShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *BatchCreateKgRelationShrinkRequest
	GetOpTenantId() *int64
	SetWorkspaceId(v string) *BatchCreateKgRelationShrinkRequest
	GetWorkspaceId() *string
}

type BatchCreateKgRelationShrinkRequest struct {
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

func (s BatchCreateKgRelationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateKgRelationShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateKgRelationShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *BatchCreateKgRelationShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *BatchCreateKgRelationShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *BatchCreateKgRelationShrinkRequest) SetCreateCommandShrink(v string) *BatchCreateKgRelationShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *BatchCreateKgRelationShrinkRequest) SetOpTenantId(v int64) *BatchCreateKgRelationShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *BatchCreateKgRelationShrinkRequest) SetWorkspaceId(v string) *BatchCreateKgRelationShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *BatchCreateKgRelationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
