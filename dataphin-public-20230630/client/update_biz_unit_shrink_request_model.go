// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBizUnitShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateBizUnitShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateBizUnitShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateBizUnitShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateBizUnitShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizUnitShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateBizUnitShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateBizUnitShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateBizUnitShrinkRequest) SetOpTenantId(v int64) *UpdateBizUnitShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateBizUnitShrinkRequest) SetUpdateCommandShrink(v string) *UpdateBizUnitShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateBizUnitShrinkRequest) Validate() error {
	return dara.Validate(s)
}
