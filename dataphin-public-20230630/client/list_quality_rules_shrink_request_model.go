// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListQualityRulesShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListQualityRulesShrinkRequest
	GetOpTenantId() *int64
}

type ListQualityRulesShrinkRequest struct {
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListQualityRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListQualityRulesShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListQualityRulesShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListQualityRulesShrinkRequest) SetListQueryShrink(v string) *ListQualityRulesShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListQualityRulesShrinkRequest) SetOpTenantId(v int64) *ListQualityRulesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListQualityRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
