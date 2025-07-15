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
	// The log cleanup policy that is used for the topic. This parameter is available only when LocalTopic is set to true. Valid values:
	//
	// 	- false: The topic uses the default log cleanup policy.
	//
	// 	- true: The topic uses the log compaction policy.
	//
	// example:
	//
	// false
	CompactTopic *bool `json:"CompactTopic,omitempty" xml:"CompactTopic,omitempty"`
	// The additional configuration.
	//
	// 	- The value must be in JSON format.
	//
	// 	- Set Key to **replications**. This value specifies the number of replicas of the topic. The value must be an integer that ranges from 1 to 3.
	//
	// 	- You can configure this parameter only if you set **LocalTopic*	- to **true*	- or specify **Open Source Edition (Local Disk)*	- as the instance edition.****
	//
	// >  If you specify replications in this parameter, **ReplicationFactor*	- does not take effect.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// {"replications": 3}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-mp919o4v****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of storage that the topic uses. Valid values:
	//
	// 	- false: The topic uses cloud storage.
	//
	// 	- true: The topic uses local storage.
	//
	// example:
	//
	// false
	LocalTopic *bool `json:"LocalTopic,omitempty" xml:"LocalTopic,omitempty"`
	// The minimum number of in-sync replicas (ISRs).
	//
	// 	- This parameter is available only when **LocalTopic*	- is set to **true**, or the instance is of the **Open Source Edition (Local Disk)**.****
	//
	// 	- The value of this parameter must be smaller than the value of ReplicationFactor.
	//
	// 	- Valid values: 1 to 3.
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
	// 	- Valid values: 1 to 360.
	//
	// 	- In the ApsaraMQ for Kafka console, you can view the number of partitions that the system recommends based on the specifications of the instance. We recommend that you specify the number that is recommended by the system as the value of this parameter to reduce the risk of data skew.
	//
	// Default values:
	//
	// 	- ApsaraMQ for Kafka V2 instance: 12
	//
	// 	- ApsaraMQ for Kafka V3 instance: 3
	//
	// example:
	//
	// 12
	PartitionNum *string `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// The region ID of the instance in which you want to create a topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the topic.
	//
	// 	- The description can contain only letters, digits, hyphens (-), and underscores (_).
	//
	// 	- The description must be 3 to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_topic_test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The number of replicas for the topic.
	//
	// 	- This parameter is available only when **LocalTopic*	- is set to **true**, or the instance is of the **Open Source Edition (Local Disk)**.****
	//
	// 	- Valid values: 1 to 3.
	//
	// > If you set this parameter to **1**, data loss may occur. Exercise caution when you configure this parameter.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 3
	ReplicationFactor *int64 `json:"ReplicationFactor,omitempty" xml:"ReplicationFactor,omitempty"`
	// The tags that you want to add to the topic.
	Tag []*CreateTopicRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The topic name.
	//
	// 	- The name can contain only letters, digits, hyphens (-), and underscores (_).
	//
	// 	- The name must be 3 to 64 characters in length. If the name that you specify contains more than 64 characters, the system automatically truncates the name.
	//
	// 	- After a topic is created, you cannot change the name of the topic.
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
	return dara.Validate(s)
}

type CreateTopicRequestTag struct {
	// The tag key.
	//
	// 	- If you do not specify this parameter, the keys of all tags are matched.
	//
	// 	- The tag key must be 1 to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	//
	// This parameter is required.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// 	- You can leave this parameter empty.
	//
	// 	- The tag value must be 1 to 128 characters in length and cannot contain http:// or https://. The tag value cannot start with aliyun or acs:.
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
