// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateResourceShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateResourceShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateResourceShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateResourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateResourceShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateResourceShrinkRequest) SetOpTenantId(v int64) *UpdateResourceShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateResourceShrinkRequest) SetUpdateCommandShrink(v string) *UpdateResourceShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateResourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
