// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardMappingToInvalidShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateStandardMappingToInvalidShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpdateStandardMappingToInvalidShrinkRequest
	GetOpUserId() *string
	SetUpdateCommandShrink(v string) *UpdateStandardMappingToInvalidShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateStandardMappingToInvalidShrinkRequest struct {
	// Tenant ID
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
	// Update Command
	//
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateStandardMappingToInvalidShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardMappingToInvalidShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateStandardMappingToInvalidShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateStandardMappingToInvalidShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpdateStandardMappingToInvalidShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateStandardMappingToInvalidShrinkRequest) SetOpTenantId(v int64) *UpdateStandardMappingToInvalidShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateStandardMappingToInvalidShrinkRequest) SetOpUserId(v string) *UpdateStandardMappingToInvalidShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateStandardMappingToInvalidShrinkRequest) SetUpdateCommandShrink(v string) *UpdateStandardMappingToInvalidShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateStandardMappingToInvalidShrinkRequest) Validate() error {
	return dara.Validate(s)
}
