// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSourceConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateDataSourceConfigShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateDataSourceConfigShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateDataSourceConfigShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateDataSourceConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceConfigShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateDataSourceConfigShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateDataSourceConfigShrinkRequest) SetOpTenantId(v int64) *UpdateDataSourceConfigShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateDataSourceConfigShrinkRequest) SetUpdateCommandShrink(v string) *UpdateDataSourceConfigShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateDataSourceConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
