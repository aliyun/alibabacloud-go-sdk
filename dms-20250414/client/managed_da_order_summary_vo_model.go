// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManagedDaOrderSummaryVO interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableQuota(v int32) *ManagedDaOrderSummaryVO
	GetAvailableQuota() *int32
	SetTotalQuota(v int32) *ManagedDaOrderSummaryVO
	GetTotalQuota() *int32
	SetTrialExpireTime(v string) *ManagedDaOrderSummaryVO
	GetTrialExpireTime() *string
	SetTrialUsed(v bool) *ManagedDaOrderSummaryVO
	GetTrialUsed() *bool
	SetUsedQuota(v int32) *ManagedDaOrderSummaryVO
	GetUsedQuota() *int32
	SetValidOrderCount(v int32) *ManagedDaOrderSummaryVO
	GetValidOrderCount() *int32
}

type ManagedDaOrderSummaryVO struct {
	AvailableQuota  *int32  `json:"availableQuota,omitempty" xml:"availableQuota,omitempty"`
	TotalQuota      *int32  `json:"totalQuota,omitempty" xml:"totalQuota,omitempty"`
	TrialExpireTime *string `json:"trialExpireTime,omitempty" xml:"trialExpireTime,omitempty"`
	TrialUsed       *bool   `json:"trialUsed,omitempty" xml:"trialUsed,omitempty"`
	UsedQuota       *int32  `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
	ValidOrderCount *int32  `json:"validOrderCount,omitempty" xml:"validOrderCount,omitempty"`
}

func (s ManagedDaOrderSummaryVO) String() string {
	return dara.Prettify(s)
}

func (s ManagedDaOrderSummaryVO) GoString() string {
	return s.String()
}

func (s *ManagedDaOrderSummaryVO) GetAvailableQuota() *int32 {
	return s.AvailableQuota
}

func (s *ManagedDaOrderSummaryVO) GetTotalQuota() *int32 {
	return s.TotalQuota
}

func (s *ManagedDaOrderSummaryVO) GetTrialExpireTime() *string {
	return s.TrialExpireTime
}

func (s *ManagedDaOrderSummaryVO) GetTrialUsed() *bool {
	return s.TrialUsed
}

func (s *ManagedDaOrderSummaryVO) GetUsedQuota() *int32 {
	return s.UsedQuota
}

func (s *ManagedDaOrderSummaryVO) GetValidOrderCount() *int32 {
	return s.ValidOrderCount
}

func (s *ManagedDaOrderSummaryVO) SetAvailableQuota(v int32) *ManagedDaOrderSummaryVO {
	s.AvailableQuota = &v
	return s
}

func (s *ManagedDaOrderSummaryVO) SetTotalQuota(v int32) *ManagedDaOrderSummaryVO {
	s.TotalQuota = &v
	return s
}

func (s *ManagedDaOrderSummaryVO) SetTrialExpireTime(v string) *ManagedDaOrderSummaryVO {
	s.TrialExpireTime = &v
	return s
}

func (s *ManagedDaOrderSummaryVO) SetTrialUsed(v bool) *ManagedDaOrderSummaryVO {
	s.TrialUsed = &v
	return s
}

func (s *ManagedDaOrderSummaryVO) SetUsedQuota(v int32) *ManagedDaOrderSummaryVO {
	s.UsedQuota = &v
	return s
}

func (s *ManagedDaOrderSummaryVO) SetValidOrderCount(v int32) *ManagedDaOrderSummaryVO {
	s.ValidOrderCount = &v
	return s
}

func (s *ManagedDaOrderSummaryVO) Validate() error {
	return dara.Validate(s)
}
