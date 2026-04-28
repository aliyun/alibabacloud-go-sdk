// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBudgetPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeBudgetPoliciesResponseBodyItems) *DescribeBudgetPoliciesResponseBody
	GetItems() []*DescribeBudgetPoliciesResponseBodyItems
	SetPageNumber(v int32) *DescribeBudgetPoliciesResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeBudgetPoliciesResponseBody
	GetPageRecordCount() *int32
	SetPageSize(v int32) *DescribeBudgetPoliciesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeBudgetPoliciesResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeBudgetPoliciesResponseBody
	GetTotalRecordCount() *int32
}

type DescribeBudgetPoliciesResponseBody struct {
	Items []*DescribeBudgetPoliciesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 24A1990B-4F6E-482B-B8CB-75C612******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeBudgetPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBudgetPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBudgetPoliciesResponseBody) GetItems() []*DescribeBudgetPoliciesResponseBodyItems {
	return s.Items
}

func (s *DescribeBudgetPoliciesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBudgetPoliciesResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeBudgetPoliciesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBudgetPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBudgetPoliciesResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeBudgetPoliciesResponseBody) SetItems(v []*DescribeBudgetPoliciesResponseBodyItems) *DescribeBudgetPoliciesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeBudgetPoliciesResponseBody) SetPageNumber(v int32) *DescribeBudgetPoliciesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBudgetPoliciesResponseBody) SetPageRecordCount(v int32) *DescribeBudgetPoliciesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeBudgetPoliciesResponseBody) SetPageSize(v int32) *DescribeBudgetPoliciesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBudgetPoliciesResponseBody) SetRequestId(v string) *DescribeBudgetPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBudgetPoliciesResponseBody) SetTotalRecordCount(v int32) *DescribeBudgetPoliciesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeBudgetPoliciesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBudgetPoliciesResponseBodyItems struct {
	// example:
	//
	// 80
	AlertThresholdPct *string `json:"AlertThresholdPct,omitempty" xml:"AlertThresholdPct,omitempty"`
	AlertTriggered    *bool   `json:"AlertTriggered,omitempty" xml:"AlertTriggered,omitempty"`
	// example:
	//
	// cg-p3gk2oh55c**
	BudgetDimensionRefId *string `json:"BudgetDimensionRefId,omitempty" xml:"BudgetDimensionRefId,omitempty"`
	// example:
	//
	// ConsumerGroup
	BudgetDimensionType *string `json:"BudgetDimensionType,omitempty" xml:"BudgetDimensionType,omitempty"`
	// example:
	//
	// 10000
	BudgetPoints *string `json:"BudgetPoints,omitempty" xml:"BudgetPoints,omitempty"`
	// example:
	//
	// 023aacc1effc4b56bb154bfbec6ba9**
	BudgetPolicyId *string `json:"BudgetPolicyId,omitempty" xml:"BudgetPolicyId,omitempty"`
	// example:
	//
	// GlobalTotal
	BudgetType *string `json:"BudgetType,omitempty" xml:"BudgetType,omitempty"`
	Exceeded   *string `json:"Exceeded,omitempty" xml:"Exceeded,omitempty"`
	// example:
	//
	// 2025-03-19T14:13:53+08:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// 2025-04-08T10:43:28+08:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// pg-xxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// example:
	//
	// 2
	ResetDayOfMonth *string `json:"ResetDayOfMonth,omitempty" xml:"ResetDayOfMonth,omitempty"`
	// example:
	//
	// Enabled
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UsedPoints *int32  `json:"UsedPoints,omitempty" xml:"UsedPoints,omitempty"`
}

func (s DescribeBudgetPoliciesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeBudgetPoliciesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeBudgetPoliciesResponseBodyItems) GetAlertThresholdPct() *string {
	return s.AlertThresholdPct
}

func (s *DescribeBudgetPoliciesResponseBodyItems) GetAlertTriggered() *bool {
	return s.AlertTriggered
}

func (s *DescribeBudgetPoliciesResponseBodyItems) GetBudgetDimensionRefId() *string {
	return s.BudgetDimensionRefId
}

func (s *DescribeBudgetPoliciesResponseBodyItems) GetBudgetDimensionType() *string {
	return s.BudgetDimensionType
}

func (s *DescribeBudgetPoliciesResponseBodyItems) GetBudgetPoints() *string {
	return s.BudgetPoints
}

func (s *DescribeBudgetPoliciesResponseBodyItems) GetBudgetPolicyId() *string {
	return s.BudgetPolicyId
}

func (s *DescribeBudgetPoliciesResponseBodyItems) GetBudgetType() *string {
	return s.BudgetType
}

func (s *DescribeBudgetPoliciesResponseBodyItems) GetExceeded() *string {
	return s.Exceeded
}

func (s *DescribeBudgetPoliciesResponseBodyItems) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeBudgetPoliciesResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeBudgetPoliciesResponseBodyItems) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DescribeBudgetPoliciesResponseBodyItems) GetResetDayOfMonth() *string {
	return s.ResetDayOfMonth
}

func (s *DescribeBudgetPoliciesResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeBudgetPoliciesResponseBodyItems) GetUsedPoints() *int32 {
	return s.UsedPoints
}

func (s *DescribeBudgetPoliciesResponseBodyItems) SetAlertThresholdPct(v string) *DescribeBudgetPoliciesResponseBodyItems {
	s.AlertThresholdPct = &v
	return s
}

func (s *DescribeBudgetPoliciesResponseBodyItems) SetAlertTriggered(v bool) *DescribeBudgetPoliciesResponseBodyItems {
	s.AlertTriggered = &v
	return s
}

func (s *DescribeBudgetPoliciesResponseBodyItems) SetBudgetDimensionRefId(v string) *DescribeBudgetPoliciesResponseBodyItems {
	s.BudgetDimensionRefId = &v
	return s
}

func (s *DescribeBudgetPoliciesResponseBodyItems) SetBudgetDimensionType(v string) *DescribeBudgetPoliciesResponseBodyItems {
	s.BudgetDimensionType = &v
	return s
}

func (s *DescribeBudgetPoliciesResponseBodyItems) SetBudgetPoints(v string) *DescribeBudgetPoliciesResponseBodyItems {
	s.BudgetPoints = &v
	return s
}

func (s *DescribeBudgetPoliciesResponseBodyItems) SetBudgetPolicyId(v string) *DescribeBudgetPoliciesResponseBodyItems {
	s.BudgetPolicyId = &v
	return s
}

func (s *DescribeBudgetPoliciesResponseBodyItems) SetBudgetType(v string) *DescribeBudgetPoliciesResponseBodyItems {
	s.BudgetType = &v
	return s
}

func (s *DescribeBudgetPoliciesResponseBodyItems) SetExceeded(v string) *DescribeBudgetPoliciesResponseBodyItems {
	s.Exceeded = &v
	return s
}

func (s *DescribeBudgetPoliciesResponseBodyItems) SetGmtCreated(v string) *DescribeBudgetPoliciesResponseBodyItems {
	s.GmtCreated = &v
	return s
}

func (s *DescribeBudgetPoliciesResponseBodyItems) SetGmtModified(v string) *DescribeBudgetPoliciesResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *DescribeBudgetPoliciesResponseBodyItems) SetGwClusterId(v string) *DescribeBudgetPoliciesResponseBodyItems {
	s.GwClusterId = &v
	return s
}

func (s *DescribeBudgetPoliciesResponseBodyItems) SetResetDayOfMonth(v string) *DescribeBudgetPoliciesResponseBodyItems {
	s.ResetDayOfMonth = &v
	return s
}

func (s *DescribeBudgetPoliciesResponseBodyItems) SetStatus(v string) *DescribeBudgetPoliciesResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeBudgetPoliciesResponseBodyItems) SetUsedPoints(v int32) *DescribeBudgetPoliciesResponseBodyItems {
	s.UsedPoints = &v
	return s
}

func (s *DescribeBudgetPoliciesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
