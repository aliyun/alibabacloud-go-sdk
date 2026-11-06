// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNumberBindingsOfInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateNumberBindingsOfInstanceShrinkRequest
	GetInstanceId() *string
	SetNumbersShrink(v string) *UpdateNumberBindingsOfInstanceShrinkRequest
	GetNumbersShrink() *string
	SetTenantId(v string) *UpdateNumberBindingsOfInstanceShrinkRequest
	GetTenantId() *string
}

type UpdateNumberBindingsOfInstanceShrinkRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4eee9bf8-1319-468f-ac82-83c50ae389f8
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of phone numbers.
	NumbersShrink *string `json:"Numbers,omitempty" xml:"Numbers,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 1234567890
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s UpdateNumberBindingsOfInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNumberBindingsOfInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateNumberBindingsOfInstanceShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateNumberBindingsOfInstanceShrinkRequest) GetNumbersShrink() *string {
	return s.NumbersShrink
}

func (s *UpdateNumberBindingsOfInstanceShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateNumberBindingsOfInstanceShrinkRequest) SetInstanceId(v string) *UpdateNumberBindingsOfInstanceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateNumberBindingsOfInstanceShrinkRequest) SetNumbersShrink(v string) *UpdateNumberBindingsOfInstanceShrinkRequest {
	s.NumbersShrink = &v
	return s
}

func (s *UpdateNumberBindingsOfInstanceShrinkRequest) SetTenantId(v string) *UpdateNumberBindingsOfInstanceShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateNumberBindingsOfInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
