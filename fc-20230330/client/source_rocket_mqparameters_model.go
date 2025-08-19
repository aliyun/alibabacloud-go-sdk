// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSourceRocketMQParameters interface {
	dara.Model
	String() string
	GoString() string
	SetAuthType(v string) *SourceRocketMQParameters
	GetAuthType() *string
	SetFilterType(v string) *SourceRocketMQParameters
	GetFilterType() *string
	SetGroupID(v string) *SourceRocketMQParameters
	GetGroupID() *string
	SetInstanceEndpoint(v string) *SourceRocketMQParameters
	GetInstanceEndpoint() *string
	SetInstanceId(v string) *SourceRocketMQParameters
	GetInstanceId() *string
	SetInstanceNetwork(v string) *SourceRocketMQParameters
	GetInstanceNetwork() *string
	SetInstancePassword(v string) *SourceRocketMQParameters
	GetInstancePassword() *string
	SetInstanceSecurityGroupId(v string) *SourceRocketMQParameters
	GetInstanceSecurityGroupId() *string
	SetInstanceType(v string) *SourceRocketMQParameters
	GetInstanceType() *string
	SetInstanceUsername(v string) *SourceRocketMQParameters
	GetInstanceUsername() *string
	SetInstanceVSwitchIds(v string) *SourceRocketMQParameters
	GetInstanceVSwitchIds() *string
	SetInstanceVpcId(v string) *SourceRocketMQParameters
	GetInstanceVpcId() *string
	SetOffset(v string) *SourceRocketMQParameters
	GetOffset() *string
	SetRegionId(v string) *SourceRocketMQParameters
	GetRegionId() *string
	SetTag(v string) *SourceRocketMQParameters
	GetTag() *string
	SetTimestamp(v int32) *SourceRocketMQParameters
	GetTimestamp() *int32
	SetTopic(v string) *SourceRocketMQParameters
	GetTopic() *string
}

type SourceRocketMQParameters struct {
	// example:
	//
	// ACL
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// example:
	//
	// Tag
	FilterType *string `json:"FilterType,omitempty" xml:"FilterType,omitempty"`
	// example:
	//
	// GID_group1
	GroupID *string `json:"GroupID,omitempty" xml:"GroupID,omitempty"`
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com
	InstanceEndpoint *string `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty"`
	// example:
	//
	// MQ_INST_164901546557****_BAAN****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// PrivateNetwork
	InstanceNetwork *string `json:"InstanceNetwork,omitempty" xml:"InstanceNetwork,omitempty"`
	// example:
	//
	// 123
	InstancePassword *string `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty"`
	// example:
	//
	// sg-hp35r2hc3a3sv8q2****
	InstanceSecurityGroupId *string `json:"InstanceSecurityGroupId,omitempty" xml:"InstanceSecurityGroupId,omitempty"`
	// example:
	//
	// Cloud_5
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 6W0xz2uPfiwp****
	InstanceUsername *string `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty"`
	// example:
	//
	// vsw-uf6gwtbn6etadpvz7****
	InstanceVSwitchIds *string `json:"InstanceVSwitchIds,omitempty" xml:"InstanceVSwitchIds,omitempty"`
	// example:
	//
	// vpc-uf6of9452b2pba82c****
	InstanceVpcId *string `json:"InstanceVpcId,omitempty" xml:"InstanceVpcId,omitempty"`
	// example:
	//
	// CONSUME_FROM_TIMESTAMP
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// test
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// example:
	//
	// 1636597951964
	Timestamp *int32 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// myTopic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s SourceRocketMQParameters) String() string {
	return dara.Prettify(s)
}

func (s SourceRocketMQParameters) GoString() string {
	return s.String()
}

func (s *SourceRocketMQParameters) GetAuthType() *string {
	return s.AuthType
}

func (s *SourceRocketMQParameters) GetFilterType() *string {
	return s.FilterType
}

func (s *SourceRocketMQParameters) GetGroupID() *string {
	return s.GroupID
}

func (s *SourceRocketMQParameters) GetInstanceEndpoint() *string {
	return s.InstanceEndpoint
}

func (s *SourceRocketMQParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SourceRocketMQParameters) GetInstanceNetwork() *string {
	return s.InstanceNetwork
}

func (s *SourceRocketMQParameters) GetInstancePassword() *string {
	return s.InstancePassword
}

func (s *SourceRocketMQParameters) GetInstanceSecurityGroupId() *string {
	return s.InstanceSecurityGroupId
}

func (s *SourceRocketMQParameters) GetInstanceType() *string {
	return s.InstanceType
}

func (s *SourceRocketMQParameters) GetInstanceUsername() *string {
	return s.InstanceUsername
}

func (s *SourceRocketMQParameters) GetInstanceVSwitchIds() *string {
	return s.InstanceVSwitchIds
}

func (s *SourceRocketMQParameters) GetInstanceVpcId() *string {
	return s.InstanceVpcId
}

func (s *SourceRocketMQParameters) GetOffset() *string {
	return s.Offset
}

func (s *SourceRocketMQParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *SourceRocketMQParameters) GetTag() *string {
	return s.Tag
}

func (s *SourceRocketMQParameters) GetTimestamp() *int32 {
	return s.Timestamp
}

func (s *SourceRocketMQParameters) GetTopic() *string {
	return s.Topic
}

func (s *SourceRocketMQParameters) SetAuthType(v string) *SourceRocketMQParameters {
	s.AuthType = &v
	return s
}

func (s *SourceRocketMQParameters) SetFilterType(v string) *SourceRocketMQParameters {
	s.FilterType = &v
	return s
}

func (s *SourceRocketMQParameters) SetGroupID(v string) *SourceRocketMQParameters {
	s.GroupID = &v
	return s
}

func (s *SourceRocketMQParameters) SetInstanceEndpoint(v string) *SourceRocketMQParameters {
	s.InstanceEndpoint = &v
	return s
}

func (s *SourceRocketMQParameters) SetInstanceId(v string) *SourceRocketMQParameters {
	s.InstanceId = &v
	return s
}

func (s *SourceRocketMQParameters) SetInstanceNetwork(v string) *SourceRocketMQParameters {
	s.InstanceNetwork = &v
	return s
}

func (s *SourceRocketMQParameters) SetInstancePassword(v string) *SourceRocketMQParameters {
	s.InstancePassword = &v
	return s
}

func (s *SourceRocketMQParameters) SetInstanceSecurityGroupId(v string) *SourceRocketMQParameters {
	s.InstanceSecurityGroupId = &v
	return s
}

func (s *SourceRocketMQParameters) SetInstanceType(v string) *SourceRocketMQParameters {
	s.InstanceType = &v
	return s
}

func (s *SourceRocketMQParameters) SetInstanceUsername(v string) *SourceRocketMQParameters {
	s.InstanceUsername = &v
	return s
}

func (s *SourceRocketMQParameters) SetInstanceVSwitchIds(v string) *SourceRocketMQParameters {
	s.InstanceVSwitchIds = &v
	return s
}

func (s *SourceRocketMQParameters) SetInstanceVpcId(v string) *SourceRocketMQParameters {
	s.InstanceVpcId = &v
	return s
}

func (s *SourceRocketMQParameters) SetOffset(v string) *SourceRocketMQParameters {
	s.Offset = &v
	return s
}

func (s *SourceRocketMQParameters) SetRegionId(v string) *SourceRocketMQParameters {
	s.RegionId = &v
	return s
}

func (s *SourceRocketMQParameters) SetTag(v string) *SourceRocketMQParameters {
	s.Tag = &v
	return s
}

func (s *SourceRocketMQParameters) SetTimestamp(v int32) *SourceRocketMQParameters {
	s.Timestamp = &v
	return s
}

func (s *SourceRocketMQParameters) SetTopic(v string) *SourceRocketMQParameters {
	s.Topic = &v
	return s
}

func (s *SourceRocketMQParameters) Validate() error {
	return dara.Validate(s)
}
