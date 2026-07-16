// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPostpaidSitePlansResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListPostpaidSitePlansResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPostpaidSitePlansResponseBody
	GetPageSize() *int32
	SetPlanInfo(v []*ListPostpaidSitePlansResponseBodyPlanInfo) *ListPostpaidSitePlansResponseBody
	GetPlanInfo() []*ListPostpaidSitePlansResponseBodyPlanInfo
	SetRequestId(v string) *ListPostpaidSitePlansResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListPostpaidSitePlansResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListPostpaidSitePlansResponseBody
	GetTotalPage() *int32
}

type ListPostpaidSitePlansResponseBody struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size. Default value: 20. Maximum value: 500. Valid values: any integer from 1 to 500.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The pay-as-you-go plan details.
	PlanInfo []*ListPostpaidSitePlansResponseBodyPlanInfo `json:"PlanInfo,omitempty" xml:"PlanInfo,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 8CD541DB-CD83-5D0C-BE94-21B794074249
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 16
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListPostpaidSitePlansResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPostpaidSitePlansResponseBody) GoString() string {
	return s.String()
}

func (s *ListPostpaidSitePlansResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPostpaidSitePlansResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPostpaidSitePlansResponseBody) GetPlanInfo() []*ListPostpaidSitePlansResponseBodyPlanInfo {
	return s.PlanInfo
}

func (s *ListPostpaidSitePlansResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPostpaidSitePlansResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPostpaidSitePlansResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListPostpaidSitePlansResponseBody) SetPageNumber(v int32) *ListPostpaidSitePlansResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPostpaidSitePlansResponseBody) SetPageSize(v int32) *ListPostpaidSitePlansResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPostpaidSitePlansResponseBody) SetPlanInfo(v []*ListPostpaidSitePlansResponseBodyPlanInfo) *ListPostpaidSitePlansResponseBody {
	s.PlanInfo = v
	return s
}

func (s *ListPostpaidSitePlansResponseBody) SetRequestId(v string) *ListPostpaidSitePlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPostpaidSitePlansResponseBody) SetTotalCount(v int32) *ListPostpaidSitePlansResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPostpaidSitePlansResponseBody) SetTotalPage(v int32) *ListPostpaidSitePlansResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListPostpaidSitePlansResponseBody) Validate() error {
	if s.PlanInfo != nil {
		for _, item := range s.PlanInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPostpaidSitePlansResponseBodyPlanInfo struct {
	// The billable methods of the plan. Valid values:
	//
	// - dps_month95: monthly 95th percentile billing method.
	//
	// example:
	//
	// dps_month95
	BillingMethod *string `json:"BillingMethod,omitempty" xml:"BillingMethod,omitempty"`
	// The payment type. Valid values:
	//
	// - PREPAY: subscription.
	//
	// - POSTPAY: pay-as-you-go.
	//
	// example:
	//
	// PREPAY
	BillingMode *string `json:"BillingMode,omitempty" xml:"BillingMode,omitempty"`
	// The acceleration regions to which sites can be added under the plan. Multiple values are separated by commas (,). Valid values:
	//
	// - domestic: the Chinese mainland.
	//
	// - overseas: global (excluding the Chinese mainland).
	//
	// - global: global (including the Chinese mainland).
	//
	// example:
	//
	// domestic
	Coverages *string `json:"Coverages,omitempty" xml:"Coverages,omitempty"`
	// The name of the plan, which serves as a unique identifier in English.
	//
	// example:
	//
	// basic
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// The plan description.
	//
	// example:
	//
	// 测试版套餐
	PlanNameCn *string `json:"PlanNameCn,omitempty" xml:"PlanNameCn,omitempty"`
	// The plan type of the plan instance. Valid values:
	//
	// - normal: fixed edition plan.
	//
	// - enterprise: enterprise edition plan.
	//
	// example:
	//
	// normal
	PlanType *string `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
	// The sale status of the plan. Valid values for enterprise edition plans:
	//
	// - saled: sold.
	//
	// - upgrading: specification change in progress.
	//
	// example:
	//
	// saled
	SaleStatus *string `json:"SaleStatus,omitempty" xml:"SaleStatus,omitempty"`
	// The site quantity quota.
	//
	// example:
	//
	// 1
	SiteQuota *string `json:"SiteQuota,omitempty" xml:"SiteQuota,omitempty"`
}

func (s ListPostpaidSitePlansResponseBodyPlanInfo) String() string {
	return dara.Prettify(s)
}

func (s ListPostpaidSitePlansResponseBodyPlanInfo) GoString() string {
	return s.String()
}

func (s *ListPostpaidSitePlansResponseBodyPlanInfo) GetBillingMethod() *string {
	return s.BillingMethod
}

func (s *ListPostpaidSitePlansResponseBodyPlanInfo) GetBillingMode() *string {
	return s.BillingMode
}

func (s *ListPostpaidSitePlansResponseBodyPlanInfo) GetCoverages() *string {
	return s.Coverages
}

func (s *ListPostpaidSitePlansResponseBodyPlanInfo) GetPlanName() *string {
	return s.PlanName
}

func (s *ListPostpaidSitePlansResponseBodyPlanInfo) GetPlanNameCn() *string {
	return s.PlanNameCn
}

func (s *ListPostpaidSitePlansResponseBodyPlanInfo) GetPlanType() *string {
	return s.PlanType
}

func (s *ListPostpaidSitePlansResponseBodyPlanInfo) GetSaleStatus() *string {
	return s.SaleStatus
}

func (s *ListPostpaidSitePlansResponseBodyPlanInfo) GetSiteQuota() *string {
	return s.SiteQuota
}

func (s *ListPostpaidSitePlansResponseBodyPlanInfo) SetBillingMethod(v string) *ListPostpaidSitePlansResponseBodyPlanInfo {
	s.BillingMethod = &v
	return s
}

func (s *ListPostpaidSitePlansResponseBodyPlanInfo) SetBillingMode(v string) *ListPostpaidSitePlansResponseBodyPlanInfo {
	s.BillingMode = &v
	return s
}

func (s *ListPostpaidSitePlansResponseBodyPlanInfo) SetCoverages(v string) *ListPostpaidSitePlansResponseBodyPlanInfo {
	s.Coverages = &v
	return s
}

func (s *ListPostpaidSitePlansResponseBodyPlanInfo) SetPlanName(v string) *ListPostpaidSitePlansResponseBodyPlanInfo {
	s.PlanName = &v
	return s
}

func (s *ListPostpaidSitePlansResponseBodyPlanInfo) SetPlanNameCn(v string) *ListPostpaidSitePlansResponseBodyPlanInfo {
	s.PlanNameCn = &v
	return s
}

func (s *ListPostpaidSitePlansResponseBodyPlanInfo) SetPlanType(v string) *ListPostpaidSitePlansResponseBodyPlanInfo {
	s.PlanType = &v
	return s
}

func (s *ListPostpaidSitePlansResponseBodyPlanInfo) SetSaleStatus(v string) *ListPostpaidSitePlansResponseBodyPlanInfo {
	s.SaleStatus = &v
	return s
}

func (s *ListPostpaidSitePlansResponseBodyPlanInfo) SetSiteQuota(v string) *ListPostpaidSitePlansResponseBodyPlanInfo {
	s.SiteQuota = &v
	return s
}

func (s *ListPostpaidSitePlansResponseBodyPlanInfo) Validate() error {
	return dara.Validate(s)
}
