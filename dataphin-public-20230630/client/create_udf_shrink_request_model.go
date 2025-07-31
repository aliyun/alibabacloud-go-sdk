// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUdfShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateUdfShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateUdfShrinkRequest
	GetOpTenantId() *int64
}

type CreateUdfShrinkRequest struct {
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateUdfShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUdfShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateUdfShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateUdfShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateUdfShrinkRequest) SetCreateCommandShrink(v string) *CreateUdfShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateUdfShrinkRequest) SetOpTenantId(v int64) *CreateUdfShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateUdfShrinkRequest) Validate() error {
	return dara.Validate(s)
}
