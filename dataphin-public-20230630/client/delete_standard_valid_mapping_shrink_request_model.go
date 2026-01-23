// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardValidMappingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommandShrink(v string) *DeleteStandardValidMappingShrinkRequest
	GetDeleteCommandShrink() *string
	SetOpTenantId(v int64) *DeleteStandardValidMappingShrinkRequest
	GetOpTenantId() *int64
}

type DeleteStandardValidMappingShrinkRequest struct {
	// This parameter is required.
	DeleteCommandShrink *string `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteStandardValidMappingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardValidMappingShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteStandardValidMappingShrinkRequest) GetDeleteCommandShrink() *string {
	return s.DeleteCommandShrink
}

func (s *DeleteStandardValidMappingShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteStandardValidMappingShrinkRequest) SetDeleteCommandShrink(v string) *DeleteStandardValidMappingShrinkRequest {
	s.DeleteCommandShrink = &v
	return s
}

func (s *DeleteStandardValidMappingShrinkRequest) SetOpTenantId(v int64) *DeleteStandardValidMappingShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteStandardValidMappingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
