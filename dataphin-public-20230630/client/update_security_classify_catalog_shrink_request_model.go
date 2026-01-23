// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityClassifyCatalogShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateSecurityClassifyCatalogShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateSecurityClassifyCatalogShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateSecurityClassifyCatalogShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateSecurityClassifyCatalogShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityClassifyCatalogShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecurityClassifyCatalogShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateSecurityClassifyCatalogShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateSecurityClassifyCatalogShrinkRequest) SetOpTenantId(v int64) *UpdateSecurityClassifyCatalogShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateSecurityClassifyCatalogShrinkRequest) SetUpdateCommandShrink(v string) *UpdateSecurityClassifyCatalogShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateSecurityClassifyCatalogShrinkRequest) Validate() error {
	return dara.Validate(s)
}
