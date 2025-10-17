// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConsumerGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerChannels(v *DescribeConsumerGroupResponseBodyConsumerChannels) *DescribeConsumerGroupResponseBody
	GetConsumerChannels() *DescribeConsumerGroupResponseBodyConsumerChannels
	SetErrCode(v string) *DescribeConsumerGroupResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeConsumerGroupResponseBody
	GetErrMessage() *string
	SetPageNumber(v int32) *DescribeConsumerGroupResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeConsumerGroupResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeConsumerGroupResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeConsumerGroupResponseBody
	GetSuccess() *string
	SetTotalRecordCount(v int32) *DescribeConsumerGroupResponseBody
	GetTotalRecordCount() *int32
}

type DescribeConsumerGroupResponseBody struct {
	// The list of consumer groups.
	ConsumerChannels *DescribeConsumerGroupResponseBodyConsumerChannels `json:"ConsumerChannels,omitempty" xml:"ConsumerChannels,omitempty" type:"Struct"`
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
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of consumer groups that can be displayed on one page.
	//
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4204E899-8193-4D7D-A4FB-3A7F9063****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
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
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeConsumerGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConsumerGroupResponseBody) GetConsumerChannels() *DescribeConsumerGroupResponseBodyConsumerChannels {
	return s.ConsumerChannels
}

func (s *DescribeConsumerGroupResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeConsumerGroupResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeConsumerGroupResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeConsumerGroupResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeConsumerGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeConsumerGroupResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeConsumerGroupResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeConsumerGroupResponseBody) SetConsumerChannels(v *DescribeConsumerGroupResponseBodyConsumerChannels) *DescribeConsumerGroupResponseBody {
	s.ConsumerChannels = v
	return s
}

func (s *DescribeConsumerGroupResponseBody) SetErrCode(v string) *DescribeConsumerGroupResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeConsumerGroupResponseBody) SetErrMessage(v string) *DescribeConsumerGroupResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeConsumerGroupResponseBody) SetPageNumber(v int32) *DescribeConsumerGroupResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeConsumerGroupResponseBody) SetPageRecordCount(v int32) *DescribeConsumerGroupResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeConsumerGroupResponseBody) SetRequestId(v string) *DescribeConsumerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConsumerGroupResponseBody) SetSuccess(v string) *DescribeConsumerGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeConsumerGroupResponseBody) SetTotalRecordCount(v int32) *DescribeConsumerGroupResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeConsumerGroupResponseBody) Validate() error {
	if s.ConsumerChannels != nil {
		if err := s.ConsumerChannels.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeConsumerGroupResponseBodyConsumerChannels struct {
	DescribeConsumerChannel []*DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel `json:"DescribeConsumerChannel,omitempty" xml:"DescribeConsumerChannel,omitempty" type:"Repeated"`
}

func (s DescribeConsumerGroupResponseBodyConsumerChannels) String() string {
	return dara.Prettify(s)
}

func (s DescribeConsumerGroupResponseBodyConsumerChannels) GoString() string {
	return s.String()
}

func (s *DescribeConsumerGroupResponseBodyConsumerChannels) GetDescribeConsumerChannel() []*DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel {
	return s.DescribeConsumerChannel
}

func (s *DescribeConsumerGroupResponseBodyConsumerChannels) SetDescribeConsumerChannel(v []*DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel) *DescribeConsumerGroupResponseBodyConsumerChannels {
	s.DescribeConsumerChannel = v
	return s
}

func (s *DescribeConsumerGroupResponseBodyConsumerChannels) Validate() error {
	if s.DescribeConsumerChannel != nil {
		for _, item := range s.DescribeConsumerChannel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel struct {
	// The ID of the consumer group.
	//
	// example:
	//
	// dtspis1110z232****
	ConsumerGroupID *string `json:"ConsumerGroupID,omitempty" xml:"ConsumerGroupID,omitempty"`
	// The name of the consumer group.
	//
	// example:
	//
	// consumergrouptest
	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" xml:"ConsumerGroupName,omitempty"`
	// The username of the consumer group.
	//
	// example:
	//
	// test
	ConsumerGroupUserName *string `json:"ConsumerGroupUserName,omitempty" xml:"ConsumerGroupUserName,omitempty"`
	// The consumption checkpoint, which is the time when the latest data record was consumed by the change tracking client. The format is *yyyy-MM-dd*T*HH:mm:ss*Z. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-10-02T12:00:00Z
	ConsumptionCheckpoint *string `json:"ConsumptionCheckpoint,omitempty" xml:"ConsumptionCheckpoint,omitempty"`
	// The message delay, which is the current time minus the timestamp of the earliest unconsumed message in the change tracking instance. Unit: seconds.
	//
	// >  If the return value of this parameter is **-1**, no client is connected to the consumer group.
	//
	// example:
	//
	// 172714
	MessageDelay *int64 `json:"MessageDelay,omitempty" xml:"MessageDelay,omitempty"`
	// The total number of unconsumed messages, which is the number of unconsumed data records plus the number of heartbeat messages.
	//
	// >  If the return value of this parameter is **-1**, no client is connected to the consumer group.
	//
	// example:
	//
	// 186600
	UnconsumedData *int64 `json:"UnconsumedData,omitempty" xml:"UnconsumedData,omitempty"`
}

func (s DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel) String() string {
	return dara.Prettify(s)
}

func (s DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel) GoString() string {
	return s.String()
}

func (s *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel) GetConsumerGroupID() *string {
	return s.ConsumerGroupID
}

func (s *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel) GetConsumerGroupName() *string {
	return s.ConsumerGroupName
}

func (s *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel) GetConsumerGroupUserName() *string {
	return s.ConsumerGroupUserName
}

func (s *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel) GetConsumptionCheckpoint() *string {
	return s.ConsumptionCheckpoint
}

func (s *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel) GetMessageDelay() *int64 {
	return s.MessageDelay
}

func (s *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel) GetUnconsumedData() *int64 {
	return s.UnconsumedData
}

func (s *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel) SetConsumerGroupID(v string) *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel {
	s.ConsumerGroupID = &v
	return s
}

func (s *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel) SetConsumerGroupName(v string) *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel {
	s.ConsumerGroupName = &v
	return s
}

func (s *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel) SetConsumerGroupUserName(v string) *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel {
	s.ConsumerGroupUserName = &v
	return s
}

func (s *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel) SetConsumptionCheckpoint(v string) *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel {
	s.ConsumptionCheckpoint = &v
	return s
}

func (s *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel) SetMessageDelay(v int64) *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel {
	s.MessageDelay = &v
	return s
}

func (s *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel) SetUnconsumedData(v int64) *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel {
	s.UnconsumedData = &v
	return s
}

func (s *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel) Validate() error {
	return dara.Validate(s)
}
