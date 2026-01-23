// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveQualityRuleSchedulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *RemoveQualityRuleSchedulesShrinkRequest
	GetOpTenantId() *int64
	SetRemoveCommandShrink(v string) *RemoveQualityRuleSchedulesShrinkRequest
	GetRemoveCommandShrink() *string
}

type RemoveQualityRuleSchedulesShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	RemoveCommandShrink *string `json:"RemoveCommand,omitempty" xml:"RemoveCommand,omitempty"`
}

func (s RemoveQualityRuleSchedulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveQualityRuleSchedulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveQualityRuleSchedulesShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *RemoveQualityRuleSchedulesShrinkRequest) GetRemoveCommandShrink() *string {
	return s.RemoveCommandShrink
}

func (s *RemoveQualityRuleSchedulesShrinkRequest) SetOpTenantId(v int64) *RemoveQualityRuleSchedulesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *RemoveQualityRuleSchedulesShrinkRequest) SetRemoveCommandShrink(v string) *RemoveQualityRuleSchedulesShrinkRequest {
	s.RemoveCommandShrink = &v
	return s
}

func (s *RemoveQualityRuleSchedulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
