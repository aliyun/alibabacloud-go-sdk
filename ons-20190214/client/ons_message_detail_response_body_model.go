// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsMessageDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OnsMessageDetailResponseBodyData) *OnsMessageDetailResponseBody
	GetData() *OnsMessageDetailResponseBodyData
	SetRequestId(v string) *OnsMessageDetailResponseBody
	GetRequestId() *string
}

type OnsMessageDetailResponseBody struct {
	// The data returned.
	Data *OnsMessageDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// EAE5BE23-37A1-4354-94D6-E44AE17E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsMessageDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsMessageDetailResponseBody) GoString() string {
	return s.String()
}

func (s *OnsMessageDetailResponseBody) GetData() *OnsMessageDetailResponseBodyData {
	return s.Data
}

func (s *OnsMessageDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsMessageDetailResponseBody) SetData(v *OnsMessageDetailResponseBodyData) *OnsMessageDetailResponseBody {
	s.Data = v
	return s
}

func (s *OnsMessageDetailResponseBody) SetRequestId(v string) *OnsMessageDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsMessageDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsMessageDetailResponseBodyData struct {
	// The string that is obtained after the message body is encrypted by using the Base 64 algorithm.
	//
	// example:
	//
	// aGVsbG8=
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The cyclic redundancy check (CRC) value of the message body.
	//
	// example:
	//
	// 907060870
	BodyCRC *int32 `json:"BodyCRC,omitempty" xml:"BodyCRC,omitempty"`
	// The information about the message body.
	//
	// example:
	//
	// hello
	BodyStr *string `json:"BodyStr,omitempty" xml:"BodyStr,omitempty"`
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
	// The ID of the ApsaraMQ for RocketMQ Instance.
	//
	// example:
	//
	// MQ_INST_184681981******_BXig0x6A
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the message.
	//
	// example:
	//
	// 1E0578FE110F18B4AAC235C05F2*****
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The attributes of the message.
	PropertyList []*OnsMessageDetailResponseBodyDataPropertyList `json:"PropertyList,omitempty" xml:"PropertyList,omitempty" type:"Repeated"`
	// The number of retries that ApsaraMQ for RocketMQ performed to send the message to consumers.
	//
	// example:
	//
	// 0
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
	// 2
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

func (s OnsMessageDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OnsMessageDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsMessageDetailResponseBodyData) GetBody() *string {
	return s.Body
}

func (s *OnsMessageDetailResponseBodyData) GetBodyCRC() *int32 {
	return s.BodyCRC
}

func (s *OnsMessageDetailResponseBodyData) GetBodyStr() *string {
	return s.BodyStr
}

func (s *OnsMessageDetailResponseBodyData) GetBornHost() *string {
	return s.BornHost
}

func (s *OnsMessageDetailResponseBodyData) GetBornTimestamp() *int64 {
	return s.BornTimestamp
}

func (s *OnsMessageDetailResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsMessageDetailResponseBodyData) GetMsgId() *string {
	return s.MsgId
}

func (s *OnsMessageDetailResponseBodyData) GetPropertyList() []*OnsMessageDetailResponseBodyDataPropertyList {
	return s.PropertyList
}

func (s *OnsMessageDetailResponseBodyData) GetReconsumeTimes() *int32 {
	return s.ReconsumeTimes
}

func (s *OnsMessageDetailResponseBodyData) GetStoreHost() *string {
	return s.StoreHost
}

func (s *OnsMessageDetailResponseBodyData) GetStoreSize() *int32 {
	return s.StoreSize
}

func (s *OnsMessageDetailResponseBodyData) GetStoreTimestamp() *int64 {
	return s.StoreTimestamp
}

func (s *OnsMessageDetailResponseBodyData) GetTopic() *string {
	return s.Topic
}

func (s *OnsMessageDetailResponseBodyData) SetBody(v string) *OnsMessageDetailResponseBodyData {
	s.Body = &v
	return s
}

func (s *OnsMessageDetailResponseBodyData) SetBodyCRC(v int32) *OnsMessageDetailResponseBodyData {
	s.BodyCRC = &v
	return s
}

func (s *OnsMessageDetailResponseBodyData) SetBodyStr(v string) *OnsMessageDetailResponseBodyData {
	s.BodyStr = &v
	return s
}

func (s *OnsMessageDetailResponseBodyData) SetBornHost(v string) *OnsMessageDetailResponseBodyData {
	s.BornHost = &v
	return s
}

func (s *OnsMessageDetailResponseBodyData) SetBornTimestamp(v int64) *OnsMessageDetailResponseBodyData {
	s.BornTimestamp = &v
	return s
}

func (s *OnsMessageDetailResponseBodyData) SetInstanceId(v string) *OnsMessageDetailResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *OnsMessageDetailResponseBodyData) SetMsgId(v string) *OnsMessageDetailResponseBodyData {
	s.MsgId = &v
	return s
}

func (s *OnsMessageDetailResponseBodyData) SetPropertyList(v []*OnsMessageDetailResponseBodyDataPropertyList) *OnsMessageDetailResponseBodyData {
	s.PropertyList = v
	return s
}

func (s *OnsMessageDetailResponseBodyData) SetReconsumeTimes(v int32) *OnsMessageDetailResponseBodyData {
	s.ReconsumeTimes = &v
	return s
}

func (s *OnsMessageDetailResponseBodyData) SetStoreHost(v string) *OnsMessageDetailResponseBodyData {
	s.StoreHost = &v
	return s
}

func (s *OnsMessageDetailResponseBodyData) SetStoreSize(v int32) *OnsMessageDetailResponseBodyData {
	s.StoreSize = &v
	return s
}

func (s *OnsMessageDetailResponseBodyData) SetStoreTimestamp(v int64) *OnsMessageDetailResponseBodyData {
	s.StoreTimestamp = &v
	return s
}

func (s *OnsMessageDetailResponseBodyData) SetTopic(v string) *OnsMessageDetailResponseBodyData {
	s.Topic = &v
	return s
}

func (s *OnsMessageDetailResponseBodyData) Validate() error {
	if s.PropertyList != nil {
		for _, item := range s.PropertyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OnsMessageDetailResponseBodyDataPropertyList struct {
	// The name of the attribute. Valid values:
	//
	// 	- **TRACE_ON**: indicates whether the trace of the message exists.
	//
	// 	- **MSG_REGION**: The region ID of the instance to which the topic belongs.
	//
	// 	- **__MESSAGE_DECODED_TIME**: The time when the message was decoded.
	//
	// 	- **__BORNHOST**: The IP address of the producer client.
	//
	// 	- **UNIQ_KEY**: The ID of the message.
	//
	// For information about the terms that are used in ApsaraMQ for RocketMQ, see [Terms](https://help.aliyun.com/document_detail/29533.html).
	//
	// example:
	//
	// cn-hangzhou
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the attribute.
	//
	// example:
	//
	// MSG_REGION
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsMessageDetailResponseBodyDataPropertyList) String() string {
	return dara.Prettify(s)
}

func (s OnsMessageDetailResponseBodyDataPropertyList) GoString() string {
	return s.String()
}

func (s *OnsMessageDetailResponseBodyDataPropertyList) GetName() *string {
	return s.Name
}

func (s *OnsMessageDetailResponseBodyDataPropertyList) GetValue() *string {
	return s.Value
}

func (s *OnsMessageDetailResponseBodyDataPropertyList) SetName(v string) *OnsMessageDetailResponseBodyDataPropertyList {
	s.Name = &v
	return s
}

func (s *OnsMessageDetailResponseBodyDataPropertyList) SetValue(v string) *OnsMessageDetailResponseBodyDataPropertyList {
	s.Value = &v
	return s
}

func (s *OnsMessageDetailResponseBodyDataPropertyList) Validate() error {
	return dara.Validate(s)
}
