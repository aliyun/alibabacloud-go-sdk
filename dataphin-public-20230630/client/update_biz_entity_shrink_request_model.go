// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBizEntityShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateBizEntityShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateBizEntityShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateBizEntityShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateBizEntityShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizEntityShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateBizEntityShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateBizEntityShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateBizEntityShrinkRequest) SetOpTenantId(v int64) *UpdateBizEntityShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateBizEntityShrinkRequest) SetUpdateCommandShrink(v string) *UpdateBizEntityShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateBizEntityShrinkRequest) Validate() error {
	return dara.Validate(s)
}
