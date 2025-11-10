// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDisasterRecoveryItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTopics(v []*UpdateDisasterRecoveryItemRequestTopics) *UpdateDisasterRecoveryItemRequest
	GetTopics() []*UpdateDisasterRecoveryItemRequestTopics
}

type UpdateDisasterRecoveryItemRequest struct {
	// The topics involved in the topic mapping.
	Topics []*UpdateDisasterRecoveryItemRequestTopics `json:"topics,omitempty" xml:"topics,omitempty" type:"Repeated"`
}

func (s UpdateDisasterRecoveryItemRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDisasterRecoveryItemRequest) GoString() string {
	return s.String()
}

func (s *UpdateDisasterRecoveryItemRequest) GetTopics() []*UpdateDisasterRecoveryItemRequestTopics {
	return s.Topics
}

func (s *UpdateDisasterRecoveryItemRequest) SetTopics(v []*UpdateDisasterRecoveryItemRequestTopics) *UpdateDisasterRecoveryItemRequest {
	s.Topics = v
	return s
}

func (s *UpdateDisasterRecoveryItemRequest) Validate() error {
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

type UpdateDisasterRecoveryItemRequestTopics struct {
	// Deprecated
	//
	// The ID of the consumer group. If you use the two-way backup mode, you must specify this parameter.
	//
	// example:
	//
	// GID_xxx
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
	// The instance ID. If you set instanceType to EXTERNAL_ROCKETMQ, the system automatically generates an ID for the instance. You can obtain the ID by querying the global message backup plan.
	//
	// example:
	//
	// rmq-cn-em93y94xxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The instance type. Valid values:
	//
	// 	- ALIYUN_ROCKETMQ: ApsaraMQ for RocketMQ instance
	//
	// 	- EXTERNAL_ROCKETMQ: open source RocketMQ cluster
	//
	// example:
	//
	// ALIYUN_ROCKETMQ
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The topic name. You must specify this parameter.
	//
	// example:
	//
	// Topic_xxx
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s UpdateDisasterRecoveryItemRequestTopics) String() string {
	return dara.Prettify(s)
}

func (s UpdateDisasterRecoveryItemRequestTopics) GoString() string {
	return s.String()
}

func (s *UpdateDisasterRecoveryItemRequestTopics) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *UpdateDisasterRecoveryItemRequestTopics) GetDeliveryOrderType() *string {
	return s.DeliveryOrderType
}

func (s *UpdateDisasterRecoveryItemRequestTopics) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateDisasterRecoveryItemRequestTopics) GetInstanceType() *string {
	return s.InstanceType
}

func (s *UpdateDisasterRecoveryItemRequestTopics) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDisasterRecoveryItemRequestTopics) GetTopicName() *string {
	return s.TopicName
}

func (s *UpdateDisasterRecoveryItemRequestTopics) SetConsumerGroupId(v string) *UpdateDisasterRecoveryItemRequestTopics {
	s.ConsumerGroupId = &v
	return s
}

func (s *UpdateDisasterRecoveryItemRequestTopics) SetDeliveryOrderType(v string) *UpdateDisasterRecoveryItemRequestTopics {
	s.DeliveryOrderType = &v
	return s
}

func (s *UpdateDisasterRecoveryItemRequestTopics) SetInstanceId(v string) *UpdateDisasterRecoveryItemRequestTopics {
	s.InstanceId = &v
	return s
}

func (s *UpdateDisasterRecoveryItemRequestTopics) SetInstanceType(v string) *UpdateDisasterRecoveryItemRequestTopics {
	s.InstanceType = &v
	return s
}

func (s *UpdateDisasterRecoveryItemRequestTopics) SetRegionId(v string) *UpdateDisasterRecoveryItemRequestTopics {
	s.RegionId = &v
	return s
}

func (s *UpdateDisasterRecoveryItemRequestTopics) SetTopicName(v string) *UpdateDisasterRecoveryItemRequestTopics {
	s.TopicName = &v
	return s
}

func (s *UpdateDisasterRecoveryItemRequestTopics) Validate() error {
	return dara.Validate(s)
}
