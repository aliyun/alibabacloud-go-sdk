// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsConsumerAccumulateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OnsConsumerAccumulateResponseBodyData) *OnsConsumerAccumulateResponseBody
	GetData() *OnsConsumerAccumulateResponseBodyData
	SetRequestId(v string) *OnsConsumerAccumulateResponseBody
	GetRequestId() *string
}

type OnsConsumerAccumulateResponseBody struct {
	// The message accumulation information about topics to which the specified consumer subscribes.
	Data *OnsConsumerAccumulateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// CE817BFF-B389-43CD-9419-95011AC9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsConsumerAccumulateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerAccumulateResponseBody) GoString() string {
	return s.String()
}

func (s *OnsConsumerAccumulateResponseBody) GetData() *OnsConsumerAccumulateResponseBodyData {
	return s.Data
}

func (s *OnsConsumerAccumulateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsConsumerAccumulateResponseBody) SetData(v *OnsConsumerAccumulateResponseBodyData) *OnsConsumerAccumulateResponseBody {
	s.Data = v
	return s
}

func (s *OnsConsumerAccumulateResponseBody) SetRequestId(v string) *OnsConsumerAccumulateResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsConsumerAccumulateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsConsumerAccumulateResponseBodyData struct {
	// The transactions per second (TPS) for message consumption performed by consumers in the group.
	//
	// example:
	//
	// 10
	ConsumeTps *float32 `json:"ConsumeTps,omitempty" xml:"ConsumeTps,omitempty"`
	// The consumption latency.
	//
	// example:
	//
	// 10000
	DelayTime         *int64                                                  `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	DetailInTopicList *OnsConsumerAccumulateResponseBodyDataDetailInTopicList `json:"DetailInTopicList,omitempty" xml:"DetailInTopicList,omitempty" type:"Struct"`
	// The point in time when the latest message consumed by a consumer in the consumer group was produced.
	//
	// example:
	//
	// 1566231000000
	LastTimestamp *int64 `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// Indicates whether the consumer group is online. The consumer group is online if one of the consumers in the group is online. Valid values:
	//
	// 	- **true**: The consumer group is online.
	//
	// 	- **false**: The consumer group is offline.
	//
	// example:
	//
	// true
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// The total number of accumulated messages in all topics to which the consumer group subscribes.
	//
	// example:
	//
	// 100
	TotalDiff *int64 `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
}

func (s OnsConsumerAccumulateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerAccumulateResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsConsumerAccumulateResponseBodyData) GetConsumeTps() *float32 {
	return s.ConsumeTps
}

func (s *OnsConsumerAccumulateResponseBodyData) GetDelayTime() *int64 {
	return s.DelayTime
}

func (s *OnsConsumerAccumulateResponseBodyData) GetDetailInTopicList() *OnsConsumerAccumulateResponseBodyDataDetailInTopicList {
	return s.DetailInTopicList
}

func (s *OnsConsumerAccumulateResponseBodyData) GetLastTimestamp() *int64 {
	return s.LastTimestamp
}

func (s *OnsConsumerAccumulateResponseBodyData) GetOnline() *bool {
	return s.Online
}

func (s *OnsConsumerAccumulateResponseBodyData) GetTotalDiff() *int64 {
	return s.TotalDiff
}

func (s *OnsConsumerAccumulateResponseBodyData) SetConsumeTps(v float32) *OnsConsumerAccumulateResponseBodyData {
	s.ConsumeTps = &v
	return s
}

func (s *OnsConsumerAccumulateResponseBodyData) SetDelayTime(v int64) *OnsConsumerAccumulateResponseBodyData {
	s.DelayTime = &v
	return s
}

func (s *OnsConsumerAccumulateResponseBodyData) SetDetailInTopicList(v *OnsConsumerAccumulateResponseBodyDataDetailInTopicList) *OnsConsumerAccumulateResponseBodyData {
	s.DetailInTopicList = v
	return s
}

func (s *OnsConsumerAccumulateResponseBodyData) SetLastTimestamp(v int64) *OnsConsumerAccumulateResponseBodyData {
	s.LastTimestamp = &v
	return s
}

func (s *OnsConsumerAccumulateResponseBodyData) SetOnline(v bool) *OnsConsumerAccumulateResponseBodyData {
	s.Online = &v
	return s
}

func (s *OnsConsumerAccumulateResponseBodyData) SetTotalDiff(v int64) *OnsConsumerAccumulateResponseBodyData {
	s.TotalDiff = &v
	return s
}

func (s *OnsConsumerAccumulateResponseBodyData) Validate() error {
	if s.DetailInTopicList != nil {
		if err := s.DetailInTopicList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsConsumerAccumulateResponseBodyDataDetailInTopicList struct {
	DetailInTopicDo []*OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo `json:"DetailInTopicDo,omitempty" xml:"DetailInTopicDo,omitempty" type:"Repeated"`
}

func (s OnsConsumerAccumulateResponseBodyDataDetailInTopicList) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerAccumulateResponseBodyDataDetailInTopicList) GoString() string {
	return s.String()
}

func (s *OnsConsumerAccumulateResponseBodyDataDetailInTopicList) GetDetailInTopicDo() []*OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo {
	return s.DetailInTopicDo
}

func (s *OnsConsumerAccumulateResponseBodyDataDetailInTopicList) SetDetailInTopicDo(v []*OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo) *OnsConsumerAccumulateResponseBodyDataDetailInTopicList {
	s.DetailInTopicDo = v
	return s
}

func (s *OnsConsumerAccumulateResponseBodyDataDetailInTopicList) Validate() error {
	if s.DetailInTopicDo != nil {
		for _, item := range s.DetailInTopicDo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo struct {
	DelayTime     *int64  `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	LastTimestamp *int64  `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	Topic         *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	TotalDiff     *int64  `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
}

func (s OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo) GoString() string {
	return s.String()
}

func (s *OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo) GetDelayTime() *int64 {
	return s.DelayTime
}

func (s *OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo) GetLastTimestamp() *int64 {
	return s.LastTimestamp
}

func (s *OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo) GetTopic() *string {
	return s.Topic
}

func (s *OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo) GetTotalDiff() *int64 {
	return s.TotalDiff
}

func (s *OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo) SetDelayTime(v int64) *OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo {
	s.DelayTime = &v
	return s
}

func (s *OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo) SetLastTimestamp(v int64) *OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo {
	s.LastTimestamp = &v
	return s
}

func (s *OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo) SetTopic(v string) *OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo {
	s.Topic = &v
	return s
}

func (s *OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo) SetTotalDiff(v int64) *OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo {
	s.TotalDiff = &v
	return s
}

func (s *OnsConsumerAccumulateResponseBodyDataDetailInTopicListDetailInTopicDo) Validate() error {
	return dara.Validate(s)
}
