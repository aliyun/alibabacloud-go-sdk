// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitQualityRuleTasksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *SubmitQualityRuleTasksShrinkRequest
	GetOpTenantId() *int64
	SetSubmitCommandShrink(v string) *SubmitQualityRuleTasksShrinkRequest
	GetSubmitCommandShrink() *string
}

type SubmitQualityRuleTasksShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	SubmitCommandShrink *string `json:"SubmitCommand,omitempty" xml:"SubmitCommand,omitempty"`
}

func (s SubmitQualityRuleTasksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitQualityRuleTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitQualityRuleTasksShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *SubmitQualityRuleTasksShrinkRequest) GetSubmitCommandShrink() *string {
	return s.SubmitCommandShrink
}

func (s *SubmitQualityRuleTasksShrinkRequest) SetOpTenantId(v int64) *SubmitQualityRuleTasksShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *SubmitQualityRuleTasksShrinkRequest) SetSubmitCommandShrink(v string) *SubmitQualityRuleTasksShrinkRequest {
	s.SubmitCommandShrink = &v
	return s
}

func (s *SubmitQualityRuleTasksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
