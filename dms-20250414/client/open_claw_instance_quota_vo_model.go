// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenClawInstanceQuotaVO interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunAccountUid(v string) *OpenClawInstanceQuotaVO
	GetAliyunAccountUid() *string
	SetDeepResearchCallQuota(v string) *OpenClawInstanceQuotaVO
	GetDeepResearchCallQuota() *string
	SetDeepResearchCallUsed(v string) *OpenClawInstanceQuotaVO
	GetDeepResearchCallUsed() *string
	SetInstanceGmtCreate(v string) *OpenClawInstanceQuotaVO
	GetInstanceGmtCreate() *string
	SetInstanceId(v string) *OpenClawInstanceQuotaVO
	GetInstanceId() *string
	SetInstanceName(v string) *OpenClawInstanceQuotaVO
	GetInstanceName() *string
	SetLastMeteringTime(v string) *OpenClawInstanceQuotaVO
	GetLastMeteringTime() *string
	SetModelCallQuota(v string) *OpenClawInstanceQuotaVO
	GetModelCallQuota() *string
	SetModelCallUsed(v string) *OpenClawInstanceQuotaVO
	GetModelCallUsed() *string
	SetRefreshDay(v string) *OpenClawInstanceQuotaVO
	GetRefreshDay() *string
	SetSkillPlanCallQuota(v string) *OpenClawInstanceQuotaVO
	GetSkillPlanCallQuota() *string
	SetSkillPlanCallUsed(v string) *OpenClawInstanceQuotaVO
	GetSkillPlanCallUsed() *string
}

type OpenClawInstanceQuotaVO struct {
	AliyunAccountUid      *string `json:"AliyunAccountUid,omitempty" xml:"AliyunAccountUid,omitempty"`
	DeepResearchCallQuota *string `json:"DeepResearchCallQuota,omitempty" xml:"DeepResearchCallQuota,omitempty"`
	DeepResearchCallUsed  *string `json:"DeepResearchCallUsed,omitempty" xml:"DeepResearchCallUsed,omitempty"`
	InstanceGmtCreate     *string `json:"InstanceGmtCreate,omitempty" xml:"InstanceGmtCreate,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName          *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	LastMeteringTime      *string `json:"LastMeteringTime,omitempty" xml:"LastMeteringTime,omitempty"`
	ModelCallQuota        *string `json:"ModelCallQuota,omitempty" xml:"ModelCallQuota,omitempty"`
	ModelCallUsed         *string `json:"ModelCallUsed,omitempty" xml:"ModelCallUsed,omitempty"`
	RefreshDay            *string `json:"RefreshDay,omitempty" xml:"RefreshDay,omitempty"`
	SkillPlanCallQuota    *string `json:"SkillPlanCallQuota,omitempty" xml:"SkillPlanCallQuota,omitempty"`
	SkillPlanCallUsed     *string `json:"SkillPlanCallUsed,omitempty" xml:"SkillPlanCallUsed,omitempty"`
}

func (s OpenClawInstanceQuotaVO) String() string {
	return dara.Prettify(s)
}

func (s OpenClawInstanceQuotaVO) GoString() string {
	return s.String()
}

func (s *OpenClawInstanceQuotaVO) GetAliyunAccountUid() *string {
	return s.AliyunAccountUid
}

func (s *OpenClawInstanceQuotaVO) GetDeepResearchCallQuota() *string {
	return s.DeepResearchCallQuota
}

func (s *OpenClawInstanceQuotaVO) GetDeepResearchCallUsed() *string {
	return s.DeepResearchCallUsed
}

func (s *OpenClawInstanceQuotaVO) GetInstanceGmtCreate() *string {
	return s.InstanceGmtCreate
}

func (s *OpenClawInstanceQuotaVO) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OpenClawInstanceQuotaVO) GetInstanceName() *string {
	return s.InstanceName
}

func (s *OpenClawInstanceQuotaVO) GetLastMeteringTime() *string {
	return s.LastMeteringTime
}

func (s *OpenClawInstanceQuotaVO) GetModelCallQuota() *string {
	return s.ModelCallQuota
}

func (s *OpenClawInstanceQuotaVO) GetModelCallUsed() *string {
	return s.ModelCallUsed
}

func (s *OpenClawInstanceQuotaVO) GetRefreshDay() *string {
	return s.RefreshDay
}

func (s *OpenClawInstanceQuotaVO) GetSkillPlanCallQuota() *string {
	return s.SkillPlanCallQuota
}

func (s *OpenClawInstanceQuotaVO) GetSkillPlanCallUsed() *string {
	return s.SkillPlanCallUsed
}

func (s *OpenClawInstanceQuotaVO) SetAliyunAccountUid(v string) *OpenClawInstanceQuotaVO {
	s.AliyunAccountUid = &v
	return s
}

func (s *OpenClawInstanceQuotaVO) SetDeepResearchCallQuota(v string) *OpenClawInstanceQuotaVO {
	s.DeepResearchCallQuota = &v
	return s
}

func (s *OpenClawInstanceQuotaVO) SetDeepResearchCallUsed(v string) *OpenClawInstanceQuotaVO {
	s.DeepResearchCallUsed = &v
	return s
}

func (s *OpenClawInstanceQuotaVO) SetInstanceGmtCreate(v string) *OpenClawInstanceQuotaVO {
	s.InstanceGmtCreate = &v
	return s
}

func (s *OpenClawInstanceQuotaVO) SetInstanceId(v string) *OpenClawInstanceQuotaVO {
	s.InstanceId = &v
	return s
}

func (s *OpenClawInstanceQuotaVO) SetInstanceName(v string) *OpenClawInstanceQuotaVO {
	s.InstanceName = &v
	return s
}

func (s *OpenClawInstanceQuotaVO) SetLastMeteringTime(v string) *OpenClawInstanceQuotaVO {
	s.LastMeteringTime = &v
	return s
}

func (s *OpenClawInstanceQuotaVO) SetModelCallQuota(v string) *OpenClawInstanceQuotaVO {
	s.ModelCallQuota = &v
	return s
}

func (s *OpenClawInstanceQuotaVO) SetModelCallUsed(v string) *OpenClawInstanceQuotaVO {
	s.ModelCallUsed = &v
	return s
}

func (s *OpenClawInstanceQuotaVO) SetRefreshDay(v string) *OpenClawInstanceQuotaVO {
	s.RefreshDay = &v
	return s
}

func (s *OpenClawInstanceQuotaVO) SetSkillPlanCallQuota(v string) *OpenClawInstanceQuotaVO {
	s.SkillPlanCallQuota = &v
	return s
}

func (s *OpenClawInstanceQuotaVO) SetSkillPlanCallUsed(v string) *OpenClawInstanceQuotaVO {
	s.SkillPlanCallUsed = &v
	return s
}

func (s *OpenClawInstanceQuotaVO) Validate() error {
	return dara.Validate(s)
}
