// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNumbersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteNumbersShrinkRequest
	GetInstanceId() *string
	SetNumbersShrink(v string) *DeleteNumbersShrinkRequest
	GetNumbersShrink() *string
	SetTenantId(v string) *DeleteNumbersShrinkRequest
	GetTenantId() *string
}

type DeleteNumbersShrinkRequest struct {
	// The instance ID.
	//
	// example:
	//
	// e5035654-1745-484a-8c5b-165f7c7bcd79
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of numbers to be deleted.
	NumbersShrink *string `json:"Numbers,omitempty" xml:"Numbers,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 1234567890
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DeleteNumbersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNumbersShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteNumbersShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteNumbersShrinkRequest) GetNumbersShrink() *string {
	return s.NumbersShrink
}

func (s *DeleteNumbersShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteNumbersShrinkRequest) SetInstanceId(v string) *DeleteNumbersShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteNumbersShrinkRequest) SetNumbersShrink(v string) *DeleteNumbersShrinkRequest {
	s.NumbersShrink = &v
	return s
}

func (s *DeleteNumbersShrinkRequest) SetTenantId(v string) *DeleteNumbersShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteNumbersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
