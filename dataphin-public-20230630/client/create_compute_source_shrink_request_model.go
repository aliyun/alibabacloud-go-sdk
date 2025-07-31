// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeSourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateComputeSourceShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateComputeSourceShrinkRequest
	GetOpTenantId() *int64
}

type CreateComputeSourceShrinkRequest struct {
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateComputeSourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeSourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateComputeSourceShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateComputeSourceShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateComputeSourceShrinkRequest) SetCreateCommandShrink(v string) *CreateComputeSourceShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateComputeSourceShrinkRequest) SetOpTenantId(v int64) *CreateComputeSourceShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateComputeSourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
