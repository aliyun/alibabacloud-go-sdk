// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCostRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCostRuleId(v string) *DescribeCostRulesRequest
	GetCostRuleId() *string
	SetEffectiveTargetType(v string) *DescribeCostRulesRequest
	GetEffectiveTargetType() *string
	SetEffectiveTargetValue(v string) *DescribeCostRulesRequest
	GetEffectiveTargetValue() *string
	SetGwClusterId(v string) *DescribeCostRulesRequest
	GetGwClusterId() *string
	SetInputTokenTier(v string) *DescribeCostRulesRequest
	GetInputTokenTier() *string
	SetModelName(v string) *DescribeCostRulesRequest
	GetModelName() *string
	SetModelServiceId(v string) *DescribeCostRulesRequest
	GetModelServiceId() *string
	SetPageNumber(v int32) *DescribeCostRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCostRulesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeCostRulesRequest
	GetRegionId() *string
}

type DescribeCostRulesRequest struct {
	// The ID of the cost rule.
	//
	// example:
	//
	// cr-xxx
	CostRuleId *string `json:"CostRuleId,omitempty" xml:"CostRuleId,omitempty"`
	// The effective target type used to filter the results. Valid values: global, consumerGroup, and consumer.
	//
	// example:
	//
	// global
	EffectiveTargetType *string `json:"EffectiveTargetType,omitempty" xml:"EffectiveTargetType,omitempty"`
	// The effective target value used to filter the results.
	//
	// example:
	//
	// user
	EffectiveTargetValue *string `json:"EffectiveTargetValue,omitempty" xml:"EffectiveTargetValue,omitempty"`
	// The gateway instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// The input token tier used to filter the results. Returns rules that include this tier. Valid values: 0-32K, 32K-64K, 64K-128K, 128K-256K, 256K-512K, 512K-1M, and 1M+.
	//
	// example:
	//
	// 0-32K
	InputTokenTier *string `json:"InputTokenTier,omitempty" xml:"InputTokenTier,omitempty"`
	// The name of the model, such as gpt-4 or qwen-turbo.
	//
	// example:
	//
	// qwen3-max
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The ID of the model service.
	//
	// example:
	//
	// ms-xxxxxx
	ModelServiceId *string `json:"ModelServiceId,omitempty" xml:"ModelServiceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCostRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCostRulesRequest) GetCostRuleId() *string {
	return s.CostRuleId
}

func (s *DescribeCostRulesRequest) GetEffectiveTargetType() *string {
	return s.EffectiveTargetType
}

func (s *DescribeCostRulesRequest) GetEffectiveTargetValue() *string {
	return s.EffectiveTargetValue
}

func (s *DescribeCostRulesRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DescribeCostRulesRequest) GetInputTokenTier() *string {
	return s.InputTokenTier
}

func (s *DescribeCostRulesRequest) GetModelName() *string {
	return s.ModelName
}

func (s *DescribeCostRulesRequest) GetModelServiceId() *string {
	return s.ModelServiceId
}

func (s *DescribeCostRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCostRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCostRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCostRulesRequest) SetCostRuleId(v string) *DescribeCostRulesRequest {
	s.CostRuleId = &v
	return s
}

func (s *DescribeCostRulesRequest) SetEffectiveTargetType(v string) *DescribeCostRulesRequest {
	s.EffectiveTargetType = &v
	return s
}

func (s *DescribeCostRulesRequest) SetEffectiveTargetValue(v string) *DescribeCostRulesRequest {
	s.EffectiveTargetValue = &v
	return s
}

func (s *DescribeCostRulesRequest) SetGwClusterId(v string) *DescribeCostRulesRequest {
	s.GwClusterId = &v
	return s
}

func (s *DescribeCostRulesRequest) SetInputTokenTier(v string) *DescribeCostRulesRequest {
	s.InputTokenTier = &v
	return s
}

func (s *DescribeCostRulesRequest) SetModelName(v string) *DescribeCostRulesRequest {
	s.ModelName = &v
	return s
}

func (s *DescribeCostRulesRequest) SetModelServiceId(v string) *DescribeCostRulesRequest {
	s.ModelServiceId = &v
	return s
}

func (s *DescribeCostRulesRequest) SetPageNumber(v int32) *DescribeCostRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCostRulesRequest) SetPageSize(v int32) *DescribeCostRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCostRulesRequest) SetRegionId(v string) *DescribeCostRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCostRulesRequest) Validate() error {
	return dara.Validate(s)
}
