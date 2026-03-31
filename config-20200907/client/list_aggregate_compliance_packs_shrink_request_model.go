// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateCompliancePacksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *ListAggregateCompliancePacksShrinkRequest
	GetAggregatorId() *string
	SetPageNumber(v int32) *ListAggregateCompliancePacksShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAggregateCompliancePacksShrinkRequest
	GetPageSize() *int32
	SetRiskLevel(v int32) *ListAggregateCompliancePacksShrinkRequest
	GetRiskLevel() *int32
	SetStatus(v string) *ListAggregateCompliancePacksShrinkRequest
	GetStatus() *string
	SetTagShrink(v string) *ListAggregateCompliancePacksShrinkRequest
	GetTagShrink() *string
}

type ListAggregateCompliancePacksShrinkRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-f632626622af0079****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 20
	PageSize  *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The status of the one or more compliance packages to be queried. Valid values:
	//
	// 	- ACTIVE: compliance packages that are available for use.
	//
	// 	- CREATING: compliance packages that are being created.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListAggregateCompliancePacksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateCompliancePacksShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAggregateCompliancePacksShrinkRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *ListAggregateCompliancePacksShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAggregateCompliancePacksShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAggregateCompliancePacksShrinkRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *ListAggregateCompliancePacksShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAggregateCompliancePacksShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListAggregateCompliancePacksShrinkRequest) SetAggregatorId(v string) *ListAggregateCompliancePacksShrinkRequest {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateCompliancePacksShrinkRequest) SetPageNumber(v int32) *ListAggregateCompliancePacksShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAggregateCompliancePacksShrinkRequest) SetPageSize(v int32) *ListAggregateCompliancePacksShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListAggregateCompliancePacksShrinkRequest) SetRiskLevel(v int32) *ListAggregateCompliancePacksShrinkRequest {
	s.RiskLevel = &v
	return s
}

func (s *ListAggregateCompliancePacksShrinkRequest) SetStatus(v string) *ListAggregateCompliancePacksShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListAggregateCompliancePacksShrinkRequest) SetTagShrink(v string) *ListAggregateCompliancePacksShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListAggregateCompliancePacksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
