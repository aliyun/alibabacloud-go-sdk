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
	SetOpUserId(v string) *DeleteBatchTaskShrinkRequest
	GetOpUserId() *string
}

type DeleteBatchTaskShrinkRequest struct {
	// The delete request.
	//
	// This parameter is required.
	DeleteCommandShrink *string `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty"`
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

func (s *DeleteBatchTaskShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *DeleteBatchTaskShrinkRequest) SetDeleteCommandShrink(v string) *DeleteBatchTaskShrinkRequest {
	s.DeleteCommandShrink = &v
	return s
}

func (s *DeleteBatchTaskShrinkRequest) SetOpTenantId(v int64) *DeleteBatchTaskShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteBatchTaskShrinkRequest) SetOpUserId(v string) *DeleteBatchTaskShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *DeleteBatchTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
