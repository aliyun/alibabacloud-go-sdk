// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetTopicStatusResponseBody
	GetCode() *int32
	SetMessage(v string) *GetTopicStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTopicStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTopicStatusResponseBody
	GetSuccess() *bool
	SetTopicStatus(v *GetTopicStatusResponseBodyTopicStatus) *GetTopicStatusResponseBody
	GetTopicStatus() *GetTopicStatusResponseBodyTopicStatus
}

type GetTopicStatusResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
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
	// The ID of the request.
	//
	// example:
	//
	// E475C7E2-8C35-46EF-BE7D-5D2A9F5D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The status information about messages in the topic.
	TopicStatus *GetTopicStatusResponseBodyTopicStatus `json:"TopicStatus,omitempty" xml:"TopicStatus,omitempty" type:"Struct"`
}

func (s GetTopicStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTopicStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetTopicStatusResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTopicStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTopicStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTopicStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTopicStatusResponseBody) GetTopicStatus() *GetTopicStatusResponseBodyTopicStatus {
	return s.TopicStatus
}

func (s *GetTopicStatusResponseBody) SetCode(v int32) *GetTopicStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetTopicStatusResponseBody) SetMessage(v string) *GetTopicStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetTopicStatusResponseBody) SetRequestId(v string) *GetTopicStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTopicStatusResponseBody) SetSuccess(v bool) *GetTopicStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetTopicStatusResponseBody) SetTopicStatus(v *GetTopicStatusResponseBodyTopicStatus) *GetTopicStatusResponseBody {
	s.TopicStatus = v
	return s
}

func (s *GetTopicStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTopicStatusResponseBodyTopicStatus struct {
	// The time when the last consumed message was generated.
	//
	// example:
	//
	// 1566470063575
	LastTimeStamp *int64 `json:"LastTimeStamp,omitempty" xml:"LastTimeStamp,omitempty"`
	// The information about offsets in the topic.
	OffsetTable *GetTopicStatusResponseBodyTopicStatusOffsetTable `json:"OffsetTable,omitempty" xml:"OffsetTable,omitempty" type:"Struct"`
	// The number of messages in the topic.
	//
	// example:
	//
	// 423
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetTopicStatusResponseBodyTopicStatus) String() string {
	return dara.Prettify(s)
}

func (s GetTopicStatusResponseBodyTopicStatus) GoString() string {
	return s.String()
}

func (s *GetTopicStatusResponseBodyTopicStatus) GetLastTimeStamp() *int64 {
	return s.LastTimeStamp
}

func (s *GetTopicStatusResponseBodyTopicStatus) GetOffsetTable() *GetTopicStatusResponseBodyTopicStatusOffsetTable {
	return s.OffsetTable
}

func (s *GetTopicStatusResponseBodyTopicStatus) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetTopicStatusResponseBodyTopicStatus) SetLastTimeStamp(v int64) *GetTopicStatusResponseBodyTopicStatus {
	s.LastTimeStamp = &v
	return s
}

func (s *GetTopicStatusResponseBodyTopicStatus) SetOffsetTable(v *GetTopicStatusResponseBodyTopicStatusOffsetTable) *GetTopicStatusResponseBodyTopicStatus {
	s.OffsetTable = v
	return s
}

func (s *GetTopicStatusResponseBodyTopicStatus) SetTotalCount(v int64) *GetTopicStatusResponseBodyTopicStatus {
	s.TotalCount = &v
	return s
}

func (s *GetTopicStatusResponseBodyTopicStatus) Validate() error {
	return dara.Validate(s)
}

type GetTopicStatusResponseBodyTopicStatusOffsetTable struct {
	OffsetTable []*GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable `json:"OffsetTable,omitempty" xml:"OffsetTable,omitempty" type:"Repeated"`
}

func (s GetTopicStatusResponseBodyTopicStatusOffsetTable) String() string {
	return dara.Prettify(s)
}

func (s GetTopicStatusResponseBodyTopicStatusOffsetTable) GoString() string {
	return s.String()
}

func (s *GetTopicStatusResponseBodyTopicStatusOffsetTable) GetOffsetTable() []*GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable {
	return s.OffsetTable
}

func (s *GetTopicStatusResponseBodyTopicStatusOffsetTable) SetOffsetTable(v []*GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) *GetTopicStatusResponseBodyTopicStatusOffsetTable {
	s.OffsetTable = v
	return s
}

func (s *GetTopicStatusResponseBodyTopicStatusOffsetTable) Validate() error {
	return dara.Validate(s)
}

type GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable struct {
	// The last time when the partition was modified.
	//
	// example:
	//
	// 1566470063547
	LastUpdateTimestamp *int64 `json:"LastUpdateTimestamp,omitempty" xml:"LastUpdateTimestamp,omitempty"`
	// The latest offset in the partition of the topic.
	//
	// example:
	//
	// 76
	MaxOffset *int64 `json:"MaxOffset,omitempty" xml:"MaxOffset,omitempty"`
	// The earliest offset in the partition of the topic.
	//
	// example:
	//
	// 0
	MinOffset *int64 `json:"MinOffset,omitempty" xml:"MinOffset,omitempty"`
	// The ID of the partition.
	//
	// example:
	//
	// 0
	Partition *int32 `json:"Partition,omitempty" xml:"Partition,omitempty"`
	// The name of the topic.
	//
	// example:
	//
	// testkafka
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) String() string {
	return dara.Prettify(s)
}

func (s GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) GoString() string {
	return s.String()
}

func (s *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) GetLastUpdateTimestamp() *int64 {
	return s.LastUpdateTimestamp
}

func (s *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) GetMaxOffset() *int64 {
	return s.MaxOffset
}

func (s *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) GetMinOffset() *int64 {
	return s.MinOffset
}

func (s *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) GetPartition() *int32 {
	return s.Partition
}

func (s *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) GetTopic() *string {
	return s.Topic
}

func (s *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) SetLastUpdateTimestamp(v int64) *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable {
	s.LastUpdateTimestamp = &v
	return s
}

func (s *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) SetMaxOffset(v int64) *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable {
	s.MaxOffset = &v
	return s
}

func (s *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) SetMinOffset(v int64) *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable {
	s.MinOffset = &v
	return s
}

func (s *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) SetPartition(v int32) *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable {
	s.Partition = &v
	return s
}

func (s *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) SetTopic(v string) *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable {
	s.Topic = &v
	return s
}

func (s *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) Validate() error {
	return dara.Validate(s)
}
