// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCostCheckResultsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssumeAliyunIdListShrink(v string) *DescribeCostCheckResultsShrinkRequest
	GetAssumeAliyunIdListShrink() *string
	SetCheckIdsShrink(v string) *DescribeCostCheckResultsShrinkRequest
	GetCheckIdsShrink() *string
	SetCheckPlanId(v int64) *DescribeCostCheckResultsShrinkRequest
	GetCheckPlanId() *int64
	SetGroupBy(v string) *DescribeCostCheckResultsShrinkRequest
	GetGroupBy() *string
	SetLanguage(v string) *DescribeCostCheckResultsShrinkRequest
	GetLanguage() *string
	SetProduct(v string) *DescribeCostCheckResultsShrinkRequest
	GetProduct() *string
	SetRegionIdsShrink(v string) *DescribeCostCheckResultsShrinkRequest
	GetRegionIdsShrink() *string
	SetResourceGroupIdListShrink(v string) *DescribeCostCheckResultsShrinkRequest
	GetResourceGroupIdListShrink() *string
	SetResourceId(v string) *DescribeCostCheckResultsShrinkRequest
	GetResourceId() *string
	SetResourceIdsShrink(v string) *DescribeCostCheckResultsShrinkRequest
	GetResourceIdsShrink() *string
	SetResourceName(v string) *DescribeCostCheckResultsShrinkRequest
	GetResourceName() *string
	SetSeverity(v int32) *DescribeCostCheckResultsShrinkRequest
	GetSeverity() *int32
	SetTagKeysShrink(v string) *DescribeCostCheckResultsShrinkRequest
	GetTagKeysShrink() *string
	SetTagListShrink(v string) *DescribeCostCheckResultsShrinkRequest
	GetTagListShrink() *string
	SetTagValuesShrink(v string) *DescribeCostCheckResultsShrinkRequest
	GetTagValuesShrink() *string
}

type DescribeCostCheckResultsShrinkRequest struct {
	AssumeAliyunIdListShrink *string `json:"AssumeAliyunIdList,omitempty" xml:"AssumeAliyunIdList,omitempty"`
	CheckIdsShrink           *string `json:"CheckIds,omitempty" xml:"CheckIds,omitempty"`
	CheckPlanId              *int64  `json:"CheckPlanId,omitempty" xml:"CheckPlanId,omitempty"`
	// example:
	//
	// Category
	GroupBy  *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// ecs
	Product                   *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RegionIdsShrink           *string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
	ResourceGroupIdListShrink *string `json:"ResourceGroupIdList,omitempty" xml:"ResourceGroupIdList,omitempty"`
	ResourceId                *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceIdsShrink         *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// example:
	//
	// SYNC_********_TASK
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// 1
	Severity        *int32  `json:"Severity,omitempty" xml:"Severity,omitempty"`
	TagKeysShrink   *string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty"`
	TagListShrink   *string `json:"TagList,omitempty" xml:"TagList,omitempty"`
	TagValuesShrink *string `json:"TagValues,omitempty" xml:"TagValues,omitempty"`
}

func (s DescribeCostCheckResultsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostCheckResultsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckResultsShrinkRequest) GetAssumeAliyunIdListShrink() *string {
	return s.AssumeAliyunIdListShrink
}

func (s *DescribeCostCheckResultsShrinkRequest) GetCheckIdsShrink() *string {
	return s.CheckIdsShrink
}

func (s *DescribeCostCheckResultsShrinkRequest) GetCheckPlanId() *int64 {
	return s.CheckPlanId
}

func (s *DescribeCostCheckResultsShrinkRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *DescribeCostCheckResultsShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeCostCheckResultsShrinkRequest) GetProduct() *string {
	return s.Product
}

func (s *DescribeCostCheckResultsShrinkRequest) GetRegionIdsShrink() *string {
	return s.RegionIdsShrink
}

func (s *DescribeCostCheckResultsShrinkRequest) GetResourceGroupIdListShrink() *string {
	return s.ResourceGroupIdListShrink
}

func (s *DescribeCostCheckResultsShrinkRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeCostCheckResultsShrinkRequest) GetResourceIdsShrink() *string {
	return s.ResourceIdsShrink
}

func (s *DescribeCostCheckResultsShrinkRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribeCostCheckResultsShrinkRequest) GetSeverity() *int32 {
	return s.Severity
}

func (s *DescribeCostCheckResultsShrinkRequest) GetTagKeysShrink() *string {
	return s.TagKeysShrink
}

func (s *DescribeCostCheckResultsShrinkRequest) GetTagListShrink() *string {
	return s.TagListShrink
}

func (s *DescribeCostCheckResultsShrinkRequest) GetTagValuesShrink() *string {
	return s.TagValuesShrink
}

func (s *DescribeCostCheckResultsShrinkRequest) SetAssumeAliyunIdListShrink(v string) *DescribeCostCheckResultsShrinkRequest {
	s.AssumeAliyunIdListShrink = &v
	return s
}

func (s *DescribeCostCheckResultsShrinkRequest) SetCheckIdsShrink(v string) *DescribeCostCheckResultsShrinkRequest {
	s.CheckIdsShrink = &v
	return s
}

func (s *DescribeCostCheckResultsShrinkRequest) SetCheckPlanId(v int64) *DescribeCostCheckResultsShrinkRequest {
	s.CheckPlanId = &v
	return s
}

func (s *DescribeCostCheckResultsShrinkRequest) SetGroupBy(v string) *DescribeCostCheckResultsShrinkRequest {
	s.GroupBy = &v
	return s
}

func (s *DescribeCostCheckResultsShrinkRequest) SetLanguage(v string) *DescribeCostCheckResultsShrinkRequest {
	s.Language = &v
	return s
}

func (s *DescribeCostCheckResultsShrinkRequest) SetProduct(v string) *DescribeCostCheckResultsShrinkRequest {
	s.Product = &v
	return s
}

func (s *DescribeCostCheckResultsShrinkRequest) SetRegionIdsShrink(v string) *DescribeCostCheckResultsShrinkRequest {
	s.RegionIdsShrink = &v
	return s
}

func (s *DescribeCostCheckResultsShrinkRequest) SetResourceGroupIdListShrink(v string) *DescribeCostCheckResultsShrinkRequest {
	s.ResourceGroupIdListShrink = &v
	return s
}

func (s *DescribeCostCheckResultsShrinkRequest) SetResourceId(v string) *DescribeCostCheckResultsShrinkRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeCostCheckResultsShrinkRequest) SetResourceIdsShrink(v string) *DescribeCostCheckResultsShrinkRequest {
	s.ResourceIdsShrink = &v
	return s
}

func (s *DescribeCostCheckResultsShrinkRequest) SetResourceName(v string) *DescribeCostCheckResultsShrinkRequest {
	s.ResourceName = &v
	return s
}

func (s *DescribeCostCheckResultsShrinkRequest) SetSeverity(v int32) *DescribeCostCheckResultsShrinkRequest {
	s.Severity = &v
	return s
}

func (s *DescribeCostCheckResultsShrinkRequest) SetTagKeysShrink(v string) *DescribeCostCheckResultsShrinkRequest {
	s.TagKeysShrink = &v
	return s
}

func (s *DescribeCostCheckResultsShrinkRequest) SetTagListShrink(v string) *DescribeCostCheckResultsShrinkRequest {
	s.TagListShrink = &v
	return s
}

func (s *DescribeCostCheckResultsShrinkRequest) SetTagValuesShrink(v string) *DescribeCostCheckResultsShrinkRequest {
	s.TagValuesShrink = &v
	return s
}

func (s *DescribeCostCheckResultsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
