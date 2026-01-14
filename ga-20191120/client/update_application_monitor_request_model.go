// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *UpdateApplicationMonitorRequest
	GetAddress() *string
	SetClientToken(v string) *UpdateApplicationMonitorRequest
	GetClientToken() *string
	SetDetectEnable(v bool) *UpdateApplicationMonitorRequest
	GetDetectEnable() *bool
	SetDetectThreshold(v int32) *UpdateApplicationMonitorRequest
	GetDetectThreshold() *int32
	SetDetectTimes(v int32) *UpdateApplicationMonitorRequest
	GetDetectTimes() *int32
	SetListenerId(v string) *UpdateApplicationMonitorRequest
	GetListenerId() *string
	SetOptionsJson(v string) *UpdateApplicationMonitorRequest
	GetOptionsJson() *string
	SetRegionId(v string) *UpdateApplicationMonitorRequest
	GetRegionId() *string
	SetSilenceTime(v int32) *UpdateApplicationMonitorRequest
	GetSilenceTime() *int32
	SetTaskId(v string) *UpdateApplicationMonitorRequest
	GetTaskId() *string
	SetTaskName(v string) *UpdateApplicationMonitorRequest
	GetTaskName() *string
}

type UpdateApplicationMonitorRequest struct {
	// The URL or IP address that is probed.
	//
	// example:
	//
	// https://www.aliyun.com
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable the automatic diagnostics feature. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	DetectEnable *bool `json:"DetectEnable,omitempty" xml:"DetectEnable,omitempty"`
	// Set the threshold that is used to trigger the automatic diagnostics feature. If the liveness of the origin in percentile drops below the specified threshold, the automatic diagnostics feature is triggered.
	//
	// Valid values: **0*	- to **100**.
	//
	// example:
	//
	// 0
	DetectThreshold *int32 `json:"DetectThreshold,omitempty" xml:"DetectThreshold,omitempty"`
	// The number of times that the threshold must be reached before the automatic diagnostics feature is triggered.
	//
	// Valid values: **1*	- to **20**.
	//
	// example:
	//
	// 1
	DetectTimes *int32 `json:"DetectTimes,omitempty" xml:"DetectTimes,omitempty"`
	// The ID of the listener that you want to modify. The listener runs the origin probing task.
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
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The silence period of the automatic diagnostics feature. This parameter specifies the interval at which the automatic diagnostics feature is triggered. If the availability rate does not return to normal after GA triggers an automatic diagnostic task, GA must wait until the silence period ends before GA can trigger another automatic diagnostic task.
	//
	// If the number of consecutive times that the availability rate drops below the threshold of automatic diagnostics reaches the value of the **DetectTimes*	- parameter, the automatic diagnostics feature is triggered. The automatic diagnostics feature is not triggered again within the silence period even if the availability rate stays below the threshold. If the availability rate does not return to normal after the silence period ends, the automatic diagnostics feature is triggered again.
	//
	// Unit: seconds. Valid values: **300*	- to **86400**.
	//
	// example:
	//
	// 300
	SilenceTime *int32 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The ID of the origin probing task that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// sm-bp1fpdjfju9k8yr1y****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the origin probing task.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// task1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s UpdateApplicationMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationMonitorRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationMonitorRequest) GetAddress() *string {
	return s.Address
}

func (s *UpdateApplicationMonitorRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateApplicationMonitorRequest) GetDetectEnable() *bool {
	return s.DetectEnable
}

func (s *UpdateApplicationMonitorRequest) GetDetectThreshold() *int32 {
	return s.DetectThreshold
}

func (s *UpdateApplicationMonitorRequest) GetDetectTimes() *int32 {
	return s.DetectTimes
}

func (s *UpdateApplicationMonitorRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *UpdateApplicationMonitorRequest) GetOptionsJson() *string {
	return s.OptionsJson
}

func (s *UpdateApplicationMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateApplicationMonitorRequest) GetSilenceTime() *int32 {
	return s.SilenceTime
}

func (s *UpdateApplicationMonitorRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateApplicationMonitorRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *UpdateApplicationMonitorRequest) SetAddress(v string) *UpdateApplicationMonitorRequest {
	s.Address = &v
	return s
}

func (s *UpdateApplicationMonitorRequest) SetClientToken(v string) *UpdateApplicationMonitorRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateApplicationMonitorRequest) SetDetectEnable(v bool) *UpdateApplicationMonitorRequest {
	s.DetectEnable = &v
	return s
}

func (s *UpdateApplicationMonitorRequest) SetDetectThreshold(v int32) *UpdateApplicationMonitorRequest {
	s.DetectThreshold = &v
	return s
}

func (s *UpdateApplicationMonitorRequest) SetDetectTimes(v int32) *UpdateApplicationMonitorRequest {
	s.DetectTimes = &v
	return s
}

func (s *UpdateApplicationMonitorRequest) SetListenerId(v string) *UpdateApplicationMonitorRequest {
	s.ListenerId = &v
	return s
}

func (s *UpdateApplicationMonitorRequest) SetOptionsJson(v string) *UpdateApplicationMonitorRequest {
	s.OptionsJson = &v
	return s
}

func (s *UpdateApplicationMonitorRequest) SetRegionId(v string) *UpdateApplicationMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateApplicationMonitorRequest) SetSilenceTime(v int32) *UpdateApplicationMonitorRequest {
	s.SilenceTime = &v
	return s
}

func (s *UpdateApplicationMonitorRequest) SetTaskId(v string) *UpdateApplicationMonitorRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateApplicationMonitorRequest) SetTaskName(v string) *UpdateApplicationMonitorRequest {
	s.TaskName = &v
	return s
}

func (s *UpdateApplicationMonitorRequest) Validate() error {
	return dara.Validate(s)
}
