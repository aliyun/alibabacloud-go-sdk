// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRegisterLineageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddRegisterLineageCommandShrink(v string) *AddRegisterLineageShrinkRequest
	GetAddRegisterLineageCommandShrink() *string
	SetOpTenantId(v int64) *AddRegisterLineageShrinkRequest
	GetOpTenantId() *int64
}

type AddRegisterLineageShrinkRequest struct {
	// This parameter is required.
	AddRegisterLineageCommandShrink *string `json:"AddRegisterLineageCommand,omitempty" xml:"AddRegisterLineageCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s AddRegisterLineageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddRegisterLineageShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddRegisterLineageShrinkRequest) GetAddRegisterLineageCommandShrink() *string {
	return s.AddRegisterLineageCommandShrink
}

func (s *AddRegisterLineageShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *AddRegisterLineageShrinkRequest) SetAddRegisterLineageCommandShrink(v string) *AddRegisterLineageShrinkRequest {
	s.AddRegisterLineageCommandShrink = &v
	return s
}

func (s *AddRegisterLineageShrinkRequest) SetOpTenantId(v int64) *AddRegisterLineageShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *AddRegisterLineageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
