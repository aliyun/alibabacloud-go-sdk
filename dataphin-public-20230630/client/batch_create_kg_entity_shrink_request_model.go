// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateKgEntityShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *BatchCreateKgEntityShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *BatchCreateKgEntityShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *BatchCreateKgEntityShrinkRequest
	GetOpUserId() *string
	SetWorkspaceId(v string) *BatchCreateKgEntityShrinkRequest
	GetWorkspaceId() *string
}

type BatchCreateKgEntityShrinkRequest struct {
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
	// The ID of the operator user.
	//
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

func (s BatchCreateKgEntityShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateKgEntityShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateKgEntityShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *BatchCreateKgEntityShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *BatchCreateKgEntityShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *BatchCreateKgEntityShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *BatchCreateKgEntityShrinkRequest) SetCreateCommandShrink(v string) *BatchCreateKgEntityShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *BatchCreateKgEntityShrinkRequest) SetOpTenantId(v int64) *BatchCreateKgEntityShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *BatchCreateKgEntityShrinkRequest) SetOpUserId(v string) *BatchCreateKgEntityShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *BatchCreateKgEntityShrinkRequest) SetWorkspaceId(v string) *BatchCreateKgEntityShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *BatchCreateKgEntityShrinkRequest) Validate() error {
	return dara.Validate(s)
}
