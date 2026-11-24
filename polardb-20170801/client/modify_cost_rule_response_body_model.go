// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCostRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAffectedModelApiIds(v []*string) *ModifyCostRuleResponseBody
	GetAffectedModelApiIds() []*string
	SetCacheCostPointsPerMillion(v string) *ModifyCostRuleResponseBody
	GetCacheCostPointsPerMillion() *string
	SetCacheCreationCostPointsPerMillion(v string) *ModifyCostRuleResponseBody
	GetCacheCreationCostPointsPerMillion() *string
	SetCostRuleId(v string) *ModifyCostRuleResponseBody
	GetCostRuleId() *string
	SetEffectiveEndTime(v string) *ModifyCostRuleResponseBody
	GetEffectiveEndTime() *string
	SetEffectiveStartTime(v string) *ModifyCostRuleResponseBody
	GetEffectiveStartTime() *string
	SetEffectiveTargetType(v string) *ModifyCostRuleResponseBody
	GetEffectiveTargetType() *string
	SetEffectiveTargetValue(v string) *ModifyCostRuleResponseBody
	GetEffectiveTargetValue() *string
	SetExplicitCacheCostPointsPerMillion(v string) *ModifyCostRuleResponseBody
	GetExplicitCacheCostPointsPerMillion() *string
	SetGmtCreated(v string) *ModifyCostRuleResponseBody
	GetGmtCreated() *string
	SetGmtModified(v string) *ModifyCostRuleResponseBody
	GetGmtModified() *string
	SetGwClusterId(v string) *ModifyCostRuleResponseBody
	GetGwClusterId() *string
	SetImageCostPointsPerCount(v string) *ModifyCostRuleResponseBody
	GetImageCostPointsPerCount() *string
	SetInputCostPointsPerMillion(v string) *ModifyCostRuleResponseBody
	GetInputCostPointsPerMillion() *string
	SetInputTokenTier(v string) *ModifyCostRuleResponseBody
	GetInputTokenTier() *string
	SetInputTokenTiers(v []*ModifyCostRuleResponseBodyInputTokenTiers) *ModifyCostRuleResponseBody
	GetInputTokenTiers() []*ModifyCostRuleResponseBodyInputTokenTiers
	SetModelName(v string) *ModifyCostRuleResponseBody
	GetModelName() *string
	SetModelServiceId(v string) *ModifyCostRuleResponseBody
	GetModelServiceId() *string
	SetModelServiceName(v string) *ModifyCostRuleResponseBody
	GetModelServiceName() *string
	SetOutputCostPointsPerMillion(v string) *ModifyCostRuleResponseBody
	GetOutputCostPointsPerMillion() *string
	SetRequestId(v string) *ModifyCostRuleResponseBody
	GetRequestId() *string
	SetVideoCostPointsPerSecond(v string) *ModifyCostRuleResponseBody
	GetVideoCostPointsPerSecond() *string
}

type ModifyCostRuleResponseBody struct {
	// The list of model API IDs affected by the rule.
	//
	// example:
	//
	// ["ma-2ze3x6i1d7g5rto28"]
	AffectedModelApiIds []*string `json:"AffectedModelApiIds,omitempty" xml:"AffectedModelApiIds,omitempty" type:"Repeated"`
	// The cache read cost points per million tokens.
	//
	// example:
	//
	// 40
	CacheCostPointsPerMillion *string `json:"CacheCostPointsPerMillion,omitempty" xml:"CacheCostPointsPerMillion,omitempty"`
	// The cache creation cost points per million tokens.
	//
	// example:
	//
	// 400
	CacheCreationCostPointsPerMillion *string `json:"CacheCreationCostPointsPerMillion,omitempty" xml:"CacheCreationCostPointsPerMillion,omitempty"`
	// The cost rule ID.
	//
	// example:
	//
	// 924d450014e64e88ac6e8486f8e990**
	CostRuleId *string `json:"CostRuleId,omitempty" xml:"CostRuleId,omitempty"`
	// The effective end time.
	//
	// example:
	//
	// 2026-10-01T00:00:00Z
	EffectiveEndTime *string `json:"EffectiveEndTime,omitempty" xml:"EffectiveEndTime,omitempty"`
	// The effective start time.
	//
	// example:
	//
	// 2026-09-01T00:00:00Z
	EffectiveStartTime *string `json:"EffectiveStartTime,omitempty" xml:"EffectiveStartTime,omitempty"`
	// The effective target type. Valid values: global, consumerGroup, and consumer.
	//
	// example:
	//
	// global
	EffectiveTargetType *string `json:"EffectiveTargetType,omitempty" xml:"EffectiveTargetType,omitempty"`
	// The effective target value. This parameter has a value when the effective target type is not global.
	//
	// example:
	//
	// c-xxx
	EffectiveTargetValue *string `json:"EffectiveTargetValue,omitempty" xml:"EffectiveTargetValue,omitempty"`
	// The explicit cache read cost points per million tokens.
	//
	// example:
	//
	// 100
	ExplicitCacheCostPointsPerMillion *string `json:"ExplicitCacheCostPointsPerMillion,omitempty" xml:"ExplicitCacheCostPointsPerMillion,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-09-01T12:00:00Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2026-09-01T12:00:00Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the gateway cluster to which the cost rule belongs.
	//
	// example:
	//
	// pg-2ze3x6i1d7g5rto28
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// The image input cost points per image.
	//
	// example:
	//
	// 5
	ImageCostPointsPerCount *string `json:"ImageCostPointsPerCount,omitempty" xml:"ImageCostPointsPerCount,omitempty"`
	// The input cost points per million tokens.
	//
	// example:
	//
	// 1000
	InputCostPointsPerMillion *string `json:"InputCostPointsPerMillion,omitempty" xml:"InputCostPointsPerMillion,omitempty"`
	// The input token tier. An empty value indicates the default tier.
	//
	// example:
	//
	// 0-32K
	InputTokenTier *string `json:"InputTokenTier,omitempty" xml:"InputTokenTier,omitempty"`
	// The tiered pricing for input tokens (priced by tier, non-cumulative). Each element contains Tier and InputCostPointsPerMillion.
	//
	// example:
	//
	// [{"Tier":"0-32K","InputCostPointsPerMillion":"1000"}]
	InputTokenTiers []*ModifyCostRuleResponseBodyInputTokenTiers `json:"InputTokenTiers,omitempty" xml:"InputTokenTiers,omitempty" type:"Repeated"`
	// The name of the pricing model.
	//
	// example:
	//
	// qwen3-max
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The model service ID.
	//
	// example:
	//
	// ms-2ze3x6i1d7g5rto28
	ModelServiceId *string `json:"ModelServiceId,omitempty" xml:"ModelServiceId,omitempty"`
	// The model service name.
	//
	// example:
	//
	// model-service-a
	ModelServiceName *string `json:"ModelServiceName,omitempty" xml:"ModelServiceName,omitempty"`
	// The output cost points per million tokens.
	//
	// example:
	//
	// 2000
	OutputCostPointsPerMillion *string `json:"OutputCostPointsPerMillion,omitempty" xml:"OutputCostPointsPerMillion,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6BD9CDE4-5E7B-4BF3-9BB8-83C73E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The video input cost points per second.
	//
	// example:
	//
	// 20
	VideoCostPointsPerSecond *string `json:"VideoCostPointsPerSecond,omitempty" xml:"VideoCostPointsPerSecond,omitempty"`
}

func (s ModifyCostRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCostRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCostRuleResponseBody) GetAffectedModelApiIds() []*string {
	return s.AffectedModelApiIds
}

func (s *ModifyCostRuleResponseBody) GetCacheCostPointsPerMillion() *string {
	return s.CacheCostPointsPerMillion
}

func (s *ModifyCostRuleResponseBody) GetCacheCreationCostPointsPerMillion() *string {
	return s.CacheCreationCostPointsPerMillion
}

func (s *ModifyCostRuleResponseBody) GetCostRuleId() *string {
	return s.CostRuleId
}

func (s *ModifyCostRuleResponseBody) GetEffectiveEndTime() *string {
	return s.EffectiveEndTime
}

func (s *ModifyCostRuleResponseBody) GetEffectiveStartTime() *string {
	return s.EffectiveStartTime
}

func (s *ModifyCostRuleResponseBody) GetEffectiveTargetType() *string {
	return s.EffectiveTargetType
}

func (s *ModifyCostRuleResponseBody) GetEffectiveTargetValue() *string {
	return s.EffectiveTargetValue
}

func (s *ModifyCostRuleResponseBody) GetExplicitCacheCostPointsPerMillion() *string {
	return s.ExplicitCacheCostPointsPerMillion
}

func (s *ModifyCostRuleResponseBody) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *ModifyCostRuleResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ModifyCostRuleResponseBody) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *ModifyCostRuleResponseBody) GetImageCostPointsPerCount() *string {
	return s.ImageCostPointsPerCount
}

func (s *ModifyCostRuleResponseBody) GetInputCostPointsPerMillion() *string {
	return s.InputCostPointsPerMillion
}

func (s *ModifyCostRuleResponseBody) GetInputTokenTier() *string {
	return s.InputTokenTier
}

func (s *ModifyCostRuleResponseBody) GetInputTokenTiers() []*ModifyCostRuleResponseBodyInputTokenTiers {
	return s.InputTokenTiers
}

func (s *ModifyCostRuleResponseBody) GetModelName() *string {
	return s.ModelName
}

func (s *ModifyCostRuleResponseBody) GetModelServiceId() *string {
	return s.ModelServiceId
}

func (s *ModifyCostRuleResponseBody) GetModelServiceName() *string {
	return s.ModelServiceName
}

func (s *ModifyCostRuleResponseBody) GetOutputCostPointsPerMillion() *string {
	return s.OutputCostPointsPerMillion
}

func (s *ModifyCostRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCostRuleResponseBody) GetVideoCostPointsPerSecond() *string {
	return s.VideoCostPointsPerSecond
}

func (s *ModifyCostRuleResponseBody) SetAffectedModelApiIds(v []*string) *ModifyCostRuleResponseBody {
	s.AffectedModelApiIds = v
	return s
}

func (s *ModifyCostRuleResponseBody) SetCacheCostPointsPerMillion(v string) *ModifyCostRuleResponseBody {
	s.CacheCostPointsPerMillion = &v
	return s
}

func (s *ModifyCostRuleResponseBody) SetCacheCreationCostPointsPerMillion(v string) *ModifyCostRuleResponseBody {
	s.CacheCreationCostPointsPerMillion = &v
	return s
}

func (s *ModifyCostRuleResponseBody) SetCostRuleId(v string) *ModifyCostRuleResponseBody {
	s.CostRuleId = &v
	return s
}

func (s *ModifyCostRuleResponseBody) SetEffectiveEndTime(v string) *ModifyCostRuleResponseBody {
	s.EffectiveEndTime = &v
	return s
}

func (s *ModifyCostRuleResponseBody) SetEffectiveStartTime(v string) *ModifyCostRuleResponseBody {
	s.EffectiveStartTime = &v
	return s
}

func (s *ModifyCostRuleResponseBody) SetEffectiveTargetType(v string) *ModifyCostRuleResponseBody {
	s.EffectiveTargetType = &v
	return s
}

func (s *ModifyCostRuleResponseBody) SetEffectiveTargetValue(v string) *ModifyCostRuleResponseBody {
	s.EffectiveTargetValue = &v
	return s
}

func (s *ModifyCostRuleResponseBody) SetExplicitCacheCostPointsPerMillion(v string) *ModifyCostRuleResponseBody {
	s.ExplicitCacheCostPointsPerMillion = &v
	return s
}

func (s *ModifyCostRuleResponseBody) SetGmtCreated(v string) *ModifyCostRuleResponseBody {
	s.GmtCreated = &v
	return s
}

func (s *ModifyCostRuleResponseBody) SetGmtModified(v string) *ModifyCostRuleResponseBody {
	s.GmtModified = &v
	return s
}

func (s *ModifyCostRuleResponseBody) SetGwClusterId(v string) *ModifyCostRuleResponseBody {
	s.GwClusterId = &v
	return s
}

func (s *ModifyCostRuleResponseBody) SetImageCostPointsPerCount(v string) *ModifyCostRuleResponseBody {
	s.ImageCostPointsPerCount = &v
	return s
}

func (s *ModifyCostRuleResponseBody) SetInputCostPointsPerMillion(v string) *ModifyCostRuleResponseBody {
	s.InputCostPointsPerMillion = &v
	return s
}

func (s *ModifyCostRuleResponseBody) SetInputTokenTier(v string) *ModifyCostRuleResponseBody {
	s.InputTokenTier = &v
	return s
}

func (s *ModifyCostRuleResponseBody) SetInputTokenTiers(v []*ModifyCostRuleResponseBodyInputTokenTiers) *ModifyCostRuleResponseBody {
	s.InputTokenTiers = v
	return s
}

func (s *ModifyCostRuleResponseBody) SetModelName(v string) *ModifyCostRuleResponseBody {
	s.ModelName = &v
	return s
}

func (s *ModifyCostRuleResponseBody) SetModelServiceId(v string) *ModifyCostRuleResponseBody {
	s.ModelServiceId = &v
	return s
}

func (s *ModifyCostRuleResponseBody) SetModelServiceName(v string) *ModifyCostRuleResponseBody {
	s.ModelServiceName = &v
	return s
}

func (s *ModifyCostRuleResponseBody) SetOutputCostPointsPerMillion(v string) *ModifyCostRuleResponseBody {
	s.OutputCostPointsPerMillion = &v
	return s
}

func (s *ModifyCostRuleResponseBody) SetRequestId(v string) *ModifyCostRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCostRuleResponseBody) SetVideoCostPointsPerSecond(v string) *ModifyCostRuleResponseBody {
	s.VideoCostPointsPerSecond = &v
	return s
}

func (s *ModifyCostRuleResponseBody) Validate() error {
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

type ModifyCostRuleResponseBodyInputTokenTiers struct {
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

func (s ModifyCostRuleResponseBodyInputTokenTiers) String() string {
	return dara.Prettify(s)
}

func (s ModifyCostRuleResponseBodyInputTokenTiers) GoString() string {
	return s.String()
}

func (s *ModifyCostRuleResponseBodyInputTokenTiers) GetInputCostPointsPerMillion() *string {
	return s.InputCostPointsPerMillion
}

func (s *ModifyCostRuleResponseBodyInputTokenTiers) GetTier() *string {
	return s.Tier
}

func (s *ModifyCostRuleResponseBodyInputTokenTiers) SetInputCostPointsPerMillion(v string) *ModifyCostRuleResponseBodyInputTokenTiers {
	s.InputCostPointsPerMillion = &v
	return s
}

func (s *ModifyCostRuleResponseBodyInputTokenTiers) SetTier(v string) *ModifyCostRuleResponseBodyInputTokenTiers {
	s.Tier = &v
	return s
}

func (s *ModifyCostRuleResponseBodyInputTokenTiers) Validate() error {
	return dara.Validate(s)
}
