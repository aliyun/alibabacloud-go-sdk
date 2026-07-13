// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCreditPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreditsPackageInfos(v []*DescribeCreditPackageResponseBodyCreditsPackageInfos) *DescribeCreditPackageResponseBody
	GetCreditsPackageInfos() []*DescribeCreditPackageResponseBodyCreditsPackageInfos
	SetIsFirstPurchase(v bool) *DescribeCreditPackageResponseBody
	GetIsFirstPurchase() *bool
	SetRequestId(v string) *DescribeCreditPackageResponseBody
	GetRequestId() *string
	SetTotalAvailableCredits(v string) *DescribeCreditPackageResponseBody
	GetTotalAvailableCredits() *string
	SetTotalCount(v int32) *DescribeCreditPackageResponseBody
	GetTotalCount() *int32
	SetTotalExhaustedCredit(v string) *DescribeCreditPackageResponseBody
	GetTotalExhaustedCredit() *string
}

type DescribeCreditPackageResponseBody struct {
	// The credit booster package information.
	CreditsPackageInfos []*DescribeCreditPackageResponseBodyCreditsPackageInfos `json:"CreditsPackageInfos,omitempty" xml:"CreditsPackageInfos,omitempty" type:"Repeated"`
	// Indicates whether this is the first purchase.
	//
	// example:
	//
	// true
	IsFirstPurchase *bool `json:"IsFirstPurchase,omitempty" xml:"IsFirstPurchase,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of available credits.
	//
	// example:
	//
	// 1000
	TotalAvailableCredits *string `json:"TotalAvailableCredits,omitempty" xml:"TotalAvailableCredits,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of consumed credits.
	//
	// example:
	//
	// 30
	TotalExhaustedCredit *string `json:"TotalExhaustedCredit,omitempty" xml:"TotalExhaustedCredit,omitempty"`
}

func (s DescribeCreditPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditPackageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCreditPackageResponseBody) GetCreditsPackageInfos() []*DescribeCreditPackageResponseBodyCreditsPackageInfos {
	return s.CreditsPackageInfos
}

func (s *DescribeCreditPackageResponseBody) GetIsFirstPurchase() *bool {
	return s.IsFirstPurchase
}

func (s *DescribeCreditPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCreditPackageResponseBody) GetTotalAvailableCredits() *string {
	return s.TotalAvailableCredits
}

func (s *DescribeCreditPackageResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCreditPackageResponseBody) GetTotalExhaustedCredit() *string {
	return s.TotalExhaustedCredit
}

func (s *DescribeCreditPackageResponseBody) SetCreditsPackageInfos(v []*DescribeCreditPackageResponseBodyCreditsPackageInfos) *DescribeCreditPackageResponseBody {
	s.CreditsPackageInfos = v
	return s
}

func (s *DescribeCreditPackageResponseBody) SetIsFirstPurchase(v bool) *DescribeCreditPackageResponseBody {
	s.IsFirstPurchase = &v
	return s
}

func (s *DescribeCreditPackageResponseBody) SetRequestId(v string) *DescribeCreditPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCreditPackageResponseBody) SetTotalAvailableCredits(v string) *DescribeCreditPackageResponseBody {
	s.TotalAvailableCredits = &v
	return s
}

func (s *DescribeCreditPackageResponseBody) SetTotalCount(v int32) *DescribeCreditPackageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCreditPackageResponseBody) SetTotalExhaustedCredit(v string) *DescribeCreditPackageResponseBody {
	s.TotalExhaustedCredit = &v
	return s
}

func (s *DescribeCreditPackageResponseBody) Validate() error {
	if s.CreditsPackageInfos != nil {
		for _, item := range s.CreditsPackageInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCreditPackageResponseBodyCreditsPackageInfos struct {
	// The number of available credits in the current credit booster package.
	//
	// example:
	//
	// 70
	AvailableCredits *string `json:"AvailableCredits,omitempty" xml:"AvailableCredits,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the credit booster package.
	//
	// example:
	//
	// crp-xagydbhfkah****
	CreditPackageId *string `json:"CreditPackageId,omitempty" xml:"CreditPackageId,omitempty"`
	// The status of the credit booster package.
	//
	// example:
	//
	// ACTIVE
	CreditPackageStatus *string `json:"CreditPackageStatus,omitempty" xml:"CreditPackageStatus,omitempty"`
	// The effective period of the credit booster package.
	//
	// example:
	//
	// 2026-04-30 00:00:00
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The number of consumed credits in the current credit booster package.
	//
	// example:
	//
	// 30
	ExhaustedCredits *string `json:"ExhaustedCredits,omitempty" xml:"ExhaustedCredits,omitempty"`
	// The expiration time of the credit booster package.
	//
	// example:
	//
	// 2026-10-30 00:00:00
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The total number of credits in the current credit booster package.
	//
	// example:
	//
	// 100
	TotalCredits *string `json:"TotalCredits,omitempty" xml:"TotalCredits,omitempty"`
}

func (s DescribeCreditPackageResponseBodyCreditsPackageInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditPackageResponseBodyCreditsPackageInfos) GoString() string {
	return s.String()
}

func (s *DescribeCreditPackageResponseBodyCreditsPackageInfos) GetAvailableCredits() *string {
	return s.AvailableCredits
}

func (s *DescribeCreditPackageResponseBodyCreditsPackageInfos) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeCreditPackageResponseBodyCreditsPackageInfos) GetCreditPackageId() *string {
	return s.CreditPackageId
}

func (s *DescribeCreditPackageResponseBodyCreditsPackageInfos) GetCreditPackageStatus() *string {
	return s.CreditPackageStatus
}

func (s *DescribeCreditPackageResponseBodyCreditsPackageInfos) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *DescribeCreditPackageResponseBodyCreditsPackageInfos) GetExhaustedCredits() *string {
	return s.ExhaustedCredits
}

func (s *DescribeCreditPackageResponseBodyCreditsPackageInfos) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeCreditPackageResponseBodyCreditsPackageInfos) GetTotalCredits() *string {
	return s.TotalCredits
}

func (s *DescribeCreditPackageResponseBodyCreditsPackageInfos) SetAvailableCredits(v string) *DescribeCreditPackageResponseBodyCreditsPackageInfos {
	s.AvailableCredits = &v
	return s
}

func (s *DescribeCreditPackageResponseBodyCreditsPackageInfos) SetCreateTime(v string) *DescribeCreditPackageResponseBodyCreditsPackageInfos {
	s.CreateTime = &v
	return s
}

func (s *DescribeCreditPackageResponseBodyCreditsPackageInfos) SetCreditPackageId(v string) *DescribeCreditPackageResponseBodyCreditsPackageInfos {
	s.CreditPackageId = &v
	return s
}

func (s *DescribeCreditPackageResponseBodyCreditsPackageInfos) SetCreditPackageStatus(v string) *DescribeCreditPackageResponseBodyCreditsPackageInfos {
	s.CreditPackageStatus = &v
	return s
}

func (s *DescribeCreditPackageResponseBodyCreditsPackageInfos) SetEffectiveTime(v string) *DescribeCreditPackageResponseBodyCreditsPackageInfos {
	s.EffectiveTime = &v
	return s
}

func (s *DescribeCreditPackageResponseBodyCreditsPackageInfos) SetExhaustedCredits(v string) *DescribeCreditPackageResponseBodyCreditsPackageInfos {
	s.ExhaustedCredits = &v
	return s
}

func (s *DescribeCreditPackageResponseBodyCreditsPackageInfos) SetExpiredTime(v string) *DescribeCreditPackageResponseBodyCreditsPackageInfos {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeCreditPackageResponseBodyCreditsPackageInfos) SetTotalCredits(v string) *DescribeCreditPackageResponseBodyCreditsPackageInfos {
	s.TotalCredits = &v
	return s
}

func (s *DescribeCreditPackageResponseBodyCreditsPackageInfos) Validate() error {
	return dara.Validate(s)
}
