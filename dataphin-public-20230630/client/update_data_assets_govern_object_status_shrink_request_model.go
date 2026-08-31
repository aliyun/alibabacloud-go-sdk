// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataAssetsGovernObjectStatusShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateDataAssetsGovernObjectStatusShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpdateDataAssetsGovernObjectStatusShrinkRequest
	GetOpUserId() *string
	SetUpdateCommandShrink(v string) *UpdateDataAssetsGovernObjectStatusShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateDataAssetsGovernObjectStatusShrinkRequest struct {
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
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateDataAssetsGovernObjectStatusShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataAssetsGovernObjectStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataAssetsGovernObjectStatusShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateDataAssetsGovernObjectStatusShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpdateDataAssetsGovernObjectStatusShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateDataAssetsGovernObjectStatusShrinkRequest) SetOpTenantId(v int64) *UpdateDataAssetsGovernObjectStatusShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateDataAssetsGovernObjectStatusShrinkRequest) SetOpUserId(v string) *UpdateDataAssetsGovernObjectStatusShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateDataAssetsGovernObjectStatusShrinkRequest) SetUpdateCommandShrink(v string) *UpdateDataAssetsGovernObjectStatusShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateDataAssetsGovernObjectStatusShrinkRequest) Validate() error {
	return dara.Validate(s)
}
