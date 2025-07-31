// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBatchTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommandShrink(v string) *DeleteBatchTaskShrinkRequest
	GetDeleteCommandShrink() *string
	SetOpTenantId(v int64) *DeleteBatchTaskShrinkRequest
	GetOpTenantId() *int64
}

type DeleteBatchTaskShrinkRequest struct {
	// This parameter is required.
	DeleteCommandShrink *string `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteBatchTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBatchTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteBatchTaskShrinkRequest) GetDeleteCommandShrink() *string {
	return s.DeleteCommandShrink
}

func (s *DeleteBatchTaskShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteBatchTaskShrinkRequest) SetDeleteCommandShrink(v string) *DeleteBatchTaskShrinkRequest {
	s.DeleteCommandShrink = &v
	return s
}

func (s *DeleteBatchTaskShrinkRequest) SetOpTenantId(v int64) *DeleteBatchTaskShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteBatchTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
