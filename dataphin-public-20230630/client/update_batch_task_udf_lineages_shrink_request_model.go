// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBatchTaskUdfLineagesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateBatchTaskUdfLineagesShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateBatchTaskUdfLineagesShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateBatchTaskUdfLineagesShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateBatchTaskUdfLineagesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTaskUdfLineagesShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateBatchTaskUdfLineagesShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateBatchTaskUdfLineagesShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateBatchTaskUdfLineagesShrinkRequest) SetOpTenantId(v int64) *UpdateBatchTaskUdfLineagesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateBatchTaskUdfLineagesShrinkRequest) SetUpdateCommandShrink(v string) *UpdateBatchTaskUdfLineagesShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateBatchTaskUdfLineagesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
