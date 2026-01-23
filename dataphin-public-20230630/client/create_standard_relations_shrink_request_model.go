// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardRelationsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateStandardRelationsShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateStandardRelationsShrinkRequest
	GetOpTenantId() *int64
}

type CreateStandardRelationsShrinkRequest struct {
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateStandardRelationsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardRelationsShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateStandardRelationsShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateStandardRelationsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateStandardRelationsShrinkRequest) SetCreateCommandShrink(v string) *CreateStandardRelationsShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateStandardRelationsShrinkRequest) SetOpTenantId(v int64) *CreateStandardRelationsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateStandardRelationsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
