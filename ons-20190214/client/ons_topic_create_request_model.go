// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTopicCreateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *OnsTopicCreateRequest
	GetInstanceId() *string
	SetMessageType(v int32) *OnsTopicCreateRequest
	GetMessageType() *int32
	SetRemark(v string) *OnsTopicCreateRequest
	GetRemark() *string
	SetTopic(v string) *OnsTopicCreateRequest
	GetTopic() *string
}

type OnsTopicCreateRequest struct {
	// The ID of the instance in which you want to create the topic.
	//
	// example:
	//
	// MQ_INST_188077086902****_BXSuW61e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of messages that you want to publish to the topic. Valid values:
	//
	// 	- **0**: normal messages
	//
	// 	- **1**: partitionally ordered messages
	//
	// 	- **2**: globally ordered messages
	//
	// 	- **4**: transactional messages
	//
	// 	- **5**: scheduled or delayed messages
	//
	// For more information about message types, see [Message types](https://help.aliyun.com/document_detail/155952.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	MessageType *int32 `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	// The description of the topic that you want to create.
	//
	// example:
	//
	// Test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The name of the topic that you want to create. The name must meet the following rules:
	//
	// 	- The name must be 3 to 64 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// 	- The topic name cannot start with CID or GID because CID and GID are reserved prefixes for group IDs.
	//
	// 	- If the ApsaraMQ for RocketMQ instance in which you want to create the topic uses a namespace, the topic name must be unique in the instance. The topic name cannot be the same as an existing topic name or a group ID in the instance. The topic name can be the same as a topic name or a group ID in another instance that uses a different namespace. For example, if Instance A and Instance B use different namespaces, a topic name in Instance A can be the same as a topic name or a group ID in Instance B.
	//
	// 	- If the instance in which you want to create the topic does not use a namespace, the topic name must be globally unique across instances and regions. The topic name cannot be the same as an existing topic name or group ID in all ApsaraMQ for RocketMQ instances that belong to your Alibaba Cloud account.
	//
	// > To check whether an instance uses a namespace, log on to the ApsaraMQ for RocketMQ console, go to the **Instance Details*	- page, and then view the value of the Namespace field in the **Basic Information*	- section.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsTopicCreateRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicCreateRequest) GoString() string {
	return s.String()
}

func (s *OnsTopicCreateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsTopicCreateRequest) GetMessageType() *int32 {
	return s.MessageType
}

func (s *OnsTopicCreateRequest) GetRemark() *string {
	return s.Remark
}

func (s *OnsTopicCreateRequest) GetTopic() *string {
	return s.Topic
}

func (s *OnsTopicCreateRequest) SetInstanceId(v string) *OnsTopicCreateRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsTopicCreateRequest) SetMessageType(v int32) *OnsTopicCreateRequest {
	s.MessageType = &v
	return s
}

func (s *OnsTopicCreateRequest) SetRemark(v string) *OnsTopicCreateRequest {
	s.Remark = &v
	return s
}

func (s *OnsTopicCreateRequest) SetTopic(v string) *OnsTopicCreateRequest {
	s.Topic = &v
	return s
}

func (s *OnsTopicCreateRequest) Validate() error {
	return dara.Validate(s)
}
