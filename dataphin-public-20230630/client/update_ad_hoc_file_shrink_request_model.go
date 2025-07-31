// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAdHocFileShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateAdHocFileShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateAdHocFileShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateAdHocFileShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateAdHocFileShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdHocFileShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAdHocFileShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateAdHocFileShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateAdHocFileShrinkRequest) SetOpTenantId(v int64) *UpdateAdHocFileShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateAdHocFileShrinkRequest) SetUpdateCommandShrink(v string) *UpdateAdHocFileShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateAdHocFileShrinkRequest) Validate() error {
	return dara.Validate(s)
}
