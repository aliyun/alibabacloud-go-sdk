// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityClassifyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommandShrink(v string) *DeleteSecurityClassifyShrinkRequest
	GetDeleteCommandShrink() *string
	SetOpTenantId(v int64) *DeleteSecurityClassifyShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *DeleteSecurityClassifyShrinkRequest
	GetOpUserId() *string
}

type DeleteSecurityClassifyShrinkRequest struct {
	// The delete instruction.
	//
	// This parameter is required.
	DeleteCommandShrink *string `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty"`
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

func (s DeleteSecurityClassifyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityClassifyShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityClassifyShrinkRequest) GetDeleteCommandShrink() *string {
	return s.DeleteCommandShrink
}

func (s *DeleteSecurityClassifyShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteSecurityClassifyShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *DeleteSecurityClassifyShrinkRequest) SetDeleteCommandShrink(v string) *DeleteSecurityClassifyShrinkRequest {
	s.DeleteCommandShrink = &v
	return s
}

func (s *DeleteSecurityClassifyShrinkRequest) SetOpTenantId(v int64) *DeleteSecurityClassifyShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteSecurityClassifyShrinkRequest) SetOpUserId(v string) *DeleteSecurityClassifyShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *DeleteSecurityClassifyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
