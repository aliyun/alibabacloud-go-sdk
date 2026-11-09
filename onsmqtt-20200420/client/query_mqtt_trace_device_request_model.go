// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMqttTraceDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *QueryMqttTraceDeviceRequest
	GetBeginTime() *int64
	SetClientId(v string) *QueryMqttTraceDeviceRequest
	GetClientId() *string
	SetCurrentPage(v int32) *QueryMqttTraceDeviceRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *QueryMqttTraceDeviceRequest
	GetEndTime() *int64
	SetInstanceId(v string) *QueryMqttTraceDeviceRequest
	GetInstanceId() *string
	SetMqttRegionId(v string) *QueryMqttTraceDeviceRequest
	GetMqttRegionId() *string
	SetPageSize(v int32) *QueryMqttTraceDeviceRequest
	GetPageSize() *int32
	SetReverse(v bool) *QueryMqttTraceDeviceRequest
	GetReverse() *bool
}

type QueryMqttTraceDeviceRequest struct {
	// The beginning of the time range to query. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1621580400000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The client ID of the device whose trace information you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test@@@producer
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The number of the page to return. Pages start from page 1. If the value of this parameter exceeds the total number of pages, an empty result is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1621584000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance. This value must match the instance ID that the client actually uses. You can view the instance ID in the **Basic Information*	- area on the **Instance Details*	- page in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// mqtt-cn-i7m26mf****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the ApsaraMQ for MQTT instance. For more information, see [Endpoints](https://help.aliyun.com/document_detail/181438.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	MqttRegionId *string `json:"MqttRegionId,omitempty" xml:"MqttRegionId,omitempty"`
	// The maximum number of MSMQ to display on each page for paging. Minimum value: 1. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to display the results in reverse chronological order. Valid values:
	//
	// - **true**: Results are displayed in reverse chronological order of device activity time. The most recent device activity is displayed first, and the earliest device activity is displayed last.
	//
	// - **false**: Results are displayed in chronological order of device activity time. The earliest device activity is displayed first, and the most recent device activity is displayed last.
	//
	// If this parameter is not specified, results are displayed in chronological order by default.
	//
	// example:
	//
	// false
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
}

func (s QueryMqttTraceDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMqttTraceDeviceRequest) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceDeviceRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *QueryMqttTraceDeviceRequest) GetClientId() *string {
	return s.ClientId
}

func (s *QueryMqttTraceDeviceRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryMqttTraceDeviceRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryMqttTraceDeviceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryMqttTraceDeviceRequest) GetMqttRegionId() *string {
	return s.MqttRegionId
}

func (s *QueryMqttTraceDeviceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMqttTraceDeviceRequest) GetReverse() *bool {
	return s.Reverse
}

func (s *QueryMqttTraceDeviceRequest) SetBeginTime(v int64) *QueryMqttTraceDeviceRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryMqttTraceDeviceRequest) SetClientId(v string) *QueryMqttTraceDeviceRequest {
	s.ClientId = &v
	return s
}

func (s *QueryMqttTraceDeviceRequest) SetCurrentPage(v int32) *QueryMqttTraceDeviceRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryMqttTraceDeviceRequest) SetEndTime(v int64) *QueryMqttTraceDeviceRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMqttTraceDeviceRequest) SetInstanceId(v string) *QueryMqttTraceDeviceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMqttTraceDeviceRequest) SetMqttRegionId(v string) *QueryMqttTraceDeviceRequest {
	s.MqttRegionId = &v
	return s
}

func (s *QueryMqttTraceDeviceRequest) SetPageSize(v int32) *QueryMqttTraceDeviceRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMqttTraceDeviceRequest) SetReverse(v bool) *QueryMqttTraceDeviceRequest {
	s.Reverse = &v
	return s
}

func (s *QueryMqttTraceDeviceRequest) Validate() error {
	return dara.Validate(s)
}
