// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCostRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeCostRulesResponseBodyItems) *DescribeCostRulesResponseBody
	GetItems() []*DescribeCostRulesResponseBodyItems
	SetPageNumber(v int32) *DescribeCostRulesResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeCostRulesResponseBody
	GetPageRecordCount() *int32
	SetPageSize(v int32) *DescribeCostRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCostRulesResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeCostRulesResponseBody
	GetTotalRecordCount() *int32
}

type DescribeCostRulesResponseBody struct {
	// The list of cost rules.
	Items []*DescribeCostRulesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries on the current page.
	//
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The number of entries per page. Valid values: 30, 50, and 100. Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 24A1990B-4F6E-482B-B8CB-75C612******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeCostRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCostRulesResponseBody) GetItems() []*DescribeCostRulesResponseBodyItems {
	return s.Items
}

func (s *DescribeCostRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCostRulesResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeCostRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCostRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCostRulesResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeCostRulesResponseBody) SetItems(v []*DescribeCostRulesResponseBodyItems) *DescribeCostRulesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeCostRulesResponseBody) SetPageNumber(v int32) *DescribeCostRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCostRulesResponseBody) SetPageRecordCount(v int32) *DescribeCostRulesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeCostRulesResponseBody) SetPageSize(v int32) *DescribeCostRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCostRulesResponseBody) SetRequestId(v string) *DescribeCostRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCostRulesResponseBody) SetTotalRecordCount(v int32) *DescribeCostRulesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeCostRulesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCostRulesResponseBodyItems struct {
	// The cache read cost points per million tokens. Default value: 0.
	//
	// example:
	//
	// 0
	CacheCostPointsPerMillion *string `json:"CacheCostPointsPerMillion,omitempty" xml:"CacheCostPointsPerMillion,omitempty"`
	// The cache creation or write cost points per million tokens.
	//
	// example:
	//
	// 400
	CacheCreationCostPointsPerMillion *string `json:"CacheCreationCostPointsPerMillion,omitempty" xml:"CacheCreationCostPointsPerMillion,omitempty"`
	// The ID of the cost rule.
	//
	// example:
	//
	// 924d450014e64e88ac6e8486f8e990**
	CostRuleId *string `json:"CostRuleId,omitempty" xml:"CostRuleId,omitempty"`
	// The end time of the effective period.
	//
	// example:
	//
	// 2026-10-01T00:00:00Z
	EffectiveEndTime *string `json:"EffectiveEndTime,omitempty" xml:"EffectiveEndTime,omitempty"`
	// The start time of the effective period.
	//
	// example:
	//
	// 2026-09-01T00:00:00Z
	EffectiveStartTime *string `json:"EffectiveStartTime,omitempty" xml:"EffectiveStartTime,omitempty"`
	// The type of the effective target.
	//
	// example:
	//
	// global
	EffectiveTargetType *string `json:"EffectiveTargetType,omitempty" xml:"EffectiveTargetType,omitempty"`
	// The value of the effective target.
	//
	// example:
	//
	// user
	EffectiveTargetValue *string `json:"EffectiveTargetValue,omitempty" xml:"EffectiveTargetValue,omitempty"`
	// The explicit cache read cost points per million tokens.
	//
	// example:
	//
	// 100
	ExplicitCacheCostPointsPerMillion *string `json:"ExplicitCacheCostPointsPerMillion,omitempty" xml:"ExplicitCacheCostPointsPerMillion,omitempty"`
	// The time when the cost rule was created.
	//
	// example:
	//
	// 2026-01-04T16:09:29+08:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the cost rule was last updated.
	//
	// example:
	//
	// 2024-10-16 16:46:20
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The gateway instance ID.
	//
	// example:
	//
	// pg-xxxxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// The image input cost points per image.
	//
	// example:
	//
	// 5
	ImageCostPointsPerCount *string `json:"ImageCostPointsPerCount,omitempty" xml:"ImageCostPointsPerCount,omitempty"`
	// The input cost points per million tokens. Default value: 0.
	//
	// example:
	//
	// 0
	InputCostPointsPerMillion *string `json:"InputCostPointsPerMillion,omitempty" xml:"InputCostPointsPerMillion,omitempty"`
	// The input token tier. An empty value indicates a legacy or default rule.
	//
	// example:
	//
	// 0-32K
	InputTokenTier *string `json:"InputTokenTier,omitempty" xml:"InputTokenTier,omitempty"`
	// The non-cumulative tiered pricing for input tokens. Each element includes the tier and the input cost points per million tokens.
	InputTokenTiers []*DescribeCostRulesResponseBodyItemsInputTokenTiers `json:"InputTokenTiers,omitempty" xml:"InputTokenTiers,omitempty" type:"Repeated"`
	// The name of the model.
	//
	// example:
	//
	// gpt-4
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The ID of the model service.
	//
	// example:
	//
	// ms-xxxxxxxxx
	ModelServiceId *string `json:"ModelServiceId,omitempty" xml:"ModelServiceId,omitempty"`
	// The model service name.
	//
	// example:
	//
	// qwen-service
	ModelServiceName *string `json:"ModelServiceName,omitempty" xml:"ModelServiceName,omitempty"`
	// The output cost points per million tokens. Default value: 0.
	//
	// example:
	//
	// 0
	OutputCostPointsPerMillion *string `json:"OutputCostPointsPerMillion,omitempty" xml:"OutputCostPointsPerMillion,omitempty"`
	// The video input cost points per second.
	//
	// example:
	//
	// 20
	VideoCostPointsPerSecond *string `json:"VideoCostPointsPerSecond,omitempty" xml:"VideoCostPointsPerSecond,omitempty"`
}

func (s DescribeCostRulesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostRulesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeCostRulesResponseBodyItems) GetCacheCostPointsPerMillion() *string {
	return s.CacheCostPointsPerMillion
}

func (s *DescribeCostRulesResponseBodyItems) GetCacheCreationCostPointsPerMillion() *string {
	return s.CacheCreationCostPointsPerMillion
}

func (s *DescribeCostRulesResponseBodyItems) GetCostRuleId() *string {
	return s.CostRuleId
}

func (s *DescribeCostRulesResponseBodyItems) GetEffectiveEndTime() *string {
	return s.EffectiveEndTime
}

func (s *DescribeCostRulesResponseBodyItems) GetEffectiveStartTime() *string {
	return s.EffectiveStartTime
}

func (s *DescribeCostRulesResponseBodyItems) GetEffectiveTargetType() *string {
	return s.EffectiveTargetType
}

func (s *DescribeCostRulesResponseBodyItems) GetEffectiveTargetValue() *string {
	return s.EffectiveTargetValue
}

func (s *DescribeCostRulesResponseBodyItems) GetExplicitCacheCostPointsPerMillion() *string {
	return s.ExplicitCacheCostPointsPerMillion
}

func (s *DescribeCostRulesResponseBodyItems) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeCostRulesResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeCostRulesResponseBodyItems) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DescribeCostRulesResponseBodyItems) GetImageCostPointsPerCount() *string {
	return s.ImageCostPointsPerCount
}

func (s *DescribeCostRulesResponseBodyItems) GetInputCostPointsPerMillion() *string {
	return s.InputCostPointsPerMillion
}

func (s *DescribeCostRulesResponseBodyItems) GetInputTokenTier() *string {
	return s.InputTokenTier
}

func (s *DescribeCostRulesResponseBodyItems) GetInputTokenTiers() []*DescribeCostRulesResponseBodyItemsInputTokenTiers {
	return s.InputTokenTiers
}

func (s *DescribeCostRulesResponseBodyItems) GetModel() *string {
	return s.Model
}

func (s *DescribeCostRulesResponseBodyItems) GetModelServiceId() *string {
	return s.ModelServiceId
}

func (s *DescribeCostRulesResponseBodyItems) GetModelServiceName() *string {
	return s.ModelServiceName
}

func (s *DescribeCostRulesResponseBodyItems) GetOutputCostPointsPerMillion() *string {
	return s.OutputCostPointsPerMillion
}

func (s *DescribeCostRulesResponseBodyItems) GetVideoCostPointsPerSecond() *string {
	return s.VideoCostPointsPerSecond
}

func (s *DescribeCostRulesResponseBodyItems) SetCacheCostPointsPerMillion(v string) *DescribeCostRulesResponseBodyItems {
	s.CacheCostPointsPerMillion = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetCacheCreationCostPointsPerMillion(v string) *DescribeCostRulesResponseBodyItems {
	s.CacheCreationCostPointsPerMillion = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetCostRuleId(v string) *DescribeCostRulesResponseBodyItems {
	s.CostRuleId = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetEffectiveEndTime(v string) *DescribeCostRulesResponseBodyItems {
	s.EffectiveEndTime = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetEffectiveStartTime(v string) *DescribeCostRulesResponseBodyItems {
	s.EffectiveStartTime = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetEffectiveTargetType(v string) *DescribeCostRulesResponseBodyItems {
	s.EffectiveTargetType = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetEffectiveTargetValue(v string) *DescribeCostRulesResponseBodyItems {
	s.EffectiveTargetValue = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetExplicitCacheCostPointsPerMillion(v string) *DescribeCostRulesResponseBodyItems {
	s.ExplicitCacheCostPointsPerMillion = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetGmtCreated(v string) *DescribeCostRulesResponseBodyItems {
	s.GmtCreated = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetGmtModified(v string) *DescribeCostRulesResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetGwClusterId(v string) *DescribeCostRulesResponseBodyItems {
	s.GwClusterId = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetImageCostPointsPerCount(v string) *DescribeCostRulesResponseBodyItems {
	s.ImageCostPointsPerCount = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetInputCostPointsPerMillion(v string) *DescribeCostRulesResponseBodyItems {
	s.InputCostPointsPerMillion = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetInputTokenTier(v string) *DescribeCostRulesResponseBodyItems {
	s.InputTokenTier = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetInputTokenTiers(v []*DescribeCostRulesResponseBodyItemsInputTokenTiers) *DescribeCostRulesResponseBodyItems {
	s.InputTokenTiers = v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetModel(v string) *DescribeCostRulesResponseBodyItems {
	s.Model = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetModelServiceId(v string) *DescribeCostRulesResponseBodyItems {
	s.ModelServiceId = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetModelServiceName(v string) *DescribeCostRulesResponseBodyItems {
	s.ModelServiceName = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetOutputCostPointsPerMillion(v string) *DescribeCostRulesResponseBodyItems {
	s.OutputCostPointsPerMillion = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetVideoCostPointsPerSecond(v string) *DescribeCostRulesResponseBodyItems {
	s.VideoCostPointsPerSecond = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) Validate() error {
	if s.InputTokenTiers != nil {
		for _, item := range s.InputTokenTiers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCostRulesResponseBodyItemsInputTokenTiers struct {
	// The input cost points per million tokens for the tier.
	//
	// example:
	//
	// 1000
	InputCostPointsPerMillion *string `json:"InputCostPointsPerMillion,omitempty" xml:"InputCostPointsPerMillion,omitempty"`
	// The token tier. Valid values: "" (default tier), 0-32K, 32K-64K, 64K-128K, 128K-256K, 256K-512K, 512K-1M, and 1M+.
	//
	// example:
	//
	// 0-32K
	Tier *string `json:"Tier,omitempty" xml:"Tier,omitempty"`
}

func (s DescribeCostRulesResponseBodyItemsInputTokenTiers) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostRulesResponseBodyItemsInputTokenTiers) GoString() string {
	return s.String()
}

func (s *DescribeCostRulesResponseBodyItemsInputTokenTiers) GetInputCostPointsPerMillion() *string {
	return s.InputCostPointsPerMillion
}

func (s *DescribeCostRulesResponseBodyItemsInputTokenTiers) GetTier() *string {
	return s.Tier
}

func (s *DescribeCostRulesResponseBodyItemsInputTokenTiers) SetInputCostPointsPerMillion(v string) *DescribeCostRulesResponseBodyItemsInputTokenTiers {
	s.InputCostPointsPerMillion = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItemsInputTokenTiers) SetTier(v string) *DescribeCostRulesResponseBodyItemsInputTokenTiers {
	s.Tier = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItemsInputTokenTiers) Validate() error {
	return dara.Validate(s)
}
