// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPromotionActivity interface {
	dara.Model
	String() string
	GoString() string
	SetActivityCode(v string) *PromotionActivity
	GetActivityCode() *string
	SetActivityName(v string) *PromotionActivity
	GetActivityName() *string
	SetActivityType(v string) *PromotionActivity
	GetActivityType() *string
	SetConsumedQuota(v int64) *PromotionActivity
	GetConsumedQuota() *int64
	SetCreateTime(v string) *PromotionActivity
	GetCreateTime() *string
	SetCreatedBy(v string) *PromotionActivity
	GetCreatedBy() *string
	SetEligibilityConfig(v string) *PromotionActivity
	GetEligibilityConfig() *string
	SetEndDate(v string) *PromotionActivity
	GetEndDate() *string
	SetOfferConfig(v string) *PromotionActivity
	GetOfferConfig() *string
	SetOfferConfigSummary(v string) *PromotionActivity
	GetOfferConfigSummary() *string
	SetRemainingQuota(v int64) *PromotionActivity
	GetRemainingQuota() *int64
	SetStartDate(v string) *PromotionActivity
	GetStartDate() *string
	SetStatus(v string) *PromotionActivity
	GetStatus() *string
	SetTotalQuota(v int64) *PromotionActivity
	GetTotalQuota() *int64
	SetTouchpointConfig(v string) *PromotionActivity
	GetTouchpointConfig() *string
	SetUpdateTime(v string) *PromotionActivity
	GetUpdateTime() *string
	SetUpdatedBy(v string) *PromotionActivity
	GetUpdatedBy() *string
	SetWarningThreshold(v int32) *PromotionActivity
	GetWarningThreshold() *int32
}

type PromotionActivity struct {
	ActivityCode       *string `json:"ActivityCode,omitempty" xml:"ActivityCode,omitempty"`
	ActivityName       *string `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	ActivityType       *string `json:"ActivityType,omitempty" xml:"ActivityType,omitempty"`
	ConsumedQuota      *int64  `json:"ConsumedQuota,omitempty" xml:"ConsumedQuota,omitempty"`
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreatedBy          *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	EligibilityConfig  *string `json:"EligibilityConfig,omitempty" xml:"EligibilityConfig,omitempty"`
	EndDate            *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	OfferConfig        *string `json:"OfferConfig,omitempty" xml:"OfferConfig,omitempty"`
	OfferConfigSummary *string `json:"OfferConfigSummary,omitempty" xml:"OfferConfigSummary,omitempty"`
	RemainingQuota     *int64  `json:"RemainingQuota,omitempty" xml:"RemainingQuota,omitempty"`
	StartDate          *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalQuota         *int64  `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	TouchpointConfig   *string `json:"TouchpointConfig,omitempty" xml:"TouchpointConfig,omitempty"`
	UpdateTime         *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdatedBy          *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	WarningThreshold   *int32  `json:"WarningThreshold,omitempty" xml:"WarningThreshold,omitempty"`
}

func (s PromotionActivity) String() string {
	return dara.Prettify(s)
}

func (s PromotionActivity) GoString() string {
	return s.String()
}

func (s *PromotionActivity) GetActivityCode() *string {
	return s.ActivityCode
}

func (s *PromotionActivity) GetActivityName() *string {
	return s.ActivityName
}

func (s *PromotionActivity) GetActivityType() *string {
	return s.ActivityType
}

func (s *PromotionActivity) GetConsumedQuota() *int64 {
	return s.ConsumedQuota
}

func (s *PromotionActivity) GetCreateTime() *string {
	return s.CreateTime
}

func (s *PromotionActivity) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *PromotionActivity) GetEligibilityConfig() *string {
	return s.EligibilityConfig
}

func (s *PromotionActivity) GetEndDate() *string {
	return s.EndDate
}

func (s *PromotionActivity) GetOfferConfig() *string {
	return s.OfferConfig
}

func (s *PromotionActivity) GetOfferConfigSummary() *string {
	return s.OfferConfigSummary
}

func (s *PromotionActivity) GetRemainingQuota() *int64 {
	return s.RemainingQuota
}

func (s *PromotionActivity) GetStartDate() *string {
	return s.StartDate
}

func (s *PromotionActivity) GetStatus() *string {
	return s.Status
}

func (s *PromotionActivity) GetTotalQuota() *int64 {
	return s.TotalQuota
}

func (s *PromotionActivity) GetTouchpointConfig() *string {
	return s.TouchpointConfig
}

func (s *PromotionActivity) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *PromotionActivity) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *PromotionActivity) GetWarningThreshold() *int32 {
	return s.WarningThreshold
}

func (s *PromotionActivity) SetActivityCode(v string) *PromotionActivity {
	s.ActivityCode = &v
	return s
}

func (s *PromotionActivity) SetActivityName(v string) *PromotionActivity {
	s.ActivityName = &v
	return s
}

func (s *PromotionActivity) SetActivityType(v string) *PromotionActivity {
	s.ActivityType = &v
	return s
}

func (s *PromotionActivity) SetConsumedQuota(v int64) *PromotionActivity {
	s.ConsumedQuota = &v
	return s
}

func (s *PromotionActivity) SetCreateTime(v string) *PromotionActivity {
	s.CreateTime = &v
	return s
}

func (s *PromotionActivity) SetCreatedBy(v string) *PromotionActivity {
	s.CreatedBy = &v
	return s
}

func (s *PromotionActivity) SetEligibilityConfig(v string) *PromotionActivity {
	s.EligibilityConfig = &v
	return s
}

func (s *PromotionActivity) SetEndDate(v string) *PromotionActivity {
	s.EndDate = &v
	return s
}

func (s *PromotionActivity) SetOfferConfig(v string) *PromotionActivity {
	s.OfferConfig = &v
	return s
}

func (s *PromotionActivity) SetOfferConfigSummary(v string) *PromotionActivity {
	s.OfferConfigSummary = &v
	return s
}

func (s *PromotionActivity) SetRemainingQuota(v int64) *PromotionActivity {
	s.RemainingQuota = &v
	return s
}

func (s *PromotionActivity) SetStartDate(v string) *PromotionActivity {
	s.StartDate = &v
	return s
}

func (s *PromotionActivity) SetStatus(v string) *PromotionActivity {
	s.Status = &v
	return s
}

func (s *PromotionActivity) SetTotalQuota(v int64) *PromotionActivity {
	s.TotalQuota = &v
	return s
}

func (s *PromotionActivity) SetTouchpointConfig(v string) *PromotionActivity {
	s.TouchpointConfig = &v
	return s
}

func (s *PromotionActivity) SetUpdateTime(v string) *PromotionActivity {
	s.UpdateTime = &v
	return s
}

func (s *PromotionActivity) SetUpdatedBy(v string) *PromotionActivity {
	s.UpdatedBy = &v
	return s
}

func (s *PromotionActivity) SetWarningThreshold(v int32) *PromotionActivity {
	s.WarningThreshold = &v
	return s
}

func (s *PromotionActivity) Validate() error {
	return dara.Validate(s)
}
