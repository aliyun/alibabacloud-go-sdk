// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataServiceAppMemberShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateDataServiceAppMemberShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateDataServiceAppMemberShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateDataServiceAppMemberShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateDataServiceAppMemberShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataServiceAppMemberShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataServiceAppMemberShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateDataServiceAppMemberShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateDataServiceAppMemberShrinkRequest) SetOpTenantId(v int64) *UpdateDataServiceAppMemberShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateDataServiceAppMemberShrinkRequest) SetUpdateCommandShrink(v string) *UpdateDataServiceAppMemberShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateDataServiceAppMemberShrinkRequest) Validate() error {
	return dara.Validate(s)
}
