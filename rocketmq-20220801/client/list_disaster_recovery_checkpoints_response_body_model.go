// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDisasterRecoveryCheckpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDisasterRecoveryCheckpointsResponseBody
	GetCode() *string
	SetData(v *ListDisasterRecoveryCheckpointsResponseBodyData) *ListDisasterRecoveryCheckpointsResponseBody
	GetData() *ListDisasterRecoveryCheckpointsResponseBodyData
	SetDynamicCode(v string) *ListDisasterRecoveryCheckpointsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListDisasterRecoveryCheckpointsResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ListDisasterRecoveryCheckpointsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDisasterRecoveryCheckpointsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDisasterRecoveryCheckpointsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDisasterRecoveryCheckpointsResponseBody
	GetSuccess() *bool
}

type ListDisasterRecoveryCheckpointsResponseBody struct {
	// Error code
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response Data
	Data *ListDisasterRecoveryCheckpointsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Dynamic error code
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Error message
	//
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID
	//
	// example:
	//
	// AF9A8B10-C426-530F-A0DD-96320B39****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Whether the operation was successful
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListDisasterRecoveryCheckpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDisasterRecoveryCheckpointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDisasterRecoveryCheckpointsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDisasterRecoveryCheckpointsResponseBody) GetData() *ListDisasterRecoveryCheckpointsResponseBodyData {
	return s.Data
}

func (s *ListDisasterRecoveryCheckpointsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListDisasterRecoveryCheckpointsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListDisasterRecoveryCheckpointsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDisasterRecoveryCheckpointsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDisasterRecoveryCheckpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDisasterRecoveryCheckpointsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDisasterRecoveryCheckpointsResponseBody) SetCode(v string) *ListDisasterRecoveryCheckpointsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBody) SetData(v *ListDisasterRecoveryCheckpointsResponseBodyData) *ListDisasterRecoveryCheckpointsResponseBody {
	s.Data = v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBody) SetDynamicCode(v string) *ListDisasterRecoveryCheckpointsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBody) SetDynamicMessage(v string) *ListDisasterRecoveryCheckpointsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBody) SetHttpStatusCode(v int32) *ListDisasterRecoveryCheckpointsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBody) SetMessage(v string) *ListDisasterRecoveryCheckpointsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBody) SetRequestId(v string) *ListDisasterRecoveryCheckpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBody) SetSuccess(v bool) *ListDisasterRecoveryCheckpointsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDisasterRecoveryCheckpointsResponseBodyData struct {
	// Paged data
	List []*ListDisasterRecoveryCheckpointsResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// Current page number
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Page size
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Total number of records
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListDisasterRecoveryCheckpointsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDisasterRecoveryCheckpointsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyData) GetList() []*ListDisasterRecoveryCheckpointsResponseBodyDataList {
	return s.List
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyData) SetList(v []*ListDisasterRecoveryCheckpointsResponseBodyDataList) *ListDisasterRecoveryCheckpointsResponseBodyData {
	s.List = v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyData) SetPageNumber(v int64) *ListDisasterRecoveryCheckpointsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyData) SetPageSize(v int64) *ListDisasterRecoveryCheckpointsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyData) SetTotalCount(v int64) *ListDisasterRecoveryCheckpointsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDisasterRecoveryCheckpointsResponseBodyDataList struct {
	// Consumption Progress ID
	//
	// example:
	//
	// 10000000xx
	CheckpointId *int64 `json:"checkpointId,omitempty" xml:"checkpointId,omitempty"`
	// Backup Mapping ID
	//
	// example:
	//
	// 10000000xx
	ItemId *int64 `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// Last synchronization time
	//
	// example:
	//
	// 1740724080343
	LastSyncTime *int64 `json:"lastSyncTime,omitempty" xml:"lastSyncTime,omitempty"`
	// Backup Plan ID
	//
	// example:
	//
	// 13000000xx
	PlanId *int64 `json:"planId,omitempty" xml:"planId,omitempty"`
	// Source consumption progress
	SourceProgress *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress `json:"sourceProgress,omitempty" xml:"sourceProgress,omitempty" type:"Struct"`
	// Target consumption progress
	TargetProgress *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress `json:"targetProgress,omitempty" xml:"targetProgress,omitempty" type:"Struct"`
}

func (s ListDisasterRecoveryCheckpointsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListDisasterRecoveryCheckpointsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataList) GetCheckpointId() *int64 {
	return s.CheckpointId
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataList) GetItemId() *int64 {
	return s.ItemId
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataList) GetLastSyncTime() *int64 {
	return s.LastSyncTime
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataList) GetPlanId() *int64 {
	return s.PlanId
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataList) GetSourceProgress() *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress {
	return s.SourceProgress
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataList) GetTargetProgress() *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress {
	return s.TargetProgress
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataList) SetCheckpointId(v int64) *ListDisasterRecoveryCheckpointsResponseBodyDataList {
	s.CheckpointId = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataList) SetItemId(v int64) *ListDisasterRecoveryCheckpointsResponseBodyDataList {
	s.ItemId = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataList) SetLastSyncTime(v int64) *ListDisasterRecoveryCheckpointsResponseBodyDataList {
	s.LastSyncTime = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataList) SetPlanId(v int64) *ListDisasterRecoveryCheckpointsResponseBodyDataList {
	s.PlanId = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataList) SetSourceProgress(v *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress) *ListDisasterRecoveryCheckpointsResponseBodyDataList {
	s.SourceProgress = v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataList) SetTargetProgress(v *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress) *ListDisasterRecoveryCheckpointsResponseBodyDataList {
	s.TargetProgress = v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress struct {
	// Consumer Group ID
	//
	// example:
	//
	// GID_TEST
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// Instance ID
	//
	// example:
	//
	// rmq-cn-3mp3vblzxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// Instance type
	//
	//   - ALIYUN_ROCKETMQ: Alibaba Cloud instance
	//
	//   - EXTERNAL_ROCKETMQ: External instance, open-source instance, open-source cluster
	//
	// example:
	//
	// ALIYUN_ROCKETMQ
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// Last fetch time
	//
	// example:
	//
	// 1740724080343
	LastFetchTime *int64 `json:"lastFetchTime,omitempty" xml:"lastFetchTime,omitempty"`
	// Consumption progress data
	ProgressData *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgressProgressData `json:"progressData,omitempty" xml:"progressData,omitempty" type:"Struct"`
	// Region ID
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The topic name.
	//
	// example:
	//
	// TOPIC_TEST
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress) String() string {
	return dara.Prettify(s)
}

func (s ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress) GoString() string {
	return s.String()
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress) GetLastFetchTime() *int64 {
	return s.LastFetchTime
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress) GetProgressData() *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgressProgressData {
	return s.ProgressData
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress) GetTopicName() *string {
	return s.TopicName
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress) SetConsumerGroupId(v string) *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress {
	s.ConsumerGroupId = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress) SetInstanceId(v string) *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress {
	s.InstanceId = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress) SetInstanceType(v string) *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress {
	s.InstanceType = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress) SetLastFetchTime(v int64) *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress {
	s.LastFetchTime = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress) SetProgressData(v *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgressProgressData) *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress {
	s.ProgressData = v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress) SetRegionId(v string) *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress {
	s.RegionId = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress) SetTopicName(v string) *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress {
	s.TopicName = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgress) Validate() error {
	return dara.Validate(s)
}

type ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgressProgressData struct {
	// Latest consumption time
	//
	// example:
	//
	// 1740724080343
	ConsumeTimestamp *int64 `json:"consumeTimestamp,omitempty" xml:"consumeTimestamp,omitempty"`
}

func (s ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgressProgressData) String() string {
	return dara.Prettify(s)
}

func (s ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgressProgressData) GoString() string {
	return s.String()
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgressProgressData) GetConsumeTimestamp() *int64 {
	return s.ConsumeTimestamp
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgressProgressData) SetConsumeTimestamp(v int64) *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgressProgressData {
	s.ConsumeTimestamp = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListSourceProgressProgressData) Validate() error {
	return dara.Validate(s)
}

type ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress struct {
	// Consumer group ID
	//
	// example:
	//
	// GID_TEST
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// Instance ID
	//
	// example:
	//
	// rmq-cn-nwy3i065xxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// Instance type
	//
	//   - ALIYUN_ROCKETMQ: Alibaba Cloud instance
	//
	//   - EXTERNAL_ROCKETMQ: External instance, open-source instance, open-source cluster
	//
	// example:
	//
	// ALIYUN_ROCKETMQ
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// Latest fetch time
	//
	// example:
	//
	// 1740724080343
	LastFetchTime *int64 `json:"lastFetchTime,omitempty" xml:"lastFetchTime,omitempty"`
	// Consumption progress data
	ProgressData *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgressProgressData `json:"progressData,omitempty" xml:"progressData,omitempty" type:"Struct"`
	// Region ID
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Topic name
	//
	// example:
	//
	// TOPIC_TEST
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress) String() string {
	return dara.Prettify(s)
}

func (s ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress) GoString() string {
	return s.String()
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress) GetLastFetchTime() *int64 {
	return s.LastFetchTime
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress) GetProgressData() *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgressProgressData {
	return s.ProgressData
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress) GetTopicName() *string {
	return s.TopicName
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress) SetConsumerGroupId(v string) *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress {
	s.ConsumerGroupId = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress) SetInstanceId(v string) *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress {
	s.InstanceId = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress) SetInstanceType(v string) *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress {
	s.InstanceType = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress) SetLastFetchTime(v int64) *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress {
	s.LastFetchTime = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress) SetProgressData(v *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgressProgressData) *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress {
	s.ProgressData = v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress) SetRegionId(v string) *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress {
	s.RegionId = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress) SetTopicName(v string) *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress {
	s.TopicName = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgress) Validate() error {
	return dara.Validate(s)
}

type ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgressProgressData struct {
	// Latest consumption time
	//
	// example:
	//
	// 1740724080343
	ConsumeTimestamp *int64 `json:"consumeTimestamp,omitempty" xml:"consumeTimestamp,omitempty"`
}

func (s ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgressProgressData) String() string {
	return dara.Prettify(s)
}

func (s ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgressProgressData) GoString() string {
	return s.String()
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgressProgressData) GetConsumeTimestamp() *int64 {
	return s.ConsumeTimestamp
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgressProgressData) SetConsumeTimestamp(v int64) *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgressProgressData {
	s.ConsumeTimestamp = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsResponseBodyDataListTargetProgressProgressData) Validate() error {
	return dara.Validate(s)
}
