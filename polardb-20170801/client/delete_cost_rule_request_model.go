// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCostRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCostRuleId(v string) *DeleteCostRuleRequest
	GetCostRuleId() *string
	SetGwClusterId(v string) *DeleteCostRuleRequest
	GetGwClusterId() *string
	SetRegionId(v string) *DeleteCostRuleRequest
	GetRegionId() *string
}

type DeleteCostRuleRequest struct {
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
	// pg-xxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCostRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCostRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteCostRuleRequest) GetCostRuleId() *string {
	return s.CostRuleId
}

func (s *DeleteCostRuleRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DeleteCostRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCostRuleRequest) SetCostRuleId(v string) *DeleteCostRuleRequest {
	s.CostRuleId = &v
	return s
}

func (s *DeleteCostRuleRequest) SetGwClusterId(v string) *DeleteCostRuleRequest {
	s.GwClusterId = &v
	return s
}

func (s *DeleteCostRuleRequest) SetRegionId(v string) *DeleteCostRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCostRuleRequest) Validate() error {
	return dara.Validate(s)
}
