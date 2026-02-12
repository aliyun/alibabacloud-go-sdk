// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsConsumerTimeSpanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OnsConsumerTimeSpanResponseBodyData) *OnsConsumerTimeSpanResponseBody
	GetData() *OnsConsumerTimeSpanResponseBodyData
	SetRequestId(v string) *OnsConsumerTimeSpanResponseBody
	GetRequestId() *string
}

type OnsConsumerTimeSpanResponseBody struct {
	// The data returned.
	Data *OnsConsumerTimeSpanResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// A07E3902-B92E-44A6-B6C5-6AA111111****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsConsumerTimeSpanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerTimeSpanResponseBody) GoString() string {
	return s.String()
}

func (s *OnsConsumerTimeSpanResponseBody) GetData() *OnsConsumerTimeSpanResponseBodyData {
	return s.Data
}

func (s *OnsConsumerTimeSpanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsConsumerTimeSpanResponseBody) SetData(v *OnsConsumerTimeSpanResponseBodyData) *OnsConsumerTimeSpanResponseBody {
	s.Data = v
	return s
}

func (s *OnsConsumerTimeSpanResponseBody) SetRequestId(v string) *OnsConsumerTimeSpanResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsConsumerTimeSpanResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsConsumerTimeSpanResponseBodyData struct {
	// The most recent point in time when a message in the topic was consumed by the customer group.
	//
	// example:
	//
	// 1570761026400
	ConsumeTimeStamp *int64 `json:"ConsumeTimeStamp,omitempty" xml:"ConsumeTimeStamp,omitempty"`
	// The ID of the instance to which the consumer group belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The point in time when the most recently stored message was published to the topic.
	//
	// example:
	//
	// 1570761026804
	MaxTimeStamp *int64 `json:"MaxTimeStamp,omitempty" xml:"MaxTimeStamp,omitempty"`
	// The point in time when the earliest stored message was published to the topic.
	//
	// example:
	//
	// 1570701231122
	MinTimeStamp *int64 `json:"MinTimeStamp,omitempty" xml:"MinTimeStamp,omitempty"`
	// The name of the topic that you want to query.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsConsumerTimeSpanResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerTimeSpanResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsConsumerTimeSpanResponseBodyData) GetConsumeTimeStamp() *int64 {
	return s.ConsumeTimeStamp
}

func (s *OnsConsumerTimeSpanResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsConsumerTimeSpanResponseBodyData) GetMaxTimeStamp() *int64 {
	return s.MaxTimeStamp
}

func (s *OnsConsumerTimeSpanResponseBodyData) GetMinTimeStamp() *int64 {
	return s.MinTimeStamp
}

func (s *OnsConsumerTimeSpanResponseBodyData) GetTopic() *string {
	return s.Topic
}

func (s *OnsConsumerTimeSpanResponseBodyData) SetConsumeTimeStamp(v int64) *OnsConsumerTimeSpanResponseBodyData {
	s.ConsumeTimeStamp = &v
	return s
}

func (s *OnsConsumerTimeSpanResponseBodyData) SetInstanceId(v string) *OnsConsumerTimeSpanResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *OnsConsumerTimeSpanResponseBodyData) SetMaxTimeStamp(v int64) *OnsConsumerTimeSpanResponseBodyData {
	s.MaxTimeStamp = &v
	return s
}

func (s *OnsConsumerTimeSpanResponseBodyData) SetMinTimeStamp(v int64) *OnsConsumerTimeSpanResponseBodyData {
	s.MinTimeStamp = &v
	return s
}

func (s *OnsConsumerTimeSpanResponseBodyData) SetTopic(v string) *OnsConsumerTimeSpanResponseBodyData {
	s.Topic = &v
	return s
}

func (s *OnsConsumerTimeSpanResponseBodyData) Validate() error {
	return dara.Validate(s)
}
