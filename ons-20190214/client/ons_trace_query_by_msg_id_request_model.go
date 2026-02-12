// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTraceQueryByMsgIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *OnsTraceQueryByMsgIdRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *OnsTraceQueryByMsgIdRequest
	GetEndTime() *int64
	SetInstanceId(v string) *OnsTraceQueryByMsgIdRequest
	GetInstanceId() *string
	SetMsgId(v string) *OnsTraceQueryByMsgIdRequest
	GetMsgId() *string
	SetTopic(v string) *OnsTraceQueryByMsgIdRequest
	GetTopic() *string
}

type OnsTraceQueryByMsgIdRequest struct {
	// The beginning of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds.
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
	// The ID of the message that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1E05791C117818B4AAC23B1BB0CE****
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The topic that contains the message you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsTraceQueryByMsgIdRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsTraceQueryByMsgIdRequest) GoString() string {
	return s.String()
}

func (s *OnsTraceQueryByMsgIdRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *OnsTraceQueryByMsgIdRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *OnsTraceQueryByMsgIdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsTraceQueryByMsgIdRequest) GetMsgId() *string {
	return s.MsgId
}

func (s *OnsTraceQueryByMsgIdRequest) GetTopic() *string {
	return s.Topic
}

func (s *OnsTraceQueryByMsgIdRequest) SetBeginTime(v int64) *OnsTraceQueryByMsgIdRequest {
	s.BeginTime = &v
	return s
}

func (s *OnsTraceQueryByMsgIdRequest) SetEndTime(v int64) *OnsTraceQueryByMsgIdRequest {
	s.EndTime = &v
	return s
}

func (s *OnsTraceQueryByMsgIdRequest) SetInstanceId(v string) *OnsTraceQueryByMsgIdRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsTraceQueryByMsgIdRequest) SetMsgId(v string) *OnsTraceQueryByMsgIdRequest {
	s.MsgId = &v
	return s
}

func (s *OnsTraceQueryByMsgIdRequest) SetTopic(v string) *OnsTraceQueryByMsgIdRequest {
	s.Topic = &v
	return s
}

func (s *OnsTraceQueryByMsgIdRequest) Validate() error {
	return dara.Validate(s)
}
