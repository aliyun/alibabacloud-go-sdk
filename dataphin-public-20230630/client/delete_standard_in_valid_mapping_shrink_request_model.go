// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardInValidMappingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommandShrink(v string) *DeleteStandardInValidMappingShrinkRequest
	GetDeleteCommandShrink() *string
	SetOpTenantId(v int64) *DeleteStandardInValidMappingShrinkRequest
	GetOpTenantId() *int64
}

type DeleteStandardInValidMappingShrinkRequest struct {
	// This parameter is required.
	DeleteCommandShrink *string `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteStandardInValidMappingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardInValidMappingShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteStandardInValidMappingShrinkRequest) GetDeleteCommandShrink() *string {
	return s.DeleteCommandShrink
}

func (s *DeleteStandardInValidMappingShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteStandardInValidMappingShrinkRequest) SetDeleteCommandShrink(v string) *DeleteStandardInValidMappingShrinkRequest {
	s.DeleteCommandShrink = &v
	return s
}

func (s *DeleteStandardInValidMappingShrinkRequest) SetOpTenantId(v int64) *DeleteStandardInValidMappingShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteStandardInValidMappingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
