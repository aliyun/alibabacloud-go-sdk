// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBatchTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateBatchTaskShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpdateBatchTaskShrinkRequest
	GetOpUserId() *string
	SetUpdateCommandShrink(v string) *UpdateBatchTaskShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateBatchTaskShrinkRequest struct {
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
	// The update request.
	//
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateBatchTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateBatchTaskShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateBatchTaskShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpdateBatchTaskShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateBatchTaskShrinkRequest) SetOpTenantId(v int64) *UpdateBatchTaskShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateBatchTaskShrinkRequest) SetOpUserId(v string) *UpdateBatchTaskShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateBatchTaskShrinkRequest) SetUpdateCommandShrink(v string) *UpdateBatchTaskShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateBatchTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
