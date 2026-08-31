// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataServiceAppShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateDataServiceAppShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpdateDataServiceAppShrinkRequest
	GetOpUserId() *string
	SetUpdateCommandShrink(v string) *UpdateDataServiceAppShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateDataServiceAppShrinkRequest struct {
	// Tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The command to update the data service application.
	//
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateDataServiceAppShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataServiceAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataServiceAppShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateDataServiceAppShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpdateDataServiceAppShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateDataServiceAppShrinkRequest) SetOpTenantId(v int64) *UpdateDataServiceAppShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateDataServiceAppShrinkRequest) SetOpUserId(v string) *UpdateDataServiceAppShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateDataServiceAppShrinkRequest) SetUpdateCommandShrink(v string) *UpdateDataServiceAppShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateDataServiceAppShrinkRequest) Validate() error {
	return dara.Validate(s)
}
