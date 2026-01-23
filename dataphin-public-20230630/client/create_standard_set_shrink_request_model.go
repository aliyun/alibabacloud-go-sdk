// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardSetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateStandardSetShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateStandardSetShrinkRequest
	GetOpTenantId() *int64
}

type CreateStandardSetShrinkRequest struct {
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateStandardSetShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardSetShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateStandardSetShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateStandardSetShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateStandardSetShrinkRequest) SetCreateCommandShrink(v string) *CreateStandardSetShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateStandardSetShrinkRequest) SetOpTenantId(v int64) *CreateStandardSetShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateStandardSetShrinkRequest) Validate() error {
	return dara.Validate(s)
}
