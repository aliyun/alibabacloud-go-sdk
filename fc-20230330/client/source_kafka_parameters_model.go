// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSourceKafkaParameters interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerGroup(v string) *SourceKafkaParameters
	GetConsumerGroup() *string
	SetInstanceId(v string) *SourceKafkaParameters
	GetInstanceId() *string
	SetNetwork(v string) *SourceKafkaParameters
	GetNetwork() *string
	SetOffsetReset(v string) *SourceKafkaParameters
	GetOffsetReset() *string
	SetRegionId(v string) *SourceKafkaParameters
	GetRegionId() *string
	SetSecurityGroupId(v string) *SourceKafkaParameters
	GetSecurityGroupId() *string
	SetTopic(v string) *SourceKafkaParameters
	GetTopic() *string
	SetVSwitchIds(v string) *SourceKafkaParameters
	GetVSwitchIds() *string
	SetVpcId(v string) *SourceKafkaParameters
	GetVpcId() *string
}

type SourceKafkaParameters struct {
	// example:
	//
	// DEFAULT_GROUP
	ConsumerGroup *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	// example:
	//
	// r-8vb64581862c****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Default
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// example:
	//
	// latest
	OffsetReset *string `json:"OffsetReset,omitempty" xml:"OffsetReset,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// sg-bp1iv19sp1msc7zot4****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// popvip_center_robot_order
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// example:
	//
	// vsw-bp179l3llg3jjxwrq72****
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// example:
	//
	// vpc-8vblalsi0vbhizr77****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s SourceKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s SourceKafkaParameters) GoString() string {
	return s.String()
}

func (s *SourceKafkaParameters) GetConsumerGroup() *string {
	return s.ConsumerGroup
}

func (s *SourceKafkaParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SourceKafkaParameters) GetNetwork() *string {
	return s.Network
}

func (s *SourceKafkaParameters) GetOffsetReset() *string {
	return s.OffsetReset
}

func (s *SourceKafkaParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *SourceKafkaParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *SourceKafkaParameters) GetTopic() *string {
	return s.Topic
}

func (s *SourceKafkaParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *SourceKafkaParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *SourceKafkaParameters) SetConsumerGroup(v string) *SourceKafkaParameters {
	s.ConsumerGroup = &v
	return s
}

func (s *SourceKafkaParameters) SetInstanceId(v string) *SourceKafkaParameters {
	s.InstanceId = &v
	return s
}

func (s *SourceKafkaParameters) SetNetwork(v string) *SourceKafkaParameters {
	s.Network = &v
	return s
}

func (s *SourceKafkaParameters) SetOffsetReset(v string) *SourceKafkaParameters {
	s.OffsetReset = &v
	return s
}

func (s *SourceKafkaParameters) SetRegionId(v string) *SourceKafkaParameters {
	s.RegionId = &v
	return s
}

func (s *SourceKafkaParameters) SetSecurityGroupId(v string) *SourceKafkaParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *SourceKafkaParameters) SetTopic(v string) *SourceKafkaParameters {
	s.Topic = &v
	return s
}

func (s *SourceKafkaParameters) SetVSwitchIds(v string) *SourceKafkaParameters {
	s.VSwitchIds = &v
	return s
}

func (s *SourceKafkaParameters) SetVpcId(v string) *SourceKafkaParameters {
	s.VpcId = &v
	return s
}

func (s *SourceKafkaParameters) Validate() error {
	return dara.Validate(s)
}
