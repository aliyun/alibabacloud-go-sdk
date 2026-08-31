// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardMappingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateStandardMappingShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateStandardMappingShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CreateStandardMappingShrinkRequest
	GetOpUserId() *string
}

type CreateStandardMappingShrinkRequest struct {
	// The create command.
	//
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
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

func (s CreateStandardMappingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardMappingShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateStandardMappingShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateStandardMappingShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateStandardMappingShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CreateStandardMappingShrinkRequest) SetCreateCommandShrink(v string) *CreateStandardMappingShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateStandardMappingShrinkRequest) SetOpTenantId(v int64) *CreateStandardMappingShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateStandardMappingShrinkRequest) SetOpUserId(v string) *CreateStandardMappingShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateStandardMappingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
