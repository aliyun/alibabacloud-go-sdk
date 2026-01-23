// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityRuleTasksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListQualityRuleTasksShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListQualityRuleTasksShrinkRequest
	GetOpTenantId() *int64
}

type ListQualityRuleTasksShrinkRequest struct {
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListQualityRuleTasksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRuleTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListQualityRuleTasksShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListQualityRuleTasksShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListQualityRuleTasksShrinkRequest) SetListQueryShrink(v string) *ListQualityRuleTasksShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListQualityRuleTasksShrinkRequest) SetOpTenantId(v int64) *ListQualityRuleTasksShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListQualityRuleTasksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
