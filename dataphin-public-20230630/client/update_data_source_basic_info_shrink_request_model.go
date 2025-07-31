// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSourceBasicInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateDataSourceBasicInfoShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateDataSourceBasicInfoShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateDataSourceBasicInfoShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateDataSourceBasicInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceBasicInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceBasicInfoShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateDataSourceBasicInfoShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateDataSourceBasicInfoShrinkRequest) SetOpTenantId(v int64) *UpdateDataSourceBasicInfoShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateDataSourceBasicInfoShrinkRequest) SetUpdateCommandShrink(v string) *UpdateDataSourceBasicInfoShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateDataSourceBasicInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
