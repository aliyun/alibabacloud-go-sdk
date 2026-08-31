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
	SetOpUserId(v string) *AddRegisterLineageShrinkRequest
	GetOpUserId() *string
}

type AddRegisterLineageShrinkRequest struct {
	// The command for registering and adding data lineage.
	//
	// This parameter is required.
	AddRegisterLineageCommandShrink *string `json:"AddRegisterLineageCommand,omitempty" xml:"AddRegisterLineageCommand,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
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

func (s *AddRegisterLineageShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *AddRegisterLineageShrinkRequest) SetAddRegisterLineageCommandShrink(v string) *AddRegisterLineageShrinkRequest {
	s.AddRegisterLineageCommandShrink = &v
	return s
}

func (s *AddRegisterLineageShrinkRequest) SetOpTenantId(v int64) *AddRegisterLineageShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *AddRegisterLineageShrinkRequest) SetOpUserId(v string) *AddRegisterLineageShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *AddRegisterLineageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
