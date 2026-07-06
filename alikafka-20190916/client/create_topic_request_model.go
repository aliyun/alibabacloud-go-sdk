// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTopicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompactTopic(v bool) *CreateTopicRequest
	GetCompactTopic() *bool
	SetConfig(v string) *CreateTopicRequest
	GetConfig() *string
	SetInstanceId(v string) *CreateTopicRequest
	GetInstanceId() *string
	SetLocalTopic(v bool) *CreateTopicRequest
	GetLocalTopic() *bool
	SetMinInsyncReplicas(v int64) *CreateTopicRequest
	GetMinInsyncReplicas() *int64
	SetPartitionNum(v string) *CreateTopicRequest
	GetPartitionNum() *string
	SetRegionId(v string) *CreateTopicRequest
	GetRegionId() *string
	SetRemark(v string) *CreateTopicRequest
	GetRemark() *string
	SetReplicationFactor(v int64) *CreateTopicRequest
	GetReplicationFactor() *int64
	SetTag(v []*CreateTopicRequestTag) *CreateTopicRequest
	GetTag() []*CreateTopicRequestTag
	SetTopic(v string) *CreateTopicRequest
	GetTopic() *string
}

type CreateTopicRequest struct {
	// The cleanup policy for the topic. This parameter is available only if the storage engine of the topic is local storage. Valid values:
	//
	// - false: The delete cleanup policy.
	//
	// - true: The compact cleanup policy.
	//
	// example:
	//
	// false
	CompactTopic *bool `json:"CompactTopic,omitempty" xml:"CompactTopic,omitempty"`
	// The advanced configurations of the topic.
	//
	// - Configure this parameter in the JSON format.
	//
	// - This parameter is available only if **LocalTopic*	- is set to **true**.
	//
	// - The following configurations are supported for reserved instances:
	//
	//   - **retention.ms**: The message retention period. The value must be an integer from 3,600,000 to 31,536,000,000. Unit: milliseconds.
	//
	//   - **max.message.bytes**: The maximum size of a message that can be sent. The value must be an integer from 1,048,576 to 10,485,760. Unit: bytes.
	//
	//   - message.timestamp.type: The timestamp type of a message. Valid values: CreateTime or LogAppendTime. CreateTime indicates that the message timestamp is the time when the producer creates the message. If you do not specify a timestamp, the client time is used. LogAppendTime indicates that the message timestamp is the time when the server stores the message. The default value is CreateTime. We recommend that you set this parameter to **LogAppendTime**.
	//
	// - The following configurations are supported for Serverless instances:
	//
	//   - **retention.hours**: The message retention period. The value is of the string type. The value must be an integer from 24 to 8,760.
	//
	//   - **max.message.bytes**: The maximum size of a message that can be sent. The value is of the string type. The value must be an integer from 1,048,576 to 10,485,760.
	//
	//   - message.timestamp.type: The timestamp type of a message. Valid values: CreateTime or LogAppendTime. CreateTime indicates that the message timestamp is the time when the producer creates the message. If you do not specify a timestamp, the client time is used. LogAppendTime indicates that the message timestamp is the time when the server stores the message. The default value is CreateTime. We recommend that you set this parameter to **LogAppendTime**.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// {"retention.ms": "3600000"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-mp919o4v****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The storage engine of the topic. Valid values:
	//
	// - false: cloud storage.
	//
	// - true: local storage.
	//
	// example:
	//
	// false
	LocalTopic *bool `json:"LocalTopic,omitempty" xml:"LocalTopic,omitempty"`
	// The minimum number of in-sync replicas (ISRs).
	//
	// - This parameter is available only if **LocalTopic*	- is set to **true**.
	//
	// - The value of this parameter must be smaller than the number of replicas for the topic.
	//
	// - The value must be an integer from 1 to 3.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 1
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitempty" xml:"MinInsyncReplicas,omitempty"`
	// The number of partitions in the topic.
	//
	// - The value must be an integer from 1 to 360.
	//
	// - The console suggests a number of partitions based on the instance type. Follow the suggestion to reduce the risk of data skew.
	//
	// Default value:
	//
	// - Reserved instance: 12
	//
	// - Serverless instance: 3
	//
	// example:
	//
	// 12
	PartitionNum *string `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// The ID of the region where the instance that contains the topic is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks on the topic.
	//
	// - The remarks can contain only letters, digits, underscores (_), and hyphens (-).
	//
	// - The remarks must be 3 to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_topic_test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The number of replicas for the topic.
	//
	// - This parameter is available only if **LocalTopic*	- is set to **true**.
	//
	// - The value must be an integer from 1 to 3.
	//
	// > If you set the number of replicas to **1**, you may lose data. Set this parameter with caution.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 3
	ReplicationFactor *int64 `json:"ReplicationFactor,omitempty" xml:"ReplicationFactor,omitempty"`
	// The list of tags.
	Tag []*CreateTopicRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the topic.
	//
	// - Reserved instance: The name can contain uppercase letters, lowercase letters, digits, underscores (_), hyphens (-), and periods (.). The name must be 3 to 64 characters in length.
	//
	// - Serverless instance: The name can contain uppercase letters, lowercase letters, digits, underscores (_), hyphens (-), and periods (.). The name must be 1 to 249 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_topic_test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s CreateTopicRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTopicRequest) GoString() string {
	return s.String()
}

func (s *CreateTopicRequest) GetCompactTopic() *bool {
	return s.CompactTopic
}

func (s *CreateTopicRequest) GetConfig() *string {
	return s.Config
}

func (s *CreateTopicRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateTopicRequest) GetLocalTopic() *bool {
	return s.LocalTopic
}

func (s *CreateTopicRequest) GetMinInsyncReplicas() *int64 {
	return s.MinInsyncReplicas
}

func (s *CreateTopicRequest) GetPartitionNum() *string {
	return s.PartitionNum
}

func (s *CreateTopicRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTopicRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateTopicRequest) GetReplicationFactor() *int64 {
	return s.ReplicationFactor
}

func (s *CreateTopicRequest) GetTag() []*CreateTopicRequestTag {
	return s.Tag
}

func (s *CreateTopicRequest) GetTopic() *string {
	return s.Topic
}

func (s *CreateTopicRequest) SetCompactTopic(v bool) *CreateTopicRequest {
	s.CompactTopic = &v
	return s
}

func (s *CreateTopicRequest) SetConfig(v string) *CreateTopicRequest {
	s.Config = &v
	return s
}

func (s *CreateTopicRequest) SetInstanceId(v string) *CreateTopicRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateTopicRequest) SetLocalTopic(v bool) *CreateTopicRequest {
	s.LocalTopic = &v
	return s
}

func (s *CreateTopicRequest) SetMinInsyncReplicas(v int64) *CreateTopicRequest {
	s.MinInsyncReplicas = &v
	return s
}

func (s *CreateTopicRequest) SetPartitionNum(v string) *CreateTopicRequest {
	s.PartitionNum = &v
	return s
}

func (s *CreateTopicRequest) SetRegionId(v string) *CreateTopicRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTopicRequest) SetRemark(v string) *CreateTopicRequest {
	s.Remark = &v
	return s
}

func (s *CreateTopicRequest) SetReplicationFactor(v int64) *CreateTopicRequest {
	s.ReplicationFactor = &v
	return s
}

func (s *CreateTopicRequest) SetTag(v []*CreateTopicRequestTag) *CreateTopicRequest {
	s.Tag = v
	return s
}

func (s *CreateTopicRequest) SetTopic(v string) *CreateTopicRequest {
	s.Topic = &v
	return s
}

func (s *CreateTopicRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateTopicRequestTag struct {
	// The tag key of the resource.
	//
	// - N specifies the number of the tag. The value of N must be an integer from 1 to 20.
	//
	// - If this parameter is left empty, all tag keys are matched.
	//
	// - The tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// This parameter is required.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// - N specifies the number of the tag. The value of N must be an integer from 1 to 20.
	//
	// - The tag value can be empty.
	//
	// - The tag value can be up to 128 characters in length. It cannot start with aliyun or acs:, and cannot contain http\\:// or https\\://.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTopicRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateTopicRequestTag) GoString() string {
	return s.String()
}

func (s *CreateTopicRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateTopicRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateTopicRequestTag) SetKey(v string) *CreateTopicRequestTag {
	s.Key = &v
	return s
}

func (s *CreateTopicRequestTag) SetValue(v string) *CreateTopicRequestTag {
	s.Value = &v
	return s
}

func (s *CreateTopicRequestTag) Validate() error {
	return dara.Validate(s)
}
