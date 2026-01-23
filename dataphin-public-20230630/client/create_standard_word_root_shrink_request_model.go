// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardWordRootShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateStandardWordRootShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateStandardWordRootShrinkRequest
	GetOpTenantId() *int64
}

type CreateStandardWordRootShrinkRequest struct {
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateStandardWordRootShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardWordRootShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateStandardWordRootShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateStandardWordRootShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateStandardWordRootShrinkRequest) SetCreateCommandShrink(v string) *CreateStandardWordRootShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateStandardWordRootShrinkRequest) SetOpTenantId(v int64) *CreateStandardWordRootShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateStandardWordRootShrinkRequest) Validate() error {
	return dara.Validate(s)
}
