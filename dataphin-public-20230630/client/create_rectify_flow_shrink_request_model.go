// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRectifyFlowShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateRectifyFlowShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateRectifyFlowShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CreateRectifyFlowShrinkRequest
	GetOpUserId() *string
}

type CreateRectifyFlowShrinkRequest struct {
	// The business request parameters.
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
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s CreateRectifyFlowShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRectifyFlowShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRectifyFlowShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateRectifyFlowShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateRectifyFlowShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CreateRectifyFlowShrinkRequest) SetCreateCommandShrink(v string) *CreateRectifyFlowShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateRectifyFlowShrinkRequest) SetOpTenantId(v int64) *CreateRectifyFlowShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateRectifyFlowShrinkRequest) SetOpUserId(v string) *CreateRectifyFlowShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateRectifyFlowShrinkRequest) Validate() error {
	return dara.Validate(s)
}
