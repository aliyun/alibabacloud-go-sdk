// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSiteFeaturesMatchPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsPassed(v bool) *CheckSiteFeaturesMatchPlanResponseBody
	GetIsPassed() *bool
	SetRequestId(v string) *CheckSiteFeaturesMatchPlanResponseBody
	GetRequestId() *string
	SetUnPassedSiteQuotas(v []*CheckSiteFeaturesMatchPlanResponseBodyUnPassedSiteQuotas) *CheckSiteFeaturesMatchPlanResponseBody
	GetUnPassedSiteQuotas() []*CheckSiteFeaturesMatchPlanResponseBodyUnPassedSiteQuotas
}

type CheckSiteFeaturesMatchPlanResponseBody struct {
	// Indicates whether the features of the current site are compatible with the target instance.
	//
	// example:
	//
	// true
	IsPassed *bool `json:"IsPassed,omitempty" xml:"IsPassed,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 65705631-908C-5D24-997C-17E0397721C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about site quotas that do not meet the requirements.
	UnPassedSiteQuotas []*CheckSiteFeaturesMatchPlanResponseBodyUnPassedSiteQuotas `json:"UnPassedSiteQuotas,omitempty" xml:"UnPassedSiteQuotas,omitempty" type:"Repeated"`
}

func (s CheckSiteFeaturesMatchPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckSiteFeaturesMatchPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CheckSiteFeaturesMatchPlanResponseBody) GetIsPassed() *bool {
	return s.IsPassed
}

func (s *CheckSiteFeaturesMatchPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckSiteFeaturesMatchPlanResponseBody) GetUnPassedSiteQuotas() []*CheckSiteFeaturesMatchPlanResponseBodyUnPassedSiteQuotas {
	return s.UnPassedSiteQuotas
}

func (s *CheckSiteFeaturesMatchPlanResponseBody) SetIsPassed(v bool) *CheckSiteFeaturesMatchPlanResponseBody {
	s.IsPassed = &v
	return s
}

func (s *CheckSiteFeaturesMatchPlanResponseBody) SetRequestId(v string) *CheckSiteFeaturesMatchPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckSiteFeaturesMatchPlanResponseBody) SetUnPassedSiteQuotas(v []*CheckSiteFeaturesMatchPlanResponseBodyUnPassedSiteQuotas) *CheckSiteFeaturesMatchPlanResponseBody {
	s.UnPassedSiteQuotas = v
	return s
}

func (s *CheckSiteFeaturesMatchPlanResponseBody) Validate() error {
	if s.UnPassedSiteQuotas != nil {
		for _, item := range s.UnPassedSiteQuotas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CheckSiteFeaturesMatchPlanResponseBodyUnPassedSiteQuotas struct {
	// The quota value of the current site.
	//
	// example:
	//
	// true
	CurrentQuotaValue *string `json:"CurrentQuotaValue,omitempty" xml:"CurrentQuotaValue,omitempty"`
	// The quota value of the target instance.
	//
	// example:
	//
	// false
	DestQuotaValue *string `json:"DestQuotaValue,omitempty" xml:"DestQuotaValue,omitempty"`
	// The quota name.
	//
	// example:
	//
	// cache_reserve
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
}

func (s CheckSiteFeaturesMatchPlanResponseBodyUnPassedSiteQuotas) String() string {
	return dara.Prettify(s)
}

func (s CheckSiteFeaturesMatchPlanResponseBodyUnPassedSiteQuotas) GoString() string {
	return s.String()
}

func (s *CheckSiteFeaturesMatchPlanResponseBodyUnPassedSiteQuotas) GetCurrentQuotaValue() *string {
	return s.CurrentQuotaValue
}

func (s *CheckSiteFeaturesMatchPlanResponseBodyUnPassedSiteQuotas) GetDestQuotaValue() *string {
	return s.DestQuotaValue
}

func (s *CheckSiteFeaturesMatchPlanResponseBodyUnPassedSiteQuotas) GetQuotaName() *string {
	return s.QuotaName
}

func (s *CheckSiteFeaturesMatchPlanResponseBodyUnPassedSiteQuotas) SetCurrentQuotaValue(v string) *CheckSiteFeaturesMatchPlanResponseBodyUnPassedSiteQuotas {
	s.CurrentQuotaValue = &v
	return s
}

func (s *CheckSiteFeaturesMatchPlanResponseBodyUnPassedSiteQuotas) SetDestQuotaValue(v string) *CheckSiteFeaturesMatchPlanResponseBodyUnPassedSiteQuotas {
	s.DestQuotaValue = &v
	return s
}

func (s *CheckSiteFeaturesMatchPlanResponseBodyUnPassedSiteQuotas) SetQuotaName(v string) *CheckSiteFeaturesMatchPlanResponseBodyUnPassedSiteQuotas {
	s.QuotaName = &v
	return s
}

func (s *CheckSiteFeaturesMatchPlanResponseBodyUnPassedSiteQuotas) Validate() error {
	return dara.Validate(s)
}
