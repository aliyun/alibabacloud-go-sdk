// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsDLQMessageGetByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OnsDLQMessageGetByIdResponseBodyData) *OnsDLQMessageGetByIdResponseBody
	GetData() *OnsDLQMessageGetByIdResponseBodyData
	SetRequestId(v string) *OnsDLQMessageGetByIdResponseBody
	GetRequestId() *string
}

type OnsDLQMessageGetByIdResponseBody struct {
	// The data returned.
	Data *OnsDLQMessageGetByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID.
	//
	// example:
	//
	// A07E3902-B92E-44A6-B6C5-6AA111111****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsDLQMessageGetByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsDLQMessageGetByIdResponseBody) GoString() string {
	return s.String()
}

func (s *OnsDLQMessageGetByIdResponseBody) GetData() *OnsDLQMessageGetByIdResponseBodyData {
	return s.Data
}

func (s *OnsDLQMessageGetByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsDLQMessageGetByIdResponseBody) SetData(v *OnsDLQMessageGetByIdResponseBodyData) *OnsDLQMessageGetByIdResponseBody {
	s.Data = v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBody) SetRequestId(v string) *OnsDLQMessageGetByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsDLQMessageGetByIdResponseBodyData struct {
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
	// The timestamp that indicates the point in time when the message was generated. Unit: milliseconds.
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
	// The ID of the dead-letter message.
	//
	// example:
	//
	// 0BC16699165C03B925DB8A404E2D****
	MsgId        *string                                           `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	PropertyList *OnsDLQMessageGetByIdResponseBodyDataPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Struct"`
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
	// The size of the message. Unit: KB.
	//
	// example:
	//
	// 407
	StoreSize *int32 `json:"StoreSize,omitempty" xml:"StoreSize,omitempty"`
	// The timestamp that indicates the point in time when the ApsaraMQ for RocketMQ broker stored the message. Unit: milliseconds.
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

func (s OnsDLQMessageGetByIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OnsDLQMessageGetByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsDLQMessageGetByIdResponseBodyData) GetBodyCRC() *int32 {
	return s.BodyCRC
}

func (s *OnsDLQMessageGetByIdResponseBodyData) GetBornHost() *string {
	return s.BornHost
}

func (s *OnsDLQMessageGetByIdResponseBodyData) GetBornTimestamp() *int64 {
	return s.BornTimestamp
}

func (s *OnsDLQMessageGetByIdResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsDLQMessageGetByIdResponseBodyData) GetMsgId() *string {
	return s.MsgId
}

func (s *OnsDLQMessageGetByIdResponseBodyData) GetPropertyList() *OnsDLQMessageGetByIdResponseBodyDataPropertyList {
	return s.PropertyList
}

func (s *OnsDLQMessageGetByIdResponseBodyData) GetReconsumeTimes() *int32 {
	return s.ReconsumeTimes
}

func (s *OnsDLQMessageGetByIdResponseBodyData) GetStoreHost() *string {
	return s.StoreHost
}

func (s *OnsDLQMessageGetByIdResponseBodyData) GetStoreSize() *int32 {
	return s.StoreSize
}

func (s *OnsDLQMessageGetByIdResponseBodyData) GetStoreTimestamp() *int64 {
	return s.StoreTimestamp
}

func (s *OnsDLQMessageGetByIdResponseBodyData) GetTopic() *string {
	return s.Topic
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetBodyCRC(v int32) *OnsDLQMessageGetByIdResponseBodyData {
	s.BodyCRC = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetBornHost(v string) *OnsDLQMessageGetByIdResponseBodyData {
	s.BornHost = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetBornTimestamp(v int64) *OnsDLQMessageGetByIdResponseBodyData {
	s.BornTimestamp = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetInstanceId(v string) *OnsDLQMessageGetByIdResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetMsgId(v string) *OnsDLQMessageGetByIdResponseBodyData {
	s.MsgId = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetPropertyList(v *OnsDLQMessageGetByIdResponseBodyDataPropertyList) *OnsDLQMessageGetByIdResponseBodyData {
	s.PropertyList = v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetReconsumeTimes(v int32) *OnsDLQMessageGetByIdResponseBodyData {
	s.ReconsumeTimes = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetStoreHost(v string) *OnsDLQMessageGetByIdResponseBodyData {
	s.StoreHost = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetStoreSize(v int32) *OnsDLQMessageGetByIdResponseBodyData {
	s.StoreSize = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetStoreTimestamp(v int64) *OnsDLQMessageGetByIdResponseBodyData {
	s.StoreTimestamp = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) SetTopic(v string) *OnsDLQMessageGetByIdResponseBodyData {
	s.Topic = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyData) Validate() error {
	if s.PropertyList != nil {
		if err := s.PropertyList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsDLQMessageGetByIdResponseBodyDataPropertyList struct {
	MessageProperty []*OnsDLQMessageGetByIdResponseBodyDataPropertyListMessageProperty `json:"MessageProperty,omitempty" xml:"MessageProperty,omitempty" type:"Repeated"`
}

func (s OnsDLQMessageGetByIdResponseBodyDataPropertyList) String() string {
	return dara.Prettify(s)
}

func (s OnsDLQMessageGetByIdResponseBodyDataPropertyList) GoString() string {
	return s.String()
}

func (s *OnsDLQMessageGetByIdResponseBodyDataPropertyList) GetMessageProperty() []*OnsDLQMessageGetByIdResponseBodyDataPropertyListMessageProperty {
	return s.MessageProperty
}

func (s *OnsDLQMessageGetByIdResponseBodyDataPropertyList) SetMessageProperty(v []*OnsDLQMessageGetByIdResponseBodyDataPropertyListMessageProperty) *OnsDLQMessageGetByIdResponseBodyDataPropertyList {
	s.MessageProperty = v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyDataPropertyList) Validate() error {
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

type OnsDLQMessageGetByIdResponseBodyDataPropertyListMessageProperty struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsDLQMessageGetByIdResponseBodyDataPropertyListMessageProperty) String() string {
	return dara.Prettify(s)
}

func (s OnsDLQMessageGetByIdResponseBodyDataPropertyListMessageProperty) GoString() string {
	return s.String()
}

func (s *OnsDLQMessageGetByIdResponseBodyDataPropertyListMessageProperty) GetName() *string {
	return s.Name
}

func (s *OnsDLQMessageGetByIdResponseBodyDataPropertyListMessageProperty) GetValue() *string {
	return s.Value
}

func (s *OnsDLQMessageGetByIdResponseBodyDataPropertyListMessageProperty) SetName(v string) *OnsDLQMessageGetByIdResponseBodyDataPropertyListMessageProperty {
	s.Name = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyDataPropertyListMessageProperty) SetValue(v string) *OnsDLQMessageGetByIdResponseBodyDataPropertyListMessageProperty {
	s.Value = &v
	return s
}

func (s *OnsDLQMessageGetByIdResponseBodyDataPropertyListMessageProperty) Validate() error {
	return dara.Validate(s)
}
