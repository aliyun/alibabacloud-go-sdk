// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTrendGroupOutputTpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *OnsTrendGroupOutputTpsRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *OnsTrendGroupOutputTpsRequest
	GetEndTime() *int64
	SetGroupId(v string) *OnsTrendGroupOutputTpsRequest
	GetGroupId() *string
	SetInstanceId(v string) *OnsTrendGroupOutputTpsRequest
	GetInstanceId() *string
	SetPeriod(v int64) *OnsTrendGroupOutputTpsRequest
	GetPeriod() *int64
	SetTopic(v string) *OnsTrendGroupOutputTpsRequest
	GetTopic() *string
	SetType(v int32) *OnsTrendGroupOutputTpsRequest
	GetType() *int32
}

type OnsTrendGroupOutputTpsRequest struct {
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
	// The ID of the consumer group that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the instance to which the consumer group you want to query belongs.
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
	// 	- **0**: the number of messages that are consumed during each sampling period.
	//
	// 	- **1**: the TPS for message consumption during each sampling period.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s OnsTrendGroupOutputTpsRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsTrendGroupOutputTpsRequest) GoString() string {
	return s.String()
}

func (s *OnsTrendGroupOutputTpsRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *OnsTrendGroupOutputTpsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *OnsTrendGroupOutputTpsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *OnsTrendGroupOutputTpsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsTrendGroupOutputTpsRequest) GetPeriod() *int64 {
	return s.Period
}

func (s *OnsTrendGroupOutputTpsRequest) GetTopic() *string {
	return s.Topic
}

func (s *OnsTrendGroupOutputTpsRequest) GetType() *int32 {
	return s.Type
}

func (s *OnsTrendGroupOutputTpsRequest) SetBeginTime(v int64) *OnsTrendGroupOutputTpsRequest {
	s.BeginTime = &v
	return s
}

func (s *OnsTrendGroupOutputTpsRequest) SetEndTime(v int64) *OnsTrendGroupOutputTpsRequest {
	s.EndTime = &v
	return s
}

func (s *OnsTrendGroupOutputTpsRequest) SetGroupId(v string) *OnsTrendGroupOutputTpsRequest {
	s.GroupId = &v
	return s
}

func (s *OnsTrendGroupOutputTpsRequest) SetInstanceId(v string) *OnsTrendGroupOutputTpsRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsTrendGroupOutputTpsRequest) SetPeriod(v int64) *OnsTrendGroupOutputTpsRequest {
	s.Period = &v
	return s
}

func (s *OnsTrendGroupOutputTpsRequest) SetTopic(v string) *OnsTrendGroupOutputTpsRequest {
	s.Topic = &v
	return s
}

func (s *OnsTrendGroupOutputTpsRequest) SetType(v int32) *OnsTrendGroupOutputTpsRequest {
	s.Type = &v
	return s
}

func (s *OnsTrendGroupOutputTpsRequest) Validate() error {
	return dara.Validate(s)
}
