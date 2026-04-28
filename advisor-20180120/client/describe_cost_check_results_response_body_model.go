// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCostCheckResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeCostCheckResultsResponseBody
	GetCode() *string
	SetData(v *DescribeCostCheckResultsResponseBodyData) *DescribeCostCheckResultsResponseBody
	GetData() *DescribeCostCheckResultsResponseBodyData
	SetMessage(v string) *DescribeCostCheckResultsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCostCheckResultsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeCostCheckResultsResponseBody
	GetSuccess() *string
}

type DescribeCostCheckResultsResponseBody struct {
	// example:
	//
	// 200
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeCostCheckResultsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 566331F9-****-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCostCheckResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostCheckResultsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckResultsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeCostCheckResultsResponseBody) GetData() *DescribeCostCheckResultsResponseBodyData {
	return s.Data
}

func (s *DescribeCostCheckResultsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCostCheckResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCostCheckResultsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeCostCheckResultsResponseBody) SetCode(v string) *DescribeCostCheckResultsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBody) SetData(v *DescribeCostCheckResultsResponseBodyData) *DescribeCostCheckResultsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCostCheckResultsResponseBody) SetMessage(v string) *DescribeCostCheckResultsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBody) SetRequestId(v string) *DescribeCostCheckResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBody) SetSuccess(v string) *DescribeCostCheckResultsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCostCheckResultsResponseBodyData struct {
	AdviceResourceCount *int32 `json:"AdviceResourceCount,omitempty" xml:"AdviceResourceCount,omitempty"`
	// example:
	//
	// Category
	GroupBy *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	// example:
	//
	// 1
	NormalCount *int32 `json:"NormalCount,omitempty" xml:"NormalCount,omitempty"`
	// example:
	//
	// 76
	ResourceCount *int32 `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
	// example:
	//
	// 4
	TotalCount *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	ViewGroup  []*DescribeCostCheckResultsResponseBodyDataViewGroup `json:"ViewGroup,omitempty" xml:"ViewGroup,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	WarningCount *int32 `json:"WarningCount,omitempty" xml:"WarningCount,omitempty"`
}

func (s DescribeCostCheckResultsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostCheckResultsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckResultsResponseBodyData) GetAdviceResourceCount() *int32 {
	return s.AdviceResourceCount
}

func (s *DescribeCostCheckResultsResponseBodyData) GetGroupBy() *string {
	return s.GroupBy
}

func (s *DescribeCostCheckResultsResponseBodyData) GetNormalCount() *int32 {
	return s.NormalCount
}

func (s *DescribeCostCheckResultsResponseBodyData) GetResourceCount() *int32 {
	return s.ResourceCount
}

func (s *DescribeCostCheckResultsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCostCheckResultsResponseBodyData) GetViewGroup() []*DescribeCostCheckResultsResponseBodyDataViewGroup {
	return s.ViewGroup
}

func (s *DescribeCostCheckResultsResponseBodyData) GetWarningCount() *int32 {
	return s.WarningCount
}

func (s *DescribeCostCheckResultsResponseBodyData) SetAdviceResourceCount(v int32) *DescribeCostCheckResultsResponseBodyData {
	s.AdviceResourceCount = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyData) SetGroupBy(v string) *DescribeCostCheckResultsResponseBodyData {
	s.GroupBy = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyData) SetNormalCount(v int32) *DescribeCostCheckResultsResponseBodyData {
	s.NormalCount = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyData) SetResourceCount(v int32) *DescribeCostCheckResultsResponseBodyData {
	s.ResourceCount = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyData) SetTotalCount(v int32) *DescribeCostCheckResultsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyData) SetViewGroup(v []*DescribeCostCheckResultsResponseBodyDataViewGroup) *DescribeCostCheckResultsResponseBodyData {
	s.ViewGroup = v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyData) SetWarningCount(v int32) *DescribeCostCheckResultsResponseBodyData {
	s.WarningCount = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyData) Validate() error {
	if s.ViewGroup != nil {
		for _, item := range s.ViewGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCostCheckResultsResponseBodyDataViewGroup struct {
	CheckItems []*DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems `json:"CheckItems,omitempty" xml:"CheckItems,omitempty" type:"Repeated"`
	// example:
	//
	// 22
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
	// example:
	//
	// 0
	GroupCount *int32 `json:"GroupCount,omitempty" xml:"GroupCount,omitempty"`
	// example:
	//
	// 1
	GroupExpectedSavingCost *float32 `json:"GroupExpectedSavingCost,omitempty" xml:"GroupExpectedSavingCost,omitempty"`
	// example:
	//
	// aut***8ainRh1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s DescribeCostCheckResultsResponseBodyDataViewGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostCheckResultsResponseBodyDataViewGroup) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroup) GetCheckItems() []*DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	return s.CheckItems
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroup) GetGroupCode() *string {
	return s.GroupCode
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroup) GetGroupCount() *int32 {
	return s.GroupCount
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroup) GetGroupExpectedSavingCost() *float32 {
	return s.GroupExpectedSavingCost
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroup) SetCheckItems(v []*DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) *DescribeCostCheckResultsResponseBodyDataViewGroup {
	s.CheckItems = v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroup) SetGroupCode(v string) *DescribeCostCheckResultsResponseBodyDataViewGroup {
	s.GroupCode = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroup) SetGroupCount(v int32) *DescribeCostCheckResultsResponseBodyDataViewGroup {
	s.GroupCount = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroup) SetGroupExpectedSavingCost(v float32) *DescribeCostCheckResultsResponseBodyDataViewGroup {
	s.GroupExpectedSavingCost = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroup) SetGroupName(v string) *DescribeCostCheckResultsResponseBodyDataViewGroup {
	s.GroupName = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroup) Validate() error {
	if s.CheckItems != nil {
		for _, item := range s.CheckItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems struct {
	// example:
	//
	// 100
	AdviceCount *int32 `json:"AdviceCount,omitempty" xml:"AdviceCount,omitempty"`
	// example:
	//
	// 1
	AdviceResourceCount *int32 `json:"AdviceResourceCount,omitempty" xml:"AdviceResourceCount,omitempty"`
	// example:
	//
	// 4
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// EbsCostIdleCheck
	CheckId   *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckName *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	// example:
	//
	// 1
	CurrentCost *float32 `json:"CurrentCost,omitempty" xml:"CurrentCost,omitempty"`
	Description *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1
	ExpectedSavingCost *float32 `json:"ExpectedSavingCost,omitempty" xml:"ExpectedSavingCost,omitempty"`
	// example:
	//
	// 1
	OptimizedCost *float32 `json:"OptimizedCost,omitempty" xml:"OptimizedCost,omitempty"`
	// example:
	//
	// slb
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// 1
	Severity *int32 `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// example:
	//
	// true
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Tips    *string `json:"Tips,omitempty" xml:"Tips,omitempty"`
}

func (s DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) GetAdviceCount() *int32 {
	return s.AdviceCount
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) GetAdviceResourceCount() *int32 {
	return s.AdviceResourceCount
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) GetCategory() *string {
	return s.Category
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) GetCheckId() *string {
	return s.CheckId
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) GetCheckName() *string {
	return s.CheckName
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) GetCurrentCost() *float32 {
	return s.CurrentCost
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) GetDescription() *string {
	return s.Description
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) GetExpectedSavingCost() *float32 {
	return s.ExpectedSavingCost
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) GetOptimizedCost() *float32 {
	return s.OptimizedCost
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) GetProduct() *string {
	return s.Product
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) GetSeverity() *int32 {
	return s.Severity
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) GetSummary() *string {
	return s.Summary
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) GetTips() *string {
	return s.Tips
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) SetAdviceCount(v int32) *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	s.AdviceCount = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) SetAdviceResourceCount(v int32) *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	s.AdviceResourceCount = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) SetCategory(v string) *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	s.Category = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) SetCheckId(v string) *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	s.CheckId = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) SetCheckName(v string) *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	s.CheckName = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) SetCurrentCost(v float32) *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	s.CurrentCost = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) SetDescription(v string) *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	s.Description = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) SetExpectedSavingCost(v float32) *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	s.ExpectedSavingCost = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) SetOptimizedCost(v float32) *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	s.OptimizedCost = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) SetProduct(v string) *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	s.Product = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) SetSeverity(v int32) *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	s.Severity = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) SetSummary(v string) *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	s.Summary = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) SetTips(v string) *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems {
	s.Tips = &v
	return s
}

func (s *DescribeCostCheckResultsResponseBodyDataViewGroupCheckItems) Validate() error {
	return dara.Validate(s)
}
