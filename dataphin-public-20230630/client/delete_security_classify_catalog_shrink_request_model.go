// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityClassifyCatalogShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommandShrink(v string) *DeleteSecurityClassifyCatalogShrinkRequest
	GetDeleteCommandShrink() *string
	SetOpTenantId(v int64) *DeleteSecurityClassifyCatalogShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *DeleteSecurityClassifyCatalogShrinkRequest
	GetOpUserId() *string
}

type DeleteSecurityClassifyCatalogShrinkRequest struct {
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

func (s DeleteSecurityClassifyCatalogShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityClassifyCatalogShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityClassifyCatalogShrinkRequest) GetDeleteCommandShrink() *string {
	return s.DeleteCommandShrink
}

func (s *DeleteSecurityClassifyCatalogShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteSecurityClassifyCatalogShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *DeleteSecurityClassifyCatalogShrinkRequest) SetDeleteCommandShrink(v string) *DeleteSecurityClassifyCatalogShrinkRequest {
	s.DeleteCommandShrink = &v
	return s
}

func (s *DeleteSecurityClassifyCatalogShrinkRequest) SetOpTenantId(v int64) *DeleteSecurityClassifyCatalogShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteSecurityClassifyCatalogShrinkRequest) SetOpUserId(v string) *DeleteSecurityClassifyCatalogShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *DeleteSecurityClassifyCatalogShrinkRequest) Validate() error {
	return dara.Validate(s)
}
