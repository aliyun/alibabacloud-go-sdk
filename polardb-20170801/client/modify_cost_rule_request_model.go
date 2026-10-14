// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCostRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCacheCostPointsPerMillion(v string) *ModifyCostRuleRequest
	GetCacheCostPointsPerMillion() *string
	SetCacheCreationCostPointsPerMillion(v string) *ModifyCostRuleRequest
	GetCacheCreationCostPointsPerMillion() *string
	SetCostRuleId(v string) *ModifyCostRuleRequest
	GetCostRuleId() *string
	SetExplicitCacheCostPointsPerMillion(v string) *ModifyCostRuleRequest
	GetExplicitCacheCostPointsPerMillion() *string
	SetGwClusterId(v string) *ModifyCostRuleRequest
	GetGwClusterId() *string
	SetImageCostPointsPerCount(v string) *ModifyCostRuleRequest
	GetImageCostPointsPerCount() *string
	SetInputCostPointsPerMillion(v string) *ModifyCostRuleRequest
	GetInputCostPointsPerMillion() *string
	SetInputTokenTiers(v string) *ModifyCostRuleRequest
	GetInputTokenTiers() *string
	SetModelName(v string) *ModifyCostRuleRequest
	GetModelName() *string
	SetModelServiceId(v string) *ModifyCostRuleRequest
	GetModelServiceId() *string
	SetOutputCostPointsPerMillion(v string) *ModifyCostRuleRequest
	GetOutputCostPointsPerMillion() *string
	SetRegionId(v string) *ModifyCostRuleRequest
	GetRegionId() *string
	SetVideoCostPointsPerSecond(v string) *ModifyCostRuleRequest
	GetVideoCostPointsPerSecond() *string
}

type ModifyCostRuleRequest struct {
	// The cache cost points per million tokens. Default value: 0.
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
	// The ID of the cost rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 924d450014e64e88ac6e8486f8e990**
	CostRuleId *string `json:"CostRuleId,omitempty" xml:"CostRuleId,omitempty"`
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
	// The input token tiered price (non-cumulative) specified as a JSON array character string. Each element contains Tier and InputCostPointsPerMillion. Valid values for Tier: "" (default tier), 0-32K, 32K-64K, 64K-128K, 128K-256K, 256K-512K, 512K-1M, and 1M+. Example: [{"Tier":"0-32K","InputCostPointsPerMillion":"1200"}]. A non-empty value replaces all existing tiers. If this parameter is not specified or is set to an empty string, the original tiers are retained. An empty array [] is not allowed.
	//
	// example:
	//
	// [{"Tier":"0-32K","InputCostPointsPerMillion":"1200"}]
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
	// 10
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

func (s ModifyCostRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCostRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyCostRuleRequest) GetCacheCostPointsPerMillion() *string {
	return s.CacheCostPointsPerMillion
}

func (s *ModifyCostRuleRequest) GetCacheCreationCostPointsPerMillion() *string {
	return s.CacheCreationCostPointsPerMillion
}

func (s *ModifyCostRuleRequest) GetCostRuleId() *string {
	return s.CostRuleId
}

func (s *ModifyCostRuleRequest) GetExplicitCacheCostPointsPerMillion() *string {
	return s.ExplicitCacheCostPointsPerMillion
}

func (s *ModifyCostRuleRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *ModifyCostRuleRequest) GetImageCostPointsPerCount() *string {
	return s.ImageCostPointsPerCount
}

func (s *ModifyCostRuleRequest) GetInputCostPointsPerMillion() *string {
	return s.InputCostPointsPerMillion
}

func (s *ModifyCostRuleRequest) GetInputTokenTiers() *string {
	return s.InputTokenTiers
}

func (s *ModifyCostRuleRequest) GetModelName() *string {
	return s.ModelName
}

func (s *ModifyCostRuleRequest) GetModelServiceId() *string {
	return s.ModelServiceId
}

func (s *ModifyCostRuleRequest) GetOutputCostPointsPerMillion() *string {
	return s.OutputCostPointsPerMillion
}

func (s *ModifyCostRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCostRuleRequest) GetVideoCostPointsPerSecond() *string {
	return s.VideoCostPointsPerSecond
}

func (s *ModifyCostRuleRequest) SetCacheCostPointsPerMillion(v string) *ModifyCostRuleRequest {
	s.CacheCostPointsPerMillion = &v
	return s
}

func (s *ModifyCostRuleRequest) SetCacheCreationCostPointsPerMillion(v string) *ModifyCostRuleRequest {
	s.CacheCreationCostPointsPerMillion = &v
	return s
}

func (s *ModifyCostRuleRequest) SetCostRuleId(v string) *ModifyCostRuleRequest {
	s.CostRuleId = &v
	return s
}

func (s *ModifyCostRuleRequest) SetExplicitCacheCostPointsPerMillion(v string) *ModifyCostRuleRequest {
	s.ExplicitCacheCostPointsPerMillion = &v
	return s
}

func (s *ModifyCostRuleRequest) SetGwClusterId(v string) *ModifyCostRuleRequest {
	s.GwClusterId = &v
	return s
}

func (s *ModifyCostRuleRequest) SetImageCostPointsPerCount(v string) *ModifyCostRuleRequest {
	s.ImageCostPointsPerCount = &v
	return s
}

func (s *ModifyCostRuleRequest) SetInputCostPointsPerMillion(v string) *ModifyCostRuleRequest {
	s.InputCostPointsPerMillion = &v
	return s
}

func (s *ModifyCostRuleRequest) SetInputTokenTiers(v string) *ModifyCostRuleRequest {
	s.InputTokenTiers = &v
	return s
}

func (s *ModifyCostRuleRequest) SetModelName(v string) *ModifyCostRuleRequest {
	s.ModelName = &v
	return s
}

func (s *ModifyCostRuleRequest) SetModelServiceId(v string) *ModifyCostRuleRequest {
	s.ModelServiceId = &v
	return s
}

func (s *ModifyCostRuleRequest) SetOutputCostPointsPerMillion(v string) *ModifyCostRuleRequest {
	s.OutputCostPointsPerMillion = &v
	return s
}

func (s *ModifyCostRuleRequest) SetRegionId(v string) *ModifyCostRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCostRuleRequest) SetVideoCostPointsPerSecond(v string) *ModifyCostRuleRequest {
	s.VideoCostPointsPerSecond = &v
	return s
}

func (s *ModifyCostRuleRequest) Validate() error {
	return dara.Validate(s)
}
