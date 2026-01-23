// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityIdentifyResultsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommandShrink(v string) *DeleteSecurityIdentifyResultsShrinkRequest
	GetDeleteCommandShrink() *string
	SetOpTenantId(v int64) *DeleteSecurityIdentifyResultsShrinkRequest
	GetOpTenantId() *int64
}

type DeleteSecurityIdentifyResultsShrinkRequest struct {
	// This parameter is required.
	DeleteCommandShrink *string `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteSecurityIdentifyResultsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityIdentifyResultsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityIdentifyResultsShrinkRequest) GetDeleteCommandShrink() *string {
	return s.DeleteCommandShrink
}

func (s *DeleteSecurityIdentifyResultsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteSecurityIdentifyResultsShrinkRequest) SetDeleteCommandShrink(v string) *DeleteSecurityIdentifyResultsShrinkRequest {
	s.DeleteCommandShrink = &v
	return s
}

func (s *DeleteSecurityIdentifyResultsShrinkRequest) SetOpTenantId(v int64) *DeleteSecurityIdentifyResultsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteSecurityIdentifyResultsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
