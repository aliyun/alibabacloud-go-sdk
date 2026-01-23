// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityLevelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateSecurityLevelShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateSecurityLevelShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateSecurityLevelShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateSecurityLevelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityLevelShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecurityLevelShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateSecurityLevelShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateSecurityLevelShrinkRequest) SetOpTenantId(v int64) *UpdateSecurityLevelShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateSecurityLevelShrinkRequest) SetUpdateCommandShrink(v string) *UpdateSecurityLevelShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateSecurityLevelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
