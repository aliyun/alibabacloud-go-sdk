// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityIdentifyResultStatusShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateSecurityIdentifyResultStatusShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateSecurityIdentifyResultStatusShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateSecurityIdentifyResultStatusShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateSecurityIdentifyResultStatusShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityIdentifyResultStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecurityIdentifyResultStatusShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateSecurityIdentifyResultStatusShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateSecurityIdentifyResultStatusShrinkRequest) SetOpTenantId(v int64) *UpdateSecurityIdentifyResultStatusShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateSecurityIdentifyResultStatusShrinkRequest) SetUpdateCommandShrink(v string) *UpdateSecurityIdentifyResultStatusShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateSecurityIdentifyResultStatusShrinkRequest) Validate() error {
	return dara.Validate(s)
}
