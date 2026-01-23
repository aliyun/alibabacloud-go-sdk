// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardRelationsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommandShrink(v string) *DeleteStandardRelationsShrinkRequest
	GetDeleteCommandShrink() *string
	SetOpTenantId(v int64) *DeleteStandardRelationsShrinkRequest
	GetOpTenantId() *int64
}

type DeleteStandardRelationsShrinkRequest struct {
	// This parameter is required.
	DeleteCommandShrink *string `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteStandardRelationsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardRelationsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteStandardRelationsShrinkRequest) GetDeleteCommandShrink() *string {
	return s.DeleteCommandShrink
}

func (s *DeleteStandardRelationsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteStandardRelationsShrinkRequest) SetDeleteCommandShrink(v string) *DeleteStandardRelationsShrinkRequest {
	s.DeleteCommandShrink = &v
	return s
}

func (s *DeleteStandardRelationsShrinkRequest) SetOpTenantId(v int64) *DeleteStandardRelationsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteStandardRelationsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
