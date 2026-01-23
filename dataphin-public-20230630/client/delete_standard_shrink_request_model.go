// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommandShrink(v string) *DeleteStandardShrinkRequest
	GetDeleteCommandShrink() *string
	SetOpTenantId(v int64) *DeleteStandardShrinkRequest
	GetOpTenantId() *int64
}

type DeleteStandardShrinkRequest struct {
	// This parameter is required.
	DeleteCommandShrink *string `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteStandardShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteStandardShrinkRequest) GetDeleteCommandShrink() *string {
	return s.DeleteCommandShrink
}

func (s *DeleteStandardShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteStandardShrinkRequest) SetDeleteCommandShrink(v string) *DeleteStandardShrinkRequest {
	s.DeleteCommandShrink = &v
	return s
}

func (s *DeleteStandardShrinkRequest) SetOpTenantId(v int64) *DeleteStandardShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteStandardShrinkRequest) Validate() error {
	return dara.Validate(s)
}
