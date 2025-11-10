// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDisasterRecoveryItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListDisasterRecoveryItemsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListDisasterRecoveryItemsResponseBody
	GetCode() *string
	SetData(v *ListDisasterRecoveryItemsResponseBodyData) *ListDisasterRecoveryItemsResponseBody
	GetData() *ListDisasterRecoveryItemsResponseBodyData
	SetDynamicCode(v string) *ListDisasterRecoveryItemsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListDisasterRecoveryItemsResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ListDisasterRecoveryItemsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDisasterRecoveryItemsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDisasterRecoveryItemsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDisasterRecoveryItemsResponseBody
	GetSuccess() *bool
}

type ListDisasterRecoveryItemsResponseBody struct {
	// The details about the access denial. This parameter is returned only if the access is denied because the Resource Access Management (RAM) user does not have the required permissions.
	//
	// example:
	//
	// xxx
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *ListDisasterRecoveryItemsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
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
	// The response code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// xxx
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C115601B-8736-5BBF-AC99-7FEAE12xxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListDisasterRecoveryItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDisasterRecoveryItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDisasterRecoveryItemsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListDisasterRecoveryItemsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDisasterRecoveryItemsResponseBody) GetData() *ListDisasterRecoveryItemsResponseBodyData {
	return s.Data
}

func (s *ListDisasterRecoveryItemsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListDisasterRecoveryItemsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListDisasterRecoveryItemsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDisasterRecoveryItemsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDisasterRecoveryItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDisasterRecoveryItemsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDisasterRecoveryItemsResponseBody) SetAccessDeniedDetail(v string) *ListDisasterRecoveryItemsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBody) SetCode(v string) *ListDisasterRecoveryItemsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBody) SetData(v *ListDisasterRecoveryItemsResponseBodyData) *ListDisasterRecoveryItemsResponseBody {
	s.Data = v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBody) SetDynamicCode(v string) *ListDisasterRecoveryItemsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBody) SetDynamicMessage(v string) *ListDisasterRecoveryItemsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBody) SetHttpStatusCode(v int32) *ListDisasterRecoveryItemsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBody) SetMessage(v string) *ListDisasterRecoveryItemsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBody) SetRequestId(v string) *ListDisasterRecoveryItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBody) SetSuccess(v bool) *ListDisasterRecoveryItemsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDisasterRecoveryItemsResponseBodyData struct {
	// The Global Replicator tasks.
	List []*ListDisasterRecoveryItemsResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The scroll ID of the request. The ID is automatically generated by the system. The result can be paginated only if this parameter is included in the pagination request.
	//
	// example:
	//
	// B13D0B07-F24B-4790-88D8-D47A38063D00
	ScrollId *string `json:"scrollId,omitempty" xml:"scrollId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 49
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListDisasterRecoveryItemsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDisasterRecoveryItemsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDisasterRecoveryItemsResponseBodyData) GetList() []*ListDisasterRecoveryItemsResponseBodyDataList {
	return s.List
}

func (s *ListDisasterRecoveryItemsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDisasterRecoveryItemsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDisasterRecoveryItemsResponseBodyData) GetScrollId() *string {
	return s.ScrollId
}

func (s *ListDisasterRecoveryItemsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDisasterRecoveryItemsResponseBodyData) SetList(v []*ListDisasterRecoveryItemsResponseBodyDataList) *ListDisasterRecoveryItemsResponseBodyData {
	s.List = v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBodyData) SetPageNumber(v int64) *ListDisasterRecoveryItemsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBodyData) SetPageSize(v int64) *ListDisasterRecoveryItemsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBodyData) SetScrollId(v string) *ListDisasterRecoveryItemsResponseBodyData {
	s.ScrollId = &v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBodyData) SetTotalCount(v int64) *ListDisasterRecoveryItemsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDisasterRecoveryItemsResponseBodyDataList struct {
	// The time when the query task was created.
	//
	// example:
	//
	// 2024-09-20 03:38:28
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The extended information.
	ExtInfo map[string]*string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// The ID of the Global Replicator task.
	//
	// example:
	//
	// 100070284
	ItemId *int64 `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// The status of the topic mapping. Valid values:
	//
	// 	- CREATING
	//
	// 	- CHANGING
	//
	// 	- RUNNING
	//
	// 	- MANUAL_STOPPED
	//
	// 	- OVERDUE_STOPPED
	//
	// example:
	//
	// RUNNING
	ItemStatus *string `json:"itemStatus,omitempty" xml:"itemStatus,omitempty"`
	// The ID of the topic mapping.
	//
	// example:
	//
	// 1300000016
	PlanId *int64 `json:"planId,omitempty" xml:"planId,omitempty"`
	// The topics involved in the topic mapping.
	Topics []*ListDisasterRecoveryItemsResponseBodyDataListTopics `json:"topics,omitempty" xml:"topics,omitempty" type:"Repeated"`
	// The time when the query task was last modified.
	//
	// example:
	//
	// 2024-10-04 02:19:44
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListDisasterRecoveryItemsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListDisasterRecoveryItemsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListDisasterRecoveryItemsResponseBodyDataList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDisasterRecoveryItemsResponseBodyDataList) GetExtInfo() map[string]*string {
	return s.ExtInfo
}

func (s *ListDisasterRecoveryItemsResponseBodyDataList) GetItemId() *int64 {
	return s.ItemId
}

func (s *ListDisasterRecoveryItemsResponseBodyDataList) GetItemStatus() *string {
	return s.ItemStatus
}

func (s *ListDisasterRecoveryItemsResponseBodyDataList) GetPlanId() *int64 {
	return s.PlanId
}

func (s *ListDisasterRecoveryItemsResponseBodyDataList) GetTopics() []*ListDisasterRecoveryItemsResponseBodyDataListTopics {
	return s.Topics
}

func (s *ListDisasterRecoveryItemsResponseBodyDataList) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListDisasterRecoveryItemsResponseBodyDataList) SetCreateTime(v string) *ListDisasterRecoveryItemsResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBodyDataList) SetExtInfo(v map[string]*string) *ListDisasterRecoveryItemsResponseBodyDataList {
	s.ExtInfo = v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBodyDataList) SetItemId(v int64) *ListDisasterRecoveryItemsResponseBodyDataList {
	s.ItemId = &v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBodyDataList) SetItemStatus(v string) *ListDisasterRecoveryItemsResponseBodyDataList {
	s.ItemStatus = &v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBodyDataList) SetPlanId(v int64) *ListDisasterRecoveryItemsResponseBodyDataList {
	s.PlanId = &v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBodyDataList) SetTopics(v []*ListDisasterRecoveryItemsResponseBodyDataListTopics) *ListDisasterRecoveryItemsResponseBodyDataList {
	s.Topics = v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBodyDataList) SetUpdateTime(v string) *ListDisasterRecoveryItemsResponseBodyDataList {
	s.UpdateTime = &v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBodyDataList) Validate() error {
	if s.Topics != nil {
		for _, item := range s.Topics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDisasterRecoveryItemsResponseBodyDataListTopics struct {
	// Deprecated
	//
	// The ID of the consumer group.
	//
	// example:
	//
	// group-test
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// The method used to deliver messages to the destination instance.
	//
	// Valid values:
	//
	// 	- Concurrently: concurrent delivery
	//
	// 	- Orderly: ordered delivery
	//
	// example:
	//
	// Concurrently
	DeliveryOrderType *string `json:"deliveryOrderType,omitempty" xml:"deliveryOrderType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-kh43w0olz0c
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The instance type.
	//
	// example:
	//
	// ALIYUN_ROCKETMQ
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// regionId
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The topic name.
	//
	// example:
	//
	// topic-test
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s ListDisasterRecoveryItemsResponseBodyDataListTopics) String() string {
	return dara.Prettify(s)
}

func (s ListDisasterRecoveryItemsResponseBodyDataListTopics) GoString() string {
	return s.String()
}

func (s *ListDisasterRecoveryItemsResponseBodyDataListTopics) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *ListDisasterRecoveryItemsResponseBodyDataListTopics) GetDeliveryOrderType() *string {
	return s.DeliveryOrderType
}

func (s *ListDisasterRecoveryItemsResponseBodyDataListTopics) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDisasterRecoveryItemsResponseBodyDataListTopics) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListDisasterRecoveryItemsResponseBodyDataListTopics) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDisasterRecoveryItemsResponseBodyDataListTopics) GetTopicName() *string {
	return s.TopicName
}

func (s *ListDisasterRecoveryItemsResponseBodyDataListTopics) SetConsumerGroupId(v string) *ListDisasterRecoveryItemsResponseBodyDataListTopics {
	s.ConsumerGroupId = &v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBodyDataListTopics) SetDeliveryOrderType(v string) *ListDisasterRecoveryItemsResponseBodyDataListTopics {
	s.DeliveryOrderType = &v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBodyDataListTopics) SetInstanceId(v string) *ListDisasterRecoveryItemsResponseBodyDataListTopics {
	s.InstanceId = &v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBodyDataListTopics) SetInstanceType(v string) *ListDisasterRecoveryItemsResponseBodyDataListTopics {
	s.InstanceType = &v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBodyDataListTopics) SetRegionId(v string) *ListDisasterRecoveryItemsResponseBodyDataListTopics {
	s.RegionId = &v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBodyDataListTopics) SetTopicName(v string) *ListDisasterRecoveryItemsResponseBodyDataListTopics {
	s.TopicName = &v
	return s
}

func (s *ListDisasterRecoveryItemsResponseBodyDataListTopics) Validate() error {
	return dara.Validate(s)
}
