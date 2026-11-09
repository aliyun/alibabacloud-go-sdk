// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMqttTraceMessagePublishRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *QueryMqttTraceMessagePublishRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *QueryMqttTraceMessagePublishRequest
	GetEndTime() *int64
	SetInstanceId(v string) *QueryMqttTraceMessagePublishRequest
	GetInstanceId() *string
	SetMqttRegionId(v string) *QueryMqttTraceMessagePublishRequest
	GetMqttRegionId() *string
	SetMsgId(v string) *QueryMqttTraceMessagePublishRequest
	GetMsgId() *string
}

type QueryMqttTraceMessagePublishRequest struct {
	// The beginning of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1618646400000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The end of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1621591200000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance. The ID must be consistent with the ID of the instance that the ApsaraMQ for MQTT client uses. You can view the instance ID in the **Basic Information*	- section on the **Instance Details*	- page that corresponds to the instance in the ApsaraMQ for MQTT console.
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
	// AC1EC0030EAB78308DB16A3EC773****
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s QueryMqttTraceMessagePublishRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMqttTraceMessagePublishRequest) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessagePublishRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *QueryMqttTraceMessagePublishRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryMqttTraceMessagePublishRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryMqttTraceMessagePublishRequest) GetMqttRegionId() *string {
	return s.MqttRegionId
}

func (s *QueryMqttTraceMessagePublishRequest) GetMsgId() *string {
	return s.MsgId
}

func (s *QueryMqttTraceMessagePublishRequest) SetBeginTime(v int64) *QueryMqttTraceMessagePublishRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryMqttTraceMessagePublishRequest) SetEndTime(v int64) *QueryMqttTraceMessagePublishRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMqttTraceMessagePublishRequest) SetInstanceId(v string) *QueryMqttTraceMessagePublishRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMqttTraceMessagePublishRequest) SetMqttRegionId(v string) *QueryMqttTraceMessagePublishRequest {
	s.MqttRegionId = &v
	return s
}

func (s *QueryMqttTraceMessagePublishRequest) SetMsgId(v string) *QueryMqttTraceMessagePublishRequest {
	s.MsgId = &v
	return s
}

func (s *QueryMqttTraceMessagePublishRequest) Validate() error {
	return dara.Validate(s)
}
