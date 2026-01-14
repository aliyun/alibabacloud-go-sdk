// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationMonitors(v []*ListApplicationMonitorResponseBodyApplicationMonitors) *ListApplicationMonitorResponseBody
	GetApplicationMonitors() []*ListApplicationMonitorResponseBodyApplicationMonitors
	SetPageNumber(v int32) *ListApplicationMonitorResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListApplicationMonitorResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListApplicationMonitorResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListApplicationMonitorResponseBody
	GetTotalCount() *int32
}

type ListApplicationMonitorResponseBody struct {
	// The list of origin probing tasks.
	ApplicationMonitors []*ListApplicationMonitorResponseBodyApplicationMonitors `json:"ApplicationMonitors,omitempty" xml:"ApplicationMonitors,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationMonitorResponseBody) GetApplicationMonitors() []*ListApplicationMonitorResponseBodyApplicationMonitors {
	return s.ApplicationMonitors
}

func (s *ListApplicationMonitorResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListApplicationMonitorResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApplicationMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationMonitorResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListApplicationMonitorResponseBody) SetApplicationMonitors(v []*ListApplicationMonitorResponseBodyApplicationMonitors) *ListApplicationMonitorResponseBody {
	s.ApplicationMonitors = v
	return s
}

func (s *ListApplicationMonitorResponseBody) SetPageNumber(v int32) *ListApplicationMonitorResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationMonitorResponseBody) SetPageSize(v int32) *ListApplicationMonitorResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListApplicationMonitorResponseBody) SetRequestId(v string) *ListApplicationMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationMonitorResponseBody) SetTotalCount(v int32) *ListApplicationMonitorResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationMonitorResponseBody) Validate() error {
	if s.ApplicationMonitors != nil {
		for _, item := range s.ApplicationMonitors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationMonitorResponseBodyApplicationMonitors struct {
	// The ID of the GA instance on which the origin probing task runs.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The URL or IP address that was probed.
	//
	// example:
	//
	// https://www.aliyun.com
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// Indicates whether the automatic diagnostics feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	DetectEnable *bool `json:"DetectEnable,omitempty" xml:"DetectEnable,omitempty"`
	// The threshold that is used to trigger the automatic diagnostics feature.
	//
	// example:
	//
	// １
	DetectThreshold *int32 `json:"DetectThreshold,omitempty" xml:"DetectThreshold,omitempty"`
	// The number of times that are required to reach the threshold before the automatic diagnostics feature can be triggered.
	//
	// example:
	//
	// １
	DetectTimes *int32 `json:"DetectTimes,omitempty" xml:"DetectTimes,omitempty"`
	// The ID of the listener on which the origin probing task runs.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The extended options of the listener protocol that is used by the origin probing task. The options vary based on the listener protocol.
	//
	// example:
	//
	// { "http_method": "get","header": "key:asd","acceptable_response_code": "500","cert_verify": true }
	OptionsJson *string `json:"OptionsJson,omitempty" xml:"OptionsJson,omitempty"`
	// The silence period of the automatic diagnostics feature. This parameter indicates the interval at which the automatic diagnostics feature is triggered. If the availability rate does not return to normal after GA triggers an automatic diagnostic task, GA must wait until the silence period ends before GA can trigger another automatic diagnostic task.
	//
	// If the number of consecutive times that the availability rate drops below the threshold of automatic diagnostics reaches the value of **DetectTimes*	- , the automatic diagnostics feature is triggered. The automatic diagnostics feature is not triggered again within the silence period even if the availability rate stays below the threshold. If the availability rate does not return to normal after the silence period ends, the automatic diagnostics feature is triggered again.
	//
	// Unit: seconds.
	//
	// example:
	//
	// 300
	SilenceTime *int32 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The status of the origin probing task. Valid values:
	//
	// 	- **active:*	- The origin probing task is running.
	//
	// 	- **inactive:*	- The origin probing task is stopped.
	//
	// 	- **init:*	- The origin probing task is being initialized.
	//
	// 	- **deleting:*	- The origin probing task is being deleted.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The origin probing task ID.
	//
	// example:
	//
	// sm-bp1fpdjfju9k8yr1y****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The origin probing task name.
	//
	// example:
	//
	// task1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ListApplicationMonitorResponseBodyApplicationMonitors) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationMonitorResponseBodyApplicationMonitors) GoString() string {
	return s.String()
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) GetAddress() *string {
	return s.Address
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) GetDetectEnable() *bool {
	return s.DetectEnable
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) GetDetectThreshold() *int32 {
	return s.DetectThreshold
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) GetDetectTimes() *int32 {
	return s.DetectTimes
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) GetOptionsJson() *string {
	return s.OptionsJson
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) GetSilenceTime() *int32 {
	return s.SilenceTime
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) GetState() *string {
	return s.State
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) GetTaskId() *string {
	return s.TaskId
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) GetTaskName() *string {
	return s.TaskName
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) SetAcceleratorId(v string) *ListApplicationMonitorResponseBodyApplicationMonitors {
	s.AcceleratorId = &v
	return s
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) SetAddress(v string) *ListApplicationMonitorResponseBodyApplicationMonitors {
	s.Address = &v
	return s
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) SetDetectEnable(v bool) *ListApplicationMonitorResponseBodyApplicationMonitors {
	s.DetectEnable = &v
	return s
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) SetDetectThreshold(v int32) *ListApplicationMonitorResponseBodyApplicationMonitors {
	s.DetectThreshold = &v
	return s
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) SetDetectTimes(v int32) *ListApplicationMonitorResponseBodyApplicationMonitors {
	s.DetectTimes = &v
	return s
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) SetListenerId(v string) *ListApplicationMonitorResponseBodyApplicationMonitors {
	s.ListenerId = &v
	return s
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) SetOptionsJson(v string) *ListApplicationMonitorResponseBodyApplicationMonitors {
	s.OptionsJson = &v
	return s
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) SetSilenceTime(v int32) *ListApplicationMonitorResponseBodyApplicationMonitors {
	s.SilenceTime = &v
	return s
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) SetState(v string) *ListApplicationMonitorResponseBodyApplicationMonitors {
	s.State = &v
	return s
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) SetTaskId(v string) *ListApplicationMonitorResponseBodyApplicationMonitors {
	s.TaskId = &v
	return s
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) SetTaskName(v string) *ListApplicationMonitorResponseBodyApplicationMonitors {
	s.TaskName = &v
	return s
}

func (s *ListApplicationMonitorResponseBodyApplicationMonitors) Validate() error {
	return dara.Validate(s)
}
