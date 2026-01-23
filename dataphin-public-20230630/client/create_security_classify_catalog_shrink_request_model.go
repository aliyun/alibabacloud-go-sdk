// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityClassifyCatalogShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateSecurityClassifyCatalogShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateSecurityClassifyCatalogShrinkRequest
	GetOpTenantId() *int64
}

type CreateSecurityClassifyCatalogShrinkRequest struct {
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateSecurityClassifyCatalogShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityClassifyCatalogShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityClassifyCatalogShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateSecurityClassifyCatalogShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateSecurityClassifyCatalogShrinkRequest) SetCreateCommandShrink(v string) *CreateSecurityClassifyCatalogShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateSecurityClassifyCatalogShrinkRequest) SetOpTenantId(v int64) *CreateSecurityClassifyCatalogShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateSecurityClassifyCatalogShrinkRequest) Validate() error {
	return dara.Validate(s)
}
