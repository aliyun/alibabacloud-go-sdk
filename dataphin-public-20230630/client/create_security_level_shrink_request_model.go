// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityLevelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateSecurityLevelShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateSecurityLevelShrinkRequest
	GetOpTenantId() *int64
}

type CreateSecurityLevelShrinkRequest struct {
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateSecurityLevelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityLevelShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityLevelShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateSecurityLevelShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateSecurityLevelShrinkRequest) SetCreateCommandShrink(v string) *CreateSecurityLevelShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateSecurityLevelShrinkRequest) SetOpTenantId(v int64) *CreateSecurityLevelShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateSecurityLevelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
