// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentServiceOrderSummaryVO interface {
	dara.Model
	String() string
	GoString() string
	SetAgentService(v string) *AgentServiceOrderSummaryVO
	GetAgentService() *string
	SetDeepResearchQuota(v int64) *AgentServiceOrderSummaryVO
	GetDeepResearchQuota() *int64
	SetModelCallQuota(v int64) *AgentServiceOrderSummaryVO
	GetModelCallQuota() *int64
	SetOrderCount(v int32) *AgentServiceOrderSummaryVO
	GetOrderCount() *int32
	SetServiceNumTotal(v int32) *AgentServiceOrderSummaryVO
	GetServiceNumTotal() *int32
	SetSkillPlanCallQuota(v int64) *AgentServiceOrderSummaryVO
	GetSkillPlanCallQuota() *int64
}

type AgentServiceOrderSummaryVO struct {
	// The name of the agent service.
	AgentService *string `json:"AgentService,omitempty" xml:"AgentService,omitempty"`
	// The deep research quota for the service.
	DeepResearchQuota *int64 `json:"DeepResearchQuota,omitempty" xml:"DeepResearchQuota,omitempty"`
	// The model call quota for the service.
	ModelCallQuota *int64 `json:"ModelCallQuota,omitempty" xml:"ModelCallQuota,omitempty"`
	// The order count for the agent service.
	OrderCount *int32 `json:"OrderCount,omitempty" xml:"OrderCount,omitempty"`
	// The total number of services.
	ServiceNumTotal *int32 `json:"ServiceNumTotal,omitempty" xml:"ServiceNumTotal,omitempty"`
	// The skill plan call quota for the service.
	SkillPlanCallQuota *int64 `json:"SkillPlanCallQuota,omitempty" xml:"SkillPlanCallQuota,omitempty"`
}

func (s AgentServiceOrderSummaryVO) String() string {
	return dara.Prettify(s)
}

func (s AgentServiceOrderSummaryVO) GoString() string {
	return s.String()
}

func (s *AgentServiceOrderSummaryVO) GetAgentService() *string {
	return s.AgentService
}

func (s *AgentServiceOrderSummaryVO) GetDeepResearchQuota() *int64 {
	return s.DeepResearchQuota
}

func (s *AgentServiceOrderSummaryVO) GetModelCallQuota() *int64 {
	return s.ModelCallQuota
}

func (s *AgentServiceOrderSummaryVO) GetOrderCount() *int32 {
	return s.OrderCount
}

func (s *AgentServiceOrderSummaryVO) GetServiceNumTotal() *int32 {
	return s.ServiceNumTotal
}

func (s *AgentServiceOrderSummaryVO) GetSkillPlanCallQuota() *int64 {
	return s.SkillPlanCallQuota
}

func (s *AgentServiceOrderSummaryVO) SetAgentService(v string) *AgentServiceOrderSummaryVO {
	s.AgentService = &v
	return s
}

func (s *AgentServiceOrderSummaryVO) SetDeepResearchQuota(v int64) *AgentServiceOrderSummaryVO {
	s.DeepResearchQuota = &v
	return s
}

func (s *AgentServiceOrderSummaryVO) SetModelCallQuota(v int64) *AgentServiceOrderSummaryVO {
	s.ModelCallQuota = &v
	return s
}

func (s *AgentServiceOrderSummaryVO) SetOrderCount(v int32) *AgentServiceOrderSummaryVO {
	s.OrderCount = &v
	return s
}

func (s *AgentServiceOrderSummaryVO) SetServiceNumTotal(v int32) *AgentServiceOrderSummaryVO {
	s.ServiceNumTotal = &v
	return s
}

func (s *AgentServiceOrderSummaryVO) SetSkillPlanCallQuota(v int64) *AgentServiceOrderSummaryVO {
	s.SkillPlanCallQuota = &v
	return s
}

func (s *AgentServiceOrderSummaryVO) Validate() error {
	return dara.Validate(s)
}
