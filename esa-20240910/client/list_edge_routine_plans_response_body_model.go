// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeRoutinePlansResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListEdgeRoutinePlansResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEdgeRoutinePlansResponseBody
	GetPageSize() *int32
	SetPlanInfo(v []*ListEdgeRoutinePlansResponseBodyPlanInfo) *ListEdgeRoutinePlansResponseBody
	GetPlanInfo() []*ListEdgeRoutinePlansResponseBodyPlanInfo
	SetRequestId(v string) *ListEdgeRoutinePlansResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListEdgeRoutinePlansResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListEdgeRoutinePlansResponseBody
	GetTotalPage() *int32
}

type ListEdgeRoutinePlansResponseBody struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 500.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The plans.
	PlanInfo []*ListEdgeRoutinePlansResponseBodyPlanInfo `json:"PlanInfo,omitempty" xml:"PlanInfo,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListEdgeRoutinePlansResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeRoutinePlansResponseBody) GoString() string {
	return s.String()
}

func (s *ListEdgeRoutinePlansResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEdgeRoutinePlansResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEdgeRoutinePlansResponseBody) GetPlanInfo() []*ListEdgeRoutinePlansResponseBodyPlanInfo {
	return s.PlanInfo
}

func (s *ListEdgeRoutinePlansResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEdgeRoutinePlansResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEdgeRoutinePlansResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListEdgeRoutinePlansResponseBody) SetPageNumber(v int32) *ListEdgeRoutinePlansResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListEdgeRoutinePlansResponseBody) SetPageSize(v int32) *ListEdgeRoutinePlansResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEdgeRoutinePlansResponseBody) SetPlanInfo(v []*ListEdgeRoutinePlansResponseBodyPlanInfo) *ListEdgeRoutinePlansResponseBody {
	s.PlanInfo = v
	return s
}

func (s *ListEdgeRoutinePlansResponseBody) SetRequestId(v string) *ListEdgeRoutinePlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEdgeRoutinePlansResponseBody) SetTotalCount(v int32) *ListEdgeRoutinePlansResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEdgeRoutinePlansResponseBody) SetTotalPage(v int32) *ListEdgeRoutinePlansResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListEdgeRoutinePlansResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListEdgeRoutinePlansResponseBodyPlanInfo struct {
	// The billing method. Valid values:
	//
	// 	- PREPAY: subscription.
	//
	// 	- POSTPAY: pay-as-you-go.
	//
	// example:
	//
	// POSTPAY
	BillingMode *string `json:"BillingMode,omitempty" xml:"BillingMode,omitempty"`
	// The maximum number of versions that each routine supports.
	//
	// example:
	//
	// 20
	ErRoutineCodeVersionQuota *string `json:"ErRoutineCodeVersionQuota,omitempty" xml:"ErRoutineCodeVersionQuota,omitempty"`
	// The maximum of routines that can be created.
	//
	// example:
	//
	// 100
	ErRoutineQuota *string `json:"ErRoutineQuota,omitempty" xml:"ErRoutineQuota,omitempty"`
	// The maximum number of websites with which each routine can be associated.
	//
	// example:
	//
	// 100
	ErRoutineRouteSiteCountQuota *string `json:"ErRoutineRouteSiteCountQuota,omitempty" xml:"ErRoutineRouteSiteCountQuota,omitempty"`
	// The payment method. Valid values:
	//
	// 	- er_free
	//
	// 	- er_pay
	//
	// example:
	//
	// er_free
	PaymentMethod *string `json:"PaymentMethod,omitempty" xml:"PaymentMethod,omitempty"`
	// The plan name.
	//
	// example:
	//
	// test_plan
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
}

func (s ListEdgeRoutinePlansResponseBodyPlanInfo) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeRoutinePlansResponseBodyPlanInfo) GoString() string {
	return s.String()
}

func (s *ListEdgeRoutinePlansResponseBodyPlanInfo) GetBillingMode() *string {
	return s.BillingMode
}

func (s *ListEdgeRoutinePlansResponseBodyPlanInfo) GetErRoutineCodeVersionQuota() *string {
	return s.ErRoutineCodeVersionQuota
}

func (s *ListEdgeRoutinePlansResponseBodyPlanInfo) GetErRoutineQuota() *string {
	return s.ErRoutineQuota
}

func (s *ListEdgeRoutinePlansResponseBodyPlanInfo) GetErRoutineRouteSiteCountQuota() *string {
	return s.ErRoutineRouteSiteCountQuota
}

func (s *ListEdgeRoutinePlansResponseBodyPlanInfo) GetPaymentMethod() *string {
	return s.PaymentMethod
}

func (s *ListEdgeRoutinePlansResponseBodyPlanInfo) GetPlanName() *string {
	return s.PlanName
}

func (s *ListEdgeRoutinePlansResponseBodyPlanInfo) SetBillingMode(v string) *ListEdgeRoutinePlansResponseBodyPlanInfo {
	s.BillingMode = &v
	return s
}

func (s *ListEdgeRoutinePlansResponseBodyPlanInfo) SetErRoutineCodeVersionQuota(v string) *ListEdgeRoutinePlansResponseBodyPlanInfo {
	s.ErRoutineCodeVersionQuota = &v
	return s
}

func (s *ListEdgeRoutinePlansResponseBodyPlanInfo) SetErRoutineQuota(v string) *ListEdgeRoutinePlansResponseBodyPlanInfo {
	s.ErRoutineQuota = &v
	return s
}

func (s *ListEdgeRoutinePlansResponseBodyPlanInfo) SetErRoutineRouteSiteCountQuota(v string) *ListEdgeRoutinePlansResponseBodyPlanInfo {
	s.ErRoutineRouteSiteCountQuota = &v
	return s
}

func (s *ListEdgeRoutinePlansResponseBodyPlanInfo) SetPaymentMethod(v string) *ListEdgeRoutinePlansResponseBodyPlanInfo {
	s.PaymentMethod = &v
	return s
}

func (s *ListEdgeRoutinePlansResponseBodyPlanInfo) SetPlanName(v string) *ListEdgeRoutinePlansResponseBodyPlanInfo {
	s.PlanName = &v
	return s
}

func (s *ListEdgeRoutinePlansResponseBodyPlanInfo) Validate() error {
	return dara.Validate(s)
}
