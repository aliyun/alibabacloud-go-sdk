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
	SetCostRuleId(v string) *ModifyCostRuleRequest
	GetCostRuleId() *string
	SetGwClusterId(v string) *ModifyCostRuleRequest
	GetGwClusterId() *string
	SetInputCostPointsPerMillion(v string) *ModifyCostRuleRequest
	GetInputCostPointsPerMillion() *string
	SetModelName(v string) *ModifyCostRuleRequest
	GetModelName() *string
	SetModelServiceId(v string) *ModifyCostRuleRequest
	GetModelServiceId() *string
	SetOutputCostPointsPerMillion(v string) *ModifyCostRuleRequest
	GetOutputCostPointsPerMillion() *string
	SetRegionId(v string) *ModifyCostRuleRequest
	GetRegionId() *string
}

type ModifyCostRuleRequest struct {
	// example:
	//
	// 0
	CacheCostPointsPerMillion *string `json:"CacheCostPointsPerMillion,omitempty" xml:"CacheCostPointsPerMillion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 924d450014e64e88ac6e8486f8e990**
	CostRuleId *string `json:"CostRuleId,omitempty" xml:"CostRuleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// example:
	//
	// 0
	InputCostPointsPerMillion *string `json:"InputCostPointsPerMillion,omitempty" xml:"InputCostPointsPerMillion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// gpt-4
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ms-xxxxxx
	ModelServiceId *string `json:"ModelServiceId,omitempty" xml:"ModelServiceId,omitempty"`
	// example:
	//
	// 10
	OutputCostPointsPerMillion *string `json:"OutputCostPointsPerMillion,omitempty" xml:"OutputCostPointsPerMillion,omitempty"`
	RegionId                   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *ModifyCostRuleRequest) GetCostRuleId() *string {
	return s.CostRuleId
}

func (s *ModifyCostRuleRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *ModifyCostRuleRequest) GetInputCostPointsPerMillion() *string {
	return s.InputCostPointsPerMillion
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

func (s *ModifyCostRuleRequest) SetCacheCostPointsPerMillion(v string) *ModifyCostRuleRequest {
	s.CacheCostPointsPerMillion = &v
	return s
}

func (s *ModifyCostRuleRequest) SetCostRuleId(v string) *ModifyCostRuleRequest {
	s.CostRuleId = &v
	return s
}

func (s *ModifyCostRuleRequest) SetGwClusterId(v string) *ModifyCostRuleRequest {
	s.GwClusterId = &v
	return s
}

func (s *ModifyCostRuleRequest) SetInputCostPointsPerMillion(v string) *ModifyCostRuleRequest {
	s.InputCostPointsPerMillion = &v
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

func (s *ModifyCostRuleRequest) Validate() error {
	return dara.Validate(s)
}
