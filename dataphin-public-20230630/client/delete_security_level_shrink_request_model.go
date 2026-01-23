// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityLevelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommandShrink(v string) *DeleteSecurityLevelShrinkRequest
	GetDeleteCommandShrink() *string
	SetOpTenantId(v int64) *DeleteSecurityLevelShrinkRequest
	GetOpTenantId() *int64
}

type DeleteSecurityLevelShrinkRequest struct {
	// This parameter is required.
	DeleteCommandShrink *string `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteSecurityLevelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityLevelShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityLevelShrinkRequest) GetDeleteCommandShrink() *string {
	return s.DeleteCommandShrink
}

func (s *DeleteSecurityLevelShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteSecurityLevelShrinkRequest) SetDeleteCommandShrink(v string) *DeleteSecurityLevelShrinkRequest {
	s.DeleteCommandShrink = &v
	return s
}

func (s *DeleteSecurityLevelShrinkRequest) SetOpTenantId(v int64) *DeleteSecurityLevelShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteSecurityLevelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
