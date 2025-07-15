// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumerProgressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetConsumerProgressResponseBody
	GetCode() *int32
	SetConsumerProgress(v *GetConsumerProgressResponseBodyConsumerProgress) *GetConsumerProgressResponseBody
	GetConsumerProgress() *GetConsumerProgressResponseBodyConsumerProgress
	SetMessage(v string) *GetConsumerProgressResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetConsumerProgressResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetConsumerProgressResponseBody
	GetSuccess() *bool
}

type GetConsumerProgressResponseBody struct {
	// The returned HTTP status code. If the request is successful, 200 is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The consumer progress of the consumer group.
	ConsumerProgress *GetConsumerProgressResponseBodyConsumerProgress `json:"ConsumerProgress,omitempty" xml:"ConsumerProgress,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 252820E1-A2E6-45F2-B4C9-1056B8CE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetConsumerProgressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerProgressResponseBody) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetConsumerProgressResponseBody) GetConsumerProgress() *GetConsumerProgressResponseBodyConsumerProgress {
	return s.ConsumerProgress
}

func (s *GetConsumerProgressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetConsumerProgressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConsumerProgressResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetConsumerProgressResponseBody) SetCode(v int32) *GetConsumerProgressResponseBody {
	s.Code = &v
	return s
}

func (s *GetConsumerProgressResponseBody) SetConsumerProgress(v *GetConsumerProgressResponseBodyConsumerProgress) *GetConsumerProgressResponseBody {
	s.ConsumerProgress = v
	return s
}

func (s *GetConsumerProgressResponseBody) SetMessage(v string) *GetConsumerProgressResponseBody {
	s.Message = &v
	return s
}

func (s *GetConsumerProgressResponseBody) SetRequestId(v string) *GetConsumerProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConsumerProgressResponseBody) SetSuccess(v bool) *GetConsumerProgressResponseBody {
	s.Success = &v
	return s
}

func (s *GetConsumerProgressResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetConsumerProgressResponseBodyConsumerProgress struct {
	// The time when the last message consumed by the consumer group was generated.
	//
	// example:
	//
	// 1566874931671
	LastTimestamp *int64 `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// The details of rebalances in the consumer group.
	RebalanceInfoList *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoList `json:"RebalanceInfoList,omitempty" xml:"RebalanceInfoList,omitempty" type:"Struct"`
	// The consumer progress of each topic to which the consumer group subscribes.
	TopicList *GetConsumerProgressResponseBodyConsumerProgressTopicList `json:"TopicList,omitempty" xml:"TopicList,omitempty" type:"Struct"`
	// The total number of unconsumed messages in all topics to which the consumer group subscribes.
	//
	// example:
	//
	// 0
	TotalDiff *int64  `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
	State     *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s GetConsumerProgressResponseBodyConsumerProgress) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerProgressResponseBodyConsumerProgress) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressResponseBodyConsumerProgress) GetLastTimestamp() *int64 {
	return s.LastTimestamp
}

func (s *GetConsumerProgressResponseBodyConsumerProgress) GetRebalanceInfoList() *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoList {
	return s.RebalanceInfoList
}

func (s *GetConsumerProgressResponseBodyConsumerProgress) GetTopicList() *GetConsumerProgressResponseBodyConsumerProgressTopicList {
	return s.TopicList
}

func (s *GetConsumerProgressResponseBodyConsumerProgress) GetTotalDiff() *int64 {
	return s.TotalDiff
}

func (s *GetConsumerProgressResponseBodyConsumerProgress) GetState() *string {
	return s.State
}

func (s *GetConsumerProgressResponseBodyConsumerProgress) SetLastTimestamp(v int64) *GetConsumerProgressResponseBodyConsumerProgress {
	s.LastTimestamp = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgress) SetRebalanceInfoList(v *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoList) *GetConsumerProgressResponseBodyConsumerProgress {
	s.RebalanceInfoList = v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgress) SetTopicList(v *GetConsumerProgressResponseBodyConsumerProgressTopicList) *GetConsumerProgressResponseBodyConsumerProgress {
	s.TopicList = v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgress) SetTotalDiff(v int64) *GetConsumerProgressResponseBodyConsumerProgress {
	s.TotalDiff = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgress) SetState(v string) *GetConsumerProgressResponseBodyConsumerProgress {
	s.State = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgress) Validate() error {
	return dara.Validate(s)
}

type GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoList struct {
	RebalanceInfoList []*GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList `json:"RebalanceInfoList,omitempty" xml:"RebalanceInfoList,omitempty" type:"Repeated"`
}

func (s GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoList) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoList) GetRebalanceInfoList() []*GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList {
	return s.RebalanceInfoList
}

func (s *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoList) SetRebalanceInfoList(v []*GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList) *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoList {
	s.RebalanceInfoList = v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoList) Validate() error {
	return dara.Validate(s)
}

type GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList struct {
	// The number of rebalances.
	//
	// example:
	//
	// 100
	Generation *int64 `json:"Generation,omitempty" xml:"Generation,omitempty"`
	// The group ID of the subscriber.
	//
	// example:
	//
	// falcon-uat
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The time when the last rebalance occurred. Unit: milliseconds.
	//
	// example:
	//
	// 1709199270
	LastRebalanceTimestamp *int64 `json:"LastRebalanceTimestamp,omitempty" xml:"LastRebalanceTimestamp,omitempty"`
	// The cause of the rebalance.
	//
	// example:
	//
	// removing member consumer-1-cd14eb9c-379b-4b8e-9bbd-76f147f8536f on LeaveGroup
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// Indicates whether new members are added to the consumer group in the rebalance.
	//
	// example:
	//
	// true
	RebalanceSuccess *bool `json:"RebalanceSuccess,omitempty" xml:"RebalanceSuccess,omitempty"`
	// The duration of the rebalance. Unit: milliseconds.
	//
	// example:
	//
	// 1
	RebalanceTimeConsuming *int64 `json:"RebalanceTimeConsuming,omitempty" xml:"RebalanceTimeConsuming,omitempty"`
}

func (s GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList) GetGeneration() *int64 {
	return s.Generation
}

func (s *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList) GetGroupId() *string {
	return s.GroupId
}

func (s *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList) GetLastRebalanceTimestamp() *int64 {
	return s.LastRebalanceTimestamp
}

func (s *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList) GetReason() *string {
	return s.Reason
}

func (s *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList) GetRebalanceSuccess() *bool {
	return s.RebalanceSuccess
}

func (s *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList) GetRebalanceTimeConsuming() *int64 {
	return s.RebalanceTimeConsuming
}

func (s *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList) SetGeneration(v int64) *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList {
	s.Generation = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList) SetGroupId(v string) *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList {
	s.GroupId = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList) SetLastRebalanceTimestamp(v int64) *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList {
	s.LastRebalanceTimestamp = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList) SetReason(v string) *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList {
	s.Reason = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList) SetRebalanceSuccess(v bool) *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList {
	s.RebalanceSuccess = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList) SetRebalanceTimeConsuming(v int64) *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList {
	s.RebalanceTimeConsuming = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList) Validate() error {
	return dara.Validate(s)
}

type GetConsumerProgressResponseBodyConsumerProgressTopicList struct {
	TopicList []*GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList `json:"TopicList,omitempty" xml:"TopicList,omitempty" type:"Repeated"`
}

func (s GetConsumerProgressResponseBodyConsumerProgressTopicList) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerProgressResponseBodyConsumerProgressTopicList) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicList) GetTopicList() []*GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList {
	return s.TopicList
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicList) SetTopicList(v []*GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList) *GetConsumerProgressResponseBodyConsumerProgressTopicList {
	s.TopicList = v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicList) Validate() error {
	return dara.Validate(s)
}

type GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList struct {
	// The time when the last consumed message in the topic was generated.
	//
	// example:
	//
	// 1566874931649
	LastTimestamp *int64 `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// The consumer offsets.
	OffsetList *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetList `json:"OffsetList,omitempty" xml:"OffsetList,omitempty" type:"Struct"`
	// The topic name.
	//
	// example:
	//
	// kafka-test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The number of unconsumed messages in the topic to which the consumer group subscribes.
	//
	// example:
	//
	// 0
	TotalDiff *int64 `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
}

func (s GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList) GetLastTimestamp() *int64 {
	return s.LastTimestamp
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList) GetOffsetList() *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetList {
	return s.OffsetList
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList) GetTopic() *string {
	return s.Topic
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList) GetTotalDiff() *int64 {
	return s.TotalDiff
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList) SetLastTimestamp(v int64) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList {
	s.LastTimestamp = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList) SetOffsetList(v *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetList) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList {
	s.OffsetList = v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList) SetTopic(v string) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList {
	s.Topic = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList) SetTotalDiff(v int64) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList {
	s.TotalDiff = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList) Validate() error {
	return dara.Validate(s)
}

type GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetList struct {
	OffsetList []*GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList `json:"OffsetList,omitempty" xml:"OffsetList,omitempty" type:"Repeated"`
}

func (s GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetList) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetList) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetList) GetOffsetList() []*GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList {
	return s.OffsetList
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetList) SetOffsetList(v []*GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetList {
	s.OffsetList = v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetList) Validate() error {
	return dara.Validate(s)
}

type GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList struct {
	// The latest offset in the partition of the topic.
	//
	// example:
	//
	// 9
	BrokerOffset *int64 `json:"BrokerOffset,omitempty" xml:"BrokerOffset,omitempty"`
	// Client ID of the application.
	//
	// example:
	//
	// client-id-KafkaConsumerDemo
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 172.16.11.3
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The consumer offset in the partition of the topic.
	//
	// example:
	//
	// 9
	ConsumerOffset *int64 `json:"ConsumerOffset,omitempty" xml:"ConsumerOffset,omitempty"`
	// The time when the last consumed message in the partition was generated.
	//
	// example:
	//
	// 1566874931649
	LastTimestamp *int64 `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// Member ID.
	//
	// example:
	//
	// client-id-KafkaConsumerDemo-70b64883-a911-4882-8084-598b958848b4
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The partition ID.
	//
	// example:
	//
	// 0
	Partition *int32 `json:"Partition,omitempty" xml:"Partition,omitempty"`
}

func (s GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) GetBrokerOffset() *int64 {
	return s.BrokerOffset
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) GetClientId() *string {
	return s.ClientId
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) GetClientIp() *string {
	return s.ClientIp
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) GetConsumerOffset() *int64 {
	return s.ConsumerOffset
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) GetLastTimestamp() *int64 {
	return s.LastTimestamp
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) GetMemberId() *string {
	return s.MemberId
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) GetPartition() *int32 {
	return s.Partition
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) SetBrokerOffset(v int64) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList {
	s.BrokerOffset = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) SetClientId(v string) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList {
	s.ClientId = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) SetClientIp(v string) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList {
	s.ClientIp = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) SetConsumerOffset(v int64) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList {
	s.ConsumerOffset = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) SetLastTimestamp(v int64) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList {
	s.LastTimestamp = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) SetMemberId(v string) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList {
	s.MemberId = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) SetPartition(v int32) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList {
	s.Partition = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) Validate() error {
	return dara.Validate(s)
}
