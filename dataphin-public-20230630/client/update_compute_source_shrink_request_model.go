// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeSourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateComputeSourceShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateComputeSourceShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateComputeSourceShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateComputeSourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeSourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateComputeSourceShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateComputeSourceShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateComputeSourceShrinkRequest) SetOpTenantId(v int64) *UpdateComputeSourceShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateComputeSourceShrinkRequest) SetUpdateCommandShrink(v string) *UpdateComputeSourceShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateComputeSourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
