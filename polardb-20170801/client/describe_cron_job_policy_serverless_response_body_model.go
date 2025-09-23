// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCronJobPolicyServerlessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeCronJobPolicyServerlessResponseBodyItems) *DescribeCronJobPolicyServerlessResponseBody
	GetItems() []*DescribeCronJobPolicyServerlessResponseBodyItems
	SetPageNumber(v int32) *DescribeCronJobPolicyServerlessResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCronJobPolicyServerlessResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCronJobPolicyServerlessResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeCronJobPolicyServerlessResponseBody
	GetTotalRecordCount() *int32
}

type DescribeCronJobPolicyServerlessResponseBody struct {
	Items []*DescribeCronJobPolicyServerlessResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9B7BFB11-C077-4FE3-B051-F69CEB******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeCronJobPolicyServerlessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCronJobPolicyServerlessResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCronJobPolicyServerlessResponseBody) GetItems() []*DescribeCronJobPolicyServerlessResponseBodyItems {
	return s.Items
}

func (s *DescribeCronJobPolicyServerlessResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCronJobPolicyServerlessResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCronJobPolicyServerlessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCronJobPolicyServerlessResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeCronJobPolicyServerlessResponseBody) SetItems(v []*DescribeCronJobPolicyServerlessResponseBodyItems) *DescribeCronJobPolicyServerlessResponseBody {
	s.Items = v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponseBody) SetPageNumber(v int32) *DescribeCronJobPolicyServerlessResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponseBody) SetPageSize(v int32) *DescribeCronJobPolicyServerlessResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponseBody) SetRequestId(v string) *DescribeCronJobPolicyServerlessResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponseBody) SetTotalRecordCount(v int32) *DescribeCronJobPolicyServerlessResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCronJobPolicyServerlessResponseBodyItems struct {
	// example:
	//
	// ModifyDBClusterServerlessConf
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// true
	AllowShutDown *string `json:"AllowShutDown,omitempty" xml:"AllowShutDown,omitempty"`
	// example:
	//
	// 0 0 8 	- 	- ?
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// example:
	//
	// pc-*****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 2024-12-04T02:25:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 12eee3eb-60bd-40ac-a403-218e02eb99c7
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 254752088000354
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 2
	ScaleApRoNumMax *string `json:"ScaleApRoNumMax,omitempty" xml:"ScaleApRoNumMax,omitempty"`
	// example:
	//
	// 1
	ScaleApRoNumMin *string `json:"ScaleApRoNumMin,omitempty" xml:"ScaleApRoNumMin,omitempty"`
	// example:
	//
	// 9
	ScaleMax *string `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// example:
	//
	// 1
	ScaleMin *string `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// example:
	//
	// 3
	ScaleRoNumMax *string `json:"ScaleRoNumMax,omitempty" xml:"ScaleRoNumMax,omitempty"`
	// example:
	//
	// 2
	ScaleRoNumMin *string `json:"ScaleRoNumMin,omitempty" xml:"ScaleRoNumMin,omitempty"`
	// example:
	//
	// 1200
	SecondsUntilAutoPause *string `json:"SecondsUntilAutoPause,omitempty" xml:"SecondsUntilAutoPause,omitempty"`
	// example:
	//
	// 70
	ServerlessRuleCpuEnlargeThreshold *string `json:"ServerlessRuleCpuEnlargeThreshold,omitempty" xml:"ServerlessRuleCpuEnlargeThreshold,omitempty"`
	// example:
	//
	// 40
	ServerlessRuleCpuShrinkThreshold *string `json:"ServerlessRuleCpuShrinkThreshold,omitempty" xml:"ServerlessRuleCpuShrinkThreshold,omitempty"`
	// example:
	//
	// normal
	ServerlessRuleMode *string `json:"ServerlessRuleMode,omitempty" xml:"ServerlessRuleMode,omitempty"`
	// example:
	//
	// 2020-06-09T18:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 3
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCronJobPolicyServerlessResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeCronJobPolicyServerlessResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) GetAction() *string {
	return s.Action
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) GetAllowShutDown() *string {
	return s.AllowShutDown
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) GetCronExpression() *string {
	return s.CronExpression
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) GetJobId() *string {
	return s.JobId
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) GetOrderId() *string {
	return s.OrderId
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) GetScaleApRoNumMax() *string {
	return s.ScaleApRoNumMax
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) GetScaleApRoNumMin() *string {
	return s.ScaleApRoNumMin
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) GetScaleMax() *string {
	return s.ScaleMax
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) GetScaleMin() *string {
	return s.ScaleMin
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) GetScaleRoNumMax() *string {
	return s.ScaleRoNumMax
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) GetScaleRoNumMin() *string {
	return s.ScaleRoNumMin
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) GetSecondsUntilAutoPause() *string {
	return s.SecondsUntilAutoPause
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) GetServerlessRuleCpuEnlargeThreshold() *string {
	return s.ServerlessRuleCpuEnlargeThreshold
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) GetServerlessRuleCpuShrinkThreshold() *string {
	return s.ServerlessRuleCpuShrinkThreshold
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) GetServerlessRuleMode() *string {
	return s.ServerlessRuleMode
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) SetAction(v string) *DescribeCronJobPolicyServerlessResponseBodyItems {
	s.Action = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) SetAllowShutDown(v string) *DescribeCronJobPolicyServerlessResponseBodyItems {
	s.AllowShutDown = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) SetCronExpression(v string) *DescribeCronJobPolicyServerlessResponseBodyItems {
	s.CronExpression = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) SetDBClusterId(v string) *DescribeCronJobPolicyServerlessResponseBodyItems {
	s.DBClusterId = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) SetEndTime(v string) *DescribeCronJobPolicyServerlessResponseBodyItems {
	s.EndTime = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) SetJobId(v string) *DescribeCronJobPolicyServerlessResponseBodyItems {
	s.JobId = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) SetOrderId(v string) *DescribeCronJobPolicyServerlessResponseBodyItems {
	s.OrderId = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) SetRegionId(v string) *DescribeCronJobPolicyServerlessResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) SetScaleApRoNumMax(v string) *DescribeCronJobPolicyServerlessResponseBodyItems {
	s.ScaleApRoNumMax = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) SetScaleApRoNumMin(v string) *DescribeCronJobPolicyServerlessResponseBodyItems {
	s.ScaleApRoNumMin = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) SetScaleMax(v string) *DescribeCronJobPolicyServerlessResponseBodyItems {
	s.ScaleMax = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) SetScaleMin(v string) *DescribeCronJobPolicyServerlessResponseBodyItems {
	s.ScaleMin = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) SetScaleRoNumMax(v string) *DescribeCronJobPolicyServerlessResponseBodyItems {
	s.ScaleRoNumMax = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) SetScaleRoNumMin(v string) *DescribeCronJobPolicyServerlessResponseBodyItems {
	s.ScaleRoNumMin = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) SetSecondsUntilAutoPause(v string) *DescribeCronJobPolicyServerlessResponseBodyItems {
	s.SecondsUntilAutoPause = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) SetServerlessRuleCpuEnlargeThreshold(v string) *DescribeCronJobPolicyServerlessResponseBodyItems {
	s.ServerlessRuleCpuEnlargeThreshold = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) SetServerlessRuleCpuShrinkThreshold(v string) *DescribeCronJobPolicyServerlessResponseBodyItems {
	s.ServerlessRuleCpuShrinkThreshold = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) SetServerlessRuleMode(v string) *DescribeCronJobPolicyServerlessResponseBodyItems {
	s.ServerlessRuleMode = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) SetStartTime(v string) *DescribeCronJobPolicyServerlessResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) SetStatus(v string) *DescribeCronJobPolicyServerlessResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
