// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityTemplatesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListQualityTemplatesShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListQualityTemplatesShrinkRequest
	GetOpTenantId() *int64
}

type ListQualityTemplatesShrinkRequest struct {
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListQualityTemplatesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQualityTemplatesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListQualityTemplatesShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListQualityTemplatesShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListQualityTemplatesShrinkRequest) SetListQueryShrink(v string) *ListQualityTemplatesShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListQualityTemplatesShrinkRequest) SetOpTenantId(v int64) *ListQualityTemplatesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListQualityTemplatesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
