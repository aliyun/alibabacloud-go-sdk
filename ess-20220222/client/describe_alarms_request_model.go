// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlarmsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmTaskId(v string) *DescribeAlarmsRequest
	GetAlarmTaskId() *string
	SetIsEnable(v bool) *DescribeAlarmsRequest
	GetIsEnable() *bool
	SetMetricName(v string) *DescribeAlarmsRequest
	GetMetricName() *string
	SetMetricType(v string) *DescribeAlarmsRequest
	GetMetricType() *string
	SetOwnerId(v int64) *DescribeAlarmsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeAlarmsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAlarmsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeAlarmsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAlarmsRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *DescribeAlarmsRequest
	GetScalingGroupId() *string
	SetState(v string) *DescribeAlarmsRequest
	GetState() *string
}

type DescribeAlarmsRequest struct {
	// The ID of the event-triggered task.
	//
	// example:
	//
	// asg-bp1hvbnmkl10vll5****_f95ce797-dc2e-4bad-9618-14fee7d1****
	AlarmTaskId *string `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
	// Specifies whether to enable the event-triggered task. Valid values:
	//
	// 	- true: enables the event-triggered task.
	//
	// 	- false: disables the event-triggered task.
	//
	// example:
	//
	// true
	IsEnable *bool `json:"IsEnable,omitempty" xml:"IsEnable,omitempty"`
	// The metric name.
	//
	// example:
	//
	// CpuUtilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The metric type. Valid values:
	//
	// 	- system: a system metric of CloudMonitor
	//
	// 	- custom: a custom metric that is reported to CloudMonitor.
	//
	// 	- hybrid: a metric of Hybrid Cloud Monitoring.
	//
	// example:
	//
	// true
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the event-triggered task.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the scaling group with which the event-triggered task is associated.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The status of the event-triggered task. Valid values:
	//
	// 	- ALARM: The alert condition is met and an alert is triggered.
	//
	// 	- OK: The alert condition is not met.
	//
	// 	- INSUFFICIENT_DATA: Auto Scaling cannot determine whether the alert condition is met due to insufficient data.
	//
	// example:
	//
	// OK
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeAlarmsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlarmsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsRequest) GetAlarmTaskId() *string {
	return s.AlarmTaskId
}

func (s *DescribeAlarmsRequest) GetIsEnable() *bool {
	return s.IsEnable
}

func (s *DescribeAlarmsRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeAlarmsRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *DescribeAlarmsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAlarmsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAlarmsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAlarmsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAlarmsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAlarmsRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeAlarmsRequest) GetState() *string {
	return s.State
}

func (s *DescribeAlarmsRequest) SetAlarmTaskId(v string) *DescribeAlarmsRequest {
	s.AlarmTaskId = &v
	return s
}

func (s *DescribeAlarmsRequest) SetIsEnable(v bool) *DescribeAlarmsRequest {
	s.IsEnable = &v
	return s
}

func (s *DescribeAlarmsRequest) SetMetricName(v string) *DescribeAlarmsRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeAlarmsRequest) SetMetricType(v string) *DescribeAlarmsRequest {
	s.MetricType = &v
	return s
}

func (s *DescribeAlarmsRequest) SetOwnerId(v int64) *DescribeAlarmsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAlarmsRequest) SetPageNumber(v int32) *DescribeAlarmsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAlarmsRequest) SetPageSize(v int32) *DescribeAlarmsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAlarmsRequest) SetRegionId(v string) *DescribeAlarmsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlarmsRequest) SetResourceOwnerAccount(v string) *DescribeAlarmsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAlarmsRequest) SetScalingGroupId(v string) *DescribeAlarmsRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeAlarmsRequest) SetState(v string) *DescribeAlarmsRequest {
	s.State = &v
	return s
}

func (s *DescribeAlarmsRequest) Validate() error {
	return dara.Validate(s)
}
