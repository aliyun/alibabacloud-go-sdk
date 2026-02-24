// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCompliancePacksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListCompliancePacksShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCompliancePacksShrinkRequest
	GetPageSize() *int32
	SetRiskLevel(v int32) *ListCompliancePacksShrinkRequest
	GetRiskLevel() *int32
	SetStatus(v string) *ListCompliancePacksShrinkRequest
	GetStatus() *string
	SetTagShrink(v string) *ListCompliancePacksShrinkRequest
	GetTagShrink() *string
}

type ListCompliancePacksShrinkRequest struct {
	// The page number.
	//
	// Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The risk level of the compliance pack. Valid values:
	//
	// - 1: high risk.
	//
	// - 2: medium risk.
	//
	// - 3: low risk.
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The status of the compliance pack. Valid values:
	//
	// - ACTIVE: The compliance pack is active.
	//
	// - CREATING: The compliance pack is being created.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the resource.
	//
	// You can attach up to 20 tags to a resource.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListCompliancePacksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCompliancePacksShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCompliancePacksShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCompliancePacksShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCompliancePacksShrinkRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *ListCompliancePacksShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListCompliancePacksShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListCompliancePacksShrinkRequest) SetPageNumber(v int32) *ListCompliancePacksShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCompliancePacksShrinkRequest) SetPageSize(v int32) *ListCompliancePacksShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListCompliancePacksShrinkRequest) SetRiskLevel(v int32) *ListCompliancePacksShrinkRequest {
	s.RiskLevel = &v
	return s
}

func (s *ListCompliancePacksShrinkRequest) SetStatus(v string) *ListCompliancePacksShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListCompliancePacksShrinkRequest) SetTagShrink(v string) *ListCompliancePacksShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListCompliancePacksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
