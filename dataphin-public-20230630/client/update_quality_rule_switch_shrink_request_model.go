// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQualityRuleSwitchShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateQualityRuleSwitchShrinkRequest
	GetOpTenantId() *int64
	SetUpdateCommandShrink(v string) *UpdateQualityRuleSwitchShrinkRequest
	GetUpdateCommandShrink() *string
}

type UpdateQualityRuleSwitchShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateQualityRuleSwitchShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateQualityRuleSwitchShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateQualityRuleSwitchShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateQualityRuleSwitchShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *UpdateQualityRuleSwitchShrinkRequest) SetOpTenantId(v int64) *UpdateQualityRuleSwitchShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateQualityRuleSwitchShrinkRequest) SetUpdateCommandShrink(v string) *UpdateQualityRuleSwitchShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *UpdateQualityRuleSwitchShrinkRequest) Validate() error {
	return dara.Validate(s)
}
