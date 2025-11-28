// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSinkRabbitMQMetaParameters interface {
	dara.Model
	String() string
	GoString() string
	SetEndpoint(v string) *SinkRabbitMQMetaParameters
	GetEndpoint() *string
	SetInstanceId(v string) *SinkRabbitMQMetaParameters
	GetInstanceId() *string
	SetInstanceType(v string) *SinkRabbitMQMetaParameters
	GetInstanceType() *string
	SetNetworkType(v string) *SinkRabbitMQMetaParameters
	GetNetworkType() *string
	SetPassword(v string) *SinkRabbitMQMetaParameters
	GetPassword() *string
	SetSecurityGroupId(v string) *SinkRabbitMQMetaParameters
	GetSecurityGroupId() *string
	SetUsername(v string) *SinkRabbitMQMetaParameters
	GetUsername() *string
	SetVSwitchIds(v string) *SinkRabbitMQMetaParameters
	GetVSwitchIds() *string
	SetVirtualHostName(v string) *SinkRabbitMQMetaParameters
	GetVirtualHostName() *string
	SetVpcId(v string) *SinkRabbitMQMetaParameters
	GetVpcId() *string
}

type SinkRabbitMQMetaParameters struct {
	Endpoint        *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType    *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	NetworkType     *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Username        *string `json:"Username,omitempty" xml:"Username,omitempty"`
	VSwitchIds      *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	VirtualHostName *string `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s SinkRabbitMQMetaParameters) String() string {
	return dara.Prettify(s)
}

func (s SinkRabbitMQMetaParameters) GoString() string {
	return s.String()
}

func (s *SinkRabbitMQMetaParameters) GetEndpoint() *string {
	return s.Endpoint
}

func (s *SinkRabbitMQMetaParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SinkRabbitMQMetaParameters) GetInstanceType() *string {
	return s.InstanceType
}

func (s *SinkRabbitMQMetaParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *SinkRabbitMQMetaParameters) GetPassword() *string {
	return s.Password
}

func (s *SinkRabbitMQMetaParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *SinkRabbitMQMetaParameters) GetUsername() *string {
	return s.Username
}

func (s *SinkRabbitMQMetaParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *SinkRabbitMQMetaParameters) GetVirtualHostName() *string {
	return s.VirtualHostName
}

func (s *SinkRabbitMQMetaParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *SinkRabbitMQMetaParameters) SetEndpoint(v string) *SinkRabbitMQMetaParameters {
	s.Endpoint = &v
	return s
}

func (s *SinkRabbitMQMetaParameters) SetInstanceId(v string) *SinkRabbitMQMetaParameters {
	s.InstanceId = &v
	return s
}

func (s *SinkRabbitMQMetaParameters) SetInstanceType(v string) *SinkRabbitMQMetaParameters {
	s.InstanceType = &v
	return s
}

func (s *SinkRabbitMQMetaParameters) SetNetworkType(v string) *SinkRabbitMQMetaParameters {
	s.NetworkType = &v
	return s
}

func (s *SinkRabbitMQMetaParameters) SetPassword(v string) *SinkRabbitMQMetaParameters {
	s.Password = &v
	return s
}

func (s *SinkRabbitMQMetaParameters) SetSecurityGroupId(v string) *SinkRabbitMQMetaParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *SinkRabbitMQMetaParameters) SetUsername(v string) *SinkRabbitMQMetaParameters {
	s.Username = &v
	return s
}

func (s *SinkRabbitMQMetaParameters) SetVSwitchIds(v string) *SinkRabbitMQMetaParameters {
	s.VSwitchIds = &v
	return s
}

func (s *SinkRabbitMQMetaParameters) SetVirtualHostName(v string) *SinkRabbitMQMetaParameters {
	s.VirtualHostName = &v
	return s
}

func (s *SinkRabbitMQMetaParameters) SetVpcId(v string) *SinkRabbitMQMetaParameters {
	s.VpcId = &v
	return s
}

func (s *SinkRabbitMQMetaParameters) Validate() error {
	return dara.Validate(s)
}
