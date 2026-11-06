// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceBindingsOfNumberShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdsShrink(v string) *UpdateInstanceBindingsOfNumberShrinkRequest
	GetInstanceIdsShrink() *string
	SetNumber(v string) *UpdateInstanceBindingsOfNumberShrinkRequest
	GetNumber() *string
	SetTenantId(v string) *UpdateInstanceBindingsOfNumberShrinkRequest
	GetTenantId() *string
}

type UpdateInstanceBindingsOfNumberShrinkRequest struct {
	// The list of instance IDs.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The bound number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 057100001234
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 1234567890
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s UpdateInstanceBindingsOfNumberShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceBindingsOfNumberShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceBindingsOfNumberShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *UpdateInstanceBindingsOfNumberShrinkRequest) GetNumber() *string {
	return s.Number
}

func (s *UpdateInstanceBindingsOfNumberShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateInstanceBindingsOfNumberShrinkRequest) SetInstanceIdsShrink(v string) *UpdateInstanceBindingsOfNumberShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *UpdateInstanceBindingsOfNumberShrinkRequest) SetNumber(v string) *UpdateInstanceBindingsOfNumberShrinkRequest {
	s.Number = &v
	return s
}

func (s *UpdateInstanceBindingsOfNumberShrinkRequest) SetTenantId(v string) *UpdateInstanceBindingsOfNumberShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateInstanceBindingsOfNumberShrinkRequest) Validate() error {
	return dara.Validate(s)
}
