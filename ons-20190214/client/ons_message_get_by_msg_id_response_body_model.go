// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsMessageGetByMsgIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OnsMessageGetByMsgIdResponseBodyData) *OnsMessageGetByMsgIdResponseBody
	GetData() *OnsMessageGetByMsgIdResponseBodyData
	SetRequestId(v string) *OnsMessageGetByMsgIdResponseBody
	GetRequestId() *string
}

type OnsMessageGetByMsgIdResponseBody struct {
	// The data returned.
	Data *OnsMessageGetByMsgIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// A07E3902-B92E-44A6-B6C5-6AA111111****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsMessageGetByMsgIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsMessageGetByMsgIdResponseBody) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByMsgIdResponseBody) GetData() *OnsMessageGetByMsgIdResponseBodyData {
	return s.Data
}

func (s *OnsMessageGetByMsgIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsMessageGetByMsgIdResponseBody) SetData(v *OnsMessageGetByMsgIdResponseBodyData) *OnsMessageGetByMsgIdResponseBody {
	s.Data = v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBody) SetRequestId(v string) *OnsMessageGetByMsgIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsMessageGetByMsgIdResponseBodyData struct {
	// The cyclic redundancy check (CRC) value of the message body.
	//
	// example:
	//
	// 914112295
	BodyCRC *int32 `json:"BodyCRC,omitempty" xml:"BodyCRC,omitempty"`
	// The producer instance that generated the message.
	//
	// example:
	//
	// ``42.120.**.**``:64646
	BornHost *string `json:"BornHost,omitempty" xml:"BornHost,omitempty"`
	// The timestamp that indicates when the message was produced.
	//
	// example:
	//
	// 1570761026630
	BornTimestamp *int64 `json:"BornTimestamp,omitempty" xml:"BornTimestamp,omitempty"`
	// The ID of the instance
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the message.
	//
	// example:
	//
	// 1E0578FE110F18B4AAC235C0C8460BA2
	MsgId        *string                                           `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	PropertyList *OnsMessageGetByMsgIdResponseBodyDataPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Struct"`
	// The number of retries that were performed to send the message to consumers.
	//
	// example:
	//
	// 1
	ReconsumeTimes *int32 `json:"ReconsumeTimes,omitempty" xml:"ReconsumeTimes,omitempty"`
	// The ApsaraMQ for RocketMQ broker that stores the message.
	//
	// example:
	//
	// 11.220.***.***:10911
	StoreHost *string `json:"StoreHost,omitempty" xml:"StoreHost,omitempty"`
	// The size of the message.
	//
	// example:
	//
	// 407
	StoreSize *int32 `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	// The timestamp that indicates when the ApsaraMQ for RocketMQ broker stored the message.
	//
	// example:
	//
	// 1570761026708
	StoreTimestamp *int64 `json:"StoreTimestamp,omitempty" xml:"StoreTimestamp,omitempty"`
	// The topic to which the message belongs.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsMessageGetByMsgIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OnsMessageGetByMsgIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByMsgIdResponseBodyData) GetBodyCRC() *int32 {
	return s.BodyCRC
}

func (s *OnsMessageGetByMsgIdResponseBodyData) GetBornHost() *string {
	return s.BornHost
}

func (s *OnsMessageGetByMsgIdResponseBodyData) GetBornTimestamp() *int64 {
	return s.BornTimestamp
}

func (s *OnsMessageGetByMsgIdResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsMessageGetByMsgIdResponseBodyData) GetMsgId() *string {
	return s.MsgId
}

func (s *OnsMessageGetByMsgIdResponseBodyData) GetPropertyList() *OnsMessageGetByMsgIdResponseBodyDataPropertyList {
	return s.PropertyList
}

func (s *OnsMessageGetByMsgIdResponseBodyData) GetReconsumeTimes() *int32 {
	return s.ReconsumeTimes
}

func (s *OnsMessageGetByMsgIdResponseBodyData) GetStoreHost() *string {
	return s.StoreHost
}

func (s *OnsMessageGetByMsgIdResponseBodyData) GetStoreSize() *int32 {
	return s.StoreSize
}

func (s *OnsMessageGetByMsgIdResponseBodyData) GetStoreTimestamp() *int64 {
	return s.StoreTimestamp
}

func (s *OnsMessageGetByMsgIdResponseBodyData) GetTopic() *string {
	return s.Topic
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetBodyCRC(v int32) *OnsMessageGetByMsgIdResponseBodyData {
	s.BodyCRC = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetBornHost(v string) *OnsMessageGetByMsgIdResponseBodyData {
	s.BornHost = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetBornTimestamp(v int64) *OnsMessageGetByMsgIdResponseBodyData {
	s.BornTimestamp = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetInstanceId(v string) *OnsMessageGetByMsgIdResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetMsgId(v string) *OnsMessageGetByMsgIdResponseBodyData {
	s.MsgId = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetPropertyList(v *OnsMessageGetByMsgIdResponseBodyDataPropertyList) *OnsMessageGetByMsgIdResponseBodyData {
	s.PropertyList = v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetReconsumeTimes(v int32) *OnsMessageGetByMsgIdResponseBodyData {
	s.ReconsumeTimes = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetStoreHost(v string) *OnsMessageGetByMsgIdResponseBodyData {
	s.StoreHost = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetStoreSize(v int32) *OnsMessageGetByMsgIdResponseBodyData {
	s.StoreSize = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetStoreTimestamp(v int64) *OnsMessageGetByMsgIdResponseBodyData {
	s.StoreTimestamp = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) SetTopic(v string) *OnsMessageGetByMsgIdResponseBodyData {
	s.Topic = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyData) Validate() error {
	if s.PropertyList != nil {
		if err := s.PropertyList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsMessageGetByMsgIdResponseBodyDataPropertyList struct {
	MessageProperty []*OnsMessageGetByMsgIdResponseBodyDataPropertyListMessageProperty `json:"MessageProperty,omitempty" xml:"MessageProperty,omitempty" type:"Repeated"`
}

func (s OnsMessageGetByMsgIdResponseBodyDataPropertyList) String() string {
	return dara.Prettify(s)
}

func (s OnsMessageGetByMsgIdResponseBodyDataPropertyList) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByMsgIdResponseBodyDataPropertyList) GetMessageProperty() []*OnsMessageGetByMsgIdResponseBodyDataPropertyListMessageProperty {
	return s.MessageProperty
}

func (s *OnsMessageGetByMsgIdResponseBodyDataPropertyList) SetMessageProperty(v []*OnsMessageGetByMsgIdResponseBodyDataPropertyListMessageProperty) *OnsMessageGetByMsgIdResponseBodyDataPropertyList {
	s.MessageProperty = v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyDataPropertyList) Validate() error {
	if s.MessageProperty != nil {
		for _, item := range s.MessageProperty {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OnsMessageGetByMsgIdResponseBodyDataPropertyListMessageProperty struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsMessageGetByMsgIdResponseBodyDataPropertyListMessageProperty) String() string {
	return dara.Prettify(s)
}

func (s OnsMessageGetByMsgIdResponseBodyDataPropertyListMessageProperty) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByMsgIdResponseBodyDataPropertyListMessageProperty) GetName() *string {
	return s.Name
}

func (s *OnsMessageGetByMsgIdResponseBodyDataPropertyListMessageProperty) GetValue() *string {
	return s.Value
}

func (s *OnsMessageGetByMsgIdResponseBodyDataPropertyListMessageProperty) SetName(v string) *OnsMessageGetByMsgIdResponseBodyDataPropertyListMessageProperty {
	s.Name = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyDataPropertyListMessageProperty) SetValue(v string) *OnsMessageGetByMsgIdResponseBodyDataPropertyListMessageProperty {
	s.Value = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponseBodyDataPropertyListMessageProperty) Validate() error {
	return dara.Validate(s)
}
