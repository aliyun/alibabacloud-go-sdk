// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityClassifyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateSecurityClassifyShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateSecurityClassifyShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateSecurityClassifyShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateSecurityClassifyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityClassifyShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecurityClassifyShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateSecurityClassifyShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateSecurityClassifyShrinkRequest) SetOpTenantId(v int64) *UpdateSecurityClassifyShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateSecurityClassifyShrinkRequest) SetUpdateCommandShrink(v string) *UpdateSecurityClassifyShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateSecurityClassifyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
