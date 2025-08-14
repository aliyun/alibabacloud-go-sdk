// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecommendTaskPageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRecommendTaskPageListResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeRecommendTaskPageListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeRecommendTaskPageListResponseBody
	GetPageSize() *int32
	SetResultObject(v []*DescribeRecommendTaskPageListResponseBodyResultObject) *DescribeRecommendTaskPageListResponseBody
	GetResultObject() []*DescribeRecommendTaskPageListResponseBodyResultObject
	SetTotalItem(v int32) *DescribeRecommendTaskPageListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeRecommendTaskPageListResponseBody
	GetTotalPage() *int32
}

type DescribeRecommendTaskPageListResponseBody struct {
	// Request ID
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Current page number
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size, with a default value of 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Returned object
	ResultObject []*DescribeRecommendTaskPageListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items
	//
	// example:
	//
	// 3
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeRecommendTaskPageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendTaskPageListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecommendTaskPageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRecommendTaskPageListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeRecommendTaskPageListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRecommendTaskPageListResponseBody) GetResultObject() []*DescribeRecommendTaskPageListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeRecommendTaskPageListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeRecommendTaskPageListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeRecommendTaskPageListResponseBody) SetRequestId(v string) *DescribeRecommendTaskPageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecommendTaskPageListResponseBody) SetCurrentPage(v int32) *DescribeRecommendTaskPageListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRecommendTaskPageListResponseBody) SetPageSize(v int32) *DescribeRecommendTaskPageListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRecommendTaskPageListResponseBody) SetResultObject(v []*DescribeRecommendTaskPageListResponseBodyResultObject) *DescribeRecommendTaskPageListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeRecommendTaskPageListResponseBody) SetTotalItem(v int32) *DescribeRecommendTaskPageListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeRecommendTaskPageListResponseBody) SetTotalPage(v int32) *DescribeRecommendTaskPageListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeRecommendTaskPageListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRecommendTaskPageListResponseBodyResultObject struct {
	// Impact indicators
	ExpectVelocities []*string `json:"expectVelocities,omitempty" xml:"expectVelocities,omitempty" type:"Repeated"`
	// Creation time.
	//
	// example:
	//
	// 1621578648000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// Number of samples
	//
	// example:
	//
	// 100
	NormalCount *int64 `json:"normalCount,omitempty" xml:"normalCount,omitempty"`
	// Number of normal samples
	//
	// example:
	//
	// 10
	NormalSize *int64 `json:"normalSize,omitempty" xml:"normalSize,omitempty"`
	// Number of records displayed on the current page.
	//
	// example:
	//
	// 100
	RiskCount *int64 `json:"riskCount,omitempty" xml:"riskCount,omitempty"`
	// Number of risk samples
	//
	// example:
	//
	// 5
	RiskSize *int64 `json:"riskSize,omitempty" xml:"riskSize,omitempty"`
	// Sample name
	//
	// example:
	//
	// 白样本
	SampleName *string `json:"sampleName,omitempty" xml:"sampleName,omitempty"`
	// Sample scenario
	//
	// example:
	//
	// account_abuse_detection
	SampleScene *string `json:"sampleScene,omitempty" xml:"sampleScene,omitempty"`
	// Task ID.
	//
	// example:
	//
	// 240c93ddffa74e38be3a00375eb3041d
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// Task name.
	//
	// example:
	//
	// 策略推荐任务
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
	// Task status.
	//
	// example:
	//
	// CREATE
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
}

func (s DescribeRecommendTaskPageListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendTaskPageListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeRecommendTaskPageListResponseBodyResultObject) GetExpectVelocities() []*string {
	return s.ExpectVelocities
}

func (s *DescribeRecommendTaskPageListResponseBodyResultObject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeRecommendTaskPageListResponseBodyResultObject) GetNormalCount() *int64 {
	return s.NormalCount
}

func (s *DescribeRecommendTaskPageListResponseBodyResultObject) GetNormalSize() *int64 {
	return s.NormalSize
}

func (s *DescribeRecommendTaskPageListResponseBodyResultObject) GetRiskCount() *int64 {
	return s.RiskCount
}

func (s *DescribeRecommendTaskPageListResponseBodyResultObject) GetRiskSize() *int64 {
	return s.RiskSize
}

func (s *DescribeRecommendTaskPageListResponseBodyResultObject) GetSampleName() *string {
	return s.SampleName
}

func (s *DescribeRecommendTaskPageListResponseBodyResultObject) GetSampleScene() *string {
	return s.SampleScene
}

func (s *DescribeRecommendTaskPageListResponseBodyResultObject) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeRecommendTaskPageListResponseBodyResultObject) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeRecommendTaskPageListResponseBodyResultObject) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeRecommendTaskPageListResponseBodyResultObject) SetExpectVelocities(v []*string) *DescribeRecommendTaskPageListResponseBodyResultObject {
	s.ExpectVelocities = v
	return s
}

func (s *DescribeRecommendTaskPageListResponseBodyResultObject) SetGmtCreate(v int64) *DescribeRecommendTaskPageListResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRecommendTaskPageListResponseBodyResultObject) SetNormalCount(v int64) *DescribeRecommendTaskPageListResponseBodyResultObject {
	s.NormalCount = &v
	return s
}

func (s *DescribeRecommendTaskPageListResponseBodyResultObject) SetNormalSize(v int64) *DescribeRecommendTaskPageListResponseBodyResultObject {
	s.NormalSize = &v
	return s
}

func (s *DescribeRecommendTaskPageListResponseBodyResultObject) SetRiskCount(v int64) *DescribeRecommendTaskPageListResponseBodyResultObject {
	s.RiskCount = &v
	return s
}

func (s *DescribeRecommendTaskPageListResponseBodyResultObject) SetRiskSize(v int64) *DescribeRecommendTaskPageListResponseBodyResultObject {
	s.RiskSize = &v
	return s
}

func (s *DescribeRecommendTaskPageListResponseBodyResultObject) SetSampleName(v string) *DescribeRecommendTaskPageListResponseBodyResultObject {
	s.SampleName = &v
	return s
}

func (s *DescribeRecommendTaskPageListResponseBodyResultObject) SetSampleScene(v string) *DescribeRecommendTaskPageListResponseBodyResultObject {
	s.SampleScene = &v
	return s
}

func (s *DescribeRecommendTaskPageListResponseBodyResultObject) SetTaskId(v int64) *DescribeRecommendTaskPageListResponseBodyResultObject {
	s.TaskId = &v
	return s
}

func (s *DescribeRecommendTaskPageListResponseBodyResultObject) SetTaskName(v string) *DescribeRecommendTaskPageListResponseBodyResultObject {
	s.TaskName = &v
	return s
}

func (s *DescribeRecommendTaskPageListResponseBodyResultObject) SetTaskStatus(v string) *DescribeRecommendTaskPageListResponseBodyResultObject {
	s.TaskStatus = &v
	return s
}

func (s *DescribeRecommendTaskPageListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
