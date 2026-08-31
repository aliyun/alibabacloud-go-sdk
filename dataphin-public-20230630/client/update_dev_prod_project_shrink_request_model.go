// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDevProdProjectShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateDevProdProjectShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpdateDevProdProjectShrinkRequest
	GetOpUserId() *string
	SetUpdateCommandShrink(v string) *UpdateDevProdProjectShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateDevProdProjectShrinkRequest struct {
	// The tenant ID.
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
	// The update command.
	//
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateDevProdProjectShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDevProdProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDevProdProjectShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateDevProdProjectShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpdateDevProdProjectShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateDevProdProjectShrinkRequest) SetOpTenantId(v int64) *UpdateDevProdProjectShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateDevProdProjectShrinkRequest) SetOpUserId(v string) *UpdateDevProdProjectShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateDevProdProjectShrinkRequest) SetUpdateCommandShrink(v string) *UpdateDevProdProjectShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateDevProdProjectShrinkRequest) Validate() error {
	return dara.Validate(s)
}
