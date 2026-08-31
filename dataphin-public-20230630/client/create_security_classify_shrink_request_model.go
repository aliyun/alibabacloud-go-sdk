// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityClassifyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateSecurityClassifyShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateSecurityClassifyShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CreateSecurityClassifyShrinkRequest
	GetOpUserId() *string
}

type CreateSecurityClassifyShrinkRequest struct {
	// The create instruction.
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
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s CreateSecurityClassifyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityClassifyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityClassifyShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateSecurityClassifyShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateSecurityClassifyShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CreateSecurityClassifyShrinkRequest) SetCreateCommandShrink(v string) *CreateSecurityClassifyShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateSecurityClassifyShrinkRequest) SetOpTenantId(v int64) *CreateSecurityClassifyShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateSecurityClassifyShrinkRequest) SetOpUserId(v string) *CreateSecurityClassifyShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateSecurityClassifyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
