// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateResourceShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateResourceShrinkRequest
	GetOpTenantId() *int64
}

type CreateResourceShrinkRequest struct {
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateResourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateResourceShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateResourceShrinkRequest) SetCreateCommandShrink(v string) *CreateResourceShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateResourceShrinkRequest) SetOpTenantId(v int64) *CreateResourceShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateResourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
