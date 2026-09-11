// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConsumerChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerChannels(v []*DescribeConsumerChannelResponseBodyConsumerChannels) *DescribeConsumerChannelResponseBody
	GetConsumerChannels() []*DescribeConsumerChannelResponseBodyConsumerChannels
	SetErrCode(v string) *DescribeConsumerChannelResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeConsumerChannelResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *DescribeConsumerChannelResponseBody
	GetHttpStatusCode() *string
	SetPageNumber(v int32) *DescribeConsumerChannelResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeConsumerChannelResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeConsumerChannelResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeConsumerChannelResponseBody
	GetSuccess() *string
	SetTotalRecordCount(v int64) *DescribeConsumerChannelResponseBody
	GetTotalRecordCount() *int64
}

type DescribeConsumerChannelResponseBody struct {
	// The list of consumer groups.
	ConsumerChannels []*DescribeConsumerChannelResponseBodyConsumerChannels `json:"ConsumerChannels,omitempty" xml:"ConsumerChannels,omitempty" type:"Repeated"`
	// The error code returned if the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of consumer groups that can be displayed on one page.
	//
	// example:
	//
	// 20
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D66140B3-C747-42B6-8315-BAF6490E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of consumer groups.
	//
	// example:
	//
	// 1
	TotalRecordCount *int64 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeConsumerChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeConsumerChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConsumerChannelResponseBody) GetConsumerChannels() []*DescribeConsumerChannelResponseBodyConsumerChannels {
	return s.ConsumerChannels
}

func (s *DescribeConsumerChannelResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeConsumerChannelResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeConsumerChannelResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeConsumerChannelResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeConsumerChannelResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeConsumerChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeConsumerChannelResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeConsumerChannelResponseBody) GetTotalRecordCount() *int64 {
	return s.TotalRecordCount
}

func (s *DescribeConsumerChannelResponseBody) SetConsumerChannels(v []*DescribeConsumerChannelResponseBodyConsumerChannels) *DescribeConsumerChannelResponseBody {
	s.ConsumerChannels = v
	return s
}

func (s *DescribeConsumerChannelResponseBody) SetErrCode(v string) *DescribeConsumerChannelResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeConsumerChannelResponseBody) SetErrMessage(v string) *DescribeConsumerChannelResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeConsumerChannelResponseBody) SetHttpStatusCode(v string) *DescribeConsumerChannelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeConsumerChannelResponseBody) SetPageNumber(v int32) *DescribeConsumerChannelResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeConsumerChannelResponseBody) SetPageRecordCount(v int32) *DescribeConsumerChannelResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeConsumerChannelResponseBody) SetRequestId(v string) *DescribeConsumerChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConsumerChannelResponseBody) SetSuccess(v string) *DescribeConsumerChannelResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeConsumerChannelResponseBody) SetTotalRecordCount(v int64) *DescribeConsumerChannelResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeConsumerChannelResponseBody) Validate() error {
	if s.ConsumerChannels != nil {
		for _, item := range s.ConsumerChannels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeConsumerChannelResponseBodyConsumerChannels struct {
	// The consumer group ID.
	//
	// example:
	//
	// dtsor2y66j4219****
	ConsumerGroupId *string `json:"ConsumerGroupId,omitempty" xml:"ConsumerGroupId,omitempty"`
	// The name of the consumer group.
	//
	// example:
	//
	// 订阅组A
	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" xml:"ConsumerGroupName,omitempty"`
	// The account of the consumer group.
	//
	// example:
	//
	// dtstest
	ConsumerGroupUserName *string `json:"ConsumerGroupUserName,omitempty" xml:"ConsumerGroupUserName,omitempty"`
	// The consumption checkpoint, which is the point in time when the client consumed the last message in the subscription channel. The time is displayed in the yyyy-MM-ddTHH:mm:ssZ format (UTC).
	//
	// example:
	//
	// 2021-06-20T12:00:00Z
	ConsumptionCheckpoint *string `json:"ConsumptionCheckpoint,omitempty" xml:"ConsumptionCheckpoint,omitempty"`
	// The message delay. This value is calculated as the timestamp of the latest data consumed by the downstream client minus the timestamp of the latest data in the change tracking task. The value is a UNIX timestamp. Unit: seconds.
	//
	// For example, if the latest data in the source database was generated at 10:00, the DTS change tracking task has read data up to 09:55, and the downstream client has consumed data up to 09:30, the message delay is the difference in UNIX timestamps between 09:55 and 09:30.
	//
	// > If this parameter returns **-1**, no client is connected to the consumer group.
	//
	// example:
	//
	// 1500
	MessageDelay *int64 `json:"MessageDelay,omitempty" xml:"MessageDelay,omitempty"`
	// The total number of unconsumed messages, which is the sum of unconsumed subscription data and heartbeat messages.
	//
	// > If this parameter returns -1, no client is connected to the consumer group.
	//
	// example:
	//
	// 186600
	UnconsumedData *int64 `json:"UnconsumedData,omitempty" xml:"UnconsumedData,omitempty"`
}

func (s DescribeConsumerChannelResponseBodyConsumerChannels) String() string {
	return dara.Prettify(s)
}

func (s DescribeConsumerChannelResponseBodyConsumerChannels) GoString() string {
	return s.String()
}

func (s *DescribeConsumerChannelResponseBodyConsumerChannels) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *DescribeConsumerChannelResponseBodyConsumerChannels) GetConsumerGroupName() *string {
	return s.ConsumerGroupName
}

func (s *DescribeConsumerChannelResponseBodyConsumerChannels) GetConsumerGroupUserName() *string {
	return s.ConsumerGroupUserName
}

func (s *DescribeConsumerChannelResponseBodyConsumerChannels) GetConsumptionCheckpoint() *string {
	return s.ConsumptionCheckpoint
}

func (s *DescribeConsumerChannelResponseBodyConsumerChannels) GetMessageDelay() *int64 {
	return s.MessageDelay
}

func (s *DescribeConsumerChannelResponseBodyConsumerChannels) GetUnconsumedData() *int64 {
	return s.UnconsumedData
}

func (s *DescribeConsumerChannelResponseBodyConsumerChannels) SetConsumerGroupId(v string) *DescribeConsumerChannelResponseBodyConsumerChannels {
	s.ConsumerGroupId = &v
	return s
}

func (s *DescribeConsumerChannelResponseBodyConsumerChannels) SetConsumerGroupName(v string) *DescribeConsumerChannelResponseBodyConsumerChannels {
	s.ConsumerGroupName = &v
	return s
}

func (s *DescribeConsumerChannelResponseBodyConsumerChannels) SetConsumerGroupUserName(v string) *DescribeConsumerChannelResponseBodyConsumerChannels {
	s.ConsumerGroupUserName = &v
	return s
}

func (s *DescribeConsumerChannelResponseBodyConsumerChannels) SetConsumptionCheckpoint(v string) *DescribeConsumerChannelResponseBodyConsumerChannels {
	s.ConsumptionCheckpoint = &v
	return s
}

func (s *DescribeConsumerChannelResponseBodyConsumerChannels) SetMessageDelay(v int64) *DescribeConsumerChannelResponseBodyConsumerChannels {
	s.MessageDelay = &v
	return s
}

func (s *DescribeConsumerChannelResponseBodyConsumerChannels) SetUnconsumedData(v int64) *DescribeConsumerChannelResponseBodyConsumerChannels {
	s.UnconsumedData = &v
	return s
}

func (s *DescribeConsumerChannelResponseBodyConsumerChannels) Validate() error {
	return dara.Validate(s)
}
