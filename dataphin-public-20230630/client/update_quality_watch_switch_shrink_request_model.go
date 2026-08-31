// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQualityWatchSwitchShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateQualityWatchSwitchShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpdateQualityWatchSwitchShrinkRequest
	GetOpUserId() *string
	SetUpdateCommandShrink(v string) *UpdateQualityWatchSwitchShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateQualityWatchSwitchShrinkRequest struct {
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
	// The update instruction.
	//
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateQualityWatchSwitchShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateQualityWatchSwitchShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateQualityWatchSwitchShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateQualityWatchSwitchShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpdateQualityWatchSwitchShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateQualityWatchSwitchShrinkRequest) SetOpTenantId(v int64) *UpdateQualityWatchSwitchShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateQualityWatchSwitchShrinkRequest) SetOpUserId(v string) *UpdateQualityWatchSwitchShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateQualityWatchSwitchShrinkRequest) SetUpdateCommandShrink(v string) *UpdateQualityWatchSwitchShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateQualityWatchSwitchShrinkRequest) Validate() error {
	return dara.Validate(s)
}
