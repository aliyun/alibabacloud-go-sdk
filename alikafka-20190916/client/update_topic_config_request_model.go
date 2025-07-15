// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTopicConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *UpdateTopicConfigRequest
	GetConfig() *string
	SetInstanceId(v string) *UpdateTopicConfigRequest
	GetInstanceId() *string
	SetRegionId(v string) *UpdateTopicConfigRequest
	GetRegionId() *string
	SetTopic(v string) *UpdateTopicConfigRequest
	GetTopic() *string
	SetValue(v string) *UpdateTopicConfigRequest
	GetValue() *string
}

type UpdateTopicConfigRequest struct {
	// The key of the topic configuration.
	//
	// 	- For reserved instances, you can modify only the configurations of the topics that use local storage.
	//
	// 	- For serverless instances, you can modify the configurations of all topics.
	//
	// 	- Reserved instances whose topics use local storage support the following keys: retention.ms, max.message.bytes, replications, message.timestamp.type, and message.timestamp.difference.max.ms.``
	//
	// 	- Serverless instances support the following keys: retention.hours, max.message.bytes, message.timestamp.type, message.timestamp.difference.max.ms.
	//
	// This parameter is required.
	//
	// example:
	//
	// replications
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The topic name.
	//
	// This parameter is required.
	//
	// example:
	//
	// dqc_test2
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The value of the topic configuration.
	//
	// 	- Serverless instances support the following values:
	//
	//     	- `retention.hours`: the message retention period. Value type: string. Valid values: 24 to 8760.
	//
	//     	- `max.message.bytes`: the maximum size of a sent message. Value type: string. Valid values: 1048576 to 10485760.
	//
	//     	- `message.timestamp.type`: the timestamp type of a message. Valid values: CreateTime and LogAppendTime. The value CreateTime specifies the timestamp that is specified by the producer when the message is sent. The value LogAppendTime specifies the time when the broker appends the message to its log. If you do not specify this parameter, the time when the message is created on the client is used.
	//
	//     	- `message.timestamp.difference.max.ms`: the maximum positive or negative difference allowed between the timestamp when the broker receives a message and the timestamp specified in the message. If you set message.timestamp.type to CreateTime, **a message is rejected*	- if the difference in timestamp exceeds the threshold. If you set message.timestamp.type to LogAppendTime, this configuration does not take effect.
	//
	// 	- Reserved instances support the following values:
	//
	//     	- `retention.ms`: the message retention period. Value type: string. Valid values: 3600000 to 31536000000.
	//
	//     	- `max.message.bytes`: the maximum size of a sent message. Value type: string. Valid values: 1048576 to 10485760.
	//
	//     	- `replications`: the number of replicas. Value type: string. Valid values: 1 to 3.
	//
	//     	- `message.timestamp.type`: the timestamp type of a message. Valid values: CreateTime and LogAppendTime. The value CreateTime specifies the timestamp that is specified by the producer when the message is sent. The value LogAppendTime specifies the time when the broker appends the message to its log. If you do not specify this parameter, the time when the message is created on the client is used.
	//
	//     	- `message.timestamp.difference.max.ms`: the maximum positive or negative difference allowed between the timestamp when the broker receives a message and the timestamp specified in the message. If you set message.timestamp.type to CreateTime, **a message is rejected*	- if the difference in timestamp exceeds the threshold. If you set message.timestamp.type to LogAppendTime, this configuration does not take effect.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateTopicConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTopicConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateTopicConfigRequest) GetConfig() *string {
	return s.Config
}

func (s *UpdateTopicConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateTopicConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateTopicConfigRequest) GetTopic() *string {
	return s.Topic
}

func (s *UpdateTopicConfigRequest) GetValue() *string {
	return s.Value
}

func (s *UpdateTopicConfigRequest) SetConfig(v string) *UpdateTopicConfigRequest {
	s.Config = &v
	return s
}

func (s *UpdateTopicConfigRequest) SetInstanceId(v string) *UpdateTopicConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateTopicConfigRequest) SetRegionId(v string) *UpdateTopicConfigRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateTopicConfigRequest) SetTopic(v string) *UpdateTopicConfigRequest {
	s.Topic = &v
	return s
}

func (s *UpdateTopicConfigRequest) SetValue(v string) *UpdateTopicConfigRequest {
	s.Value = &v
	return s
}

func (s *UpdateTopicConfigRequest) Validate() error {
	return dara.Validate(s)
}
