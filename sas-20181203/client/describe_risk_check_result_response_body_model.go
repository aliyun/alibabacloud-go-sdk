// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskCheckResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeRiskCheckResultResponseBody
	GetCount() *int32
	SetCurrentPage(v int32) *DescribeRiskCheckResultResponseBody
	GetCurrentPage() *int32
	SetList(v []*DescribeRiskCheckResultResponseBodyList) *DescribeRiskCheckResultResponseBody
	GetList() []*DescribeRiskCheckResultResponseBodyList
	SetPageCount(v int32) *DescribeRiskCheckResultResponseBody
	GetPageCount() *int32
	SetPageSize(v int32) *DescribeRiskCheckResultResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeRiskCheckResultResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeRiskCheckResultResponseBody
	GetTotalCount() *int32
}

type DescribeRiskCheckResultResponseBody struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The check items.
	List []*DescribeRiskCheckResultResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The total number of pages returned.
	//
	// example:
	//
	// 20
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// The number of entries returned per page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// AD271C07-4ACE-413D-AA9B-F14FD3B7717F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRiskCheckResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskCheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckResultResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeRiskCheckResultResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeRiskCheckResultResponseBody) GetList() []*DescribeRiskCheckResultResponseBodyList {
	return s.List
}

func (s *DescribeRiskCheckResultResponseBody) GetPageCount() *int32 {
	return s.PageCount
}

func (s *DescribeRiskCheckResultResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRiskCheckResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRiskCheckResultResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRiskCheckResultResponseBody) SetCount(v int32) *DescribeRiskCheckResultResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBody) SetCurrentPage(v int32) *DescribeRiskCheckResultResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBody) SetList(v []*DescribeRiskCheckResultResponseBodyList) *DescribeRiskCheckResultResponseBody {
	s.List = v
	return s
}

func (s *DescribeRiskCheckResultResponseBody) SetPageCount(v int32) *DescribeRiskCheckResultResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBody) SetPageSize(v int32) *DescribeRiskCheckResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBody) SetRequestId(v string) *DescribeRiskCheckResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBody) SetTotalCount(v int32) *DescribeRiskCheckResultResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRiskCheckResultResponseBodyList struct {
	// The number of affected assets.
	//
	// example:
	//
	// 0
	AffectedCount *int32 `json:"AffectedCount,omitempty" xml:"AffectedCount,omitempty"`
	// The timestamp when the last check was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1639429164000
	CheckTime *int64 `json:"CheckTime,omitempty" xml:"CheckTime,omitempty"`
	// The ID of the check item. For more information about the check item, see the check item table in the "Response parameters" section of this topic.
	//
	// example:
	//
	// 1
	ItemId *int64 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// The time when the next check will be performed.
	//
	// example:
	//
	// 0
	RemainingTime *int32 `json:"RemainingTime,omitempty" xml:"RemainingTime,omitempty"`
	// Indicates whether the risks that are detected based on the check item can be fixed. Valid values:
	//
	// 	- **enabled**: yes
	//
	// 	- **disabled**: no
	//
	// example:
	//
	// disabled
	RepairStatus *string `json:"RepairStatus,omitempty" xml:"RepairStatus,omitempty"`
	// The type of the affected assets.
	//
	// example:
	//
	// ECS
	RiskAssertType *string `json:"RiskAssertType,omitempty" xml:"RiskAssertType,omitempty"`
	// An array that consists of the details about the check item.
	RiskItemResources []*DescribeRiskCheckResultResponseBodyListRiskItemResources `json:"RiskItemResources,omitempty" xml:"RiskItemResources,omitempty" type:"Repeated"`
	// The risk level of the check item. Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The sequence number in the check results. The check items are sorted based on the sequence number.
	//
	// example:
	//
	// 1
	Sort *int32 `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// Indicates whether the check item is supported by the edition of Security Center that you purchase. Valid values:
	//
	// 	- **enabled**: yes
	//
	// 	- **disable**: no
	//
	// example:
	//
	// enabled
	StartStatus *string `json:"StartStatus,omitempty" xml:"StartStatus,omitempty"`
	// The status of the check results. Valid values:
	//
	// 	- **pass**
	//
	// 	- **failed**
	//
	// 	- **running**
	//
	// 	- **waiting**
	//
	// 	- **ignored**
	//
	// 	- **falsePositive**
	//
	// example:
	//
	// pass
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the check task.
	//
	// example:
	//
	// 15384933
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the check item.
	//
	// example:
	//
	// RDS - Whitelist Configuration
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The type of the check item. Valid values:
	//
	// 	- Identity authentication and permissions
	//
	// 	- Network access control
	//
	// 	- Log audit
	//
	// 	- Data security
	//
	// 	- Monitoring and alerting
	//
	// 	- Basic security protection
	//
	// example:
	//
	// Log audit
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeRiskCheckResultResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskCheckResultResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckResultResponseBodyList) GetAffectedCount() *int32 {
	return s.AffectedCount
}

func (s *DescribeRiskCheckResultResponseBodyList) GetCheckTime() *int64 {
	return s.CheckTime
}

func (s *DescribeRiskCheckResultResponseBodyList) GetItemId() *int64 {
	return s.ItemId
}

func (s *DescribeRiskCheckResultResponseBodyList) GetRemainingTime() *int32 {
	return s.RemainingTime
}

func (s *DescribeRiskCheckResultResponseBodyList) GetRepairStatus() *string {
	return s.RepairStatus
}

func (s *DescribeRiskCheckResultResponseBodyList) GetRiskAssertType() *string {
	return s.RiskAssertType
}

func (s *DescribeRiskCheckResultResponseBodyList) GetRiskItemResources() []*DescribeRiskCheckResultResponseBodyListRiskItemResources {
	return s.RiskItemResources
}

func (s *DescribeRiskCheckResultResponseBodyList) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeRiskCheckResultResponseBodyList) GetSort() *int32 {
	return s.Sort
}

func (s *DescribeRiskCheckResultResponseBodyList) GetStartStatus() *string {
	return s.StartStatus
}

func (s *DescribeRiskCheckResultResponseBodyList) GetStatus() *string {
	return s.Status
}

func (s *DescribeRiskCheckResultResponseBodyList) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeRiskCheckResultResponseBodyList) GetTitle() *string {
	return s.Title
}

func (s *DescribeRiskCheckResultResponseBodyList) GetType() *string {
	return s.Type
}

func (s *DescribeRiskCheckResultResponseBodyList) SetAffectedCount(v int32) *DescribeRiskCheckResultResponseBodyList {
	s.AffectedCount = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) SetCheckTime(v int64) *DescribeRiskCheckResultResponseBodyList {
	s.CheckTime = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) SetItemId(v int64) *DescribeRiskCheckResultResponseBodyList {
	s.ItemId = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) SetRemainingTime(v int32) *DescribeRiskCheckResultResponseBodyList {
	s.RemainingTime = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) SetRepairStatus(v string) *DescribeRiskCheckResultResponseBodyList {
	s.RepairStatus = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) SetRiskAssertType(v string) *DescribeRiskCheckResultResponseBodyList {
	s.RiskAssertType = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) SetRiskItemResources(v []*DescribeRiskCheckResultResponseBodyListRiskItemResources) *DescribeRiskCheckResultResponseBodyList {
	s.RiskItemResources = v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) SetRiskLevel(v string) *DescribeRiskCheckResultResponseBodyList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) SetSort(v int32) *DescribeRiskCheckResultResponseBodyList {
	s.Sort = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) SetStartStatus(v string) *DescribeRiskCheckResultResponseBodyList {
	s.StartStatus = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) SetStatus(v string) *DescribeRiskCheckResultResponseBodyList {
	s.Status = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) SetTaskId(v int64) *DescribeRiskCheckResultResponseBodyList {
	s.TaskId = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) SetTitle(v string) *DescribeRiskCheckResultResponseBodyList {
	s.Title = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) SetType(v string) *DescribeRiskCheckResultResponseBodyList {
	s.Type = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) Validate() error {
	if s.RiskItemResources != nil {
		for _, item := range s.RiskItemResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRiskCheckResultResponseBodyListRiskItemResources struct {
	// The details about the check results.
	ContentResource map[string]interface{} `json:"ContentResource,omitempty" xml:"ContentResource,omitempty"`
	// The title in the details. Valid values:
	//
	// 	- **bestPractice**: description
	//
	// 	- **influence**: risk
	//
	// 	- **suggestion**: solution
	//
	// 	- **helpResource**: reference
	//
	// example:
	//
	// bestPractice
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
}

func (s DescribeRiskCheckResultResponseBodyListRiskItemResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskCheckResultResponseBodyListRiskItemResources) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckResultResponseBodyListRiskItemResources) GetContentResource() map[string]interface{} {
	return s.ContentResource
}

func (s *DescribeRiskCheckResultResponseBodyListRiskItemResources) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribeRiskCheckResultResponseBodyListRiskItemResources) SetContentResource(v map[string]interface{}) *DescribeRiskCheckResultResponseBodyListRiskItemResources {
	s.ContentResource = v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyListRiskItemResources) SetResourceName(v string) *DescribeRiskCheckResultResponseBodyListRiskItemResources {
	s.ResourceName = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyListRiskItemResources) Validate() error {
	return dara.Validate(s)
}
