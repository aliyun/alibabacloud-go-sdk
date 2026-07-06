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
	// - You can modify the configurations only for topics that use the local storage engine on reserved instances. You cannot modify the configurations for topics that use the cloud storage engine.
	//
	// - You can modify the configurations of topics for Serverless instances.
	//
	// - For `local topics` on reserved instances, the supported keys are \\`retention.ms\\`, \\`max.message.bytes\\`, \\`message.timestamp.type\\`, and \\`message.timestamp.difference.max.ms\\`.
	//
	// - For Serverless instances, the supported keys are \\`retention.hours\\`, \\`max.message.bytes\\`, \\`message.timestamp.type\\`, and \\`message.timestamp.difference.max.ms\\`.
	//
	// This parameter is required.
	//
	// example:
	//
	// retention.ms
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
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
	// - The following configurations are supported for Serverless instances:
	//
	//   - `retention.hours` specifies the message retention period. The value must be a string. The value must be in the range of 24 to 8,760.
	//
	//   - `max.message.bytes` specifies the maximum message size. The value must be a string. The value must be in the range of 1,048,576 to 10,485,760.
	//
	//   - `message.timestamp.type` specifies the message timestamp type. \\`CreateTime\\` indicates the timestamp that is specified by the producer when the message is sent. If no timestamp is specified, the time when the message is created on the client is used. \\`LogAppendTime\\` indicates the time when the message is stored on the server. Valid values: \\`CreateTime\\` and \\`LogAppendTime\\`.
	//
	//   - `message.timestamp.difference.max.ms` specifies the maximum allowed difference between the timestamp of the server that receives the message and the timestamp specified in the message. If \\`message.timestamp.type\\` is set to \\`CreateTime\\` and the time difference exceeds this threshold, **the message is rejected**. This configuration does not take effect if \\`message.timestamp.type\\` is set to \\`LogAppendTime\\`.
	//
	// - The following configurations are supported for reserved instances:
	//
	//   - `retention.ms` specifies the message retention period. The value must be a string. The value must be in the range of 3,600,000 to 31,536,000,000.
	//
	//   - `max.message.bytes` specifies the maximum message size. The value must be a string. The value must be in the range of 1,048,576 to 10,485,760.
	//
	//   - `message.timestamp.type` specifies the message timestamp type. \\`CreateTime\\` indicates the timestamp that is specified by the producer when the message is sent. If no timestamp is specified, the time when the message is created on the client is used. \\`LogAppendTime\\` indicates the time when the message is stored on the server. Valid values: \\`CreateTime\\` and \\`LogAppendTime\\`.
	//
	//   - `message.timestamp.difference.max.ms` specifies the maximum allowed difference between the timestamp of the server that receives the message and the timestamp specified in the message. If \\`message.timestamp.type\\` is set to \\`CreateTime\\` and the time difference exceeds this threshold, **the message is rejected**. This configuration does not take effect if \\`message.timestamp.type\\` is set to \\`LogAppendTime\\`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3600000
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
