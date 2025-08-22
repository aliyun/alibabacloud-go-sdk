// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEdition(v string) *DescribeDcdnWafServiceResponseBody
	GetEdition() *string
	SetEnabled(v string) *DescribeDcdnWafServiceResponseBody
	GetEnabled() *string
	SetOpeningTime(v string) *DescribeDcdnWafServiceResponseBody
	GetOpeningTime() *string
	SetRequestBillingType(v string) *DescribeDcdnWafServiceResponseBody
	GetRequestBillingType() *string
	SetRequestId(v string) *DescribeDcdnWafServiceResponseBody
	GetRequestId() *string
	SetRuleBillingType(v string) *DescribeDcdnWafServiceResponseBody
	GetRuleBillingType() *string
	SetStatus(v string) *DescribeDcdnWafServiceResponseBody
	GetStatus() *string
}

type DescribeDcdnWafServiceResponseBody struct {
	// The edition of WAF.
	//
	// example:
	//
	// dcdnwaf_afterpay
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The status of WAF. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	Enabled *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The time when WAF was enabled.
	//
	// example:
	//
	// 2021-09-26T16:00:00Z
	OpeningTime *string `json:"OpeningTime,omitempty" xml:"OpeningTime,omitempty"`
	// The metering method for requests.
	//
	// example:
	//
	// dcdn_waf_req
	RequestBillingType *string `json:"RequestBillingType,omitempty" xml:"RequestBillingType,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4A95CA90-E0F2-1BF6-99E0-8C1510CAF649
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The metering method for rules. You are charged for the number of SeCUs.
	//
	// example:
	//
	// dcdn_waf_rule
	RuleBillingType *string `json:"RuleBillingType,omitempty" xml:"RuleBillingType,omitempty"`
	// The status of WAF. Valid values:
	//
	// 	- Normal
	//
	// 	- WaitForExpire
	//
	// 	- Expired
	//
	// 	- Released
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDcdnWafServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafServiceResponseBody) GetEdition() *string {
	return s.Edition
}

func (s *DescribeDcdnWafServiceResponseBody) GetEnabled() *string {
	return s.Enabled
}

func (s *DescribeDcdnWafServiceResponseBody) GetOpeningTime() *string {
	return s.OpeningTime
}

func (s *DescribeDcdnWafServiceResponseBody) GetRequestBillingType() *string {
	return s.RequestBillingType
}

func (s *DescribeDcdnWafServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnWafServiceResponseBody) GetRuleBillingType() *string {
	return s.RuleBillingType
}

func (s *DescribeDcdnWafServiceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeDcdnWafServiceResponseBody) SetEdition(v string) *DescribeDcdnWafServiceResponseBody {
	s.Edition = &v
	return s
}

func (s *DescribeDcdnWafServiceResponseBody) SetEnabled(v string) *DescribeDcdnWafServiceResponseBody {
	s.Enabled = &v
	return s
}

func (s *DescribeDcdnWafServiceResponseBody) SetOpeningTime(v string) *DescribeDcdnWafServiceResponseBody {
	s.OpeningTime = &v
	return s
}

func (s *DescribeDcdnWafServiceResponseBody) SetRequestBillingType(v string) *DescribeDcdnWafServiceResponseBody {
	s.RequestBillingType = &v
	return s
}

func (s *DescribeDcdnWafServiceResponseBody) SetRequestId(v string) *DescribeDcdnWafServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnWafServiceResponseBody) SetRuleBillingType(v string) *DescribeDcdnWafServiceResponseBody {
	s.RuleBillingType = &v
	return s
}

func (s *DescribeDcdnWafServiceResponseBody) SetStatus(v string) *DescribeDcdnWafServiceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDcdnWafServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
