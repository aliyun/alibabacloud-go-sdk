// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAssetAttributesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateAssetAttributesShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpdateAssetAttributesShrinkRequest
	GetOpUserId() *string
	SetUpdateCommandShrink(v string) *UpdateAssetAttributesShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateAssetAttributesShrinkRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The update command.
	//
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateAssetAttributesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAssetAttributesShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAssetAttributesShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateAssetAttributesShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpdateAssetAttributesShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateAssetAttributesShrinkRequest) SetOpTenantId(v int64) *UpdateAssetAttributesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateAssetAttributesShrinkRequest) SetOpUserId(v string) *UpdateAssetAttributesShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateAssetAttributesShrinkRequest) SetUpdateCommandShrink(v string) *UpdateAssetAttributesShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateAssetAttributesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
