// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateBatchTaskShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateBatchTaskShrinkRequest
	GetOpTenantId() *int64
}

type CreateBatchTaskShrinkRequest struct {
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateBatchTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateBatchTaskShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateBatchTaskShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateBatchTaskShrinkRequest) SetCreateCommandShrink(v string) *CreateBatchTaskShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateBatchTaskShrinkRequest) SetOpTenantId(v int64) *CreateBatchTaskShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateBatchTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
