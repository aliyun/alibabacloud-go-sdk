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
	// The list of tasks.
	Items []*DescribeCronJobPolicyServerlessResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records on each page. Valid values:
	//
	// - **30**
	//
	// - **50**
	//
	// - **100**
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9B7BFB11-C077-4FE3-B051-F69CEB******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
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

type DescribeCronJobPolicyServerlessResponseBodyItems struct {
	// A system parameter. Set the value to **ModifyDBClusterServerlessConf**.
	//
	// example:
	//
	// ModifyDBClusterServerlessConf
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// Specifies whether to enable No-activity Suspension. Valid values:
	//
	// - **true**: enabled
	//
	// - **false**: disabled (default)
	//
	// example:
	//
	// true
	AllowShutDown *string `json:"AllowShutDown,omitempty" xml:"AllowShutDown,omitempty"`
	// The Cron expression for the scheduled task.
	//
	// example:
	//
	// 0 0 8 	- 	- ?
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// pc-*****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the task. The time is in the yyyy-MM-ddTHH:mm:ssZ format and in UTC.
	//
	// example:
	//
	// 2024-12-04T02:25:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the scheduled task.
	//
	// example:
	//
	// 12eee3eb-60bd-40ac-a403-218e02eb99c7
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 254752088000354
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The maximum number of read-only IMCI nodes. Valid values: 1 to 15.
	//
	// example:
	//
	// 2
	ScaleApRoNumMax *string `json:"ScaleApRoNumMax,omitempty" xml:"ScaleApRoNumMax,omitempty"`
	// The minimum number of read-only IMCI nodes. Valid values: 0 to 15.
	//
	// example:
	//
	// 1
	ScaleApRoNumMin *string `json:"ScaleApRoNumMin,omitempty" xml:"ScaleApRoNumMin,omitempty"`
	// The maximum capacity. The value must be from 1 to 32. Unit: PCU.
	//
	// example:
	//
	// 9
	ScaleMax *string `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// The minimum capacity. The value must be from 0.25 to 32 and less than or equal to the maximum capacity. Unit: PolarDB Capacity Unit (PCU).
	//
	// example:
	//
	// 1
	ScaleMin *string `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// The maximum number of read-only nodes. The value must be greater than or equal to the minimum value. Valid values: 0 to 15.
	//
	// example:
	//
	// 3
	ScaleRoNumMax *string `json:"ScaleRoNumMax,omitempty" xml:"ScaleRoNumMax,omitempty"`
	// The minimum number of read-only nodes. Valid values: 0 to 15.
	//
	// example:
	//
	// 2
	ScaleRoNumMin *string `json:"ScaleRoNumMin,omitempty" xml:"ScaleRoNumMin,omitempty"`
	// The detection period for No-activity Suspension. The value must be a multiple of 5. Valid values: 5 to 1440. Unit: minutes.
	//
	// example:
	//
	// 1200
	SecondsUntilAutoPause *string `json:"SecondsUntilAutoPause,omitempty" xml:"SecondsUntilAutoPause,omitempty"`
	// The CPU utilization threshold for scaling up. Valid values: 40 to 100. Unit: %.
	//
	// example:
	//
	// 70
	ServerlessRuleCpuEnlargeThreshold *string `json:"ServerlessRuleCpuEnlargeThreshold,omitempty" xml:"ServerlessRuleCpuEnlargeThreshold,omitempty"`
	// The CPU utilization threshold for scaling down. Valid values: 10 to 100. Unit: %. The difference between the scale-up threshold and the scale-down threshold must be 30 or greater.
	//
	// example:
	//
	// 40
	ServerlessRuleCpuShrinkThreshold *string `json:"ServerlessRuleCpuShrinkThreshold,omitempty" xml:"ServerlessRuleCpuShrinkThreshold,omitempty"`
	// The elasticity sensitivity. Valid values:
	//
	// - normal: standard
	//
	// - flexible: sensitive
	//
	// example:
	//
	// normal
	ServerlessRuleMode *string `json:"ServerlessRuleMode,omitempty" xml:"ServerlessRuleMode,omitempty"`
	// The start time of the task. The time is in the `yyyy-MM-ddTHH:mmZ` format and in UTC.
	//
	// example:
	//
	// 2020-06-09T18:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the task. Valid values:
	//
	// - **working**: The scheduled task is running.
	//
	// - **finish**: The scheduled task is complete.
	//
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
