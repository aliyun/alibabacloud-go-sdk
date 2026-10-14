// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCostRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCacheCostPointsPerMillion(v string) *CreateCostRuleRequest
	GetCacheCostPointsPerMillion() *string
	SetCacheCreationCostPointsPerMillion(v string) *CreateCostRuleRequest
	GetCacheCreationCostPointsPerMillion() *string
	SetEffectiveStartTime(v string) *CreateCostRuleRequest
	GetEffectiveStartTime() *string
	SetEffectiveTargetType(v string) *CreateCostRuleRequest
	GetEffectiveTargetType() *string
	SetEffectiveTargetValue(v string) *CreateCostRuleRequest
	GetEffectiveTargetValue() *string
	SetExplicitCacheCostPointsPerMillion(v string) *CreateCostRuleRequest
	GetExplicitCacheCostPointsPerMillion() *string
	SetGwClusterId(v string) *CreateCostRuleRequest
	GetGwClusterId() *string
	SetImageCostPointsPerCount(v string) *CreateCostRuleRequest
	GetImageCostPointsPerCount() *string
	SetInputCostPointsPerMillion(v string) *CreateCostRuleRequest
	GetInputCostPointsPerMillion() *string
	SetInputTokenTiers(v string) *CreateCostRuleRequest
	GetInputTokenTiers() *string
	SetModelName(v string) *CreateCostRuleRequest
	GetModelName() *string
	SetModelServiceId(v string) *CreateCostRuleRequest
	GetModelServiceId() *string
	SetOutputCostPointsPerMillion(v string) *CreateCostRuleRequest
	GetOutputCostPointsPerMillion() *string
	SetRegionId(v string) *CreateCostRuleRequest
	GetRegionId() *string
	SetVideoCostPointsPerSecond(v string) *CreateCostRuleRequest
	GetVideoCostPointsPerSecond() *string
}

type CreateCostRuleRequest struct {
	// The cache read cost points per million tokens. Default value: 0.
	//
	// example:
	//
	// 0
	CacheCostPointsPerMillion *string `json:"CacheCostPointsPerMillion,omitempty" xml:"CacheCostPointsPerMillion,omitempty"`
	// The cache creation cost points per million tokens.
	//
	// example:
	//
	// 400
	CacheCreationCostPointsPerMillion *string `json:"CacheCreationCostPointsPerMillion,omitempty" xml:"CacheCreationCostPointsPerMillion,omitempty"`
	// The effective start time in the yyyy-MM-dd HH:mm:ss format. If this parameter is not specified, the rule takes effect immediately.
	//
	// example:
	//
	// 2026-10-14 00:00:00
	EffectiveStartTime *string `json:"EffectiveStartTime,omitempty" xml:"EffectiveStartTime,omitempty"`
	// The effective target type. Valid values: global, consumerGroup, and consumer. Default value: global.
	//
	// example:
	//
	// global
	EffectiveTargetType *string `json:"EffectiveTargetType,omitempty" xml:"EffectiveTargetType,omitempty"`
	// The effective target value. This parameter is required when the effective target type is not global.
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
	// The gateway instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxx
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
	// The tiered price for input tokens (non-cumulative). The value is a JSON array character string. Each element contains Tier and InputCostPointsPerMillion. Valid values for Tier: "" (default tier), 0-32K, 32K-64K, 64K-128K, 128K-256K, 256K-512K, 512K-1M, and 1M+. Example: [{"Tier":"0-32K","InputCostPointsPerMillion":"1000"},{"Tier":"32K-64K","InputCostPointsPerMillion":"1800"}]. If this parameter is specified, it takes precedence over InputCostPointsPerMillion. An empty array [] is not allowed.
	//
	// example:
	//
	// [{"Tier":"0-32K","InputCostPointsPerMillion":"1000"},{"Tier":"32K-64K","InputCostPointsPerMillion":"1800"}]
	InputTokenTiers *string `json:"InputTokenTiers,omitempty" xml:"InputTokenTiers,omitempty"`
	// The name of the model, such as gpt-4 or qwen-turbo.
	//
	// This parameter is required.
	//
	// example:
	//
	// gpt-4
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The ID of the model service.
	//
	// This parameter is required.
	//
	// example:
	//
	// ms-xxxxxx
	ModelServiceId *string `json:"ModelServiceId,omitempty" xml:"ModelServiceId,omitempty"`
	// The output cost points per million tokens. Default value: 0.
	//
	// example:
	//
	// 0
	OutputCostPointsPerMillion *string `json:"OutputCostPointsPerMillion,omitempty" xml:"OutputCostPointsPerMillion,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The video input cost points per second.
	//
	// example:
	//
	// 20
	VideoCostPointsPerSecond *string `json:"VideoCostPointsPerSecond,omitempty" xml:"VideoCostPointsPerSecond,omitempty"`
}

func (s CreateCostRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCostRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateCostRuleRequest) GetCacheCostPointsPerMillion() *string {
	return s.CacheCostPointsPerMillion
}

func (s *CreateCostRuleRequest) GetCacheCreationCostPointsPerMillion() *string {
	return s.CacheCreationCostPointsPerMillion
}

func (s *CreateCostRuleRequest) GetEffectiveStartTime() *string {
	return s.EffectiveStartTime
}

func (s *CreateCostRuleRequest) GetEffectiveTargetType() *string {
	return s.EffectiveTargetType
}

func (s *CreateCostRuleRequest) GetEffectiveTargetValue() *string {
	return s.EffectiveTargetValue
}

func (s *CreateCostRuleRequest) GetExplicitCacheCostPointsPerMillion() *string {
	return s.ExplicitCacheCostPointsPerMillion
}

func (s *CreateCostRuleRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *CreateCostRuleRequest) GetImageCostPointsPerCount() *string {
	return s.ImageCostPointsPerCount
}

func (s *CreateCostRuleRequest) GetInputCostPointsPerMillion() *string {
	return s.InputCostPointsPerMillion
}

func (s *CreateCostRuleRequest) GetInputTokenTiers() *string {
	return s.InputTokenTiers
}

func (s *CreateCostRuleRequest) GetModelName() *string {
	return s.ModelName
}

func (s *CreateCostRuleRequest) GetModelServiceId() *string {
	return s.ModelServiceId
}

func (s *CreateCostRuleRequest) GetOutputCostPointsPerMillion() *string {
	return s.OutputCostPointsPerMillion
}

func (s *CreateCostRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCostRuleRequest) GetVideoCostPointsPerSecond() *string {
	return s.VideoCostPointsPerSecond
}

func (s *CreateCostRuleRequest) SetCacheCostPointsPerMillion(v string) *CreateCostRuleRequest {
	s.CacheCostPointsPerMillion = &v
	return s
}

func (s *CreateCostRuleRequest) SetCacheCreationCostPointsPerMillion(v string) *CreateCostRuleRequest {
	s.CacheCreationCostPointsPerMillion = &v
	return s
}

func (s *CreateCostRuleRequest) SetEffectiveStartTime(v string) *CreateCostRuleRequest {
	s.EffectiveStartTime = &v
	return s
}

func (s *CreateCostRuleRequest) SetEffectiveTargetType(v string) *CreateCostRuleRequest {
	s.EffectiveTargetType = &v
	return s
}

func (s *CreateCostRuleRequest) SetEffectiveTargetValue(v string) *CreateCostRuleRequest {
	s.EffectiveTargetValue = &v
	return s
}

func (s *CreateCostRuleRequest) SetExplicitCacheCostPointsPerMillion(v string) *CreateCostRuleRequest {
	s.ExplicitCacheCostPointsPerMillion = &v
	return s
}

func (s *CreateCostRuleRequest) SetGwClusterId(v string) *CreateCostRuleRequest {
	s.GwClusterId = &v
	return s
}

func (s *CreateCostRuleRequest) SetImageCostPointsPerCount(v string) *CreateCostRuleRequest {
	s.ImageCostPointsPerCount = &v
	return s
}

func (s *CreateCostRuleRequest) SetInputCostPointsPerMillion(v string) *CreateCostRuleRequest {
	s.InputCostPointsPerMillion = &v
	return s
}

func (s *CreateCostRuleRequest) SetInputTokenTiers(v string) *CreateCostRuleRequest {
	s.InputTokenTiers = &v
	return s
}

func (s *CreateCostRuleRequest) SetModelName(v string) *CreateCostRuleRequest {
	s.ModelName = &v
	return s
}

func (s *CreateCostRuleRequest) SetModelServiceId(v string) *CreateCostRuleRequest {
	s.ModelServiceId = &v
	return s
}

func (s *CreateCostRuleRequest) SetOutputCostPointsPerMillion(v string) *CreateCostRuleRequest {
	s.OutputCostPointsPerMillion = &v
	return s
}

func (s *CreateCostRuleRequest) SetRegionId(v string) *CreateCostRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCostRuleRequest) SetVideoCostPointsPerSecond(v string) *CreateCostRuleRequest {
	s.VideoCostPointsPerSecond = &v
	return s
}

func (s *CreateCostRuleRequest) Validate() error {
	return dara.Validate(s)
}
