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
	SetUpdateCommandShrink(v string) *UpdateBatchTaskShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateBatchTaskShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
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

func (s *UpdateBatchTaskShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateBatchTaskShrinkRequest) SetOpTenantId(v int64) *UpdateBatchTaskShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateBatchTaskShrinkRequest) SetUpdateCommandShrink(v string) *UpdateBatchTaskShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateBatchTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
