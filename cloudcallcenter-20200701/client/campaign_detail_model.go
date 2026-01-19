// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCampaignDetail interface {
	dara.Model
	String() string
	GoString() string
	SetActualEndTime(v int64) *CampaignDetail
	GetActualEndTime() *int64
	SetActualStartTime(v int64) *CampaignDetail
	GetActualStartTime() *int64
	SetCasesAborted(v int64) *CampaignDetail
	GetCasesAborted() *int64
	SetCasesConnected(v int64) *CampaignDetail
	GetCasesConnected() *int64
	SetCasesUncompleted(v int64) *CampaignDetail
	GetCasesUncompleted() *int64
	SetCompletedRate(v int64) *CampaignDetail
	GetCompletedRate() *int64
	SetCreateTime(v int64) *CampaignDetail
	GetCreateTime() *int64
	SetId(v string) *CampaignDetail
	GetId() *string
	SetMaxAttemptCount(v int64) *CampaignDetail
	GetMaxAttemptCount() *int64
	SetMinAttemptInterval(v int64) *CampaignDetail
	GetMinAttemptInterval() *int64
	SetName(v string) *CampaignDetail
	GetName() *string
	SetPlanedEndTime(v int64) *CampaignDetail
	GetPlanedEndTime() *int64
	SetPlanedStartTime(v int64) *CampaignDetail
	GetPlanedStartTime() *int64
	SetQueueName(v string) *CampaignDetail
	GetQueueName() *string
	SetState(v string) *CampaignDetail
	GetState() *string
	SetTotalCases(v int64) *CampaignDetail
	GetTotalCases() *int64
	SetUpdateTime(v int64) *CampaignDetail
	GetUpdateTime() *int64
}

type CampaignDetail struct {
	ActualEndTime      *int64  `json:"ActualEndTime,omitempty" xml:"ActualEndTime,omitempty"`
	ActualStartTime    *int64  `json:"ActualStartTime,omitempty" xml:"ActualStartTime,omitempty"`
	CasesAborted       *int64  `json:"CasesAborted,omitempty" xml:"CasesAborted,omitempty"`
	CasesConnected     *int64  `json:"CasesConnected,omitempty" xml:"CasesConnected,omitempty"`
	CasesUncompleted   *int64  `json:"CasesUncompleted,omitempty" xml:"CasesUncompleted,omitempty"`
	CompletedRate      *int64  `json:"CompletedRate,omitempty" xml:"CompletedRate,omitempty"`
	CreateTime         *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Id                 *string `json:"Id,omitempty" xml:"Id,omitempty"`
	MaxAttemptCount    *int64  `json:"MaxAttemptCount,omitempty" xml:"MaxAttemptCount,omitempty"`
	MinAttemptInterval *int64  `json:"MinAttemptInterval,omitempty" xml:"MinAttemptInterval,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PlanedEndTime      *int64  `json:"PlanedEndTime,omitempty" xml:"PlanedEndTime,omitempty"`
	PlanedStartTime    *int64  `json:"PlanedStartTime,omitempty" xml:"PlanedStartTime,omitempty"`
	QueueName          *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	State              *string `json:"State,omitempty" xml:"State,omitempty"`
	TotalCases         *int64  `json:"TotalCases,omitempty" xml:"TotalCases,omitempty"`
	UpdateTime         *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s CampaignDetail) String() string {
	return dara.Prettify(s)
}

func (s CampaignDetail) GoString() string {
	return s.String()
}

func (s *CampaignDetail) GetActualEndTime() *int64 {
	return s.ActualEndTime
}

func (s *CampaignDetail) GetActualStartTime() *int64 {
	return s.ActualStartTime
}

func (s *CampaignDetail) GetCasesAborted() *int64 {
	return s.CasesAborted
}

func (s *CampaignDetail) GetCasesConnected() *int64 {
	return s.CasesConnected
}

func (s *CampaignDetail) GetCasesUncompleted() *int64 {
	return s.CasesUncompleted
}

func (s *CampaignDetail) GetCompletedRate() *int64 {
	return s.CompletedRate
}

func (s *CampaignDetail) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CampaignDetail) GetId() *string {
	return s.Id
}

func (s *CampaignDetail) GetMaxAttemptCount() *int64 {
	return s.MaxAttemptCount
}

func (s *CampaignDetail) GetMinAttemptInterval() *int64 {
	return s.MinAttemptInterval
}

func (s *CampaignDetail) GetName() *string {
	return s.Name
}

func (s *CampaignDetail) GetPlanedEndTime() *int64 {
	return s.PlanedEndTime
}

func (s *CampaignDetail) GetPlanedStartTime() *int64 {
	return s.PlanedStartTime
}

func (s *CampaignDetail) GetQueueName() *string {
	return s.QueueName
}

func (s *CampaignDetail) GetState() *string {
	return s.State
}

func (s *CampaignDetail) GetTotalCases() *int64 {
	return s.TotalCases
}

func (s *CampaignDetail) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *CampaignDetail) SetActualEndTime(v int64) *CampaignDetail {
	s.ActualEndTime = &v
	return s
}

func (s *CampaignDetail) SetActualStartTime(v int64) *CampaignDetail {
	s.ActualStartTime = &v
	return s
}

func (s *CampaignDetail) SetCasesAborted(v int64) *CampaignDetail {
	s.CasesAborted = &v
	return s
}

func (s *CampaignDetail) SetCasesConnected(v int64) *CampaignDetail {
	s.CasesConnected = &v
	return s
}

func (s *CampaignDetail) SetCasesUncompleted(v int64) *CampaignDetail {
	s.CasesUncompleted = &v
	return s
}

func (s *CampaignDetail) SetCompletedRate(v int64) *CampaignDetail {
	s.CompletedRate = &v
	return s
}

func (s *CampaignDetail) SetCreateTime(v int64) *CampaignDetail {
	s.CreateTime = &v
	return s
}

func (s *CampaignDetail) SetId(v string) *CampaignDetail {
	s.Id = &v
	return s
}

func (s *CampaignDetail) SetMaxAttemptCount(v int64) *CampaignDetail {
	s.MaxAttemptCount = &v
	return s
}

func (s *CampaignDetail) SetMinAttemptInterval(v int64) *CampaignDetail {
	s.MinAttemptInterval = &v
	return s
}

func (s *CampaignDetail) SetName(v string) *CampaignDetail {
	s.Name = &v
	return s
}

func (s *CampaignDetail) SetPlanedEndTime(v int64) *CampaignDetail {
	s.PlanedEndTime = &v
	return s
}

func (s *CampaignDetail) SetPlanedStartTime(v int64) *CampaignDetail {
	s.PlanedStartTime = &v
	return s
}

func (s *CampaignDetail) SetQueueName(v string) *CampaignDetail {
	s.QueueName = &v
	return s
}

func (s *CampaignDetail) SetState(v string) *CampaignDetail {
	s.State = &v
	return s
}

func (s *CampaignDetail) SetTotalCases(v int64) *CampaignDetail {
	s.TotalCases = &v
	return s
}

func (s *CampaignDetail) SetUpdateTime(v int64) *CampaignDetail {
	s.UpdateTime = &v
	return s
}

func (s *CampaignDetail) Validate() error {
	return dara.Validate(s)
}
