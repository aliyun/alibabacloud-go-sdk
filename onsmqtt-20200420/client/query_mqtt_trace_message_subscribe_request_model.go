// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMqttTraceMessageSubscribeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *QueryMqttTraceMessageSubscribeRequest
	GetBeginTime() *int64
	SetClientId(v string) *QueryMqttTraceMessageSubscribeRequest
	GetClientId() *string
	SetCurrentPage(v int32) *QueryMqttTraceMessageSubscribeRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *QueryMqttTraceMessageSubscribeRequest
	GetEndTime() *int64
	SetInstanceId(v string) *QueryMqttTraceMessageSubscribeRequest
	GetInstanceId() *string
	SetMqttRegionId(v string) *QueryMqttTraceMessageSubscribeRequest
	GetMqttRegionId() *string
	SetMsgId(v string) *QueryMqttTraceMessageSubscribeRequest
	GetMsgId() *string
	SetPageSize(v int32) *QueryMqttTraceMessageSubscribeRequest
	GetPageSize() *int32
	SetReverse(v bool) *QueryMqttTraceMessageSubscribeRequest
	GetReverse() *bool
}

type QueryMqttTraceMessageSubscribeRequest struct {
	// The beginning of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1621936800000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The ID of the client that subscribes to the message. If you do not specify this parameter, the IDs of all clients that subscribe to the message are returned.
	//
	// example:
	//
	// GID_test@@@consumer
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
	// 1618646400000
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
	// The message ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// AC1EC1B33D5978308DB17F3245E4****
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
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
	// - **true**: The returned results are displayed in reverse order of the time when messages are delivered. This means that the latest consumed message is displayed as the first entry and the earliest consumed message is displayed as the last entry.
	//
	// - **false**: The returned results are displayed in order of the time when messages are delivered. This means that the earliest consumed message is displayed as the first entry and the latest consumed message is displayed as the last entry.
	//
	// If you do not specify this parameter, the returned results are displayed in order of time when messages are delivered.
	//
	// example:
	//
	// false
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
}

func (s QueryMqttTraceMessageSubscribeRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMqttTraceMessageSubscribeRequest) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessageSubscribeRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *QueryMqttTraceMessageSubscribeRequest) GetClientId() *string {
	return s.ClientId
}

func (s *QueryMqttTraceMessageSubscribeRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryMqttTraceMessageSubscribeRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryMqttTraceMessageSubscribeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryMqttTraceMessageSubscribeRequest) GetMqttRegionId() *string {
	return s.MqttRegionId
}

func (s *QueryMqttTraceMessageSubscribeRequest) GetMsgId() *string {
	return s.MsgId
}

func (s *QueryMqttTraceMessageSubscribeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMqttTraceMessageSubscribeRequest) GetReverse() *bool {
	return s.Reverse
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetBeginTime(v int64) *QueryMqttTraceMessageSubscribeRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetClientId(v string) *QueryMqttTraceMessageSubscribeRequest {
	s.ClientId = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetCurrentPage(v int32) *QueryMqttTraceMessageSubscribeRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetEndTime(v int64) *QueryMqttTraceMessageSubscribeRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetInstanceId(v string) *QueryMqttTraceMessageSubscribeRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetMqttRegionId(v string) *QueryMqttTraceMessageSubscribeRequest {
	s.MqttRegionId = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetMsgId(v string) *QueryMqttTraceMessageSubscribeRequest {
	s.MsgId = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetPageSize(v int32) *QueryMqttTraceMessageSubscribeRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetReverse(v bool) *QueryMqttTraceMessageSubscribeRequest {
	s.Reverse = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeRequest) Validate() error {
	return dara.Validate(s)
}
