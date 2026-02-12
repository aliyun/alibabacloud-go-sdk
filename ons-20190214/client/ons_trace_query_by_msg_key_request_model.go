// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTraceQueryByMsgKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *OnsTraceQueryByMsgKeyRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *OnsTraceQueryByMsgKeyRequest
	GetEndTime() *int64
	SetInstanceId(v string) *OnsTraceQueryByMsgKeyRequest
	GetInstanceId() *string
	SetMsgKey(v string) *OnsTraceQueryByMsgKeyRequest
	GetMsgKey() *string
	SetTopic(v string) *OnsTraceQueryByMsgKeyRequest
	GetTopic() *string
}

type OnsTraceQueryByMsgKeyRequest struct {
	// The start of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1570852800000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The end of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1570968000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the ApsaraMQ for RocketMQ instance that contains the specified topic.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The key of the message that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// ORDERID_100
	MsgKey *string `json:"MsgKey,omitempty" xml:"MsgKey,omitempty"`
	// The topic that contains the message you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsTraceQueryByMsgKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsTraceQueryByMsgKeyRequest) GoString() string {
	return s.String()
}

func (s *OnsTraceQueryByMsgKeyRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *OnsTraceQueryByMsgKeyRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *OnsTraceQueryByMsgKeyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsTraceQueryByMsgKeyRequest) GetMsgKey() *string {
	return s.MsgKey
}

func (s *OnsTraceQueryByMsgKeyRequest) GetTopic() *string {
	return s.Topic
}

func (s *OnsTraceQueryByMsgKeyRequest) SetBeginTime(v int64) *OnsTraceQueryByMsgKeyRequest {
	s.BeginTime = &v
	return s
}

func (s *OnsTraceQueryByMsgKeyRequest) SetEndTime(v int64) *OnsTraceQueryByMsgKeyRequest {
	s.EndTime = &v
	return s
}

func (s *OnsTraceQueryByMsgKeyRequest) SetInstanceId(v string) *OnsTraceQueryByMsgKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsTraceQueryByMsgKeyRequest) SetMsgKey(v string) *OnsTraceQueryByMsgKeyRequest {
	s.MsgKey = &v
	return s
}

func (s *OnsTraceQueryByMsgKeyRequest) SetTopic(v string) *OnsTraceQueryByMsgKeyRequest {
	s.Topic = &v
	return s
}

func (s *OnsTraceQueryByMsgKeyRequest) Validate() error {
	return dara.Validate(s)
}
