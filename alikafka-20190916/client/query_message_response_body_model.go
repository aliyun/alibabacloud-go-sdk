// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryMessageResponseBody
	GetCode() *int32
	SetMessage(v string) *QueryMessageResponseBody
	GetMessage() *string
	SetMessageList(v []*QueryMessageResponseBodyMessageList) *QueryMessageResponseBody
	GetMessageList() []*QueryMessageResponseBodyMessageList
	SetRequestId(v string) *QueryMessageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryMessageResponseBody
	GetSuccess() *bool
}

type QueryMessageResponseBody struct {
	// The returned HTTP status code. If the request is successful, 200 is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The messages.
	MessageList []*QueryMessageResponseBodyMessageList `json:"MessageList,omitempty" xml:"MessageList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ABA4A7FD-E10F-45C7-9774-A5236015****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMessageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMessageResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryMessageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryMessageResponseBody) GetMessageList() []*QueryMessageResponseBodyMessageList {
	return s.MessageList
}

func (s *QueryMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMessageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMessageResponseBody) SetCode(v int32) *QueryMessageResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMessageResponseBody) SetMessage(v string) *QueryMessageResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMessageResponseBody) SetMessageList(v []*QueryMessageResponseBodyMessageList) *QueryMessageResponseBody {
	s.MessageList = v
	return s
}

func (s *QueryMessageResponseBody) SetRequestId(v string) *QueryMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMessageResponseBody) SetSuccess(v bool) *QueryMessageResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMessageResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryMessageResponseBodyMessageList struct {
	// The check value of the chaincode.
	//
	// example:
	//
	// 0
	Checksum *int64 `json:"Checksum,omitempty" xml:"Checksum,omitempty"`
	// The message key.
	//
	// example:
	//
	// this is key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Indicates whether the key is truncated.
	//
	// example:
	//
	// false
	KeyTruncated *bool `json:"KeyTruncated,omitempty" xml:"KeyTruncated,omitempty"`
	// The consumer offset of the partition.
	//
	// example:
	//
	// 1
	Offset *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The partition ID.
	//
	// example:
	//
	// 0
	Partition *int64 `json:"Partition,omitempty" xml:"Partition,omitempty"`
	// The size of the key after serialization. Unit: bytes.
	//
	// example:
	//
	// 11
	SerializedKeySize *int32 `json:"SerializedKeySize,omitempty" xml:"SerializedKeySize,omitempty"`
	// The size of the value after serialization. Unit: bytes.
	//
	// example:
	//
	// 20
	SerializedValueSize *int32 `json:"SerializedValueSize,omitempty" xml:"SerializedValueSize,omitempty"`
	// The time when the message was created. The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1705482172640
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The time type.
	//
	// example:
	//
	// CreateTime
	TimestampType *string `json:"TimestampType,omitempty" xml:"TimestampType,omitempty"`
	// The topic name.
	//
	// example:
	//
	// dqc_test2
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The truncated size of the message key. Unit: bytes.
	//
	// >  A maximum of 1 KB of content can be displayed for each message. Content that exceeds 1 KB is automatically truncated. For more information, see [Query messages](https://help.aliyun.com/document_detail/113172.html).
	//
	// example:
	//
	// 0
	TruncatedKeySize *int32 `json:"TruncatedKeySize,omitempty" xml:"TruncatedKeySize,omitempty"`
	// The truncated size of the message value. Unit: bytes.
	//
	// >  A maximum of 1 KB of content can be displayed for each message. Content that exceeds 1 KB is automatically truncated. For more information, see [Query messages](https://help.aliyun.com/document_detail/113172.html).
	//
	// example:
	//
	// 0
	TruncatedValueSize *int32 `json:"TruncatedValueSize,omitempty" xml:"TruncatedValueSize,omitempty"`
	// The message value.
	//
	// example:
	//
	// Welcome to Ali kafka
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// Indicates whether the value is truncated.
	//
	// example:
	//
	// false
	ValueTruncated *bool `json:"ValueTruncated,omitempty" xml:"ValueTruncated,omitempty"`
}

func (s QueryMessageResponseBodyMessageList) String() string {
	return dara.Prettify(s)
}

func (s QueryMessageResponseBodyMessageList) GoString() string {
	return s.String()
}

func (s *QueryMessageResponseBodyMessageList) GetChecksum() *int64 {
	return s.Checksum
}

func (s *QueryMessageResponseBodyMessageList) GetKey() *string {
	return s.Key
}

func (s *QueryMessageResponseBodyMessageList) GetKeyTruncated() *bool {
	return s.KeyTruncated
}

func (s *QueryMessageResponseBodyMessageList) GetOffset() *int64 {
	return s.Offset
}

func (s *QueryMessageResponseBodyMessageList) GetPartition() *int64 {
	return s.Partition
}

func (s *QueryMessageResponseBodyMessageList) GetSerializedKeySize() *int32 {
	return s.SerializedKeySize
}

func (s *QueryMessageResponseBodyMessageList) GetSerializedValueSize() *int32 {
	return s.SerializedValueSize
}

func (s *QueryMessageResponseBodyMessageList) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *QueryMessageResponseBodyMessageList) GetTimestampType() *string {
	return s.TimestampType
}

func (s *QueryMessageResponseBodyMessageList) GetTopic() *string {
	return s.Topic
}

func (s *QueryMessageResponseBodyMessageList) GetTruncatedKeySize() *int32 {
	return s.TruncatedKeySize
}

func (s *QueryMessageResponseBodyMessageList) GetTruncatedValueSize() *int32 {
	return s.TruncatedValueSize
}

func (s *QueryMessageResponseBodyMessageList) GetValue() *string {
	return s.Value
}

func (s *QueryMessageResponseBodyMessageList) GetValueTruncated() *bool {
	return s.ValueTruncated
}

func (s *QueryMessageResponseBodyMessageList) SetChecksum(v int64) *QueryMessageResponseBodyMessageList {
	s.Checksum = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) SetKey(v string) *QueryMessageResponseBodyMessageList {
	s.Key = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) SetKeyTruncated(v bool) *QueryMessageResponseBodyMessageList {
	s.KeyTruncated = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) SetOffset(v int64) *QueryMessageResponseBodyMessageList {
	s.Offset = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) SetPartition(v int64) *QueryMessageResponseBodyMessageList {
	s.Partition = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) SetSerializedKeySize(v int32) *QueryMessageResponseBodyMessageList {
	s.SerializedKeySize = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) SetSerializedValueSize(v int32) *QueryMessageResponseBodyMessageList {
	s.SerializedValueSize = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) SetTimestamp(v int64) *QueryMessageResponseBodyMessageList {
	s.Timestamp = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) SetTimestampType(v string) *QueryMessageResponseBodyMessageList {
	s.TimestampType = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) SetTopic(v string) *QueryMessageResponseBodyMessageList {
	s.Topic = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) SetTruncatedKeySize(v int32) *QueryMessageResponseBodyMessageList {
	s.TruncatedKeySize = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) SetTruncatedValueSize(v int32) *QueryMessageResponseBodyMessageList {
	s.TruncatedValueSize = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) SetValue(v string) *QueryMessageResponseBodyMessageList {
	s.Value = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) SetValueTruncated(v bool) *QueryMessageResponseBodyMessageList {
	s.ValueTruncated = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) Validate() error {
	return dara.Validate(s)
}
