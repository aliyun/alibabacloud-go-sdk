// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSourceRabbitMQMetaParameters interface {
	dara.Model
	String() string
	GoString() string
	SetEndpoint(v string) *SourceRabbitMQMetaParameters
	GetEndpoint() *string
	SetInnerNameSuffix(v string) *SourceRabbitMQMetaParameters
	GetInnerNameSuffix() *string
	SetInstanceId(v string) *SourceRabbitMQMetaParameters
	GetInstanceId() *string
	SetInstanceType(v string) *SourceRabbitMQMetaParameters
	GetInstanceType() *string
	SetMaxHops(v string) *SourceRabbitMQMetaParameters
	GetMaxHops() *string
	SetNetworkType(v string) *SourceRabbitMQMetaParameters
	GetNetworkType() *string
	SetOrderConsume(v string) *SourceRabbitMQMetaParameters
	GetOrderConsume() *string
	SetPassword(v string) *SourceRabbitMQMetaParameters
	GetPassword() *string
	SetRegex(v string) *SourceRabbitMQMetaParameters
	GetRegex() *string
	SetSecurityGroupId(v string) *SourceRabbitMQMetaParameters
	GetSecurityGroupId() *string
	SetSyncMetaType(v string) *SourceRabbitMQMetaParameters
	GetSyncMetaType() *string
	SetUsername(v string) *SourceRabbitMQMetaParameters
	GetUsername() *string
	SetVSwitchIds(v string) *SourceRabbitMQMetaParameters
	GetVSwitchIds() *string
	SetVirtualHostName(v string) *SourceRabbitMQMetaParameters
	GetVirtualHostName() *string
	SetVpcId(v string) *SourceRabbitMQMetaParameters
	GetVpcId() *string
}

type SourceRabbitMQMetaParameters struct {
	// The access endpoint of the RabbitMQ cluster.
	//
	// example:
	//
	// rabbitmq-cn-a1b2c3d4e5f-cn-hangzhou-vpc.amqp.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The added suffix.
	//
	// example:
	//
	// rabbitmq-cn-a1b2c3d4e5f_example_123e4567-e89b-42d3-a456-426614174000
	InnerNameSuffix *string `json:"InnerNameSuffix,omitempty" xml:"InnerNameSuffix,omitempty"`
	// The ID of the ApsaraMQ for RabbitMQ (MSMQ message queue) instance.
	//
	// example:
	//
	// rabbitmq-cn-a1b2c3d4e5f
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the source RabbitMQ instance.
	//
	// example:
	//
	// Cloud
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The maximum number of hops allowed for message forwarding in a federation task. Messages that exceed this limit are dropped.
	//
	// example:
	//
	// 1
	MaxHops *string `json:"MaxHops,omitempty" xml:"MaxHops,omitempty"`
	// The network type.
	//
	// example:
	//
	// PrivateNetwork
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// Specifies whether to consume messages in order.
	//
	// example:
	//
	// false
	OrderConsume *string `json:"OrderConsume,omitempty" xml:"OrderConsume,omitempty"`
	// The password used to connect to the RabbitMQ cluster.
	//
	// example:
	//
	// yourPassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The regular expression used to filter exchanges.
	//
	// example:
	//
	// ^example_.*
	Regex *string `json:"Regex,omitempty" xml:"Regex,omitempty"`
	// The ID of the security group to which the machines of the RabbitMQ cluster belong.
	//
	// example:
	//
	// sg-bp1****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The type of synchronized metadata.
	//
	// example:
	//
	// queue
	SyncMetaType *string `json:"SyncMetaType,omitempty" xml:"SyncMetaType,omitempty"`
	// The username used to connect to the RabbitMQ cluster.
	//
	// example:
	//
	// admin
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The ID of the vSwitch in the VPC where the RabbitMQ cluster resides.
	//
	// example:
	//
	// vsw-bp1****
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The name of the virtual host (vhost) in the RabbitMQ cluster.
	//
	// example:
	//
	// test-001
	VirtualHostName *string `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
	// The ID of the VPC where the RabbitMQ cluster resides.
	//
	// example:
	//
	// vpc-8vb****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s SourceRabbitMQMetaParameters) String() string {
	return dara.Prettify(s)
}

func (s SourceRabbitMQMetaParameters) GoString() string {
	return s.String()
}

func (s *SourceRabbitMQMetaParameters) GetEndpoint() *string {
	return s.Endpoint
}

func (s *SourceRabbitMQMetaParameters) GetInnerNameSuffix() *string {
	return s.InnerNameSuffix
}

func (s *SourceRabbitMQMetaParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SourceRabbitMQMetaParameters) GetInstanceType() *string {
	return s.InstanceType
}

func (s *SourceRabbitMQMetaParameters) GetMaxHops() *string {
	return s.MaxHops
}

func (s *SourceRabbitMQMetaParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *SourceRabbitMQMetaParameters) GetOrderConsume() *string {
	return s.OrderConsume
}

func (s *SourceRabbitMQMetaParameters) GetPassword() *string {
	return s.Password
}

func (s *SourceRabbitMQMetaParameters) GetRegex() *string {
	return s.Regex
}

func (s *SourceRabbitMQMetaParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *SourceRabbitMQMetaParameters) GetSyncMetaType() *string {
	return s.SyncMetaType
}

func (s *SourceRabbitMQMetaParameters) GetUsername() *string {
	return s.Username
}

func (s *SourceRabbitMQMetaParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *SourceRabbitMQMetaParameters) GetVirtualHostName() *string {
	return s.VirtualHostName
}

func (s *SourceRabbitMQMetaParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *SourceRabbitMQMetaParameters) SetEndpoint(v string) *SourceRabbitMQMetaParameters {
	s.Endpoint = &v
	return s
}

func (s *SourceRabbitMQMetaParameters) SetInnerNameSuffix(v string) *SourceRabbitMQMetaParameters {
	s.InnerNameSuffix = &v
	return s
}

func (s *SourceRabbitMQMetaParameters) SetInstanceId(v string) *SourceRabbitMQMetaParameters {
	s.InstanceId = &v
	return s
}

func (s *SourceRabbitMQMetaParameters) SetInstanceType(v string) *SourceRabbitMQMetaParameters {
	s.InstanceType = &v
	return s
}

func (s *SourceRabbitMQMetaParameters) SetMaxHops(v string) *SourceRabbitMQMetaParameters {
	s.MaxHops = &v
	return s
}

func (s *SourceRabbitMQMetaParameters) SetNetworkType(v string) *SourceRabbitMQMetaParameters {
	s.NetworkType = &v
	return s
}

func (s *SourceRabbitMQMetaParameters) SetOrderConsume(v string) *SourceRabbitMQMetaParameters {
	s.OrderConsume = &v
	return s
}

func (s *SourceRabbitMQMetaParameters) SetPassword(v string) *SourceRabbitMQMetaParameters {
	s.Password = &v
	return s
}

func (s *SourceRabbitMQMetaParameters) SetRegex(v string) *SourceRabbitMQMetaParameters {
	s.Regex = &v
	return s
}

func (s *SourceRabbitMQMetaParameters) SetSecurityGroupId(v string) *SourceRabbitMQMetaParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *SourceRabbitMQMetaParameters) SetSyncMetaType(v string) *SourceRabbitMQMetaParameters {
	s.SyncMetaType = &v
	return s
}

func (s *SourceRabbitMQMetaParameters) SetUsername(v string) *SourceRabbitMQMetaParameters {
	s.Username = &v
	return s
}

func (s *SourceRabbitMQMetaParameters) SetVSwitchIds(v string) *SourceRabbitMQMetaParameters {
	s.VSwitchIds = &v
	return s
}

func (s *SourceRabbitMQMetaParameters) SetVirtualHostName(v string) *SourceRabbitMQMetaParameters {
	s.VirtualHostName = &v
	return s
}

func (s *SourceRabbitMQMetaParameters) SetVpcId(v string) *SourceRabbitMQMetaParameters {
	s.VpcId = &v
	return s
}

func (s *SourceRabbitMQMetaParameters) Validate() error {
	return dara.Validate(s)
}
