// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRegisterLineageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteRegisterLineageCommandShrink(v string) *DeleteRegisterLineageShrinkRequest
	GetDeleteRegisterLineageCommandShrink() *string
	SetOpTenantId(v int64) *DeleteRegisterLineageShrinkRequest
	GetOpTenantId() *int64
}

type DeleteRegisterLineageShrinkRequest struct {
	// This parameter is required.
	DeleteRegisterLineageCommandShrink *string `json:"DeleteRegisterLineageCommand,omitempty" xml:"DeleteRegisterLineageCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteRegisterLineageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRegisterLineageShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteRegisterLineageShrinkRequest) GetDeleteRegisterLineageCommandShrink() *string {
	return s.DeleteRegisterLineageCommandShrink
}

func (s *DeleteRegisterLineageShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteRegisterLineageShrinkRequest) SetDeleteRegisterLineageCommandShrink(v string) *DeleteRegisterLineageShrinkRequest {
	s.DeleteRegisterLineageCommandShrink = &v
	return s
}

func (s *DeleteRegisterLineageShrinkRequest) SetOpTenantId(v int64) *DeleteRegisterLineageShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteRegisterLineageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
