// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCostCheckAdvicesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssumeAliyunIdListShrink(v string) *DescribeCostCheckAdvicesShrinkRequest
	GetAssumeAliyunIdListShrink() *string
	SetCheckId(v string) *DescribeCostCheckAdvicesShrinkRequest
	GetCheckId() *string
	SetCheckPlanId(v int64) *DescribeCostCheckAdvicesShrinkRequest
	GetCheckPlanId() *int64
	SetLanguage(v string) *DescribeCostCheckAdvicesShrinkRequest
	GetLanguage() *string
	SetPageNumber(v int32) *DescribeCostCheckAdvicesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCostCheckAdvicesShrinkRequest
	GetPageSize() *int32
	SetRegionIdsShrink(v string) *DescribeCostCheckAdvicesShrinkRequest
	GetRegionIdsShrink() *string
	SetResourceGroupIdListShrink(v string) *DescribeCostCheckAdvicesShrinkRequest
	GetResourceGroupIdListShrink() *string
	SetResourceId(v string) *DescribeCostCheckAdvicesShrinkRequest
	GetResourceId() *string
	SetResourceIdsShrink(v string) *DescribeCostCheckAdvicesShrinkRequest
	GetResourceIdsShrink() *string
	SetResourceName(v string) *DescribeCostCheckAdvicesShrinkRequest
	GetResourceName() *string
	SetSeverity(v string) *DescribeCostCheckAdvicesShrinkRequest
	GetSeverity() *string
	SetTagKeysShrink(v string) *DescribeCostCheckAdvicesShrinkRequest
	GetTagKeysShrink() *string
	SetTagListShrink(v string) *DescribeCostCheckAdvicesShrinkRequest
	GetTagListShrink() *string
	SetTagValuesShrink(v string) *DescribeCostCheckAdvicesShrinkRequest
	GetTagValuesShrink() *string
}

type DescribeCostCheckAdvicesShrinkRequest struct {
	AssumeAliyunIdListShrink *string `json:"AssumeAliyunIdList,omitempty" xml:"AssumeAliyunIdList,omitempty"`
	// example:
	//
	// EcsCostLowUtilizationCheck
	CheckId     *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckPlanId *int64  `json:"CheckPlanId,omitempty" xml:"CheckPlanId,omitempty"`
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 6
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize                  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionIdsShrink           *string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
	ResourceGroupIdListShrink *string `json:"ResourceGroupIdList,omitempty" xml:"ResourceGroupIdList,omitempty"`
	ResourceId                *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceIdsShrink         *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// example:
	//
	// test
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// 1
	Severity        *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	TagKeysShrink   *string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty"`
	TagListShrink   *string `json:"TagList,omitempty" xml:"TagList,omitempty"`
	TagValuesShrink *string `json:"TagValues,omitempty" xml:"TagValues,omitempty"`
}

func (s DescribeCostCheckAdvicesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostCheckAdvicesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckAdvicesShrinkRequest) GetAssumeAliyunIdListShrink() *string {
	return s.AssumeAliyunIdListShrink
}

func (s *DescribeCostCheckAdvicesShrinkRequest) GetCheckId() *string {
	return s.CheckId
}

func (s *DescribeCostCheckAdvicesShrinkRequest) GetCheckPlanId() *int64 {
	return s.CheckPlanId
}

func (s *DescribeCostCheckAdvicesShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeCostCheckAdvicesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCostCheckAdvicesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCostCheckAdvicesShrinkRequest) GetRegionIdsShrink() *string {
	return s.RegionIdsShrink
}

func (s *DescribeCostCheckAdvicesShrinkRequest) GetResourceGroupIdListShrink() *string {
	return s.ResourceGroupIdListShrink
}

func (s *DescribeCostCheckAdvicesShrinkRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeCostCheckAdvicesShrinkRequest) GetResourceIdsShrink() *string {
	return s.ResourceIdsShrink
}

func (s *DescribeCostCheckAdvicesShrinkRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribeCostCheckAdvicesShrinkRequest) GetSeverity() *string {
	return s.Severity
}

func (s *DescribeCostCheckAdvicesShrinkRequest) GetTagKeysShrink() *string {
	return s.TagKeysShrink
}

func (s *DescribeCostCheckAdvicesShrinkRequest) GetTagListShrink() *string {
	return s.TagListShrink
}

func (s *DescribeCostCheckAdvicesShrinkRequest) GetTagValuesShrink() *string {
	return s.TagValuesShrink
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetAssumeAliyunIdListShrink(v string) *DescribeCostCheckAdvicesShrinkRequest {
	s.AssumeAliyunIdListShrink = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetCheckId(v string) *DescribeCostCheckAdvicesShrinkRequest {
	s.CheckId = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetCheckPlanId(v int64) *DescribeCostCheckAdvicesShrinkRequest {
	s.CheckPlanId = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetLanguage(v string) *DescribeCostCheckAdvicesShrinkRequest {
	s.Language = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetPageNumber(v int32) *DescribeCostCheckAdvicesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetPageSize(v int32) *DescribeCostCheckAdvicesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetRegionIdsShrink(v string) *DescribeCostCheckAdvicesShrinkRequest {
	s.RegionIdsShrink = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetResourceGroupIdListShrink(v string) *DescribeCostCheckAdvicesShrinkRequest {
	s.ResourceGroupIdListShrink = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetResourceId(v string) *DescribeCostCheckAdvicesShrinkRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetResourceIdsShrink(v string) *DescribeCostCheckAdvicesShrinkRequest {
	s.ResourceIdsShrink = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetResourceName(v string) *DescribeCostCheckAdvicesShrinkRequest {
	s.ResourceName = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetSeverity(v string) *DescribeCostCheckAdvicesShrinkRequest {
	s.Severity = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetTagKeysShrink(v string) *DescribeCostCheckAdvicesShrinkRequest {
	s.TagKeysShrink = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetTagListShrink(v string) *DescribeCostCheckAdvicesShrinkRequest {
	s.TagListShrink = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) SetTagValuesShrink(v string) *DescribeCostCheckAdvicesShrinkRequest {
	s.TagValuesShrink = &v
	return s
}

func (s *DescribeCostCheckAdvicesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
