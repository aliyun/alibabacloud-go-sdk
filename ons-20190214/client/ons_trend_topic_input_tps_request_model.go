// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTrendTopicInputTpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *OnsTrendTopicInputTpsRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *OnsTrendTopicInputTpsRequest
	GetEndTime() *int64
	SetInstanceId(v string) *OnsTrendTopicInputTpsRequest
	GetInstanceId() *string
	SetPeriod(v int64) *OnsTrendTopicInputTpsRequest
	GetPeriod() *int64
	SetTopic(v string) *OnsTrendTopicInputTpsRequest
	GetTopic() *string
	SetType(v int32) *OnsTrendTopicInputTpsRequest
	GetType() *int32
}

type OnsTrendTopicInputTpsRequest struct {
	// The timestamp that indicates the beginning of the time range to query. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1570852800000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The timestamp that indicates the end of the time range to query. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1570868400000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the instance to which the topic you want to query belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The sampling period. Unit: minutes. Valid values: 1, 5, and 10.
	//
	// example:
	//
	// 10
	Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The name of the topic that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The type of information that you want to query. Valid values:
	//
	// 	- **0**: the number of messages that are published to the topic during each sampling period.
	//
	// 	- **1**: the TPS for message publishing in the topic during each sampling period.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s OnsTrendTopicInputTpsRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsTrendTopicInputTpsRequest) GoString() string {
	return s.String()
}

func (s *OnsTrendTopicInputTpsRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *OnsTrendTopicInputTpsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *OnsTrendTopicInputTpsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsTrendTopicInputTpsRequest) GetPeriod() *int64 {
	return s.Period
}

func (s *OnsTrendTopicInputTpsRequest) GetTopic() *string {
	return s.Topic
}

func (s *OnsTrendTopicInputTpsRequest) GetType() *int32 {
	return s.Type
}

func (s *OnsTrendTopicInputTpsRequest) SetBeginTime(v int64) *OnsTrendTopicInputTpsRequest {
	s.BeginTime = &v
	return s
}

func (s *OnsTrendTopicInputTpsRequest) SetEndTime(v int64) *OnsTrendTopicInputTpsRequest {
	s.EndTime = &v
	return s
}

func (s *OnsTrendTopicInputTpsRequest) SetInstanceId(v string) *OnsTrendTopicInputTpsRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsTrendTopicInputTpsRequest) SetPeriod(v int64) *OnsTrendTopicInputTpsRequest {
	s.Period = &v
	return s
}

func (s *OnsTrendTopicInputTpsRequest) SetTopic(v string) *OnsTrendTopicInputTpsRequest {
	s.Topic = &v
	return s
}

func (s *OnsTrendTopicInputTpsRequest) SetType(v int32) *OnsTrendTopicInputTpsRequest {
	s.Type = &v
	return s
}

func (s *OnsTrendTopicInputTpsRequest) Validate() error {
	return dara.Validate(s)
}
