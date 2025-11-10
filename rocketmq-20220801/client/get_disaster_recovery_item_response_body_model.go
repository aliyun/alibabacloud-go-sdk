// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDisasterRecoveryItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetDisasterRecoveryItemResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetDisasterRecoveryItemResponseBody
	GetCode() *string
	SetData(v *GetDisasterRecoveryItemResponseBodyData) *GetDisasterRecoveryItemResponseBody
	GetData() *GetDisasterRecoveryItemResponseBodyData
	SetDynamicCode(v string) *GetDisasterRecoveryItemResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetDisasterRecoveryItemResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *GetDisasterRecoveryItemResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDisasterRecoveryItemResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDisasterRecoveryItemResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDisasterRecoveryItemResponseBody
	GetSuccess() *bool
}

type GetDisasterRecoveryItemResponseBody struct {
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
	// The data returned.
	Data *GetDisasterRecoveryItemResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// The HTTP status code.
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
	// Request ID
	//
	// example:
	//
	// AF9A8B10-C426-530F-A0DD-96320B39****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDisasterRecoveryItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDisasterRecoveryItemResponseBody) GoString() string {
	return s.String()
}

func (s *GetDisasterRecoveryItemResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetDisasterRecoveryItemResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDisasterRecoveryItemResponseBody) GetData() *GetDisasterRecoveryItemResponseBodyData {
	return s.Data
}

func (s *GetDisasterRecoveryItemResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetDisasterRecoveryItemResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetDisasterRecoveryItemResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDisasterRecoveryItemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDisasterRecoveryItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDisasterRecoveryItemResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDisasterRecoveryItemResponseBody) SetAccessDeniedDetail(v string) *GetDisasterRecoveryItemResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetDisasterRecoveryItemResponseBody) SetCode(v string) *GetDisasterRecoveryItemResponseBody {
	s.Code = &v
	return s
}

func (s *GetDisasterRecoveryItemResponseBody) SetData(v *GetDisasterRecoveryItemResponseBodyData) *GetDisasterRecoveryItemResponseBody {
	s.Data = v
	return s
}

func (s *GetDisasterRecoveryItemResponseBody) SetDynamicCode(v string) *GetDisasterRecoveryItemResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetDisasterRecoveryItemResponseBody) SetDynamicMessage(v string) *GetDisasterRecoveryItemResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetDisasterRecoveryItemResponseBody) SetHttpStatusCode(v int32) *GetDisasterRecoveryItemResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDisasterRecoveryItemResponseBody) SetMessage(v string) *GetDisasterRecoveryItemResponseBody {
	s.Message = &v
	return s
}

func (s *GetDisasterRecoveryItemResponseBody) SetRequestId(v string) *GetDisasterRecoveryItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDisasterRecoveryItemResponseBody) SetSuccess(v bool) *GetDisasterRecoveryItemResponseBody {
	s.Success = &v
	return s
}

func (s *GetDisasterRecoveryItemResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDisasterRecoveryItemResponseBodyData struct {
	// The time when the topic mapping task was created.
	//
	// example:
	//
	// 2024-06-24 02:57:31
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Additional Information
	ExtInfo map[string]*string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// The ID of the topic mapping
	//
	// example:
	//
	// 100070284
	ItemId *int64 `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// The topic mapping task status.
	//
	// example:
	//
	// RUNNING
	ItemStatus *string `json:"itemStatus,omitempty" xml:"itemStatus,omitempty"`
	// The ID of the global message backup plan.
	//
	// example:
	//
	// 1300000016
	PlanId *int64 `json:"planId,omitempty" xml:"planId,omitempty"`
	// Topics included in the backup mapping
	Topics []*GetDisasterRecoveryItemResponseBodyDataTopics `json:"topics,omitempty" xml:"topics,omitempty" type:"Repeated"`
	// The time when the topic mapping task was last updated.
	//
	// example:
	//
	// 2024-09-26 02:13:10
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetDisasterRecoveryItemResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDisasterRecoveryItemResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDisasterRecoveryItemResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetDisasterRecoveryItemResponseBodyData) GetExtInfo() map[string]*string {
	return s.ExtInfo
}

func (s *GetDisasterRecoveryItemResponseBodyData) GetItemId() *int64 {
	return s.ItemId
}

func (s *GetDisasterRecoveryItemResponseBodyData) GetItemStatus() *string {
	return s.ItemStatus
}

func (s *GetDisasterRecoveryItemResponseBodyData) GetPlanId() *int64 {
	return s.PlanId
}

func (s *GetDisasterRecoveryItemResponseBodyData) GetTopics() []*GetDisasterRecoveryItemResponseBodyDataTopics {
	return s.Topics
}

func (s *GetDisasterRecoveryItemResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetDisasterRecoveryItemResponseBodyData) SetCreateTime(v string) *GetDisasterRecoveryItemResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetDisasterRecoveryItemResponseBodyData) SetExtInfo(v map[string]*string) *GetDisasterRecoveryItemResponseBodyData {
	s.ExtInfo = v
	return s
}

func (s *GetDisasterRecoveryItemResponseBodyData) SetItemId(v int64) *GetDisasterRecoveryItemResponseBodyData {
	s.ItemId = &v
	return s
}

func (s *GetDisasterRecoveryItemResponseBodyData) SetItemStatus(v string) *GetDisasterRecoveryItemResponseBodyData {
	s.ItemStatus = &v
	return s
}

func (s *GetDisasterRecoveryItemResponseBodyData) SetPlanId(v int64) *GetDisasterRecoveryItemResponseBodyData {
	s.PlanId = &v
	return s
}

func (s *GetDisasterRecoveryItemResponseBodyData) SetTopics(v []*GetDisasterRecoveryItemResponseBodyDataTopics) *GetDisasterRecoveryItemResponseBodyData {
	s.Topics = v
	return s
}

func (s *GetDisasterRecoveryItemResponseBodyData) SetUpdateTime(v string) *GetDisasterRecoveryItemResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetDisasterRecoveryItemResponseBodyData) Validate() error {
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

type GetDisasterRecoveryItemResponseBodyDataTopics struct {
	// Deprecated
	//
	// The consumer group ID.
	//
	// example:
	//
	// xxx_reserve_group
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// The order in which messages are delivered to the target instance. The parameter values ​​are as follows:
	//
	//   - Concurrently: concurrent delivery
	//
	//   - Orderly: sequential delivery
	//
	// example:
	//
	// Concurrently
	DeliveryOrderType *string `json:"deliveryOrderType,omitempty" xml:"deliveryOrderType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-wwo3xxx
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
	// order_push_xxx
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s GetDisasterRecoveryItemResponseBodyDataTopics) String() string {
	return dara.Prettify(s)
}

func (s GetDisasterRecoveryItemResponseBodyDataTopics) GoString() string {
	return s.String()
}

func (s *GetDisasterRecoveryItemResponseBodyDataTopics) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *GetDisasterRecoveryItemResponseBodyDataTopics) GetDeliveryOrderType() *string {
	return s.DeliveryOrderType
}

func (s *GetDisasterRecoveryItemResponseBodyDataTopics) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDisasterRecoveryItemResponseBodyDataTopics) GetInstanceType() *string {
	return s.InstanceType
}

func (s *GetDisasterRecoveryItemResponseBodyDataTopics) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDisasterRecoveryItemResponseBodyDataTopics) GetTopicName() *string {
	return s.TopicName
}

func (s *GetDisasterRecoveryItemResponseBodyDataTopics) SetConsumerGroupId(v string) *GetDisasterRecoveryItemResponseBodyDataTopics {
	s.ConsumerGroupId = &v
	return s
}

func (s *GetDisasterRecoveryItemResponseBodyDataTopics) SetDeliveryOrderType(v string) *GetDisasterRecoveryItemResponseBodyDataTopics {
	s.DeliveryOrderType = &v
	return s
}

func (s *GetDisasterRecoveryItemResponseBodyDataTopics) SetInstanceId(v string) *GetDisasterRecoveryItemResponseBodyDataTopics {
	s.InstanceId = &v
	return s
}

func (s *GetDisasterRecoveryItemResponseBodyDataTopics) SetInstanceType(v string) *GetDisasterRecoveryItemResponseBodyDataTopics {
	s.InstanceType = &v
	return s
}

func (s *GetDisasterRecoveryItemResponseBodyDataTopics) SetRegionId(v string) *GetDisasterRecoveryItemResponseBodyDataTopics {
	s.RegionId = &v
	return s
}

func (s *GetDisasterRecoveryItemResponseBodyDataTopics) SetTopicName(v string) *GetDisasterRecoveryItemResponseBodyDataTopics {
	s.TopicName = &v
	return s
}

func (s *GetDisasterRecoveryItemResponseBodyDataTopics) Validate() error {
	return dara.Validate(s)
}
