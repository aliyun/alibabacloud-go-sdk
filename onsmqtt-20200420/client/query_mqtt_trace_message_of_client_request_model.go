// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMqttTraceMessageOfClientRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *QueryMqttTraceMessageOfClientRequest
	GetBeginTime() *int64
	SetClientId(v string) *QueryMqttTraceMessageOfClientRequest
	GetClientId() *string
	SetCurrentPage(v int32) *QueryMqttTraceMessageOfClientRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *QueryMqttTraceMessageOfClientRequest
	GetEndTime() *int64
	SetInstanceId(v string) *QueryMqttTraceMessageOfClientRequest
	GetInstanceId() *string
	SetMqttRegionId(v string) *QueryMqttTraceMessageOfClientRequest
	GetMqttRegionId() *string
	SetPageSize(v int32) *QueryMqttTraceMessageOfClientRequest
	GetPageSize() *int32
	SetReverse(v bool) *QueryMqttTraceMessageOfClientRequest
	GetReverse() *bool
}

type QueryMqttTraceMessageOfClientRequest struct {
	// The beginning of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1618646400000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The client ID of the device whose messages you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test@@@producer
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The number of the page to return. Pages start from page 1. If the input parameter value is greater than the total number of pages, the returned result is empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1621591200000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance. The ID must be consistent with the ID of the instance that the ApsaraMQ for MQTT client uses. You can view the instance ID in the **Basic Information*	- section of the **Instance Details*	- page that corresponds to the instance in the ApsaraMQ for MQTT console.
	//
	// This parameter is required.
	//
	// example:
	//
	// mqtt-cn-i7m26mf****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the ApsaraMQ for MQTT instance resides. For more information, see [Endpoints](https://help.aliyun.com/document_detail/181438.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	MqttRegionId *string `json:"MqttRegionId,omitempty" xml:"MqttRegionId,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether the returned results are displayed in reverse chronological order. Valid values:
	//
	// 	- **true**: The returned results are displayed in reverse order of the time when messages are sent or received. This means that the latest sent or received message is displayed as the first entry and the earliest sent or received message is displayed as the last entry.
	//
	// 	- **false**: The returned results are displayed in order of the time when messages are sent or received. This means that the earliest sent or received message is displayed as the first entry and the latest sent or received message is displayed as the last entry.
	//
	// If this parameter is not specified, the returned results are displayed in order of the time when messages are sent or received.
	//
	// example:
	//
	// false
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
}

func (s QueryMqttTraceMessageOfClientRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMqttTraceMessageOfClientRequest) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessageOfClientRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *QueryMqttTraceMessageOfClientRequest) GetClientId() *string {
	return s.ClientId
}

func (s *QueryMqttTraceMessageOfClientRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryMqttTraceMessageOfClientRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryMqttTraceMessageOfClientRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryMqttTraceMessageOfClientRequest) GetMqttRegionId() *string {
	return s.MqttRegionId
}

func (s *QueryMqttTraceMessageOfClientRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMqttTraceMessageOfClientRequest) GetReverse() *bool {
	return s.Reverse
}

func (s *QueryMqttTraceMessageOfClientRequest) SetBeginTime(v int64) *QueryMqttTraceMessageOfClientRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientRequest) SetClientId(v string) *QueryMqttTraceMessageOfClientRequest {
	s.ClientId = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientRequest) SetCurrentPage(v int32) *QueryMqttTraceMessageOfClientRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientRequest) SetEndTime(v int64) *QueryMqttTraceMessageOfClientRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientRequest) SetInstanceId(v string) *QueryMqttTraceMessageOfClientRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientRequest) SetMqttRegionId(v string) *QueryMqttTraceMessageOfClientRequest {
	s.MqttRegionId = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientRequest) SetPageSize(v int32) *QueryMqttTraceMessageOfClientRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientRequest) SetReverse(v bool) *QueryMqttTraceMessageOfClientRequest {
	s.Reverse = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientRequest) Validate() error {
	return dara.Validate(s)
}
