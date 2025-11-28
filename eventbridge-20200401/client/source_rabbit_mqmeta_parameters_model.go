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
	Endpoint        *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	InnerNameSuffix *string `json:"InnerNameSuffix,omitempty" xml:"InnerNameSuffix,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType    *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	MaxHops         *string `json:"MaxHops,omitempty" xml:"MaxHops,omitempty"`
	NetworkType     *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OrderConsume    *string `json:"OrderConsume,omitempty" xml:"OrderConsume,omitempty"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Regex           *string `json:"Regex,omitempty" xml:"Regex,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Username        *string `json:"Username,omitempty" xml:"Username,omitempty"`
	VSwitchIds      *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	VirtualHostName *string `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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
