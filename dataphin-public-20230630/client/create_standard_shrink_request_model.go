// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateStandardShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateStandardShrinkRequest
	GetOpTenantId() *int64
}

type CreateStandardShrinkRequest struct {
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateStandardShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateStandardShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateStandardShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateStandardShrinkRequest) SetCreateCommandShrink(v string) *CreateStandardShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateStandardShrinkRequest) SetOpTenantId(v int64) *CreateStandardShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateStandardShrinkRequest) Validate() error {
	return dara.Validate(s)
}
