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
	SetGwClusterId(v string) *CreateCostRuleRequest
	GetGwClusterId() *string
	SetInputCostPointsPerMillion(v string) *CreateCostRuleRequest
	GetInputCostPointsPerMillion() *string
	SetModelName(v string) *CreateCostRuleRequest
	GetModelName() *string
	SetModelServiceId(v string) *CreateCostRuleRequest
	GetModelServiceId() *string
	SetOutputCostPointsPerMillion(v string) *CreateCostRuleRequest
	GetOutputCostPointsPerMillion() *string
	SetRegionId(v string) *CreateCostRuleRequest
	GetRegionId() *string
}

type CreateCostRuleRequest struct {
	// The number of cost points per million cache tokens. The default value is 0.
	//
	// example:
	//
	// 0
	CacheCostPointsPerMillion *string `json:"CacheCostPointsPerMillion,omitempty" xml:"CacheCostPointsPerMillion,omitempty"`
	// The gateway instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// The number of cost points per million input tokens. The default value is 0.
	//
	// example:
	//
	// 0
	InputCostPointsPerMillion *string `json:"InputCostPointsPerMillion,omitempty" xml:"InputCostPointsPerMillion,omitempty"`
	// The name of the model, such as `gpt-4` or `qwen-turbo`.
	//
	// This parameter is required.
	//
	// example:
	//
	// gpt-4
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The model service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ms-xxxxxx
	ModelServiceId *string `json:"ModelServiceId,omitempty" xml:"ModelServiceId,omitempty"`
	// The number of cost points per million output tokens. The default value is 0.
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

func (s *CreateCostRuleRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *CreateCostRuleRequest) GetInputCostPointsPerMillion() *string {
	return s.InputCostPointsPerMillion
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

func (s *CreateCostRuleRequest) SetCacheCostPointsPerMillion(v string) *CreateCostRuleRequest {
	s.CacheCostPointsPerMillion = &v
	return s
}

func (s *CreateCostRuleRequest) SetGwClusterId(v string) *CreateCostRuleRequest {
	s.GwClusterId = &v
	return s
}

func (s *CreateCostRuleRequest) SetInputCostPointsPerMillion(v string) *CreateCostRuleRequest {
	s.InputCostPointsPerMillion = &v
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

func (s *CreateCostRuleRequest) Validate() error {
	return dara.Validate(s)
}
