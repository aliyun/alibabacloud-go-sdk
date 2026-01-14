// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *CreateApplicationMonitorRequest
	GetAcceleratorId() *string
	SetAddress(v string) *CreateApplicationMonitorRequest
	GetAddress() *string
	SetClientToken(v string) *CreateApplicationMonitorRequest
	GetClientToken() *string
	SetDetectEnable(v bool) *CreateApplicationMonitorRequest
	GetDetectEnable() *bool
	SetDetectThreshold(v int32) *CreateApplicationMonitorRequest
	GetDetectThreshold() *int32
	SetDetectTimes(v int32) *CreateApplicationMonitorRequest
	GetDetectTimes() *int32
	SetListenerId(v string) *CreateApplicationMonitorRequest
	GetListenerId() *string
	SetOptionsJson(v string) *CreateApplicationMonitorRequest
	GetOptionsJson() *string
	SetRegionId(v string) *CreateApplicationMonitorRequest
	GetRegionId() *string
	SetSilenceTime(v int32) *CreateApplicationMonitorRequest
	GetSilenceTime() *int32
	SetTaskName(v string) *CreateApplicationMonitorRequest
	GetTaskName() *string
}

type CreateApplicationMonitorRequest struct {
	// The ID of the GA instance on which you want to perform the origin probing task.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The URL or IP address that you want to probe.
	//
	// This parameter is required.
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
	// The threshold that is used to trigger the automatic diagnostics feature. Unit: percentage.
	//
	// Valid values: **0*	- to **100**.
	//
	// The default value is **0**, which indicates that the automatic diagnostics feature is not triggered.
	//
	// example:
	//
	// 0
	DetectThreshold *int32 `json:"DetectThreshold,omitempty" xml:"DetectThreshold,omitempty"`
	// The number of times that the threshold must be reached before the automatic diagnostics feature is triggered.
	//
	// Valid values: **1*	- to **20**. Default value: **1**.
	//
	// example:
	//
	// 1
	DetectTimes *int32 `json:"DetectTimes,omitempty" xml:"DetectTimes,omitempty"`
	// The ID of the listener on which you want to perform the origin probing task.
	//
	// This parameter is required.
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
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The silence period of the automatic diagnostics feature. This parameter specifies the interval at which the automatic diagnostics feature is triggered. If the availability rate does not return to normal after GA triggers an automatic diagnostic task, GA must wait until the silence period ends before GA can trigger another automatic diagnostic task.
	//
	// If the number of consecutive times that the availability rate drops below the threshold of automatic diagnostics reaches the value of **DetectTimes**, the automatic diagnostics feature is triggered. The automatic diagnostics feature is not triggered again within the silence period even if the availability rate remains below the threshold. If the availability rate does not return to normal after the silence period ends, the automatic diagnostics feature is triggered again.
	//
	// Unit: seconds. Valid values: **300*	- to **86400**. Default value: **300**.
	//
	// example:
	//
	// 300
	SilenceTime *int32 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The name of the origin probing task.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// task1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateApplicationMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationMonitorRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationMonitorRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *CreateApplicationMonitorRequest) GetAddress() *string {
	return s.Address
}

func (s *CreateApplicationMonitorRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateApplicationMonitorRequest) GetDetectEnable() *bool {
	return s.DetectEnable
}

func (s *CreateApplicationMonitorRequest) GetDetectThreshold() *int32 {
	return s.DetectThreshold
}

func (s *CreateApplicationMonitorRequest) GetDetectTimes() *int32 {
	return s.DetectTimes
}

func (s *CreateApplicationMonitorRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *CreateApplicationMonitorRequest) GetOptionsJson() *string {
	return s.OptionsJson
}

func (s *CreateApplicationMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApplicationMonitorRequest) GetSilenceTime() *int32 {
	return s.SilenceTime
}

func (s *CreateApplicationMonitorRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateApplicationMonitorRequest) SetAcceleratorId(v string) *CreateApplicationMonitorRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateApplicationMonitorRequest) SetAddress(v string) *CreateApplicationMonitorRequest {
	s.Address = &v
	return s
}

func (s *CreateApplicationMonitorRequest) SetClientToken(v string) *CreateApplicationMonitorRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateApplicationMonitorRequest) SetDetectEnable(v bool) *CreateApplicationMonitorRequest {
	s.DetectEnable = &v
	return s
}

func (s *CreateApplicationMonitorRequest) SetDetectThreshold(v int32) *CreateApplicationMonitorRequest {
	s.DetectThreshold = &v
	return s
}

func (s *CreateApplicationMonitorRequest) SetDetectTimes(v int32) *CreateApplicationMonitorRequest {
	s.DetectTimes = &v
	return s
}

func (s *CreateApplicationMonitorRequest) SetListenerId(v string) *CreateApplicationMonitorRequest {
	s.ListenerId = &v
	return s
}

func (s *CreateApplicationMonitorRequest) SetOptionsJson(v string) *CreateApplicationMonitorRequest {
	s.OptionsJson = &v
	return s
}

func (s *CreateApplicationMonitorRequest) SetRegionId(v string) *CreateApplicationMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApplicationMonitorRequest) SetSilenceTime(v int32) *CreateApplicationMonitorRequest {
	s.SilenceTime = &v
	return s
}

func (s *CreateApplicationMonitorRequest) SetTaskName(v string) *CreateApplicationMonitorRequest {
	s.TaskName = &v
	return s
}

func (s *CreateApplicationMonitorRequest) Validate() error {
	return dara.Validate(s)
}
