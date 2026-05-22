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
	PageNumber *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PlanInfo   []*ListEdgeRoutinePlansResponseBodyPlanInfo `json:"PlanInfo,omitempty" xml:"PlanInfo,omitempty" type:"Repeated"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                                      `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
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

type ListEdgeRoutinePlansResponseBodyPlanInfo struct {
	BillingMode                  *string `json:"BillingMode,omitempty" xml:"BillingMode,omitempty"`
	ErRoutineCodeVersionQuota    *string `json:"ErRoutineCodeVersionQuota,omitempty" xml:"ErRoutineCodeVersionQuota,omitempty"`
	ErRoutineQuota               *string `json:"ErRoutineQuota,omitempty" xml:"ErRoutineQuota,omitempty"`
	ErRoutineRouteSiteCountQuota *string `json:"ErRoutineRouteSiteCountQuota,omitempty" xml:"ErRoutineRouteSiteCountQuota,omitempty"`
	PaymentMethod                *string `json:"PaymentMethod,omitempty" xml:"PaymentMethod,omitempty"`
	PlanName                     *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
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
