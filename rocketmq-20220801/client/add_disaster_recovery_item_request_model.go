// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDisasterRecoveryItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTopics(v []*AddDisasterRecoveryItemRequestTopics) *AddDisasterRecoveryItemRequest
	GetTopics() []*AddDisasterRecoveryItemRequestTopics
}

type AddDisasterRecoveryItemRequest struct {
	// Topics included in the backup mapping. Required.
	Topics []*AddDisasterRecoveryItemRequestTopics `json:"topics,omitempty" xml:"topics,omitempty" type:"Repeated"`
}

func (s AddDisasterRecoveryItemRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDisasterRecoveryItemRequest) GoString() string {
	return s.String()
}

func (s *AddDisasterRecoveryItemRequest) GetTopics() []*AddDisasterRecoveryItemRequestTopics {
	return s.Topics
}

func (s *AddDisasterRecoveryItemRequest) SetTopics(v []*AddDisasterRecoveryItemRequestTopics) *AddDisasterRecoveryItemRequest {
	s.Topics = v
	return s
}

func (s *AddDisasterRecoveryItemRequest) Validate() error {
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

type AddDisasterRecoveryItemRequestTopics struct {
	// Deprecated
	//
	// Consumer group ID, required for ACTIVE_ACTIVE bidirectional backup
	//
	// example:
	//
	// GID_xxx
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
	// Instance ID, an instance ID will be automatically generated when `instanceType` is `EXTERNAL_ROCKETMQ`, and it can be obtained by querying the backup plan.
	//
	// example:
	//
	// rmq-cn-em93y94xxx
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
	// Region ID
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Disaster recovery topic name, required
	//
	// example:
	//
	// Topic_xxx
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s AddDisasterRecoveryItemRequestTopics) String() string {
	return dara.Prettify(s)
}

func (s AddDisasterRecoveryItemRequestTopics) GoString() string {
	return s.String()
}

func (s *AddDisasterRecoveryItemRequestTopics) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *AddDisasterRecoveryItemRequestTopics) GetDeliveryOrderType() *string {
	return s.DeliveryOrderType
}

func (s *AddDisasterRecoveryItemRequestTopics) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddDisasterRecoveryItemRequestTopics) GetInstanceType() *string {
	return s.InstanceType
}

func (s *AddDisasterRecoveryItemRequestTopics) GetRegionId() *string {
	return s.RegionId
}

func (s *AddDisasterRecoveryItemRequestTopics) GetTopicName() *string {
	return s.TopicName
}

func (s *AddDisasterRecoveryItemRequestTopics) SetConsumerGroupId(v string) *AddDisasterRecoveryItemRequestTopics {
	s.ConsumerGroupId = &v
	return s
}

func (s *AddDisasterRecoveryItemRequestTopics) SetDeliveryOrderType(v string) *AddDisasterRecoveryItemRequestTopics {
	s.DeliveryOrderType = &v
	return s
}

func (s *AddDisasterRecoveryItemRequestTopics) SetInstanceId(v string) *AddDisasterRecoveryItemRequestTopics {
	s.InstanceId = &v
	return s
}

func (s *AddDisasterRecoveryItemRequestTopics) SetInstanceType(v string) *AddDisasterRecoveryItemRequestTopics {
	s.InstanceType = &v
	return s
}

func (s *AddDisasterRecoveryItemRequestTopics) SetRegionId(v string) *AddDisasterRecoveryItemRequestTopics {
	s.RegionId = &v
	return s
}

func (s *AddDisasterRecoveryItemRequestTopics) SetTopicName(v string) *AddDisasterRecoveryItemRequestTopics {
	s.TopicName = &v
	return s
}

func (s *AddDisasterRecoveryItemRequestTopics) Validate() error {
	return dara.Validate(s)
}
