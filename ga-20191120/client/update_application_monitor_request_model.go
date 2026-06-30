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
	// The URL or IP address to be probed.
	//
	// example:
	//
	// https://www.aliyun.com
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- value as the **ClientToken*	- value. The **RequestId*	- value of each API request is different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable the automatic diagnostics feature. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false*	- (default): Disabled.
	//
	// example:
	//
	// false
	DetectEnable *bool `json:"DetectEnable,omitempty" xml:"DetectEnable,omitempty"`
	// The threshold that triggers automatic diagnostics. When the origin availability rate falls below this threshold, automatic diagnostics is triggered.
	//
	// Valid values: **0*	- to **100**.
	//
	// example:
	//
	// 0
	DetectThreshold *int32 `json:"DetectThreshold,omitempty" xml:"DetectThreshold,omitempty"`
	// The number of consecutive times that the availability rate must fall below the threshold before automatic diagnostics is triggered.
	//
	// Valid values: **1*	- to **20**.
	//
	// example:
	//
	// 1
	DetectTimes *int32 `json:"DetectTimes,omitempty" xml:"DetectTimes,omitempty"`
	// The instance ID of the listener associated with the origin probing task that you want to modify.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The advanced extension options for the listener protocol type of the origin probing task. Different listener protocol types correspond to different extension options.
	//
	// example:
	//
	// { "http_method": "get","header": "key:asd","acceptable_response_code": "500","cert_verify": true }
	OptionsJson *string `json:"OptionsJson,omitempty" xml:"OptionsJson,omitempty"`
	// The region ID of the Alibaba Cloud Global Accelerator (GA) instance. Set the value to **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The silence period for automatic diagnostics. This specifies the interval between automatic diagnostics triggers when the availability rate does not recover to normal after diagnostics is triggered.
	//
	// When the availability rate falls below the automatic diagnostics threshold for consecutive times (as specified by **DetectTimes**), automatic diagnostics is triggered. If the availability rate remains below the threshold during the silence period, automatic diagnostics is not triggered again within the silence period. If the availability rate has not recovered after the silence period expires, automatic diagnostics is triggered again.
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
	// The name must be 1 to 128 characters in length and must start with a letter or a Chinese character. It can contain digits, periods (.), underscores (_), and hyphens (-).
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
