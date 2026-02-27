// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataServiceAppGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateDataServiceAppGroupShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateDataServiceAppGroupShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateDataServiceAppGroupShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateDataServiceAppGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataServiceAppGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataServiceAppGroupShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateDataServiceAppGroupShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateDataServiceAppGroupShrinkRequest) SetOpTenantId(v int64) *UpdateDataServiceAppGroupShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateDataServiceAppGroupShrinkRequest) SetUpdateCommandShrink(v string) *UpdateDataServiceAppGroupShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateDataServiceAppGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
