// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCostRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAffectedModelApiIds(v []*string) *CreateCostRuleResponseBody
	GetAffectedModelApiIds() []*string
	SetCacheCostPointsPerMillion(v string) *CreateCostRuleResponseBody
	GetCacheCostPointsPerMillion() *string
	SetCacheCreationCostPointsPerMillion(v string) *CreateCostRuleResponseBody
	GetCacheCreationCostPointsPerMillion() *string
	SetCostRuleId(v string) *CreateCostRuleResponseBody
	GetCostRuleId() *string
	SetEffectiveEndTime(v string) *CreateCostRuleResponseBody
	GetEffectiveEndTime() *string
	SetEffectiveStartTime(v string) *CreateCostRuleResponseBody
	GetEffectiveStartTime() *string
	SetEffectiveTargetType(v string) *CreateCostRuleResponseBody
	GetEffectiveTargetType() *string
	SetEffectiveTargetValue(v string) *CreateCostRuleResponseBody
	GetEffectiveTargetValue() *string
	SetExplicitCacheCostPointsPerMillion(v string) *CreateCostRuleResponseBody
	GetExplicitCacheCostPointsPerMillion() *string
	SetGmtCreated(v string) *CreateCostRuleResponseBody
	GetGmtCreated() *string
	SetGmtModified(v string) *CreateCostRuleResponseBody
	GetGmtModified() *string
	SetGwClusterId(v string) *CreateCostRuleResponseBody
	GetGwClusterId() *string
	SetImageCostPointsPerCount(v string) *CreateCostRuleResponseBody
	GetImageCostPointsPerCount() *string
	SetInputCostPointsPerMillion(v string) *CreateCostRuleResponseBody
	GetInputCostPointsPerMillion() *string
	SetInputTokenTier(v string) *CreateCostRuleResponseBody
	GetInputTokenTier() *string
	SetInputTokenTiers(v []*CreateCostRuleResponseBodyInputTokenTiers) *CreateCostRuleResponseBody
	GetInputTokenTiers() []*CreateCostRuleResponseBodyInputTokenTiers
	SetModelName(v string) *CreateCostRuleResponseBody
	GetModelName() *string
	SetModelServiceId(v string) *CreateCostRuleResponseBody
	GetModelServiceId() *string
	SetModelServiceName(v string) *CreateCostRuleResponseBody
	GetModelServiceName() *string
	SetOutputCostPointsPerMillion(v string) *CreateCostRuleResponseBody
	GetOutputCostPointsPerMillion() *string
	SetRequestId(v string) *CreateCostRuleResponseBody
	GetRequestId() *string
	SetVideoCostPointsPerSecond(v string) *CreateCostRuleResponseBody
	GetVideoCostPointsPerSecond() *string
}

type CreateCostRuleResponseBody struct {
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
	// The cache creation (write) cost points per million tokens.
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
	// The modification time.
	//
	// example:
	//
	// 2026-09-01T12:00:00Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the gateway instance to which the cost rule belongs.
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
	// The tiered pricing for input tokens. Pricing is applied per tier and is not cumulative. Each element contains Tier and InputCostPointsPerMillion.
	//
	// example:
	//
	// [{"Tier":"0-32K","InputCostPointsPerMillion":"1000"}]
	InputTokenTiers []*CreateCostRuleResponseBodyInputTokenTiers `json:"InputTokenTiers,omitempty" xml:"InputTokenTiers,omitempty" type:"Repeated"`
	// The name of the pricing model.
	//
	// example:
	//
	// qwen3-max
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The ID of the model service.
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
	// 580EF224-9647-59E7-9950-D9EBFD6A2921
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The video input cost points per second.
	//
	// example:
	//
	// 20
	VideoCostPointsPerSecond *string `json:"VideoCostPointsPerSecond,omitempty" xml:"VideoCostPointsPerSecond,omitempty"`
}

func (s CreateCostRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCostRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCostRuleResponseBody) GetAffectedModelApiIds() []*string {
	return s.AffectedModelApiIds
}

func (s *CreateCostRuleResponseBody) GetCacheCostPointsPerMillion() *string {
	return s.CacheCostPointsPerMillion
}

func (s *CreateCostRuleResponseBody) GetCacheCreationCostPointsPerMillion() *string {
	return s.CacheCreationCostPointsPerMillion
}

func (s *CreateCostRuleResponseBody) GetCostRuleId() *string {
	return s.CostRuleId
}

func (s *CreateCostRuleResponseBody) GetEffectiveEndTime() *string {
	return s.EffectiveEndTime
}

func (s *CreateCostRuleResponseBody) GetEffectiveStartTime() *string {
	return s.EffectiveStartTime
}

func (s *CreateCostRuleResponseBody) GetEffectiveTargetType() *string {
	return s.EffectiveTargetType
}

func (s *CreateCostRuleResponseBody) GetEffectiveTargetValue() *string {
	return s.EffectiveTargetValue
}

func (s *CreateCostRuleResponseBody) GetExplicitCacheCostPointsPerMillion() *string {
	return s.ExplicitCacheCostPointsPerMillion
}

func (s *CreateCostRuleResponseBody) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *CreateCostRuleResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *CreateCostRuleResponseBody) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *CreateCostRuleResponseBody) GetImageCostPointsPerCount() *string {
	return s.ImageCostPointsPerCount
}

func (s *CreateCostRuleResponseBody) GetInputCostPointsPerMillion() *string {
	return s.InputCostPointsPerMillion
}

func (s *CreateCostRuleResponseBody) GetInputTokenTier() *string {
	return s.InputTokenTier
}

func (s *CreateCostRuleResponseBody) GetInputTokenTiers() []*CreateCostRuleResponseBodyInputTokenTiers {
	return s.InputTokenTiers
}

func (s *CreateCostRuleResponseBody) GetModelName() *string {
	return s.ModelName
}

func (s *CreateCostRuleResponseBody) GetModelServiceId() *string {
	return s.ModelServiceId
}

func (s *CreateCostRuleResponseBody) GetModelServiceName() *string {
	return s.ModelServiceName
}

func (s *CreateCostRuleResponseBody) GetOutputCostPointsPerMillion() *string {
	return s.OutputCostPointsPerMillion
}

func (s *CreateCostRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCostRuleResponseBody) GetVideoCostPointsPerSecond() *string {
	return s.VideoCostPointsPerSecond
}

func (s *CreateCostRuleResponseBody) SetAffectedModelApiIds(v []*string) *CreateCostRuleResponseBody {
	s.AffectedModelApiIds = v
	return s
}

func (s *CreateCostRuleResponseBody) SetCacheCostPointsPerMillion(v string) *CreateCostRuleResponseBody {
	s.CacheCostPointsPerMillion = &v
	return s
}

func (s *CreateCostRuleResponseBody) SetCacheCreationCostPointsPerMillion(v string) *CreateCostRuleResponseBody {
	s.CacheCreationCostPointsPerMillion = &v
	return s
}

func (s *CreateCostRuleResponseBody) SetCostRuleId(v string) *CreateCostRuleResponseBody {
	s.CostRuleId = &v
	return s
}

func (s *CreateCostRuleResponseBody) SetEffectiveEndTime(v string) *CreateCostRuleResponseBody {
	s.EffectiveEndTime = &v
	return s
}

func (s *CreateCostRuleResponseBody) SetEffectiveStartTime(v string) *CreateCostRuleResponseBody {
	s.EffectiveStartTime = &v
	return s
}

func (s *CreateCostRuleResponseBody) SetEffectiveTargetType(v string) *CreateCostRuleResponseBody {
	s.EffectiveTargetType = &v
	return s
}

func (s *CreateCostRuleResponseBody) SetEffectiveTargetValue(v string) *CreateCostRuleResponseBody {
	s.EffectiveTargetValue = &v
	return s
}

func (s *CreateCostRuleResponseBody) SetExplicitCacheCostPointsPerMillion(v string) *CreateCostRuleResponseBody {
	s.ExplicitCacheCostPointsPerMillion = &v
	return s
}

func (s *CreateCostRuleResponseBody) SetGmtCreated(v string) *CreateCostRuleResponseBody {
	s.GmtCreated = &v
	return s
}

func (s *CreateCostRuleResponseBody) SetGmtModified(v string) *CreateCostRuleResponseBody {
	s.GmtModified = &v
	return s
}

func (s *CreateCostRuleResponseBody) SetGwClusterId(v string) *CreateCostRuleResponseBody {
	s.GwClusterId = &v
	return s
}

func (s *CreateCostRuleResponseBody) SetImageCostPointsPerCount(v string) *CreateCostRuleResponseBody {
	s.ImageCostPointsPerCount = &v
	return s
}

func (s *CreateCostRuleResponseBody) SetInputCostPointsPerMillion(v string) *CreateCostRuleResponseBody {
	s.InputCostPointsPerMillion = &v
	return s
}

func (s *CreateCostRuleResponseBody) SetInputTokenTier(v string) *CreateCostRuleResponseBody {
	s.InputTokenTier = &v
	return s
}

func (s *CreateCostRuleResponseBody) SetInputTokenTiers(v []*CreateCostRuleResponseBodyInputTokenTiers) *CreateCostRuleResponseBody {
	s.InputTokenTiers = v
	return s
}

func (s *CreateCostRuleResponseBody) SetModelName(v string) *CreateCostRuleResponseBody {
	s.ModelName = &v
	return s
}

func (s *CreateCostRuleResponseBody) SetModelServiceId(v string) *CreateCostRuleResponseBody {
	s.ModelServiceId = &v
	return s
}

func (s *CreateCostRuleResponseBody) SetModelServiceName(v string) *CreateCostRuleResponseBody {
	s.ModelServiceName = &v
	return s
}

func (s *CreateCostRuleResponseBody) SetOutputCostPointsPerMillion(v string) *CreateCostRuleResponseBody {
	s.OutputCostPointsPerMillion = &v
	return s
}

func (s *CreateCostRuleResponseBody) SetRequestId(v string) *CreateCostRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCostRuleResponseBody) SetVideoCostPointsPerSecond(v string) *CreateCostRuleResponseBody {
	s.VideoCostPointsPerSecond = &v
	return s
}

func (s *CreateCostRuleResponseBody) Validate() error {
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

type CreateCostRuleResponseBodyInputTokenTiers struct {
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

func (s CreateCostRuleResponseBodyInputTokenTiers) String() string {
	return dara.Prettify(s)
}

func (s CreateCostRuleResponseBodyInputTokenTiers) GoString() string {
	return s.String()
}

func (s *CreateCostRuleResponseBodyInputTokenTiers) GetInputCostPointsPerMillion() *string {
	return s.InputCostPointsPerMillion
}

func (s *CreateCostRuleResponseBodyInputTokenTiers) GetTier() *string {
	return s.Tier
}

func (s *CreateCostRuleResponseBodyInputTokenTiers) SetInputCostPointsPerMillion(v string) *CreateCostRuleResponseBodyInputTokenTiers {
	s.InputCostPointsPerMillion = &v
	return s
}

func (s *CreateCostRuleResponseBodyInputTokenTiers) SetTier(v string) *CreateCostRuleResponseBodyInputTokenTiers {
	s.Tier = &v
	return s
}

func (s *CreateCostRuleResponseBodyInputTokenTiers) Validate() error {
	return dara.Validate(s)
}
