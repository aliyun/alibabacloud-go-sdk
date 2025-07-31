// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTenantComputeEngineShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateTenantComputeEngineShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateTenantComputeEngineShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateTenantComputeEngineShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateTenantComputeEngineShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTenantComputeEngineShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTenantComputeEngineShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateTenantComputeEngineShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateTenantComputeEngineShrinkRequest) SetOpTenantId(v int64) *UpdateTenantComputeEngineShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateTenantComputeEngineShrinkRequest) SetUpdateCommandShrink(v string) *UpdateTenantComputeEngineShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateTenantComputeEngineShrinkRequest) Validate() error {
	return dara.Validate(s)
}
