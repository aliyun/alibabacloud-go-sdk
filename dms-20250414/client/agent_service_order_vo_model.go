// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentServiceOrderVO interface {
	dara.Model
	String() string
	GoString() string
	SetAgentService(v string) *AgentServiceOrderVO
	GetAgentService() *string
	SetDeepResearchQuota(v int32) *AgentServiceOrderVO
	GetDeepResearchQuota() *int32
	SetDeepResearchUsed(v int32) *AgentServiceOrderVO
	GetDeepResearchUsed() *int32
	SetExpireTime(v string) *AgentServiceOrderVO
	GetExpireTime() *string
	SetGmtCreate(v string) *AgentServiceOrderVO
	GetGmtCreate() *string
	SetGmtModified(v string) *AgentServiceOrderVO
	GetGmtModified() *string
	SetModelCallQuota(v int32) *AgentServiceOrderVO
	GetModelCallQuota() *int32
	SetModelCallUsed(v int32) *AgentServiceOrderVO
	GetModelCallUsed() *int32
	SetOrderInstanceId(v string) *AgentServiceOrderVO
	GetOrderInstanceId() *string
	SetServiceNum(v int32) *AgentServiceOrderVO
	GetServiceNum() *int32
	SetSkillPlanCallQuota(v int32) *AgentServiceOrderVO
	GetSkillPlanCallQuota() *int32
	SetSkillPlanCallUsed(v int32) *AgentServiceOrderVO
	GetSkillPlanCallUsed() *int32
	SetStatus(v string) *AgentServiceOrderVO
	GetStatus() *string
}

type AgentServiceOrderVO struct {
	// The agent service.
	AgentService *string `json:"AgentService,omitempty" xml:"AgentService,omitempty"`
	// The deep research quota.
	DeepResearchQuota *int32 `json:"DeepResearchQuota,omitempty" xml:"DeepResearchQuota,omitempty"`
	// The quantity of deep research used.
	DeepResearchUsed *int32 `json:"DeepResearchUsed,omitempty" xml:"DeepResearchUsed,omitempty"`
	// The service expiration time, specified in UTC.
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The order creation time, specified in UTC.
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The order modification time, specified in UTC.
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The model call quota.
	ModelCallQuota *int32 `json:"ModelCallQuota,omitempty" xml:"ModelCallQuota,omitempty"`
	// The number of model calls used.
	ModelCallUsed *int32 `json:"ModelCallUsed,omitempty" xml:"ModelCallUsed,omitempty"`
	// The order instance ID.
	OrderInstanceId *string `json:"OrderInstanceId,omitempty" xml:"OrderInstanceId,omitempty"`
	// The number of services.
	ServiceNum *int32 `json:"ServiceNum,omitempty" xml:"ServiceNum,omitempty"`
	// The skill plan call quota.
	SkillPlanCallQuota *int32 `json:"SkillPlanCallQuota,omitempty" xml:"SkillPlanCallQuota,omitempty"`
	// The number of skill plan calls used.
	SkillPlanCallUsed *int32 `json:"SkillPlanCallUsed,omitempty" xml:"SkillPlanCallUsed,omitempty"`
	// The order status.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s AgentServiceOrderVO) String() string {
	return dara.Prettify(s)
}

func (s AgentServiceOrderVO) GoString() string {
	return s.String()
}

func (s *AgentServiceOrderVO) GetAgentService() *string {
	return s.AgentService
}

func (s *AgentServiceOrderVO) GetDeepResearchQuota() *int32 {
	return s.DeepResearchQuota
}

func (s *AgentServiceOrderVO) GetDeepResearchUsed() *int32 {
	return s.DeepResearchUsed
}

func (s *AgentServiceOrderVO) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *AgentServiceOrderVO) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *AgentServiceOrderVO) GetGmtModified() *string {
	return s.GmtModified
}

func (s *AgentServiceOrderVO) GetModelCallQuota() *int32 {
	return s.ModelCallQuota
}

func (s *AgentServiceOrderVO) GetModelCallUsed() *int32 {
	return s.ModelCallUsed
}

func (s *AgentServiceOrderVO) GetOrderInstanceId() *string {
	return s.OrderInstanceId
}

func (s *AgentServiceOrderVO) GetServiceNum() *int32 {
	return s.ServiceNum
}

func (s *AgentServiceOrderVO) GetSkillPlanCallQuota() *int32 {
	return s.SkillPlanCallQuota
}

func (s *AgentServiceOrderVO) GetSkillPlanCallUsed() *int32 {
	return s.SkillPlanCallUsed
}

func (s *AgentServiceOrderVO) GetStatus() *string {
	return s.Status
}

func (s *AgentServiceOrderVO) SetAgentService(v string) *AgentServiceOrderVO {
	s.AgentService = &v
	return s
}

func (s *AgentServiceOrderVO) SetDeepResearchQuota(v int32) *AgentServiceOrderVO {
	s.DeepResearchQuota = &v
	return s
}

func (s *AgentServiceOrderVO) SetDeepResearchUsed(v int32) *AgentServiceOrderVO {
	s.DeepResearchUsed = &v
	return s
}

func (s *AgentServiceOrderVO) SetExpireTime(v string) *AgentServiceOrderVO {
	s.ExpireTime = &v
	return s
}

func (s *AgentServiceOrderVO) SetGmtCreate(v string) *AgentServiceOrderVO {
	s.GmtCreate = &v
	return s
}

func (s *AgentServiceOrderVO) SetGmtModified(v string) *AgentServiceOrderVO {
	s.GmtModified = &v
	return s
}

func (s *AgentServiceOrderVO) SetModelCallQuota(v int32) *AgentServiceOrderVO {
	s.ModelCallQuota = &v
	return s
}

func (s *AgentServiceOrderVO) SetModelCallUsed(v int32) *AgentServiceOrderVO {
	s.ModelCallUsed = &v
	return s
}

func (s *AgentServiceOrderVO) SetOrderInstanceId(v string) *AgentServiceOrderVO {
	s.OrderInstanceId = &v
	return s
}

func (s *AgentServiceOrderVO) SetServiceNum(v int32) *AgentServiceOrderVO {
	s.ServiceNum = &v
	return s
}

func (s *AgentServiceOrderVO) SetSkillPlanCallQuota(v int32) *AgentServiceOrderVO {
	s.SkillPlanCallQuota = &v
	return s
}

func (s *AgentServiceOrderVO) SetSkillPlanCallUsed(v int32) *AgentServiceOrderVO {
	s.SkillPlanCallUsed = &v
	return s
}

func (s *AgentServiceOrderVO) SetStatus(v string) *AgentServiceOrderVO {
	s.Status = &v
	return s
}

func (s *AgentServiceOrderVO) Validate() error {
	return dara.Validate(s)
}
