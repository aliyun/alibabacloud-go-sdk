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
	// The cleanup policy configured when the storage engine of the topic is set to local storage. Valid values:
	//
	// - false: delete cleanup policy.
	//
	// - true: compact cleanup policy.
	//
	// example:
	//
	// false
	CompactTopic *bool `json:"CompactTopic,omitempty" xml:"CompactTopic,omitempty"`
	// The supplementary configuration.
	//
	// - Must be in JSON format.
	//
	//
	// - This parameter takes effect only when **LocalTopic*	- is set to **true**.
	//
	// - Supported configurations for reserved instances:
	//
	//   -   **retention.ms*	- (message retention period): ranges from 3600000 to 31536000000 milliseconds.
	//
	//   - **max.message.bytes*	- (maximum message size): ranges from 1048576 to 10485760 bytes.
	//
	//   - **message.timestamp.type**: specifies the type of message timestamp. CreateTime indicates the timestamp specified by the producer when sending a message. If not specified, it is the message creation time on the client. LogAppendTime indicates the time when the message is written to disk on the server. Valid values: CreateTime or LogAppendTime. Default value: CreateTime. We recommend LogAppendTime.
	//
	//  - Supported configurations for Serverless instances:
	//
	//    -  **retention.hours*	- (message retention period): value type is String. Valid values: 24 to 8760.
	//
	//    -  **max.message.bytes*	- (maximum message size): value type is String. Valid values: 1048576 to 10485760.
	//
	//    -  **message.timestamp.type*	- (type of message timestamp): CreateTime indicates the timestamp specified by the producer when sending a message. If not specified, it is the message creation time on the client. LogAppendTime indicates the time when the message is written to disk on the server. Valid values: CreateTime or LogAppendTime. Default value: CreateTime. We recommend LogAppendTime.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// {"retention.ms": "3600000"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The instance ID.
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
	// The minimum number of in-sync replicas (ISR).
	//
	// - This parameter takes effect only when **LocalTopic*	- is set to **true**.
	//
	// - The value must be less than the number of topic replicas.
	//
	// - The number of in-sync replicas ranges from 1 to 3.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 1
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitempty" xml:"MinInsyncReplicas,omitempty"`
	// The number of partitions for the topic.
	//
	// - The number of partitions ranges from 1 to 360.
	//
	// - The console provides different configuration suggestions based on the instance edition. Configure the number of partitions based on the console suggestions to reduce the risk of data skew.
	//
	// Default value:
	//
	// - Reserved instances: 12
	//
	// - Serverless instances: 3
	//
	// example:
	//
	// 12
	PartitionNum *string `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// The region ID of the instance to which the topic belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the topic.
	//
	// - Can contain only letters, digits, underscores (_), and hyphens (-).
	//
	// - Must be 3 to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_topic_test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The number of replicas for the topic.
	//
	// - This parameter takes effect only when **LocalTopic*	- is set to **true**.
	//
	// - The number of replicas ranges from 1 to 3.
	//
	// > If the number of replicas is set to **1**, data loss may occur. Set this parameter with caution.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 3
	ReplicationFactor *int64 `json:"ReplicationFactor,omitempty" xml:"ReplicationFactor,omitempty"`
	// The tag list.
	Tag []*CreateTopicRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the topic.
	//
	// - Reserved instances:
	//
	// Supports uppercase and lowercase letters, digits, underscores (_), hyphens (-), and periods (.). The name must be 3 to 64 characters in length.
	//
	// - Serverless instances:
	//
	// Supports uppercase and lowercase letters, digits, underscores (_), hyphens (-), and periods (.). The name must be 1 to 249 characters in length.
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
	// - N ranges from 1 to 20.
	//
	// - If this parameter is left empty, all tag keys are matched.
	//
	// - The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`, or contain `http://` or `https://`.
	//
	// This parameter is required.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// - N ranges from 1 to 20.
	//
	// - This parameter can be left empty.
	//
	// - The tag value can be up to 128 characters in length and cannot start with aliyun or acs:, or contain http:// or https://.
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
